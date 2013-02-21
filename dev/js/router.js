/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
 source code is governed by a BSD-style license that can be found in
 the LICENSE file.
 */

// This is the routing mechanism.
function Router($routeProvider) {
		$routeProvider
				.when('/lists/', {
						controller:ListsAllCtrl, 
						templateUrl: 'lists/all.html'
				})
				.when('/lists/view/:id', {
						controller:ListsViewCtrl, 
						templateUrl: 'lists/view.html'
				})
				.when('/recipes/', {
						controller:RecipesAllCtrl, 
						templateUrl: 'recipes/all.html'
				})
				.when('/recipes/view/:id', {
						controller:RecipesViewCtrl, 
						templateUrl: 'recipes/view.html'
				})
				.otherwise({redirectTo: '/lists/'});
}
