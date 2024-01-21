package user

import (
	"taurus_go_demo/entity/template/internal"
	"time"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/field"
)

type IDType struct {
	field.IntStorage[int64]
	mutator *entity.Mutation
}

func newIDType(m *entity.Mutation) *IDType {
	t := &IDType{}
	t.mutator = m
	return t
}

func (t *IDType) Set(v int64) {
	t.IntStorage.Set(v)
	internal.SetEntityState(t.mutator, entity.Modified)
}

func (t *IDType) Get() *int64 {
	return t.IntStorage.Get()
}

type UUIDType struct {
	field.StringStorage
	mutator *entity.Mutation
}

func newUUIDType(m *entity.Mutation) *UUIDType {
	t := &UUIDType{}
	t.mutator = m
	return t
}

func (t *UUIDType) Set(v string) {
	t.StringStorage.Set(v)
	internal.SetEntityState(t.mutator, entity.Modified)
}

func (t *UUIDType) Get() *string {
	return t.StringStorage.Get()
}

type NameType struct {
	field.StringStorage
	mutator *entity.Mutation
}

func newNameType(m *entity.Mutation) *NameType {
	t := &NameType{}
	t.mutator = m
	return t
}

func (t *NameType) Set(v string) {
	t.StringStorage.Set(v)
	internal.SetEntityState(t.mutator, entity.Modified)
}

func (t *NameType) Get() *string {
	return t.StringStorage.Get()
}

type EmailType struct {
	field.StringStorage
	mutator *entity.Mutation
}

func newEmailType(m *entity.Mutation) *EmailType {
	t := &EmailType{}
	t.mutator = m
	return t
}

func (t *EmailType) Set(v string) {
	t.StringStorage.Set(v)
	internal.SetEntityState(t.mutator, entity.Modified)
}

func (t *EmailType) Get() *string {
	return t.StringStorage.Get()
}

type CreatedTimeType struct {
	field.TimestampStorage
	mutator *entity.Mutation
}

func newCreatedTimeType(m *entity.Mutation) *CreatedTimeType {
	t := &CreatedTimeType{}
	t.mutator = m
	return t
}

func (t *CreatedTimeType) Set(v time.Time) {
	t.TimestampStorage.Set(v)
	internal.SetEntityState(t.mutator, entity.Modified)
}

func (t *CreatedTimeType) Get() *time.Time {
	return t.TimestampStorage.Get()
}

type DescType struct {
	field.StringStorage
	mutator *entity.Mutation
}

func newDescType(m *entity.Mutation) *DescType {
	t := &DescType{}
	t.mutator = m
	return t
}

func (t *DescType) Set(v string) {
	t.StringStorage.Set(v)
	internal.SetEntityState(t.mutator, entity.Modified)
	t.mutator.SetFields(FieldDesc.Name.String())
}

func (t *DescType) Get() *string {
	return t.StringStorage.Get()
}
