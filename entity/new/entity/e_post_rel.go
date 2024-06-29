// Code generated by taurus_go/entity, DO NOT EDIT.

package entity

import (
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity/entitysql"
)

type postEntityRelation struct {
	blogEntityRel

	authorEntityRel
	Config   internal.EntityConfig
	relation entitysql.RelationDesc
	children []rel
}

func newPostEntityRelation(config internal.EntityConfig, desc entitysql.RelationDesc) *postEntityRelation {
	return &postEntityRelation{
		Config:   config,
		relation: desc,
		children: []rel{},
	}
}

func (r *postEntityRelation) Where(predicates ...entitysql.PredicateFunc) *postEntityRelation {
	r.relation.Predicates = append(r.relation.Predicates, predicates...)
	return r
}

func (r *postEntityRelation) Include(rels ...postEntityRel) *postEntityRelation {
	// Create a slice of type Rel with the same length as r.children
	newRels := make([]rel, len(rels))
	for i, r := range rels {
		// Convert each postEntityRel to Rel and store it in the new slice
		newRels[i] = rel(r)
	}
	r.children = append(r.children, newRels...)
	return r
}

func (r *postEntityRelation) Desc() (entitysql.RelationDesc, []rel, internal.EntityConfig) {
	return r.relation, r.children, r.Config
}

func (r *postEntityRelation) reset() {
	for _, child := range r.children {
		child.reset()
	}
	r.relation.Reset()
	r.children = []rel{}
}
