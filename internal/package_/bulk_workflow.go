package package_

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/oklog/run"
	temporalsdk_activity "go.temporal.io/sdk/activity"
	temporalsdk_temporal "go.temporal.io/sdk/temporal"
	temporalsdk_workflow "go.temporal.io/sdk/workflow"

	"github.com/artefactual-labs/enduro/internal/api/gen/package_"
)

const (
	BulkWorkflowName              = "package-bulk-workflow"
	BulkWorkflowID                = "package-bulk-workflow"
	BulkWorkflowStateQueryHandler = "package-bulk-state"
	BulkActivityName              = "package-bulk-activity"
)

// BulkProgress reports bulk operation status - delivered as heartbeats.
type BulkProgress struct {
	CurrentID uint
	Count     uint
	Max       uint
}

type BulkWorkflowOperation string

const (
	BulkWorkflowOperationRetry   BulkWorkflowOperation = "retry"
	BulkWorkflowOperationCancel  BulkWorkflowOperation = "cancel"
	BulkWorkflowOperationAbandon BulkWorkflowOperation = "abandon"
)

type BulkWorkflowInput struct {
	// Status of packages where bulk is performed.
	Status Status

	// Type of operation that is performed, e.g. "retry", "cancel"...
	Operation BulkWorkflowOperation

	// Max. number of packages affected. Zero means no cap established.
	Size uint
}

// BulkWorkflow is a Temporal workflow that performs bulk operations.
func BulkWorkflow(ctx temporalsdk_workflow.Context, params BulkWorkflowInput) error {
	opts := temporalsdk_workflow.WithActivityOptions(ctx, temporalsdk_workflow.ActivityOptions{
		StartToCloseTimeout: time.Hour * 24 * 365,
		WaitForCancellation: true,
		HeartbeatTimeout:    time.Second * 5,
		RetryPolicy: &temporalsdk_temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	})

	return temporalsdk_workflow.ExecuteActivity(opts, BulkActivityName, params).Get(opts, nil)
}

type BulkActivity struct {
	pkgsvc Service
}

func NewBulkActivity(pkgsvc Service) *BulkActivity {
	return &BulkActivity{
		pkgsvc: pkgsvc,
	}
}

func (a *BulkActivity) Execute(ctx context.Context, params BulkWorkflowInput) error {
	var group run.Group

	// One actor does the work while updating progress.
	// The other one sends the heartbeats.
	progress := &BulkProgress{}
	var mu sync.RWMutex

	{
		cancel := make(chan struct{})

		group.Add(
			func() error {
				ticker := time.NewTicker(time.Second * 1)
				defer ticker.Stop()
				for {
					select {
					case <-ctx.Done():
						return ctx.Err()
					case <-cancel:
						return nil
					case <-ticker.C:
						mu.RLock()
						cp := progress
						mu.RUnlock()

						temporalsdk_activity.RecordHeartbeat(ctx, cp)
					}
				}
			},
			func(error) {
				close(cancel)
			},
		)
	}

	{
		cancel := make(chan struct{})

		group.Add(
			func() error {
				var nextCursor *string
				status := params.Status.String()
				var count uint
				for {
					select {
					case <-cancel:
						return nil
					default:
						res, err := a.pkgsvc.Goa().List(ctx, &package_.ListPayload{
							Status: &status,
							Cursor: nextCursor,
						})
						if err != nil {
							return err
						}

						for _, item := range res.Items {
							mu.Lock()
							progress = &BulkProgress{
								CurrentID: item.ID,
								Count:     count + 1,
								Max:       params.Size,
							}
							mu.Unlock()

							switch params.Operation {
							case BulkWorkflowOperationRetry:
								err = a.Retry(ctx, item.ID)
							default:
								return fmt.Errorf("bulk %s not supported yet", params.Operation)
							}
							if err != nil {
								return fmt.Errorf("error executing bulk %s (failed on package %d): %v", params.Operation, item.ID, err)
							}

							// Stop when cap is reached.
							count++
							if params.Size > 0 && count == params.Size {
								return nil
							}

							// Making sure that we don't overwork the system.
							time.Sleep(time.Millisecond * 50)
						}

						nextCursor = res.NextCursor
						if nextCursor == nil {
							return nil
						}
					}
				}
			},
			func(error) {
				close(cancel)
			},
		)
	}

	return group.Run()
}

func (a *BulkActivity) Retry(ctx context.Context, ID uint) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	err := a.pkgsvc.Goa().Retry(ctx, &package_.RetryPayload{ID: ID})

	// User may have already started it manually.
	if temporalsdk_temporal.IsWorkflowExecutionAlreadyStartedError(err) {
		return nil
	}

	// TODO: ignore the error returned when the workflow history does not exist,
	// which is something that we used to do in Cadence.

	return err
}
