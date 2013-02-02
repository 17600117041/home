// Package list provides functionality for managing lists.
package list

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
	"strings"
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

// ListItem is a single item in a List. 
type ListItem struct {
	// The string representing the list item.
	Name string

	// The state of completion of this list item..
	Completed bool

	// If delete is true, when this item is merged with other items,
	// this item will be removed from the merged list. See
	// MergeListItems for more information.
	Delete bool
}

// List is the structure used to save, get, and delete lists from the
// datastore.
type List struct {
	// The name of the list.
	Name string

	// A URL safe version of the datastores key for this list item. It
	// is not stored in the datastore.
	Key string

	// This is the time the list was last modified.
	LastModified time.Time

	// The items in the list. This is used by the application.
	Items []ListItem `datastore:"-"`

	// This is the string representation for the list of items. The
	// datastore won't save lists of structs, so we convert it on the
	// fly and it's saved here.
	Sitems []string `json:"-"`
}

// ConvertSitems takes the Sitems and translates them into Items and
// replaces the current Items with the translated values.
func (l *List) ConvertSitems() {
	l.Items = make([]ListItem, 0, len(l.Sitems))
	for _, s := range l.Sitems {
		parts := strings.SplitN(s, "|", 2)

		completed := false
		if parts[0] == "true" {
			completed = true
		}

		l.Items = append(l.Items, ListItem{
			Name:      parts[1],
			Completed: completed,
		})
	}
}

// ConvertItems translates the Items list into their string
// representation and saves them to Sitems.
func (l *List) ConvertItems() {
	l.Sitems = make([]string, 0, len(l.Items))
	for _, i := range l.Items {
		l.Sitems = append(l.Sitems,
			fmt.Sprintf("%v|%s", i.Completed, i.Name))
	}
}

// RemoveItem looks through this List's Items and returns a ListItem
// matching the given name or nil if it was not found. If an item was
// found, it is removed from the list.
func (l *List) RemoveItem(name string) *ListItem {
	for i, li := range l.Items {
		if li.Name == name {
			// Remove the item.
			l.Items = append(l.Items[:i], l.Items[i+1:]...)

			return &li
		}
	}

	return nil
}

// Merge combines modifies the ListItems in this List with the given
// List. The incoming list is considered to be the authority on the
// order. The only exception is that any ListItem with Delete set to
// true will not be in the merged list.
func (l *List) Merge(m *List) {
	var nlst []ListItem

	for _, r := range m.Items {
		// Get the accompanying list item.
		s := l.RemoveItem(r.Name)

		// If it is marked for deletion, we should skip this item.
		if (s != nil && s.Delete) || r.Delete {
			continue
		}

		// Save a new ListItem to the new []ListItem.
		nlst = append(nlst, ListItem{
			Name:      r.Name,
			Completed: r.Completed,
			Delete:    false,
		})
	}

	// Add the remaining items from this list.
	for _, r := range l.Items {
		// If it is marked for deletion, we should skip this item.
		if r.Delete {
			continue
		}

		// Save a new ListItem to the new []ListItem.
		nlst = append(nlst, ListItem{
			Name:      r.Name,
			Completed: r.Completed,
			Delete:    false,
		})
	}

	l.Items = nlst
}
