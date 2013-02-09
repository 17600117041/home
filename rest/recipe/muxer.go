// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package recipe

import (
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
)

// MakeMuxer creates a http.Handler to manage all recipe operations. If
// a prefix is given, that prefix will be the first part of the
// url. This muxer will provide the following handlers and return a
// RESTful 404 on all others.
//
//  GET     prefix + /        Get all recipes.
//  POST    prefix + /        Create a new recipe.
//
//  GET     prefix + /{key}/  Get the recipe for the given key.
//  PUT     prefix + /{key}/  Update the recipe with the given key.
//  DELETE  prefix + /{key}/  Delete the recipe with the given key.
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

	// Get all recipes.
	m.HandleFunc("/", GetAllRecipes).Methods("GET")

	// Make a new recipe.
	m.HandleFunc("/", PostRecipe).Methods("POST")

	// Singe recipe operations.
	m.HandleFunc("/{key}/", GetRecipe).Methods("GET")
	m.HandleFunc("/{key}/", PutRecipe).Methods("PUT")
	m.HandleFunc("/{key}/", DeleteRecipe).Methods("DELETE")

	// Everything else fails.
	m.HandleFunc("/{path:.*}", gorca.NotFoundFunc)

	return m
}
