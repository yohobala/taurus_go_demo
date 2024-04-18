// Code generated by taurus_go, DO NOT EDIT.

package entity

import (
	"context"
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"

	"taurus_go_demo/entity/new/entity/author"
)

// AuthorEntityQuery is the query action for the AuthorEntity.
type AuthorEntityQuery struct {
	config     *AuthorEntityConfig
	ctx        *entitysql.QueryContext
	predicates []func(*entitysql.Predicate)
	scanner    []*internal.QueryScanner
	scannerTotal int
}

// First returns the first result of the query.
func (o *AuthorEntityQuery) First(ctx context.Context) (*AuthorEntity, error) {
	result, err := o.Single(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// NewAuthorEntityQuery creates a new AuthorEntityQuery.
func NewAuthorEntityQuery(c *internal.Dialect, t entity.Tracker, ms *aothorMutations) *AuthorEntityQuery {
	return &AuthorEntityQuery{
		config: &AuthorEntityConfig{
			Dialect:    c,
			aothorMutations: ms,
		},
		ctx: &entitysql.QueryContext{},
	}
}

func (o *AuthorEntityQuery) Where(predicates ...func(*entitysql.Predicate)) *AuthorEntityQuery {
	o.predicates = append(o.predicates, predicates...)
	return o
}

// Limit sets the limit of the query.
func (o *AuthorEntityQuery) Limit(limit int) *AuthorEntityQuery {
	o.ctx.Limit = &limit
	return o
}

// ToList returns the list of results of the query.
func (o *AuthorEntityQuery) ToList(ctx context.Context) ([]*AuthorEntity, error) {
	return o.sqlAll(ctx)
}

// Single returns the single result of the query.
func (o *AuthorEntityQuery) Single(ctx context.Context) (*AuthorEntity, error) {
	limit := 1
	o.ctx.Limit = &limit
	return o.sqlSingle(ctx)
}

func (o *AuthorEntityQuery) sqlSingle(ctx context.Context) (*AuthorEntity, error) {
	var (
		spec = o.querySpec()
		res  = o.config.New()
	)
	switch res := res.(type) {
	case *AuthorEntity:
		spec.Scan = func(rows dialect.Rows, fields []entitysql.ScannerField) error {
			builder := entitysql.NewScannerBuilder(o.scannerTotal)
			builder.Append(0,res.scan(fields)...)
			return rows.Scan(builder.Flatten()...)
		}
		if err := entitysql.NewQuery(ctx, o.config.Driver, spec); err != nil {
			return nil, err
		}
		if err := res.setUnchanged(); err != nil {
			return nil, err
		}
		return res, nil
	default:
		return nil, entity.Err_0100030006
	}
}

func (o *AuthorEntityQuery) sqlAll(ctx context.Context) ([]*AuthorEntity, error) {
	var (
		spec = o.querySpec()
		res  = []*AuthorEntity{}
	)
	spec.Scan = func(rows dialect.Rows, fields []entitysql.ScannerField) error {
		e := o.config.New()
		switch e := e.(type) {
		case *AuthorEntity:
			builder := entitysql.NewScannerBuilder(o.scannerTotal)
			builder.Append(0,e.scan([]entitysql.ScannerField{})...)
			for _, s := range o.scanner {
				e.createRel(builder, s)
			}

			if err := rows.Scan(builder.Flatten()...); err != nil {
				return err
			} else {
				res = append(res, e)
				return nil
			}
		default:
			return entity.Err_0100030006
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

func (o *AuthorEntityQuery) querySpec() *entitysql.QuerySpec {
	s := entitysql.NewQuerySpec(author.Entity, author.Columns)
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