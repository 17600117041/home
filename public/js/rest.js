// This is a module for interacting with the rest framework.
angular.module('rest', ['ngResource'])
		.factory('User', function($resource) {
								 var User = $resource("/rest/user/", {}, {});
								 
								 return User;
						 })
		.service('List', function($http) {
								 this.create = function(data, call) {
										 return $http.post("/rest/list/", data)
												 .then(function (response) {
																	 data.Key = response.data.Key;
																	 call(data);
															 });
								 };

								 this.save = function(data, call) {
										 return $http.put("/rest/list/" + data.Key + "/", data)
												 .then(function (response) {
																	 call(response.data);
															 });
								 };

								 this.delete = function(key) {
										 return $http.delete("/rest/list/" + key + "/")
												 .then(function(response) {
																	 return response.data;
															 });
								 };
								 
								 this.getall = function(call) {
										 return $http.get("/rest/list/")
												 .then(function(response) {
																	 call(response.data);
															 });
								 };


								 this.get = function(key, call) {
										 return $http.get("/rest/list/" + key + "/")
												 .then(function(response) {
																	 call(response.data);
															 });
								 };
						 });