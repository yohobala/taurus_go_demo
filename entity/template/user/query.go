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
	config     *UserEntityConfig
	ctx        *entitysql.QueryContext
	predicates []func(*entitysql.Predicate)
}

// NewUserEntityQuery creates a new UserEntityQuery.
func NewUserEntityQuery(c *internal.Config, t entity.Tracker, ms *mutations) *UserEntityQuery {
	return &UserEntityQuery{
		config: &UserEntityConfig{
			Config:    c,
			mutations: ms,
		},
		ctx: &entitysql.QueryContext{},
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
		res  = New(o.config.Config, o.config.mutations)
	)
	spec.Scan = func(rows dialect.Rows, fields []entitysql.FieldName) error {
		return scan(res, fields, rows)
	}
	if err := entitysql.NewQuery(ctx, o.config.Driver, spec); err != nil {
		return nil, err
	}
	if err := res.setUnchanged(); err != nil {
		return nil, err
	}
	return res, nil
}

func (o *UserEntityQuery) sqlAll(ctx context.Context) ([]*UserEntity, error) {
	var (
		spec = o.querySpec()
		res  = []*UserEntity{}
	)
	spec.Scan = func(rows dialect.Rows, fields []entitysql.FieldName) error {
		e := New(o.config.Config, o.config.mutations)
		if err := scan(e, fields, rows); err != nil {
			return err
		} else {
			res = append(res, e)
			return nil
		}
	}
	if err := entitysql.NewQuery(ctx, o.config.Driver, spec); err != nil {
		return nil, err
	}
	for _, e := range res {
		if err := e.setUnchanged(); err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (o *UserEntityQuery) querySpec() *entitysql.QuerySpec {
	s := entitysql.NewQuerySpec(Entity, columns)
	if o.ctx.Limit != nil {
		s.Limit = *o.ctx.Limit
	}
	if fields := o.ctx.Fields; len(fields) > 0 {
		s.Entity.Columns = make([]entitysql.FieldName, 0, len(fields))
		s.Entity.Columns = append(s.Entity.Columns, fields...)
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
