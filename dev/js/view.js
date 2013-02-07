/*
	Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
	source code is governed by a BSD-style license that can be found in
	the LICENSE file.
*/

// ViewCtrl is the controller for viewing and updating lists.
function ViewCtrl($scope, $location, $routeParams, $timeout, List) {

		// Cancel the update checks when we leave.
		$scope.$on('$destroy', function() {
									 $timeout.cancel($scope.timer);
							 });

		// Check for changes every 30 seconds.
		$scope.checkupdate = function() {
				List.checkupdate($scope.list, function(u) {
														 $scope.updatable = u;

														 if ($scope.updatable && !$scope.dirty) {
																 $scope.updateoverwrite();
														 }
												 });
				$scope.timer = $timeout($scope.checkupdate, 30000);
		};
		$scope.timer = $timeout($scope.checkupdate, 30000);

		// This is called when the update button is pressed.
		$scope.update = function() {
				if ($scope.dirty) {
						// Ask them if they want to merge their changes 
						// if there are changes.
						$('#updateModal').modal();
				} else {
						// Otherwise, we can just update.
						List.get($routeParams.id, function(l) {
												 $scope.list = l;
										 });
						$scope.updatable = false;
				}
		};

		// This is called when they want to merge their changes with 
		// the latest list.
		$scope.updatemerge = function() {
				$scope.save();	
				$scope.updatable = false;
				$('#updateModal').modal('hide');
		};

		// This is called when they want to just get the updated list 
		// and lose their changes.
		$scope.updateoverwrite = function() {
				List.get($routeParams.id, function(l) {
										 $scope.list = l;
								 });
				$scope.updatable = false;
				$('#updateModal').modal('hide');
		};

		// When the order of the list items change, this is called to 
		// update the internal array that is storing the list items.
		$scope.sort = function(event, ui) {
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

		// Add a new item to the list of items.
		$scope.add = function() {
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


		// Bring up the modal to clean completed items.
		$scope.clean = function(key, event) {
				$('#cleanModal').modal();
		};

		// Mark completed items for deletion.
		$scope.sureclean = function() {
				$scope.list.Items
						.forEach(function (e) {
												 if (e.Completed) {
														 e.Delete = true;
												 }
										 });
				
				$scope.dirty = true;
				$('#cleanModal').modal('hide');
		};

		// Save changes to the list and update the list.
		$scope.save = function() {
				List.save($scope.list, function(l){
											$scope.list = l;
											$scope.dirty = false;
									});
		};

		// A helper function that marks the list dirty when an item is
		// checked.
		$scope.check = function(item) {
				item.Completed = !item.Completed;
				$scope.dirty = true;
		};

		// A helper function that filters the list to not show items
		// scheduled for deletion.
		$scope.noshow = function(item) {
				return !item.Delete;
		};
		
		// Stop propogating a click. This is mainly used by the list
		// item input text boxes so they don't change the completed 
		// state of the item when they are clicked.
		$scope.noclick = function(event) {
				// Don't let it fall through to the view.
				event.stopPropagation();
		};

		// This changes the state of the item being edited. There 
		// are two special values: -1 is for no item, and -2 is the 
		// title of the list.
		$scope.edit = function(id, event) {
				$scope.editing = id;
				$scope.dirty = true;
				// Don't let it fall through to the view.
				if (event != undefined) {
						event.stopPropagation();		
				}
		};

		// Get the list items.
		List.get($routeParams.id, function(l) {
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
}
ViewCtrl.$inject = ['$scope', '$location', '$routeParams', '$timeout', 'List'];