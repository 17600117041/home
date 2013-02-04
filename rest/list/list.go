// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

// Package list provides functionality for managing lists.
package list

import (
	"time"
)

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

// ItemsList is a list of Items that has the ability reset the orders
// of it's itms based on the current position in the array.
type ItemsList []*Item

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

	// The items in the list. This is used by the web
	// application. Within the datastore, these are children of the
	// list. The get, put, etc. functions will manage this automagically
	// if you use them.
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

// Merge combines the Items in this List with the given List. The
// given list is considered the authority on the order of list
// items. The items are then given a new order based on their position
// in the array. A list of keys is returned of items that were removed.
func (l *List) Merge(m *List) []string {
	// This is the new list.
	nlst := ItemsList{}

	// This is the items marked for deletion. They'll need to be removed
	// from the datastore for the deletes to persist.
	del := []string{}

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

		// Save a new Item to the new list.
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

		// Add the item.
		nlst = append(nlst, &Item{
			Key:       r.Key,
			Order:     r.Order,
			Name:      r.Name,
			Completed: r.Completed,
			Delete:    false,
		})
	}

	// Set the list order.
	nlst.SetOrders()

	// Set the list to our newly merged sorted list.
	l.Items = nlst

	return del
}
