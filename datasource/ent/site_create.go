// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/renoinn/bookmark-go/datasource/ent/bookmark"
	"github.com/renoinn/bookmark-go/datasource/ent/site"
)

// SiteCreate is the builder for creating a Site entity.
type SiteCreate struct {
	config
	mutation *SiteMutation
	hooks    []Hook
}

// SetURL sets the "url" field.
func (sc *SiteCreate) SetURL(s string) *SiteCreate {
	sc.mutation.SetURL(s)
	return sc
}

// SetTitle sets the "title" field.
func (sc *SiteCreate) SetTitle(s string) *SiteCreate {
	sc.mutation.SetTitle(s)
	return sc
}

// AddBookmarkFromIDs adds the "bookmark_from" edge to the Bookmark entity by IDs.
func (sc *SiteCreate) AddBookmarkFromIDs(ids ...int) *SiteCreate {
	sc.mutation.AddBookmarkFromIDs(ids...)
	return sc
}

// AddBookmarkFrom adds the "bookmark_from" edges to the Bookmark entity.
func (sc *SiteCreate) AddBookmarkFrom(b ...*Bookmark) *SiteCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return sc.AddBookmarkFromIDs(ids...)
}

// Mutation returns the SiteMutation object of the builder.
func (sc *SiteCreate) Mutation() *SiteMutation {
	return sc.mutation
}

// Save creates the Site in the database.
func (sc *SiteCreate) Save(ctx context.Context) (*Site, error) {
	var (
		err  error
		node *Site
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SiteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Site)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SiteMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SiteCreate) SaveX(ctx context.Context) *Site {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SiteCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SiteCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SiteCreate) check() error {
	if _, ok := sc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Site.url"`)}
	}
	if v, ok := sc.mutation.URL(); ok {
		if err := site.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Site.url": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Site.title"`)}
	}
	if v, ok := sc.mutation.Title(); ok {
		if err := site.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Site.title": %w`, err)}
		}
	}
	return nil
}

func (sc *SiteCreate) sqlSave(ctx context.Context) (*Site, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *SiteCreate) createSpec() (*Site, *sqlgraph.CreateSpec) {
	var (
		_node = &Site{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: site.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: site.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.URL(); ok {
		_spec.SetField(site.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := sc.mutation.Title(); ok {
		_spec.SetField(site.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if nodes := sc.mutation.BookmarkFromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   site.BookmarkFromTable,
			Columns: []string{site.BookmarkFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bookmark.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SiteCreateBulk is the builder for creating many Site entities in bulk.
type SiteCreateBulk struct {
	config
	builders []*SiteCreate
}

// Save creates the Site entities in the database.
func (scb *SiteCreateBulk) Save(ctx context.Context) ([]*Site, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Site, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SiteMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SiteCreateBulk) SaveX(ctx context.Context) []*Site {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SiteCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SiteCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
