// Code generated by goa v3.7.6, DO NOT EDIT.
//
// storage endpoints
//
// Command:
// $ goa-v3.7.6 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package storage

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "storage" service endpoints.
type Endpoints struct {
	Submit     goa.Endpoint
	Update     goa.Endpoint
	Download   goa.Endpoint
	List       goa.Endpoint
	Move       goa.Endpoint
	MoveStatus goa.Endpoint
	Reject     goa.Endpoint
	Show       goa.Endpoint
}

// NewEndpoints wraps the methods of the "storage" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Submit:     NewSubmitEndpoint(s),
		Update:     NewUpdateEndpoint(s),
		Download:   NewDownloadEndpoint(s),
		List:       NewListEndpoint(s),
		Move:       NewMoveEndpoint(s),
		MoveStatus: NewMoveStatusEndpoint(s),
		Reject:     NewRejectEndpoint(s),
		Show:       NewShowEndpoint(s),
	}
}

// Use applies the given middleware to all the "storage" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Submit = m(e.Submit)
	e.Update = m(e.Update)
	e.Download = m(e.Download)
	e.List = m(e.List)
	e.Move = m(e.Move)
	e.MoveStatus = m(e.MoveStatus)
	e.Reject = m(e.Reject)
	e.Show = m(e.Show)
}

// NewSubmitEndpoint returns an endpoint function that calls the method
// "submit" of service "storage".
func NewSubmitEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SubmitPayload)
		return s.Submit(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "storage".
func NewUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdatePayload)
		return nil, s.Update(ctx, p)
	}
}

// NewDownloadEndpoint returns an endpoint function that calls the method
// "download" of service "storage".
func NewDownloadEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DownloadPayload)
		return s.Download(ctx, p)
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "storage".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := s.List(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredLocationCollection(res, "default")
		return vres, nil
	}
}

// NewMoveEndpoint returns an endpoint function that calls the method "move" of
// service "storage".
func NewMoveEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MovePayload)
		return nil, s.Move(ctx, p)
	}
}

// NewMoveStatusEndpoint returns an endpoint function that calls the method
// "move_status" of service "storage".
func NewMoveStatusEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MoveStatusPayload)
		return s.MoveStatus(ctx, p)
	}
}

// NewRejectEndpoint returns an endpoint function that calls the method
// "reject" of service "storage".
func NewRejectEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RejectPayload)
		return nil, s.Reject(ctx, p)
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "storage".
func NewShowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		res, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredStoragePackage(res, "default")
		return vres, nil
	}
}
