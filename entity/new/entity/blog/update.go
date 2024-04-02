// Code generated by taurus_go, DO NOT EDIT.

package blog

import (
	"context"
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// BlogEntityUpdate is the update action for the BlogEntity.
type BlogEntityUpdate struct {
	config     *internal.Config
	ctx        *entitysql.QueryContext
	tracker    entity.Tracker
	es         []*BlogEntity
	predicates [][]func(*entitysql.Predicate)
	sets       []map[string]entitysql.CaseSpec
	total      int
	batchIndex []int
}

// NewBlogEntityUpdate creates a new BlogEntityUpdate.
func NewBlogEntityUpdate(c *internal.Config, es ...*BlogEntity) *BlogEntityUpdate {
	return &BlogEntityUpdate{
		config:     c,
		ctx:        &entitysql.QueryContext{},
		es:         es,
		predicates: [][]func(*entitysql.Predicate){},
		batchIndex: []int{0},
	}
}

func (o *BlogEntityUpdate) update(ctx context.Context, tx dialect.Tx) error {
	return o.sqlUpdate(ctx, tx)
}

func (o *BlogEntityUpdate) sqlUpdate(ctx context.Context, tx dialect.Tx) error {
	var (
		spec, err = o.updateSpec()
		res       = o.es
		cursor    = 0
	)
	if err != nil {
		return err
	}
	spec.Scan = func(rows dialect.Rows, fields []entitysql.FieldName) error {
		e := res[cursor]
		cursor++
		if err := scan(e, fields, rows); err != nil {
			return err
		} else {
			res = append(res, e)
			return e.setUnchanged()
		}
	}
	return entitysql.NewUpdate(ctx, tx, spec)
}

func (o *BlogEntityUpdate) updateSpec() (*entitysql.UpdateSpec, error) {
	spec := entitysql.NewUpdateSpec(Entity, columns)
	if len(o.predicates) != len(o.sets) {
		return nil, entity.Err_0100030005
	}
	if err := o.setEntity(spec); err != nil {
		return nil, err
	}
	o.mergeArgs(spec)
	return spec, nil
}

// setEntity 用于在updateSpec中设置[]*BlogEntity的配置，
// 一般来说这个setEntity里的entity都是通过状态追踪，自动添加的。
func (o *BlogEntityUpdate) setEntity(spec *entitysql.UpdateSpec) error {
	predID := &PredID{}
	num := 0
	for i, e := range o.es {
		fields := e.config.Mutation.Fields()
		if len(fields) == 0 {
			return entity.Err_0100030002.Sprintf(e.config.Tag)
		}
		o.predicates = append(o.predicates, []func(*entitysql.Predicate){})
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
			case FieldID.Name.String():
				o.sets[index][FieldID.Name.String()] = entitysql.CaseSpec{
					Value: e.ID.Value(),
					When:  predID.EQ(*e.ID.Get()),
				}
				num++
			case FieldUUID.Name.String():
				o.sets[index][FieldUUID.Name.String()] = entitysql.CaseSpec{
					Value: e.UUID.Value(),
					When:  predID.EQ(*e.ID.Get()),
				}
				num++
			case FieldDesc.Name.String():
				o.sets[index][FieldDesc.Name.String()] = entitysql.CaseSpec{
					Value: e.Desc.Value(),
					When:  predID.EQ(*e.ID.Get()),
				}
				num++
			case FieldCreatedTime.Name.String():
				o.sets[index][FieldCreatedTime.Name.String()] = entitysql.CaseSpec{
					Value: e.CreatedTime.Value(),
					When:  predID.EQ(*e.ID.Get()),
				}
				num++
			}
		}
		if (o.total+num)/entity.BatchSize > len(o.batchIndex) {
			o.batchIndex = append(o.batchIndex, len(o.predicates))
		} else {
			o.batchIndex[len(o.batchIndex)-1] = len(o.predicates)
		}
		o.total += num
	}
	return nil
}

func (o *BlogEntityUpdate) mergeArgs(spec *entitysql.UpdateSpec) {
	for i, end := range o.batchIndex {
		var begin int
		if i == 0 {
			begin = 0
		} else {
			begin = o.batchIndex[i-1]
		}
		pred := []func(*entitysql.Predicate){}
		set := map[string][]entitysql.CaseSpec{}
		for _, ps := range o.predicates[begin:end] {
			pred = append(pred, ps...)
		}
		for _, ss := range o.sets[begin:end] {
			for k, v := range ss {
				set[k] = append(set[k], v)
			}
		}
		spec.Predicate = append(spec.Predicate, func(p *entitysql.Predicate) {
			for _, f := range pred {
				f(p)
			}
		})
		spec.Sets = append(spec.Sets, set)
	}
}
