// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

// Package rest provides a RESTful interface for interacting with the
// list app.
package rest

import (
	"github.com/icub3d/gorca"
	"github.com/icub3d/list/rest/list"
	"github.com/icub3d/list/rest/user"
	"net/http"
)

func init() {
	// Manage the user 
	http.Handle("/rest/user/", user.MakeMuxer("/rest/user/"))
	http.Handle("/rest/list/", list.MakeMuxer("/rest/list/"))

	http.HandleFunc("/", gorca.NotFoundFunc)

}
