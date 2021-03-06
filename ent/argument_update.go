// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/knabben/ggql/ent/argument"
	"github.com/knabben/ggql/ent/predicate"
)

// ArgumentUpdate is the builder for updating Argument entities.
type ArgumentUpdate struct {
	config
	name        *string
	description *string
	type_kind   *string
	type_name   *string
	predicates  []predicate.Argument
}

// Where adds a new predicate for the builder.
func (au *ArgumentUpdate) Where(ps ...predicate.Argument) *ArgumentUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetName sets the name field.
func (au *ArgumentUpdate) SetName(s string) *ArgumentUpdate {
	au.name = &s
	return au
}

// SetDescription sets the description field.
func (au *ArgumentUpdate) SetDescription(s string) *ArgumentUpdate {
	au.description = &s
	return au
}

// SetTypeKind sets the type_kind field.
func (au *ArgumentUpdate) SetTypeKind(s string) *ArgumentUpdate {
	au.type_kind = &s
	return au
}

// SetTypeName sets the type_name field.
func (au *ArgumentUpdate) SetTypeName(s string) *ArgumentUpdate {
	au.type_name = &s
	return au
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *ArgumentUpdate) Save(ctx context.Context) (int, error) {
	return au.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArgumentUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArgumentUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArgumentUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ArgumentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   argument.Table,
			Columns: argument.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: argument.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := au.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: argument.FieldName,
		})
	}
	if value := au.description; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: argument.FieldDescription,
		})
	}
	if value := au.type_kind; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: argument.FieldTypeKind,
		})
	}
	if value := au.type_name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: argument.FieldTypeName,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ArgumentUpdateOne is the builder for updating a single Argument entity.
type ArgumentUpdateOne struct {
	config
	id          int
	name        *string
	description *string
	type_kind   *string
	type_name   *string
}

// SetName sets the name field.
func (auo *ArgumentUpdateOne) SetName(s string) *ArgumentUpdateOne {
	auo.name = &s
	return auo
}

// SetDescription sets the description field.
func (auo *ArgumentUpdateOne) SetDescription(s string) *ArgumentUpdateOne {
	auo.description = &s
	return auo
}

// SetTypeKind sets the type_kind field.
func (auo *ArgumentUpdateOne) SetTypeKind(s string) *ArgumentUpdateOne {
	auo.type_kind = &s
	return auo
}

// SetTypeName sets the type_name field.
func (auo *ArgumentUpdateOne) SetTypeName(s string) *ArgumentUpdateOne {
	auo.type_name = &s
	return auo
}

// Save executes the query and returns the updated entity.
func (auo *ArgumentUpdateOne) Save(ctx context.Context) (*Argument, error) {
	return auo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArgumentUpdateOne) SaveX(ctx context.Context) *Argument {
	a, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// Exec executes the query on the entity.
func (auo *ArgumentUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArgumentUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ArgumentUpdateOne) sqlSave(ctx context.Context) (a *Argument, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   argument.Table,
			Columns: argument.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  auo.id,
				Type:   field.TypeInt,
				Column: argument.FieldID,
			},
		},
	}
	if value := auo.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: argument.FieldName,
		})
	}
	if value := auo.description; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: argument.FieldDescription,
		})
	}
	if value := auo.type_kind; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: argument.FieldTypeKind,
		})
	}
	if value := auo.type_name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: argument.FieldTypeName,
		})
	}
	a = &Argument{config: auo.config}
	_spec.Assign = a.assignValues
	_spec.ScanValues = a.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}
