package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Int("count").
			Default(0),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return nil
	/*
		return []ent.Edge{
			edge.From("bookmark", Bookmark.Type).
				Ref("tag"),
			edge.To("user", User.Type),
		}
	*/
}
