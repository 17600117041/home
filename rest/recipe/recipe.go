// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

// Package recipe provides functionality for managing recipes.
package recipe

import (
	"time"
)

// Recipe is a list of ingredients and directions.
type Recipe struct {
	// A URL safe version of the datastores key for this recipe item.
	Key string `datastore:",noindex"`

	// The title of the recipe.
	Name string

	// The URL where the recipe was originally pulled from.
	URL string `datastore:",noindex"`

	// This is the time the recipe was last modified.
	LastModified time.Time

	// The list of ingredients. 
	Ingredients []string `datastore:",noindex"`

	// The list of directions. 
	Directions []string `datastore:",noindex"`
}
