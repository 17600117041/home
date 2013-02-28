/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
 source code is governed by a BSD-style license that can be found in
 the LICENSE file.
 */

function HomeCtrl($scope, Links) {
		$scope.static_links = [
				{
						Url: "#/lists/",
						Name: "Lists",
						Icon: "icon-list",
						Description: "Lists of things to do, places to go, and " +
								"people to see."
				}, {
						Url: "#/recipes/",
						Name: "Recipes",
						Icon: "icon-book",
						Description: "An online recipe book."
				}, {
						Url: "#/links/",
						Name: "Links",
						Icon: "icon-globe",
						Description: "An online repository of links."
				}
		];

		
		$scope.links = $scope.static_links;
		Links.search("*home*", function(links) {
				$scope.links = $scope.static_links.concat(links);
		});
}
HomeCtrl.$inject = ['$scope', 'Links'];
