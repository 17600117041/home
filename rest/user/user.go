// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

// Package user provides functionality for getting information about
// the currently logged in user.
package user

import (
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
)

// MakeMuxer creates a http.Handler to manage all user operations. If
// a prefix is given, that prefix will be the first part of the
// url. This muxer will provide the following handlers and return a
// RESTful 404 on all others.
//
//  GET     prefix + /      Get the currently logged in user.
func MakeMuxer(prefix string) http.Handler {
	var m *mux.Router

	// Pass through the prefix if we got one.
	if prefix == "" {
		m = mux.NewRouter()
	} else {
		m = mux.NewRouter().PathPrefix(prefix).Subrouter()
	}

	m.HandleFunc("/", GetUser).Methods("GET")

	m.HandleFunc("/{path:.*}", gorca.NotFoundFunc)

	return m
}

// UserInfo is the information about the current user.
type UserInfo struct {
	// The users e-mail address.
	Email string

	// The url to use to log the user out.
	LogoutURL string
}

// GetUser sends a JSON response with the current user information.
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Get the current user.
	u, ok := gorca.GetUserOrUnexpected(w, r)
	if !ok {
		return
	}

	// Get their logout URL.
	logout, ok := gorca.GetUserLogoutURL(w, r, "/")
	if !ok {
		return
	}

	// Make the user struct we'll return.
	userInfo := UserInfo{
		Email:     u.Email,
		LogoutURL: logout,
	}

	// Write the user information out.
	gorca.WriteJSON(w, r, userInfo)
}
