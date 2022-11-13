// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/riadafridishibly/vault-app/ent/passworditem"
	"github.com/riadafridishibly/vault-app/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypePasswordItem = "PasswordItem"
)

// PasswordItemMutation represents an operation that mutates the PasswordItem nodes in the graph.
type PasswordItemMutation struct {
	config
	op            Op
	typ           string
	id            *int
	avatar        *string
	description   *string
	site_name     *string
	site_url      *string
	username      *string
	username_type *string
	password      *string
	tags          *[]string
	appendtags    []string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*PasswordItem, error)
	predicates    []predicate.PasswordItem
}

var _ ent.Mutation = (*PasswordItemMutation)(nil)

// passworditemOption allows management of the mutation configuration using functional options.
type passworditemOption func(*PasswordItemMutation)

// newPasswordItemMutation creates new mutation for the PasswordItem entity.
func newPasswordItemMutation(c config, op Op, opts ...passworditemOption) *PasswordItemMutation {
	m := &PasswordItemMutation{
		config:        c,
		op:            op,
		typ:           TypePasswordItem,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPasswordItemID sets the ID field of the mutation.
func withPasswordItemID(id int) passworditemOption {
	return func(m *PasswordItemMutation) {
		var (
			err   error
			once  sync.Once
			value *PasswordItem
		)
		m.oldValue = func(ctx context.Context) (*PasswordItem, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().PasswordItem.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPasswordItem sets the old PasswordItem of the mutation.
func withPasswordItem(node *PasswordItem) passworditemOption {
	return func(m *PasswordItemMutation) {
		m.oldValue = func(context.Context) (*PasswordItem, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PasswordItemMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PasswordItemMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PasswordItemMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PasswordItemMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().PasswordItem.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetAvatar sets the "avatar" field.
func (m *PasswordItemMutation) SetAvatar(s string) {
	m.avatar = &s
}

// Avatar returns the value of the "avatar" field in the mutation.
func (m *PasswordItemMutation) Avatar() (r string, exists bool) {
	v := m.avatar
	if v == nil {
		return
	}
	return *v, true
}

// OldAvatar returns the old "avatar" field's value of the PasswordItem entity.
// If the PasswordItem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PasswordItemMutation) OldAvatar(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAvatar is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAvatar requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAvatar: %w", err)
	}
	return oldValue.Avatar, nil
}

// ResetAvatar resets all changes to the "avatar" field.
func (m *PasswordItemMutation) ResetAvatar() {
	m.avatar = nil
}

// SetDescription sets the "description" field.
func (m *PasswordItemMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *PasswordItemMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the PasswordItem entity.
// If the PasswordItem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PasswordItemMutation) OldDescription(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *PasswordItemMutation) ResetDescription() {
	m.description = nil
}

// SetSiteName sets the "site_name" field.
func (m *PasswordItemMutation) SetSiteName(s string) {
	m.site_name = &s
}

// SiteName returns the value of the "site_name" field in the mutation.
func (m *PasswordItemMutation) SiteName() (r string, exists bool) {
	v := m.site_name
	if v == nil {
		return
	}
	return *v, true
}

// OldSiteName returns the old "site_name" field's value of the PasswordItem entity.
// If the PasswordItem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PasswordItemMutation) OldSiteName(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSiteName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSiteName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSiteName: %w", err)
	}
	return oldValue.SiteName, nil
}

// ResetSiteName resets all changes to the "site_name" field.
func (m *PasswordItemMutation) ResetSiteName() {
	m.site_name = nil
}

// SetSiteURL sets the "site_url" field.
func (m *PasswordItemMutation) SetSiteURL(s string) {
	m.site_url = &s
}

// SiteURL returns the value of the "site_url" field in the mutation.
func (m *PasswordItemMutation) SiteURL() (r string, exists bool) {
	v := m.site_url
	if v == nil {
		return
	}
	return *v, true
}

// OldSiteURL returns the old "site_url" field's value of the PasswordItem entity.
// If the PasswordItem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PasswordItemMutation) OldSiteURL(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSiteURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSiteURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSiteURL: %w", err)
	}
	return oldValue.SiteURL, nil
}

// ResetSiteURL resets all changes to the "site_url" field.
func (m *PasswordItemMutation) ResetSiteURL() {
	m.site_url = nil
}

// SetUsername sets the "username" field.
func (m *PasswordItemMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *PasswordItemMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the PasswordItem entity.
// If the PasswordItem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PasswordItemMutation) OldUsername(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *PasswordItemMutation) ResetUsername() {
	m.username = nil
}

// SetUsernameType sets the "username_type" field.
func (m *PasswordItemMutation) SetUsernameType(s string) {
	m.username_type = &s
}

// UsernameType returns the value of the "username_type" field in the mutation.
func (m *PasswordItemMutation) UsernameType() (r string, exists bool) {
	v := m.username_type
	if v == nil {
		return
	}
	return *v, true
}

// OldUsernameType returns the old "username_type" field's value of the PasswordItem entity.
// If the PasswordItem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PasswordItemMutation) OldUsernameType(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsernameType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsernameType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsernameType: %w", err)
	}
	return oldValue.UsernameType, nil
}

// ResetUsernameType resets all changes to the "username_type" field.
func (m *PasswordItemMutation) ResetUsernameType() {
	m.username_type = nil
}

// SetPassword sets the "password" field.
func (m *PasswordItemMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *PasswordItemMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the PasswordItem entity.
// If the PasswordItem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PasswordItemMutation) OldPassword(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *PasswordItemMutation) ResetPassword() {
	m.password = nil
}

// SetTags sets the "tags" field.
func (m *PasswordItemMutation) SetTags(s []string) {
	m.tags = &s
	m.appendtags = nil
}

// Tags returns the value of the "tags" field in the mutation.
func (m *PasswordItemMutation) Tags() (r []string, exists bool) {
	v := m.tags
	if v == nil {
		return
	}
	return *v, true
}

// OldTags returns the old "tags" field's value of the PasswordItem entity.
// If the PasswordItem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PasswordItemMutation) OldTags(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTags is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTags requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTags: %w", err)
	}
	return oldValue.Tags, nil
}

// AppendTags adds s to the "tags" field.
func (m *PasswordItemMutation) AppendTags(s []string) {
	m.appendtags = append(m.appendtags, s...)
}

// AppendedTags returns the list of values that were appended to the "tags" field in this mutation.
func (m *PasswordItemMutation) AppendedTags() ([]string, bool) {
	if len(m.appendtags) == 0 {
		return nil, false
	}
	return m.appendtags, true
}

// ClearTags clears the value of the "tags" field.
func (m *PasswordItemMutation) ClearTags() {
	m.tags = nil
	m.appendtags = nil
	m.clearedFields[passworditem.FieldTags] = struct{}{}
}

// TagsCleared returns if the "tags" field was cleared in this mutation.
func (m *PasswordItemMutation) TagsCleared() bool {
	_, ok := m.clearedFields[passworditem.FieldTags]
	return ok
}

// ResetTags resets all changes to the "tags" field.
func (m *PasswordItemMutation) ResetTags() {
	m.tags = nil
	m.appendtags = nil
	delete(m.clearedFields, passworditem.FieldTags)
}

// SetCreatedAt sets the "created_at" field.
func (m *PasswordItemMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *PasswordItemMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the PasswordItem entity.
// If the PasswordItem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PasswordItemMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *PasswordItemMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *PasswordItemMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *PasswordItemMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the PasswordItem entity.
// If the PasswordItem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PasswordItemMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *PasswordItemMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the PasswordItemMutation builder.
func (m *PasswordItemMutation) Where(ps ...predicate.PasswordItem) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *PasswordItemMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (PasswordItem).
func (m *PasswordItemMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PasswordItemMutation) Fields() []string {
	fields := make([]string, 0, 10)
	if m.avatar != nil {
		fields = append(fields, passworditem.FieldAvatar)
	}
	if m.description != nil {
		fields = append(fields, passworditem.FieldDescription)
	}
	if m.site_name != nil {
		fields = append(fields, passworditem.FieldSiteName)
	}
	if m.site_url != nil {
		fields = append(fields, passworditem.FieldSiteURL)
	}
	if m.username != nil {
		fields = append(fields, passworditem.FieldUsername)
	}
	if m.username_type != nil {
		fields = append(fields, passworditem.FieldUsernameType)
	}
	if m.password != nil {
		fields = append(fields, passworditem.FieldPassword)
	}
	if m.tags != nil {
		fields = append(fields, passworditem.FieldTags)
	}
	if m.created_at != nil {
		fields = append(fields, passworditem.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, passworditem.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PasswordItemMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case passworditem.FieldAvatar:
		return m.Avatar()
	case passworditem.FieldDescription:
		return m.Description()
	case passworditem.FieldSiteName:
		return m.SiteName()
	case passworditem.FieldSiteURL:
		return m.SiteURL()
	case passworditem.FieldUsername:
		return m.Username()
	case passworditem.FieldUsernameType:
		return m.UsernameType()
	case passworditem.FieldPassword:
		return m.Password()
	case passworditem.FieldTags:
		return m.Tags()
	case passworditem.FieldCreatedAt:
		return m.CreatedAt()
	case passworditem.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PasswordItemMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case passworditem.FieldAvatar:
		return m.OldAvatar(ctx)
	case passworditem.FieldDescription:
		return m.OldDescription(ctx)
	case passworditem.FieldSiteName:
		return m.OldSiteName(ctx)
	case passworditem.FieldSiteURL:
		return m.OldSiteURL(ctx)
	case passworditem.FieldUsername:
		return m.OldUsername(ctx)
	case passworditem.FieldUsernameType:
		return m.OldUsernameType(ctx)
	case passworditem.FieldPassword:
		return m.OldPassword(ctx)
	case passworditem.FieldTags:
		return m.OldTags(ctx)
	case passworditem.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case passworditem.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown PasswordItem field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PasswordItemMutation) SetField(name string, value ent.Value) error {
	switch name {
	case passworditem.FieldAvatar:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAvatar(v)
		return nil
	case passworditem.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case passworditem.FieldSiteName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSiteName(v)
		return nil
	case passworditem.FieldSiteURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSiteURL(v)
		return nil
	case passworditem.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case passworditem.FieldUsernameType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsernameType(v)
		return nil
	case passworditem.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case passworditem.FieldTags:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTags(v)
		return nil
	case passworditem.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case passworditem.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown PasswordItem field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PasswordItemMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PasswordItemMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PasswordItemMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown PasswordItem numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PasswordItemMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(passworditem.FieldTags) {
		fields = append(fields, passworditem.FieldTags)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PasswordItemMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PasswordItemMutation) ClearField(name string) error {
	switch name {
	case passworditem.FieldTags:
		m.ClearTags()
		return nil
	}
	return fmt.Errorf("unknown PasswordItem nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PasswordItemMutation) ResetField(name string) error {
	switch name {
	case passworditem.FieldAvatar:
		m.ResetAvatar()
		return nil
	case passworditem.FieldDescription:
		m.ResetDescription()
		return nil
	case passworditem.FieldSiteName:
		m.ResetSiteName()
		return nil
	case passworditem.FieldSiteURL:
		m.ResetSiteURL()
		return nil
	case passworditem.FieldUsername:
		m.ResetUsername()
		return nil
	case passworditem.FieldUsernameType:
		m.ResetUsernameType()
		return nil
	case passworditem.FieldPassword:
		m.ResetPassword()
		return nil
	case passworditem.FieldTags:
		m.ResetTags()
		return nil
	case passworditem.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case passworditem.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown PasswordItem field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PasswordItemMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PasswordItemMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PasswordItemMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PasswordItemMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PasswordItemMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PasswordItemMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PasswordItemMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown PasswordItem unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PasswordItemMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown PasswordItem edge %s", name)
}
