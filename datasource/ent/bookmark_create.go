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
	"github.com/renoinn/bookmark-go/datasource/ent/tag"
	"github.com/renoinn/bookmark-go/datasource/ent/user"
)

// BookmarkCreate is the builder for creating a Bookmark entity.
type BookmarkCreate struct {
	config
	mutation *BookmarkMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (bc *BookmarkCreate) SetTitle(s string) *BookmarkCreate {
	bc.mutation.SetTitle(s)
	return bc
}

// SetNote sets the "note" field.
func (bc *BookmarkCreate) SetNote(s string) *BookmarkCreate {
	bc.mutation.SetNote(s)
	return bc
}

// AddSiteIDs adds the "site" edge to the Site entity by IDs.
func (bc *BookmarkCreate) AddSiteIDs(ids ...int) *BookmarkCreate {
	bc.mutation.AddSiteIDs(ids...)
	return bc
}

// AddSite adds the "site" edges to the Site entity.
func (bc *BookmarkCreate) AddSite(s ...*Site) *BookmarkCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bc.AddSiteIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (bc *BookmarkCreate) AddUserIDs(ids ...int) *BookmarkCreate {
	bc.mutation.AddUserIDs(ids...)
	return bc
}

// AddUser adds the "user" edges to the User entity.
func (bc *BookmarkCreate) AddUser(u ...*User) *BookmarkCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bc.AddUserIDs(ids...)
}

// AddTagIDs adds the "tag" edge to the Tag entity by IDs.
func (bc *BookmarkCreate) AddTagIDs(ids ...int) *BookmarkCreate {
	bc.mutation.AddTagIDs(ids...)
	return bc
}

// AddTag adds the "tag" edges to the Tag entity.
func (bc *BookmarkCreate) AddTag(t ...*Tag) *BookmarkCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bc.AddTagIDs(ids...)
}

// Mutation returns the BookmarkMutation object of the builder.
func (bc *BookmarkCreate) Mutation() *BookmarkMutation {
	return bc.mutation
}

// Save creates the Bookmark in the database.
func (bc *BookmarkCreate) Save(ctx context.Context) (*Bookmark, error) {
	var (
		err  error
		node *Bookmark
	)
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookmarkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, bc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Bookmark)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BookmarkMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookmarkCreate) SaveX(ctx context.Context) *Bookmark {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BookmarkCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BookmarkCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookmarkCreate) check() error {
	if _, ok := bc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Bookmark.title"`)}
	}
	if v, ok := bc.mutation.Title(); ok {
		if err := bookmark.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Bookmark.title": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Note(); !ok {
		return &ValidationError{Name: "note", err: errors.New(`ent: missing required field "Bookmark.note"`)}
	}
	if v, ok := bc.mutation.Note(); ok {
		if err := bookmark.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf(`ent: validator failed for field "Bookmark.note": %w`, err)}
		}
	}
	return nil
}

func (bc *BookmarkCreate) sqlSave(ctx context.Context) (*Bookmark, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BookmarkCreate) createSpec() (*Bookmark, *sqlgraph.CreateSpec) {
	var (
		_node = &Bookmark{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bookmark.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bookmark.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.Title(); ok {
		_spec.SetField(bookmark.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := bc.mutation.Note(); ok {
		_spec.SetField(bookmark.FieldNote, field.TypeString, value)
		_node.Note = value
	}
	if nodes := bc.mutation.SiteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bookmark.SiteTable,
			Columns: bookmark.SitePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   bookmark.UserTable,
			Columns: bookmark.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookmark.TagTable,
			Columns: bookmark.TagPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
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

// BookmarkCreateBulk is the builder for creating many Bookmark entities in bulk.
type BookmarkCreateBulk struct {
	config
	builders []*BookmarkCreate
}

// Save creates the Bookmark entities in the database.
func (bcb *BookmarkCreateBulk) Save(ctx context.Context) ([]*Bookmark, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Bookmark, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookmarkMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookmarkCreateBulk) SaveX(ctx context.Context) []*Bookmark {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BookmarkCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BookmarkCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
