// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package link

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
)

// PostLink creates a new link from the POSTed data.
func PostLink(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	datastore.RunInTransaction(c, func(c appengine.Context) error {

		// Get the link from the body.
		var l Link
		if !gorca.UnmarshalFromBodyOrFail(c, w, r, &l) {
			return fmt.Errorf("unmarshalling")
		}

		// Create a new link in the datastore.
		if !NewLinkHelper(c, w, r, &l) {
			return fmt.Errorf("generating new link")
		}

		// Return the updated link back.
		gorca.WriteJSON(c, w, r, l)

		return nil
	}, nil)

}

// PutLink saves the link for the given tag.
func PutLink(w http.ResponseWriter, r *http.Request) {
	// Get the context.
	c := appengine.NewContext(r)

	// Get the Key.
	vars := mux.Vars(r)
	key := vars["key"]

	datastore.RunInTransaction(c, func(c appengine.Context) error {

		// Get the original link.
		ol, ok := GetLinkHelper(c, w, r, key)
		if !ok {
			return fmt.Errorf("getting original link")
		}

		// Get the new link from the body.
		var nl Link
		if !gorca.UnmarshalFromBodyOrFail(c, w, r, &nl) {
			return fmt.Errorf("unmarshalling")
		}

		// Update the values.
		ol.Name = nl.Name
		ol.Icon = nl.Icon
		ol.Url = nl.Url
		ol.Description = nl.Description
		ol.Tags = nl.Tags

		if !PutLinkHelper(c, w, r, ol) {
			return fmt.Errorf("putting link")
		}

		// Return the updated link back.
		gorca.WriteJSON(c, w, r, ol)

		return nil
	}, nil)

}

// DeleteLink deletes the link for the given tag. The currently
// logged in user must own the link. Otherwise, an unauthorized error
// is returned.
func DeleteLink(w http.ResponseWriter, r *http.Request) {
	// Get the context.
	c := appengine.NewContext(r)

	// Get the Key.
	vars := mux.Vars(r)
	key := vars["key"]

	datastore.RunInTransaction(c, func(c appengine.Context) error {

		if !gorca.DeleteStringKeyAndAncestors(c, w, r, "Link", key) {
			return fmt.Errorf("deleting link and items")
		}

		gorca.WriteSuccessMessage(c, w, r)

		return nil

	}, nil)
}
