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
