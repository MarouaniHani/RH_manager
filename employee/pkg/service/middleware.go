package service

import (
	"context"
	io "mgo/employee/pkg/io"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(EmployeeService) EmployeeService

type loggingMiddleware struct {
	logger log.Logger
	next   EmployeeService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a EmployeeService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next EmployeeService) EmployeeService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context) (e []io.Employee, error error) {
	defer func() {
		l.logger.Log("method", "Get", "e", e, "error", error)
	}()
	return l.next.Get(ctx)
}
func (l loggingMiddleware) Add(ctx context.Context, employee io.Employee) (e io.Employee, error error) {
	defer func() {
		l.logger.Log("method", "Add", "employee", employee, "e", e, "error", error)
	}()
	return l.next.Add(ctx, employee)
}
func (l loggingMiddleware) Delete(ctx context.Context, id string) (error error) {
	defer func() {
		l.logger.Log("method", "Delete", "id", id, "error", error)
	}()
	return l.next.Delete(ctx, id)
}

func (l loggingMiddleware) GetByID(ctx context.Context, id string) (e io.Employee, error error) {
	defer func() {
		l.logger.Log("method", "GetByID", "id", id, "e", e, "error", error)
	}()
	return l.next.GetByID(ctx, id)
}

func (l loggingMiddleware) GetByCreteria(ctx context.Context, creteria string) (e []io.Employee, error error) {
	defer func() {
		l.logger.Log("method", "GetByCreteria", "creteria", creteria, "e", e, "error", error)
	}()
	return l.next.GetByCreteria(ctx, creteria)
}

func (l loggingMiddleware) GetByMultiCriteria(ctx context.Context, urlMap string) (e []io.Employee, error error) {
	defer func() {
		l.logger.Log("method", "GetByMultiCriteria", "urlMap", urlMap, "e", e, "error", error)
	}()
	return l.next.GetByMultiCriteria(ctx, urlMap)
}
