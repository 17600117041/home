/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
 source code is governed by a BSD-style license that can be found in
 the LICENSE file.
 */

// LinksCtrl is the controller for viewing all links.
function LinksAllCtrl($scope, $location, Links) {
		// del prepare the delete values that might be used if the user
		// verifies they want to delete an item.
		$scope.del = function(index, key) {
				$scope.data.index = index;
				$scope.data.key = key;
		};
		
		// sure performs the actual delete.
		$scope.sure = function() {
				Links.del($scope.data.key, function() {
						$scope.links.splice($scope.data.index, 1);
				});
		};
		
		// save creates the new link and redirects to that link.
		$scope.create = function() {
				$location.path('/links/new/');
		};

		// To start off, we should get all the links.
		Links.getall(function (links) {
				$scope.links = links;
		});


		$scope.search = function() {
				Links.search($scope.data.search, function (links) {
						$scope.links = links;
				});
		};

		$scope.data = {};
}
LinksAllCtrl.$inject = ['$scope', '$location', 'Links'];
