/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
 source code is governed by a BSD-style license that can be found in
 the LICENSE file.
 */

// RecipesAllCtrl is the controller for viewing all recipes.
function RecipesAllCtrl($scope, $location, Recipes) {
		// setcopy handles clicking the copy button. It prepares the key to be
		// copied and the new name.
		$scope.setcopy = function(key, name) {
				$scope.data.key = key;
				$scope.data.name = "Copy of " + name;
		};

		// copy actually makes the copy of the recipe.
		$scope.copy = function() {
				Recipes.get($scope.data.key, function(l) {
						l.Name = $scope.data.name;
						Recipes.create(l, function(nl) {
								$location.path('/recipes/view/' + nl.Key + '/');
						});
				});
		};
		
		// del prepare the delete values that might be used if
		// the user verifies they want to delete a recipe.
		$scope.del = function(index, key) {
				$scope.data.index = index;
				$scope.data.key = key;
		};
		
		// sure performs the actual delete.
		$scope.sure = function() {
				Recipes.del($scope.data.key, function() {
						$scope.recipes.splice($scope.data.index, 1);
				});
		};
		
		// save creates the new recipe and redirects you to that recipe.
		$scope.save = function() {
				Recipes.create({"Name": $scope.data.name, "URL": $scope.data.URL},
											 function (l) {
													 $location.path('/recipes/view/' + l.Key + '/');
											 });
		};

		// To start off, we should get all the recipes.
		Recipes.getall(function (recipes) {
				$scope.recipes = recipes;
		});

		$scope.data = {};
}
RecipesAllCtrl.$inject = ['$scope', '$location', 'Recipes'];
