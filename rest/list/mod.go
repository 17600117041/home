// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package list

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
	"time"
)

// PostList creates a new list
func PostList(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	datastore.RunInTransaction(c, func(c appengine.Context) error {

		// Get the list from the body.
		var l List
		if !gorca.UnmarshalFromBodyOrFail(w, r, &l) {
			return fmt.Errorf("unmarshalling")
		}

		// Generate a new key for this list.
		id, _, err := datastore.AllocateIDs(c, "List", nil, 1)
		if err != nil {
			gorca.LogAndUnexpected(w, r, err)
			return fmt.Errorf("allocating list id")
		}
		key := datastore.NewKey(c, "List", "", id, nil)

		// Save its string value to the list.
		l.Key = key.Encode()
		l.LastModified = time.Now()

		// Save the list.
		if _, err := datastore.Put(c, key, &l); err != nil {
			gorca.LogAndUnexpected(w, r, err)
			return fmt.Errorf("saving list")

		}

		// Generate a new key set for each of the list items.
		low, high, err := datastore.AllocateIDs(c, "Item", key, len(l.Items))
		if err != nil || high-low != int64(len(l.Items)) {
			gorca.LogAndUnexpected(w, r, err)
			return fmt.Errorf("allocating item ids")
		}

		// Add each of the items in the list.
		for i, item := range l.Items {
			// Make a key.
			ikey := datastore.NewKey(c, "Item", "", low+int64(i), key)

			// Save the keys string value.
			item.Key = ikey.Encode()
			item.Order = i

			// Save the item.
			if _, err := datastore.Put(c, ikey, item); err != nil {
				gorca.LogAndUnexpected(w, r, err)
				return fmt.Errorf("saving item")
			}
		}

		// Return the updated list back.
		gorca.WriteJSON(w, r, l)

		return nil
	}, nil)

}

// PutList saves the list for the given tag. The currently logged in
// user must own the list or the list must have been shared with the
// user. Otherwise, an unauthorized error is returned. Additionally,
// if the user is not the owner, only the items in the list can be
// modified.
func PutList(w http.ResponseWriter, r *http.Request) {
	// Get the context.
	c := appengine.NewContext(r)

	// Get the Key.
	vars := mux.Vars(r)
	key := vars["key"]

	datastore.RunInTransaction(c, func(c appengine.Context) error {

		ol, olkey, ok := getListHelper(w, r, key)
		if !ok {
			return fmt.Errorf("getting original list")
		}

		// Get the new list from the body.
		var nl List
		if !gorca.UnmarshalFromBodyOrFail(w, r, &nl) {
			return fmt.Errorf("unmarshalling")
		}

		// Merge the new list into the old list.
		delskeys := ol.Merge(&nl)

		// Delete the removed keys.
		delkeys := make([]*datastore.Key, 0, len(delskeys))
		for _, k := range delskeys {
			key, err := datastore.DecodeKey(k)
			if err != nil {
				gorca.LogAndUnexpected(w, r, err)
				return fmt.Errorf("decoding item key")
			}

			delkeys = append(delkeys, key)
		}

		// Delete all the removed items.
		if err := datastore.DeleteMulti(c, delkeys); err != nil {
			gorca.LogAndUnexpected(w, r, err)
			return fmt.Errorf("deleting items")
		}

		// Set the last modified time.
		ol.LastModified = time.Now()

		// Save the changed list.
		if _, err := datastore.Put(c, olkey, ol); err != nil {
			gorca.LogAndUnexpected(w, r, err)
			return fmt.Errorf("saving list")
		}

		// Save the list items.
		for _, item := range ol.Items {
			var ikey *datastore.Key
			var err error

			if item.Key != "" {
				ikey, err = datastore.DecodeKey(item.Key)
				if err != nil {
					gorca.LogAndUnexpected(w, r, err)
					return fmt.Errorf("decoding item key")
				}
			} else {
				// Generate a new key for this item.
				low, _, err := datastore.AllocateIDs(c, "Item", olkey, 1)
				if err != nil {
					gorca.LogAndUnexpected(w, r, err)
					return fmt.Errorf("generating item key")
				}

				ikey = datastore.NewKey(c, "Item", "", low, olkey)
				item.Key = ikey.Encode()
			}

			// Save the list item.
			if _, err := datastore.Put(c, ikey, item); err != nil {
				gorca.LogAndUnexpected(w, r, err)
				return fmt.Errorf("saving item key")
			}
		}

		// Return the updated list back.
		gorca.WriteJSON(w, r, ol)

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

		// Decode the string version of the key.
		k, err := datastore.DecodeKey(key)
		if err != nil {
			gorca.LogAndUnexpected(w, r, err)
			return fmt.Errorf("decoding key")

		}

		// Get all of the items for this list and delete them.
		q := datastore.NewQuery("Item").Ancestor(k).KeysOnly()
		keys, err := q.GetAll(c, nil)
		if err != nil {
			gorca.LogAndUnexpected(w, r, err)
			return fmt.Errorf("getting items")
		}

		// Delete all the items.
		if err := datastore.DeleteMulti(c, keys); err != nil {
			gorca.LogAndUnexpected(w, r, err)
			return fmt.Errorf("deleting items")
		}

		// Delete the key.
		if err := datastore.Delete(c, k); err != nil {
			gorca.LogAndUnexpected(w, r, err)
			return fmt.Errorf("deleting list")
		}

		gorca.WriteSuccessMessage(w, r)
		return nil

	}, nil)

}
