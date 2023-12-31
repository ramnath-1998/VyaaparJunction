// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/ramnath.1998/vyaaparjunction/ent/order"
	"github.com/ramnath.1998/vyaaparjunction/ent/product"
	"github.com/ramnath.1998/vyaaparjunction/ent/productcategory"
	"github.com/ramnath.1998/vyaaparjunction/ent/schema"
	"github.com/ramnath.1998/vyaaparjunction/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescIdentifier is the schema descriptor for identifier field.
	orderDescIdentifier := orderFields[1].Descriptor()
	// order.DefaultIdentifier holds the default value on creation for the identifier field.
	order.DefaultIdentifier = orderDescIdentifier.Default.(func() uuid.UUID)
	// orderDescCreatedOn is the schema descriptor for createdOn field.
	orderDescCreatedOn := orderFields[2].Descriptor()
	// order.DefaultCreatedOn holds the default value on creation for the createdOn field.
	order.DefaultCreatedOn = orderDescCreatedOn.Default.(func() time.Time)
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescProductName is the schema descriptor for productName field.
	productDescProductName := productFields[0].Descriptor()
	// product.ProductNameValidator is a validator for the "productName" field. It is called by the builders before save.
	product.ProductNameValidator = productDescProductName.Validators[0].(func(string) error)
	// productDescIdentifier is the schema descriptor for Identifier field.
	productDescIdentifier := productFields[1].Descriptor()
	// product.DefaultIdentifier holds the default value on creation for the Identifier field.
	product.DefaultIdentifier = productDescIdentifier.Default.(func() uuid.UUID)
	// productDescCreatedOn is the schema descriptor for createdOn field.
	productDescCreatedOn := productFields[2].Descriptor()
	// product.DefaultCreatedOn holds the default value on creation for the createdOn field.
	product.DefaultCreatedOn = productDescCreatedOn.Default.(func() time.Time)
	productcategoryFields := schema.ProductCategory{}.Fields()
	_ = productcategoryFields
	// productcategoryDescCategoryName is the schema descriptor for categoryName field.
	productcategoryDescCategoryName := productcategoryFields[0].Descriptor()
	// productcategory.CategoryNameValidator is a validator for the "categoryName" field. It is called by the builders before save.
	productcategory.CategoryNameValidator = productcategoryDescCategoryName.Validators[0].(func(string) error)
	// productcategoryDescIdentifier is the schema descriptor for identifier field.
	productcategoryDescIdentifier := productcategoryFields[1].Descriptor()
	// productcategory.DefaultIdentifier holds the default value on creation for the identifier field.
	productcategory.DefaultIdentifier = productcategoryDescIdentifier.Default.(func() uuid.UUID)
	// productcategoryDescCreatedOn is the schema descriptor for createdOn field.
	productcategoryDescCreatedOn := productcategoryFields[2].Descriptor()
	// productcategory.DefaultCreatedOn holds the default value on creation for the createdOn field.
	productcategory.DefaultCreatedOn = productcategoryDescCreatedOn.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[2].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescIdentifier is the schema descriptor for identifier field.
	userDescIdentifier := userFields[3].Descriptor()
	// user.DefaultIdentifier holds the default value on creation for the identifier field.
	user.DefaultIdentifier = userDescIdentifier.Default.(func() uuid.UUID)
	// userDescCreatedOn is the schema descriptor for createdOn field.
	userDescCreatedOn := userFields[4].Descriptor()
	// user.DefaultCreatedOn holds the default value on creation for the createdOn field.
	user.DefaultCreatedOn = userDescCreatedOn.Default.(func() time.Time)
}
