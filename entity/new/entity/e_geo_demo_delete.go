// Code generated by taurus_go/entity, DO NOT EDIT.

package entity

import (
	"context"
	"taurus_go_demo/entity/new/entity/geo_demo"
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/zodileap/taurus_go/entity"
	"github.com/zodileap/taurus_go/entity/dialect"
	"github.com/zodileap/taurus_go/entity/entitysql"
)

// GeoEntityDelete is the delete action for the GeoEntity.
type GeoEntityDelete struct {
	config     *internal.Dialect
	es         []*GeoEntity
	predicates []entitysql.PredicateFunc
}

// newGeoEntityDelete creates a new GeoEntityDelete.
func newGeoEntityDelete(c *internal.Dialect, es ...*GeoEntity) *GeoEntityDelete {
	return &GeoEntityDelete{
		config: c,
		es:     es,
	}
}

// Where adds a predicate to the delete action.
func (o *GeoEntityDelete) Where(predicates ...entitysql.PredicateFunc) *GeoEntityDelete {
	o.predicates = append(o.predicates, predicates...)
	return o
}

func (o *GeoEntityDelete) delete(ctx context.Context, tx dialect.Tx) error {
	return o.sqlDelete(ctx, tx)
}

func (o *GeoEntityDelete) sqlDelete(ctx context.Context, tx dialect.Tx) error {
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

func (o *GeoEntityDelete) deleteSpec() (*entitysql.DeleteSpec, error) {
	spec := entitysql.NewDeleteSpec(geo_demo.Entity)
	if ps := o.predicates; len(ps) > 0 {
		spec.Predicate = func(p *entitysql.Predicate) {
			for _, f := range ps {
				f(p)
			}
		}
	}
	predID := &geo_demo.PredID{}
	if o.predicates == nil {
		o.predicates = make([]entitysql.PredicateFunc, 0, len(o.es))
	}
	for i, e := range o.es {
		if i >= 1 {
			o.predicates = append(o.predicates, entitysql.Or)
		}
		o.predicates = append(o.predicates, predID.EQ(e.Id.Get()))
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
