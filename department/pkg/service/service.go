package service

import (
	"context"
	"mgo/db"
	"mgo/department/pkg/io"

	"gopkg.in/mgo.v2/bson"
)

// DepartmentService describes the service.
type DepartmentService interface {
	Get(ctx context.Context) (d []io.Department, error error)
	Add(ctx context.Context, department io.Department) (d io.Department, error error)
	Delete(ctx context.Context, id string) (error error)
	GetByID(ctx context.Context, id string) (d io.Department, error error)
}

type basicDepartmentService struct{}

func (b *basicDepartmentService) Get(ctx context.Context) (d []io.Department, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return d, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("departments")
	error = c.Find(nil).All(&d)
	return d, error
}
func (b *basicDepartmentService) Add(ctx context.Context, department io.Department) (d io.Department, error error) {
	department.ID = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return d, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("departments")
	error = c.Insert(&department)
	return department, error
}
func (b *basicDepartmentService) Delete(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("departments")
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
func (b *basicDepartmentService) GetByID(ctx context.Context, id string) (d io.Department, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return d, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("departments")
	error = c.FindId(bson.ObjectIdHex(id)).One(&d)
	return d, error
}

// NewBasicDepartmentService returns a naive, stateless implementation of DepartmentService.
func NewBasicDepartmentService() DepartmentService {
	return &basicDepartmentService{}
}

// New returns a DepartmentService with all of the expected middleware wired in.
func New(middleware []Middleware) DepartmentService {
	var svc DepartmentService = NewBasicDepartmentService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
