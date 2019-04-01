package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type Employee struct {
	Id                   bson.ObjectId `json:"id" bson:"_id"`
	EmployeeName         string        `json:"EmployeeName" bson:"EmployeeName"`
	EmployeeEmail        string        `json:"EmployeeEmail" bson:"EmployeeEmail"`
	Address              string        `json:"Address" bson:"Address"`
	ZipCode              int           `json:"ZipCode" bson:"ZipCode"`
	EmployeeBirthDate    string        `json:"EmployeeBirthDate" bson:"EmployeeBirthDate"`
	EmployeeNumTel       int           `json:"EmployeeNumTel" bson:"EmployeeNumTel"`
	EmergencyContactName string        `json:"EmergencyContactName" bson:"EmergencyContactName"`
	EmergencyContactTel  int           `json:"EmergencyContactTel" bson:"EmergencyContactTel"`
	EmployeeStartDate    string        `json:"EmployeeStartDate" bson:"EmployeeStartDate"`
	EmployeeSalary       float64       `json:"EmployeeSalary" bson:"EmployeeSalary"`
	EmployeeIban         int           `json:"EmployeeIban" bson:"EmployeeIban"`
	EmployeeBic          int           `json:"EmployeeBic" bson:"EmployeeBic"`
	DepartmentID         bson.ObjectId `json:"DepartmentID" bson:"DepartmentID"`
}

func (t Employee) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
