package user

import (
	"context"
	"taurus_go_demo/entity/template/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// UserEntityCreate is the create action for the UserEntity.
type UserEntityCreate struct {
	*internal.Config
	tracker entity.Tracker
	e       *UserEntity
}

// NewUserEntityCreate creates a new UserEntityCreate.
func NewUserEntityCreate(c *internal.Config, e *UserEntity, t entity.Tracker) *UserEntityCreate {
	return &UserEntityCreate{
		Config:  c,
		e:       e,
		tracker: t,
	}
}

func (o *UserEntityCreate) create(ctx context.Context) (*UserEntity, error) {
	return o.sqlCreate(ctx)
}

func (o *UserEntityCreate) sqlCreate(ctx context.Context) (*UserEntity, error) {
	var (
		spec, err = o.createSpec()
		res       = o.e
	)
	if err != nil {
		return nil, err
	}
	spec.Scan = func(rows dialect.Rows, fields []entitysql.FieldName) error {
		return scan(res, fields, rows)
	}
	if err := entitysql.NewCreate(ctx, o.Driver, spec); err != nil {
		return nil, err
	}
	setUnchanged(o.tracker, res)
	return res, nil
}

func (o *UserEntityCreate) createSpec() (*entitysql.CreateSpec, error) {
	returning := []entitysql.FieldName{
		FieldID.Name,
		FieldCreatedTime.Name,
		FieldDesc.Name,
	}
	spec := entitysql.NewCreateSpec(Entity, returning)
	spec.Fields = make([]*entitysql.FieldSpec, 0, len(rows))
	for i := range rows {
		switch rows[i] {
		case FieldUUID.Name:
			if err := spec.CheckRequired(FieldUUID.Name, o.e.UUID); err != nil {
				return nil, err
			}
			spec.Fields = append(spec.Fields, &entitysql.FieldSpec{
				Column: FieldUUID.Name.String(),
				Value:  o.e.UUID.Value(),
			})
		case FieldName.Name:
			if err := spec.CheckRequired(FieldName.Name, o.e.Name); err != nil {
				return nil, err
			}
			spec.Fields = append(spec.Fields, &entitysql.FieldSpec{
				Column: FieldName.Name.String(),
				Value:  o.e.Name.Value(),
			})
		}
	}
	spec.Returning = returning
	return spec, nil
}
