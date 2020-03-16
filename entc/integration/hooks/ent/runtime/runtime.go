// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/facebookincubator/ent/entc/integration/hooks/ent/card"
	"github.com/facebookincubator/ent/entc/integration/hooks/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	cardHooks := schema.Card{}.Hooks()
	for i, h := range cardHooks {
		card.Hooks[i] = h
	}
	cardFields := schema.Card{}.Fields()
	_ = cardFields
	// cardDescNumber is the schema descriptor for number field.
	cardDescNumber := cardFields[0].Descriptor()
	// card.DefaultNumber holds the default value on creation for the number field.
	card.DefaultNumber = cardDescNumber.Default.(string)
	// card.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	card.NumberValidator = cardDescNumber.Validators[0].(func(string) error)
	// cardDescCreatedAt is the schema descriptor for created_at field.
	cardDescCreatedAt := cardFields[2].Descriptor()
	// card.DefaultCreatedAt holds the default value on creation for the created_at field.
	card.DefaultCreatedAt = cardDescCreatedAt.Default.(func() time.Time)
}

const (
	Version = "(devel)" // Version of ent codegen.
)
