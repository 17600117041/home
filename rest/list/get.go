package list

import (
	"appengine"
	"appengine/datastore"
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
)

// GetAllLists fetches all of the lists.
func GetAllLists(w http.ResponseWriter, r *http.Request) {
	// Create the query.
	c := appengine.NewContext(r)
	q := datastore.NewQuery("List").Order("-LastModified")

	// Fetch the lists. We only want subset of the data, so we make a
	// struct with the fields we want.
	lists := []List{}
	if _, err := q.GetAll(c, &lists); err != nil {
		gorca.LogAndUnexpected(w, r, err)
		return
	}

	for _, list := range lists {
		list.Items = nil
		list.Sitems = nil
	}

	// Write the lists as JSON.
	gorca.WriteJSON(w, r, lists)
}

// GetList fetches the list for the given tag. The currently
// logged in user must own the list or the list must have been shared
// with the user. Otherwise, an unauthorized error is returned.
func GetList(w http.ResponseWriter, r *http.Request) {
	// Get the context.
	c := appengine.NewContext(r)

	// Get the Key.
	vars := mux.Vars(r)
	key := vars["key"]

	// Decode the string version of the key.
	k, err := datastore.DecodeKey(key)
	if err != nil {
		gorca.LogAndUnexpected(w, r, err)
		return
	}

	// Get the list by key.
	var l List
	if err := datastore.Get(c, k, &l); err != nil {
		gorca.LogAndNotFound(w, r, err)
		return
	}

	l.ConvertSitems()

	// Write the lists as JSON.
	gorca.WriteJSON(w, r, l)
}
