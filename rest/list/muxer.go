package list

import (
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
)

// MakeMuxer creates a http.Handler to manage all list operations. If
// a prefix is given, that prefix will be the first part of the
// url. This muxer will provide the following handlers and return a
// RESTful 404 on all others.
//
//  GET     prefix + /        Get all lists.
//  POST    prefix + /        Create a new list.
//
//  GET     prefix + /{key}/  Get the list for the given key.
//  PUT     prefix + /{key}/  Update the list with the given key.
//  DELETE  prefix + /{key}/  Delete the list with the given key.
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

	// Get all lists.
	m.HandleFunc("/", GetAllLists).Methods("GET")

	// Make a new list.
	m.HandleFunc("/", PostList).Methods("POST")

	// Singe list operations.
	m.HandleFunc("/{key}/", GetList).Methods("GET")
	m.HandleFunc("/{key}/", PutList).Methods("PUT")
	m.HandleFunc("/{key}/", DeleteList).Methods("DELETE")

	// Everything else fails.
	m.HandleFunc("/{path:.*}", gorca.NotFoundFunc)

	return m
}
