// Code generated by taurus_go, DO NOT EDIT.

package blog

import (
	"fmt"
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

type BlogEntity struct {
	config      *BlogEntityConfig
	ID          *IDType
	UUID        *UUIDType
	Desc        *DescType
	CreatedTime *CreatedTimeType
}

type BlogEntityConfig struct {
	*mutations
	*entity.Mutation
	*internal.Config
}

// New creates a new BlogEntity, but does not add tracking.
func New(c *internal.Config, ms *mutations) *BlogEntity {
	b := entity.NewMutation(entity.Detached)
	e := &BlogEntity{
		config: &BlogEntityConfig{
			Mutation:  b,
			Config:    c,
			mutations: ms,
		},
	}
	e.setState(entity.Detached)
	e.ID = newIDType(e.config)
	e.UUID = newUUIDType(e.config)
	e.Desc = newDescType(e.config)
	e.CreatedTime = newCreatedTimeType(e.config)
	return e
}

// String implements the fmt.Stringer interface.
func (e *BlogEntity) String() string {
	return fmt.Sprintf("{ ID: %v, UUID: %v, Desc: %v, CreatedTime: %v }",
		e.ID,
		e.UUID,
		e.Desc,
		e.CreatedTime,
	)
}

// State returns the state of the BlogEntity.
func (e *BlogEntity) State() entity.EntityState {
	return e.config.State()
}

func (e *BlogEntity) remove() error {
	return e.setState(entity.Deleted)
}

// create creates a new BlogEntity and adds tracking.
func (e *BlogEntity) create(uuid string, options ...func(*BlogEntity)) (*BlogEntity, error) {
	e.setState(entity.Added)
	e.UUID.Set(uuid)
	for _, option := range options {
		option(e)
	}
	return e, nil
}

func (e *BlogEntity) setUnchanged() error {
	return e.setState(entity.Unchanged)
}

func (e *BlogEntity) setState(state entity.EntityState) error {
	return e.config.mutations.SetEntityState(e, state)
}

func scan(e *BlogEntity, fields []entitysql.FieldName, rows dialect.Rows) error {
	args := make([]interface{}, len(fields))
	for i := range fields {
		switch fields[i] {
		case FieldID.Name:
			args[i] = e.ID
		case FieldUUID.Name:
			args[i] = e.UUID
		case FieldDesc.Name:
			args[i] = e.Desc
		case FieldCreatedTime.Name:
			args[i] = e.CreatedTime
		}
	}
	if err := rows.Scan(args...); err != nil {
		return err
	}
	return nil
}
