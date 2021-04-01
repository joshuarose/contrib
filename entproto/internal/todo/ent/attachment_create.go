// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/todo/ent/attachment"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AttachmentCreate is the builder for creating a Attachment entity.
type AttachmentCreate struct {
	config
	mutation *AttachmentMutation
	hooks    []Hook
}

// SetID sets the "id" field.
func (ac *AttachmentCreate) SetID(u uuid.UUID) *AttachmentCreate {
	ac.mutation.SetID(u)
	return ac
}

// Mutation returns the AttachmentMutation object of the builder.
func (ac *AttachmentCreate) Mutation() *AttachmentMutation {
	return ac.mutation
}

// Save creates the Attachment in the database.
func (ac *AttachmentCreate) Save(ctx context.Context) (*Attachment, error) {
	var (
		err  error
		node *Attachment
	)
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AttachmentCreate) SaveX(ctx context.Context) *Attachment {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (ac *AttachmentCreate) check() error {
	return nil
}

func (ac *AttachmentCreate) sqlSave(ctx context.Context) (*Attachment, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (ac *AttachmentCreate) createSpec() (*Attachment, *sqlgraph.CreateSpec) {
	var (
		_node = &Attachment{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: attachment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: attachment.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	return _node, _spec
}

// AttachmentCreateBulk is the builder for creating many Attachment entities in bulk.
type AttachmentCreateBulk struct {
	config
	builders []*AttachmentCreate
}

// Save creates the Attachment entities in the database.
func (acb *AttachmentCreateBulk) Save(ctx context.Context) ([]*Attachment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Attachment, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttachmentMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AttachmentCreateBulk) SaveX(ctx context.Context) []*Attachment {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
