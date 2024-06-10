// Code generated by taurus_go, DO NOT EDIT.

package entity

import (
	"fmt"
	"taurus_go_demo/entity/new/entity/field_demo"
	"taurus_go_demo/entity/new/entity/internal"
	"time"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

type FieldDemoEntity struct {
	internal.Entity
	config *fieldDemoEntityConfig

	// Int64F Int64 field
	Int64F *fieldDemoInt64F

	// VarF Varchar field
	VarF *fieldDemoVarF

	// BoolF Bool field
	BoolF *fieldDemoBoolF

	// IntArrayF Int array field
	IntArrayF *fieldDemoIntArrayF

	// Intarray2F Int array2 field
	Intarray2F *fieldDemoIntarray2F

	// BoolArrayF Bool array field
	BoolArrayF *fieldDemoBoolArrayF

	// TimeF Time field
	TimeF *fieldDemoTimeF

	// TimeArrayF Time array field
	TimeArrayF *fieldDemoTimeArrayF
}

// fieldDemoEntityConfig holds the configuration for the FieldDemoEntity.
type fieldDemoEntityConfig struct {
	internal.EntityConfig
	*internal.Dialect
	*entity.Mutation
	*fieldDemoEntityMutations
	name string
}

func newFieldDemoEntityConfig(c *internal.Dialect) *fieldDemoEntityConfig {
	return &fieldDemoEntityConfig{
		Dialect:                  c,
		fieldDemoEntityMutations: newFieldDemoEntityMutations(),
		name:                     "field_demo",
	}
}

// New creates a new FieldDemoEntity, but does not add tracking.
func (c *fieldDemoEntityConfig) New() internal.Entity {
	b := entity.NewMutation(entity.Detached)
	e := &FieldDemoEntity{
		config: &fieldDemoEntityConfig{
			Mutation:                 b,
			Dialect:                  c.Dialect,
			fieldDemoEntityMutations: c.fieldDemoEntityMutations,
		},
	}
	e.setState(entity.Detached)
	e.Int64F = newFieldDemoInt64F(e.config)
	e.VarF = newFieldDemoVarF(e.config)
	e.BoolF = newFieldDemoBoolF(e.config)
	e.IntArrayF = newFieldDemoIntArrayF(e.config)
	e.Intarray2F = newFieldDemoIntarray2F(e.config)
	e.BoolArrayF = newFieldDemoBoolArrayF(e.config)
	e.TimeF = newFieldDemoTimeF(e.config)
	e.TimeArrayF = newFieldDemoTimeArrayF(e.config)
	return e
}

func (c *fieldDemoEntityConfig) Desc() internal.EntityConfigDesc {
	return internal.EntityConfigDesc{
		Name: c.name,
	}
}

// String implements the fmt.Stringer interface.
func (e *FieldDemoEntity) String() string {
	return fmt.Sprintf("{ Int64F: %v, VarF: %v, BoolF: %v, IntArrayF: %v, Intarray2F: %v, BoolArrayF: %v, TimeF: %v, TimeArrayF: %v}",
		e.Int64F,
		e.VarF,
		e.BoolF,
		e.IntArrayF,
		e.Intarray2F,
		e.BoolArrayF,
		e.TimeF,
		e.TimeArrayF,
	)
}

// State returns the state of the FieldDemoEntity.
func (e *FieldDemoEntity) State() entity.EntityState {
	return e.config.State()
}

// remove removes the FieldDemoEntity from the database.
func (e *FieldDemoEntity) remove() error {
	return e.setState(entity.Deleted)
}

// create creates a new FieldDemoEntity and adds tracking.
func (e *FieldDemoEntity) create(int64_f int64, var_f string, bool_f bool, int_array_f []int64, int_array2_f [][]int64, bool_array_f []bool, time_f time.Time, time_array_f []time.Time, options ...func(*FieldDemoEntity)) (*FieldDemoEntity, error) {
	e.setState(entity.Added)
	e.Int64F.Set(int64_f)
	e.VarF.Set(var_f)
	e.BoolF.Set(bool_f)
	e.IntArrayF.Set(int_array_f)
	e.Intarray2F.Set(int_array2_f)
	e.BoolArrayF.Set(bool_array_f)
	e.TimeF.Set(time_f)
	e.TimeArrayF.Set(time_array_f)
	for _, option := range options {
		option(e)
	}
	return e, nil
}

// setUnchanged sets the state of the FieldDemoEntity to unchanged.
func (e *FieldDemoEntity) setUnchanged() error {
	return e.setState(entity.Unchanged)
}

// setState sets the state of the FieldDemoEntity.
func (e *FieldDemoEntity) setState(state entity.EntityState) error {
	return e.config.fieldDemoEntityMutations.SetEntityState(e, state)
}

// scan scans the database for the FieldDemoEntity.
func (e *FieldDemoEntity) scan(fields []entitysql.ScannerField) []any {
	if len(fields) == 0 {
		args := make([]any, len(field_demo.Columns))
		for i, c := range field_demo.Columns {
			switch c.String() {
			case field_demo.FieldInt64F.Name.String():
				v := e.Int64F
				v.Set(*new(int64))
				args[i] = v
			case field_demo.FieldVarF.Name.String():
				v := e.VarF
				v.Set(*new(string))
				args[i] = v
			case field_demo.FieldBoolF.Name.String():
				v := e.BoolF
				v.Set(*new(bool))
				args[i] = v
			case field_demo.FieldIntArrayF.Name.String():
				v := e.IntArrayF
				v.Set(*new([]int64))
				args[i] = v
			case field_demo.FieldIntarray2F.Name.String():
				v := e.Intarray2F
				v.Set(*new([][]int64))
				args[i] = v
			case field_demo.FieldBoolArrayF.Name.String():
				v := e.BoolArrayF
				v.Set(*new([]bool))
				args[i] = v
			case field_demo.FieldTimeF.Name.String():
				v := e.TimeF
				v.Set(*new(time.Time))
				args[i] = v
			case field_demo.FieldTimeArrayF.Name.String():
				v := e.TimeArrayF
				v.Set(*new([]time.Time))
				args[i] = v
			}
		}
		return args
	} else {
		args := make([]any, len(fields))
		for i := range fields {
			switch fields[i].String() {
			case field_demo.FieldInt64F.Name.String():
				v := e.Int64F
				v.Set(*new(int64))
				args[i] = v
			case field_demo.FieldVarF.Name.String():
				v := e.VarF
				v.Set(*new(string))
				args[i] = v
			case field_demo.FieldBoolF.Name.String():
				v := e.BoolF
				v.Set(*new(bool))
				args[i] = v
			case field_demo.FieldIntArrayF.Name.String():
				v := e.IntArrayF
				v.Set(*new([]int64))
				args[i] = v
			case field_demo.FieldIntarray2F.Name.String():
				v := e.Intarray2F
				v.Set(*new([][]int64))
				args[i] = v
			case field_demo.FieldBoolArrayF.Name.String():
				v := e.BoolArrayF
				v.Set(*new([]bool))
				args[i] = v
			case field_demo.FieldTimeF.Name.String():
				v := e.TimeF
				v.Set(*new(time.Time))
				args[i] = v
			case field_demo.FieldTimeArrayF.Name.String():
				v := e.TimeArrayF
				v.Set(*new([]time.Time))
				args[i] = v
			}
		}
		return args
	}
}

func (e *FieldDemoEntity) createRel(buidler *entitysql.ScannerBuilder, scanner *internal.QueryScanner) {
	switch scanner.Config.Desc().Name {
	}
}

func mergeFieldDemoEntity(es []*FieldDemoEntity, e *FieldDemoEntity) []*FieldDemoEntity {
	if e == nil {
		return es
	}
	if len(es) == 0 {
		es = append(es, e)
	} else {
		v := es[len(es)-1]

		if e.Int64F.Get() != nil {
			if v.Int64F.Get() != nil && *v.Int64F.Get() == *e.Int64F.Get() {
			} else {
				es = append(es, e)
			}
		}
	}
	return es
}