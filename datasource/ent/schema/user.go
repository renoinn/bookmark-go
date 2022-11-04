package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// User Annotations
func (User) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "users"},
    }
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field {
        field.String("name"),
        field.String("email").Unique(),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
