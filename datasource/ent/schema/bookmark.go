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
		field.Int("user_id"),
		field.Int("site_id"),
		field.String("note").
			MaxLen(1000),
	}
}

// Edges of the Bookmark.
func (Bookmark) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("have_site", Site.Type).
			Ref("bookmark_from").
            Field("site_id").
            Required().
            Unique(),
		edge.From("owner", User.Type).
			Ref("bookmarks").
            Field("user_id").
            Required().
            Unique(),
        edge.To("tags", Tag.Type),
	}
}
