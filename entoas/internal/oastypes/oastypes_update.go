// Code generated by entc, DO NOT EDIT.

package oastypes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"entgo.io/contrib/entoas/internal/oastypes/oastypes"
	"entgo.io/contrib/entoas/internal/oastypes/predicate"
	"entgo.io/contrib/entoas/internal/oastypes/schema"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OASTypesUpdate is the builder for updating OASTypes entities.
type OASTypesUpdate struct {
	config
	hooks    []Hook
	mutation *OASTypesMutation
}

// Where appends a list predicates to the OASTypesUpdate builder.
func (otu *OASTypesUpdate) Where(ps ...predicate.OASTypes) *OASTypesUpdate {
	otu.mutation.Where(ps...)
	return otu
}

// SetInt sets the "int" field.
func (otu *OASTypesUpdate) SetInt(i int) *OASTypesUpdate {
	otu.mutation.ResetInt()
	otu.mutation.SetInt(i)
	return otu
}

// AddInt adds i to the "int" field.
func (otu *OASTypesUpdate) AddInt(i int) *OASTypesUpdate {
	otu.mutation.AddInt(i)
	return otu
}

// SetInt8 sets the "int8" field.
func (otu *OASTypesUpdate) SetInt8(i int8) *OASTypesUpdate {
	otu.mutation.ResetInt8()
	otu.mutation.SetInt8(i)
	return otu
}

// AddInt8 adds i to the "int8" field.
func (otu *OASTypesUpdate) AddInt8(i int8) *OASTypesUpdate {
	otu.mutation.AddInt8(i)
	return otu
}

// SetInt16 sets the "int16" field.
func (otu *OASTypesUpdate) SetInt16(i int16) *OASTypesUpdate {
	otu.mutation.ResetInt16()
	otu.mutation.SetInt16(i)
	return otu
}

// AddInt16 adds i to the "int16" field.
func (otu *OASTypesUpdate) AddInt16(i int16) *OASTypesUpdate {
	otu.mutation.AddInt16(i)
	return otu
}

// SetInt32 sets the "int32" field.
func (otu *OASTypesUpdate) SetInt32(i int32) *OASTypesUpdate {
	otu.mutation.ResetInt32()
	otu.mutation.SetInt32(i)
	return otu
}

// AddInt32 adds i to the "int32" field.
func (otu *OASTypesUpdate) AddInt32(i int32) *OASTypesUpdate {
	otu.mutation.AddInt32(i)
	return otu
}

// SetInt64 sets the "int64" field.
func (otu *OASTypesUpdate) SetInt64(i int64) *OASTypesUpdate {
	otu.mutation.ResetInt64()
	otu.mutation.SetInt64(i)
	return otu
}

// AddInt64 adds i to the "int64" field.
func (otu *OASTypesUpdate) AddInt64(i int64) *OASTypesUpdate {
	otu.mutation.AddInt64(i)
	return otu
}

// SetUint sets the "uint" field.
func (otu *OASTypesUpdate) SetUint(u uint) *OASTypesUpdate {
	otu.mutation.ResetUint()
	otu.mutation.SetUint(u)
	return otu
}

// AddUint adds u to the "uint" field.
func (otu *OASTypesUpdate) AddUint(u int) *OASTypesUpdate {
	otu.mutation.AddUint(u)
	return otu
}

// SetUint8 sets the "uint8" field.
func (otu *OASTypesUpdate) SetUint8(u uint8) *OASTypesUpdate {
	otu.mutation.ResetUint8()
	otu.mutation.SetUint8(u)
	return otu
}

// AddUint8 adds u to the "uint8" field.
func (otu *OASTypesUpdate) AddUint8(u int8) *OASTypesUpdate {
	otu.mutation.AddUint8(u)
	return otu
}

// SetUint16 sets the "uint16" field.
func (otu *OASTypesUpdate) SetUint16(u uint16) *OASTypesUpdate {
	otu.mutation.ResetUint16()
	otu.mutation.SetUint16(u)
	return otu
}

// AddUint16 adds u to the "uint16" field.
func (otu *OASTypesUpdate) AddUint16(u int16) *OASTypesUpdate {
	otu.mutation.AddUint16(u)
	return otu
}

// SetUint32 sets the "uint32" field.
func (otu *OASTypesUpdate) SetUint32(u uint32) *OASTypesUpdate {
	otu.mutation.ResetUint32()
	otu.mutation.SetUint32(u)
	return otu
}

// AddUint32 adds u to the "uint32" field.
func (otu *OASTypesUpdate) AddUint32(u int32) *OASTypesUpdate {
	otu.mutation.AddUint32(u)
	return otu
}

// SetUint64 sets the "uint64" field.
func (otu *OASTypesUpdate) SetUint64(u uint64) *OASTypesUpdate {
	otu.mutation.ResetUint64()
	otu.mutation.SetUint64(u)
	return otu
}

// AddUint64 adds u to the "uint64" field.
func (otu *OASTypesUpdate) AddUint64(u int64) *OASTypesUpdate {
	otu.mutation.AddUint64(u)
	return otu
}

// SetFloat32 sets the "float32" field.
func (otu *OASTypesUpdate) SetFloat32(f float32) *OASTypesUpdate {
	otu.mutation.ResetFloat32()
	otu.mutation.SetFloat32(f)
	return otu
}

// AddFloat32 adds f to the "float32" field.
func (otu *OASTypesUpdate) AddFloat32(f float32) *OASTypesUpdate {
	otu.mutation.AddFloat32(f)
	return otu
}

// SetFloat64 sets the "float64" field.
func (otu *OASTypesUpdate) SetFloat64(f float64) *OASTypesUpdate {
	otu.mutation.ResetFloat64()
	otu.mutation.SetFloat64(f)
	return otu
}

// AddFloat64 adds f to the "float64" field.
func (otu *OASTypesUpdate) AddFloat64(f float64) *OASTypesUpdate {
	otu.mutation.AddFloat64(f)
	return otu
}

// SetStringField sets the "string_field" field.
func (otu *OASTypesUpdate) SetStringField(s string) *OASTypesUpdate {
	otu.mutation.SetStringField(s)
	return otu
}

// SetBool sets the "bool" field.
func (otu *OASTypesUpdate) SetBool(b bool) *OASTypesUpdate {
	otu.mutation.SetBool(b)
	return otu
}

// SetUUID sets the "uuid" field.
func (otu *OASTypesUpdate) SetUUID(u uuid.UUID) *OASTypesUpdate {
	otu.mutation.SetUUID(u)
	return otu
}

// SetTime sets the "time" field.
func (otu *OASTypesUpdate) SetTime(t time.Time) *OASTypesUpdate {
	otu.mutation.SetTime(t)
	return otu
}

// SetText sets the "text" field.
func (otu *OASTypesUpdate) SetText(s string) *OASTypesUpdate {
	otu.mutation.SetText(s)
	return otu
}

// SetState sets the "state" field.
func (otu *OASTypesUpdate) SetState(o oastypes.State) *OASTypesUpdate {
	otu.mutation.SetState(o)
	return otu
}

// SetStrings sets the "strings" field.
func (otu *OASTypesUpdate) SetStrings(s []string) *OASTypesUpdate {
	otu.mutation.SetStrings(s)
	return otu
}

// SetInts sets the "ints" field.
func (otu *OASTypesUpdate) SetInts(i []int) *OASTypesUpdate {
	otu.mutation.SetInts(i)
	return otu
}

// SetFloats sets the "floats" field.
func (otu *OASTypesUpdate) SetFloats(f []float64) *OASTypesUpdate {
	otu.mutation.SetFloats(f)
	return otu
}

// SetBytes sets the "bytes" field.
func (otu *OASTypesUpdate) SetBytes(b []byte) *OASTypesUpdate {
	otu.mutation.SetBytes(b)
	return otu
}

// SetNicknames sets the "nicknames" field.
func (otu *OASTypesUpdate) SetNicknames(s []string) *OASTypesUpdate {
	otu.mutation.SetNicknames(s)
	return otu
}

// SetJSONSlice sets the "json_slice" field.
func (otu *OASTypesUpdate) SetJSONSlice(h []http.Dir) *OASTypesUpdate {
	otu.mutation.SetJSONSlice(h)
	return otu
}

// SetJSONObj sets the "json_obj" field.
func (otu *OASTypesUpdate) SetJSONObj(u url.URL) *OASTypesUpdate {
	otu.mutation.SetJSONObj(u)
	return otu
}

// SetOther sets the "other" field.
func (otu *OASTypesUpdate) SetOther(s *schema.Link) *OASTypesUpdate {
	otu.mutation.SetOther(s)
	return otu
}

// Mutation returns the OASTypesMutation object of the builder.
func (otu *OASTypesUpdate) Mutation() *OASTypesMutation {
	return otu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (otu *OASTypesUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(otu.hooks) == 0 {
		if err = otu.check(); err != nil {
			return 0, err
		}
		affected, err = otu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OASTypesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = otu.check(); err != nil {
				return 0, err
			}
			otu.mutation = mutation
			affected, err = otu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(otu.hooks) - 1; i >= 0; i-- {
			if otu.hooks[i] == nil {
				return 0, fmt.Errorf("oastypes: uninitialized hook (forgotten import oastypes/runtime?)")
			}
			mut = otu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, otu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (otu *OASTypesUpdate) SaveX(ctx context.Context) int {
	affected, err := otu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (otu *OASTypesUpdate) Exec(ctx context.Context) error {
	_, err := otu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otu *OASTypesUpdate) ExecX(ctx context.Context) {
	if err := otu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (otu *OASTypesUpdate) check() error {
	if v, ok := otu.mutation.State(); ok {
		if err := oastypes.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`oastypes: validator failed for field "OASTypes.state": %w`, err)}
		}
	}
	return nil
}

func (otu *OASTypesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oastypes.Table,
			Columns: oastypes.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oastypes.FieldID,
			},
		},
	}
	if ps := otu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := otu.mutation.Int(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oastypes.FieldInt,
		})
	}
	if value, ok := otu.mutation.AddedInt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oastypes.FieldInt,
		})
	}
	if value, ok := otu.mutation.Int8(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: oastypes.FieldInt8,
		})
	}
	if value, ok := otu.mutation.AddedInt8(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: oastypes.FieldInt8,
		})
	}
	if value, ok := otu.mutation.Int16(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: oastypes.FieldInt16,
		})
	}
	if value, ok := otu.mutation.AddedInt16(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: oastypes.FieldInt16,
		})
	}
	if value, ok := otu.mutation.Int32(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: oastypes.FieldInt32,
		})
	}
	if value, ok := otu.mutation.AddedInt32(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: oastypes.FieldInt32,
		})
	}
	if value, ok := otu.mutation.Int64(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: oastypes.FieldInt64,
		})
	}
	if value, ok := otu.mutation.AddedInt64(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: oastypes.FieldInt64,
		})
	}
	if value, ok := otu.mutation.Uint(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: oastypes.FieldUint,
		})
	}
	if value, ok := otu.mutation.AddedUint(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: oastypes.FieldUint,
		})
	}
	if value, ok := otu.mutation.Uint8(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: oastypes.FieldUint8,
		})
	}
	if value, ok := otu.mutation.AddedUint8(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: oastypes.FieldUint8,
		})
	}
	if value, ok := otu.mutation.Uint16(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint16,
			Value:  value,
			Column: oastypes.FieldUint16,
		})
	}
	if value, ok := otu.mutation.AddedUint16(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint16,
			Value:  value,
			Column: oastypes.FieldUint16,
		})
	}
	if value, ok := otu.mutation.Uint32(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oastypes.FieldUint32,
		})
	}
	if value, ok := otu.mutation.AddedUint32(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oastypes.FieldUint32,
		})
	}
	if value, ok := otu.mutation.Uint64(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: oastypes.FieldUint64,
		})
	}
	if value, ok := otu.mutation.AddedUint64(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: oastypes.FieldUint64,
		})
	}
	if value, ok := otu.mutation.Float32(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: oastypes.FieldFloat32,
		})
	}
	if value, ok := otu.mutation.AddedFloat32(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: oastypes.FieldFloat32,
		})
	}
	if value, ok := otu.mutation.Float64(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: oastypes.FieldFloat64,
		})
	}
	if value, ok := otu.mutation.AddedFloat64(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: oastypes.FieldFloat64,
		})
	}
	if value, ok := otu.mutation.StringField(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oastypes.FieldStringField,
		})
	}
	if value, ok := otu.mutation.Bool(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: oastypes.FieldBool,
		})
	}
	if value, ok := otu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: oastypes.FieldUUID,
		})
	}
	if value, ok := otu.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oastypes.FieldTime,
		})
	}
	if value, ok := otu.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oastypes.FieldText,
		})
	}
	if value, ok := otu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: oastypes.FieldState,
		})
	}
	if value, ok := otu.mutation.Strings(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oastypes.FieldStrings,
		})
	}
	if value, ok := otu.mutation.Ints(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oastypes.FieldInts,
		})
	}
	if value, ok := otu.mutation.Floats(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oastypes.FieldFloats,
		})
	}
	if value, ok := otu.mutation.Bytes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: oastypes.FieldBytes,
		})
	}
	if value, ok := otu.mutation.Nicknames(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oastypes.FieldNicknames,
		})
	}
	if value, ok := otu.mutation.JSONSlice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oastypes.FieldJSONSlice,
		})
	}
	if value, ok := otu.mutation.JSONObj(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oastypes.FieldJSONObj,
		})
	}
	if value, ok := otu.mutation.Other(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: oastypes.FieldOther,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, otu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oastypes.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OASTypesUpdateOne is the builder for updating a single OASTypes entity.
type OASTypesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OASTypesMutation
}

// SetInt sets the "int" field.
func (otuo *OASTypesUpdateOne) SetInt(i int) *OASTypesUpdateOne {
	otuo.mutation.ResetInt()
	otuo.mutation.SetInt(i)
	return otuo
}

// AddInt adds i to the "int" field.
func (otuo *OASTypesUpdateOne) AddInt(i int) *OASTypesUpdateOne {
	otuo.mutation.AddInt(i)
	return otuo
}

// SetInt8 sets the "int8" field.
func (otuo *OASTypesUpdateOne) SetInt8(i int8) *OASTypesUpdateOne {
	otuo.mutation.ResetInt8()
	otuo.mutation.SetInt8(i)
	return otuo
}

// AddInt8 adds i to the "int8" field.
func (otuo *OASTypesUpdateOne) AddInt8(i int8) *OASTypesUpdateOne {
	otuo.mutation.AddInt8(i)
	return otuo
}

// SetInt16 sets the "int16" field.
func (otuo *OASTypesUpdateOne) SetInt16(i int16) *OASTypesUpdateOne {
	otuo.mutation.ResetInt16()
	otuo.mutation.SetInt16(i)
	return otuo
}

// AddInt16 adds i to the "int16" field.
func (otuo *OASTypesUpdateOne) AddInt16(i int16) *OASTypesUpdateOne {
	otuo.mutation.AddInt16(i)
	return otuo
}

// SetInt32 sets the "int32" field.
func (otuo *OASTypesUpdateOne) SetInt32(i int32) *OASTypesUpdateOne {
	otuo.mutation.ResetInt32()
	otuo.mutation.SetInt32(i)
	return otuo
}

// AddInt32 adds i to the "int32" field.
func (otuo *OASTypesUpdateOne) AddInt32(i int32) *OASTypesUpdateOne {
	otuo.mutation.AddInt32(i)
	return otuo
}

// SetInt64 sets the "int64" field.
func (otuo *OASTypesUpdateOne) SetInt64(i int64) *OASTypesUpdateOne {
	otuo.mutation.ResetInt64()
	otuo.mutation.SetInt64(i)
	return otuo
}

// AddInt64 adds i to the "int64" field.
func (otuo *OASTypesUpdateOne) AddInt64(i int64) *OASTypesUpdateOne {
	otuo.mutation.AddInt64(i)
	return otuo
}

// SetUint sets the "uint" field.
func (otuo *OASTypesUpdateOne) SetUint(u uint) *OASTypesUpdateOne {
	otuo.mutation.ResetUint()
	otuo.mutation.SetUint(u)
	return otuo
}

// AddUint adds u to the "uint" field.
func (otuo *OASTypesUpdateOne) AddUint(u int) *OASTypesUpdateOne {
	otuo.mutation.AddUint(u)
	return otuo
}

// SetUint8 sets the "uint8" field.
func (otuo *OASTypesUpdateOne) SetUint8(u uint8) *OASTypesUpdateOne {
	otuo.mutation.ResetUint8()
	otuo.mutation.SetUint8(u)
	return otuo
}

// AddUint8 adds u to the "uint8" field.
func (otuo *OASTypesUpdateOne) AddUint8(u int8) *OASTypesUpdateOne {
	otuo.mutation.AddUint8(u)
	return otuo
}

// SetUint16 sets the "uint16" field.
func (otuo *OASTypesUpdateOne) SetUint16(u uint16) *OASTypesUpdateOne {
	otuo.mutation.ResetUint16()
	otuo.mutation.SetUint16(u)
	return otuo
}

// AddUint16 adds u to the "uint16" field.
func (otuo *OASTypesUpdateOne) AddUint16(u int16) *OASTypesUpdateOne {
	otuo.mutation.AddUint16(u)
	return otuo
}

// SetUint32 sets the "uint32" field.
func (otuo *OASTypesUpdateOne) SetUint32(u uint32) *OASTypesUpdateOne {
	otuo.mutation.ResetUint32()
	otuo.mutation.SetUint32(u)
	return otuo
}

// AddUint32 adds u to the "uint32" field.
func (otuo *OASTypesUpdateOne) AddUint32(u int32) *OASTypesUpdateOne {
	otuo.mutation.AddUint32(u)
	return otuo
}

// SetUint64 sets the "uint64" field.
func (otuo *OASTypesUpdateOne) SetUint64(u uint64) *OASTypesUpdateOne {
	otuo.mutation.ResetUint64()
	otuo.mutation.SetUint64(u)
	return otuo
}

// AddUint64 adds u to the "uint64" field.
func (otuo *OASTypesUpdateOne) AddUint64(u int64) *OASTypesUpdateOne {
	otuo.mutation.AddUint64(u)
	return otuo
}

// SetFloat32 sets the "float32" field.
func (otuo *OASTypesUpdateOne) SetFloat32(f float32) *OASTypesUpdateOne {
	otuo.mutation.ResetFloat32()
	otuo.mutation.SetFloat32(f)
	return otuo
}

// AddFloat32 adds f to the "float32" field.
func (otuo *OASTypesUpdateOne) AddFloat32(f float32) *OASTypesUpdateOne {
	otuo.mutation.AddFloat32(f)
	return otuo
}

// SetFloat64 sets the "float64" field.
func (otuo *OASTypesUpdateOne) SetFloat64(f float64) *OASTypesUpdateOne {
	otuo.mutation.ResetFloat64()
	otuo.mutation.SetFloat64(f)
	return otuo
}

// AddFloat64 adds f to the "float64" field.
func (otuo *OASTypesUpdateOne) AddFloat64(f float64) *OASTypesUpdateOne {
	otuo.mutation.AddFloat64(f)
	return otuo
}

// SetStringField sets the "string_field" field.
func (otuo *OASTypesUpdateOne) SetStringField(s string) *OASTypesUpdateOne {
	otuo.mutation.SetStringField(s)
	return otuo
}

// SetBool sets the "bool" field.
func (otuo *OASTypesUpdateOne) SetBool(b bool) *OASTypesUpdateOne {
	otuo.mutation.SetBool(b)
	return otuo
}

// SetUUID sets the "uuid" field.
func (otuo *OASTypesUpdateOne) SetUUID(u uuid.UUID) *OASTypesUpdateOne {
	otuo.mutation.SetUUID(u)
	return otuo
}

// SetTime sets the "time" field.
func (otuo *OASTypesUpdateOne) SetTime(t time.Time) *OASTypesUpdateOne {
	otuo.mutation.SetTime(t)
	return otuo
}

// SetText sets the "text" field.
func (otuo *OASTypesUpdateOne) SetText(s string) *OASTypesUpdateOne {
	otuo.mutation.SetText(s)
	return otuo
}

// SetState sets the "state" field.
func (otuo *OASTypesUpdateOne) SetState(o oastypes.State) *OASTypesUpdateOne {
	otuo.mutation.SetState(o)
	return otuo
}

// SetStrings sets the "strings" field.
func (otuo *OASTypesUpdateOne) SetStrings(s []string) *OASTypesUpdateOne {
	otuo.mutation.SetStrings(s)
	return otuo
}

// SetInts sets the "ints" field.
func (otuo *OASTypesUpdateOne) SetInts(i []int) *OASTypesUpdateOne {
	otuo.mutation.SetInts(i)
	return otuo
}

// SetFloats sets the "floats" field.
func (otuo *OASTypesUpdateOne) SetFloats(f []float64) *OASTypesUpdateOne {
	otuo.mutation.SetFloats(f)
	return otuo
}

// SetBytes sets the "bytes" field.
func (otuo *OASTypesUpdateOne) SetBytes(b []byte) *OASTypesUpdateOne {
	otuo.mutation.SetBytes(b)
	return otuo
}

// SetNicknames sets the "nicknames" field.
func (otuo *OASTypesUpdateOne) SetNicknames(s []string) *OASTypesUpdateOne {
	otuo.mutation.SetNicknames(s)
	return otuo
}

// SetJSONSlice sets the "json_slice" field.
func (otuo *OASTypesUpdateOne) SetJSONSlice(h []http.Dir) *OASTypesUpdateOne {
	otuo.mutation.SetJSONSlice(h)
	return otuo
}

// SetJSONObj sets the "json_obj" field.
func (otuo *OASTypesUpdateOne) SetJSONObj(u url.URL) *OASTypesUpdateOne {
	otuo.mutation.SetJSONObj(u)
	return otuo
}

// SetOther sets the "other" field.
func (otuo *OASTypesUpdateOne) SetOther(s *schema.Link) *OASTypesUpdateOne {
	otuo.mutation.SetOther(s)
	return otuo
}

// Mutation returns the OASTypesMutation object of the builder.
func (otuo *OASTypesUpdateOne) Mutation() *OASTypesMutation {
	return otuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (otuo *OASTypesUpdateOne) Select(field string, fields ...string) *OASTypesUpdateOne {
	otuo.fields = append([]string{field}, fields...)
	return otuo
}

// Save executes the query and returns the updated OASTypes entity.
func (otuo *OASTypesUpdateOne) Save(ctx context.Context) (*OASTypes, error) {
	var (
		err  error
		node *OASTypes
	)
	if len(otuo.hooks) == 0 {
		if err = otuo.check(); err != nil {
			return nil, err
		}
		node, err = otuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OASTypesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = otuo.check(); err != nil {
				return nil, err
			}
			otuo.mutation = mutation
			node, err = otuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(otuo.hooks) - 1; i >= 0; i-- {
			if otuo.hooks[i] == nil {
				return nil, fmt.Errorf("oastypes: uninitialized hook (forgotten import oastypes/runtime?)")
			}
			mut = otuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, otuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (otuo *OASTypesUpdateOne) SaveX(ctx context.Context) *OASTypes {
	node, err := otuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (otuo *OASTypesUpdateOne) Exec(ctx context.Context) error {
	_, err := otuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otuo *OASTypesUpdateOne) ExecX(ctx context.Context) {
	if err := otuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (otuo *OASTypesUpdateOne) check() error {
	if v, ok := otuo.mutation.State(); ok {
		if err := oastypes.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`oastypes: validator failed for field "OASTypes.state": %w`, err)}
		}
	}
	return nil
}

func (otuo *OASTypesUpdateOne) sqlSave(ctx context.Context) (_node *OASTypes, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oastypes.Table,
			Columns: oastypes.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oastypes.FieldID,
			},
		},
	}
	id, ok := otuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`oastypes: missing "OASTypes.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := otuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oastypes.FieldID)
		for _, f := range fields {
			if !oastypes.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("oastypes: invalid field %q for query", f)}
			}
			if f != oastypes.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := otuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := otuo.mutation.Int(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oastypes.FieldInt,
		})
	}
	if value, ok := otuo.mutation.AddedInt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oastypes.FieldInt,
		})
	}
	if value, ok := otuo.mutation.Int8(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: oastypes.FieldInt8,
		})
	}
	if value, ok := otuo.mutation.AddedInt8(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: oastypes.FieldInt8,
		})
	}
	if value, ok := otuo.mutation.Int16(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: oastypes.FieldInt16,
		})
	}
	if value, ok := otuo.mutation.AddedInt16(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: oastypes.FieldInt16,
		})
	}
	if value, ok := otuo.mutation.Int32(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: oastypes.FieldInt32,
		})
	}
	if value, ok := otuo.mutation.AddedInt32(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: oastypes.FieldInt32,
		})
	}
	if value, ok := otuo.mutation.Int64(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: oastypes.FieldInt64,
		})
	}
	if value, ok := otuo.mutation.AddedInt64(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: oastypes.FieldInt64,
		})
	}
	if value, ok := otuo.mutation.Uint(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: oastypes.FieldUint,
		})
	}
	if value, ok := otuo.mutation.AddedUint(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: oastypes.FieldUint,
		})
	}
	if value, ok := otuo.mutation.Uint8(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: oastypes.FieldUint8,
		})
	}
	if value, ok := otuo.mutation.AddedUint8(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: oastypes.FieldUint8,
		})
	}
	if value, ok := otuo.mutation.Uint16(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint16,
			Value:  value,
			Column: oastypes.FieldUint16,
		})
	}
	if value, ok := otuo.mutation.AddedUint16(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint16,
			Value:  value,
			Column: oastypes.FieldUint16,
		})
	}
	if value, ok := otuo.mutation.Uint32(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oastypes.FieldUint32,
		})
	}
	if value, ok := otuo.mutation.AddedUint32(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: oastypes.FieldUint32,
		})
	}
	if value, ok := otuo.mutation.Uint64(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: oastypes.FieldUint64,
		})
	}
	if value, ok := otuo.mutation.AddedUint64(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: oastypes.FieldUint64,
		})
	}
	if value, ok := otuo.mutation.Float32(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: oastypes.FieldFloat32,
		})
	}
	if value, ok := otuo.mutation.AddedFloat32(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: oastypes.FieldFloat32,
		})
	}
	if value, ok := otuo.mutation.Float64(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: oastypes.FieldFloat64,
		})
	}
	if value, ok := otuo.mutation.AddedFloat64(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: oastypes.FieldFloat64,
		})
	}
	if value, ok := otuo.mutation.StringField(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oastypes.FieldStringField,
		})
	}
	if value, ok := otuo.mutation.Bool(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: oastypes.FieldBool,
		})
	}
	if value, ok := otuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: oastypes.FieldUUID,
		})
	}
	if value, ok := otuo.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oastypes.FieldTime,
		})
	}
	if value, ok := otuo.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oastypes.FieldText,
		})
	}
	if value, ok := otuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: oastypes.FieldState,
		})
	}
	if value, ok := otuo.mutation.Strings(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oastypes.FieldStrings,
		})
	}
	if value, ok := otuo.mutation.Ints(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oastypes.FieldInts,
		})
	}
	if value, ok := otuo.mutation.Floats(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oastypes.FieldFloats,
		})
	}
	if value, ok := otuo.mutation.Bytes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: oastypes.FieldBytes,
		})
	}
	if value, ok := otuo.mutation.Nicknames(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oastypes.FieldNicknames,
		})
	}
	if value, ok := otuo.mutation.JSONSlice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oastypes.FieldJSONSlice,
		})
	}
	if value, ok := otuo.mutation.JSONObj(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oastypes.FieldJSONObj,
		})
	}
	if value, ok := otuo.mutation.Other(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: oastypes.FieldOther,
		})
	}
	_node = &OASTypes{config: otuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, otuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oastypes.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}