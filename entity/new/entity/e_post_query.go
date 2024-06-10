// Code generated by taurus_go, DO NOT EDIT.

package entity

import (
	"context"
	"taurus_go_demo/entity/new/entity/internal"
	"taurus_go_demo/entity/new/entity/post"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// postEntityQuery is the query action for the postEntity.
type postEntityQuery struct {
	config       *postEntityConfig
	ctx          *entitysql.QueryContext
	predicates   []entitysql.PredicateFunc
	rels         []postEntityRel
	order        []post.OrderTerm
	scanner      []*internal.QueryScanner
	scannerTotal int
}

// First returns the first result of the query.
func (o *postEntityQuery) First(ctx context.Context) (*PostEntity, error) {
	result, err := o.Single(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// newPostEntityQuery creates a new postEntityQuery.
func newPostEntityQuery(c *internal.Dialect, t entity.Tracker, ms *postEntityMutations) *postEntityQuery {
	return &postEntityQuery{
		config: &postEntityConfig{
			Dialect:             c,
			postEntityMutations: ms,
		},
		ctx:          &entitysql.QueryContext{},
		predicates:   []entitysql.PredicateFunc{},
		rels:         []postEntityRel{},
		order:        []post.OrderTerm{},
		scanner:      []*internal.QueryScanner{},
		scannerTotal: 0,
	}
}

func (o *postEntityQuery) Where(predicates ...entitysql.PredicateFunc) *postEntityQuery {
	o.predicates = append(o.predicates, predicates...)
	return o
}

// Limit sets the limit of the query.
func (o *postEntityQuery) Limit(limit int) *postEntityQuery {
	o.ctx.Limit = &limit
	return o
}

func (o *postEntityQuery) Order(term ...post.OrderTerm) *postEntityQuery {
	o.order = append(o.order, term...)
	return o
}

func (o *postEntityQuery) Include(rels ...postEntityRel) *postEntityQuery {
	o.rels = append(o.rels, rels...)
	return o
}

// ToList returns the list of results of the query.
func (o *postEntityQuery) ToList(ctx context.Context) ([]*PostEntity, error) {
	return o.sqlAll(ctx)
}

// Single returns the single result of the query.
func (o *postEntityQuery) Single(ctx context.Context) (*PostEntity, error) {
	limit := 1
	o.ctx.Limit = &limit
	return o.sqlSingle(ctx)
}

func (o *postEntityQuery) sqlSingle(ctx context.Context) (*PostEntity, error) {
	var (
		spec = o.querySpec()
		res  *PostEntity
	)
	spec.Scan = func(rows dialect.Rows, fields []entitysql.ScannerField) error {
		e := o.config.New()
		switch e := e.(type) {
		case *PostEntity:
			builder := entitysql.NewScannerBuilder(o.scannerTotal + 1)
			builder.Append(0, e.scan(fields)...)
			for _, s := range o.scanner {
				e.createRel(builder, s)
			}
			if err := rows.Scan(builder.Flatten()...); err != nil {
				return err
			} else {
				res = e
				return nil
			}
		default:
			return entity.Err_0100030006
		}
	}
	if err := entitysql.NewQuery(ctx, o.config.Driver, spec); err != nil {
		return nil, err
	}
	if res != nil {
		if err := res.setUnchanged(); err != nil {
			return nil, err
		}
	}
	for _, r := range o.rels {
		r.reset()
	}
	return res, nil
}

func (o *postEntityQuery) sqlAll(ctx context.Context) ([]*PostEntity, error) {
	var (
		spec = o.querySpec()
		res  = []*PostEntity{}
	)
	spec.Scan = func(rows dialect.Rows, fields []entitysql.ScannerField) error {
		e := o.config.New()
		switch e := e.(type) {
		case *PostEntity:
			builder := entitysql.NewScannerBuilder(o.scannerTotal + 1)
			builder.Append(0, e.scan([]entitysql.ScannerField{})...)
			for _, s := range o.scanner {
				e.createRel(builder, s)
			}

			if err := rows.Scan(builder.Flatten()...); err != nil {
				return err
			} else {
				res = mergePostEntity(res, e)
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
	for _, r := range o.rels {
		rel := r
		rel.reset()
	}
	return res, nil
}

func (o *postEntityQuery) querySpec() *entitysql.QuerySpec {
	s := entitysql.NewQuerySpec(post.Entity, post.Columns)
	if o.ctx.Limit != nil {
		s.Limit = *o.ctx.Limit
	}
	if fields := o.ctx.Fields; len(fields) > 0 {
		s.Entity.Columns = make([]entitysql.FieldName, 0, len(fields))
		s.Entity.Columns = append(s.Entity.Columns, fields...)
	}
	if ps := o.predicates; len(ps) > 0 {
		s.Predicate = func(p *entitysql.Predicate, as string) {
			for _, f := range ps {
				f(p, as)
			}
		}
	}
	if rs := o.rels; len(rs) > 0 {
		s.Rels = make([]entitysql.Relation, 0, len(rs))
		s.Orders = append(s.Orders, post.ByPrimary)
		for _, r := range rs {
			rel := r
			s.Rels = append(s.Rels, func(s *entitysql.Selector) {
				o.scanner = o.addRels(s, s.Table(), rel, o.scanner)
			})
		}
	}
	for _, o := range o.order {
		s.Orders = append(s.Orders, func(order *entitysql.Order) {
			o.Apply(order)
		})
	}
	return s
}

func (o *postEntityQuery) addRels(s *entitysql.Selector, t *entitysql.SelectTable, rel rel, scanner []*internal.QueryScanner) []*internal.QueryScanner {
	desc, children, config := rel.Desc()
	join := entitysql.AddRelBySelector(s, t, desc)
	_, tableNum := join.GetAs()
	qs := internal.QueryScanner{Config: config, Children: []*internal.QueryScanner{}, TableNum: tableNum}
	scanner = append(scanner, &qs)
	o.scannerTotal++
	if len(children) > 0 {
		for _, c := range children {
			qs.Children = o.addRels(s, join, c, qs.Children)
		}
	}
	return scanner
}
