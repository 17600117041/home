// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package recipe

import (
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
	"time"
)

// PostRecipe creates a new recipe from the POSTed data. If a URL was
// given and the ingredients and diretions are blank, the information
// will be pulled from a URL if possible.
func PostRecipe(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	datastore.RunInTransaction(c, func(c appengine.Context) error {

		// Get the recipe from the body.
		var recipe Recipe
		if !gorca.UnmarshalFromBodyOrFail(c, w, r, &recipe) {
			return fmt.Errorf("unmarshalling")
		}

		// Create a new recipe in the datastore.
		if !NewRecipeHelper(c, w, r, &recipe) {
			return fmt.Errorf("generating new recipe")
		}

		// Return the updated recipe back.
		gorca.WriteJSON(c, w, r, recipe)

		return nil
	}, nil)

}

// PutRecipe saves the recipe for the given tag. 
func PutRecipe(w http.ResponseWriter, r *http.Request) {
	// Get the context.
	c := appengine.NewContext(r)

	// Get the Key.
	vars := mux.Vars(r)
	key := vars["key"]

	datastore.RunInTransaction(c, func(c appengine.Context) error {

		// Get the original recipe.
		or, ok := GetRecipeHelper(c, w, r, key)
		if !ok {
			return fmt.Errorf("getting original recipe")
		}

		// Get the new recipe from the body.
		var nr Recipe
		if !gorca.UnmarshalFromBodyOrFail(c, w, r, &nr) {
			return fmt.Errorf("unmarshalling")
		}

		// Merge the new ingredients into the old recipe and remove
		// deleted keys.
		delskeys := or.Ingredients.Merge(&(nr.Ingredients))
		if !gorca.DeleteStringKeys(c, w, r, delskeys) {
			return fmt.Errorf("deleting ingredient keys")
		}

		// Merge the new directions into the old recipe and remove deleted keys.
		delskeys = or.Diretions.Merge(&(nr.Directions))
		if !gorca.DeleteStringKeys(c, w, r, delskeys) {
			return fmt.Errorf("deleting direction keys")
		}

		// Update the values.
		or.LastModified = time.Now()
		or.Title = nr.Name

		if !PutRecipeHelper(c, w, r, or) {
			return fmt.Errorf("putting recipe")
		}

		// Remove it from memcache
		memcache.Set(c, &memcache.Item{
			Key:   key,
			Value: []byte(fmt.Sprintf("%d", or.LastModified.Unix())),
		})

		// Return the updated recipe back.
		gorca.WriteJSON(c, w, r, or)

		return nil
	}, nil)

}

// DeleteRecipe deletes the recipe for the given tag. The currently
// logged in user must own the recipe. Otherwise, an unauthorized error
// is returned.
func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	// Get the context.
	c := appengine.NewContext(r)

	// Get the Key.
	vars := mux.Vars(r)
	key := vars["key"]

	datastore.RunInTransaction(c, func(c appengine.Context) error {

		if !gorca.DeleteStringKeyAndAllAncestors(c, w, r,
			[]string{"List", "Items"}, key) {
			return fmt.Errorf("deleting recipe and items")
		}

		// Remove it from memcache
		memcache.Delete(c, key)

		gorca.WriteSuccessMessage(c, w, r)

		return nil

	}, nil)
}
