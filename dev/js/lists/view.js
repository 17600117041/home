/*
	Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
	source code is governed by a BSD-style license that can be found in
	the LICENSE file.
*/

// ListsViewCtrl is the controller for viewing and updating lists.
function ListsViewCtrl($scope, $routeParams, $timeout, Lists) {

		// Cancel the update checks when we leave.
		$scope.$on('$destroy', function() {
									 $timeout.cancel($scope.timer);
							 });

		// Check for changes every 30 seconds.
		$scope.checkupdate = function() {
				Lists.checkupdate($scope.list, function(u) {
														 $scope.updatable = u;

														 if ($scope.updatable && !$scope.dirty) {
																 $scope.overwrite();
														 }
												 });
				$scope.timer = $timeout($scope.checkupdate, 30000);
		};
		$scope.timer = $timeout($scope.checkupdate, 30000);

		// merge is called when they want to merge their changes with 
		// the latest list.
		$scope.merge = function() {
				$scope.save();	
				$scope.updatable = false;
		};

		// overwrite is called when they want to just get the updated list 
		// and lose their changes.
		$scope.overwrite = function() {
				Lists.get($routeParams.id, function(l) {
										 $scope.list = l;
								 });
				$scope.updatable = false;
		};

		// sort updates the order of the list items. This is called to
		// update the internal array that is storing the list items.
		$scope.sort = function() {
				var ids = $("#sortablelist").sortable("toArray");
				var items = new Array();

				// Make a new list from the array.
				ids.forEach(function(e, i, a) {
												var eid = e.replace("list-item-", "");
												var id = parseInt(eid);
												items.push($scope.list.Items[id]);
										});

				// Set the list ot be the newly made list.				
				$scope.list.Items = items;
				$scope.dirty = true;
		};

		// add adds a new item to the top list of items.
		$scope.add = function() {
				// Make a list if we don't have one.
				if ($scope.list.Items == null) {
						$scope.list.Items = new Array();
				}

				// We unshift to add it to the top.
				$scope.list.Items.unshift({
																	 "Name": $scope.newitem,
																	 "Completed": false,
																	 "Delete": false
															 });
				$scope.dirty = true;
				$scope.newitem = "";
		};


		// sure marks completed items for deletion.
		$scope.sure = function() {
				$scope.list.Items
						.forEach(function (e) {
												 if (e.Completed) {
														 e.Delete = true;
												 }
										 });
				$scope.dirty = true;
		};

		// save saves changes to the list and update the list.
		$scope.save = function() {
				Lists.save($scope.list, function(l){
											$scope.list = l;
											$scope.dirty = false;
									});
		};

		// dirty is a helper function that marks the list dirty when an
		// item is checked.
		$scope.dirtify = function() {
				$scope.dirty = true;
		};

		// noshow is a helper function that filters the list to not show
		// items scheduled for deletion.
		$scope.noshow = function(item) {
				return !item.Delete;
		};
		
		// edit changes the state of the item being edited. There 
		// are two special values: -1 is for no item, and -2 is the 
		// title of the list.
		$scope.edit = function(id) {
				$scope.editing = id;
				$scope.dirty = true;
		};

		// Get the list items.
		Lists.get($routeParams.id, function(l) {
								 $scope.list = l;

								 // Make the list sortable.
								 $("#sortablelist")
										 .sortable({
																	 handle: '.drag-icon',
																	 stop: function() {
																			 $scope.$apply($scope.sort());
																	 }
															 });
						 });

		// Start out with cursor in the add text box.
		$('#newitem').focus();

		// We start out not editing any list.
		$scope.editing = -1;

		// The list should start clean.
		$scope.dirty = false;
}
ListsViewCtrl.$inject = ['$scope', '$routeParams', '$timeout', 'Lists'];
