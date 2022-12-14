// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/renoinn/bookmark-go/datasource/ent/site"
)

// Site is the model entity for the Site schema.
type Site struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SiteQuery when eager-loading is set.
	Edges SiteEdges `json:"edges"`
}

// SiteEdges holds the relations/edges for other nodes in the graph.
type SiteEdges struct {
	// BookmarkFrom holds the value of the bookmark_from edge.
	BookmarkFrom []*Bookmark `json:"bookmark_from,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BookmarkFromOrErr returns the BookmarkFrom value or an error if the edge
// was not loaded in eager-loading.
func (e SiteEdges) BookmarkFromOrErr() ([]*Bookmark, error) {
	if e.loadedTypes[0] {
		return e.BookmarkFrom, nil
	}
	return nil, &NotLoadedError{edge: "bookmark_from"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Site) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case site.FieldID:
			values[i] = new(sql.NullInt64)
		case site.FieldURL, site.FieldTitle:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Site", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Site fields.
func (s *Site) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case site.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case site.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				s.URL = value.String
			}
		case site.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		}
	}
	return nil
}

// QueryBookmarkFrom queries the "bookmark_from" edge of the Site entity.
func (s *Site) QueryBookmarkFrom() *BookmarkQuery {
	return (&SiteClient{config: s.config}).QueryBookmarkFrom(s)
}

// Update returns a builder for updating this Site.
// Note that you need to call Site.Unwrap() before calling this method if this Site
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Site) Update() *SiteUpdateOne {
	return (&SiteClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Site entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Site) Unwrap() *Site {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Site is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Site) String() string {
	var builder strings.Builder
	builder.WriteString("Site(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("url=")
	builder.WriteString(s.URL)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(s.Title)
	builder.WriteByte(')')
	return builder.String()
}

// Sites is a parsable slice of Site.
type Sites []*Site

func (s Sites) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
