// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package parsers

import (
	"regexp"
)

// AllRecipesDotComParser is a RecipeParser for the site
// allrecipes.com.
type AllRecipesDotComParser struct{}

// GetName gets the title of the recipe.
func (a AllRecipesDotComParser) GetName(data []byte) string {
	return getFirstH1(data)
}

// GetDirections gets the directions of the recipe.
func (a AllRecipesDotComParser) GetDirections(data []byte) []string {

	// It looks like the directions are the only ordered list, so just
	// get the first.
	re := regexp.MustCompile(`(?sU)<ol>.*</ol>`)
	ol := re.Find(data)

	// Get each list item, removing HTML cruft.
	re = regexp.MustCompile(`(?sU)<li><span class="plaincharacterwrap break">(.*)</span></li>`)
	directions := re.FindAllSubmatch(ol, -1)

	list := []string{}
	for _, dir := range directions {
		// If there is a submatch, append it. Otherwise skip it.
		if len(dir) > 1 {
			list = append(list, cleanField(string(dir[1])))
		}
	}

	return list
}

// GetIngredients gets the ingredients of the recipe.
func (a AllRecipesDotComParser) GetIngredients(data []byte) []string {

	// We can find the unordered list by classname.
	re := regexp.MustCompile(`(?sU)<ul class="ingredient-wrap.*</ul>`)
	uls := re.FindAll(data, -1)

	list := []string{}
	for _, ul := range uls {

		// Each item in the list has a <p> tag with the ingredients.
		re = regexp.MustCompile(`(?sU)<p class="fl-ing" itemprop="ingredients">.*</p>`)
		ingredients := re.FindAll(ul, -1)

		// These are the regular expressions that parse out the amount and
		// name from the <p> ingredients.
		reamt := regexp.MustCompile(`(?sU)<span id="lblIngAmount" class="ingredient-amount">(.*)</span>`)
		rename := regexp.MustCompile(`(?sU)<span id="lblIngName" class="ingredient-name">(.*)</span>`)

		for _, ingredient := range ingredients {
			amount := reamt.FindSubmatch(ingredient)
			name := rename.FindSubmatch(ingredient)

			// This will eventually be the entire item as a single string.
			item := ""

			// Only add the amount if the submatch was found.
			if len(amount) > 1 {
				item = string(amount[1])
			}

			// Only add the name if the submatch was found.
			if len(name) > 1 {
				// Add a space if we got an amount.
				if item != "" {
					item += " "
				}

				// Append the name.
				item += string(name[1])
			}

			// Add the complete item.
			if item != "" {
				list = append(list, cleanField(item))
			}
		}
	}
	return list
}
