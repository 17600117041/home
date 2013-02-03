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
		gorca.LogAndUnexpected(w, r, err)
		return
	}

	// Write the lists as JSON.
	gorca.WriteJSON(w, r, lists)
}

// GetList fetches the list for the given tag. The currently
// logged in user must own the list or the list must have been shared
// with the user. Otherwise, an unauthorized error is returned.
func GetList(w http.ResponseWriter, r *http.Request) {
	// Get the Key.
	vars := mux.Vars(r)
	key := vars["key"]

	l, _, ok := getListHelper(w, r, key)
	if !ok {
		return
	}

	// Write the lists as JSON.
	gorca.WriteJSON(w, r, l)
}

// getListHelper is a helper function that retrieves a list and it's
// items from the datastore. If a failure occured, false is returned
// and a response was returned to the request. This case should be
// terminal.
func getListHelper(w http.ResponseWriter, r *http.Request, key string) (*List, *datastore.Key, bool) {
	// Get the context.
	c := appengine.NewContext(r)

	// Decode the string version of the key.
	k, err := datastore.DecodeKey(key)
	if err != nil {
		gorca.LogAndUnexpected(w, r, err)
		return nil, nil, false
	}

	// Get the list by key.
	var l List
	if err := datastore.Get(c, k, &l); err != nil {
		gorca.LogAndNotFound(w, r, err)
		return nil, nil, false
	}

	// Get all of the items for the list.
	var li ItemsList
	q := datastore.NewQuery("Item").Ancestor(k).Order("Order")
	if _, err := q.GetAll(c, &li); err != nil {
		gorca.LogAndUnexpected(w, r, err)
		return nil, nil, false
	}

	l.Items = li

	return &l, k, true
}
