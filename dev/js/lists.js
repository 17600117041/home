/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
 source code is governed by a BSD-style license that can be found in
 the LICENSE file.
 */

// ListsCtrl is the controller for viewing all lists.
function ListsCtrl($scope, $location, List) {
		// The entire list box is clickable and this function handles that
		// by changing the location path.
		$scope.view = function(key) {
				$location.path('/view/' + key + '/');
		};

		// This handles clicking the copy button. It pops up the modal and
		// prepares it's name.
		$scope.copymodal = function(key, name, event) {
				$scope.copyKey = key;
				$scope.copyName = "Copy of " + name;

				$('#copyModal').modal();
				
				// Don't let it fall through to the view.
				event.stopPropagation();
		};

		// This function actually makes the copy of the list.
		$scope.copy = function() {
				$('#copyModal').modal('hide');
				
				List.get($scope.copyKey, function(l) {
						l.Name = $scope.copyName;
						List.create(l, function(nl) {
								$location.path('/view/' + nl.Key + '/');
						});
				});
				
		};
		
		// This function opens up the modal to verify that 
		// they want to delete the list.
		$scope.del = function(index, key, event) {
				$scope.delIndex = index;
				$scope.delKey = key;
				
				$('#deleteModal').modal();

				// Don't let it fall through to the view.
				event.stopPropagation();
		};
		
		// This function performs the actual delete.
		$scope.sure = function() {
				$('#deleteModal').modal('hide');
				List.del($scope.delKey, function() {
						$scope.lists.splice($scope.delIndex, 1);
				});
		};
		
		// This function opens up the new modal box.
		$scope.create = function() {
				$('#newModal').modal();
		};

		// If any of the modals click cancel, this is 
		// called to close it.
		$scope.back = function() {
				$('#newModal').modal('hide');
				$('#copyModal').modal('hide');
		};

		// This creates the new list and redirects you to 
		// that list.
		$scope.save = function() {
				$('#newModal').modal('hide');
				List.create({"Name": $scope.name}, function (l) {
						$location.path('/view/' + l.Key + '/');
				});
		};

		// To start off, we should get all the lists.
		List.getall(function (lists) {
				$scope.lists = lists;
		});

}
ListsCtrl.$inject = ['$scope', '$location', 'List'];