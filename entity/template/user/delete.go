package user

import (
	"context"
	"taurus_go_demo/entity/template/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// UserEntityDelete is the delete action for the UserEntity.
type UserEntityDelete struct {
	*internal.Config
	es         []*UserEntity
	predicates []func(*entitysql.Predicate)
}

// NewUserEntityDelete creates a new UserEntityDelete.
func NewUserEntityDelete(c *internal.Config, es ...*UserEntity) *UserEntityDelete {
	return &UserEntityDelete{
		Config: c,
		es:     es,
	}
}

func (o *UserEntityDelete) Where(predicates ...func(*entitysql.Predicate)) *UserEntityDelete {
	o.predicates = append(o.predicates, predicates...)
	return o
}

func (o *UserEntityDelete) delete(ctx context.Context, tx dialect.Tx) error {
	return o.sqlDelete(ctx, tx)
}

func (o *UserEntityDelete) sqlDelete(ctx context.Context, tx dialect.Tx) error {
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

func (o *UserEntityDelete) deleteSpec() (*entitysql.DeleteSpec, error) {
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
