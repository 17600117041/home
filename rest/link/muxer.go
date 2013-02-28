// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package link

import (
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
)

// MakeMuxer creates a http.Handler to manage all link operations. If
// a prefix is given, that prefix will be the first part of the
// url. This muxer will provide the following handlers and return a
// RESTful 404 on all others.
//
//  GET     prefix + /        Get all links.
//  POST    prefix + /        Create a new link.
//
//  GET     prefix + /{key}/  Get the link for the given key.
//  PUT     prefix + /{key}/  Update the link with the given key.
//  DELETE  prefix + /{key}/  Delete the link with the given key.
//
// See the functions related to these urls for more details on what
// each one does and the requirements behind it.
func MakeMuxer(prefix string) http.Handler {
	var m *mux.Router

	// Pass through the prefix if we got one.
	if prefix == "" {
		m = mux.NewRouter()
	} else {
		m = mux.NewRouter().PathPrefix(prefix).Subrouter()
	}

	// Get all links.
	m.HandleFunc("/", GetAllLinks).Methods("GET")

	// Make a new link.
	m.HandleFunc("/", PostLink).Methods("POST")

	// Singe link operations.
	m.HandleFunc("/{key}/", GetLink).Methods("GET")
	m.HandleFunc("/{key}/", PutLink).Methods("PUT")
	m.HandleFunc("/{key}/", DeleteLink).Methods("DELETE")

	// Everything else fails.
	m.HandleFunc("/{path:.*}", gorca.NotFoundFunc)

	return m
}
