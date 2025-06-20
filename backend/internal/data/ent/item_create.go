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
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/attachment"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/group"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/item"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/itemfield"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/label"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/location"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/maintenanceentry"
)

// ItemCreate is the builder for creating a Item entity.
type ItemCreate struct {
	config
	mutation *ItemMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ic *ItemCreate) SetCreatedAt(t time.Time) *ItemCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *ItemCreate) SetNillableCreatedAt(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *ItemCreate) SetUpdatedAt(t time.Time) *ItemCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *ItemCreate) SetNillableUpdatedAt(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetName sets the "name" field.
func (ic *ItemCreate) SetName(s string) *ItemCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetDescription sets the "description" field.
func (ic *ItemCreate) SetDescription(s string) *ItemCreate {
	ic.mutation.SetDescription(s)
	return ic
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ic *ItemCreate) SetNillableDescription(s *string) *ItemCreate {
	if s != nil {
		ic.SetDescription(*s)
	}
	return ic
}

// SetImportRef sets the "import_ref" field.
func (ic *ItemCreate) SetImportRef(s string) *ItemCreate {
	ic.mutation.SetImportRef(s)
	return ic
}

// SetNillableImportRef sets the "import_ref" field if the given value is not nil.
func (ic *ItemCreate) SetNillableImportRef(s *string) *ItemCreate {
	if s != nil {
		ic.SetImportRef(*s)
	}
	return ic
}

// SetNotes sets the "notes" field.
func (ic *ItemCreate) SetNotes(s string) *ItemCreate {
	ic.mutation.SetNotes(s)
	return ic
}

// SetNillableNotes sets the "notes" field if the given value is not nil.
func (ic *ItemCreate) SetNillableNotes(s *string) *ItemCreate {
	if s != nil {
		ic.SetNotes(*s)
	}
	return ic
}

// SetQuantity sets the "quantity" field.
func (ic *ItemCreate) SetQuantity(i int) *ItemCreate {
	ic.mutation.SetQuantity(i)
	return ic
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (ic *ItemCreate) SetNillableQuantity(i *int) *ItemCreate {
	if i != nil {
		ic.SetQuantity(*i)
	}
	return ic
}

// SetInsured sets the "insured" field.
func (ic *ItemCreate) SetInsured(b bool) *ItemCreate {
	ic.mutation.SetInsured(b)
	return ic
}

// SetNillableInsured sets the "insured" field if the given value is not nil.
func (ic *ItemCreate) SetNillableInsured(b *bool) *ItemCreate {
	if b != nil {
		ic.SetInsured(*b)
	}
	return ic
}

// SetArchived sets the "archived" field.
func (ic *ItemCreate) SetArchived(b bool) *ItemCreate {
	ic.mutation.SetArchived(b)
	return ic
}

// SetNillableArchived sets the "archived" field if the given value is not nil.
func (ic *ItemCreate) SetNillableArchived(b *bool) *ItemCreate {
	if b != nil {
		ic.SetArchived(*b)
	}
	return ic
}

// SetArchivedAt sets the "archived_at" field.
func (ic *ItemCreate) SetArchivedAt(t time.Time) *ItemCreate {
	ic.mutation.SetArchivedAt(t)
	return ic
}

// SetNillableArchivedAt sets the "archived_at" field if the given value is not nil.
func (ic *ItemCreate) SetNillableArchivedAt(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetArchivedAt(*t)
	}
	return ic
}

// SetAssetID sets the "asset_id" field.
func (ic *ItemCreate) SetAssetID(i int) *ItemCreate {
	ic.mutation.SetAssetID(i)
	return ic
}

// SetNillableAssetID sets the "asset_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableAssetID(i *int) *ItemCreate {
	if i != nil {
		ic.SetAssetID(*i)
	}
	return ic
}

// SetSyncChildItemsLocations sets the "sync_child_items_locations" field.
func (ic *ItemCreate) SetSyncChildItemsLocations(b bool) *ItemCreate {
	ic.mutation.SetSyncChildItemsLocations(b)
	return ic
}

// SetNillableSyncChildItemsLocations sets the "sync_child_items_locations" field if the given value is not nil.
func (ic *ItemCreate) SetNillableSyncChildItemsLocations(b *bool) *ItemCreate {
	if b != nil {
		ic.SetSyncChildItemsLocations(*b)
	}
	return ic
}

// SetSerialNumber sets the "serial_number" field.
func (ic *ItemCreate) SetSerialNumber(s string) *ItemCreate {
	ic.mutation.SetSerialNumber(s)
	return ic
}

// SetNillableSerialNumber sets the "serial_number" field if the given value is not nil.
func (ic *ItemCreate) SetNillableSerialNumber(s *string) *ItemCreate {
	if s != nil {
		ic.SetSerialNumber(*s)
	}
	return ic
}

// SetModelNumber sets the "model_number" field.
func (ic *ItemCreate) SetModelNumber(s string) *ItemCreate {
	ic.mutation.SetModelNumber(s)
	return ic
}

// SetNillableModelNumber sets the "model_number" field if the given value is not nil.
func (ic *ItemCreate) SetNillableModelNumber(s *string) *ItemCreate {
	if s != nil {
		ic.SetModelNumber(*s)
	}
	return ic
}

// SetManufacturer sets the "manufacturer" field.
func (ic *ItemCreate) SetManufacturer(s string) *ItemCreate {
	ic.mutation.SetManufacturer(s)
	return ic
}

// SetNillableManufacturer sets the "manufacturer" field if the given value is not nil.
func (ic *ItemCreate) SetNillableManufacturer(s *string) *ItemCreate {
	if s != nil {
		ic.SetManufacturer(*s)
	}
	return ic
}

// SetLifetimeWarranty sets the "lifetime_warranty" field.
func (ic *ItemCreate) SetLifetimeWarranty(b bool) *ItemCreate {
	ic.mutation.SetLifetimeWarranty(b)
	return ic
}

// SetNillableLifetimeWarranty sets the "lifetime_warranty" field if the given value is not nil.
func (ic *ItemCreate) SetNillableLifetimeWarranty(b *bool) *ItemCreate {
	if b != nil {
		ic.SetLifetimeWarranty(*b)
	}
	return ic
}

// SetWarrantyExpires sets the "warranty_expires" field.
func (ic *ItemCreate) SetWarrantyExpires(t time.Time) *ItemCreate {
	ic.mutation.SetWarrantyExpires(t)
	return ic
}

// SetNillableWarrantyExpires sets the "warranty_expires" field if the given value is not nil.
func (ic *ItemCreate) SetNillableWarrantyExpires(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetWarrantyExpires(*t)
	}
	return ic
}

// SetWarrantyDetails sets the "warranty_details" field.
func (ic *ItemCreate) SetWarrantyDetails(s string) *ItemCreate {
	ic.mutation.SetWarrantyDetails(s)
	return ic
}

// SetNillableWarrantyDetails sets the "warranty_details" field if the given value is not nil.
func (ic *ItemCreate) SetNillableWarrantyDetails(s *string) *ItemCreate {
	if s != nil {
		ic.SetWarrantyDetails(*s)
	}
	return ic
}

// SetPurchaseTime sets the "purchase_time" field.
func (ic *ItemCreate) SetPurchaseTime(t time.Time) *ItemCreate {
	ic.mutation.SetPurchaseTime(t)
	return ic
}

// SetNillablePurchaseTime sets the "purchase_time" field if the given value is not nil.
func (ic *ItemCreate) SetNillablePurchaseTime(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetPurchaseTime(*t)
	}
	return ic
}

// SetPurchaseFrom sets the "purchase_from" field.
func (ic *ItemCreate) SetPurchaseFrom(s string) *ItemCreate {
	ic.mutation.SetPurchaseFrom(s)
	return ic
}

// SetNillablePurchaseFrom sets the "purchase_from" field if the given value is not nil.
func (ic *ItemCreate) SetNillablePurchaseFrom(s *string) *ItemCreate {
	if s != nil {
		ic.SetPurchaseFrom(*s)
	}
	return ic
}

// SetPurchasePrice sets the "purchase_price" field.
func (ic *ItemCreate) SetPurchasePrice(f float64) *ItemCreate {
	ic.mutation.SetPurchasePrice(f)
	return ic
}

// SetNillablePurchasePrice sets the "purchase_price" field if the given value is not nil.
func (ic *ItemCreate) SetNillablePurchasePrice(f *float64) *ItemCreate {
	if f != nil {
		ic.SetPurchasePrice(*f)
	}
	return ic
}

// SetSoldTime sets the "sold_time" field.
func (ic *ItemCreate) SetSoldTime(t time.Time) *ItemCreate {
	ic.mutation.SetSoldTime(t)
	return ic
}

// SetNillableSoldTime sets the "sold_time" field if the given value is not nil.
func (ic *ItemCreate) SetNillableSoldTime(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetSoldTime(*t)
	}
	return ic
}

// SetSoldTo sets the "sold_to" field.
func (ic *ItemCreate) SetSoldTo(s string) *ItemCreate {
	ic.mutation.SetSoldTo(s)
	return ic
}

// SetNillableSoldTo sets the "sold_to" field if the given value is not nil.
func (ic *ItemCreate) SetNillableSoldTo(s *string) *ItemCreate {
	if s != nil {
		ic.SetSoldTo(*s)
	}
	return ic
}

// SetSoldPrice sets the "sold_price" field.
func (ic *ItemCreate) SetSoldPrice(f float64) *ItemCreate {
	ic.mutation.SetSoldPrice(f)
	return ic
}

// SetNillableSoldPrice sets the "sold_price" field if the given value is not nil.
func (ic *ItemCreate) SetNillableSoldPrice(f *float64) *ItemCreate {
	if f != nil {
		ic.SetSoldPrice(*f)
	}
	return ic
}

// SetSoldNotes sets the "sold_notes" field.
func (ic *ItemCreate) SetSoldNotes(s string) *ItemCreate {
	ic.mutation.SetSoldNotes(s)
	return ic
}

// SetNillableSoldNotes sets the "sold_notes" field if the given value is not nil.
func (ic *ItemCreate) SetNillableSoldNotes(s *string) *ItemCreate {
	if s != nil {
		ic.SetSoldNotes(*s)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *ItemCreate) SetID(u uuid.UUID) *ItemCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableID(u *uuid.UUID) *ItemCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (ic *ItemCreate) SetGroupID(id uuid.UUID) *ItemCreate {
	ic.mutation.SetGroupID(id)
	return ic
}

// SetGroup sets the "group" edge to the Group entity.
func (ic *ItemCreate) SetGroup(g *Group) *ItemCreate {
	return ic.SetGroupID(g.ID)
}

// SetParentID sets the "parent" edge to the Item entity by ID.
func (ic *ItemCreate) SetParentID(id uuid.UUID) *ItemCreate {
	ic.mutation.SetParentID(id)
	return ic
}

// SetNillableParentID sets the "parent" edge to the Item entity by ID if the given value is not nil.
func (ic *ItemCreate) SetNillableParentID(id *uuid.UUID) *ItemCreate {
	if id != nil {
		ic = ic.SetParentID(*id)
	}
	return ic
}

// SetParent sets the "parent" edge to the Item entity.
func (ic *ItemCreate) SetParent(i *Item) *ItemCreate {
	return ic.SetParentID(i.ID)
}

// AddChildIDs adds the "children" edge to the Item entity by IDs.
func (ic *ItemCreate) AddChildIDs(ids ...uuid.UUID) *ItemCreate {
	ic.mutation.AddChildIDs(ids...)
	return ic
}

// AddChildren adds the "children" edges to the Item entity.
func (ic *ItemCreate) AddChildren(i ...*Item) *ItemCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ic.AddChildIDs(ids...)
}

// AddLabelIDs adds the "label" edge to the Label entity by IDs.
func (ic *ItemCreate) AddLabelIDs(ids ...uuid.UUID) *ItemCreate {
	ic.mutation.AddLabelIDs(ids...)
	return ic
}

// AddLabel adds the "label" edges to the Label entity.
func (ic *ItemCreate) AddLabel(l ...*Label) *ItemCreate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ic.AddLabelIDs(ids...)
}

// SetLocationID sets the "location" edge to the Location entity by ID.
func (ic *ItemCreate) SetLocationID(id uuid.UUID) *ItemCreate {
	ic.mutation.SetLocationID(id)
	return ic
}

// SetNillableLocationID sets the "location" edge to the Location entity by ID if the given value is not nil.
func (ic *ItemCreate) SetNillableLocationID(id *uuid.UUID) *ItemCreate {
	if id != nil {
		ic = ic.SetLocationID(*id)
	}
	return ic
}

// SetLocation sets the "location" edge to the Location entity.
func (ic *ItemCreate) SetLocation(l *Location) *ItemCreate {
	return ic.SetLocationID(l.ID)
}

// AddFieldIDs adds the "fields" edge to the ItemField entity by IDs.
func (ic *ItemCreate) AddFieldIDs(ids ...uuid.UUID) *ItemCreate {
	ic.mutation.AddFieldIDs(ids...)
	return ic
}

// AddFields adds the "fields" edges to the ItemField entity.
func (ic *ItemCreate) AddFields(i ...*ItemField) *ItemCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ic.AddFieldIDs(ids...)
}

// AddMaintenanceEntryIDs adds the "maintenance_entries" edge to the MaintenanceEntry entity by IDs.
func (ic *ItemCreate) AddMaintenanceEntryIDs(ids ...uuid.UUID) *ItemCreate {
	ic.mutation.AddMaintenanceEntryIDs(ids...)
	return ic
}

// AddMaintenanceEntries adds the "maintenance_entries" edges to the MaintenanceEntry entity.
func (ic *ItemCreate) AddMaintenanceEntries(m ...*MaintenanceEntry) *ItemCreate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ic.AddMaintenanceEntryIDs(ids...)
}

// AddAttachmentIDs adds the "attachments" edge to the Attachment entity by IDs.
func (ic *ItemCreate) AddAttachmentIDs(ids ...uuid.UUID) *ItemCreate {
	ic.mutation.AddAttachmentIDs(ids...)
	return ic
}

// AddAttachments adds the "attachments" edges to the Attachment entity.
func (ic *ItemCreate) AddAttachments(a ...*Attachment) *ItemCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ic.AddAttachmentIDs(ids...)
}

// Mutation returns the ItemMutation object of the builder.
func (ic *ItemCreate) Mutation() *ItemMutation {
	return ic.mutation
}

// Save creates the Item in the database.
func (ic *ItemCreate) Save(ctx context.Context) (*Item, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ItemCreate) SaveX(ctx context.Context) *Item {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ItemCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ItemCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *ItemCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := item.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := item.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ic.mutation.Quantity(); !ok {
		v := item.DefaultQuantity
		ic.mutation.SetQuantity(v)
	}
	if _, ok := ic.mutation.Insured(); !ok {
		v := item.DefaultInsured
		ic.mutation.SetInsured(v)
	}
	if _, ok := ic.mutation.Archived(); !ok {
		v := item.DefaultArchived
		ic.mutation.SetArchived(v)
	}
	if _, ok := ic.mutation.AssetID(); !ok {
		v := item.DefaultAssetID
		ic.mutation.SetAssetID(v)
	}
	if _, ok := ic.mutation.SyncChildItemsLocations(); !ok {
		v := item.DefaultSyncChildItemsLocations
		ic.mutation.SetSyncChildItemsLocations(v)
	}
	if _, ok := ic.mutation.LifetimeWarranty(); !ok {
		v := item.DefaultLifetimeWarranty
		ic.mutation.SetLifetimeWarranty(v)
	}
	if _, ok := ic.mutation.PurchasePrice(); !ok {
		v := item.DefaultPurchasePrice
		ic.mutation.SetPurchasePrice(v)
	}
	if _, ok := ic.mutation.SoldPrice(); !ok {
		v := item.DefaultSoldPrice
		ic.mutation.SetSoldPrice(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := item.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ItemCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Item.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Item.updated_at"`)}
	}
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Item.name"`)}
	}
	if v, ok := ic.mutation.Name(); ok {
		if err := item.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Item.name": %w`, err)}
		}
	}
	if v, ok := ic.mutation.Description(); ok {
		if err := item.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Item.description": %w`, err)}
		}
	}
	if v, ok := ic.mutation.ImportRef(); ok {
		if err := item.ImportRefValidator(v); err != nil {
			return &ValidationError{Name: "import_ref", err: fmt.Errorf(`ent: validator failed for field "Item.import_ref": %w`, err)}
		}
	}
	if v, ok := ic.mutation.Notes(); ok {
		if err := item.NotesValidator(v); err != nil {
			return &ValidationError{Name: "notes", err: fmt.Errorf(`ent: validator failed for field "Item.notes": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "Item.quantity"`)}
	}
	if _, ok := ic.mutation.Insured(); !ok {
		return &ValidationError{Name: "insured", err: errors.New(`ent: missing required field "Item.insured"`)}
	}
	if _, ok := ic.mutation.Archived(); !ok {
		return &ValidationError{Name: "archived", err: errors.New(`ent: missing required field "Item.archived"`)}
	}
	if _, ok := ic.mutation.AssetID(); !ok {
		return &ValidationError{Name: "asset_id", err: errors.New(`ent: missing required field "Item.asset_id"`)}
	}
	if _, ok := ic.mutation.SyncChildItemsLocations(); !ok {
		return &ValidationError{Name: "sync_child_items_locations", err: errors.New(`ent: missing required field "Item.sync_child_items_locations"`)}
	}
	if v, ok := ic.mutation.SerialNumber(); ok {
		if err := item.SerialNumberValidator(v); err != nil {
			return &ValidationError{Name: "serial_number", err: fmt.Errorf(`ent: validator failed for field "Item.serial_number": %w`, err)}
		}
	}
	if v, ok := ic.mutation.ModelNumber(); ok {
		if err := item.ModelNumberValidator(v); err != nil {
			return &ValidationError{Name: "model_number", err: fmt.Errorf(`ent: validator failed for field "Item.model_number": %w`, err)}
		}
	}
	if v, ok := ic.mutation.Manufacturer(); ok {
		if err := item.ManufacturerValidator(v); err != nil {
			return &ValidationError{Name: "manufacturer", err: fmt.Errorf(`ent: validator failed for field "Item.manufacturer": %w`, err)}
		}
	}
	if _, ok := ic.mutation.LifetimeWarranty(); !ok {
		return &ValidationError{Name: "lifetime_warranty", err: errors.New(`ent: missing required field "Item.lifetime_warranty"`)}
	}
	if v, ok := ic.mutation.WarrantyDetails(); ok {
		if err := item.WarrantyDetailsValidator(v); err != nil {
			return &ValidationError{Name: "warranty_details", err: fmt.Errorf(`ent: validator failed for field "Item.warranty_details": %w`, err)}
		}
	}
	if _, ok := ic.mutation.PurchasePrice(); !ok {
		return &ValidationError{Name: "purchase_price", err: errors.New(`ent: missing required field "Item.purchase_price"`)}
	}
	if _, ok := ic.mutation.SoldPrice(); !ok {
		return &ValidationError{Name: "sold_price", err: errors.New(`ent: missing required field "Item.sold_price"`)}
	}
	if v, ok := ic.mutation.SoldNotes(); ok {
		if err := item.SoldNotesValidator(v); err != nil {
			return &ValidationError{Name: "sold_notes", err: fmt.Errorf(`ent: validator failed for field "Item.sold_notes": %w`, err)}
		}
	}
	if len(ic.mutation.GroupIDs()) == 0 {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "Item.group"`)}
	}
	return nil
}

func (ic *ItemCreate) sqlSave(ctx context.Context) (*Item, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
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
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *ItemCreate) createSpec() (*Item, *sqlgraph.CreateSpec) {
	var (
		_node = &Item{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(item.Table, sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID))
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(item.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(item.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(item.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.Description(); ok {
		_spec.SetField(item.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ic.mutation.ImportRef(); ok {
		_spec.SetField(item.FieldImportRef, field.TypeString, value)
		_node.ImportRef = value
	}
	if value, ok := ic.mutation.Notes(); ok {
		_spec.SetField(item.FieldNotes, field.TypeString, value)
		_node.Notes = value
	}
	if value, ok := ic.mutation.Quantity(); ok {
		_spec.SetField(item.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if value, ok := ic.mutation.Insured(); ok {
		_spec.SetField(item.FieldInsured, field.TypeBool, value)
		_node.Insured = value
	}
	if value, ok := ic.mutation.Archived(); ok {
		_spec.SetField(item.FieldArchived, field.TypeBool, value)
		_node.Archived = value
	}
	if value, ok := ic.mutation.ArchivedAt(); ok {
		_spec.SetField(item.FieldArchivedAt, field.TypeTime, value)
		_node.ArchivedAt = &value
	}
	if value, ok := ic.mutation.AssetID(); ok {
		_spec.SetField(item.FieldAssetID, field.TypeInt, value)
		_node.AssetID = value
	}
	if value, ok := ic.mutation.SyncChildItemsLocations(); ok {
		_spec.SetField(item.FieldSyncChildItemsLocations, field.TypeBool, value)
		_node.SyncChildItemsLocations = value
	}
	if value, ok := ic.mutation.SerialNumber(); ok {
		_spec.SetField(item.FieldSerialNumber, field.TypeString, value)
		_node.SerialNumber = value
	}
	if value, ok := ic.mutation.ModelNumber(); ok {
		_spec.SetField(item.FieldModelNumber, field.TypeString, value)
		_node.ModelNumber = value
	}
	if value, ok := ic.mutation.Manufacturer(); ok {
		_spec.SetField(item.FieldManufacturer, field.TypeString, value)
		_node.Manufacturer = value
	}
	if value, ok := ic.mutation.LifetimeWarranty(); ok {
		_spec.SetField(item.FieldLifetimeWarranty, field.TypeBool, value)
		_node.LifetimeWarranty = value
	}
	if value, ok := ic.mutation.WarrantyExpires(); ok {
		_spec.SetField(item.FieldWarrantyExpires, field.TypeTime, value)
		_node.WarrantyExpires = value
	}
	if value, ok := ic.mutation.WarrantyDetails(); ok {
		_spec.SetField(item.FieldWarrantyDetails, field.TypeString, value)
		_node.WarrantyDetails = value
	}
	if value, ok := ic.mutation.PurchaseTime(); ok {
		_spec.SetField(item.FieldPurchaseTime, field.TypeTime, value)
		_node.PurchaseTime = value
	}
	if value, ok := ic.mutation.PurchaseFrom(); ok {
		_spec.SetField(item.FieldPurchaseFrom, field.TypeString, value)
		_node.PurchaseFrom = value
	}
	if value, ok := ic.mutation.PurchasePrice(); ok {
		_spec.SetField(item.FieldPurchasePrice, field.TypeFloat64, value)
		_node.PurchasePrice = value
	}
	if value, ok := ic.mutation.SoldTime(); ok {
		_spec.SetField(item.FieldSoldTime, field.TypeTime, value)
		_node.SoldTime = value
	}
	if value, ok := ic.mutation.SoldTo(); ok {
		_spec.SetField(item.FieldSoldTo, field.TypeString, value)
		_node.SoldTo = value
	}
	if value, ok := ic.mutation.SoldPrice(); ok {
		_spec.SetField(item.FieldSoldPrice, field.TypeFloat64, value)
		_node.SoldPrice = value
	}
	if value, ok := ic.mutation.SoldNotes(); ok {
		_spec.SetField(item.FieldSoldNotes, field.TypeString, value)
		_node.SoldNotes = value
	}
	if nodes := ic.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.GroupTable,
			Columns: []string{item.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_items = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.ParentTable,
			Columns: []string{item.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.item_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   item.ChildrenTable,
			Columns: []string{item.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.LabelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.LabelTable,
			Columns: item.LabelPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(label.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.LocationTable,
			Columns: []string{item.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.location_items = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.FieldsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   item.FieldsTable,
			Columns: []string{item.FieldsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(itemfield.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.MaintenanceEntriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   item.MaintenanceEntriesTable,
			Columns: []string{item.MaintenanceEntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(maintenanceentry.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.AttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   item.AttachmentsTable,
			Columns: []string{item.AttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ItemCreateBulk is the builder for creating many Item entities in bulk.
type ItemCreateBulk struct {
	config
	err      error
	builders []*ItemCreate
}

// Save creates the Item entities in the database.
func (icb *ItemCreateBulk) Save(ctx context.Context) ([]*Item, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Item, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ItemCreateBulk) SaveX(ctx context.Context) []*Item {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ItemCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ItemCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
