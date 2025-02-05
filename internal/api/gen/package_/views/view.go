// Code generated by goa v3.7.7, DO NOT EDIT.
//
// package views
//
// Command:
// $ goa-v3.7.7 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// EnduroMonitorEvent is the viewed result type that is projected based on a
// view.
type EnduroMonitorEvent struct {
	// Type to project
	Projected *EnduroMonitorEventView
	// View to render
	View string
}

// EnduroStoredPackage is the viewed result type that is projected based on a
// view.
type EnduroStoredPackage struct {
	// Type to project
	Projected *EnduroStoredPackageView
	// View to render
	View string
}

// EnduroPackageWorkflowStatus is the viewed result type that is projected
// based on a view.
type EnduroPackageWorkflowStatus struct {
	// Type to project
	Projected *EnduroPackageWorkflowStatusView
	// View to render
	View string
}

// EnduroPackagePreservationActions is the viewed result type that is projected
// based on a view.
type EnduroPackagePreservationActions struct {
	// Type to project
	Projected *EnduroPackagePreservationActionsView
	// View to render
	View string
}

// EnduroMonitorEventView is a type that runs validations on a projected type.
type EnduroMonitorEventView struct {
	MonitorPingEvent            *EnduroMonitorPingEventView
	PackageCreatedEvent         *EnduroPackageCreatedEventView
	PackageDeletedEvent         *EnduroPackageDeletedEventView
	PackageUpdatedEvent         *EnduroPackageUpdatedEventView
	PackageStatusUpdatedEvent   *EnduroPackageStatusUpdatedEventView
	PackageLocationUpdatedEvent *EnduroPackageLocationUpdatedEventView
}

// EnduroMonitorPingEventView is a type that runs validations on a projected
// type.
type EnduroMonitorPingEventView struct {
	Message *string
}

// EnduroPackageCreatedEventView is a type that runs validations on a projected
// type.
type EnduroPackageCreatedEventView struct {
	// Identifier of package
	ID   *uint
	Item *EnduroStoredPackageView
}

// EnduroStoredPackageView is a type that runs validations on a projected type.
type EnduroStoredPackageView struct {
	// Identifier of package
	ID *uint
	// Name of the package
	Name *string
	// Location of the package
	Location *string
	// Status of the package
	Status *string
	// Identifier of processing workflow
	WorkflowID *string
	// Identifier of latest processing workflow run
	RunID *string
	// Identifier of Archivematica AIP
	AipID *string
	// Creation datetime
	CreatedAt *string
	// Start datetime
	StartedAt *string
	// Completion datetime
	CompletedAt *string
}

// EnduroPackageDeletedEventView is a type that runs validations on a projected
// type.
type EnduroPackageDeletedEventView struct {
	// Identifier of package
	ID *uint
}

// EnduroPackageUpdatedEventView is a type that runs validations on a projected
// type.
type EnduroPackageUpdatedEventView struct {
	// Identifier of package
	ID   *uint
	Item *EnduroStoredPackageView
}

// EnduroPackageStatusUpdatedEventView is a type that runs validations on a
// projected type.
type EnduroPackageStatusUpdatedEventView struct {
	// Identifier of package
	ID     *uint
	Status *string
}

// EnduroPackageLocationUpdatedEventView is a type that runs validations on a
// projected type.
type EnduroPackageLocationUpdatedEventView struct {
	// Identifier of package
	ID       *uint
	Location *string
}

// EnduroPackageWorkflowStatusView is a type that runs validations on a
// projected type.
type EnduroPackageWorkflowStatusView struct {
	Status  *string
	History EnduroPackageWorkflowHistoryCollectionView
}

// EnduroPackageWorkflowHistoryCollectionView is a type that runs validations
// on a projected type.
type EnduroPackageWorkflowHistoryCollectionView []*EnduroPackageWorkflowHistoryView

// EnduroPackageWorkflowHistoryView is a type that runs validations on a
// projected type.
type EnduroPackageWorkflowHistoryView struct {
	// Identifier of package
	ID *uint
	// Type of the event
	Type *string
	// Contents of the event
	Details interface{}
}

// EnduroPackagePreservationActionsView is a type that runs validations on a
// projected type.
type EnduroPackagePreservationActionsView struct {
	Actions EnduroPackagePreservationActionsActionCollectionView
}

// EnduroPackagePreservationActionsActionCollectionView is a type that runs
// validations on a projected type.
type EnduroPackagePreservationActionsActionCollectionView []*EnduroPackagePreservationActionsActionView

// EnduroPackagePreservationActionsActionView is a type that runs validations
// on a projected type.
type EnduroPackagePreservationActionsActionView struct {
	ID          *uint
	ActionID    *string
	Name        *string
	Status      *string
	StartedAt   *string
	CompletedAt *string
}

var (
	// EnduroMonitorEventMap is a map indexing the attribute names of
	// EnduroMonitorEvent by view name.
	EnduroMonitorEventMap = map[string][]string{
		"default": {
			"monitor_ping_event",
			"package_created_event",
			"package_deleted_event",
			"package_updated_event",
			"package_status_updated_event",
			"package_location_updated_event",
		},
	}
	// EnduroStoredPackageMap is a map indexing the attribute names of
	// EnduroStoredPackage by view name.
	EnduroStoredPackageMap = map[string][]string{
		"default": {
			"id",
			"name",
			"location",
			"status",
			"workflow_id",
			"run_id",
			"aip_id",
			"created_at",
			"started_at",
			"completed_at",
		},
	}
	// EnduroPackageWorkflowStatusMap is a map indexing the attribute names of
	// EnduroPackageWorkflowStatus by view name.
	EnduroPackageWorkflowStatusMap = map[string][]string{
		"default": {
			"status",
			"history",
		},
	}
	// EnduroPackagePreservationActionsMap is a map indexing the attribute names of
	// EnduroPackagePreservationActions by view name.
	EnduroPackagePreservationActionsMap = map[string][]string{
		"default": {
			"actions",
		},
	}
	// EnduroMonitorPingEventMap is a map indexing the attribute names of
	// EnduroMonitorPingEvent by view name.
	EnduroMonitorPingEventMap = map[string][]string{
		"default": {
			"message",
		},
	}
	// EnduroPackageCreatedEventMap is a map indexing the attribute names of
	// EnduroPackageCreatedEvent by view name.
	EnduroPackageCreatedEventMap = map[string][]string{
		"default": {
			"id",
			"item",
		},
	}
	// EnduroPackageDeletedEventMap is a map indexing the attribute names of
	// EnduroPackageDeletedEvent by view name.
	EnduroPackageDeletedEventMap = map[string][]string{
		"default": {
			"id",
		},
	}
	// EnduroPackageUpdatedEventMap is a map indexing the attribute names of
	// EnduroPackageUpdatedEvent by view name.
	EnduroPackageUpdatedEventMap = map[string][]string{
		"default": {
			"id",
			"item",
		},
	}
	// EnduroPackageStatusUpdatedEventMap is a map indexing the attribute names of
	// EnduroPackageStatusUpdatedEvent by view name.
	EnduroPackageStatusUpdatedEventMap = map[string][]string{
		"default": {
			"id",
			"status",
		},
	}
	// EnduroPackageLocationUpdatedEventMap is a map indexing the attribute names
	// of EnduroPackageLocationUpdatedEvent by view name.
	EnduroPackageLocationUpdatedEventMap = map[string][]string{
		"default": {
			"id",
			"location",
		},
	}
	// EnduroPackageWorkflowHistoryCollectionMap is a map indexing the attribute
	// names of EnduroPackageWorkflowHistoryCollection by view name.
	EnduroPackageWorkflowHistoryCollectionMap = map[string][]string{
		"default": {
			"id",
			"type",
			"details",
		},
	}
	// EnduroPackageWorkflowHistoryMap is a map indexing the attribute names of
	// EnduroPackageWorkflowHistory by view name.
	EnduroPackageWorkflowHistoryMap = map[string][]string{
		"default": {
			"id",
			"type",
			"details",
		},
	}
	// EnduroPackagePreservationActionsActionCollectionMap is a map indexing the
	// attribute names of EnduroPackagePreservationActionsActionCollection by view
	// name.
	EnduroPackagePreservationActionsActionCollectionMap = map[string][]string{
		"default": {
			"id",
			"action_id",
			"name",
			"status",
			"started_at",
			"completed_at",
		},
	}
	// EnduroPackagePreservationActionsActionMap is a map indexing the attribute
	// names of EnduroPackagePreservationActionsAction by view name.
	EnduroPackagePreservationActionsActionMap = map[string][]string{
		"default": {
			"id",
			"action_id",
			"name",
			"status",
			"started_at",
			"completed_at",
		},
	}
)

// ValidateEnduroMonitorEvent runs the validations defined on the viewed result
// type EnduroMonitorEvent.
func ValidateEnduroMonitorEvent(result *EnduroMonitorEvent) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateEnduroMonitorEventView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateEnduroStoredPackage runs the validations defined on the viewed
// result type EnduroStoredPackage.
func ValidateEnduroStoredPackage(result *EnduroStoredPackage) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateEnduroStoredPackageView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateEnduroPackageWorkflowStatus runs the validations defined on the
// viewed result type EnduroPackageWorkflowStatus.
func ValidateEnduroPackageWorkflowStatus(result *EnduroPackageWorkflowStatus) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateEnduroPackageWorkflowStatusView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateEnduroPackagePreservationActions runs the validations defined on the
// viewed result type EnduroPackagePreservationActions.
func ValidateEnduroPackagePreservationActions(result *EnduroPackagePreservationActions) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateEnduroPackagePreservationActionsView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateEnduroMonitorEventView runs the validations defined on
// EnduroMonitorEventView using the "default" view.
func ValidateEnduroMonitorEventView(result *EnduroMonitorEventView) (err error) {

	if result.MonitorPingEvent != nil {
		if err2 := ValidateEnduroMonitorPingEventView(result.MonitorPingEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.PackageCreatedEvent != nil {
		if err2 := ValidateEnduroPackageCreatedEventView(result.PackageCreatedEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.PackageDeletedEvent != nil {
		if err2 := ValidateEnduroPackageDeletedEventView(result.PackageDeletedEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.PackageUpdatedEvent != nil {
		if err2 := ValidateEnduroPackageUpdatedEventView(result.PackageUpdatedEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.PackageStatusUpdatedEvent != nil {
		if err2 := ValidateEnduroPackageStatusUpdatedEventView(result.PackageStatusUpdatedEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.PackageLocationUpdatedEvent != nil {
		if err2 := ValidateEnduroPackageLocationUpdatedEventView(result.PackageLocationUpdatedEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroMonitorPingEventView runs the validations defined on
// EnduroMonitorPingEventView using the "default" view.
func ValidateEnduroMonitorPingEventView(result *EnduroMonitorPingEventView) (err error) {

	return
}

// ValidateEnduroPackageCreatedEventView runs the validations defined on
// EnduroPackageCreatedEventView using the "default" view.
func ValidateEnduroPackageCreatedEventView(result *EnduroPackageCreatedEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Item != nil {
		if err2 := ValidateEnduroStoredPackageView(result.Item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroStoredPackageView runs the validations defined on
// EnduroStoredPackageView using the "default" view.
func ValidateEnduroStoredPackageView(result *EnduroStoredPackageView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "new" || *result.Status == "in progress" || *result.Status == "done" || *result.Status == "error" || *result.Status == "unknown" || *result.Status == "queued" || *result.Status == "pending" || *result.Status == "abandoned") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"new", "in progress", "done", "error", "unknown", "queued", "pending", "abandoned"}))
		}
	}
	if result.WorkflowID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.workflow_id", *result.WorkflowID, goa.FormatUUID))
	}
	if result.RunID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.run_id", *result.RunID, goa.FormatUUID))
	}
	if result.AipID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.aip_id", *result.AipID, goa.FormatUUID))
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started_at", *result.StartedAt, goa.FormatDateTime))
	}
	if result.CompletedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.completed_at", *result.CompletedAt, goa.FormatDateTime))
	}
	return
}

// ValidateEnduroPackageDeletedEventView runs the validations defined on
// EnduroPackageDeletedEventView using the "default" view.
func ValidateEnduroPackageDeletedEventView(result *EnduroPackageDeletedEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	return
}

// ValidateEnduroPackageUpdatedEventView runs the validations defined on
// EnduroPackageUpdatedEventView using the "default" view.
func ValidateEnduroPackageUpdatedEventView(result *EnduroPackageUpdatedEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Item != nil {
		if err2 := ValidateEnduroStoredPackageView(result.Item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackageStatusUpdatedEventView runs the validations defined on
// EnduroPackageStatusUpdatedEventView using the "default" view.
func ValidateEnduroPackageStatusUpdatedEventView(result *EnduroPackageStatusUpdatedEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "new" || *result.Status == "in progress" || *result.Status == "done" || *result.Status == "error" || *result.Status == "unknown" || *result.Status == "queued" || *result.Status == "pending" || *result.Status == "abandoned") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"new", "in progress", "done", "error", "unknown", "queued", "pending", "abandoned"}))
		}
	}
	return
}

// ValidateEnduroPackageLocationUpdatedEventView runs the validations defined
// on EnduroPackageLocationUpdatedEventView using the "default" view.
func ValidateEnduroPackageLocationUpdatedEventView(result *EnduroPackageLocationUpdatedEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Location == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("location", "result"))
	}
	return
}

// ValidateEnduroPackageWorkflowStatusView runs the validations defined on
// EnduroPackageWorkflowStatusView using the "default" view.
func ValidateEnduroPackageWorkflowStatusView(result *EnduroPackageWorkflowStatusView) (err error) {

	if result.History != nil {
		if err2 := ValidateEnduroPackageWorkflowHistoryCollectionView(result.History); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackageWorkflowHistoryCollectionView runs the validations
// defined on EnduroPackageWorkflowHistoryCollectionView using the "default"
// view.
func ValidateEnduroPackageWorkflowHistoryCollectionView(result EnduroPackageWorkflowHistoryCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateEnduroPackageWorkflowHistoryView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackageWorkflowHistoryView runs the validations defined on
// EnduroPackageWorkflowHistoryView using the "default" view.
func ValidateEnduroPackageWorkflowHistoryView(result *EnduroPackageWorkflowHistoryView) (err error) {

	return
}

// ValidateEnduroPackagePreservationActionsView runs the validations defined on
// EnduroPackagePreservationActionsView using the "default" view.
func ValidateEnduroPackagePreservationActionsView(result *EnduroPackagePreservationActionsView) (err error) {

	if result.Actions != nil {
		if err2 := ValidateEnduroPackagePreservationActionsActionCollectionView(result.Actions); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackagePreservationActionsActionCollectionView runs the
// validations defined on EnduroPackagePreservationActionsActionCollectionView
// using the "default" view.
func ValidateEnduroPackagePreservationActionsActionCollectionView(result EnduroPackagePreservationActionsActionCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateEnduroPackagePreservationActionsActionView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackagePreservationActionsActionView runs the validations
// defined on EnduroPackagePreservationActionsActionView using the "default"
// view.
func ValidateEnduroPackagePreservationActionsActionView(result *EnduroPackagePreservationActionsActionView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.ActionID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("action_id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.StartedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("started_at", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "unspecified" || *result.Status == "complete" || *result.Status == "processing" || *result.Status == "failed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"unspecified", "complete", "processing", "failed"}))
		}
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started_at", *result.StartedAt, goa.FormatDateTime))
	}
	if result.CompletedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.completed_at", *result.CompletedAt, goa.FormatDateTime))
	}
	return
}
