// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/knabben/ggql/ent/objecttype"
	"github.com/knabben/ggql/ent/predicate"
)

// ObjectTypeQuery is the builder for querying ObjectType entities.
type ObjectTypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.ObjectType
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (otq *ObjectTypeQuery) Where(ps ...predicate.ObjectType) *ObjectTypeQuery {
	otq.predicates = append(otq.predicates, ps...)
	return otq
}

// Limit adds a limit step to the query.
func (otq *ObjectTypeQuery) Limit(limit int) *ObjectTypeQuery {
	otq.limit = &limit
	return otq
}

// Offset adds an offset step to the query.
func (otq *ObjectTypeQuery) Offset(offset int) *ObjectTypeQuery {
	otq.offset = &offset
	return otq
}

// Order adds an order step to the query.
func (otq *ObjectTypeQuery) Order(o ...Order) *ObjectTypeQuery {
	otq.order = append(otq.order, o...)
	return otq
}

// First returns the first ObjectType entity in the query. Returns *NotFoundError when no objecttype was found.
func (otq *ObjectTypeQuery) First(ctx context.Context) (*ObjectType, error) {
	ots, err := otq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ots) == 0 {
		return nil, &NotFoundError{objecttype.Label}
	}
	return ots[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (otq *ObjectTypeQuery) FirstX(ctx context.Context) *ObjectType {
	ot, err := otq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ot
}

// FirstID returns the first ObjectType id in the query. Returns *NotFoundError when no id was found.
func (otq *ObjectTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = otq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{objecttype.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (otq *ObjectTypeQuery) FirstXID(ctx context.Context) int {
	id, err := otq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only ObjectType entity in the query, returns an error if not exactly one entity was returned.
func (otq *ObjectTypeQuery) Only(ctx context.Context) (*ObjectType, error) {
	ots, err := otq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ots) {
	case 1:
		return ots[0], nil
	case 0:
		return nil, &NotFoundError{objecttype.Label}
	default:
		return nil, &NotSingularError{objecttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (otq *ObjectTypeQuery) OnlyX(ctx context.Context) *ObjectType {
	ot, err := otq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ot
}

// OnlyID returns the only ObjectType id in the query, returns an error if not exactly one id was returned.
func (otq *ObjectTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = otq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{objecttype.Label}
	default:
		err = &NotSingularError{objecttype.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (otq *ObjectTypeQuery) OnlyXID(ctx context.Context) int {
	id, err := otq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ObjectTypes.
func (otq *ObjectTypeQuery) All(ctx context.Context) ([]*ObjectType, error) {
	return otq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (otq *ObjectTypeQuery) AllX(ctx context.Context) []*ObjectType {
	ots, err := otq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ots
}

// IDs executes the query and returns a list of ObjectType ids.
func (otq *ObjectTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := otq.Select(objecttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (otq *ObjectTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := otq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (otq *ObjectTypeQuery) Count(ctx context.Context) (int, error) {
	return otq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (otq *ObjectTypeQuery) CountX(ctx context.Context) int {
	count, err := otq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (otq *ObjectTypeQuery) Exist(ctx context.Context) (bool, error) {
	return otq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (otq *ObjectTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := otq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (otq *ObjectTypeQuery) Clone() *ObjectTypeQuery {
	return &ObjectTypeQuery{
		config:     otq.config,
		limit:      otq.limit,
		offset:     otq.offset,
		order:      append([]Order{}, otq.order...),
		unique:     append([]string{}, otq.unique...),
		predicates: append([]predicate.ObjectType{}, otq.predicates...),
		// clone intermediate query.
		sql: otq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ObjectType.Query().
//		GroupBy(objecttype.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (otq *ObjectTypeQuery) GroupBy(field string, fields ...string) *ObjectTypeGroupBy {
	group := &ObjectTypeGroupBy{config: otq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = otq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.ObjectType.Query().
//		Select(objecttype.FieldName).
//		Scan(ctx, &v)
//
func (otq *ObjectTypeQuery) Select(field string, fields ...string) *ObjectTypeSelect {
	selector := &ObjectTypeSelect{config: otq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = otq.sqlQuery()
	return selector
}

func (otq *ObjectTypeQuery) sqlAll(ctx context.Context) ([]*ObjectType, error) {
	var (
		nodes = []*ObjectType{}
		_spec = otq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &ObjectType{config: otq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, otq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (otq *ObjectTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := otq.querySpec()
	return sqlgraph.CountNodes(ctx, otq.driver, _spec)
}

func (otq *ObjectTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := otq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (otq *ObjectTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   objecttype.Table,
			Columns: objecttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: objecttype.FieldID,
			},
		},
		From:   otq.sql,
		Unique: true,
	}
	if ps := otq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := otq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := otq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := otq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (otq *ObjectTypeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(otq.driver.Dialect())
	t1 := builder.Table(objecttype.Table)
	selector := builder.Select(t1.Columns(objecttype.Columns...)...).From(t1)
	if otq.sql != nil {
		selector = otq.sql
		selector.Select(selector.Columns(objecttype.Columns...)...)
	}
	for _, p := range otq.predicates {
		p(selector)
	}
	for _, p := range otq.order {
		p(selector)
	}
	if offset := otq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := otq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ObjectTypeGroupBy is the builder for group-by ObjectType entities.
type ObjectTypeGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (otgb *ObjectTypeGroupBy) Aggregate(fns ...Aggregate) *ObjectTypeGroupBy {
	otgb.fns = append(otgb.fns, fns...)
	return otgb
}

// Scan applies the group-by query and scan the result into the given value.
func (otgb *ObjectTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	return otgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (otgb *ObjectTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := otgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (otgb *ObjectTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(otgb.fields) > 1 {
		return nil, errors.New("ent: ObjectTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := otgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (otgb *ObjectTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := otgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (otgb *ObjectTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(otgb.fields) > 1 {
		return nil, errors.New("ent: ObjectTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := otgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (otgb *ObjectTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := otgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (otgb *ObjectTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(otgb.fields) > 1 {
		return nil, errors.New("ent: ObjectTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := otgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (otgb *ObjectTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := otgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (otgb *ObjectTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(otgb.fields) > 1 {
		return nil, errors.New("ent: ObjectTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := otgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (otgb *ObjectTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := otgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (otgb *ObjectTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := otgb.sqlQuery().Query()
	if err := otgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (otgb *ObjectTypeGroupBy) sqlQuery() *sql.Selector {
	selector := otgb.sql
	columns := make([]string, 0, len(otgb.fields)+len(otgb.fns))
	columns = append(columns, otgb.fields...)
	for _, fn := range otgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(otgb.fields...)
}

// ObjectTypeSelect is the builder for select fields of ObjectType entities.
type ObjectTypeSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ots *ObjectTypeSelect) Scan(ctx context.Context, v interface{}) error {
	return ots.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ots *ObjectTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ots.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ots *ObjectTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ots.fields) > 1 {
		return nil, errors.New("ent: ObjectTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ots.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ots *ObjectTypeSelect) StringsX(ctx context.Context) []string {
	v, err := ots.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ots *ObjectTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ots.fields) > 1 {
		return nil, errors.New("ent: ObjectTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ots.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ots *ObjectTypeSelect) IntsX(ctx context.Context) []int {
	v, err := ots.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ots *ObjectTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ots.fields) > 1 {
		return nil, errors.New("ent: ObjectTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ots.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ots *ObjectTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ots.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ots *ObjectTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ots.fields) > 1 {
		return nil, errors.New("ent: ObjectTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ots.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ots *ObjectTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := ots.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ots *ObjectTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ots.sqlQuery().Query()
	if err := ots.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ots *ObjectTypeSelect) sqlQuery() sql.Querier {
	selector := ots.sql
	selector.Select(selector.Columns(ots.fields...)...)
	return selector
}