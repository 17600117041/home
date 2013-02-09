/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this source
 code is governed by a BSD-style license that can be found in the
 LICENSE file.
 */

/* This is an interface to the RESTful Recipe service. We have 
 * to use a service here because go and angular don't play nicely 
 * together with trailing slashes. 
 */
function RecipesService($http, Alerts) {
		// Note: All of these functions accept callback
		// functions in which the resulting data from the
		// server is returned.

		// Create a new recipe.
		this.create = function(data, scall, ecall) {
				var promise = $http.post("/rest/recipe/", data);
				var error = {
						type: "error",
						strong: "Failed!",
						message: "Could not create a new recipe. Try again in a few minutes."
				};
				var success = {
						type: "success",
						strong: "Success!",
						message: "Your new recipe is ready to use."
				};
				Alerts.handle(promise, error, success, scall, ecall);
				
				return promise;
		};

		// Save an existing recipe.
		this.save = function(data, scall, ecall) {
				var promise = $http.put("/rest/recipe/" + data.Key + "/", data);
				var error = {
						type: "info",
						strong: "Unable to save!",
						message: "Could not save your recipe. Try again in a few minutes."
				};
				Alerts.handle(promise, error, undefined, scall, ecall);
				
				return promise;
		};

		// Delete the recipe with the given key.
		this.del = function(key, scall, ecall) {
				var promise = $http({
						method: 'DELETE', 
						url:"/rest/recipe/" + key + "/"}
													 );
				var error = {
						type: "error",
						strong: "Failed!",
						message: "Could not delete the recipe. Try again in a few minutes."
				};
				var success = {
						type: "success",
						strong: "Success!",
						message: "The recipe has been deleted."
				};
				Alerts.handle(promise, error, success, scall, ecall);
				
				return promise;
		};
		
		// Get all recipes.
		this.getall = function(scall, ecall) {
				var promise = $http.get("/rest/recipe/");
				var error = {
						type: "warning",
						strong: "Warning!",
						message: "Unable to retrieve recipes. Try again in a few minutes."
				};
				Alerts.handle(promise, error, undefined, scall, ecall);
				
				return promise;
		};


		// Get a specific recipe.
		this.get = function(key, scall, ecall) {
				var promise = $http.get("/rest/recipe/" + key + "/");
				var error = {
						type: "warning",
						strong: "Warning!",
						message: "Unable to retrieve the recipe. Try again in a few minutes."
				};
				Alerts.handle(promise, error, undefined, scall, ecall);
				
				return promise;
		};
}
