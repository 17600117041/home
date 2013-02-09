package recipe

import (
	"appengine"
	"appengine/datastore"
	"github.com/icub3d/gorca"
	"net/http"
	"time"
)

// NewRecipeHelper is a helper function that creates a new recipe in the
// datastore for the given recipe. The given recipe is updated with the
// keys. If a failure occured, false is returned and a response was
// returned to the request. This case should be terminal.
func NewRecipeHelper(c appengine.Context, w http.ResponseWriter,
	r *http.Request, l *Recipe) bool {

	// Blank out the recipe key and the item keys.
	l.Key = ""
	for _, item := range l.Items {
		item.Key = ""
	}

	// Set the date to now.
	l.LastModified = time.Now()

	// Save the new recipe.
	if !PutRecipeHelper(c, w, r, l) {
		return false
	}

	return true
}

// GetRecipeHelper is a helper function that retrieves a recipe and it's
// items from the datastore. If a failure occured, false is returned
// and a response was returned to the request. This case should be
// terminal.
func GetRecipeHelper(c appengine.Context, w http.ResponseWriter,
	r *http.Request, key string) (*Recipe, bool) {

	// Decode the string version of the key.
	k, err := datastore.DecodeKey(key)
	if err != nil {
		gorca.LogAndUnexpected(c, w, r, err)
		return nil, false
	}

	// Get the recipe by key.
	var l Recipe
	if err := datastore.Get(c, k, &l); err != nil {
		gorca.LogAndNotFound(c, w, r, err)
		return nil, false
	}

	// Get all of the items for the recipe.
	var li ItemsRecipe
	q := datastore.NewQuery("Item").Ancestor(k).Order("Order")
	if _, err := q.GetAll(c, &li); err != nil {
		gorca.LogAndUnexpected(c, w, r, err)
		return nil, false
	}

	l.Items = li

	return &l, true
}

// PutRecipeHelepr saves the recipe and it's items to the datastore. If
// the recipe or any of it's items don't have a key, a key will be made
// for it.
func PutRecipeHelper(c appengine.Context, w http.ResponseWriter,
	r *http.Request, l *Recipe) bool {

	// This is the recipe of keys we are going to PutMulti.
	keys := make([]string, 0, len(l.Items)+1)

	// This is the recipe of things we are going to put.
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
		// Make a key and save it's value to the recipe.
		skey, lkey, ok = gorca.NewKey(c, w, r, "Recipe", nil)
		if !ok {
			return false
		}

		l.Key = skey
	}

	// Add the recipe itself.
	keys = append(keys, l.Key)
	values = append(values, l)

	// Add the items in the recipe.
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
