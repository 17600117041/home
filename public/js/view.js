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

														 // If there are no changes and we have
														 // changes, let's save them.
														 if (u == false && $scope.dirty) {
																 $scope.save();
														 }
												 });
				$scope.timer = $timeout($scope.checkupdate, 30000);
		};
		$scope.timer = $timeout($scope.checkupdate, 30000);

		$scope.update = function() {
				if ($scope.dirty) {
						// Ask them if they want to merge their changes.
						$('#updateModal').modal();
				} else {
						List.get($routeParams.id, function(l) {
												 $scope.list = l;
										 });
						$scope.updatable = false;
				}
		};

		$scope.updatemerge = function() {
				$scope.save();	
				$scope.updatable = false;
				$('#updateModal').modal('hide');
		};

		$scope.updateoverwrite = function() {
				List.get($routeParams.id, function(l) {
										 $scope.list = l;
								 });
				$scope.updatable = false;
				$('#updateModal').modal('hide');
		};

		$scope.sort = function(event, ui) {
				var ids = $("#sortablelist").sortable("toArray");
				var items = new Array();
				ids.forEach(function(e, i, a) {
												var eid = e.replace("list-item-", "");
												var id = parseInt(eid);
												items.push($scope.list.Items[id]);
										});
				
				$scope.list.Items = items;
				$scope.dirty = true;
		};

		// Add a new item to the list of items.
		$scope.add = function() {
				if ($scope.list.Items == null) {
						$scope.list.Items = new Array();
				}

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
		
		$scope.noclick = function(event) {
				// Don't let it fall through to the view.
				event.stopPropagation();
		};

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

		$('#newitem').focus();
		$scope.editing = -1;
}