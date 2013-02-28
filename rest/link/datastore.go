// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package link

import (
	"appengine"
	"appengine/datastore"
	"github.com/icub3d/gorca"
	"net/http"
)

// NewLinkHelper is a helper function that creates a new link in the
// datastore for the given link. The given link is updated with new
// keys. If a failure occured, false is returned and a response was
// returned to the request. This case should be terminal.
func NewLinkHelper(c appengine.Context, w http.ResponseWriter,
	r *http.Request, l *Link) bool {

	// Blank out the link key.
	l.Key = ""

	// Save the new link.
	if !PutLinkHelper(c, w, r, l) {
		return false
	}

	return true
}

// GetLinkHelper is a helper function that retrieves a link from the
// datastore. If a failure occured, false is returned and a response
// was returned to the request. This case should be terminal.
func GetLinkHelper(c appengine.Context, w http.ResponseWriter,
	r *http.Request, key string) (*Link, bool) {

	// Decode the string version of the key.
	k, err := datastore.DecodeKey(key)
	if err != nil {
		gorca.LogAndUnexpected(c, w, r, err)
		return nil, false
	}

	// Get the link by key.
	var l Link
	if err := datastore.Get(c, k, &l); err != nil {
		gorca.LogAndNotFound(c, w, r, err)
		return nil, false
	}

	return &l, true
}

// PutLinkHelepr saves the link. If the link or any of it's items
// don't have a key, a key will be made for it.
func PutLinkHelper(c appengine.Context, w http.ResponseWriter,
	r *http.Request, l *Link) bool {

	// This is the link of keys we are going to PutMulti.
	keys := make([]string, 0, 1)

	// This is the link of things we are going to put.
	values := make([]interface{}, 0, 1)

	// TODO The words in the name should be part in the tags.

	// We need the key to generate new keys and we might not even have
	// one.
	var skey string
	var ok bool
	if l.Key == "" {
		// Make a key and save it's value to the link.
		skey, _, ok = gorca.NewKey(c, w, r, "Link", nil)
		if !ok {
			return false
		}

		l.Key = skey
	}

	// Add the link itself.
	keys = append(keys, l.Key)
	values = append(values, l)

	// Save them all.
	return gorca.PutStringKeys(c, w, r, keys, values)
}
