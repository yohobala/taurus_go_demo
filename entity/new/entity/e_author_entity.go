// Code generated by taurus_go/entity, DO NOT EDIT.

package entity

import (
	"fmt"
	"taurus_go_demo/entity/new/entity/author"
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/zodileap/taurus_go/entity"
	"github.com/zodileap/taurus_go/entity/entitysql"
)

type AuthorEntity struct {
	internal.Entity `json:"-"`
	config          *authorentityConfig
	Id              *author_ID // Id Author primary key
	Name            *author_Name

	Posts []*PostEntity `json:"posts,omitempty"`
}

// authorentityConfig holds the configuration for the AuthorEntity.
type authorentityConfig struct {
	internal.EntityConfig
	*internal.Dialect
	*entity.Mutation
	*authorentityMutations
	name string
}

func newAuthorEntityConfig(c *internal.Dialect) *authorentityConfig {
	return &authorentityConfig{
		Dialect:               c,
		authorentityMutations: newAuthorEntityMutations(),
		name:                  "author",
	}
}

// New creates a new AuthorEntity, but does not add tracking.
func (c *authorentityConfig) New() internal.Entity {
	b := entity.NewMutation(entity.Detached)
	e := &AuthorEntity{
		config: &authorentityConfig{
			Mutation:              b,
			Dialect:               c.Dialect,
			authorentityMutations: c.authorentityMutations,
		},
	}
	e.setState(entity.Detached)
	e.Id = newAuthor_ID(e.config)
	e.Name = newAuthor_Name(e.config)
	return e
}

func (c *authorentityConfig) Desc() internal.EntityConfigDesc {
	return internal.EntityConfigDesc{
		Name: c.name,
	}
}

// String implements the fmt.Stringer interface.
func (e *AuthorEntity) String() string {
	return fmt.Sprintf("{ Id: %v, Name: %v, Posts: %v}",
		e.Id,
		e.Name,
		e.Posts,
	)
}

// State returns the state of the AuthorEntity.
func (e *AuthorEntity) State() entity.EntityState {
	return e.config.State()
}

// remove removes the AuthorEntity from the database.
func (e *AuthorEntity) remove() error {
	return e.setState(entity.Deleted)
}

// create creates a new AuthorEntity and adds tracking.
func (e *AuthorEntity) create(name string, options ...func(*AuthorEntity)) (*AuthorEntity, error) {
	e.setState(entity.Added)
	e.Name.Set(name)
	for _, option := range options {
		option(e)
	}
	return e, nil
}

// setUnchanged sets the state of the AuthorEntity to unchanged.
func (e *AuthorEntity) setUnchanged() error {
	return e.setState(entity.Unchanged)
}

// setState sets the state of the AuthorEntity.
func (e *AuthorEntity) setState(state entity.EntityState) error {
	return e.config.authorentityMutations.SetEntityState(e, state)
}

// scan scans the database for the AuthorEntity.
func (e *AuthorEntity) scan(fields []entitysql.ScannerField) []any {
	if len(fields) == 0 {
		args := make([]any, len(author.Columns))
		for i, c := range author.Columns {
			switch c.String() {
			case author.FieldID.Name.String():
				v := e.Id
				v.Set(*new(int64))
				args[i] = v
			case author.FieldName.Name.String():
				v := e.Name
				v.Set(*new(string))
				args[i] = v
			}
		}
		return args
	} else {
		args := make([]any, len(fields))
		for i := range fields {
			switch fields[i].String() {
			case author.FieldID.Name.String():
				v := e.Id
				v.Set(*new(int64))
				args[i] = v
			case author.FieldName.Name.String():
				v := e.Name
				v.Set(*new(string))
				args[i] = v
			}
		}
		return args
	}
}

func (e *AuthorEntity) createRel(buidler *entitysql.ScannerBuilder, scanner *internal.QueryScanner) {
	switch scanner.Config.Desc().Name {
	case "post":
		postentity := scanner.Config.New().(*PostEntity)
		buidler.Append(scanner.TableNum-1, postentity.scan([]entitysql.ScannerField{})...)
		e.Posts = append(e.Posts, postentity)
		for _, c := range scanner.Children {
			postentity.createRel(buidler, c)
		}
	}
}

func mergeAuthorEntity(es []*AuthorEntity, e *AuthorEntity) []*AuthorEntity {
	if e == nil {
		return es
	}
	if len(es) == 0 {
		es = append(es, e)
	} else {
		v := es[len(es)-1]

		if v.Id.Get() == e.Id.Get() {
			for _, post := range e.Posts {
				posts := mergePostEntity(v.Posts, post)
				if len(posts) > 0 {
					v.Posts = posts
				}
			}
		} else {
			es = append(es, e)
		}
	}
	return es
}
