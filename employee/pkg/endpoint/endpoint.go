package endpoint

import (
	"context"
	io "mgo/employee/pkg/io"
	service "mgo/employee/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	E     []io.Employee `json:"e"`
	Error error         `json:"error"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		e, error := s.Get(ctx)
		return GetResponse{
			E:     e,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Error
}

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	Employee io.Employee `json:"employee"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	E     io.Employee `json:"e"`
	Error error       `json:"error"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		e, error := s.Add(ctx, req.Employee)
		return AddResponse{
			E:     e,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.Error
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	Id string `json:"id"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	Error error `json:"error"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		error := s.Delete(ctx, req.Id)
		return DeleteResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r DeleteResponse) Failed() error {
	return r.Error
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (en Endpoints) Get(ctx context.Context) (e []io.Employee, error error) {
	request := GetRequest{}
	response, err := en.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).E, response.(GetResponse).Error
}

// Add implements Service. Primarily useful in a client.
func (en Endpoints) Add(ctx context.Context, employee io.Employee) (e io.Employee, error error) {
	request := AddRequest{Employee: employee}
	response, err := en.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).E, response.(AddResponse).Error
}

// Delete implements Service. Primarily useful in a client.
func (en Endpoints) Delete(ctx context.Context, id string) (error error) {
	request := DeleteRequest{Id: id}
	response, err := en.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).Error
}
