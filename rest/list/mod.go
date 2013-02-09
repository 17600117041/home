// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package list

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

// PostList creates a new list from the POSTed data.
func PostList(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	datastore.RunInTransaction(c, func(c appengine.Context) error {

		// Get the list from the body.
		var l List
		if !gorca.UnmarshalFromBodyOrFail(c, w, r, &l) {
			return fmt.Errorf("unmarshalling")
		}

		// Create a new list in the datastore.
		if !NewListHelper(c, w, r, &l) {
			return fmt.Errorf("generating new list")
		}

		// Return the updated list back.
		gorca.WriteJSON(c, w, r, l)

		return nil
	}, nil)

}

// PutList saves the list for the given tag.
func PutList(w http.ResponseWriter, r *http.Request) {
	// Get the context.
	c := appengine.NewContext(r)

	// Get the Key.
	vars := mux.Vars(r)
	key := vars["key"]

	datastore.RunInTransaction(c, func(c appengine.Context) error {

		// Get the original list.
		ol, ok := GetListHelper(c, w, r, key)
		if !ok {
			return fmt.Errorf("getting original list")
		}

		// Get the new list from the body.
		var nl List
		if !gorca.UnmarshalFromBodyOrFail(c, w, r, &nl) {
			return fmt.Errorf("unmarshalling")
		}

		// Merge the new list into the old list and remove deleted keys.
		delskeys := ol.Merge(&nl)
		if !gorca.DeleteStringKeys(c, w, r, delskeys) {
			return fmt.Errorf("deleting keys")
		}

		// Update the values.
		ol.LastModified = time.Now()
		ol.Name = nl.Name

		if !PutListHelper(c, w, r, ol) {
			return fmt.Errorf("putting list")
		}

		// Remove it from memcache
		memcache.Set(c, &memcache.Item{
			Key:   key,
			Value: []byte(fmt.Sprintf("%d", ol.LastModified.Unix())),
		})

		// Return the updated list back.
		gorca.WriteJSON(c, w, r, ol)

		return nil
	}, nil)

}

// DeleteList deletes the list for the given tag. The currently
// logged in user must own the list. Otherwise, an unauthorized error
// is returned.
func DeleteList(w http.ResponseWriter, r *http.Request) {
	// Get the context.
	c := appengine.NewContext(r)

	// Get the Key.
	vars := mux.Vars(r)
	key := vars["key"]

	datastore.RunInTransaction(c, func(c appengine.Context) error {

		if !gorca.DeleteStringKeyAndAncestors(c, w, r, "Item", key) {
			return fmt.Errorf("deleting list and items")
		}

		// Remove it from memcache
		memcache.Delete(c, key)

		gorca.WriteSuccessMessage(c, w, r)

		return nil

	}, nil)
}
