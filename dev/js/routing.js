/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
 source code is governed by a BSD-style license that can be found in
 the LICENSE file.
 */

// This is the routing mechanism.
angular.module('home', ['rest'])
		.config(['$routeProvider', 
						 function ($routeProvider) {
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
						 
						])
// This is used to auto-focus the items then they are switched
// from a span to a text box. 
		.directive('ngHasfocus', 
							 function() {
									 return function(scope, element, attrs) {
											 scope.$watch(attrs.ngHasfocus, function (nVal, oVal) {
													 if (nVal) {
															 $(element[0]).show();
															 $(element[0]).focus();
															 $(element[0]).select();
													 }
											 });
											 
											 element.bind('blur', function() {
													 scope.$apply("edit(-1);");
											 });
											 
											 element.bind('keydown', function (e) {
													 if (e.which == 13) {
															 scope.$apply("edit(-1);");
													 }
											 });
									 };
							 });

