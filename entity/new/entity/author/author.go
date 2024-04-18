// Code generated by taurus_go, DO NOT EDIT.

package author

import (
	"github.com/yohobala/taurus_go/entity/entitysql"
)

const (
	Entity = "author"
)

var (
	FieldID = entitysql.Field{
		Name:     "id",
		Primary:  1,
		Default:  true,
		Required: true,
	}
	FieldName = entitysql.Field{
		Name:     "name",
		Primary:  0,
		Default:  false,
		Required: true,
	}
)

var (
	Columns = []entitysql.FieldName{
		FieldID.Name,
		FieldName.Name,
	}
)