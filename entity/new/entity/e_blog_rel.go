// Code generated by taurus_go/entity, DO NOT EDIT.

package entity

import (
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity/entitysql"
)

type blogEntityRelation struct {
	postEntityRel
	Config   internal.EntityConfig
	relation entitysql.RelationDesc
	children []rel
}

func newBlogEntityRelation(config internal.EntityConfig, desc entitysql.RelationDesc) *blogEntityRelation {
	return &blogEntityRelation{
		Config:   config,
		relation: desc,
		children: []rel{},
	}
}

func (r *blogEntityRelation) Where(predicates ...entitysql.PredicateFunc) *blogEntityRelation {
	r.relation.Predicates = append(r.relation.Predicates, predicates...)
	return r
}

func (r *blogEntityRelation) Include(rels ...blogEntityRel) *blogEntityRelation {
	// Create a slice of type Rel with the same length as r.children
	newRels := make([]rel, len(rels))
	for i, r := range rels {
		// Convert each blogEntityRel to Rel and store it in the new slice
		newRels[i] = rel(r)
	}
	r.children = append(r.children, newRels...)
	return r
}

func (r *blogEntityRelation) Desc() (entitysql.RelationDesc, []rel, internal.EntityConfig) {
	return r.relation, r.children, r.Config
}

func (r *blogEntityRelation) reset() {
	for _, child := range r.children {
		child.reset()
	}
	r.relation.Reset()
	r.children = []rel{}
}
