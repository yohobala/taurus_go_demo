// Code generated by taurus_go/entity, DO NOT EDIT.

package entity

import (
	"context"
	"taurus_go_demo/entity/new/entity/author"
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// authorEntityBuilder is a builder for the AuthorEntity entity.
//
// The builder is used to create, update, and delete AuthorEntity entities.
type authorEntityBuilder struct {
	config  *authorEntityConfig
	tracker entity.Tracker

	// ID Author primary key
	ID author.PredID

	Name author.PredName
	// ByID configures the query to sort results based on the 'id' field of the entity.
	// Sorting entities in ascending order by default.
	ByID author.ByID
	// ByName configures the query to sort results based on the 'name' field of the entity.
	// Sorting entities in ascending order by default.
	ByName author.ByName

	// Posts configures the query to include data from the 'post' table.
	// The method modifies the existing query to include a LEFT JOIN clause.
	// Posts be used as an argument to the Include method。
	Posts *postEntityRelation
}

// newauthorEntityBuilder creates a new authorEntityBuilder.
func newAuthorEntityBuilder(c *authorEntityConfig, t entity.Tracker, Posts postEntityRelation) *authorEntityBuilder {
	return &authorEntityBuilder{
		config:  c,
		tracker: t,
		Posts:   &Posts,
	}
}

// Create creates a new UserEntity，and add it to the tracker.
// Required parameters are fields that have no default value but are required,
// and options are fields that can be left empty by calling WithFieldName.
func (b *authorEntityBuilder) Create(name string, options ...func(*AuthorEntity)) (*AuthorEntity, error) {
	e := b.config.New()
	switch t := e.(type) {
	case *AuthorEntity:
		return t.create(name, options...)
	default:
		return nil, entity.Err_0100030006
	}
}

func (b *authorEntityBuilder) Remove(e *AuthorEntity) error {
	if e.config.Mutation == nil {
		return nil
	}
	return e.remove()
}

// First returns the first AuthorEntity.
func (s *authorEntityBuilder) First(ctx context.Context) (*AuthorEntity, error) {
	query := s.initQuery()
	return query.First(ctx)
}

func (s *authorEntityBuilder) ToList(ctx context.Context) ([]*AuthorEntity, error) {
	query := s.initQuery()
	return query.ToList(ctx)
}

func (s *authorEntityBuilder) Include(rels ...authorEntityRel) *AuthorEntityQuery {
	query := s.initQuery()
	return query.Include(rels...)
}

func (s *authorEntityBuilder) Order(o ...author.OrderTerm) *AuthorEntityQuery {
	query := s.initQuery()
	return query.Order(o...)
}

func (s *authorEntityBuilder) Where(conditions ...entitysql.PredicateFunc) *AuthorEntityQuery {
	query := s.initQuery()
	return query.Where(conditions...)
}

// Exec executes all the authorEntityMutations for the AuthorEntity.
func (s *authorEntityBuilder) Exec(ctx context.Context, tx dialect.Tx) error {
	if len(s.config.authorEntityMutations.Addeds) > 0 {
		e := s.config.authorEntityMutations.Get(entity.Added)
		n := newAuthorEntityCreate(s.config.Dialect, e...)
		if err := n.create(ctx, tx); err != nil {
			return err
		}
	}
	if len(s.config.authorEntityMutations.Modifieds) > 0 {
		e := s.config.authorEntityMutations.Get(entity.Modified)
		n := newAuthorEntityUpdate(s.config.Dialect, e...)
		if err := n.update(ctx, tx); err != nil {
			return err
		}
	}
	if len(s.config.authorEntityMutations.Deleteds) > 0 {
		e := s.config.authorEntityMutations.Get(entity.Deleted)
		n := newAuthorEntityDelete(s.config.Dialect, e...)
		if err := n.delete(ctx, tx); err != nil {
			return err
		}
	}
	return nil
}

func (s *authorEntityBuilder) initQuery() *AuthorEntityQuery {
	return newAuthorEntityQuery(s.config.Dialect, s.tracker, s.config.authorEntityMutations)
}

// authorEntityMutations is a collection of AuthorEntity mutation.
type authorEntityMutations struct {
	Detacheds  map[string]*AuthorEntity
	Unchangeds map[string]*AuthorEntity
	Deleteds   map[string]*AuthorEntity
	Modifieds  map[string]*AuthorEntity
	Addeds     map[string]*AuthorEntity
}

// newAuthorEntityMutations creates a new mutations.
func newAuthorEntityMutations() *authorEntityMutations {
	return &authorEntityMutations{
		Detacheds:  make(map[string]*AuthorEntity),
		Unchangeds: make(map[string]*AuthorEntity),
		Deleteds:   make(map[string]*AuthorEntity),
		Modifieds:  make(map[string]*AuthorEntity),
		Addeds:     make(map[string]*AuthorEntity),
	}
}

// Get returns all the AuthorEntity in the specified state.
func (ms *authorEntityMutations) Get(state entity.EntityState) []*AuthorEntity {
	switch state {
	case entity.Detached:
		s := make([]*AuthorEntity, 0, len(ms.Detacheds))
		for _, m := range ms.Detacheds {
			s = append(s, m)
		}
		return s
	case entity.Unchanged:
		s := make([]*AuthorEntity, 0, len(ms.Unchangeds))
		for _, m := range ms.Unchangeds {
			s = append(s, m)
		}
		return s
	case entity.Deleted:
		s := make([]*AuthorEntity, 0, len(ms.Deleteds))
		for _, m := range ms.Deleteds {
			s = append(s, m)
		}
		return s
	case entity.Modified:
		s := make([]*AuthorEntity, 0, len(ms.Modifieds))
		for _, m := range ms.Modifieds {
			s = append(s, m)
		}
		return s
	case entity.Added:
		s := make([]*AuthorEntity, 0, len(ms.Addeds))
		for _, m := range ms.Addeds {
			s = append(s, m)
		}
		return s
	}
	return nil
}

// SetEntityState sets the state of the entity.
func (ms *authorEntityMutations) SetEntityState(e *AuthorEntity, state entity.EntityState) error {
	m := e.config.Mutation
	ms.set(e, state)
	if err := internal.SetEntityState(m, state); err != nil {
		return err
	}
	return nil
}

// ChangeEntityState attempts to set the desired entity state,
// but will not do so if the conditions are not met.
func (ms *authorEntityMutations) ChangeEntityState(m *entity.Mutation, state entity.EntityState) {
	e := ms.getEntity(m)
	ms.set(e, state)
	if err := internal.SetEntityState(m, state); err != nil {
		return
	}
}

// getEntity returns the entity in the specified state.
func (ms *authorEntityMutations) getEntity(m *entity.Mutation) *AuthorEntity {
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
func (ms *authorEntityMutations) set(e *AuthorEntity, state entity.EntityState) {
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
