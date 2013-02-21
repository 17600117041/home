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


		// split parses the given string and returns the parts of the
		// string: quantity, unit, and name. For example, '1 cup butter'
		// would return {quantity: '1', unit: 'cup', name, 'butter'}.
		$scope.split = function(str) {
				str = str.toLowerCase();

				// break the string out by word.
				var chunks = str.split(' ');

				// These are our states in the state machine.
				var QTY = 1, UNIT = 2, NAME = 3, MODIFIERS = 4, INPAREN = 5;
				var state = QTY, prev = 0;

				// These are the arrays we'll be storing information in.
				var qty = [];
				var units = [];
				var name = [];
				var modifiers = [];

				// Loop through each chunk.
				for (var x = 0; x < chunks.length; x++) {
						// Clean up the end. We don't want the punctuation.
						var e = chunks[x].replace(/[\.,]$/, '');

						// If it's an empty string, ignore it.
						if (e == "") continue;

						// First check to the paren states.
						if (state == INPAREN) {
								if (e.lastIndexOf(')') == e.length - 1) {
										state = prev;
								}

								continue;
						} 
						
						if (e.indexOf('(') == 0 && state != NAME) {
								prev = state;
								state = INPAREN;
								continue;
						}	
						
						// Check the quantity state.
						if (state == QTY) {
								// See if we are still getting a number.
								if (/^[-0-9\/\.]*$/.test(e)) {

										// We are going to just take the higher value of a
										// range for now. It might be good to eventually
										// be able to sum up the lower and upper bound
										// quantities.
										if (e.indexOf('-') > -1)
												e = e.substr(e.indexOf('-') + 1);

										qty.push(e);
										continue;
								} else if (e == "+" || e == "plus") {
										// There is another quantity and unit, so we
										// should start over.
										units.push('');
										continue;
								} else {
										state = UNIT;
								}
						}

						// Check the unit state. The QTY state might fall the
						// current value through to here.
						if (state == UNIT) {
								if ($.inArray(e, window.Units) > -1) {
										units.push(e);
										continue;
								} else if (e == "+" || e == "plus") {
										// There is another quantity and unit, so we
										// should start over.
										if (qty.length != units.length) qty.push('1');
										state = QTY;
										continue;
								} else {
										state = MODIFIERS;
								}
						}

						// Check the modifiers state. The UNIT state might fall
						// the current value through to here.
						if (state == MODIFIERS) {
								if ($.inArray(e, window.Modifiers) > -1) {
										modifiers.push(e);
										continue;
								} else {
										state = NAME;
								}
						}

						// Check the name state. The MODIFIERS state might fall
						// the current value through to here.
						if (state == NAME) {
								// If we find another modifier at this point, we can
								// drop it. We do this so something like '1 cup
								// butter, clarified' becomes '1 cup butter'.
								if ($.inArray(e, window.Modifiers) > -1)
										break;

								name.push(e);
						}

				}

				// Join the modifiers and name
				var joinedname = modifiers.sort().join(', ') + " " + name.join(' ');

				// Attempt to remove the end of the last word it it's pluralized.
				joinedname = joinedname.trim().replace(/es$/, '');
				if (/s$/.test(joinedname)) {
						if (!/[aeious]s$/.test(joinedname)) {
								joinedname = joinedname.replace(/s$/, '');
						}
				}

				if (qty.length == 0)
						qty.push("1");

				if (units.length == 0)
						units.push('');

				return {
						name: joinedname,
						quantity: qty,
						units: units
				};
		};

		// normalize attempts to combine the given quantities into a
		// single quantity.
		$scope.normalize = function(quantities) {
				var combined = {};
				
				for (var x = 0; x < quantities.length; x++) {
						var quantity = quantities[x];

						// Initilize to 0 if we haven't seen this unit yet.
						if (combined[quantity.units] == null || combined[quantity.units] == undefined) {
								combined[quantity.units] = 0;
						}

						// Add the quantity to the unit.
						combined[quantity.units] += eval(quantity.quantity);
				}

				// Join all the separate units. We should eventually attempt
				// to convert them to the same single unit.
				var total = "";
				for (var unit in combined) {
						total += combined[unit].toFixed(2) + " " + unit + " + ";
				}

				return total.replace(/ \+ $/, '').replace(/\.00/, '');
		};

		// combine is called when the 'Merge' button is pushed. It tries
		// to find similar elements and combine their parts.
		$scope.combine = function() {
				$scope.dirty = true;
 
				// This is the list of items and their key and current
				// quantities.
				var merged = {};

				$scope.list.Items.forEach(function(e, i) {
						if (e.Completed || e.Delete) return;
						var parts = $scope.split(e.Name);

						// Add the item if we don't have it.
						if (merged[parts.name] == undefined || merged[parts.name] == null) {
								merged[parts.name] = {
										quantities: [],
										index: i,
										key: e.Key
								};
						} else {
								// Mark it completed to schedule for deletion. We
								// should only do it if the item was already found.
								e.Delete = true;
						}

						// Add the quantitys and units to the item.
						var found = merged[parts.name];
								
						// Add each item one at at a time.
						for (var x = 0; x < parts.quantity.length; x++) {
								var unit = "";
								if (x < parts.units.length)
										unit = parts.units[x];
								
								found.quantities.push({quantity: parts.quantity[x], units: unit});
						}
						
				});

				// Now we need to merge the quantities and save them to the right item.
				for (var name in merged) {
						var n = $scope.normalize(merged[name].quantities)
								+ " " + name;
						n = n.trim();
						n = n.replace('/  /', ' ');
						$scope.list.Items[merged[name].index].Name = n;
				}
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

		// We start out not editing any list.
		$scope.editing = -1;

		// The list should start clean.
		$scope.dirty = false;
}
ListsViewCtrl.$inject = ['$scope', '$routeParams', '$timeout', 'Lists'];
