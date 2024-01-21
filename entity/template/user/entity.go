package user

import (
	"context"
	"fmt"
	"taurus_go_demo/entity/template/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

type UserEntity struct {
	*internal.Config
	*entity.Mutation
	tracker     entity.Tracker
	ID          *IDType
	UUID        *UUIDType
	Name        *NameType
	Email       *EmailType
	CreatedTime *CreatedTimeType
	Desc        *DescType
}

// New creates a new UserEntity, but does not add tracking.
func New(c *internal.Config, t entity.Tracker) *UserEntity {
	e := &UserEntity{
		Config:  c,
		tracker: t,
	}
	b := entity.NewMutation(entity.Detached)
	e.Mutation = b
	e.ID = newIDType(b)
	e.UUID = newUUIDType(b)
	e.Name = newNameType(b)
	e.Email = newEmailType(b)
	e.CreatedTime = newCreatedTimeType(b)
	e.Desc = newDescType(b)
	return e
}

// String implements the fmt.Stringer interface.
func (e *UserEntity) String() string {
	return fmt.Sprintf("{ ID: %v, UUID: %v, Name: %v, Email: %v, CreatedTime: %v, Desc: %v }",
		e.ID,
		e.UUID,
		e.Name,
		e.Email,
		e.CreatedTime,
		e.Desc,
	)
}

// Exec executes all the mutations for the UserEntity.
func (e *UserEntity) Exec() error {
	switch e.State() {
	case entity.Added:
		n := NewUserEntityCreate(e.Config, e, e.tracker)
		_, err := n.create(context.Background())
		return err
	case entity.Modified:
		n := NewUserEntityUpdate(e.Config, e, e.tracker)
		_, err := n.update(context.Background())
		return err
	case entity.Deleted:
		n := NewUserEntityDelete(e.Config, e, e.tracker)
		return n.delete(context.Background())
	}
	return nil
}

func (e *UserEntity) remove() error {
	return internal.SetEntityState(e.Mutation, entity.Deleted)
}

// create creates a new UserEntity and adds tracking.
func (e *UserEntity) create(uuid string, name string, options ...func(*UserEntity)) (*UserEntity, error) {
	e.Mutation.SetState(entity.Added)
	e.UUID.Set(uuid)
	e.Name.Set(name)
	for _, option := range options {
		option(e)
	}
	return e, nil
}

func scan(e *UserEntity, fields []entitysql.FieldName, rows dialect.Rows) error {
	args := make([]interface{}, len(fields))
	for i := range fields {
		switch fields[i] {
		case FieldID.Name:
			args[i] = e.ID
		case FieldUUID.Name:
			args[i] = e.UUID
		case FieldName.Name:
			args[i] = e.Name
		case FieldEmail.Name:
			args[i] = e.Email
		case FieldCreatedTime.Name:
			args[i] = e.CreatedTime
		case FieldDesc.Name:
			args[i] = e.Desc
		}
	}

	if err := rows.Scan(args...); err != nil {
		return err
	}
	return nil
}

func setUnchanged(t entity.Tracker, m entity.Mutator) {
	m.SetState(entity.Unchanged)
	t.Add(m)
}
