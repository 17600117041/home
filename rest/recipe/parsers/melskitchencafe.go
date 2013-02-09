// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package parsers

import (
	"regexp"
	"strings"
)

// MelsKitchenCafeDotComParser is a RecipeParser for the site
// melskitchencafe.com.
type MelsKitchenCafeDotComParser struct{}

// GetTitle gets the title of the recipe.
func (a MelsKitchenCafeDotComParser) GetTitle(data []byte) string {
	return getFirstH1(data)
}

// GetDirections gets the directions of the recipe.
func (a MelsKitchenCafeDotComParser) GetDirections(data []byte) []string {

	// It looks like we start with DIRECTIONS and end with the Recipe
	// Source.
	re := regexp.MustCompile(`(?sU)<p>DIRECTIONS:<br />(.*)</p>[\n ]*<p><em><strong>Recipe Source`)

	// Get the first submatch.
	directions := re.FindSubmatch(data)
	if len(directions) < 2 {
		return nil
	}
	dirs := directions[1]

	// Get rid of all the HTML cruft.
	re = regexp.MustCompile(`<[^<]*>`)
	dirs = re.ReplaceAll(dirs, []byte(""))

	// Make a list of directions removing empty spaces.
	list := []string{}
	for _, dir := range strings.Split(string(dirs), "\n") {
		if dir == "" {
			continue
		}

		list = append(list, cleanField(dir))
	}

	return list
}

// GetIngredients gets the ingredients of the recipe.
func (a MelsKitchenCafeDotComParser) GetIngredients(data []byte) []string {

	// It looks like it starts with INGREDIENTS and ends just before
	// DIRECTIONS.
	re := regexp.MustCompile(`(?sU)<p>INGREDIENTS:<br />(.*)</p>[\n ]*<p>DIRECTIONS`)

	// Pull out the ingredients or bail.
	ingredients := re.FindSubmatch(data)
	if len(ingredients) < 2 {
		return nil
	}
	ings := ingredients[1]

	// Get rid of all the line breaks.
	re = regexp.MustCompile(`<br />`)
	ings = re.ReplaceAll(ings, []byte(""))

	// Get rid of any other tags, and lines that are all tags.
	re = regexp.MustCompile(`<.*>`)
	ings = re.ReplaceAll(ings, []byte(""))

	// Make a list of ingredients skipping empty lines.
	list := []string{}
	for _, ing := range strings.Split(string(ings), "\n") {
		if ing == "" {
			continue
		}

		list = append(list, cleanField(ing))
	}

	return list
}
