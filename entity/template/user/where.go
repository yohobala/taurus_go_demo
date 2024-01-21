package user

import (
	"time"

	"github.com/yohobala/taurus_go/entity/entitysql"
)

type PredID struct {
}

func (f *PredID) EQ(id int64) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldID.Name.String(), id)
	}
}

type PredUUID struct {
}

func (f *PredUUID) EQ(uuid string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldUUID.Name.String(), uuid)
	}
}

type PredName struct {
}

func (f *PredName) EQ(name string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldName.Name.String(), name)
	}
}

type PredEmail struct {
}

func (f *PredEmail) EQ(email string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldEmail.Name.String(), email)
	}
}

type PredCreatedTime struct {
}

func (f *PredCreatedTime) EQ(created_time time.Time) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldCreatedTime.Name.String(), created_time)
	}
}

type PredDesc struct {
}

func (f *PredDesc) EQ(desc string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldDesc.Name.String(), desc)
	}
}
