/*
	Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
	source code is governed by a BSD-style license that can be found in
	the LICENSE file.
*/

// UserCtrl is the controller for part of the site that lists all of
// the lists.
function AlertsCtrl($scope, Alerts) {
		$scope.alerts = Alerts.alerts;
		
		$scope.$watch(Alerts.alerts, function() {
											$scope.alerts = Alerts.alerts;
									});

		$scope.remove = function(index) {
				Alerts.remove(index);
		};
}
