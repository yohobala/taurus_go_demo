package user

import (
	"context"
	"taurus_go_demo/entity/template/internal"

	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// UserEntityCreate is the create action for the UserEntity.
type UserEntityCreate struct {
	*internal.Config
	es []*UserEntity
}

// NewUserEntityCreate creates a new UserEntityCreate.
func NewUserEntityCreate(c *internal.Config, es ...*UserEntity) *UserEntityCreate {
	return &UserEntityCreate{
		Config: c,
		es:     es,
	}
}

func (o *UserEntityCreate) create(ctx context.Context, tx dialect.Tx) error {
	return o.sqlCreate(ctx, tx)
}

func (o *UserEntityCreate) sqlCreate(ctx context.Context, tx dialect.Tx) error {
	var (
		spec, err = o.createSpec()
		res       = o.es
		cursor    = 0
	)
	if err != nil {
		return err
	}
	spec.Scan = func(rows dialect.Rows, fields []entitysql.FieldName) error {
		e := res[cursor]
		cursor++
		if err := scan(e, fields, rows); err != nil {
			return err
		} else {
			res = append(res, e)
			return e.setUnchanged()
		}
	}
	return entitysql.NewCreate(ctx, tx, spec)
}

func (o *UserEntityCreate) createSpec() (*entitysql.CreateSpec, error) {
	returning := []entitysql.FieldName{
		FieldID.Name,
		FieldCreatedTime.Name,
		FieldDesc.Name,
	}
	spec := entitysql.NewCreateSpec(Entity, columns)
	spec.Fields = make([][]*entitysql.FieldSpec, 0, len(o.es))
	for _, e := range o.es {
		fields := make([]*entitysql.FieldSpec, 0, len(columns))
		for j := range columns {
			switch columns[j] {
			case FieldUUID.Name:
				if err := spec.CheckRequired(FieldUUID.Name, e.UUID); err != nil {
					return nil, err
				}
				fields = append(fields, &entitysql.FieldSpec{
					Column: FieldUUID.Name.String(),
					Value:  e.UUID.Value(),
				})
			case FieldName.Name:
				if err := spec.CheckRequired(FieldName.Name, e.Name); err != nil {
					return nil, err
				}
				fields = append(fields, &entitysql.FieldSpec{
					Column: FieldName.Name.String(),
					Value:  e.Name.Value(),
				})
			case FieldEmail.Name:
				fields = append(fields, &entitysql.FieldSpec{
					Column: FieldEmail.Name.String(),
					Value:  e.Email.Value(),
				})
			case FieldDesc.Name:
				fields = append(fields, &entitysql.FieldSpec{
					Column: FieldDesc.Name.String(),
					Value:  e.Desc.Value(),
				})
			}
		}
		spec.Fields = append(spec.Fields, fields)
	}
	spec.Returning = returning
	return spec, nil
}
