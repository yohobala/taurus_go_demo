// Code generated by taurus_go/entity, DO NOT EDIT.

package field_demo

import (
	"time"

	"github.com/yohobala/taurus_go/entity/entitysql"
)

type PredInt64F struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredInt64F) EQ(int64_f int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldInt64F.Name.String(), p.Builder.FindAs(Entity), int64_f)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredInt64F) NEQ(int64_f int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldInt64F.Name.String(), p.Builder.FindAs(Entity), int64_f)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredInt64F) GT(int64_f int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldInt64F.Name.String(), p.Builder.FindAs(Entity), int64_f)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredInt64F) GTE(int64_f int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldInt64F.Name.String(), p.Builder.FindAs(Entity), int64_f)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredInt64F) LT(int64_f int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldInt64F.Name.String(), p.Builder.FindAs(Entity), int64_f)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredInt64F) LTE(int64_f int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldInt64F.Name.String(), p.Builder.FindAs(Entity), int64_f)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredInt64F) In(int64_fs ...int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(int64_fs))
		for i := range v {
			v[i] = int64_fs[i]
		}
		p.In(FieldInt64F.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredInt64F) NotIn(int64_fs ...int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(int64_fs))
		for i := range v {
			v[i] = int64_fs[i]
		}
		p.NotIn(FieldInt64F.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredInt64F) Like(int64_f string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldInt64F.Name.String(), p.Builder.FindAs(Entity), int64_f)
	}
}

type PredVarF struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredVarF) EQ(var_f string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldVarF.Name.String(), p.Builder.FindAs(Entity), var_f)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredVarF) NEQ(var_f string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldVarF.Name.String(), p.Builder.FindAs(Entity), var_f)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredVarF) GT(var_f string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldVarF.Name.String(), p.Builder.FindAs(Entity), var_f)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredVarF) GTE(var_f string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldVarF.Name.String(), p.Builder.FindAs(Entity), var_f)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredVarF) LT(var_f string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldVarF.Name.String(), p.Builder.FindAs(Entity), var_f)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredVarF) LTE(var_f string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldVarF.Name.String(), p.Builder.FindAs(Entity), var_f)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredVarF) In(var_fs ...string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(var_fs))
		for i := range v {
			v[i] = var_fs[i]
		}
		p.In(FieldVarF.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredVarF) NotIn(var_fs ...string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(var_fs))
		for i := range v {
			v[i] = var_fs[i]
		}
		p.NotIn(FieldVarF.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredVarF) Like(var_f string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldVarF.Name.String(), p.Builder.FindAs(Entity), var_f)
	}
}

type PredBoolF struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredBoolF) EQ(bool_f bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldBoolF.Name.String(), p.Builder.FindAs(Entity), bool_f)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredBoolF) NEQ(bool_f bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldBoolF.Name.String(), p.Builder.FindAs(Entity), bool_f)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredBoolF) GT(bool_f bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldBoolF.Name.String(), p.Builder.FindAs(Entity), bool_f)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredBoolF) GTE(bool_f bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldBoolF.Name.String(), p.Builder.FindAs(Entity), bool_f)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredBoolF) LT(bool_f bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldBoolF.Name.String(), p.Builder.FindAs(Entity), bool_f)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredBoolF) LTE(bool_f bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldBoolF.Name.String(), p.Builder.FindAs(Entity), bool_f)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredBoolF) In(bool_fs ...bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(bool_fs))
		for i := range v {
			v[i] = bool_fs[i]
		}
		p.In(FieldBoolF.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredBoolF) NotIn(bool_fs ...bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(bool_fs))
		for i := range v {
			v[i] = bool_fs[i]
		}
		p.NotIn(FieldBoolF.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredBoolF) Like(bool_f string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldBoolF.Name.String(), p.Builder.FindAs(Entity), bool_f)
	}
}

type PredIntArrayF struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredIntArrayF) EQ(int_array_f []int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldIntArrayF.Name.String(), p.Builder.FindAs(Entity), int_array_f)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredIntArrayF) NEQ(int_array_f []int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldIntArrayF.Name.String(), p.Builder.FindAs(Entity), int_array_f)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredIntArrayF) GT(int_array_f []int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldIntArrayF.Name.String(), p.Builder.FindAs(Entity), int_array_f)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredIntArrayF) GTE(int_array_f []int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldIntArrayF.Name.String(), p.Builder.FindAs(Entity), int_array_f)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredIntArrayF) LT(int_array_f []int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldIntArrayF.Name.String(), p.Builder.FindAs(Entity), int_array_f)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredIntArrayF) LTE(int_array_f []int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldIntArrayF.Name.String(), p.Builder.FindAs(Entity), int_array_f)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredIntArrayF) In(int_array_fs ...[]int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(int_array_fs))
		for i := range v {
			v[i] = int_array_fs[i]
		}
		p.In(FieldIntArrayF.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredIntArrayF) NotIn(int_array_fs ...[]int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(int_array_fs))
		for i := range v {
			v[i] = int_array_fs[i]
		}
		p.NotIn(FieldIntArrayF.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredIntArrayF) Like(int_array_f string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldIntArrayF.Name.String(), p.Builder.FindAs(Entity), int_array_f)
	}
}

type PredIntarray2F struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredIntarray2F) EQ(int_array2_f [][]int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldIntarray2F.Name.String(), p.Builder.FindAs(Entity), int_array2_f)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredIntarray2F) NEQ(int_array2_f [][]int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldIntarray2F.Name.String(), p.Builder.FindAs(Entity), int_array2_f)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredIntarray2F) GT(int_array2_f [][]int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldIntarray2F.Name.String(), p.Builder.FindAs(Entity), int_array2_f)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredIntarray2F) GTE(int_array2_f [][]int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldIntarray2F.Name.String(), p.Builder.FindAs(Entity), int_array2_f)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredIntarray2F) LT(int_array2_f [][]int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldIntarray2F.Name.String(), p.Builder.FindAs(Entity), int_array2_f)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredIntarray2F) LTE(int_array2_f [][]int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldIntarray2F.Name.String(), p.Builder.FindAs(Entity), int_array2_f)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredIntarray2F) In(int_array2_fs ...[][]int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(int_array2_fs))
		for i := range v {
			v[i] = int_array2_fs[i]
		}
		p.In(FieldIntarray2F.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredIntarray2F) NotIn(int_array2_fs ...[][]int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(int_array2_fs))
		for i := range v {
			v[i] = int_array2_fs[i]
		}
		p.NotIn(FieldIntarray2F.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredIntarray2F) Like(int_array2_f string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldIntarray2F.Name.String(), p.Builder.FindAs(Entity), int_array2_f)
	}
}

type PredBoolArrayF struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredBoolArrayF) EQ(bool_array_f []bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldBoolArrayF.Name.String(), p.Builder.FindAs(Entity), bool_array_f)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredBoolArrayF) NEQ(bool_array_f []bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldBoolArrayF.Name.String(), p.Builder.FindAs(Entity), bool_array_f)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredBoolArrayF) GT(bool_array_f []bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldBoolArrayF.Name.String(), p.Builder.FindAs(Entity), bool_array_f)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredBoolArrayF) GTE(bool_array_f []bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldBoolArrayF.Name.String(), p.Builder.FindAs(Entity), bool_array_f)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredBoolArrayF) LT(bool_array_f []bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldBoolArrayF.Name.String(), p.Builder.FindAs(Entity), bool_array_f)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredBoolArrayF) LTE(bool_array_f []bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldBoolArrayF.Name.String(), p.Builder.FindAs(Entity), bool_array_f)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredBoolArrayF) In(bool_array_fs ...[]bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(bool_array_fs))
		for i := range v {
			v[i] = bool_array_fs[i]
		}
		p.In(FieldBoolArrayF.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredBoolArrayF) NotIn(bool_array_fs ...[]bool) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(bool_array_fs))
		for i := range v {
			v[i] = bool_array_fs[i]
		}
		p.NotIn(FieldBoolArrayF.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredBoolArrayF) Like(bool_array_f string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldBoolArrayF.Name.String(), p.Builder.FindAs(Entity), bool_array_f)
	}
}

type PredTimeF struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredTimeF) EQ(time_f time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldTimeF.Name.String(), p.Builder.FindAs(Entity), time_f)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredTimeF) NEQ(time_f time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldTimeF.Name.String(), p.Builder.FindAs(Entity), time_f)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredTimeF) GT(time_f time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldTimeF.Name.String(), p.Builder.FindAs(Entity), time_f)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredTimeF) GTE(time_f time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldTimeF.Name.String(), p.Builder.FindAs(Entity), time_f)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredTimeF) LT(time_f time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldTimeF.Name.String(), p.Builder.FindAs(Entity), time_f)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredTimeF) LTE(time_f time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldTimeF.Name.String(), p.Builder.FindAs(Entity), time_f)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredTimeF) In(time_fs ...time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(time_fs))
		for i := range v {
			v[i] = time_fs[i]
		}
		p.In(FieldTimeF.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredTimeF) NotIn(time_fs ...time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(time_fs))
		for i := range v {
			v[i] = time_fs[i]
		}
		p.NotIn(FieldTimeF.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredTimeF) Like(time_f string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldTimeF.Name.String(), p.Builder.FindAs(Entity), time_f)
	}
}

type PredTimeArrayF struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredTimeArrayF) EQ(time_array_f []time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldTimeArrayF.Name.String(), p.Builder.FindAs(Entity), time_array_f)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredTimeArrayF) NEQ(time_array_f []time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldTimeArrayF.Name.String(), p.Builder.FindAs(Entity), time_array_f)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredTimeArrayF) GT(time_array_f []time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldTimeArrayF.Name.String(), p.Builder.FindAs(Entity), time_array_f)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredTimeArrayF) GTE(time_array_f []time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldTimeArrayF.Name.String(), p.Builder.FindAs(Entity), time_array_f)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredTimeArrayF) LT(time_array_f []time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldTimeArrayF.Name.String(), p.Builder.FindAs(Entity), time_array_f)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredTimeArrayF) LTE(time_array_f []time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldTimeArrayF.Name.String(), p.Builder.FindAs(Entity), time_array_f)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredTimeArrayF) In(time_array_fs ...[]time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(time_array_fs))
		for i := range v {
			v[i] = time_array_fs[i]
		}
		p.In(FieldTimeArrayF.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredTimeArrayF) NotIn(time_array_fs ...[]time.Time) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(time_array_fs))
		for i := range v {
			v[i] = time_array_fs[i]
		}
		p.NotIn(FieldTimeArrayF.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredTimeArrayF) Like(time_array_f string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldTimeArrayF.Name.String(), p.Builder.FindAs(Entity), time_array_f)
	}
}
