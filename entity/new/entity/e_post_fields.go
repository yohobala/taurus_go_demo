// Code generated by taurus_go/entity, DO NOT EDIT.

package entity

import (
	"taurus_go_demo/entity/new/entity/post"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/field"
)

// postID is ID field
type postID struct {
	field.IntStorage[int64]
	config *postEntityConfig
}

// newpostID creates a new postID
func newPostID(c *postEntityConfig) *postID {
	t := &postID{}
	t.config = c
	return t
}

// Set sets the value of ID field
func (t *postID) Set(v int64) {
	t.IntStorage.Set(v)
	if t.config.State() == entity.Unchanged || t.config.State() == entity.Modified {
		t.config.postEntityMutations.ChangeEntityState(t.config.Mutation, entity.Modified)
		t.config.Mutation.SetFields(post.FieldID.Name.String())
	}
}

// Get gets the value of ID field
func (t *postID) Get() *int64 {
	return t.IntStorage.Get()
}

// postContent is Content field
type postContent struct {
	field.StringStorage[string]
	config *postEntityConfig
}

// newpostContent creates a new postContent
func newPostContent(c *postEntityConfig) *postContent {
	t := &postContent{}
	t.config = c
	return t
}

// Set sets the value of Content field
func (t *postContent) Set(v string) {
	t.StringStorage.Set(v)
	if t.config.State() == entity.Unchanged || t.config.State() == entity.Modified {
		t.config.postEntityMutations.ChangeEntityState(t.config.Mutation, entity.Modified)
		t.config.Mutation.SetFields(post.FieldContent.Name.String())
	}
}

// Get gets the value of Content field
func (t *postContent) Get() *string {
	return t.StringStorage.Get()
}

// postBlogID is BlogID field
type postBlogID struct {
	field.IntStorage[int64]
	config *postEntityConfig
}

// newpostBlogID creates a new postBlogID
func newPostBlogID(c *postEntityConfig) *postBlogID {
	t := &postBlogID{}
	t.config = c
	return t
}

// Set sets the value of BlogID field
func (t *postBlogID) Set(v int64) {
	t.IntStorage.Set(v)
	if t.config.State() == entity.Unchanged || t.config.State() == entity.Modified {
		t.config.postEntityMutations.ChangeEntityState(t.config.Mutation, entity.Modified)
		t.config.Mutation.SetFields(post.FieldBlogID.Name.String())
	}
}

// Get gets the value of BlogID field
func (t *postBlogID) Get() *int64 {
	return t.IntStorage.Get()
}

// postAuthorID is AuthorID field
type postAuthorID struct {
	field.IntStorage[int64]
	config *postEntityConfig
}

// newpostAuthorID creates a new postAuthorID
func newPostAuthorID(c *postEntityConfig) *postAuthorID {
	t := &postAuthorID{}
	t.config = c
	return t
}

// Set sets the value of AuthorID field
func (t *postAuthorID) Set(v int64) {
	t.IntStorage.Set(v)
	if t.config.State() == entity.Unchanged || t.config.State() == entity.Modified {
		t.config.postEntityMutations.ChangeEntityState(t.config.Mutation, entity.Modified)
		t.config.Mutation.SetFields(post.FieldAuthorID.Name.String())
	}
}

// Get gets the value of AuthorID field
func (t *postAuthorID) Get() *int64 {
	return t.IntStorage.Get()
}
