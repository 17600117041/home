// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

// Package recipe provides functionality for managing recipes.
package recipe

import (
	"github.com/icub3d/list/rest/list"
	"time"
)

// Recipe is a list of ingredients and directions.
type Recipe struct {
	// A URL safe version of the datastores key for this recipe item.
	Key string `datastore:",noindex"`

	// The title of the recipe.
	Title string

	// The URL where the recipe was originally pulled from.
	URL string `datastore:",noindex"`

	// This is the time the recipe was last modified.
	LastModified time.Time

	// A URL safe version of the datastores key for the list of
	// ingredients.
	IngredientsKey string `datastore:",noindex"`

	// A URL safe version of the datastores key for the list of
	// directions.
	DirectionsKey string `datastore:",noindex"`

	// The list of ingredients. This will be managed internally.
	Ingredients list.List `datastore:"-"`

	// The list of directions. This will be managed internally.
	Directions list.List `datastore:"-"`
}
