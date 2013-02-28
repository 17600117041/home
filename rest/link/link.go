// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

// Package link provides functionality for managing links.
package link

// Link is the structure used to save, get, and delete links from the
// datastore.
type Link struct {
	// The name of the link.
	Name string

	// A URL safe version of the datastores key for this link. It
	// is not stored in the datastore.
	Key string `datastore:",noindex"`

	// The Url is the location of the link.
	Url string `datastore:",noindex"`

	// The Icon is the bootstrap icon to display. See:
	// http://twitter.github.com/bootstrap/base-css.html#icons
	Icon string `datastore:",noindex"`

	// The Description is the description of the link
	Description string `datastore:",noindex"`

	// Tags are the searchable elements of a link.
	Tags []string
}
