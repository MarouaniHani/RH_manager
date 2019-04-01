package service

import (
	"context"
	"mgo/db"
	"mgo/employee/pkg/io"
	"mgo/utils"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

// EmployeeService describes the service.
type EmployeeService interface {
	Get(ctx context.Context) (e []io.Employee, error error)
	Add(ctx context.Context, employee io.Employee) (e io.Employee, error error)
	Delete(ctx context.Context, id string) (error error)
	GetByID(ctx context.Context, id string) (e io.Employee, error error)
	GetByCreteria(ctx context.Context, creteria string) (e []io.Employee, error error)
	GetByMultiCriteria(ctx context.Context, urlMap string) (e []io.Employee, error error)
}

type basicEmployeeService struct{}

func (b *basicEmployeeService) Get(ctx context.Context) (e []io.Employee, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return e, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employees")
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
	c := session.DB("Linda_app").C("employees")
	error = c.Insert(&employee)
	return employee, error
}
func (b *basicEmployeeService) Delete(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employees")
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
func (b *basicEmployeeService) GetByID(ctx context.Context, id string) (e io.Employee, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return e, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employees")
	error = c.FindId(bson.ObjectIdHex(id)).One(&e)
	return e, error
}

func (b *basicEmployeeService) GetByCreteria(ctx context.Context, creteria string) (e []io.Employee, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return e, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employees")
	error1 := c.Find(bson.M{"EmployeeName": creteria}).All(&e)
	if error1 == nil {
		error = error1
	}
	return e, error
}

func (b *basicEmployeeService) GetByMultiCriteria(ctx context.Context, urlMap string) (e []io.Employee, error error) {

	session, err := db.GetMongoSession()
	if err != nil {
		return e, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employees")
	qm := utils.QlSeparator(urlMap)
	if qm["EmployeeName"] != "" {
		error1 := c.Find(bson.M{"EmployeeName": qm["EmployeeName"]}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["ZipCode"] != "" {
		zipcode, _ := strconv.Atoi(qm["ZipCode"])
		error1 := c.Find(bson.M{"ZipCode": zipcode}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmployeeEmail"] != "" {
		error1 := c.Find(bson.M{"EmployeeEmail": qm["EmployeeEmail"]}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["Address"] != "" {
		error1 := c.Find(bson.M{"Address": qm["Address"]}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmployeeBirthDate"] != "" {
		error1 := c.Find(bson.M{"EmployeeBirthDate": qm["EmployeeBirthDate"]}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmployeeNumTel"] != "" {
		numTel, _ := strconv.Atoi(qm["EmployeeNumTel"])
		error1 := c.Find(bson.M{"EmployeeNumTel": numTel}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmergencyContactName"] != "" {
		error1 := c.Find(bson.M{"EmergencyContactName": qm["EmergencyContactName"]}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmergencyContactTel"] != "" {
		contactTel, _ := strconv.Atoi(qm["EmergencyContactTel"])
		error1 := c.Find(bson.M{"EmergencyContactTel": contactTel}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmployeeStartDate"] != "" {
		error1 := c.Find(bson.M{"EmployeeStartDate": qm["EmployeeStartDate"]}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmployeeSalary"] != "" {
		salary, _ := strconv.Atoi(qm["EmployeeSalary"])
		error1 := c.Find(bson.M{"EmployeeSalary": salary}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmployeeIban"] != "" {
		iban, _ := strconv.Atoi(qm["EmployeeIban"])
		error1 := c.Find(bson.M{"EmployeeIban": iban}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmployeeBic"] != "" {
		bic, _ := strconv.Atoi(qm["EmployeeBic"])
		error1 := c.Find(bson.M{"EmployeeBic": bic}).All(&e)
		if error1 == nil {
			error = error1
		}
	}

	return e, error
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
