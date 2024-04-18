// Code generated by taurus_go, DO NOT EDIT.

package entity

import (
	"context"

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
	tracker entity.Tracker
	Authors *AuthorEntityBuilder
	Blogs   *BlogEntityBuilder
	Posts   *PostEntityBuilder
}

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

func (d *User) init() {
	authorEntityConfig := newAuthorEntityConfig(d.Dialect)
	blogEntityConfig := newBlogEntityConfig(d.Dialect)
	postEntityConfig := newPostEntityConfig(d.Dialect)

	d.Authors = newAuthorEntityBuilder(
		authorEntityConfig,
		d.tracker,
		*NewPostEntityRelation(
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
		*NewPostEntityRelation(
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

	d.Posts = newPostEntityBuilder(
		postEntityConfig,
		d.tracker,
		*NewBlogEntityRelation(
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

		*NewAuthorEntityRelation(
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
