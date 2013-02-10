/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
 source code is governed by a BSD-style license that can be found in
 the LICENSE file.
 */

// RecipesViewCtrl is the controller for viewing and updating recipes.
function RecipesViewCtrl($scope, $routeParams, $timeout, 
												 $location, Recipes, Lists) {

		// sort is called When the order of the recipe items change. This
		// is called to update the internal array that is storing the
		// recipe items. The name passed differentiates between
		// ingredients and directions.
		$scope.sort = function(name) {
				var ids = $("#sortable" + name).sortable("toArray");
				var items = new Array();

				// Make a new list from the array.
				ids.forEach(function(e, i, a) {
						var eid = e.replace("recipe-"+name+"-", "");
						var id = parseInt(eid);
						items.push($scope[name][id]);
				});

				// Set the list to be the newly made list.
				$scope[name] = items;
				$scope.dirty = true;
		};

		// add adds a new item to the recipe. The given type is used to
		// add it to the right place: ingredients or directions.
		$scope.add = function(type) {
				// Create the list if we don't have one yet.
				if ($scope[type] == null) {
						$scope[type] = new Array();
				}

				// We push to add it to the bottom.
				$scope[type].push({value:$scope.toadd[type]});
				$scope.dirty = true;
				$scope.toadd[type] = "";
		};

		// getlists loads a list of current lists for the add to list
		// modal.
		$scope.getlists = function() {
				Lists.getall(function (lists) {
						$scope.lists = lists;
				});
		};

		// copy makes a copy of each of the ingredients and push it onto
		// the list. It then save the list.
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

		// del removes the given item from the list.
		$scope.del = function(type, index) {
				$scope[type].splice(index, 1);
				$scope.dirty = true;
		};

		// Save saves changes to the recipe back to the datastore.
		$scope.save = function() {
				$scope.saveobj();
				Recipes.save($scope.recipe, function(l){
						$scope.recipe = l;
						$scope.makeobj();
						$scope.dirty = false;
				});
		};

		// edit changes the state of the item being edited. There are two
		// special values: -1 is for no item, and -2 is the title of the
		// recipe. The type determines withere it's the ingredient or
		// direction that should be edited.
		$scope.edit = function(id, type) {
				$scope.editingtype = type;
				$scope.editing = id;
				$scope.dirty = true;
		};

		// makeobj converts the []string of ingredients to objects that
		// angularjs can handle. This should be called when we get a new
		// recipe from the datastore.
		$scope.makeobj = function() {
				// Get an object array of the ingredients with the string
				// saved to the object value.
				var ingredients = new Array();
				if ($scope.recipe.Ingredients != undefined &&
						$scope.recipe.Ingredients != null) {
						$scope.recipe.Ingredients.forEach(function(e, i, a) {
								ingredients.push({value: e});
						});
				}

				// Get an object array of the directions with the string
				// saved to the object value.
				var directions = new Array();
				if ($scope.recipe.Directions != undefined &&
						$scope.recipe.Directions != null) {
						$scope.recipe.Directions.forEach(function(e, i, a) {
								directions.push({value: e});
						});
				}

				// Set the scopes lists to the lists we just made.
				$scope.Ingredients = ingredients;
				$scope.Directions = directions;
		};

		// saveobj converts the objects back to a []string so go can save
		// them. This should be called before we save the recipe to the
		// datastore.
		$scope.saveobj = function() {
				// Generate a []string of the ingredients from the objects
				// values.
				var ingredients = new Array();
				if ($scope.Ingredients != undefined &&
						$scope.Ingredients != null) {
						$scope.Ingredients.forEach(function(e, i, a) {
								ingredients.push(e.value);
						});
				}

				// Generate a []string of the directions from the objects
				// values.
				var directions = new Array();
				if ($scope.Directions != undefined &&
						$scope.Directions != null) {
						$scope.Directions.forEach(function(e, i, a) {
								directions.push(e.value);
						});
				}

				// Set the recipes lists to the lists we just made.
				$scope.recipe.Ingredients = ingredients;
				$scope.recipe.Directions = directions;
		};

		// Get the recipe items.
		Recipes.get($routeParams.id, function(l) {
				$scope.recipe = l;
				$scope.makeobj();

				// Make the ingredients sortable.
				$("#sortableIngredients")
						.sortable({
								handle: '.drag-icon',
								stop: function() {
										$scope.$apply($scope.sort('Ingredients'));
								}
						});

				// Make the directions sortable.
				$("#sortableDirections")
						.sortable({
								handle: '.drag-icon',
								stop: function() {
										$scope.$apply($scope.sort('Directions'));
								}
						});
		});


		$scope.togglecleanup = function() {
				$scope.cleanup = !$scope.cleanup;
		};

		// We start out not editing any recipe.
		$scope.editingtype = "";
		$scope.editing = -1;
		$scope.cleanup = false;

		// The recipe should start clean.
		$scope.dirty = false;
		$scope.toadd = {
				"Ingredients": "",
				"Directions": ""
		};

		// Add shit+enter to save for the direction box.
		$('#newdirection').on('keyup', function (event) {
				if (event.which == 13 && event.shiftKey) {
						$scope.$apply($scope.add('Directions'));
				}
		});
}
RecipesViewCtrl.$inject = ['$scope', '$routeParams', '$timeout', '$location', 'Recipes', 'Lists'];
