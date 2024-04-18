// Code generated by taurus_go, DO NOT EDIT.

package entity

import (
	"context"
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"

	"taurus_go_demo/entity/new/entity/blog"
)

// BlogEntityCreate is the create action for the BlogEntity.
type BlogEntityCreate struct {
	config *internal.Dialect
	es     []*BlogEntity
}

// NewBlogEntityCreate creates a new BlogEntityCreate.
func NewBlogEntityCreate(c *internal.Dialect, es ...*BlogEntity) *BlogEntityCreate {
	return &BlogEntityCreate{
		config: c,
		es:     es,
	}
}

// create executes the create action.
func (o *BlogEntityCreate) create(ctx context.Context, tx dialect.Tx) error {
	return o.sqlCreate(ctx, tx)
}

// sqlCreate executes the SQL create action.
func (o *BlogEntityCreate) sqlCreate(ctx context.Context, tx dialect.Tx) error {
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
		args := e.scan( fields)
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
func (o *BlogEntityCreate) createSpec() (*entitysql.CreateSpec, error) {
	returning := []entitysql.FieldName{
		blog.FieldID.Name,
		blog.FieldCreatedTime.Name,
	}
	spec := entitysql.NewCreateSpec(blog.Entity, blog.Columns)
	spec.Fields = make([][]*entitysql.FieldSpec, 0, len(o.es))
	for _, e := range o.es {
		fields := make([]*entitysql.FieldSpec, 0, len(blog.Columns))
		for j := range blog.Columns {
			switch blog.Columns[j] {
			case blog.FieldUUID.Name:
				if err := spec.CheckRequired(blog.FieldUUID.Name, e.UUID); err != nil {
					return nil, err
				}
				fields = append(fields, &entitysql.FieldSpec{
					Column: blog.FieldUUID.Name.String(),
					Value:  e.UUID.Value(),
				})
			case blog.FieldDesc.Name:
				fields = append(fields, &entitysql.FieldSpec{
					Column: blog.FieldDesc.Name.String(),
					Value:  e.Desc.Value(),
				})
			case blog.FieldCreatedTime.Name:
				fields = append(fields, &entitysql.FieldSpec{
					Column: blog.FieldCreatedTime.Name.String(),
					Value:  e.CreatedTime.Value(),
				})
			}
		}
		spec.Fields = append(spec.Fields, fields)
	}
	spec.Returning = returning
	return spec, nil
}