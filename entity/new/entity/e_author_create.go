// Code generated by taurus_go/entity, DO NOT EDIT.

package entity

import (
	"context"
	"taurus_go_demo/entity/new/entity/author"
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// authorEntityCreate is the create action for the authorEntity.
type authorEntityCreate struct {
	config *internal.Dialect
	es     []*AuthorEntity
}

// newAuthorEntityCreate creates a new authorEntityCreate.
func newAuthorEntityCreate(c *internal.Dialect, es ...*AuthorEntity) *authorEntityCreate {
	return &authorEntityCreate{
		config: c,
		es:     es,
	}
}

// create executes the create action.
func (o *authorEntityCreate) create(ctx context.Context, tx dialect.Tx) error {
	return o.sqlCreate(ctx, tx)
}

// sqlCreate executes the SQL create action.
func (o *authorEntityCreate) sqlCreate(ctx context.Context, tx dialect.Tx) error {
	var (
		spec, err = o.createSpec()
		res       = o.es
		cursor    = 0
	)
	if err != nil {
		return err
	}
	spec.Scan = func(rows dialect.Rows, fields []entitysql.ScannerField) error {
		e := res[cursor]
		cursor++
		args := e.scan(fields)
		if err := rows.Scan(args...); err != nil {
			return err
		} else {
			res = append(res, e)
			return e.setUnchanged()
		}
	}
	return entitysql.NewCreate(ctx, tx, spec)
}

// createSpec creates the create action spec. It checks for required fields and sets the returning fields.
func (o *authorEntityCreate) createSpec() (*entitysql.CreateSpec, error) {
	returning := []entitysql.FieldName{
		author.FieldID.Name,
	}
	entity := author.Entity
	columns := author.Columns
	spec := entitysql.NewCreateSpec(entity, columns)
	spec.Fields = make([][]*entitysql.FieldSpec, 0, len(o.es))
	for _, e := range o.es {
		fields := make([]*entitysql.FieldSpec, 0, len(author.Columns))
		for j := range author.Columns {
			switch author.Columns[j] {
			case author.FieldName.Name:
				v, err := e.Name.SqlParam(o.config.Driver.Dialect())
				if err != nil {
					return nil, err
				}
				if err := spec.CheckRequired(o.config.Driver.Dialect(), author.FieldName.Name, e.Name); err != nil {
					return nil, err
				}
				fieldSpace := entitysql.NewFieldSpec(author.FieldName.Name)
				fieldSpace.Param = v
				fieldSpace.ParamFormat = e.Name.SqlFormatParam()
				fields = append(fields, &fieldSpace)
			}
		}
		spec.Fields = append(spec.Fields, fields)
	}
	spec.Returning = returning
	return spec, nil
}
