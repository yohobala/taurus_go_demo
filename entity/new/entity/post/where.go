// Code generated by taurus_go, DO NOT EDIT.

package post

import "github.com/yohobala/taurus_go/entity/entitysql"

type PredID struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredID) EQ(id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.EQ(FieldID.Name.String(), as, id)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredID) NEQ(id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.NEQ(FieldID.Name.String(), as, id)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredID) GT(id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.GT(FieldID.Name.String(), as, id)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredID) GTE(id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.GTE(FieldID.Name.String(), as, id)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredID) LT(id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.LT(FieldID.Name.String(), as, id)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredID) LTE(id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.LTE(FieldID.Name.String(), as, id)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredID) In(ids ...int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		p.In(FieldID.Name.String(), as, v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredID) NotIn(ids ...int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		p.NotIn(FieldID.Name.String(), as, v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredID) Like(id string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.Like(FieldID.Name.String(), as, id)
	}
}

type PredContent struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredContent) EQ(content string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.EQ(FieldContent.Name.String(), as, content)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredContent) NEQ(content string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.NEQ(FieldContent.Name.String(), as, content)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredContent) GT(content string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.GT(FieldContent.Name.String(), as, content)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredContent) GTE(content string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.GTE(FieldContent.Name.String(), as, content)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredContent) LT(content string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.LT(FieldContent.Name.String(), as, content)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredContent) LTE(content string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.LTE(FieldContent.Name.String(), as, content)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredContent) In(contents ...string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		v := make([]any, len(contents))
		for i := range v {
			v[i] = contents[i]
		}
		p.In(FieldContent.Name.String(), as, v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredContent) NotIn(contents ...string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		v := make([]any, len(contents))
		for i := range v {
			v[i] = contents[i]
		}
		p.NotIn(FieldContent.Name.String(), as, v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredContent) Like(content string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.Like(FieldContent.Name.String(), as, content)
	}
}

type PredBlogID struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredBlogID) EQ(blog_id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.EQ(FieldBlogID.Name.String(), as, blog_id)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredBlogID) NEQ(blog_id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.NEQ(FieldBlogID.Name.String(), as, blog_id)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredBlogID) GT(blog_id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.GT(FieldBlogID.Name.String(), as, blog_id)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredBlogID) GTE(blog_id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.GTE(FieldBlogID.Name.String(), as, blog_id)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredBlogID) LT(blog_id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.LT(FieldBlogID.Name.String(), as, blog_id)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredBlogID) LTE(blog_id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.LTE(FieldBlogID.Name.String(), as, blog_id)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredBlogID) In(blog_ids ...int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		v := make([]any, len(blog_ids))
		for i := range v {
			v[i] = blog_ids[i]
		}
		p.In(FieldBlogID.Name.String(), as, v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredBlogID) NotIn(blog_ids ...int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		v := make([]any, len(blog_ids))
		for i := range v {
			v[i] = blog_ids[i]
		}
		p.NotIn(FieldBlogID.Name.String(), as, v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredBlogID) Like(blog_id string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.Like(FieldBlogID.Name.String(), as, blog_id)
	}
}

type PredAuthorID struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredAuthorID) EQ(author_id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.EQ(FieldAuthorID.Name.String(), as, author_id)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredAuthorID) NEQ(author_id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.NEQ(FieldAuthorID.Name.String(), as, author_id)
	}
}

// GT returns a function that sets the predicate to check if the field is greater than the given value.
// Operator ">"
func (f *PredAuthorID) GT(author_id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.GT(FieldAuthorID.Name.String(), as, author_id)
	}
}

// GTE returns a function that sets the predicate to check if the field is greater than or equal to the given value.
// Operator ">="
func (f *PredAuthorID) GTE(author_id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.GTE(FieldAuthorID.Name.String(), as, author_id)
	}
}

// LT returns a function that sets the predicate to check if the field is less than the given value.
// Operator "<"
func (f *PredAuthorID) LT(author_id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.LT(FieldAuthorID.Name.String(), as, author_id)
	}
}

// LTE returns a function that sets the predicate to check if the field is less than or equal to the given value.
// Operator "<="
func (f *PredAuthorID) LTE(author_id int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.LTE(FieldAuthorID.Name.String(), as, author_id)
	}
}

// In returns a function that sets the predicate to check if the field is in the given values.
// Operator "IN"
func (f *PredAuthorID) In(author_ids ...int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		v := make([]any, len(author_ids))
		for i := range v {
			v[i] = author_ids[i]
		}
		p.In(FieldAuthorID.Name.String(), as, v...)
	}
}

// NotIn returns a function that sets the predicate to check if the field is not in the given values.
// Operator "NOT IN"
func (f *PredAuthorID) NotIn(author_ids ...int64) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		v := make([]any, len(author_ids))
		for i := range v {
			v[i] = author_ids[i]
		}
		p.NotIn(FieldAuthorID.Name.String(), as, v...)
	}
}

// Like returns a function that sets the predicate to check if the field is like the given value.
// Operator "LIKE"
func (f *PredAuthorID) Like(author_id string) entitysql.PredicateFunc {
	return func(p *entitysql.Predicate, as string) {
		p.Like(FieldAuthorID.Name.String(), as, author_id)
	}
}
