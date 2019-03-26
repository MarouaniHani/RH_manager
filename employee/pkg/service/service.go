package service

import (
	"context"
	"mgo/employee/pkg/db"
	"mgo/employee/pkg/io"

	"gopkg.in/mgo.v2/bson"
)

// EmployeeService describes the service.
type EmployeeService interface {
	Get(ctx context.Context) (e []io.Employee, error error)
	Add(ctx context.Context, employee io.Employee) (e io.Employee, error error)
	Delete(ctx context.Context, id string) (error error)
}

type basicEmployeeService struct{}

func (b *basicEmployeeService) Get(ctx context.Context) (e []io.Employee, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return e, err
	}
	defer session.Close()
	c := session.DB("employee_app").C("employees")
	error = c.Find(nil).All(&e)
	return e, error
}
func (b *basicEmployeeService) Add(ctx context.Context, employee io.Employee) (e io.Employee, error error) {
	employee.Id = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return e, err
	}
	defer session.Close()
	c := session.DB("employee_app").C("employees")
	error = c.Insert(&employee)
	return employee, error
}
func (b *basicEmployeeService) Delete(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("employee_app").C("employees")
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}

// NewBasicEmployeeService returns a naive, stateless implementation of EmployeeService.
func NewBasicEmployeeService() EmployeeService {
	return &basicEmployeeService{}
}

// New returns a EmployeeService with all of the expected middleware wired in.
func New(middleware []Middleware) EmployeeService {
	var svc EmployeeService = NewBasicEmployeeService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
