/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
 source code is governed by a BSD-style license that can be found in
 the LICENSE file.
 */

// This is the routing mechanism.
function Router($routeProvider) {
		$routeProvider
				.when('/', {
						controller: HomeCtrl, 
						templateUrl: 'partials/home.html'
				})
				.when('/links/', {
						controller: LinksAllCtrl, 
						templateUrl: 'partials/links/all.html'
				})
				.when('/links/:id', {
						controller: LinksEditCtrl, 
						templateUrl: 'partials/links/edit.html'
				})
				.when('/lists/', {
						controller: ListsAllCtrl, 
						templateUrl: 'partials/lists/all.html'
				})
				.when('/lists/view/:id', {
						controller: ListsViewCtrl, 
						templateUrl: 'partials/lists/view.html'
				})
				.when('/recipes/', {
						controller: RecipesAllCtrl, 
						templateUrl: 'partials/recipes/all.html'
				})
				.when('/recipes/view/:id', {
						controller: RecipesViewCtrl, 
						templateUrl: 'partials/recipes/view.html'
				})
				.otherwise({redirectTo: '/'});
}
