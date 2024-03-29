package user

import (
	"context"
	"taurus_go_demo/entity/template/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

type UserEntityBuilder struct {
	*internal.Config
	*mutations
	tracker     entity.Tracker
	ID          PredID
	UUID        PredUUID
	Name        PredName
	Email       PredEmail
	CreatedTime PredCreatedTime
	Desc        PredDesc
}

func NewUserEntityBuilder(c *internal.Config, t entity.Tracker) *UserEntityBuilder {
	return &UserEntityBuilder{
		Config:    c,
		tracker:   t,
		mutations: newMutations(),
	}
}

// New creates a new UserEntity，and add it to the tracker.
// Required parameters are fields that have no default value but are required,
// and options are fields that can be left empty by calling WithFieldName.
func (b *UserEntityBuilder) New(uuid string, name string, options ...func(*UserEntity)) (*UserEntity, error) {
	e := New(b.Config, b.mutations)
	return e.create(uuid, name, options...)
}

func (b *UserEntityBuilder) Remove(e *UserEntity) error {
	if e.config.Mutation == nil {
		return nil
	}
	return e.remove()
}

// First returns the first UserEntity.
func (s *UserEntityBuilder) First(ctx context.Context) (*UserEntity, error) {
	query := s.initQuery()
	return query.First(ctx)
}

func (s *UserEntityBuilder) ToList(ctx context.Context) ([]*UserEntity, error) {
	query := s.initQuery()
	return query.ToList(ctx)
}

func (s *UserEntityBuilder) Where(conditions ...func(*entitysql.Predicate)) *UserEntityQuery {
	query := s.initQuery()
	return query.Where(conditions...)
}

// WithEmail sets the "email" field of the UserEntity.
func (s *UserEntityBuilder) WithEmail(email string) func(*UserEntity) {
	return func(e *UserEntity) {
		e.Email.Set(email)
	}
}

// WithDesc sets the "desc" field of the UserEntity.
func (s *UserEntityBuilder) WithDesc(desc string) func(*UserEntity) {
	return func(e *UserEntity) {
		e.Desc.Set(desc)
	}
}

// Exec executes all the mutations for the UserEntity.
func (s *UserEntityBuilder) Exec(ctx context.Context, tx dialect.Tx) error {
	if len(s.mutations.Addeds) > 0 {
		e := s.mutations.Get(entity.Added)
		n := NewUserEntityCreate(s.Config, e...)
		if err := n.create(ctx, tx); err != nil {
			return err
		}
	}
	if len(s.mutations.Modifieds) > 0 {
		e := s.mutations.Get(entity.Modified)
		n := NewUserEntityUpdate(s.Config, e...)
		if err := n.update(ctx, tx); err != nil {
			return err
		}
	}
	if len(s.mutations.Deleteds) > 0 {
		e := s.mutations.Get(entity.Deleted)
		n := NewUserEntityDelete(s.Config, e...)
		if err := n.delete(ctx, tx); err != nil {
			return err
		}
	}
	return nil
}

func (s *UserEntityBuilder) initQuery() *UserEntityQuery {
	return NewUserEntityQuery(s.Config, s.tracker, s.mutations)
}

type mutations struct {
	Detacheds  map[string]*UserEntity
	Unchangeds map[string]*UserEntity
	Deleteds   map[string]*UserEntity
	Modifieds  map[string]*UserEntity
	Addeds     map[string]*UserEntity
}

func newMutations() *mutations {
	return &mutations{
		Detacheds:  make(map[string]*UserEntity),
		Unchangeds: make(map[string]*UserEntity),
		Deleteds:   make(map[string]*UserEntity),
		Modifieds:  make(map[string]*UserEntity),
		Addeds:     make(map[string]*UserEntity),
	}
}

func (ms *mutations) Get(state entity.EntityState) []*UserEntity {
	switch state {
	case entity.Detached:
		s := make([]*UserEntity, 0, len(ms.Detacheds))
		for _, m := range ms.Detacheds {
			s = append(s, m)
		}
		return s
	case entity.Unchanged:
		s := make([]*UserEntity, 0, len(ms.Unchangeds))
		for _, m := range ms.Unchangeds {
			s = append(s, m)
		}
		return s
	case entity.Deleted:
		s := make([]*UserEntity, 0, len(ms.Deleteds))
		for _, m := range ms.Deleteds {
			s = append(s, m)
		}
		return s
	case entity.Modified:
		s := make([]*UserEntity, 0, len(ms.Modifieds))
		for _, m := range ms.Modifieds {
			s = append(s, m)
		}
		return s
	case entity.Added:
		s := make([]*UserEntity, 0, len(ms.Addeds))
		for _, m := range ms.Addeds {
			s = append(s, m)
		}
		return s
	}
	return nil
}

// SetEntityState
func (ms *mutations) SetEntityState(e *UserEntity, state entity.EntityState) error {
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
	if err := internal.SetEntityState(m, state); err != nil {
		return
	}
	ms.set(e, m.State())
}

func (ms *mutations) getEntity(m *entity.Mutation) *UserEntity {
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

// Set 设置实体的状态。
func (ms *mutations) set(e *UserEntity, state entity.EntityState) {
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
