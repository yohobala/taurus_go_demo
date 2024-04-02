// Code generated by taurus_go, DO NOT EDIT.

package blog

import (
	"context"
	"taurus_go_demo/entity/new/entity/internal"
	"time"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// BlogEntityBuilder is a builder for the BlogEntity entity.
//
// The builder is used to create, update, and delete BlogEntity entities.
type BlogEntityBuilder struct {
	config    *internal.Config
	mutations *mutations
	tracker   entity.Tracker

	// ID Blog primary key
	ID PredID

	UUID PredUUID

	Desc PredDesc

	CreatedTime PredCreatedTime
}

// NewBlogEntityBuilder creates a new BlogEntityBuilder.
func NewBlogEntityBuilder(c *internal.Config, t entity.Tracker) *BlogEntityBuilder {
	return &BlogEntityBuilder{
		config:    c,
		tracker:   t,
		mutations: newMutations(),
	}
}

// New creates a new UserEntity，and add it to the tracker.
// Required parameters are fields that have no default value but are required,
// and options are fields that can be left empty by calling WithFieldName.
func (b *BlogEntityBuilder) New(uuid string, options ...func(*BlogEntity)) (*BlogEntity, error) {
	e := New(b.config, b.mutations)
	return e.create(uuid, options...)
}

func (b *BlogEntityBuilder) Remove(e *BlogEntity) error {
	if e.config.Mutation == nil {
		return nil
	}
	return e.remove()
}

// First returns the first BlogEntity.
func (s *BlogEntityBuilder) First(ctx context.Context) (*BlogEntity, error) {
	query := s.initQuery()
	return query.First(ctx)
}

func (s *BlogEntityBuilder) ToList(ctx context.Context) ([]*BlogEntity, error) {
	query := s.initQuery()
	return query.ToList(ctx)
}

func (s *BlogEntityBuilder) Where(conditions ...func(*entitysql.Predicate)) *BlogEntityQuery {
	query := s.initQuery()
	return query.Where(conditions...)
}

// WithDesc sets the "desc" field of the BlogEntity.
func (s *BlogEntityBuilder) WithDesc(desc string) func(*BlogEntity) {
	return func(e *BlogEntity) {
		e.Desc.Set(desc)
	}
}

// WithCreatedTime sets the "created_time" field of the BlogEntity.
func (s *BlogEntityBuilder) WithCreatedTime(createdtime time.Time) func(*BlogEntity) {
	return func(e *BlogEntity) {
		e.CreatedTime.Set(createdtime)
	}
}

// Exec executes all the mutations for the BlogEntity.
func (s *BlogEntityBuilder) Exec(ctx context.Context, tx dialect.Tx) error {
	if len(s.mutations.Addeds) > 0 {
		e := s.mutations.Get(entity.Added)
		n := NewBlogEntityCreate(s.config, e...)
		if err := n.create(ctx, tx); err != nil {
			return err
		}
	}
	if len(s.mutations.Modifieds) > 0 {
		e := s.mutations.Get(entity.Modified)
		n := NewBlogEntityUpdate(s.config, e...)
		if err := n.update(ctx, tx); err != nil {
			return err
		}
	}
	if len(s.mutations.Deleteds) > 0 {
		e := s.mutations.Get(entity.Deleted)
		n := NewBlogEntityDelete(s.config, e...)
		if err := n.delete(ctx, tx); err != nil {
			return err
		}
	}
	return nil
}

func (s *BlogEntityBuilder) initQuery() *BlogEntityQuery {
	return NewBlogEntityQuery(s.config, s.tracker, s.mutations)
}

// mutations is a collection of BlogEntity mutation.
type mutations struct {
	Detacheds  map[string]*BlogEntity
	Unchangeds map[string]*BlogEntity
	Deleteds   map[string]*BlogEntity
	Modifieds  map[string]*BlogEntity
	Addeds     map[string]*BlogEntity
}

// newMutations creates a new mutations.
func newMutations() *mutations {
	return &mutations{
		Detacheds:  make(map[string]*BlogEntity),
		Unchangeds: make(map[string]*BlogEntity),
		Deleteds:   make(map[string]*BlogEntity),
		Modifieds:  make(map[string]*BlogEntity),
		Addeds:     make(map[string]*BlogEntity),
	}
}

// Get returns all the BlogEntity in the specified state.
func (ms *mutations) Get(state entity.EntityState) []*BlogEntity {
	switch state {
	case entity.Detached:
		s := make([]*BlogEntity, 0, len(ms.Detacheds))
		for _, m := range ms.Detacheds {
			s = append(s, m)
		}
		return s
	case entity.Unchanged:
		s := make([]*BlogEntity, 0, len(ms.Unchangeds))
		for _, m := range ms.Unchangeds {
			s = append(s, m)
		}
		return s
	case entity.Deleted:
		s := make([]*BlogEntity, 0, len(ms.Deleteds))
		for _, m := range ms.Deleteds {
			s = append(s, m)
		}
		return s
	case entity.Modified:
		s := make([]*BlogEntity, 0, len(ms.Modifieds))
		for _, m := range ms.Modifieds {
			s = append(s, m)
		}
		return s
	case entity.Added:
		s := make([]*BlogEntity, 0, len(ms.Addeds))
		for _, m := range ms.Addeds {
			s = append(s, m)
		}
		return s
	}
	return nil
}

// SetEntityState sets the state of the entity.
func (ms *mutations) SetEntityState(e *BlogEntity, state entity.EntityState) error {
	m := e.config.Mutation
	ms.set(e, state)
	if err := internal.SetEntityState(m, state); err != nil {
		return err
	}
	return nil
}

// ChangeEntityState attempts to set the desired entity state,
// but will not do so if the conditions are not met.
func (ms *mutations) ChangeEntityState(m *entity.Mutation, state entity.EntityState) {
	e := ms.getEntity(m)
	ms.set(e, state)
	if err := internal.SetEntityState(m, state); err != nil {
		return
	}
}

// getEntity returns the entity in the specified state.
func (ms *mutations) getEntity(m *entity.Mutation) *BlogEntity {
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
func (ms *mutations) set(e *BlogEntity, state entity.EntityState) {
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
