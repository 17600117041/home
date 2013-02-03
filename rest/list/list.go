// Package list provides functionality for managing lists.
package list

import (
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
	"time"
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

	m.HandleFunc("/", GetAllLists).Methods("GET")
	m.HandleFunc("/", PostList).Methods("POST")

	m.HandleFunc("/{key}/", GetList).Methods("GET")
	m.HandleFunc("/{key}/", PutList).Methods("PUT")
	m.HandleFunc("/{key}/", DeleteList).Methods("DELETE")

	m.HandleFunc("/{path:.*}", gorca.NotFoundFunc)

	return m
}

// Item is a single item in a List. 
type Item struct {
	// A URL safe version of the datastores key for this list item.
	Key string

	// The order of this item within the list. Lower numbers come before
	// higher numbers.
	Order int

	// The string representing the list item.
	Name string

	// The state of completion of this list item..
	Completed bool

	// If delete is true, when this item is merged with other items,
	// this item will be removed from the merged list. See
	// Merge for more information.
	Delete bool
}

// ItemsList is a list of Items that have the methods necessary to
// be sorted.
type ItemsList []*Item

// Len returns the length of this ItemsList.
func (li ItemsList) Len() int {
	return len(li)
}

// Less returns true if the Order of li[i] < li[j].
func (li ItemsList) Less(i, j int) bool {
	return li[i].Order < li[j].Order
}

// Swap switches the values of li[i] and li[j].
func (li ItemsList) Swap(i, j int) {
	li[i], li[j] = li[j], li[i]
}

// SetOrders fixes the order of the items by setting each Items order
// to its current position in the array.
func (li ItemsList) SetOrders() {
	for k, r := range li {
		r.Order = k
	}
}

// List is the structure used to save, get, and delete lists from the
// datastore.
type List struct {
	// The name of the list.
	Name string

	// A URL safe version of the datastores key for this list. It
	// is not stored in the datastore.
	Key string

	// This is the time the list was last modified.
	LastModified time.Time

	// The items in the list. This is used by the web application. The
	// datastore functions will retrieve these values for us.
	Items ItemsList `datastore:"-"`
}

// RemoveItem looks through this List's Items and returns an Item
// matching the given key or nil if it was not found. If an item was
// found, it is removed from the list.
func (l *List) RemoveItem(key string) *Item {
	for i, li := range l.Items {
		if li.Key == key {
			// Remove the item.
			l.Items = append(l.Items[:i], l.Items[i+1:]...)

			return li
		}
	}

	return nil
}

// Merge combines the Items in this List with the given List. This
// list items are then sorted by their order and given a new order
// based on the sorting. The list name is also changed. A list of keys
// is returned of items that should be deleted from the datastore.
func (l *List) Merge(m *List) []string {
	nlst := ItemsList{}

	del := []string{}

	l.Name = m.Name

	for _, r := range m.Items {
		// Get the accompanying list item.
		s := l.RemoveItem(r.Key)

		// If it is marked for deletion, we should skip this item.
		if s != nil && s.Delete {
			del = append(del, s.Key)
			continue
		}

		if r.Delete {
			del = append(del, r.Key)
			continue
		}

		// Save a new ListItem to the new list.
		nlst = append(nlst, &Item{
			Key:       r.Key,
			Order:     r.Order,
			Name:      r.Name,
			Completed: r.Completed,
			Delete:    false,
		})
	}

	// Add the remaining items from this list.
	for _, r := range l.Items {
		// If it is marked for deletion, we should skip this item.
		if r.Delete {
			del = append(del, r.Key)
			continue
		}

		nlst = append(nlst, &Item{
			Key:       r.Key,
			Order:     r.Order,
			Name:      r.Name,
			Completed: r.Completed,
			Delete:    false,
		})
	}

	// sort the list.
	nlst.SetOrders()

	// Set the list to our newly merged sorted list.
	l.Items = nlst

	return del
}
