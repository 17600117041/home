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
				$scope.copyKey = key;
				$scope.copyName = "Copy of " + name;
		};

		// copy actually makes the copy of the recipe.
		$scope.copy = function() {
				Recipes.get($scope.copyKey, function(l) {
						l.Name = $scope.copyName;
						Recipes.create(l, function(nl) {
								$location.path('/recipes/view/' + nl.Key + '/');
						});
				});
		};
		
		// del prepare the delete values that might be used if
		// the user verifies they want to delete a recipe.
		$scope.del = function(index, key) {
				$scope.delIndex = index;
				$scope.delKey = key;
		};
		
		// sure performs the actual delete.
		$scope.sure = function() {
				Recipes.del($scope.delKey, function() {
						$scope.recipes.splice($scope.delIndex, 1);
				});
		};
		
		// save creates the new recipe and redirects you to that recipe.
		$scope.save = function() {
				Recipes.create({"Name": $scope.name, "URL": $scope.URL},
											 function (l) {
													 $location.path('/recipes/view/' + l.Key + '/');
											 });
		};

		// To start off, we should get all the recipes.
		Recipes.getall(function (recipes) {
				$scope.recipes = recipes;
		});
}
RecipesAllCtrl.$inject = ['$scope', '$location', 'Recipes'];
