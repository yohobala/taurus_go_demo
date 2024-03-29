package user

import (
	"time"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/field"
)

type IDType struct {
	field.IntStorage[int64]
	config *UserEntityConfig
}

func newIDType(c *UserEntityConfig) *IDType {
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
	config *UserEntityConfig
}

func newUUIDType(c *UserEntityConfig) *UUIDType {
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

type NameType struct {
	field.StringStorage
	config *UserEntityConfig
}

func newNameType(c *UserEntityConfig) *NameType {
	t := &NameType{}
	t.config = c
	return t
}

func (t *NameType) Set(v string) {
	t.StringStorage.Set(v)
	t.config.mutations.ChangeEntityState(t.config.Mutation, entity.Modified)
	t.config.Mutation.SetFields(FieldName.Name.String())
}

func (t *NameType) Get() *string {
	return t.StringStorage.Get()
}

type EmailType struct {
	field.StringStorage
	config *UserEntityConfig
}

func newEmailType(c *UserEntityConfig) *EmailType {
	t := &EmailType{}
	t.config = c
	return t
}

func (t *EmailType) Set(v string) {
	t.StringStorage.Set(v)
	t.config.mutations.ChangeEntityState(t.config.Mutation, entity.Modified)
	t.config.Mutation.SetFields(FieldEmail.Name.String())
}

func (t *EmailType) Get() *string {
	return t.StringStorage.Get()
}

type CreatedTimeType struct {
	field.TimestampStorage
	config *UserEntityConfig
}

func newCreatedTimeType(c *UserEntityConfig) *CreatedTimeType {
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

type DescType struct {
	field.StringStorage
	config *UserEntityConfig
}

func newDescType(c *UserEntityConfig) *DescType {
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
