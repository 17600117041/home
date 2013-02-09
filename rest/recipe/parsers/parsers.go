// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

// Package parsers provides an interface for parsing html pages (or
// other data) for recipes.
package parsers

import (
	"fmt"
	"net/url"
	"strings"
)

// RecipeParser is an interface that cleans information from data
// about a recipe.
type RecipeParser interface {
	// GetName gets the title of the recipe.
	GetName(data []byte) string

	// GetIngredients returns a list of ingredients for the recipe.
	GetIngredients(data []byte) []string

	// GetDirections returns a list of directions for the recipe.
	GetDirections(data []byte) []string
}

// GetParserForURL attempts to match the given URL with a known recipe
// parser. 
func GetParserForURL(URL string) (RecipeParser, error) {
	u, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}

	switch {
	case strings.Contains(u.Host, "allrecipes.com"):
		return AllRecipesDotComParser{}, nil
	case strings.Contains(u.Host, "melskitchencafe.com"):
		return MelsKitchenCafeDotComParser{}, nil
	}

	return nil, fmt.Errorf("no parser found: %s", URL)
}
