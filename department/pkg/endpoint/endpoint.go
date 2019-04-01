package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	io "mgo/department/pkg/io"
	service "mgo/department/pkg/service"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	D     []io.Department `json:"d"`
	Error error           `json:"error"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.DepartmentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		d, error := s.Get(ctx)
		return GetResponse{
			D:     d,
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
	Department io.Department `json:"department"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	D     io.Department `json:"d"`
	Error error         `json:"error"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.DepartmentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		d, error := s.Add(ctx, req.Department)
		return AddResponse{
			D:     d,
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
func MakeDeleteEndpoint(s service.DepartmentService) endpoint.Endpoint {
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

// GetByIDRequest collects the request parameters for the GetByID method.
type GetByIDRequest struct {
	Id string `json:"id"`
}

// GetByIDResponse collects the response parameters for the GetByID method.
type GetByIDResponse struct {
	D     io.Department `json:"d"`
	Error error         `json:"error"`
}

// MakeGetByIDEndpoint returns an endpoint that invokes GetByID on the service.
func MakeGetByIDEndpoint(s service.DepartmentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		d, error := s.GetByID(ctx, req.Id)
		return GetByIDResponse{
			D:     d,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByIDResponse) Failed() error {
	return r.Error
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context) (d []io.Department, error error) {
	request := GetRequest{}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).D, response.(GetResponse).Error
}

// Add implements Service. Primarily useful in a client.
func (e Endpoints) Add(ctx context.Context, department io.Department) (d io.Department, error error) {
	request := AddRequest{Department: department}
	response, err := e.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).D, response.(AddResponse).Error
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, id string) (error error) {
	request := DeleteRequest{Id: id}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).Error
}

// GetByID implements Service. Primarily useful in a client.
func (e Endpoints) GetByID(ctx context.Context, id string) (d io.Department, error error) {
	request := GetByIDRequest{Id: id}
	response, err := e.GetByIDEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByIDResponse).D, response.(GetByIDResponse).Error
}
