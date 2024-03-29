// Code generated by taurus_go, DO NOT EDIT.

package blog

import (
	"time"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/field"
)

type IDType struct {
	field.IntStorage[int64]
	config *BlogEntityConfig
}

func newIDType(c *BlogEntityConfig) *IDType {
	t := &IDType{}
	t.config = c
	return t
}

func (t *IDType) Set(v int64) {
	t.IntStorage.Set(v)
	t.config.mutations.ChangeEntityState(t.config.Mutation, entity.Modified)
	t.config.Mutation.SetFields(FieldID.Name.String())
}

func (t *IDType) Get() *int64 {
	return t.IntStorage.Get()
}

type UUIDType struct {
	field.StringStorage
	config *BlogEntityConfig
}

func newUUIDType(c *BlogEntityConfig) *UUIDType {
	t := &UUIDType{}
	t.config = c
	return t
}

func (t *UUIDType) Set(v string) {
	t.StringStorage.Set(v)
	t.config.mutations.ChangeEntityState(t.config.Mutation, entity.Modified)
	t.config.Mutation.SetFields(FieldUUID.Name.String())
}

func (t *UUIDType) Get() *string {
	return t.StringStorage.Get()
}

type DescType struct {
	field.StringStorage
	config *BlogEntityConfig
}

func newDescType(c *BlogEntityConfig) *DescType {
	t := &DescType{}
	t.config = c
	return t
}

func (t *DescType) Set(v string) {
	t.StringStorage.Set(v)
	t.config.mutations.ChangeEntityState(t.config.Mutation, entity.Modified)
	t.config.Mutation.SetFields(FieldDesc.Name.String())
}

func (t *DescType) Get() *string {
	return t.StringStorage.Get()
}

type CreatedTimeType struct {
	field.TimestampStorage
	config *BlogEntityConfig
}

func newCreatedTimeType(c *BlogEntityConfig) *CreatedTimeType {
	t := &CreatedTimeType{}
	t.config = c
	return t
}

func (t *CreatedTimeType) Set(v time.Time) {
	t.TimestampStorage.Set(v)
	t.config.mutations.ChangeEntityState(t.config.Mutation, entity.Modified)
	t.config.Mutation.SetFields(FieldCreatedTime.Name.String())
}

func (t *CreatedTimeType) Get() *time.Time {
	return t.TimestampStorage.Get()
}
