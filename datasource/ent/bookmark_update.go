// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/renoinn/bookmark-go/datasource/ent/bookmark"
	"github.com/renoinn/bookmark-go/datasource/ent/predicate"
	"github.com/renoinn/bookmark-go/datasource/ent/tag"
	"github.com/renoinn/bookmark-go/datasource/ent/user"
)

// BookmarkUpdate is the builder for updating Bookmark entities.
type BookmarkUpdate struct {
	config
	hooks    []Hook
	mutation *BookmarkMutation
}

// Where appends a list predicates to the BookmarkUpdate builder.
func (bu *BookmarkUpdate) Where(ps ...predicate.Bookmark) *BookmarkUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUserID sets the "user_id" field.
func (bu *BookmarkUpdate) SetUserID(i int) *BookmarkUpdate {
	bu.mutation.SetUserID(i)
	return bu
}

// SetURL sets the "url" field.
func (bu *BookmarkUpdate) SetURL(s string) *BookmarkUpdate {
	bu.mutation.SetURL(s)
	return bu
}

// SetTitle sets the "title" field.
func (bu *BookmarkUpdate) SetTitle(s string) *BookmarkUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetNote sets the "note" field.
func (bu *BookmarkUpdate) SetNote(s string) *BookmarkUpdate {
	bu.mutation.SetNote(s)
	return bu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (bu *BookmarkUpdate) SetOwnerID(id int) *BookmarkUpdate {
	bu.mutation.SetOwnerID(id)
	return bu
}

// SetOwner sets the "owner" edge to the User entity.
func (bu *BookmarkUpdate) SetOwner(u *User) *BookmarkUpdate {
	return bu.SetOwnerID(u.ID)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (bu *BookmarkUpdate) AddTagIDs(ids ...int) *BookmarkUpdate {
	bu.mutation.AddTagIDs(ids...)
	return bu
}

// AddTags adds the "tags" edges to the Tag entity.
func (bu *BookmarkUpdate) AddTags(t ...*Tag) *BookmarkUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.AddTagIDs(ids...)
}

// Mutation returns the BookmarkMutation object of the builder.
func (bu *BookmarkUpdate) Mutation() *BookmarkMutation {
	return bu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (bu *BookmarkUpdate) ClearOwner() *BookmarkUpdate {
	bu.mutation.ClearOwner()
	return bu
}

// ClearTags clears all "tags" edges to the Tag entity.
func (bu *BookmarkUpdate) ClearTags() *BookmarkUpdate {
	bu.mutation.ClearTags()
	return bu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (bu *BookmarkUpdate) RemoveTagIDs(ids ...int) *BookmarkUpdate {
	bu.mutation.RemoveTagIDs(ids...)
	return bu
}

// RemoveTags removes "tags" edges to Tag entities.
func (bu *BookmarkUpdate) RemoveTags(t ...*Tag) *BookmarkUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.RemoveTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookmarkUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		if err = bu.check(); err != nil {
			return 0, err
		}
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookmarkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bu.check(); err != nil {
				return 0, err
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookmarkUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookmarkUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookmarkUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BookmarkUpdate) check() error {
	if v, ok := bu.mutation.URL(); ok {
		if err := bookmark.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Bookmark.url": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Title(); ok {
		if err := bookmark.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Bookmark.title": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Note(); ok {
		if err := bookmark.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf(`ent: validator failed for field "Bookmark.note": %w`, err)}
		}
	}
	if _, ok := bu.mutation.OwnerID(); bu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Bookmark.owner"`)
	}
	return nil
}

func (bu *BookmarkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bookmark.Table,
			Columns: bookmark.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bookmark.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.URL(); ok {
		_spec.SetField(bookmark.FieldURL, field.TypeString, value)
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.SetField(bookmark.FieldTitle, field.TypeString, value)
	}
	if value, ok := bu.mutation.Note(); ok {
		_spec.SetField(bookmark.FieldNote, field.TypeString, value)
	}
	if bu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.OwnerTable,
			Columns: []string{bookmark.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.OwnerTable,
			Columns: []string{bookmark.OwnerColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookmark.TagsTable,
			Columns: bookmark.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !bu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookmark.TagsTable,
			Columns: bookmark.TagsPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookmark.TagsTable,
			Columns: bookmark.TagsPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookmark.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// BookmarkUpdateOne is the builder for updating a single Bookmark entity.
type BookmarkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookmarkMutation
}

// SetUserID sets the "user_id" field.
func (buo *BookmarkUpdateOne) SetUserID(i int) *BookmarkUpdateOne {
	buo.mutation.SetUserID(i)
	return buo
}

// SetURL sets the "url" field.
func (buo *BookmarkUpdateOne) SetURL(s string) *BookmarkUpdateOne {
	buo.mutation.SetURL(s)
	return buo
}

// SetTitle sets the "title" field.
func (buo *BookmarkUpdateOne) SetTitle(s string) *BookmarkUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetNote sets the "note" field.
func (buo *BookmarkUpdateOne) SetNote(s string) *BookmarkUpdateOne {
	buo.mutation.SetNote(s)
	return buo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (buo *BookmarkUpdateOne) SetOwnerID(id int) *BookmarkUpdateOne {
	buo.mutation.SetOwnerID(id)
	return buo
}

// SetOwner sets the "owner" edge to the User entity.
func (buo *BookmarkUpdateOne) SetOwner(u *User) *BookmarkUpdateOne {
	return buo.SetOwnerID(u.ID)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (buo *BookmarkUpdateOne) AddTagIDs(ids ...int) *BookmarkUpdateOne {
	buo.mutation.AddTagIDs(ids...)
	return buo
}

// AddTags adds the "tags" edges to the Tag entity.
func (buo *BookmarkUpdateOne) AddTags(t ...*Tag) *BookmarkUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.AddTagIDs(ids...)
}

// Mutation returns the BookmarkMutation object of the builder.
func (buo *BookmarkUpdateOne) Mutation() *BookmarkMutation {
	return buo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (buo *BookmarkUpdateOne) ClearOwner() *BookmarkUpdateOne {
	buo.mutation.ClearOwner()
	return buo
}

// ClearTags clears all "tags" edges to the Tag entity.
func (buo *BookmarkUpdateOne) ClearTags() *BookmarkUpdateOne {
	buo.mutation.ClearTags()
	return buo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (buo *BookmarkUpdateOne) RemoveTagIDs(ids ...int) *BookmarkUpdateOne {
	buo.mutation.RemoveTagIDs(ids...)
	return buo
}

// RemoveTags removes "tags" edges to Tag entities.
func (buo *BookmarkUpdateOne) RemoveTags(t ...*Tag) *BookmarkUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.RemoveTagIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookmarkUpdateOne) Select(field string, fields ...string) *BookmarkUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Bookmark entity.
func (buo *BookmarkUpdateOne) Save(ctx context.Context) (*Bookmark, error) {
	var (
		err  error
		node *Bookmark
	)
	if len(buo.hooks) == 0 {
		if err = buo.check(); err != nil {
			return nil, err
		}
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookmarkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buo.check(); err != nil {
				return nil, err
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, buo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (buo *BookmarkUpdateOne) SaveX(ctx context.Context) *Bookmark {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookmarkUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookmarkUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BookmarkUpdateOne) check() error {
	if v, ok := buo.mutation.URL(); ok {
		if err := bookmark.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Bookmark.url": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Title(); ok {
		if err := bookmark.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Bookmark.title": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Note(); ok {
		if err := bookmark.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf(`ent: validator failed for field "Bookmark.note": %w`, err)}
		}
	}
	if _, ok := buo.mutation.OwnerID(); buo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Bookmark.owner"`)
	}
	return nil
}

func (buo *BookmarkUpdateOne) sqlSave(ctx context.Context) (_node *Bookmark, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bookmark.Table,
			Columns: bookmark.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bookmark.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Bookmark.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bookmark.FieldID)
		for _, f := range fields {
			if !bookmark.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bookmark.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.URL(); ok {
		_spec.SetField(bookmark.FieldURL, field.TypeString, value)
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.SetField(bookmark.FieldTitle, field.TypeString, value)
	}
	if value, ok := buo.mutation.Note(); ok {
		_spec.SetField(bookmark.FieldNote, field.TypeString, value)
	}
	if buo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.OwnerTable,
			Columns: []string{bookmark.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.OwnerTable,
			Columns: []string{bookmark.OwnerColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookmark.TagsTable,
			Columns: bookmark.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !buo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookmark.TagsTable,
			Columns: bookmark.TagsPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookmark.TagsTable,
			Columns: bookmark.TagsPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Bookmark{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookmark.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
