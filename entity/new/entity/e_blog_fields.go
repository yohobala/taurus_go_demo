// Code generated by taurus_go/entity, DO NOT EDIT.

package entity

import (
	"taurus_go_demo/entity/new/entity/blog"
	"time"

	"github.com/zodileap/taurus_go/entity"
	"github.com/zodileap/taurus_go/entity/field"
)

// blog_ID is Id field
type blog_ID struct {
	field.IntStorage[int64]
	config *blogentityConfig
}

// newblog_ID creates a new blog_ID
func newBlog_ID(c *blogentityConfig) *blog_ID {
	t := &blog_ID{}
	t.config = c
	return t
}

// Set sets the value of Id field
func (t *blog_ID) Set(v int64) {
	t.IntStorage.Set(v)
	if t.config.State() == entity.Unchanged || t.config.State() == entity.Modified {
		t.config.blogentityMutations.ChangeEntityState(t.config.Mutation, entity.Modified)
		t.config.Mutation.SetFields(blog.FieldID.Name.String())
	}
}

// Get gets the value of Id field
//
// If the field is required, it returns the value type; otherwise, it returns a pointer type.
func (t *blog_ID) Get() int64 {
	return *t.IntStorage.Get()
}

// blog_UUID is Uuid field
type blog_UUID struct {
	field.StringStorage[string]
	config *blogentityConfig
}

// newblog_UUID creates a new blog_UUID
func newBlog_UUID(c *blogentityConfig) *blog_UUID {
	t := &blog_UUID{}
	t.config = c
	return t
}

// Set sets the value of Uuid field
func (t *blog_UUID) Set(v string) {
	t.StringStorage.Set(v)
	if t.config.State() == entity.Unchanged || t.config.State() == entity.Modified {
		t.config.blogentityMutations.ChangeEntityState(t.config.Mutation, entity.Modified)
		t.config.Mutation.SetFields(blog.FieldUUID.Name.String())
	}
}

// Get gets the value of Uuid field
//
// If the field is required, it returns the value type; otherwise, it returns a pointer type.
func (t *blog_UUID) Get() string {
	return *t.StringStorage.Get()
}

// blog_Description is Description field
type blog_Description struct {
	field.StringStorage[string]
	config *blogentityConfig
}

// newblog_Description creates a new blog_Description
func newBlog_Description(c *blogentityConfig) *blog_Description {
	t := &blog_Description{}
	t.config = c
	return t
}

// Set sets the value of Description field
func (t *blog_Description) Set(v string) {
	t.StringStorage.Set(v)
	if t.config.State() == entity.Unchanged || t.config.State() == entity.Modified {
		t.config.blogentityMutations.ChangeEntityState(t.config.Mutation, entity.Modified)
		t.config.Mutation.SetFields(blog.FieldDescription.Name.String())
	}
}

// Get gets the value of Description field
//
// If the field is required, it returns the value type; otherwise, it returns a pointer type.

func (t *blog_Description) Get() *string {
	return t.StringStorage.Get()
}

// blog_CreatedTime is CreatedTime field
type blog_CreatedTime struct {
	field.TimestampStorage[time.Time]
	config *blogentityConfig
}

// newblog_CreatedTime creates a new blog_CreatedTime
func newBlog_CreatedTime(c *blogentityConfig) *blog_CreatedTime {
	t := &blog_CreatedTime{}
	t.config = c
	return t
}

// Set sets the value of CreatedTime field
func (t *blog_CreatedTime) Set(v time.Time) {
	t.TimestampStorage.Set(v)
	if t.config.State() == entity.Unchanged || t.config.State() == entity.Modified {
		t.config.blogentityMutations.ChangeEntityState(t.config.Mutation, entity.Modified)
		t.config.Mutation.SetFields(blog.FieldCreatedTime.Name.String())
	}
}

// Get gets the value of CreatedTime field
//
// If the field is required, it returns the value type; otherwise, it returns a pointer type.

func (t *blog_CreatedTime) Get() *time.Time {
	return t.TimestampStorage.Get()
}
