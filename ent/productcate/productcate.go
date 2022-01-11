// Code generated by entc, DO NOT EDIT.

package productcate

import (
	"time"
)

const (
	// Label holds the string label denoting the productcate type in the database.
	Label = "product_cate"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// Table holds the table name of the productcate in the database.
	Table = "product_cate"
	// ProductsTable is the table that holds the products relation/edge.
	ProductsTable = "product"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "product"
	// ProductsColumn is the table column denoting the products relation/edge.
	ProductsColumn = "cate_id"
)

// Columns holds all SQL columns for productcate fields.
var Columns = []string{
	FieldID,
	FieldParentID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldSort,
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
)
