// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/renoinn/bookmark-go/datasource/ent/bookmark"
	"github.com/renoinn/bookmark-go/datasource/ent/site"
	"github.com/renoinn/bookmark-go/datasource/ent/user"
)

// Bookmark is the model entity for the Bookmark schema.
type Bookmark struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// SiteID holds the value of the "site_id" field.
	SiteID int `json:"site_id,omitempty"`
	// Note holds the value of the "note" field.
	Note string `json:"note,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookmarkQuery when eager-loading is set.
	Edges BookmarkEdges `json:"edges"`
}

// BookmarkEdges holds the relations/edges for other nodes in the graph.
type BookmarkEdges struct {
	// HaveSite holds the value of the have_site edge.
	HaveSite *Site `json:"have_site,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// HaveSiteOrErr returns the HaveSite value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookmarkEdges) HaveSiteOrErr() (*Site, error) {
	if e.loadedTypes[0] {
		if e.HaveSite == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: site.Label}
		}
		return e.HaveSite, nil
	}
	return nil, &NotLoadedError{edge: "have_site"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookmarkEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e BookmarkEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[2] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bookmark) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bookmark.FieldID, bookmark.FieldUserID, bookmark.FieldSiteID:
			values[i] = new(sql.NullInt64)
		case bookmark.FieldNote:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Bookmark", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bookmark fields.
func (b *Bookmark) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bookmark.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case bookmark.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				b.UserID = int(value.Int64)
			}
		case bookmark.FieldSiteID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field site_id", values[i])
			} else if value.Valid {
				b.SiteID = int(value.Int64)
			}
		case bookmark.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				b.Note = value.String
			}
		}
	}
	return nil
}

// QueryHaveSite queries the "have_site" edge of the Bookmark entity.
func (b *Bookmark) QueryHaveSite() *SiteQuery {
	return (&BookmarkClient{config: b.config}).QueryHaveSite(b)
}

// QueryOwner queries the "owner" edge of the Bookmark entity.
func (b *Bookmark) QueryOwner() *UserQuery {
	return (&BookmarkClient{config: b.config}).QueryOwner(b)
}

// QueryTags queries the "tags" edge of the Bookmark entity.
func (b *Bookmark) QueryTags() *TagQuery {
	return (&BookmarkClient{config: b.config}).QueryTags(b)
}

// Update returns a builder for updating this Bookmark.
// Note that you need to call Bookmark.Unwrap() before calling this method if this Bookmark
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bookmark) Update() *BookmarkUpdateOne {
	return (&BookmarkClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Bookmark entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Bookmark) Unwrap() *Bookmark {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bookmark is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bookmark) String() string {
	var builder strings.Builder
	builder.WriteString("Bookmark(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", b.UserID))
	builder.WriteString(", ")
	builder.WriteString("site_id=")
	builder.WriteString(fmt.Sprintf("%v", b.SiteID))
	builder.WriteString(", ")
	builder.WriteString("note=")
	builder.WriteString(b.Note)
	builder.WriteByte(')')
	return builder.String()
}

// Bookmarks is a parsable slice of Bookmark.
type Bookmarks []*Bookmark

func (b Bookmarks) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
