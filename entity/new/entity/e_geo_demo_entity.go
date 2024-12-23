// Code generated by taurus_go/entity, DO NOT EDIT.

package entity

import (
	"fmt"
	"taurus_go_demo/entity/new/entity/geo_demo"
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/datautil/geo"
	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

type GeoEntity struct {
	internal.Entity `json:"-"`
	config          *geoEntityConfig

	// ID 主键。
	ID *geoDemoID

	// Point 点
	Point *geoDemoPoint

	// LineString 线
	LineString *geoDemoLineString

	// Polygon 多边形
	Polygon *geoDemoPolygon

	// MultiPoint 多点
	MultiPoint *geoDemoMultiPoint

	// MultiLineString 多线
	MultiLineString *geoDemoMultiLineString

	// MultiPolygon 多多边形
	MultiPolygon *geoDemoMultiPolygon

	// CircularString 圆弧
	CircularString *geoDemoCircularString

	// PointJson 点
	PointJson *geoDemoPointJson

	// LineStringJson 线
	LineStringJson *geoDemoLineStringJson

	// PolygonJson 多边形
	PolygonJson *geoDemoPolygonJson

	// MultiPointJson 多点
	MultiPointJson *geoDemoMultiPointJson

	// MultiLineStringJson 多线
	MultiLineStringJson *geoDemoMultiLineStringJson

	// MultiPolygonJson 多多边形
	MultiPolygonJson *geoDemoMultiPolygonJson
}

// geoEntityConfig holds the configuration for the GeoEntity.
type geoEntityConfig struct {
	internal.EntityConfig
	*internal.Dialect
	*entity.Mutation
	*geoEntityMutations
	name string
}

func newGeoEntityConfig(c *internal.Dialect) *geoEntityConfig {
	return &geoEntityConfig{
		Dialect:            c,
		geoEntityMutations: newGeoEntityMutations(),
		name:               "geo_demo",
	}
}

// New creates a new GeoEntity, but does not add tracking.
func (c *geoEntityConfig) New() internal.Entity {
	b := entity.NewMutation(entity.Detached)
	e := &GeoEntity{
		config: &geoEntityConfig{
			Mutation:           b,
			Dialect:            c.Dialect,
			geoEntityMutations: c.geoEntityMutations,
		},
	}
	e.setState(entity.Detached)
	e.ID = newGeoDemoID(e.config)
	e.Point = newGeoDemoPoint(e.config)
	e.LineString = newGeoDemoLineString(e.config)
	e.Polygon = newGeoDemoPolygon(e.config)
	e.MultiPoint = newGeoDemoMultiPoint(e.config)
	e.MultiLineString = newGeoDemoMultiLineString(e.config)
	e.MultiPolygon = newGeoDemoMultiPolygon(e.config)
	e.CircularString = newGeoDemoCircularString(e.config)
	e.PointJson = newGeoDemoPointJson(e.config)
	e.LineStringJson = newGeoDemoLineStringJson(e.config)
	e.PolygonJson = newGeoDemoPolygonJson(e.config)
	e.MultiPointJson = newGeoDemoMultiPointJson(e.config)
	e.MultiLineStringJson = newGeoDemoMultiLineStringJson(e.config)
	e.MultiPolygonJson = newGeoDemoMultiPolygonJson(e.config)
	return e
}

func (c *geoEntityConfig) Desc() internal.EntityConfigDesc {
	return internal.EntityConfigDesc{
		Name: c.name,
	}
}

// String implements the fmt.Stringer interface.
func (e *GeoEntity) String() string {
	return fmt.Sprintf("{ ID: %v, Point: %v, LineString: %v, Polygon: %v, MultiPoint: %v, MultiLineString: %v, MultiPolygon: %v, CircularString: %v, PointJson: %v, LineStringJson: %v, PolygonJson: %v, MultiPointJson: %v, MultiLineStringJson: %v, MultiPolygonJson: %v}",
		e.ID,
		e.Point,
		e.LineString,
		e.Polygon,
		e.MultiPoint,
		e.MultiLineString,
		e.MultiPolygon,
		e.CircularString,
		e.PointJson,
		e.LineStringJson,
		e.PolygonJson,
		e.MultiPointJson,
		e.MultiLineStringJson,
		e.MultiPolygonJson,
	)
}

// State returns the state of the GeoEntity.
func (e *GeoEntity) State() entity.EntityState {
	return e.config.State()
}

// remove removes the GeoEntity from the database.
func (e *GeoEntity) remove() error {
	return e.setState(entity.Deleted)
}

// create creates a new GeoEntity and adds tracking.
func (e *GeoEntity) create(options ...func(*GeoEntity)) (*GeoEntity, error) {
	e.setState(entity.Added)
	for _, option := range options {
		option(e)
	}
	return e, nil
}

// setUnchanged sets the state of the GeoEntity to unchanged.
func (e *GeoEntity) setUnchanged() error {
	return e.setState(entity.Unchanged)
}

// setState sets the state of the GeoEntity.
func (e *GeoEntity) setState(state entity.EntityState) error {
	return e.config.geoEntityMutations.SetEntityState(e, state)
}

// scan scans the database for the GeoEntity.
func (e *GeoEntity) scan(fields []entitysql.ScannerField) []any {
	if len(fields) == 0 {
		args := make([]any, len(geo_demo.Columns))
		for i, c := range geo_demo.Columns {
			switch c.String() {
			case geo_demo.FieldID.Name.String():
				v := e.ID
				v.Set(*new(int64))
				args[i] = v
			case geo_demo.FieldPoint.Name.String():
				v := e.Point
				v.Set(*new(*geo.Point))
				args[i] = v
			case geo_demo.FieldLineString.Name.String():
				v := e.LineString
				v.Set(*new(*geo.LineString))
				args[i] = v
			case geo_demo.FieldPolygon.Name.String():
				v := e.Polygon
				v.Set(*new(*geo.Polygon))
				args[i] = v
			case geo_demo.FieldMultiPoint.Name.String():
				v := e.MultiPoint
				v.Set(*new(*geo.MultiPoint))
				args[i] = v
			case geo_demo.FieldMultiLineString.Name.String():
				v := e.MultiLineString
				v.Set(*new(*geo.MultiLineString))
				args[i] = v
			case geo_demo.FieldMultiPolygon.Name.String():
				v := e.MultiPolygon
				v.Set(*new(*geo.MultiPolygon))
				args[i] = v
			case geo_demo.FieldCircularString.Name.String():
				v := e.CircularString
				v.Set(*new(*geo.CircularString))
				args[i] = v
			case geo_demo.FieldPointJson.Name.String():
				v := e.PointJson
				v.Set(*new(*geo.Point))
				args[i] = v
			case geo_demo.FieldLineStringJson.Name.String():
				v := e.LineStringJson
				v.Set(*new(*geo.LineString))
				args[i] = v
			case geo_demo.FieldPolygonJson.Name.String():
				v := e.PolygonJson
				v.Set(*new(*geo.Polygon))
				args[i] = v
			case geo_demo.FieldMultiPointJson.Name.String():
				v := e.MultiPointJson
				v.Set(*new(*geo.MultiPoint))
				args[i] = v
			case geo_demo.FieldMultiLineStringJson.Name.String():
				v := e.MultiLineStringJson
				v.Set(*new(*geo.MultiLineString))
				args[i] = v
			case geo_demo.FieldMultiPolygonJson.Name.String():
				v := e.MultiPolygonJson
				v.Set(*new(*geo.MultiPolygon))
				args[i] = v
			}
		}
		return args
	} else {
		args := make([]any, len(fields))
		for i := range fields {
			switch fields[i].String() {
			case geo_demo.FieldID.Name.String():
				v := e.ID
				v.Set(*new(int64))
				args[i] = v
			case geo_demo.FieldPoint.Name.String():
				v := e.Point
				v.Set(*new(*geo.Point))
				args[i] = v
			case geo_demo.FieldLineString.Name.String():
				v := e.LineString
				v.Set(*new(*geo.LineString))
				args[i] = v
			case geo_demo.FieldPolygon.Name.String():
				v := e.Polygon
				v.Set(*new(*geo.Polygon))
				args[i] = v
			case geo_demo.FieldMultiPoint.Name.String():
				v := e.MultiPoint
				v.Set(*new(*geo.MultiPoint))
				args[i] = v
			case geo_demo.FieldMultiLineString.Name.String():
				v := e.MultiLineString
				v.Set(*new(*geo.MultiLineString))
				args[i] = v
			case geo_demo.FieldMultiPolygon.Name.String():
				v := e.MultiPolygon
				v.Set(*new(*geo.MultiPolygon))
				args[i] = v
			case geo_demo.FieldCircularString.Name.String():
				v := e.CircularString
				v.Set(*new(*geo.CircularString))
				args[i] = v
			case geo_demo.FieldPointJson.Name.String():
				v := e.PointJson
				v.Set(*new(*geo.Point))
				args[i] = v
			case geo_demo.FieldLineStringJson.Name.String():
				v := e.LineStringJson
				v.Set(*new(*geo.LineString))
				args[i] = v
			case geo_demo.FieldPolygonJson.Name.String():
				v := e.PolygonJson
				v.Set(*new(*geo.Polygon))
				args[i] = v
			case geo_demo.FieldMultiPointJson.Name.String():
				v := e.MultiPointJson
				v.Set(*new(*geo.MultiPoint))
				args[i] = v
			case geo_demo.FieldMultiLineStringJson.Name.String():
				v := e.MultiLineStringJson
				v.Set(*new(*geo.MultiLineString))
				args[i] = v
			case geo_demo.FieldMultiPolygonJson.Name.String():
				v := e.MultiPolygonJson
				v.Set(*new(*geo.MultiPolygon))
				args[i] = v
			}
		}
		return args
	}
}

func (e *GeoEntity) createRel(buidler *entitysql.ScannerBuilder, scanner *internal.QueryScanner) {
	switch scanner.Config.Desc().Name {
	}
}

func mergeGeoEntity(es []*GeoEntity, e *GeoEntity) []*GeoEntity {
	if e == nil {
		return es
	}
	if len(es) == 0 {
		es = append(es, e)
	} else {
		v := es[len(es)-1]

		if e.ID.Get() != nil {
			if v.ID.Get() != nil && *v.ID.Get() == *e.ID.Get() {
			} else {
				es = append(es, e)
			}
		}
	}
	return es
}
