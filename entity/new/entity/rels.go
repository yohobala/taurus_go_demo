// Code generated by taurus_go/entity, DO NOT EDIT.

package entity

import (
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity/entitysql"
)

type rel interface {
	Desc() (entitysql.RelationDesc, []rel, internal.EntityConfig)
	reset()
}

type authorEntityRel interface {
	rel
}

type blogEntityRel interface {
	rel
}

type fieldDemoEntityRel interface {
	rel
}

type geoEntityRel interface {
	rel
}

type postEntityRel interface {
	rel
}
