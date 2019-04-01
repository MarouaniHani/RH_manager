package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	io "mgo/department/pkg/io"
)

// Middleware describes a service middleware.
type Middleware func(DepartmentService) DepartmentService

type loggingMiddleware struct {
	logger log.Logger
	next   DepartmentService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a DepartmentService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next DepartmentService) DepartmentService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context) (d []io.Department, error error) {
	defer func() {
		l.logger.Log("method", "Get", "d", d, "error", error)
	}()
	return l.next.Get(ctx)
}
func (l loggingMiddleware) Add(ctx context.Context, department io.Department) (d io.Department, error error) {
	defer func() {
		l.logger.Log("method", "Add", "department", department, "d", d, "error", error)
	}()
	return l.next.Add(ctx, department)
}
func (l loggingMiddleware) Delete(ctx context.Context, id string) (error error) {
	defer func() {
		l.logger.Log("method", "Delete", "id", id, "error", error)
	}()
	return l.next.Delete(ctx, id)
}
func (l loggingMiddleware) GetByID(ctx context.Context, id string) (d io.Department, error error) {
	defer func() {
		l.logger.Log("method", "GetByID", "id", id, "d", d, "error", error)
	}()
	return l.next.GetByID(ctx, id)
}
