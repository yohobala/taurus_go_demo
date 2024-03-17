// Code generated by taurus_go, DO NOT EDIT.

package blog

import (
	"context"
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// BlogEntityDelete is the delete action for the BlogEntity.
type BlogEntityDelete struct {
	*internal.Config
	es         []*BlogEntity
	predicates []func(*entitysql.Predicate)
}

// NewBlogEntityDelete creates a new BlogEntityDelete.
func NewBlogEntityDelete(c *internal.Config, es ...*BlogEntity) *BlogEntityDelete {
	return &BlogEntityDelete{
		Config: c,
		es:     es,
	}
}

func (o *BlogEntityDelete) Where(predicates ...func(*entitysql.Predicate)) *BlogEntityDelete {
	o.predicates = append(o.predicates, predicates...)
	return o
}

func (o *BlogEntityDelete) delete(ctx context.Context, tx dialect.Tx) error {
	return o.sqlDelete(ctx, tx)
}

func (o *BlogEntityDelete) sqlDelete(ctx context.Context, tx dialect.Tx) error {
	var (
		spec, err = o.deleteSpec()
		affected  = int64(0)
	)
	if err != nil {
		return err
	}
	spec.Affected = &affected
	if err := entitysql.NewDelete(ctx, tx, spec); err != nil {
		return err
	}
	for _, e := range o.es {
		if err := e.setState(entity.Detached); err != nil {
			return err
		}
	}
	return nil
}

func (o *BlogEntityDelete) deleteSpec() (*entitysql.DeleteSpec, error) {
	spec := entitysql.NewDeleteSpec(Entity)
	if ps := o.predicates; len(ps) > 0 {
		spec.Predicate = func(p *entitysql.Predicate) {
			for _, f := range ps {
				f(p)
			}
		}
	}
	predID := &PredID{}
	if o.predicates == nil {
		o.predicates = make([]func(*entitysql.Predicate), 0, len(o.es))
	}
	for i, e := range o.es {
		if e.ID.Get() != nil {
			if i >= 1 {
				o.predicates = append(o.predicates, entitysql.Or)
			}
			o.predicates = append(o.predicates, predID.EQ(*e.ID.Get()))
		}
	}
	if ps := o.predicates; len(ps) > 0 {
		spec.Predicate = func(p *entitysql.Predicate) {
			for _, f := range ps {
				f(p)
			}
		}
	}
	return spec, nil
}
