package user

import (
	"context"
	"taurus_go_demo/entity/template/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// UserEntityUpdate is the update action for the UserEntity.
type UserEntityUpdate struct {
	*internal.Config
	ctx *entitysql.QueryContext
	es  []*UserEntity
	// predicates 用于存放每个entity的条件,和sets一一对应，
	// 如果没有条件，会添加一个空的[]func(*entitysql.Predicate)
	predicates [][]func(*entitysql.Predicate)
	sets       []map[string]entitysql.CaseSpec
	total      int
	batchIndex []int
}

// NewUserEntityUpdate creates a new UserEntityUpdate.
func NewUserEntityUpdate(c *internal.Config, es ...*UserEntity) *UserEntityUpdate {
	return &UserEntityUpdate{
		Config:     c,
		ctx:        &entitysql.QueryContext{},
		es:         es,
		predicates: [][]func(*entitysql.Predicate){},
	}
}

// Where 更新的条件，会默认添加主键的值作为条件，所以如果有添加主键的值，会被覆盖
// func (o *UserEntityUpdate) Where(predicates []func(*entitysql.Predicate)) *UserEntityUpdate {
// 	o.addPredicates(predicates...)
// 	return o
// }

func (o *UserEntityUpdate) update(ctx context.Context, tx dialect.Tx) error {
	return o.sqlUpdate(ctx, tx)
}

func (o *UserEntityUpdate) sqlUpdate(ctx context.Context, tx dialect.Tx) error {
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

func (o *UserEntityUpdate) updateSpec() (*entitysql.UpdateSpec, error) {
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

// setEntity 用于在updateSpec中设置[]*UserEntity的配置，
// 一般来说这个setEntity里的entity都是通过状态追踪，自动添加的。
func (o *UserEntityUpdate) setEntity(spec *entitysql.UpdateSpec) error {
	// TODO： 这里只要一个主键
	predID := &PredID{}
	num := 0
	for i, e := range o.es {
		fields := e.config.Mutation.Fields()
		if len(fields) == 0 {
			return entity.Err_0100030002.Sprintf(e.Name)
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
			case FieldName.Name.String():
				o.sets[index][FieldName.Name.String()] = entitysql.CaseSpec{
					Value: e.Name.Value(),
					When:  predID.EQ(*e.ID.Get()),
				}
				num++
			case FieldEmail.Name.String():
				o.sets[index][FieldEmail.Name.String()] = entitysql.CaseSpec{
					Value: e.Email.Value(),
					When:  predID.EQ(*e.ID.Get()),
				}
				num++
			case FieldCreatedTime.Name.String():
				o.sets[index][FieldCreatedTime.Name.String()] = entitysql.CaseSpec{
					Value: e.CreatedTime.Value(),
					When:  predID.EQ(*e.ID.Get()),
				}
				num++
			case FieldDesc.Name.String():
				o.sets[index][FieldDesc.Name.String()] = entitysql.CaseSpec{
					Value: e.Desc.Value(),
					When:  predID.EQ(*e.ID.Get()),
				}
				num++
			}
		}
		if (o.total+num)/entity.BatchSize > len(o.batchIndex) {
			o.batchIndex = append(o.batchIndex, len(o.predicates)-1)
		}
		o.total += num
	}
	return nil
}

func (o *UserEntityUpdate) mergeArgs(spec *entitysql.UpdateSpec) {
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
