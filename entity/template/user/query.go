package user

import (
	"context"
	"taurus_go_demo/entity/template/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// UserEntityQuery is the query action for the UserEntity.
type UserEntityQuery struct {
	*internal.Config
	tracker    entity.Tracker
	ctx        *entitysql.QueryContext
	predicates []func(*entitysql.Predicate)
}

// NewUserEntityQuery creates a new UserEntityQuery.
func NewUserEntityQuery(c *internal.Config, t entity.Tracker) *UserEntityQuery {
	return &UserEntityQuery{
		Config:  c,
		ctx:     &entitysql.QueryContext{},
		tracker: t,
	}
}

func (o *UserEntityQuery) Where(predicates ...func(*entitysql.Predicate)) *UserEntityQuery {
	o.predicates = append(o.predicates, predicates...)
	return o
}

// Limit sets the limit of the query.
func (o *UserEntityQuery) Limit(limit int) *UserEntityQuery {
	o.ctx.Limit = &limit
	return o
}

// First returns the first result of the query.
func (o *UserEntityQuery) First(ctx context.Context) (*UserEntity, error) {
	result, err := o.Single(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ToList returns the list of results of the query.
func (o *UserEntityQuery) ToList(ctx context.Context) ([]*UserEntity, error) {
	return o.sqlAll(ctx)
}

// Single returns the single result of the query.
func (o *UserEntityQuery) Single(ctx context.Context) (*UserEntity, error) {
	limit := 1
	o.ctx.Limit = &limit
	return o.sqlSingle(ctx)
}

func (o *UserEntityQuery) sqlSingle(ctx context.Context) (*UserEntity, error) {
	var (
		spec = o.querySpec()
		res  = New(o.Config, o.tracker)
	)
	spec.Scan = func(rows dialect.Rows, fields []entitysql.FieldName) error {
		return scan(res, fields, rows)
	}
	if err := entitysql.NewQuery(ctx, o.Driver, spec); err != nil {
		return nil, err
	}
	setUnchanged(o.tracker, res)
	return res, nil
}

func (o *UserEntityQuery) sqlAll(ctx context.Context) ([]*UserEntity, error) {
	var (
		spec = o.querySpec()
		res  = []*UserEntity{}
	)
	spec.Scan = func(rows dialect.Rows, fields []entitysql.FieldName) error {
		e := New(o.Config, o.tracker)
		if err := scan(e, fields, rows); err != nil {
			return err
		} else {
			res = append(res, e)
			return nil
		}
	}
	if err := entitysql.NewQuery(ctx, o.Driver, spec); err != nil {
		return nil, err
	}
	for _, e := range res {
		setUnchanged(o.tracker, e)
	}
	return res, nil
}

func (o *UserEntityQuery) querySpec() *entitysql.QuerySpec {
	s := entitysql.NewQuerySpec(Entity, rows)
	if o.ctx.Limit != nil {
		s.Limit = *o.ctx.Limit
	}
	if fields := o.ctx.Fields; len(fields) > 0 {
		s.Entity.Rows = make([]entitysql.FieldName, 0, len(fields))
		s.Entity.Rows = append(s.Entity.Rows, fields...)
	}
	if ps := o.predicates; len(ps) > 0 {
		s.Predicate = func(p *entitysql.Predicate) {
			for _, f := range ps {
				f(p)
			}
		}
	}
	return s
}
