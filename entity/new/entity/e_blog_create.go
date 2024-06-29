// Code generated by taurus_go/entity, DO NOT EDIT.

package entity

import (
	"context"
	"taurus_go_demo/entity/new/entity/blog"
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// blogEntityCreate is the create action for the blogEntity.
type blogEntityCreate struct {
	config *internal.Dialect
	es     []*BlogEntity
}

// newBlogEntityCreate creates a new blogEntityCreate.
func newBlogEntityCreate(c *internal.Dialect, es ...*BlogEntity) *blogEntityCreate {
	return &blogEntityCreate{
		config: c,
		es:     es,
	}
}

// create executes the create action.
func (o *blogEntityCreate) create(ctx context.Context, tx dialect.Tx) error {
	return o.sqlCreate(ctx, tx)
}

// sqlCreate executes the SQL create action.
func (o *blogEntityCreate) sqlCreate(ctx context.Context, tx dialect.Tx) error {
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
func (o *blogEntityCreate) createSpec() (*entitysql.CreateSpec, error) {
	returning := []entitysql.FieldName{
		blog.FieldID.Name,
		blog.FieldCreatedTime.Name,
	}
	entity := blog.Entity
	columns := blog.Columns
	spec := entitysql.NewCreateSpec(entity, columns)
	spec.Fields = make([][]*entitysql.FieldSpec, 0, len(o.es))
	for _, e := range o.es {
		fields := make([]*entitysql.FieldSpec, 0, len(blog.Columns))
		for j := range blog.Columns {
			switch blog.Columns[j] {
			case blog.FieldUUID.Name:
				v, err := e.UUID.SqlParam(o.config.Driver.Dialect())
				if err != nil {
					return nil, err
				}
				if err := spec.CheckRequired(o.config.Driver.Dialect(), blog.FieldUUID.Name, e.UUID); err != nil {
					return nil, err
				}
				fieldSpace := entitysql.NewFieldSpec(blog.FieldUUID.Name)
				fieldSpace.Param = v
				fieldSpace.Format = e.UUID.SqlFormatParam()
				fields = append(fields, &fieldSpace)
			case blog.FieldDesc.Name:
				v, err := e.Desc.SqlParam(o.config.Driver.Dialect())
				if err != nil {
					return nil, err
				}
				fieldSpace := entitysql.NewFieldSpec(blog.FieldDesc.Name)
				fieldSpace.Param = v
				fieldSpace.Format = e.Desc.SqlFormatParam()
				fields = append(fields, &fieldSpace)
			case blog.FieldCreatedTime.Name:
				v, err := e.CreatedTime.SqlParam(o.config.Driver.Dialect())
				if err != nil {
					return nil, err
				}
				fieldSpace := entitysql.NewFieldSpec(blog.FieldCreatedTime.Name)
				fieldSpace.Param = v
				fieldSpace.Format = e.CreatedTime.SqlFormatParam()
				fields = append(fields, &fieldSpace)
			}
		}
		spec.Fields = append(spec.Fields, fields)
	}
	spec.Returning = returning
	return spec, nil
}
