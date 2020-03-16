// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/facebookincubator/ent/entc/integration/ent/card"
	"github.com/facebookincubator/ent/entc/integration/ent/schema"

	"github.com/facebookincubator/ent/entc/integration/ent/fieldtype"

	"github.com/facebookincubator/ent/entc/integration/ent/file"

	"github.com/facebookincubator/ent/entc/integration/ent/group"

	"github.com/facebookincubator/ent/entc/integration/ent/groupinfo"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/entc/integration/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	cardMixin := schema.Card{}.Mixin()
	cardMixinFields := [...][]ent.Field{
		cardMixin[0].Fields(),
	}
	cardFields := schema.Card{}.Fields()
	_ = cardFields
	// cardDescCreateTime is the schema descriptor for create_time field.
	cardDescCreateTime := cardMixinFields[0][0].Descriptor()
	// card.DefaultCreateTime holds the default value on creation for the create_time field.
	card.DefaultCreateTime = cardDescCreateTime.Default.(func() time.Time)
	// cardDescUpdateTime is the schema descriptor for update_time field.
	cardDescUpdateTime := cardMixinFields[0][1].Descriptor()
	// card.DefaultUpdateTime holds the default value on creation for the update_time field.
	card.DefaultUpdateTime = cardDescUpdateTime.Default.(func() time.Time)
	// card.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	card.UpdateDefaultUpdateTime = cardDescUpdateTime.UpdateDefault.(func() time.Time)
	// cardDescNumber is the schema descriptor for number field.
	cardDescNumber := cardFields[0].Descriptor()
	// card.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	card.NumberValidator = cardDescNumber.Validators[0].(func(string) error)
	// cardDescName is the schema descriptor for name field.
	cardDescName := cardFields[1].Descriptor()
	// card.NameValidator is a validator for the "name" field. It is called by the builders before save.
	card.NameValidator = cardDescName.Validators[0].(func(string) error)
	fieldtypeFields := schema.FieldType{}.Fields()
	_ = fieldtypeFields
	// fieldtypeDescValidateOptionalInt32 is the schema descriptor for validate_optional_int32 field.
	fieldtypeDescValidateOptionalInt32 := fieldtypeFields[15].Descriptor()
	// fieldtype.ValidateOptionalInt32Validator is a validator for the "validate_optional_int32" field. It is called by the builders before save.
	fieldtype.ValidateOptionalInt32Validator = fieldtypeDescValidateOptionalInt32.Validators[0].(func(int32) error)
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescSize is the schema descriptor for size field.
	fileDescSize := fileFields[0].Descriptor()
	// file.DefaultSize holds the default value on creation for the size field.
	file.DefaultSize = fileDescSize.Default.(int)
	// file.SizeValidator is a validator for the "size" field. It is called by the builders before save.
	file.SizeValidator = fileDescSize.Validators[0].(func(int) error)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescActive is the schema descriptor for active field.
	groupDescActive := groupFields[0].Descriptor()
	// group.DefaultActive holds the default value on creation for the active field.
	group.DefaultActive = groupDescActive.Default.(bool)
	// groupDescType is the schema descriptor for type field.
	groupDescType := groupFields[2].Descriptor()
	// group.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	group.TypeValidator = func() func(string) error {
		validators := groupDescType.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_type string) error {
			for _, fn := range fns {
				if err := fn(_type); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// groupDescMaxUsers is the schema descriptor for max_users field.
	groupDescMaxUsers := groupFields[3].Descriptor()
	// group.DefaultMaxUsers holds the default value on creation for the max_users field.
	group.DefaultMaxUsers = groupDescMaxUsers.Default.(int)
	// group.MaxUsersValidator is a validator for the "max_users" field. It is called by the builders before save.
	group.MaxUsersValidator = groupDescMaxUsers.Validators[0].(func(int) error)
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[4].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = func() func(string) error {
		validators := groupDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	groupinfoFields := schema.GroupInfo{}.Fields()
	_ = groupinfoFields
	// groupinfoDescMaxUsers is the schema descriptor for max_users field.
	groupinfoDescMaxUsers := groupinfoFields[1].Descriptor()
	// groupinfo.DefaultMaxUsers holds the default value on creation for the max_users field.
	groupinfo.DefaultMaxUsers = groupinfoDescMaxUsers.Default.(int)
	userMixin := schema.User{}.Mixin()
	userMixinFields := [...][]ent.Field{
		userMixin[0].Fields(),
	}
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescOptionalInt is the schema descriptor for optional_int field.
	userDescOptionalInt := userMixinFields[0][0].Descriptor()
	// user.OptionalIntValidator is a validator for the "optional_int" field. It is called by the builders before save.
	user.OptionalIntValidator = userDescOptionalInt.Validators[0].(func(int) error)
	// userDescLast is the schema descriptor for last field.
	userDescLast := userFields[2].Descriptor()
	// user.DefaultLast holds the default value on creation for the last field.
	user.DefaultLast = userDescLast.Default.(string)
}
