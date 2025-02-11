// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/janus/ent/inboundwebhook"
	"github.com/lbrictson/janus/ent/predicate"
)

// InboundWebhookDelete is the builder for deleting a InboundWebhook entity.
type InboundWebhookDelete struct {
	config
	hooks    []Hook
	mutation *InboundWebhookMutation
}

// Where appends a list predicates to the InboundWebhookDelete builder.
func (iwd *InboundWebhookDelete) Where(ps ...predicate.InboundWebhook) *InboundWebhookDelete {
	iwd.mutation.Where(ps...)
	return iwd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (iwd *InboundWebhookDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, iwd.sqlExec, iwd.mutation, iwd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (iwd *InboundWebhookDelete) ExecX(ctx context.Context) int {
	n, err := iwd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (iwd *InboundWebhookDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(inboundwebhook.Table, sqlgraph.NewFieldSpec(inboundwebhook.FieldID, field.TypeInt))
	if ps := iwd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, iwd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	iwd.mutation.done = true
	return affected, err
}

// InboundWebhookDeleteOne is the builder for deleting a single InboundWebhook entity.
type InboundWebhookDeleteOne struct {
	iwd *InboundWebhookDelete
}

// Where appends a list predicates to the InboundWebhookDelete builder.
func (iwdo *InboundWebhookDeleteOne) Where(ps ...predicate.InboundWebhook) *InboundWebhookDeleteOne {
	iwdo.iwd.mutation.Where(ps...)
	return iwdo
}

// Exec executes the deletion query.
func (iwdo *InboundWebhookDeleteOne) Exec(ctx context.Context) error {
	n, err := iwdo.iwd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{inboundwebhook.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (iwdo *InboundWebhookDeleteOne) ExecX(ctx context.Context) {
	if err := iwdo.Exec(ctx); err != nil {
		panic(err)
	}
}
