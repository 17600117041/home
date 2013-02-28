/*
	Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
	source code is governed by a BSD-style license that can be found in
	the LICENSE file.
*/

// LinksEditCtrl is the controller for viewing and updating links.
function LinksEditCtrl($scope, $routeParams, $location, Links) {

		$scope.back = function () {
				history.back();
		};

		// save saves changes to the link and update the link.
		$scope.save = function() {
				if ($scope.link.Key == undefined || $scope.link.Key == null ||
						$scope.link.Key == "")
						{
								Links.create($scope.link, function(l){
										$location.path("/links/");
								});

						} else {
								Links.save($scope.link, function(l){
										$location.path("/links/");
								});
						}
		};

		// Get the link.
		$scope.link = {"Tags": []};
		if ($routeParams.id != "new") {
				Links.get($routeParams.id, function(l) {
						$scope.link = l;
				});
		}
}

LinksEditCtrl.$inject = ['$scope', '$routeParams', '$location', 'Links'];
