package list

import (
	"appengine"
	"appengine/datastore"
	"github.com/icub3d/gorca"
	"net/http"
	"time"
)

// NewListHelper is a helper function that creates a new list in the
// datastore for the given list. The given list is updated with the
// keys. If a failure occured, false is returned and a response was
// returned to the request. This case should be terminal.
func NewListHelper(c appengine.Context, w http.ResponseWriter,
	r *http.Request, l *List) bool {

	// Blank out the list key and the item keys.
	l.Key = ""
	for _, item := range l.Items {
		item.Key = ""
	}

	// Set the date to now.
	l.LastModified = time.Now()

	// Save the new list.
	if !PutListHelper(c, w, r, l) {
		return false
	}

	return true
}

// GetListHelper is a helper function that retrieves a list and it's
// items from the datastore. If a failure occured, false is returned
// and a response was returned to the request. This case should be
// terminal.
func GetListHelper(c appengine.Context, w http.ResponseWriter,
	r *http.Request, key string) (*List, bool) {

	// Decode the string version of the key.
	k, err := datastore.DecodeKey(key)
	if err != nil {
		gorca.LogAndUnexpected(c, w, r, err)
		return nil, false
	}

	// Get the list by key.
	var l List
	if err := datastore.Get(c, k, &l); err != nil {
		gorca.LogAndNotFound(c, w, r, err)
		return nil, false
	}

	// Get all of the items for the list.
	var li ItemsList
	q := datastore.NewQuery("Item").Ancestor(k).Order("Order")
	if _, err := q.GetAll(c, &li); err != nil {
		gorca.LogAndUnexpected(c, w, r, err)
		return nil, false
	}

	l.Items = li

	return &l, true
}

// PutListHelepr saves the list and it's items to the datastore. If
// the list or any of it's items don't have a key, a key will be made
// for it.
func PutListHelper(c appengine.Context, w http.ResponseWriter,
	r *http.Request, l *List) bool {

	// This is the list of keys we are going to PutMulti.
	keys := make([]string, 0, len(l.Items)+1)

	// This is the list of things we are going to put.
	values := make([]interface{}, 0, len(l.Items)+1)

	// We need the key to generate new keys and we might not even have
	// one.
	var lkey *datastore.Key
	var skey string
	var ok bool
	if l.Key != "" {
		// Just convert the key we already have.
		lkey, ok = gorca.StringToKey(c, w, r, l.Key)
		if !ok {
			return false
		}
	} else {
		// Make a key and save it's value to the list.
		skey, lkey, ok = gorca.NewKey(c, w, r, "List", nil)
		if !ok {
			return false
		}

		l.Key = skey
	}

	// Add the list itself.
	keys = append(keys, l.Key)
	values = append(values, l)

	// Add the items in the list.
	for _, item := range l.Items {
		if item.Key != "" {
			keys = append(keys, item.Key)
		} else {
			skey, _, ok := gorca.NewKey(c, w, r, "Item", lkey)
			if !ok {
				return false
			}

			item.Key = skey
			keys = append(keys, skey)
		}

		values = append(values, item)
	}

	// Save them all.
	return gorca.PutStringKeys(c, w, r, keys, values)
}
