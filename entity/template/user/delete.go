package user

import (
	"context"
	"taurus_go_demo/entity/template/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// UserEntityDelete is the delete action for the UserEntity.
type UserEntityDelete struct {
	*internal.Config
	tracker    entity.Tracker
	e          *UserEntity
	predicates []func(*entitysql.Predicate)
}

// NewUserEntityDelete creates a new UserEntityDelete.
func NewUserEntityDelete(c *internal.Config, e *UserEntity, t entity.Tracker) *UserEntityDelete {
	return &UserEntityDelete{
		Config:  c,
		e:       e,
		tracker: t,
	}
}

func (o *UserEntityDelete) Where(predicates ...func(*entitysql.Predicate)) *UserEntityDelete {
	o.predicates = append(o.predicates, predicates...)
	return o
}

func (o *UserEntityDelete) delete(ctx context.Context) error {
	return o.sqlDelete(ctx)
}

func (o *UserEntityDelete) sqlDelete(ctx context.Context) error {
	var (
		spec, err = o.deleteSpec()
		affected  = int64(0)
	)
	if err != nil {
		return err
	}
	spec.Affected = &affected
	if err := entitysql.NewDelete(ctx, o.Driver, spec); err != nil {
		return err
	}
	return internal.SetEntityState(o.e.Mutation, entity.Detached)
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
	id := &PredID{}
	o.predicates = append(o.predicates, id.EQ(*o.e.ID.Get()))
	uuid := &PredUUID{}
	o.predicates = append(o.predicates, entitysql.OpAnd, uuid.EQ(*o.e.UUID.Get()))
	if ps := o.predicates; len(ps) > 0 {
		spec.Predicate = func(p *entitysql.Predicate) {
			for _, f := range ps {
				f(p)
			}
		}
	}
	return spec, nil
}
