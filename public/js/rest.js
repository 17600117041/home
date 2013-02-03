/*
	Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
	source code is governed by a BSD-style license that can be found in
	the LICENSE file.
*/

// This is a module for interacting with the rest framework.
angular.module('rest', ['ngResource'])
    /* This is an interface to the RESTful User service. */
		.factory('User', function($resource) {
								 var User = $resource("/rest/user/", {}, {});
								 
								 return User;
						 })
    /* This is an interface to the RESTful List service. We have 
		 * to use a service here because go and angular don't play nicely 
		 * together with trailing slashes. 
		 */
		.service('List', function($http) {
								 // Note: All of these functions accept callback
								 // functions in which the resulting data from the
								 // server is returned.

								 // Create a new list.
								 this.create = function(data, call) {
										 return $http.post("/rest/list/", data)
												 .then(function (response) {
																	 call(response.data);
															 });
								 };

								 // Save an existing list.
								 this.save = function(data, call) {
										 return $http.put("/rest/list/" + data.Key + "/", data)
												 .then(function (response) {
																	 call(response.data);
															 });
								 };

								 // Check to see if a list has been modified.
								 this.checkupdate = function(data, call) {
										 return $http.get("/rest/list/" + data.Key + "/")
												 .then(function (response) {
																	 // We are just going to compare LastModified dates.
																	 var sdate = Date.parseRFC3339(response.data.LastModified);
																	 var mdate = Date.parseRFC3339(data.LastModified);

																	 call(sdate > mdate);
															 });
										 
								 };

								 // Delete the list wit the given key.
								 this.delete = function(key, call) {
										 return $http.delete("/rest/list/" + key + "/")
												 .then(function(response) {
																	 call(response.data);
															 });
								 };
								 
								 // Get all lists.
								 this.getall = function(call) {
										 return $http.get("/rest/list/")
												 .then(function(response) {
																	 call(response.data);
															 });
								 };


								 // Get a specific list.
								 this.get = function(key, call) {
										 return $http.get("/rest/list/" + key + "/")
												 .then(function(response) {
																	 call(response.data);
															 });
								 };
						 });