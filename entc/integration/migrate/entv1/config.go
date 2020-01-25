// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv1

import (
	"context"

	"github.com/facebookincubator/ent/dialect"
)

// Option function to configure the client.
type Option func(*config)

// Config is the configuration for the client and its builder.
type config struct {
	// driver used for executing database requests.
	driver dialect.Driver
	// debug enable a debug logging.
	debug bool
	// log used for logging on debug mode.
	log func(...interface{})
	// debugWithContext enable a debug logging with context.
	debugWithContext bool
	// logCtx used for logging with context on debug mode.
	logCtx func(context.Context, ...interface{})
}

// Options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
	if c.debugWithContext {
		c.driver = dialect.DebugWithContext(c.driver, c.logCtx)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// DebugWithContext enables debug logging with context on the ent.Driver.
func DebugWithContext() Option {
	return func(c *config) {
		c.debugWithContext = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...interface{})) Option {
	return func(c *config) {
		c.log = fn
	}
}

// LogCtx sets the contextual logging function for debug mode.
func LogCtx(fn func(context.Context, ...interface{})) Option {
	return func(c *config) {
		c.logCtx = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}
