// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/item"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/maintenanceentry"
)

// MaintenanceEntryCreate is the builder for creating a MaintenanceEntry entity.
type MaintenanceEntryCreate struct {
	config
	mutation *MaintenanceEntryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mec *MaintenanceEntryCreate) SetCreatedAt(t time.Time) *MaintenanceEntryCreate {
	mec.mutation.SetCreatedAt(t)
	return mec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mec *MaintenanceEntryCreate) SetNillableCreatedAt(t *time.Time) *MaintenanceEntryCreate {
	if t != nil {
		mec.SetCreatedAt(*t)
	}
	return mec
}

// SetUpdatedAt sets the "updated_at" field.
func (mec *MaintenanceEntryCreate) SetUpdatedAt(t time.Time) *MaintenanceEntryCreate {
	mec.mutation.SetUpdatedAt(t)
	return mec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mec *MaintenanceEntryCreate) SetNillableUpdatedAt(t *time.Time) *MaintenanceEntryCreate {
	if t != nil {
		mec.SetUpdatedAt(*t)
	}
	return mec
}

// SetItemID sets the "item_id" field.
func (mec *MaintenanceEntryCreate) SetItemID(u uuid.UUID) *MaintenanceEntryCreate {
	mec.mutation.SetItemID(u)
	return mec
}

// SetDate sets the "date" field.
func (mec *MaintenanceEntryCreate) SetDate(t time.Time) *MaintenanceEntryCreate {
	mec.mutation.SetDate(t)
	return mec
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (mec *MaintenanceEntryCreate) SetNillableDate(t *time.Time) *MaintenanceEntryCreate {
	if t != nil {
		mec.SetDate(*t)
	}
	return mec
}

// SetScheduledDate sets the "scheduled_date" field.
func (mec *MaintenanceEntryCreate) SetScheduledDate(t time.Time) *MaintenanceEntryCreate {
	mec.mutation.SetScheduledDate(t)
	return mec
}

// SetNillableScheduledDate sets the "scheduled_date" field if the given value is not nil.
func (mec *MaintenanceEntryCreate) SetNillableScheduledDate(t *time.Time) *MaintenanceEntryCreate {
	if t != nil {
		mec.SetScheduledDate(*t)
	}
	return mec
}

// SetName sets the "name" field.
func (mec *MaintenanceEntryCreate) SetName(s string) *MaintenanceEntryCreate {
	mec.mutation.SetName(s)
	return mec
}

// SetDescription sets the "description" field.
func (mec *MaintenanceEntryCreate) SetDescription(s string) *MaintenanceEntryCreate {
	mec.mutation.SetDescription(s)
	return mec
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mec *MaintenanceEntryCreate) SetNillableDescription(s *string) *MaintenanceEntryCreate {
	if s != nil {
		mec.SetDescription(*s)
	}
	return mec
}

// SetCost sets the "cost" field.
func (mec *MaintenanceEntryCreate) SetCost(f float64) *MaintenanceEntryCreate {
	mec.mutation.SetCost(f)
	return mec
}

// SetNillableCost sets the "cost" field if the given value is not nil.
func (mec *MaintenanceEntryCreate) SetNillableCost(f *float64) *MaintenanceEntryCreate {
	if f != nil {
		mec.SetCost(*f)
	}
	return mec
}

// SetMeasurement sets the "measurement" field.
func (mec *MaintenanceEntryCreate) SetMeasurement(s string) *MaintenanceEntryCreate {
	mec.mutation.SetMeasurement(s)
	return mec
}

// SetNillableMeasurement sets the "measurement" field if the given value is not nil.
func (mec *MaintenanceEntryCreate) SetNillableMeasurement(s *string) *MaintenanceEntryCreate {
	if s != nil {
		mec.SetMeasurement(*s)
	}
	return mec
}

// SetID sets the "id" field.
func (mec *MaintenanceEntryCreate) SetID(u uuid.UUID) *MaintenanceEntryCreate {
	mec.mutation.SetID(u)
	return mec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mec *MaintenanceEntryCreate) SetNillableID(u *uuid.UUID) *MaintenanceEntryCreate {
	if u != nil {
		mec.SetID(*u)
	}
	return mec
}

// SetItem sets the "item" edge to the Item entity.
func (mec *MaintenanceEntryCreate) SetItem(i *Item) *MaintenanceEntryCreate {
	return mec.SetItemID(i.ID)
}

// Mutation returns the MaintenanceEntryMutation object of the builder.
func (mec *MaintenanceEntryCreate) Mutation() *MaintenanceEntryMutation {
	return mec.mutation
}

// Save creates the MaintenanceEntry in the database.
func (mec *MaintenanceEntryCreate) Save(ctx context.Context) (*MaintenanceEntry, error) {
	mec.defaults()
	return withHooks(ctx, mec.sqlSave, mec.mutation, mec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mec *MaintenanceEntryCreate) SaveX(ctx context.Context) *MaintenanceEntry {
	v, err := mec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mec *MaintenanceEntryCreate) Exec(ctx context.Context) error {
	_, err := mec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mec *MaintenanceEntryCreate) ExecX(ctx context.Context) {
	if err := mec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mec *MaintenanceEntryCreate) defaults() {
	if _, ok := mec.mutation.CreatedAt(); !ok {
		v := maintenanceentry.DefaultCreatedAt()
		mec.mutation.SetCreatedAt(v)
	}
	if _, ok := mec.mutation.UpdatedAt(); !ok {
		v := maintenanceentry.DefaultUpdatedAt()
		mec.mutation.SetUpdatedAt(v)
	}
	if _, ok := mec.mutation.Cost(); !ok {
		v := maintenanceentry.DefaultCost
		mec.mutation.SetCost(v)
	}
	if _, ok := mec.mutation.ID(); !ok {
		v := maintenanceentry.DefaultID()
		mec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mec *MaintenanceEntryCreate) check() error {
	if _, ok := mec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MaintenanceEntry.created_at"`)}
	}
	if _, ok := mec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MaintenanceEntry.updated_at"`)}
	}
	if _, ok := mec.mutation.ItemID(); !ok {
		return &ValidationError{Name: "item_id", err: errors.New(`ent: missing required field "MaintenanceEntry.item_id"`)}
	}
	if _, ok := mec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "MaintenanceEntry.name"`)}
	}
	if v, ok := mec.mutation.Name(); ok {
		if err := maintenanceentry.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MaintenanceEntry.name": %w`, err)}
		}
	}
	if v, ok := mec.mutation.Description(); ok {
		if err := maintenanceentry.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "MaintenanceEntry.description": %w`, err)}
		}
	}
	if _, ok := mec.mutation.Cost(); !ok {
		return &ValidationError{Name: "cost", err: errors.New(`ent: missing required field "MaintenanceEntry.cost"`)}
	}
	if len(mec.mutation.ItemIDs()) == 0 {
		return &ValidationError{Name: "item", err: errors.New(`ent: missing required edge "MaintenanceEntry.item"`)}
	}
	return nil
}

func (mec *MaintenanceEntryCreate) sqlSave(ctx context.Context) (*MaintenanceEntry, error) {
	if err := mec.check(); err != nil {
		return nil, err
	}
	_node, _spec := mec.createSpec()
	if err := sqlgraph.CreateNode(ctx, mec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	mec.mutation.id = &_node.ID
	mec.mutation.done = true
	return _node, nil
}

func (mec *MaintenanceEntryCreate) createSpec() (*MaintenanceEntry, *sqlgraph.CreateSpec) {
	var (
		_node = &MaintenanceEntry{config: mec.config}
		_spec = sqlgraph.NewCreateSpec(maintenanceentry.Table, sqlgraph.NewFieldSpec(maintenanceentry.FieldID, field.TypeUUID))
	)
	if id, ok := mec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mec.mutation.CreatedAt(); ok {
		_spec.SetField(maintenanceentry.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mec.mutation.UpdatedAt(); ok {
		_spec.SetField(maintenanceentry.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mec.mutation.Date(); ok {
		_spec.SetField(maintenanceentry.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	if value, ok := mec.mutation.ScheduledDate(); ok {
		_spec.SetField(maintenanceentry.FieldScheduledDate, field.TypeTime, value)
		_node.ScheduledDate = value
	}
	if value, ok := mec.mutation.Name(); ok {
		_spec.SetField(maintenanceentry.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mec.mutation.Description(); ok {
		_spec.SetField(maintenanceentry.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := mec.mutation.Cost(); ok {
		_spec.SetField(maintenanceentry.FieldCost, field.TypeFloat64, value)
		_node.Cost = value
	}
	if value, ok := mec.mutation.Measurement(); ok {
		_spec.SetField(maintenanceentry.FieldMeasurement, field.TypeString, value)
		_node.Measurement = value
	}
	if nodes := mec.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   maintenanceentry.ItemTable,
			Columns: []string{maintenanceentry.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ItemID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MaintenanceEntryCreateBulk is the builder for creating many MaintenanceEntry entities in bulk.
type MaintenanceEntryCreateBulk struct {
	config
	err      error
	builders []*MaintenanceEntryCreate
}

// Save creates the MaintenanceEntry entities in the database.
func (mecb *MaintenanceEntryCreateBulk) Save(ctx context.Context) ([]*MaintenanceEntry, error) {
	if mecb.err != nil {
		return nil, mecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mecb.builders))
	nodes := make([]*MaintenanceEntry, len(mecb.builders))
	mutators := make([]Mutator, len(mecb.builders))
	for i := range mecb.builders {
		func(i int, root context.Context) {
			builder := mecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MaintenanceEntryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, mecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mecb *MaintenanceEntryCreateBulk) SaveX(ctx context.Context) []*MaintenanceEntry {
	v, err := mecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mecb *MaintenanceEntryCreateBulk) Exec(ctx context.Context) error {
	_, err := mecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mecb *MaintenanceEntryCreateBulk) ExecX(ctx context.Context) {
	if err := mecb.Exec(ctx); err != nil {
		panic(err)
	}
}
