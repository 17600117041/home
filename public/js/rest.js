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
																	 call(response.data);
															 });
								 };

								 this.save = function(data, call) {
										 return $http.put("/rest/list/" + data.Key + "/", data)
												 .then(function (response) {
																	 call(response.data);
															 });
								 };

								 this.checkupdate = function(data, call) {
										 return $http.get("/rest/list/" + data.Key + "/")
												 .then(function (response) {
																	 var sdate = Date.parseRFC3339(response.data.LastModified);
																	 var mdate = Date.parseRFC3339(data.LastModified);

																	 call(sdate > mdate);
															 });
										 
								 };

								 this.delete = function(key, call) {
										 return $http.delete("/rest/list/" + key + "/")
												 .then(function(response) {
																	 call(response.data);
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