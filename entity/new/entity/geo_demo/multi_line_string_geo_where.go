//Code generated by taurus_go/geo, DO NOT EDIT.

package geo_demo

import (
	"fmt"

	"github.com/yohobala/taurus_go/datautil/geo"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// Geo provides an interface for the Predicate structure to integrate the geospatial query capabilities of the PostGIS extension.
// Before using any PostGIS functions, they must be wrapped with the Geo function to ensure they are correctly utilized in SQL query construction.
// This method accepts a function implementing the geo.PostGISFunc interface, which specifies how to use PostGIS's geospatial operations,
// such as ST_Contains, ST_MakePoint, etc., to modify the Predicate object and generate a SQL statement that meets the requirements.
//
// Params:
//
//   - f: geo.PostGISFunc, a function that implements specific geospatial operations. This function specifies the specific operations on entitysql.Predicate,
//     to form a valid SQL statement.
//
// Returns:
//
//	0: Returns a function of type entitysql.PredicateFunc, which can be directly used to build Where query conditions in the database.
//
// Example:
//
//	g, err := db.Geos.Where(
//	    db.Geos.Point.Geo(
//	        db.Geos.Point.ST_Contains(
//	            db.Geos.Point.GeoColumn(),
//	            db.Geos.Point.ST_SetSRID(
//	                db.Geos.Point.ST_MakePoint(3, 4),
//	                new(geo.S4326),
//	            ),
//	        )),
//	)
//
// In this example, the ST_Contains query is wrapped by the Geo function to determine whether a specific geospatial coordinate contains a point.
// Sql: `SELECT "id", ST_AsText("point"), ST_AsText("line_string") FROM "geo_demo" AS "t1" WHERE ST_Contains("point", ST_SetSRID(ST_MakePoint(3, 4), 4326)) LIMIT 1`
//
// ExamplePath taurus_go_demo/entity/use/geo_test.go
func (p *PredMultiLineString) Geo(f geo.PostGISFunc) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate) {
		f.Pred(p, p.Builder.FindAs(Entity))
	}
}

// GeoColumn returns the PostGISFunc object for the column.
// This function is used to specify the column name in the PostGIS function.
// For example, in SQL: `Such as ST_Contains("point", ST_SetSRID(ST_MakePoint(3, 4), 4326))`.
func (p *PredMultiLineString) GeoColumn() geo.PostGISFunc {
	return geo.PostGISFunc{
		Entity: Entity,
		Column: string(FieldMultiLineString.Name),
	}
}

// ST_Contains is a PostGIS function.
// This function is used to determine whether the first geometry contains the second geometry.
// For example, in SQL: `ST_Contains("point", ST_SetSRID(ST_MakePoint(3, 4), 4326))`.
func (p *PredMultiLineString) ST_Contains(f1 geo.PostGISFunc, f2 geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Contains",
		Children: []geo.PostGISFunc{f1, f2},
	}
}

// ST_Crosses is a PostGIS function.
// ST_Crosses is used to determine whether the first geometry crosses the second geometry.
// ST_Crosses is generally applicable to line-to-line, line-to-surface, and surface-to-line scenarios, but are not suitable for surface-to-surface situations.
// For example, in SQL: `ST_Crosses("point", ST_SetSRID(ST_MakePoint(3, 4), 4326))`.
func (p *PredMultiLineString) ST_Crosses(f1 geo.PostGISFunc, f2 geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Crosses",
		Children: []geo.PostGISFunc{f1, f2},
	}
}

// ST_Disjoint is a PostGIS function.
// This function is used to determine whether the first geometry is disjoint from the second geometry.
// For example, in SQL: `ST_Disjoint("point", ST_SetSRID(ST_MakePoint(3, 4), 4326))`.
func (p *PredMultiLineString) ST_Disjoint(f1 geo.PostGISFunc, f2 geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Disjoint",
		Children: []geo.PostGISFunc{f1, f2},
	}
}

// ST_Equals is a PostGIS function.
// This function is used to determine whether the first geometry is equal to the second geometry.
// For example, in SQL: `ST_Equals("point", ST_SetSRID(ST_MakePoint(3, 4), 4326))`.
func (p *PredMultiLineString) ST_Equals(f1 geo.PostGISFunc, f2 geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Equals",
		Children: []geo.PostGISFunc{f1, f2},
	}
}

// ST_Intersects is a PostGIS function.
// This function is used to determine whether the first geometry intersects the second geometry.
// For example, in SQL: `ST_Intersects("point", ST_SetSRID(ST_MakePoint(3, 4), 4326))`.
func (p *PredMultiLineString) ST_Intersects(f1 geo.PostGISFunc, f2 geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Intersects",
		Children: []geo.PostGISFunc{f1, f2},
	}
}

// ST_Overlaps is a PostGIS function.
// This function is used to determine whether the first geometry overlaps the second geometry.
// For example, in SQL: `ST_Overlaps("point", ST_SetSRID(ST_MakePoint(3, 4), 4326))`.
func (p *PredMultiLineString) ST_Overlaps(f1 geo.PostGISFunc, f2 geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Overlaps",
		Children: []geo.PostGISFunc{f1, f2},
	}
}

// ST_Touches is a PostGIS function.
// This function is used to determine whether the first geometry touches the second geometry.
// For example, in SQL: `ST_Touches("point", ST_SetSRID(ST_MakePoint(3, 4), 4326))`.
func (p *PredMultiLineString) ST_Touches(f1 geo.PostGISFunc, f2 geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Touches",
		Children: []geo.PostGISFunc{f1, f2},
	}
}

// ST_Within is a PostGIS function.
// This function is used to determine whether the first geometry is within the second geometry.
// For example, in SQL: `ST_Within("point", ST_SetSRID(ST_MakePoint(3, 4), 4326))`.
func (p *PredMultiLineString) ST_Within(f1 geo.PostGISFunc, f2 geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Within",
		Children: []geo.PostGISFunc{f1, f2},
	}
}

// ST_Distance is a PostGIS function.
// This function is used to calculate the distance between two geometries.
// For example, in SQL: `ST_Distance("point", ST_SetSRID(ST_MakePoint(3, 4), 4326))`.
func (p *PredMultiLineString) ST_Distance(f1 geo.PostGISFunc, f2 geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Distance",
		Children: []geo.PostGISFunc{f1, f2},
	}
}

// ST_DWithin is a PostGIS function.
// This function is used to determine whether the first geometry is within a specified distance of the second geometry.
// For example, in SQL: `ST_DWithin("point", ST_SetSRID(ST_MakePoint(3, 4), 4326), 1)`.
func (p *PredMultiLineString) ST_DWithin(f1 geo.PostGISFunc, f2 geo.PostGISFunc, distance float64) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_DWithin",
		Value:    []any{distance},
		Children: []geo.PostGISFunc{f1, f2},
	}
}

// ST_Length is a PostGIS function.
// This function is used to calculate the length of a geometry.
// For example, in SQL: `ST_Length("lineString")`.
func (p *PredMultiLineString) ST_Length(f geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Length",
		Children: []geo.PostGISFunc{f},
	}
}

// ST_Area is a PostGIS function.
// This function is used to calculate the area of a geometry.
// For example, in SQL: `ST_Area("polygon")`.
func (p *PredMultiLineString) ST_Area(f geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Area",
		Children: []geo.PostGISFunc{f},
	}
}

// ST_Buffer is a PostGIS function.
// This function is used to create a buffer around a geometry.
// For example, in SQL: `ST_Buffer("point", 1)`.
func (p *PredMultiLineString) ST_Buffer(f geo.PostGISFunc, distance float64) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Buffer",
		Value:    []any{distance},
		Children: []geo.PostGISFunc{f},
	}
}

// ST_MakePoint is a PostGIS function.
// This function is used to create a point geometry.
// For example, in SQL: `ST_MakePoint(3, 4)`.
func (p *PredMultiLineString) ST_MakePoint(lng float64, lat float64) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:  "ST_MakePoint",
		Value: []any{lng, lat},
	}
}

// ST_MakeLine is a PostGIS function.
// This function is used to create a line geometry.
// For example, in SQL: `ST_MakeLine(ST_MakePoint(3, 4), ST_MakePoint(5, 6))`.
func (p *PredMultiLineString) ST_MakeLine(f1 geo.PostGISFunc, f2 geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_MakeLine",
		Children: []geo.PostGISFunc{f1, f2},
	}
}

// ST_MakePolygon is a PostGIS function.
// This function is used to create a polygon geometry.
// For example, in SQL: `ST_MakePolygon(ST_GeomFromText('LINESTRING(0 0, 4 0, 2 4, 0 0)')`.
func (p *PredMultiLineString) ST_MakePolygon(f geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_MakePolygon",
		Children: []geo.PostGISFunc{f},
	}
}

// ST_MakeEnvelope is a PostGIS function.
// This function is used to create a polygon geometry.
// For example, in SQL: `ST_MakeEnvelope(3, 4, 5, 6, 4326)`.
func (p *PredMultiLineString) ST_MakeEnvelope(xmin float64, ymin float64, xmax float64, ymax float64, srid geo.SRID) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:  "ST_MakeEnvelope",
		Value: []any{xmin, ymin, xmax, ymax, srid},
	}
}

// ST_Multi	 is a PostGIS function.
// This function is used to create a multi-geometry.
// For example, in SQL: `ST_Multi(ST_GeomFromText('LINESTRING(0 0, 4 0)','LINESTRING(4 0, 6 6)'))`.
func (p *PredMultiLineString) ST_Multi(f geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Multi",
		Children: []geo.PostGISFunc{f},
	}
}

// ST_Transform is a PostGIS function.
// This function is used to transform a geometry from one SRID to another.
// For example, in SQL: `ST_Transform(ST_MakePoint(3, 4), 4326)`.
func (p *PredMultiLineString) ST_Transform(f geo.PostGISFunc, srid geo.SRID) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Transform",
		Value:    []any{srid},
		Children: []geo.PostGISFunc{f},
	}
}

// ST_SetSRID is a PostGIS function.
// This function is used to set the SRID of a geometry.
// For example, in SQL: `ST_SetSRID(ST_MakePoint(3, 4), 4326)`.
func (p *PredMultiLineString) ST_SetSRID(f geo.PostGISFunc, srid geo.SRID) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_SetSRID",
		Value:    []any{srid.String()},
		Children: []geo.PostGISFunc{f},
	}
}

// ST_GeomFromText is a PostGIS function.
// This function is used to create a geometry from a WKT string.
// For example, in SQL: `ST_GeomFromText('POINT(3 4)')`.
func (p *PredMultiLineString) ST_GeomFromText(wkt string) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:  "ST_GeomFromText",
		Value: []any{fmt.Sprintf("'%v'", wkt)},
	}
}

// ST_Union is a PostGIS function.
// This function is used to merge multiple geometries into a single geometry.
// For example, in SQL: `ST_Union("point", "lineString")`.
func (p *PredMultiLineString) ST_Union(f1 geo.PostGISFunc, f2 geo.PostGISFunc) geo.PostGISFunc {
	return geo.PostGISFunc{
		Name:     "ST_Union",
		Children: []geo.PostGISFunc{f1, f2},
	}
}
