// Code generated by taurus_go/entity, DO NOT EDIT.

package geo_demo

import (
	"github.com/yohobala/taurus_go/encoding/geo"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

type PredID struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredID) EQ(id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldID.Name.String(), p.Builder.FindAs(Entity), id)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredID) NEQ(id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldID.Name.String(), p.Builder.FindAs(Entity), id)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredID) GT(id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldID.Name.String(), p.Builder.FindAs(Entity), id)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredID) GTE(id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldID.Name.String(), p.Builder.FindAs(Entity), id)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredID) LT(id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldID.Name.String(), p.Builder.FindAs(Entity), id)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredID) LTE(id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldID.Name.String(), p.Builder.FindAs(Entity), id)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredID) In(ids ...int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		p.In(FieldID.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredID) NotIn(ids ...int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		p.NotIn(FieldID.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredID) Like(id string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldID.Name.String(), p.Builder.FindAs(Entity), id)
	}
}

type PredPoint struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredPoint) EQ(point geo.Point) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldPoint.Name.String(), p.Builder.FindAs(Entity), point)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredPoint) NEQ(point geo.Point) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldPoint.Name.String(), p.Builder.FindAs(Entity), point)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredPoint) GT(point geo.Point) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldPoint.Name.String(), p.Builder.FindAs(Entity), point)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredPoint) GTE(point geo.Point) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldPoint.Name.String(), p.Builder.FindAs(Entity), point)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredPoint) LT(point geo.Point) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldPoint.Name.String(), p.Builder.FindAs(Entity), point)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredPoint) LTE(point geo.Point) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldPoint.Name.String(), p.Builder.FindAs(Entity), point)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredPoint) In(points ...geo.Point) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(points))
		for i := range v {
			v[i] = points[i]
		}
		p.In(FieldPoint.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredPoint) NotIn(points ...geo.Point) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(points))
		for i := range v {
			v[i] = points[i]
		}
		p.NotIn(FieldPoint.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredPoint) Like(point string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldPoint.Name.String(), p.Builder.FindAs(Entity), point)
	}
}

type PredLineString struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredLineString) EQ(line_string geo.LineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldLineString.Name.String(), p.Builder.FindAs(Entity), line_string)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredLineString) NEQ(line_string geo.LineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldLineString.Name.String(), p.Builder.FindAs(Entity), line_string)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredLineString) GT(line_string geo.LineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldLineString.Name.String(), p.Builder.FindAs(Entity), line_string)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredLineString) GTE(line_string geo.LineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldLineString.Name.String(), p.Builder.FindAs(Entity), line_string)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredLineString) LT(line_string geo.LineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldLineString.Name.String(), p.Builder.FindAs(Entity), line_string)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredLineString) LTE(line_string geo.LineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldLineString.Name.String(), p.Builder.FindAs(Entity), line_string)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredLineString) In(line_strings ...geo.LineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(line_strings))
		for i := range v {
			v[i] = line_strings[i]
		}
		p.In(FieldLineString.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredLineString) NotIn(line_strings ...geo.LineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(line_strings))
		for i := range v {
			v[i] = line_strings[i]
		}
		p.NotIn(FieldLineString.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredLineString) Like(line_string string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldLineString.Name.String(), p.Builder.FindAs(Entity), line_string)
	}
}

type PredPolygon struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredPolygon) EQ(polygon geo.Polygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldPolygon.Name.String(), p.Builder.FindAs(Entity), polygon)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredPolygon) NEQ(polygon geo.Polygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldPolygon.Name.String(), p.Builder.FindAs(Entity), polygon)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredPolygon) GT(polygon geo.Polygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldPolygon.Name.String(), p.Builder.FindAs(Entity), polygon)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredPolygon) GTE(polygon geo.Polygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldPolygon.Name.String(), p.Builder.FindAs(Entity), polygon)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredPolygon) LT(polygon geo.Polygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldPolygon.Name.String(), p.Builder.FindAs(Entity), polygon)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredPolygon) LTE(polygon geo.Polygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldPolygon.Name.String(), p.Builder.FindAs(Entity), polygon)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredPolygon) In(polygons ...geo.Polygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(polygons))
		for i := range v {
			v[i] = polygons[i]
		}
		p.In(FieldPolygon.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredPolygon) NotIn(polygons ...geo.Polygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(polygons))
		for i := range v {
			v[i] = polygons[i]
		}
		p.NotIn(FieldPolygon.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredPolygon) Like(polygon string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldPolygon.Name.String(), p.Builder.FindAs(Entity), polygon)
	}
}

type PredMultiPoint struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredMultiPoint) EQ(multi_point geo.MultiPoint) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldMultiPoint.Name.String(), p.Builder.FindAs(Entity), multi_point)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredMultiPoint) NEQ(multi_point geo.MultiPoint) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldMultiPoint.Name.String(), p.Builder.FindAs(Entity), multi_point)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredMultiPoint) GT(multi_point geo.MultiPoint) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldMultiPoint.Name.String(), p.Builder.FindAs(Entity), multi_point)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredMultiPoint) GTE(multi_point geo.MultiPoint) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldMultiPoint.Name.String(), p.Builder.FindAs(Entity), multi_point)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredMultiPoint) LT(multi_point geo.MultiPoint) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldMultiPoint.Name.String(), p.Builder.FindAs(Entity), multi_point)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredMultiPoint) LTE(multi_point geo.MultiPoint) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldMultiPoint.Name.String(), p.Builder.FindAs(Entity), multi_point)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredMultiPoint) In(multi_points ...geo.MultiPoint) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(multi_points))
		for i := range v {
			v[i] = multi_points[i]
		}
		p.In(FieldMultiPoint.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredMultiPoint) NotIn(multi_points ...geo.MultiPoint) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(multi_points))
		for i := range v {
			v[i] = multi_points[i]
		}
		p.NotIn(FieldMultiPoint.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredMultiPoint) Like(multi_point string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldMultiPoint.Name.String(), p.Builder.FindAs(Entity), multi_point)
	}
}

// IsNull returns a function that sets the predicate to check if the field is null.
// Operator "IS NULL"
func (f *PredMultiPoint) IsNull() entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.IsNull(FieldMultiPoint.Name.String(), p.Builder.FindAs(Entity))
	}
}

// NotNull returns a function that sets the predicate to check if the field is not null.
// Operator "IS NOT NULL"
func (f *PredMultiPoint) NotNull() entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NotNull(FieldMultiPoint.Name.String(), p.Builder.FindAs(Entity))
	}
}

type PredMultiLineString struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredMultiLineString) EQ(multi_line_string geo.MultiLineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldMultiLineString.Name.String(), p.Builder.FindAs(Entity), multi_line_string)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredMultiLineString) NEQ(multi_line_string geo.MultiLineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldMultiLineString.Name.String(), p.Builder.FindAs(Entity), multi_line_string)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredMultiLineString) GT(multi_line_string geo.MultiLineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldMultiLineString.Name.String(), p.Builder.FindAs(Entity), multi_line_string)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredMultiLineString) GTE(multi_line_string geo.MultiLineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldMultiLineString.Name.String(), p.Builder.FindAs(Entity), multi_line_string)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredMultiLineString) LT(multi_line_string geo.MultiLineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldMultiLineString.Name.String(), p.Builder.FindAs(Entity), multi_line_string)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredMultiLineString) LTE(multi_line_string geo.MultiLineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldMultiLineString.Name.String(), p.Builder.FindAs(Entity), multi_line_string)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredMultiLineString) In(multi_line_strings ...geo.MultiLineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(multi_line_strings))
		for i := range v {
			v[i] = multi_line_strings[i]
		}
		p.In(FieldMultiLineString.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredMultiLineString) NotIn(multi_line_strings ...geo.MultiLineString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(multi_line_strings))
		for i := range v {
			v[i] = multi_line_strings[i]
		}
		p.NotIn(FieldMultiLineString.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredMultiLineString) Like(multi_line_string string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldMultiLineString.Name.String(), p.Builder.FindAs(Entity), multi_line_string)
	}
}

// IsNull returns a function that sets the predicate to check if the field is null.
// Operator "IS NULL"
func (f *PredMultiLineString) IsNull() entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.IsNull(FieldMultiLineString.Name.String(), p.Builder.FindAs(Entity))
	}
}

// NotNull returns a function that sets the predicate to check if the field is not null.
// Operator "IS NOT NULL"
func (f *PredMultiLineString) NotNull() entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NotNull(FieldMultiLineString.Name.String(), p.Builder.FindAs(Entity))
	}
}

type PredMultiPolygon struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredMultiPolygon) EQ(multi_polygon geo.MultiPolygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldMultiPolygon.Name.String(), p.Builder.FindAs(Entity), multi_polygon)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredMultiPolygon) NEQ(multi_polygon geo.MultiPolygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldMultiPolygon.Name.String(), p.Builder.FindAs(Entity), multi_polygon)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredMultiPolygon) GT(multi_polygon geo.MultiPolygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldMultiPolygon.Name.String(), p.Builder.FindAs(Entity), multi_polygon)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredMultiPolygon) GTE(multi_polygon geo.MultiPolygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldMultiPolygon.Name.String(), p.Builder.FindAs(Entity), multi_polygon)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredMultiPolygon) LT(multi_polygon geo.MultiPolygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldMultiPolygon.Name.String(), p.Builder.FindAs(Entity), multi_polygon)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredMultiPolygon) LTE(multi_polygon geo.MultiPolygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldMultiPolygon.Name.String(), p.Builder.FindAs(Entity), multi_polygon)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredMultiPolygon) In(multi_polygons ...geo.MultiPolygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(multi_polygons))
		for i := range v {
			v[i] = multi_polygons[i]
		}
		p.In(FieldMultiPolygon.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredMultiPolygon) NotIn(multi_polygons ...geo.MultiPolygon) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(multi_polygons))
		for i := range v {
			v[i] = multi_polygons[i]
		}
		p.NotIn(FieldMultiPolygon.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredMultiPolygon) Like(multi_polygon string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldMultiPolygon.Name.String(), p.Builder.FindAs(Entity), multi_polygon)
	}
}

// IsNull returns a function that sets the predicate to check if the field is null.
// Operator "IS NULL"
func (f *PredMultiPolygon) IsNull() entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.IsNull(FieldMultiPolygon.Name.String(), p.Builder.FindAs(Entity))
	}
}

// NotNull returns a function that sets the predicate to check if the field is not null.
// Operator "IS NOT NULL"
func (f *PredMultiPolygon) NotNull() entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NotNull(FieldMultiPolygon.Name.String(), p.Builder.FindAs(Entity))
	}
}

type PredCircularString struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredCircularString) EQ(circular_string geo.CircularString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldCircularString.Name.String(), p.Builder.FindAs(Entity), circular_string)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredCircularString) NEQ(circular_string geo.CircularString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldCircularString.Name.String(), p.Builder.FindAs(Entity), circular_string)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredCircularString) GT(circular_string geo.CircularString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GT(FieldCircularString.Name.String(), p.Builder.FindAs(Entity), circular_string)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredCircularString) GTE(circular_string geo.CircularString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldCircularString.Name.String(), p.Builder.FindAs(Entity), circular_string)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredCircularString) LT(circular_string geo.CircularString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LT(FieldCircularString.Name.String(), p.Builder.FindAs(Entity), circular_string)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredCircularString) LTE(circular_string geo.CircularString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldCircularString.Name.String(), p.Builder.FindAs(Entity), circular_string)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredCircularString) In(circular_strings ...geo.CircularString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(circular_strings))
		for i := range v {
			v[i] = circular_strings[i]
		}
		p.In(FieldCircularString.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredCircularString) NotIn(circular_strings ...geo.CircularString) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(circular_strings))
		for i := range v {
			v[i] = circular_strings[i]
		}
		p.NotIn(FieldCircularString.Name.String(), p.Builder.FindAs(Entity), v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredCircularString) Like(circular_string string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.Like(FieldCircularString.Name.String(), p.Builder.FindAs(Entity), circular_string)
	}
}

// IsNull returns a function that sets the predicate to check if the field is null.
// Operator "IS NULL"
func (f *PredCircularString) IsNull() entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.IsNull(FieldCircularString.Name.String(), p.Builder.FindAs(Entity))
	}
}

// NotNull returns a function that sets the predicate to check if the field is not null.
// Operator "IS NOT NULL"
func (f *PredCircularString) NotNull() entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		p.NotNull(FieldCircularString.Name.String(), p.Builder.FindAs(Entity))
	}
}
