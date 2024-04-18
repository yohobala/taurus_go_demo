// Code generated by taurus_go, DO NOT EDIT.

package entity

import (
	"context"
	"taurus_go_demo/entity/new/entity/author"
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// AuthorEntityDelete is the delete action for the AuthorEntity.
type AuthorEntityDelete struct {
	config     *internal.Dialect
	es         []*AuthorEntity
	predicates []entitysql.PredicateFunc
}

// NewAuthorEntityDelete creates a new AuthorEntityDelete.
func NewAuthorEntityDelete(c *internal.Dialect, es ...*AuthorEntity) *AuthorEntityDelete {
	return &AuthorEntityDelete{
		config: c,
		es:     es,
	}
}

// Where adds a predicate to the delete action.
func (o *AuthorEntityDelete) Where(predicates ...entitysql.PredicateFunc) *AuthorEntityDelete {
	o.predicates = append(o.predicates, predicates...)
	return o
}

func (o *AuthorEntityDelete) delete(ctx context.Context, tx dialect.Tx) error {
	return o.sqlDelete(ctx, tx)
}

func (o *AuthorEntityDelete) sqlDelete(ctx context.Context, tx dialect.Tx) error {
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

func (o *AuthorEntityDelete) deleteSpec() (*entitysql.DeleteSpec, error) {
	spec := entitysql.NewDeleteSpec(author.Entity)
	if ps := o.predicates; len(ps) > 0 {
		spec.Predicate = func(p *entitysql.Predicate, as string) {
			for _, f := range ps {
				f(p, as)
			}
		}
	}
	predID := &author.PredID{}
	if o.predicates == nil {
		o.predicates = make([]entitysql.PredicateFunc, 0, len(o.es))
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
		spec.Predicate = func(p *entitysql.Predicate, as string) {
			for _, f := range ps {
				f(p, as)
			}
		}
	}
	return spec, nil
}
