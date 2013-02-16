/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
 source code is governed by a BSD-style license that can be found in
 the LICENSE file.
 */

// ListsCtrl is the controller for viewing all lists.
function ListsAllCtrl($scope, $location, Lists) {
		// setcopy handles clicking the copy button. It prepares the key
		// to be copied and the new name.
		$scope.setcopy = function(key, name) {
				$scope.data.key = key;
				$scope.data.name = "Copy of " + name;
		};

		// copy actually makes the copy of the list and redirects to the
		// new list.
		$scope.copy = function() {
				Lists.get($scope.data.key, function(l) {
						l.Name = $scope.data.name;
						Lists.create(l, function(nl) {
								$location.path('/lists/view/' + nl.Key + '/');
						});
				});
		};
		
		// del prepare the delete values that might be used if the user
		// verifies they want to delete an item.
		$scope.del = function(index, key) {
				$scope.data.index = index;
				$scope.data.key = key;
		};
		
		// sure performs the actual delete.
		$scope.sure = function() {
				Lists.del($scope.data.key, function() {
						$scope.lists.splice($scope.data.index, 1);
				});
		};
		
		// save creates the new list and redirects to that list.
		$scope.save = function() {
				Lists.create({"Name": $scope.data.name}, function (l) {
						$location.path('/lists/view/' + l.Key + '/');
				});
		};

		// To start off, we should get all the lists.
		Lists.getall(function (lists) {
				$scope.lists = lists;
		});

		// This is where we'll get our form information from the modals.
		$scope.data = {};
}
ListsAllCtrl.$inject = ['$scope', '$location', 'Lists'];
