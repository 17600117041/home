package recipe

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"github.com/icub3d/gorca"
	"github.com/icub3d/home/rest/recipe/parsers"
	"io/ioutil"
	"net/http"
	"time"
)

// NewRecipeHelper is a helper function that creates a new recipe in
// the datastore for the given recipe. If the URL field is not empty
// and the direction and ingredients are, an attempt is made to parse
// the recipe from the URL. If a failure occured, false is returned
// and a response was returned to the request. This case should be
// terminal.
func NewRecipeHelper(c appengine.Context, w http.ResponseWriter,
	r *http.Request, recipe *Recipe) bool {

	// Blank out the recipe key and the item keys.
	recipe.Key = ""

	// Set the date to now.
	recipe.LastModified = time.Now()

	if recipe.URL != "" && len(recipe.Ingredients) == 0 &&
		len(recipe.Directions) == 0 {

		ok := updateRecipeFromURL(c, w, r, recipe)
		if !ok {
			return false
		}
	}

	// Save the new recipe.
	if !PutRecipeHelper(c, w, r, recipe) {
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
	var recipe Recipe
	if err := datastore.Get(c, k, &recipe); err != nil {
		gorca.LogAndNotFound(c, w, r, err)
		return nil, false
	}

	return &recipe, true
}

// PutRecipeHelepr saves the recipe to the datastore. If the recipe
// doesn't have a key, a key will be made for it.
func PutRecipeHelper(c appengine.Context, w http.ResponseWriter,
	r *http.Request, l *Recipe) bool {

	// We may need to make a key.
	if l.Key == "" {
		// Make a key and save it's value to the recipe.
		skey, _, ok := gorca.NewKey(c, w, r, "Recipe", nil)
		if !ok {
			return false
		}

		l.Key = skey
	}

	// Save them all.
	return gorca.PutStringKeys(c, w, r, []string{l.Key}, []interface{}{l})
}

// updateRecipeFromURL is a helper function that attempts to parse the
// recipe URL to get the recipe data. If an error occurs, false is
// returned and a proper message will have been sent as a
// response. This case should be terminal. If a parser isn't available
// for the URL, no error is returned, but nothing is changed in the
// recipe.
func updateRecipeFromURL(c appengine.Context, w http.ResponseWriter,
	r *http.Request, recipe *Recipe) bool {

	p, err := parsers.GetParserForURL(recipe.URL)
	if err != nil {
		gorca.LogAndUnexpected(c, w, r, err)
		return false
	}

	if p == nil {
		gorca.Log(c, r, "warn", "no parser found for: %s", recipe.URL)
		return true
	}

	client := urlfetch.Client(c)
	resp, err := client.Get(recipe.URL)
	if err != nil {
		gorca.LogAndUnexpected(c, w, r, err)
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gorca.LogAndUnexpected(c, w, r, err)
		return false
	}

	recipe.Name = p.GetName(body)
	recipe.Ingredients = p.GetIngredients(body)
	recipe.Directions = p.GetDirections(body)

	return true
}
