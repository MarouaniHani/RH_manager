package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type Department struct {
	ID             bson.ObjectId `json:"id" bson:"_id"`
	DepartmentName string        `json:"DepartmentName" bson:"DepartmentName"`
}

func (t Department) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
