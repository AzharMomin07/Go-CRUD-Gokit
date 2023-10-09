package model

import (
	"AdminPortal/internal/dbmodels"
	"github.com/volatiletech/null/v8"
)

type Org struct {
	Id       int64       `json:"id"`
	Name     null.String `json:"fullName"`
	Location null.String `json:"location"`
}

func (o Org) ModelToDb() dbmodels.ORG {
	dbmodel := dbmodels.ORG{}
	dbmodel.ID = o.Id
	dbmodel.Name = o.Name
	dbmodel.Location = o.Location
	return dbmodel
}

func DbToModel(org dbmodels.ORG) Org {
	model := Org{}
	model.Id = org.ID
	model.Name = org.Name
	model.Location = org.Location
	return model
}
