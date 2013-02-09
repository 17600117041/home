/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
 source code is governed by a BSD-style license that can be found in
 the LICENSE file.
 */

// RecipesViewCtrl is the controller for viewing and updating recipes.
function RecipesViewCtrl($scope, $routeParams, $timeout, $location, Recipes, Lists) {

		// When the order of the recipe items change, this is called to 
		// update the internal array that is storing the recipe items.
		$scope.sort = function(name) {
				var ids = $("#sortable" + name).sortable("toArray");
				var items = new Array();

				// Make a new recipe from the array.
				ids.forEach(function(e, i, a) {
						var eid = e.replace("recipe-"+name+"-", "");
						var id = parseInt(eid);
						items.push($scope[name][id]);
				});

				// Set the recipe ot be the newly made recipe.				
				$scope[name] = items;
				$scope.dirty = true;
		};

		// Add a new item to the recipe.
		$scope.add = function(type) {
				if ($scope[type] == null) {
						$scope[type] = new Array();
				}

				// We push to add it to the bottom.
				$scope[type].push({value:$scope.toadd[type]});
				$scope.dirty = true;
				$scope.toadd[type] = "";
		};

		// If we add this to a list, we need to get all the lists.
		$scope.getlists = function() {
				Lists.getall(function (lists) {
						$scope.lists = lists;
				});
		};

		// Make a copy of each of the ingredients and push it onto the
		// list. Then save the list.
		$scope.copy = function() {
				$scope.copyList.Items = new Array();
				$scope.Ingredients.forEach(function(e, i, a) {
						$scope.copyList.Items.push({
								Name: e.value,
								Completed: false,
								Delete: false
						});
				});

				Lists.save($scope.copyList, function(data) {
						$location.path('/lists/view/' + data.Key + '/');
				});
		};


		// Save changes to the recipe and update the recipe.
		$scope.save = function() {
				$scope.saveobj();
				Recipes.save($scope.recipe, function(l){
						$scope.recipe = l;
						$scope.makeobj();
						$scope.dirty = false;
				});
		};

		// This changes the state of the item being edited. There 
		// are two special values: -1 is for no item, and -2 is the 
		// title of the recipe.
		$scope.edit = function(id, type) {
				$scope.editingtype = type;
				$scope.editing = id;
				$scope.dirty = true;
		};

		// Convert the []string of ingredients to objects that angularjs
		// can handle.
		$scope.makeobj = function() {
				var ingredients = new Array();
				if ($scope.recipe.Ingredients != undefined &&
						$scope.recipe.Ingredients != null) {
						$scope.recipe.Ingredients.forEach(function(e, i, a) {
								ingredients.push({value: e});
						});
				}
				var directions = new Array();
				if ($scope.recipe.Directions != undefined &&
						$scope.recipe.Directions != null) {
						$scope.recipe.Directions.forEach(function(e, i, a) {
								directions.push({value: e});
						});
				}
				$scope.Ingredients = ingredients;
				$scope.Directions = directions;
		};

		// Convert the objects back to a []string so go can save them.
		$scope.saveobj = function() {
				var ingredients = new Array();
				if ($scope.Ingredients != undefined &&
						$scope.Ingredients != null) {
						$scope.Ingredients.forEach(function(e, i, a) {
								ingredients.push(e.value);
						});
				}
				var directions = new Array();
				if ($scope.Directions != undefined &&
						$scope.Directions != null) {
						$scope.Directions.forEach(function(e, i, a) {
								directions.push(e.value);
						});
				}
				$scope.recipe.Ingredients = ingredients;
				$scope.recipe.Directions = directions;
		};

		// Get the recipe items.
		Recipes.get($routeParams.id, function(l) {
				$scope.recipe = l;
				$scope.makeobj();

				// Make the recipe sortable.
				$("#sortableIngredients")
						.sortable({
								handle: '.drag-icon',
								stop: function() {
										$scope.$apply($scope.sort('Ingredients'));
								}
						});

				// Make the recipe sortable.
				$("#sortableDirections")
						.sortable({
								handle: '.drag-icon',
								stop: function() {
										$scope.$apply($scope.sort('Directions'));
								}
						});
		});

		// We start out not editing any recipe.
		$scope.editingtype = "";
		$scope.editing = -1;

		// The recipe should start clean.
		$scope.dirty = false;
		$scope.toadd = {
				"Ingredients": "",
				"Directions": ""
		};
}
RecipesViewCtrl.$inject = ['$scope', '$routeParams', '$timeout', '$location', 'Recipes', 'Lists'];
