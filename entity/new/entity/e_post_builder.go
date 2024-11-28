// Code generated by taurus_go/entity, DO NOT EDIT.

package entity

import (
	"context"
	"taurus_go_demo/entity/new/entity/internal"
	"taurus_go_demo/entity/new/entity/post"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// postEntityBuilder is a builder for the PostEntity entity.
//
// The builder is used to create, update, and delete PostEntity entities.
type postEntityBuilder struct {
	config  *postEntityConfig
	tracker entity.Tracker

	// ID Post primary key
	ID post.PredID

	Content post.PredContent

	BlogID post.PredBlogID

	AuthorID post.PredAuthorID
	// ByID configures the query to sort results based on the 'id' field of the entity.
	// Sorting entities in ascending order by default.
	ByID post.ByID
	// ByContent configures the query to sort results based on the 'content' field of the entity.
	// Sorting entities in ascending order by default.
	ByContent post.ByContent
	// ByBlogID configures the query to sort results based on the 'blog_id' field of the entity.
	// Sorting entities in ascending order by default.
	ByBlogID post.ByBlogID
	// ByAuthorID configures the query to sort results based on the 'author_id' field of the entity.
	// Sorting entities in ascending order by default.
	ByAuthorID post.ByAuthorID

	// Blog configures the query to include data from the 'blog' table.
	// The method modifies the existing query to include a LEFT JOIN clause.
	// Blog be used as an argument to the Include method。
	Blog *blogEntityRelation

	// Author configures the query to include data from the 'author' table.
	// The method modifies the existing query to include a LEFT JOIN clause.
	// Author be used as an argument to the Include method。
	Author *authorEntityRelation
}

// newpostEntityBuilder creates a new postEntityBuilder.
func newPostEntityBuilder(c *postEntityConfig, t entity.Tracker, Blog blogEntityRelation, Author authorEntityRelation) *postEntityBuilder {
	return &postEntityBuilder{
		config:  c,
		tracker: t,
		Blog:    &Blog, Author: &Author,
	}
}

// Create creates a new UserEntity，and add it to the tracker.
// Required parameters are fields that have no default value but are required,
// and options are fields that can be left empty by calling WithFieldName.
func (b *postEntityBuilder) Create(content string, blog_id int64, author_id int64, options ...func(*PostEntity)) (*PostEntity, error) {
	e := b.config.New()
	switch t := e.(type) {
	case *PostEntity:
		return t.create(content, blog_id, author_id, options...)
	default:
		return nil, entity.Err_0100030006
	}
}

func (b *postEntityBuilder) Remove(e *PostEntity) error {
	if e.config.Mutation == nil {
		return nil
	}
	return e.remove()
}

// First returns the first PostEntity.
func (s *postEntityBuilder) First(ctx context.Context) (*PostEntity, error) {
	query := s.initQuery()
	return query.First(ctx)
}

func (s *postEntityBuilder) ToList(ctx context.Context) ([]*PostEntity, error) {
	query := s.initQuery()
	return query.ToList(ctx)
}

func (s *postEntityBuilder) Include(rels ...postEntityRel) *PostEntityQuery {
	query := s.initQuery()
	return query.Include(rels...)
}

func (s *postEntityBuilder) Order(o ...post.OrderTerm) *PostEntityQuery {
	query := s.initQuery()
	return query.Order(o...)
}

func (s *postEntityBuilder) Where(conditions ...entitysql.PredicateFunc) *PostEntityQuery {
	query := s.initQuery()
	return query.Where(conditions...)
}

// Exec executes all the postEntityMutations for the PostEntity.
func (s *postEntityBuilder) Exec(ctx context.Context, tx dialect.Tx) error {
	if len(s.config.postEntityMutations.Addeds) > 0 {
		e := s.config.postEntityMutations.Get(entity.Added)
		n := newPostEntityCreate(s.config.Dialect, e...)
		if err := n.create(ctx, tx); err != nil {
			return err
		}
	}
	if len(s.config.postEntityMutations.Modifieds) > 0 {
		e := s.config.postEntityMutations.Get(entity.Modified)
		n := newPostEntityUpdate(s.config.Dialect, e...)
		if err := n.update(ctx, tx); err != nil {
			return err
		}
	}
	if len(s.config.postEntityMutations.Deleteds) > 0 {
		e := s.config.postEntityMutations.Get(entity.Deleted)
		n := newPostEntityDelete(s.config.Dialect, e...)
		if err := n.delete(ctx, tx); err != nil {
			return err
		}
	}
	return nil
}

func (s *postEntityBuilder) initQuery() *PostEntityQuery {
	return newPostEntityQuery(s.config.Dialect, s.tracker, s.config.postEntityMutations)
}

// postEntityMutations is a collection of PostEntity mutation.
type postEntityMutations struct {
	Detacheds  map[string]*PostEntity
	Unchangeds map[string]*PostEntity
	Deleteds   map[string]*PostEntity
	Modifieds  map[string]*PostEntity
	Addeds     map[string]*PostEntity
}

// newPostEntityMutations creates a new mutations.
func newPostEntityMutations() *postEntityMutations {
	return &postEntityMutations{
		Detacheds:  make(map[string]*PostEntity),
		Unchangeds: make(map[string]*PostEntity),
		Deleteds:   make(map[string]*PostEntity),
		Modifieds:  make(map[string]*PostEntity),
		Addeds:     make(map[string]*PostEntity),
	}
}

// Get returns all the PostEntity in the specified state.
func (ms *postEntityMutations) Get(state entity.EntityState) []*PostEntity {
	switch state {
	case entity.Detached:
		s := make([]*PostEntity, 0, len(ms.Detacheds))
		for _, m := range ms.Detacheds {
			s = append(s, m)
		}
		return s
	case entity.Unchanged:
		s := make([]*PostEntity, 0, len(ms.Unchangeds))
		for _, m := range ms.Unchangeds {
			s = append(s, m)
		}
		return s
	case entity.Deleted:
		s := make([]*PostEntity, 0, len(ms.Deleteds))
		for _, m := range ms.Deleteds {
			s = append(s, m)
		}
		return s
	case entity.Modified:
		s := make([]*PostEntity, 0, len(ms.Modifieds))
		for _, m := range ms.Modifieds {
			s = append(s, m)
		}
		return s
	case entity.Added:
		s := make([]*PostEntity, 0, len(ms.Addeds))
		for _, m := range ms.Addeds {
			s = append(s, m)
		}
		return s
	}
	return nil
}

// SetEntityState sets the state of the entity.
func (ms *postEntityMutations) SetEntityState(e *PostEntity, state entity.EntityState) error {
	m := e.config.Mutation
	ms.set(e, state)
	if err := internal.SetEntityState(m, state); err != nil {
		return err
	}
	return nil
}

// ChangeEntityState attempts to set the desired entity state,
// but will not do so if the conditions are not met.
func (ms *postEntityMutations) ChangeEntityState(m *entity.Mutation, state entity.EntityState) {
	e := ms.getEntity(m)
	ms.set(e, state)
	if err := internal.SetEntityState(m, state); err != nil {
		return
	}
}

// getEntity returns the entity in the specified state.
func (ms *postEntityMutations) getEntity(m *entity.Mutation) *PostEntity {
	key := m.Key()
	switch m.State() {
	case entity.Detached:
		return ms.Detacheds[key]
	case entity.Unchanged:
		return ms.Unchangeds[key]
	case entity.Deleted:
		return ms.Deleteds[key]
	case entity.Modified:
		return ms.Modifieds[key]
	case entity.Added:
		return ms.Addeds[key]
	}
	return nil
}

// Set sets the entity in the specified state.
func (ms *postEntityMutations) set(e *PostEntity, state entity.EntityState) {
	m := e.config.Mutation
	key := m.Key()
	switch m.State() {
	case entity.Detached:
		delete(ms.Detacheds, key)
	case entity.Unchanged:
		delete(ms.Unchangeds, key)
	case entity.Deleted:
		delete(ms.Deleteds, key)
	case entity.Modified:
		delete(ms.Modifieds, key)
	case entity.Added:
		delete(ms.Addeds, key)
	}
	if state >= 0 {
		switch state {
		case entity.Detached:
			ms.Detacheds[key] = e
		case entity.Unchanged:
			ms.Unchangeds[key] = e
		case entity.Deleted:
			ms.Deleteds[key] = e
		case entity.Modified:
			ms.Modifieds[key] = e
		case entity.Added:
			ms.Addeds[key] = e
		}
	}
}
