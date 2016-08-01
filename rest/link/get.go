// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package link

import (
	"appengine"
	"appengine/datastore"
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
)

// GetAllLinks fetches all of the links.
func GetAllLinks(w http.ResponseWriter, r *http.Request) {
	// Create the query.
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Link")

	search := r.FormValue("search")
	if search != "" {
		q = q.Filter("Tags =", search)
	}

	q = q.Order("Name")

	// Fetch the links. 
	links := []Link{}
	if _, err := q.GetAll(c, &links); err != nil {
		gorca.LogAndUnexpected(c, w, r, err)
		return
	}

	// Write the links as JSON.
	gorca.WriteJSON(c, w, r, links)
}

// GetLink fetches the link for the given tag.
func GetLink(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// Get the Key.
	vars := mux.Vars(r)
	key := vars["key"]

	l, ok := GetLinkHelper(c, w, r, key)
	if !ok {
		return
	}

	// Write the links as JSON.
	gorca.WriteJSON(c, w, r, l)
}
