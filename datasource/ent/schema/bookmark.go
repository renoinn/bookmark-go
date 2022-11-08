package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Bookmark holds the schema definition for the Bookmark entity.
type Bookmark struct {
	ent.Schema
}

// Fields of the Bookmark.
func (Bookmark) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty(),
		field.String("note").
			MaxLen(1000),
	}
}

// Edges of the Bookmark.
func (Bookmark) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("site", Site.Type).
			Ref("bookmark"),
		edge.From("user", User.Type).
			Ref("bookmark"),
		edge.To("tag", Tag.Type),
	}
}
