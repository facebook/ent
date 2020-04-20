// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package enttest

import (
	"context"

	"github.com/facebookincubator/ent/entc/integration/migrate/entv2"
	_ "github.com/facebookincubator/ent/entc/integration/migrate/entv2/runtime"

	"github.com/facebookincubator/ent/dialect/sql/schema"
)

// TestingT is the interface that is shared between
// testing.T and testing.B and used by enttest.
type TestingT interface {
	FailNow()
	Error(...interface{})
}

// Open calls entv2.Open and auto-run migration.
func Open(t TestingT, driverName, dataSourceName string, options ...schema.MigrateOption) *entv2.Client {
	c, err := entv2.Open(driverName, dataSourceName)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := c.Schema.Create(context.Background(), options...); err != nil {
		t.Error(err)
		t.FailNow()
	}
	return c
}
