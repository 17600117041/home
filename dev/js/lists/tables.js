/*
	Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
	source code is governed by a BSD-style license that can be found in
	the LICENSE file.
*/

(function (insertion) {
		var units = [
				'tablespoon',	'tablespoons', 'tbsp', 'tbsps',
				'teaspoon',	'teaspoons', 'tsp',	'tsps',
				'pinch', 'pinches',
				'small', 'medium', 'large',
				'clove',	'cloves',
				'pound', 'pounds',
				'lb',	'lbs',
				'cup', 'cups',
				'ounce', 'ounces', 'oz',	'ozs', 'ozes',
				'can', 'cans',
				'dozen',
				'fl',	'fluid',
				'pint',	'pints', 'p',
				'liter', 'liters', 'l',
				'quart', 'quarts', 'q',	'qt',
				'gallon', 'gallons', 'gal', 
				'ml', 'mls', 'milliliter', 'milliliters',
				'grams', 'gram', 'g',
				'kg', 'kgs', 'kilogram', 'kilograms',
				'inch', 'inches', 'in', 'ins',
				'feet', 'foot', 'ft'
				];

		var modifiers = [
				'clarified',
				'boneless',
				'skinless',
				'halves',
				'halved',
				'such',
				'dry',
				'fresh',
				'frozen',
				'membranes',
				'fine',
				'finely',
				'freshly',
				'minced',
				'diced',
				'cubed',
				'small',
				'large',
				'medium',
				'rinsed',
				'and',
				'drained',
				'defrosted',
				'chopped',
				'membranes',
				'peeled',
				'seeded',
				'cut',
				'for',
				'sliced'
				];


		insertion.Units = units;
		insertion.Modifiers = modifiers;
})(window || this);
