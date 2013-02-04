// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package list

import (
	"appengine"
	"appengine/datastore"
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
)

// GetAllLists fetches all of the lists.
func GetAllLists(w http.ResponseWriter, r *http.Request) {
	// Create the query.
	c := appengine.NewContext(r)
	q := datastore.NewQuery("List").Order("-LastModified")

	// Fetch the lists. 
	lists := []List{}
	if _, err := q.GetAll(c, &lists); err != nil {
		gorca.LogAndUnexpected(c, w, r, err)
		return
	}

	// Write the lists as JSON.
	gorca.WriteJSON(c, w, r, lists)
}

// GetList fetches the list for the given tag.
func GetList(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// Get the Key.
	vars := mux.Vars(r)
	key := vars["key"]

	l, ok := GetListHelper(c, w, r, key)
	if !ok {
		return
	}

	// Write the lists as JSON.
	gorca.WriteJSON(c, w, r, l)
}
