// Code generated by taurus_go/entity, DO NOT EDIT.

package entity

import (
	"context"
	"fmt"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/entitysql"

	"taurus_go_demo/entity/new/entity/author"
	"taurus_go_demo/entity/new/entity/blog"
	"taurus_go_demo/entity/new/entity/internal"
	"taurus_go_demo/entity/new/entity/post"
)

const UserTag = "UserTag"

// User  is an struct of the database
type User struct {
	*internal.Dialect
	tracker    entity.Tracker
	Authors    *authorEntityBuilder
	Blogs      *blogEntityBuilder
	FieldDemos *fieldDemoEntityBuilder
	// Geos Geo的类型测试
	Geos  *geoEntityBuilder
	Posts *postEntityBuilder
}

type userEntityFlag interface {
	isUserEntity()
}

func (e *AuthorEntity) isUserEntity()    {}
func (e *BlogEntity) isUserEntity()      {}
func (e *FieldDemoEntity) isUserEntity() {}
func (e *GeoEntity) isUserEntity()       {}
func (e *PostEntity) isUserEntity()      {}

// NewUser creates a new User instance.
func NewUser() (*User, error) {
	dialect, err := internal.NewDialect(UserTag)
	if err != nil {
		return nil, err
	}
	user := &User{
		Dialect: dialect,
		tracker: &entity.Tracking{},
	}
	user.init()
	return user, nil
}

// Close closes the database.
func (d *User) Close() error {
	return d.Driver.Close()
}

// Save saves all changes to the database.
func (d *User) Save(ctx context.Context) error {
	tx, err := d.Dialect.MayTx(ctx)
	if err != nil {
		return err
	}
	if err := func() error {
		for _, m := range d.tracker.Mutators() {
			if err := m.Exec(ctx, tx); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return entitysql.Rollback(tx, err)
	}
	return tx.Commit()
}

// Remove will remove the entity from the database. The changes will be saved when Save is called.
func (d *User) Remove(e userEntityFlag) error {
	switch e.(type) {
	case *AuthorEntity:
		d.Authors.Remove(e.(*AuthorEntity))
	case *BlogEntity:
		d.Blogs.Remove(e.(*BlogEntity))
	case *FieldDemoEntity:
		d.FieldDemos.Remove(e.(*FieldDemoEntity))
	case *GeoEntity:
		d.Geos.Remove(e.(*GeoEntity))
	case *PostEntity:
		d.Posts.Remove(e.(*PostEntity))
	default:
		return fmt.Errorf("database User does not support entity type %T", e)
	}
	return nil
}

func (d *User) init() {
	authorEntityConfig := newAuthorEntityConfig(d.Dialect)
	blogEntityConfig := newBlogEntityConfig(d.Dialect)
	fieldDemoEntityConfig := newFieldDemoEntityConfig(d.Dialect)
	geoEntityConfig := newGeoEntityConfig(d.Dialect)
	postEntityConfig := newPostEntityConfig(d.Dialect)

	d.Authors = newAuthorEntityBuilder(
		authorEntityConfig,
		d.tracker,
		*newPostEntityRelation(
			postEntityConfig,
			entitysql.RelationDesc{
				Orders: []entitysql.OrderFunc{
					post.ByPrimary,
				},
				To: entitysql.RelationTable{
					Table:   "author",
					Field:   "id",
					Columns: author.Columns,
				},
				Join: entitysql.RelationTable{
					Table:   "post",
					Field:   "author_id",
					Columns: post.Columns,
				},
			},
		),
	)
	d.tracker.Add(d.Authors)

	d.Blogs = newBlogEntityBuilder(
		blogEntityConfig,
		d.tracker,
		*newPostEntityRelation(
			postEntityConfig,
			entitysql.RelationDesc{
				Orders: []entitysql.OrderFunc{
					post.ByPrimary,
				},
				To: entitysql.RelationTable{
					Table:   "blog",
					Field:   "id",
					Columns: blog.Columns,
				},
				Join: entitysql.RelationTable{
					Table:   "post",
					Field:   "blog_id",
					Columns: post.Columns,
				},
			},
		),
	)
	d.tracker.Add(d.Blogs)

	d.FieldDemos = newFieldDemoEntityBuilder(fieldDemoEntityConfig, d.tracker)
	d.tracker.Add(d.FieldDemos)

	d.Geos = newGeoEntityBuilder(geoEntityConfig, d.tracker)
	d.tracker.Add(d.Geos)

	d.Posts = newPostEntityBuilder(
		postEntityConfig,
		d.tracker,
		*newBlogEntityRelation(
			blogEntityConfig,
			entitysql.RelationDesc{
				Orders: []entitysql.OrderFunc{
					blog.ByPrimary,
				},
				To: entitysql.RelationTable{
					Table:   "post",
					Field:   "blog_id",
					Columns: post.Columns,
				},
				Join: entitysql.RelationTable{
					Table:   "blog",
					Field:   "id",
					Columns: blog.Columns,
				},
			},
		),

		*newAuthorEntityRelation(
			authorEntityConfig,
			entitysql.RelationDesc{
				Orders: []entitysql.OrderFunc{
					author.ByPrimary,
				},
				To: entitysql.RelationTable{
					Table:   "post",
					Field:   "author_id",
					Columns: post.Columns,
				},
				Join: entitysql.RelationTable{
					Table:   "author",
					Field:   "id",
					Columns: author.Columns,
				},
			},
		),
	)
	d.tracker.Add(d.Posts)
}
