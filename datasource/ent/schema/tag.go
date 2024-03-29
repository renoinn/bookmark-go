package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("name").
			NotEmpty(),
		field.Int("count").
			Default(0),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("tags").
			Field("user_id").
			Unique().
			Required(),
		edge.From("bookmarks", Bookmark.Type).
			Ref("tags"),
	}
}
