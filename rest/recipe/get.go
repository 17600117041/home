// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package recipe

import (
	"appengine"
	"appengine/datastore"
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
)

// GetAllRecipes fetches all of the recipes.
func GetAllRecipes(w http.ResponseWriter, r *http.Request) {

	// Create the query.
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Recipe").Order("Name")

	// Fetch the recipes. 
	recipes := []Recipe{}
	if _, err := q.GetAll(c, &recipes); err != nil {
		gorca.LogAndUnexpected(c, w, r, err)
		return
	}

	// We don't need the ingredients or directions
	for _, recipe := range recipes {
		recipe.Ingredients = []string{}
		recipe.Directions = []string{}
	}

	// Write the recipes as JSON.
	gorca.WriteJSON(c, w, r, recipes)
}

// GetRecipe fetches the recipe for the given tag.
func GetRecipe(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// Get the Key.
	vars := mux.Vars(r)
	key := vars["key"]

	recipe, ok := GetRecipeHelper(c, w, r, key)
	if !ok {
		return
	}

	// Write the recipes as JSON.
	gorca.WriteJSON(c, w, r, recipe)
}
