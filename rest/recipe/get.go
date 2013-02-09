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
	"strconv"
	"time"
)

// GetAllRecipes fetches all of the recipes.
func GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	// Create the query.
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Recipe").Order("-LastModified")

	// Fetch the recipes. 
	recipes := []Recipe{}
	if _, err := q.GetAll(c, &recipes); err != nil {
		gorca.LogAndUnexpected(c, w, r, err)
		return
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

	// Handle Not-modified
	ut := r.FormValue("date")
	if ut != "" {
		notmod(c, w, r, key, ut)
		return
	}

	recipe, ok := GetRecipeHelper(c, w, r, key)
	if !ok {
		return
	}

	// Save the results to memcache.
	item := &memcache.Item{
		Key:   key,
		Value: []byte(fmt.Sprintf("%d", recipe.LastModified.Unix())),
	}
	if err := memcache.Set(c, item); err != nil {
		gorca.Log(c, r, "error", "failed to set memcache: %v", err)
	}

	// Write the recipes as JSON.
	gorca.WriteJSON(c, w, r, recipe)
}

// notmod checks to see if the cached date for the key is newer than
// the date given from the url. This call is terminal. It will always
// respond to the request. If the dates are equal, then this function
// sends a 304 Not Modified. If an error occurs, the error is logged
// and sent back.
func notmod(c appengine.Context, w http.ResponseWriter,
	r *http.Request, key string, date string) {

	// Convert the given string.
	i, err := strconv.ParseInt(date, 10, 64)
	if err != nil {
		gorca.LogAndFailed(c, w, r, err)
		return
	}
	t := time.Unix(i, 0)

	// Try to get the key from memcache
	item, err := memcache.Get(c, key)
	if err != nil && err != memcache.ErrCacheMiss {
		gorca.LogAndFailed(c, w, r, err)
		return
	}

	var mt time.Time

	// Check to see if it's simply not there.
	if err == memcache.ErrCacheMiss {
		gorca.Log(c, r, "info", "failed to get memcache: %s", key)

		// Try to get the recipe.
		recipe, ok := GetRecipeHelper(c, w, r, key)
		if !ok {
			return
		}

		// Save the results to memcache.
		item := &memcache.Item{
			Key:   key,
			Value: []byte(fmt.Sprintf("%d", recipe.LastModified.Unix())),
		}
		if err := memcache.Set(c, item); err != nil {
			gorca.Log(c, r, "error", "failed to set memcache: %v", err)
		}

		mt = recipe.LastModified
	} else {
		// Convert the memcache string.
		mi, err := strconv.ParseInt(string(item.Value), 10, 64)
		if err != nil {
			gorca.LogAndFailed(c, w, r, err)
			return
		}

		mt = time.Unix(mi, 0)

	}

	if mt.Equal(t) {
		// Write out the not modified.
		w.WriteHeader(http.StatusNotModified)
		return
	}

	// Write out that we modified.
	gorca.WriteMessage(c, w, r, "success", "Modified.",
		http.StatusOK)
	return
}
