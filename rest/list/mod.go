package list

import (
	"appengine"
	"appengine/datastore"
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
	"time"
)

// PostList creates a new list
func PostList(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// Get the list from the body.
	var l List
	if !gorca.UnmarshalFromBodyOrFail(w, r, &l) {
		return
	}

	// Generate a new key for this list.
	id, _, err := datastore.AllocateIDs(c, "List", nil, 1)
	if err != nil {
		gorca.LogAndUnexpected(w, r, err)
		return
	}
	key := datastore.NewKey(c, "List", "", id, nil)

	// Save its string value to the list.
	l.Key = key.Encode()
	l.LastModified = time.Now()
	l.ConvertItems()

	// Save the list.
	if _, err := datastore.Put(c, key, &l); err != nil {
		gorca.LogAndUnexpected(w, r, err)
		return
	}

	// Return the updated list back.
	gorca.WriteJSON(w, r, l)
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

	// Decode the string version of the key.
	k, err := datastore.DecodeKey(key)
	if err != nil {
		gorca.LogAndUnexpected(w, r, err)
		return
	}

	// Get the old list by key.
	var ol List
	if err := datastore.Get(c, k, &ol); err != nil {
		gorca.LogAndNotFound(w, r, err)
		return
	}
	ol.ConvertSitems()

	// Get the new list from the body.
	var nl List
	if !gorca.UnmarshalFromBodyOrFail(w, r, &nl) {
		return
	}

	// Merge the new list into the old list.
	ol.Merge(&nl)

	// Set the last modified time.
	ol.LastModified = time.Now()

	// Save the changed list.
	ol.ConvertItems()
	if _, err := datastore.Put(c, k, &ol); err != nil {
		gorca.LogAndUnexpected(w, r, err)
		return
	}

	// Return the updated list back.
	gorca.WriteJSON(w, r, ol)
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

	// Decode the string version of the key.
	k, err := datastore.DecodeKey(key)
	if err != nil {
		gorca.LogAndUnexpected(w, r, err)
		return
	}

	// Delete the key.
	if err := datastore.Delete(c, k); err != nil {
		gorca.LogAndUnexpected(w, r, err)
		return
	}

	gorca.WriteSuccessMessage(w, r)
}
