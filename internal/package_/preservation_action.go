package package_

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	goapackage "github.com/artefactual-labs/enduro/internal/api/gen/package_"
)

type PreservationActionStatus uint

const (
	StatusUnspecified PreservationActionStatus = iota
	StatusComplete
	StatusProcessing
	StatusFailed
)

func NewPreservationActionStatus(status string) PreservationActionStatus {
	var s PreservationActionStatus

	switch strings.ToLower(status) {
	case "processing":
		s = StatusProcessing
	case "complete":
		s = StatusComplete
	case "failed":
		s = StatusFailed
	default:
		s = StatusUnspecified
	}

	return s
}

func (p PreservationActionStatus) String() string {
	switch p {
	case StatusProcessing:
		return "processing"
	case StatusComplete:
		return "complete"
	case StatusFailed:
		return "failed"
	}
	return "unspecified"
}

func (p PreservationActionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *PreservationActionStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*p = NewPreservationActionStatus(s)

	return nil
}

// PreservationAction represents a preservation action in the preservation_action table.
type PreservationAction struct {
	ID          uint                     `db:"id"`
	ActionID    string                   `db:"action_id"`
	Name        string                   `db:"name"`
	Status      PreservationActionStatus `db:"status"`
	StartedAt   sql.NullTime             `db:"started_at"`
	CompletedAt sql.NullTime             `db:"completed_at"`
	PackageID   uint                     `db:"package_id"`
}

func (w *goaWrapper) PreservationActions(ctx context.Context, payload *goapackage.PreservationActionsPayload) (*goapackage.EnduroPackagePreservationActions, error) {
	var err error
	var goacol *goapackage.EnduroStoredPackage
	if goacol, err = w.Show(ctx, &goapackage.ShowPayload{ID: payload.ID}); err != nil {
		return nil, err
	}

	query := "SELECT id, action_id, name, status, CONVERT_TZ(started_at, @@session.time_zone, '+00:00') AS started_at, CONVERT_TZ(completed_at, @@session.time_zone, '+00:00') AS completed_at FROM preservation_action WHERE package_id = (?)"
	args := []interface{}{goacol.ID}

	query = w.db.Rebind(query)
	rows, err := w.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying the database: %w", err)
	}
	defer rows.Close()

	preservation_actions := []*goapackage.EnduroPackagePreservationActionsAction{}
	for rows.Next() {
		pa := PreservationAction{}
		if err := rows.StructScan(&pa); err != nil {
			return nil, fmt.Errorf("error scanning database result: %w", err)
		}
		preservation_actions = append(preservation_actions, &goapackage.EnduroPackagePreservationActionsAction{
			ID:          pa.ID,
			ActionID:    pa.ActionID,
			Name:        pa.Name,
			Status:      pa.Status.String(),
			StartedAt:   pa.StartedAt.Time.Format(time.RFC3339),
			CompletedAt: formatOptionalTime(pa.CompletedAt),
		})
	}

	result := &goapackage.EnduroPackagePreservationActions{
		Actions: preservation_actions,
	}

	return result, nil
}

func (svc *packageImpl) CreatePreservationAction(ctx context.Context, pa *PreservationAction) error {
	query := `INSERT INTO preservation_action (action_id, name, status, started_at, completed_at, package_id) VALUES ((?), (?), (?), (?), (?), (?))`
	args := []interface{}{
		pa.ActionID,
		pa.Name,
		pa.Status,
		pa.StartedAt,
		pa.CompletedAt,
		pa.PackageID,
	}

	query = svc.db.Rebind(query)
	res, err := svc.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error inserting preservation action: %w", err)
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return fmt.Errorf("error retrieving insert ID: %w", err)
	}

	pa.ID = uint(id)

	// publishEvent(ctx, svc.events, EventTypePackageUpdated, pa.PackageID)

	return nil
}
