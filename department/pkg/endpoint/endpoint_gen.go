// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "mgo/department/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetEndpoint     endpoint.Endpoint
	AddEndpoint     endpoint.Endpoint
	DeleteEndpoint  endpoint.Endpoint
	GetByIDEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.DepartmentService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		AddEndpoint:     MakeAddEndpoint(s),
		DeleteEndpoint:  MakeDeleteEndpoint(s),
		GetByIDEndpoint: MakeGetByIDEndpoint(s),
		GetEndpoint:     MakeGetEndpoint(s),
	}
	for _, m := range mdw["Get"] {
		eps.GetEndpoint = m(eps.GetEndpoint)
	}
	for _, m := range mdw["Add"] {
		eps.AddEndpoint = m(eps.AddEndpoint)
	}
	for _, m := range mdw["Delete"] {
		eps.DeleteEndpoint = m(eps.DeleteEndpoint)
	}
	for _, m := range mdw["GetByID"] {
		eps.GetByIDEndpoint = m(eps.GetByIDEndpoint)
	}
	return eps
}
