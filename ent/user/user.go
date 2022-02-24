// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldPass holds the string denoting the pass field in the database.
	FieldPass = "pass"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldLogin holds the string denoting the login field in the database.
	FieldLogin = "login"
	// EdgeDetail holds the string denoting the detail edge name in mutations.
	EdgeDetail = "detail"
	// Table holds the table name of the user in the database.
	Table = "user"
	// DetailTable is the table that holds the detail relation/edge.
	DetailTable = "user_detail"
	// DetailInverseTable is the table name for the UserDetail entity.
	// It exists in this package in order to avoid circular dependency with the "userdetail" package.
	DetailInverseTable = "user_detail"
	// DetailColumn is the table column denoting the detail relation/edge.
	DetailColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUsername,
	FieldMobile,
	FieldPass,
	FieldUUID,
	FieldRole,
	FieldActive,
	FieldState,
	FieldLogin,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// DefaultLogin holds the default value on creation for the "login" field.
	DefaultLogin func() time.Time
)

// State defines the type for the "state" enum field.
type State string

// State values.
const (
	StateOn  State = "on"
	StateOff State = "off"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateOn, StateOff:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for state field: %q", s)
	}
}