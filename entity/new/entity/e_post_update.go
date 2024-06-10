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

// postEntityUpdate is the update action for the postEntity.
type postEntityUpdate struct {
	config     *internal.Dialect
	ctx        *entitysql.QueryContext
	tracker    entity.Tracker
	es         []*PostEntity
	predicates [][]entitysql.PredicateFunc
	sets       []map[string]entitysql.CaseSpec
	total      int
	batchIndex []int
}

// newPostEntityUpdate creates a new postEntityUpdate.
func newPostEntityUpdate(c *internal.Dialect, es ...*PostEntity) *postEntityUpdate {
	return &postEntityUpdate{
		config:     c,
		ctx:        &entitysql.QueryContext{},
		es:         es,
		predicates: [][]entitysql.PredicateFunc{},
		batchIndex: []int{0},
	}
}

func (o *postEntityUpdate) update(ctx context.Context, tx dialect.Tx) error {
	return o.sqlUpdate(ctx, tx)
}

func (o *postEntityUpdate) sqlUpdate(ctx context.Context, tx dialect.Tx) error {
	var (
		spec, err = o.updateSpec()
		res       = o.es
		cursor    = 0
	)
	if err != nil {
		return err
	}
	spec.Scan = func(rows dialect.Rows, fields []entitysql.ScannerField) error {
		e := res[cursor]
		cursor++
		args := e.scan(fields)
		if err := rows.Scan(args...); err != nil {
			return err
		} else {
			res = append(res, e)
			return e.setUnchanged()
		}
	}
	return entitysql.NewUpdate(ctx, tx, spec)
}

func (o *postEntityUpdate) updateSpec() (*entitysql.UpdateSpec, error) {
	spec := entitysql.NewUpdateSpec(post.Entity, post.Columns)
	if len(o.predicates) != len(o.sets) {
		return nil, entity.Err_0100030005
	}
	if err := o.setEntity(spec); err != nil {
		return nil, err
	}
	o.mergeArgs(spec)
	return spec, nil
}

// setEntity 用于在updateSpec中设置[]*postEntity的配置，
// 一般来说这个setEntity里的entity都是通过状态追踪，自动添加的。
func (o *postEntityUpdate) setEntity(spec *entitysql.UpdateSpec) error {
	predID := &post.PredID{}
	num := 0
	for i, e := range o.es {
		fields := e.config.Mutation.Fields()
		if len(fields) == 0 {
			return entity.Err_0100030002.Sprintf(e.config.Tag)
		}
		o.predicates = append(o.predicates, []entitysql.PredicateFunc{})
		o.sets = append(o.sets, map[string]entitysql.CaseSpec{})
		// 因为判断过predicates和set长度，所以这里默认等长
		index := len(o.predicates) - 1
		if i > 0 {
			o.predicates[index] = append(o.predicates[index], entitysql.Or, predID.EQ(*e.ID.Get()))
		} else {
			o.predicates[index] = append(o.predicates[index], predID.EQ(*e.ID.Get()))
		}
		num++
		for _, f := range fields {
			switch f {
			case post.FieldID.Name.String():
				v, err := e.ID.Value(o.config.Driver.Dialect())
				if err != nil {
					return err
				}
				o.sets[index][post.FieldID.Name.String()] = entitysql.CaseSpec{
					Value: v,
					When:  predID.EQ(*e.ID.Get()),
				}
				num++
			case post.FieldContent.Name.String():
				v, err := e.Content.Value(o.config.Driver.Dialect())
				if err != nil {
					return err
				}
				o.sets[index][post.FieldContent.Name.String()] = entitysql.CaseSpec{
					Value: v,
					When:  predID.EQ(*e.ID.Get()),
				}
				num++
			case post.FieldBlogID.Name.String():
				v, err := e.BlogID.Value(o.config.Driver.Dialect())
				if err != nil {
					return err
				}
				o.sets[index][post.FieldBlogID.Name.String()] = entitysql.CaseSpec{
					Value: v,
					When:  predID.EQ(*e.ID.Get()),
				}
				num++
			case post.FieldAuthorID.Name.String():
				v, err := e.AuthorID.Value(o.config.Driver.Dialect())
				if err != nil {
					return err
				}
				o.sets[index][post.FieldAuthorID.Name.String()] = entitysql.CaseSpec{
					Value: v,
					When:  predID.EQ(*e.ID.Get()),
				}
				num++
			}
		}
		batchSize := *(entity.GetConfig().BatchSize)
		if (o.total+num)/batchSize > len(o.batchIndex) {
			o.batchIndex = append(o.batchIndex, len(o.predicates))
		} else {
			o.batchIndex[len(o.batchIndex)-1] = len(o.predicates)
		}
		o.total += num
	}
	return nil
}

func (o *postEntityUpdate) mergeArgs(spec *entitysql.UpdateSpec) {
	for i, end := range o.batchIndex {
		var begin int
		if i == 0 {
			begin = 0
		} else {
			begin = o.batchIndex[i-1]
		}
		pred := []entitysql.PredicateFunc{}
		set := map[string][]entitysql.CaseSpec{}
		for _, ps := range o.predicates[begin:end] {
			pred = append(pred, ps...)
		}
		for _, ss := range o.sets[begin:end] {
			for k, v := range ss {
				set[k] = append(set[k], v)
			}
		}
		spec.Predicate = append(spec.Predicate, func(p *entitysql.Predicate, as string) {
			for _, f := range pred {
				f(p, as)
			}
		})
		spec.Sets = append(spec.Sets, set)
	}
}
