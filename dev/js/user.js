/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
 source code is governed by a BSD-style license that can be found in
 the LICENSE file.
 */

// UserCtrl is the controller for part of the site that lists all of
// the lists.
function UserCtrl($scope, User) {
		User.get(function(data) {
								 $scope.user = data;
						 });
}
UserCtrl.$inject = ['$scope', 'User'];