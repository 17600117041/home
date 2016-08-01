/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this source
 code is governed by a BSD-style license that can be found in the
 LICENSE file.
 */

/* LinksService is an interface to the RESTful Link service. We have
 * to use a service here because go and angular don't play nicely
 * together with trailing slashes.
 */
function LinksService($http, Alerts) {
		// Note: All of these functions accept callback
		// functions in which the resulting data from the
		// server is returned.

		// Create a new link.
		this.create = function(data, scall, ecall) {
				var promise = $http.post("/rest/link/", data);
				var error = {
						type: "error",
						strong: "Failed!",
						message: "Could not create a new link. Try again in a few minutes."
				};
				var success = {
						type: "success",
						strong: "Success!",
						message: "Your new link is ready to use."
				};
				Alerts.handle(promise, error, success, scall, ecall);
				
				return promise;
		};

		// Save an existing link.
		this.save = function(data, scall, ecall) {
				var promise = $http.put("/rest/link/" + data.Key + "/", data);
				var error = {
						type: "info",
						strong: "Unable to save!",
						message: "Could not save your link. Try again in a few minutes."
				};
				Alerts.handle(promise, error, undefined, scall, ecall);
				
				return promise;
		};

		// Delete the link with the given key.
		this.del = function(key, scall, ecall) {
				var promise = $http({
						method: 'DELETE', 
						url:"/rest/link/" + key + "/"}
													 );
				var error = {
						type: "error",
						strong: "Failed!",
						message: "Could not delete the link. Try again in a few minutes."
				};
				var success = {
						type: "success",
						strong: "Success!",
						message: "The link has been deleted."
				};
				Alerts.handle(promise, error, success, scall, ecall);
				
				return promise;
		};
		
		// Get all links.
		this.getall = function(scall, ecall) {
				var promise = $http.get("/rest/link/");
				var error = {
						type: "warning",
						strong: "Warning!",
						message: "Unable to retrieve links. Try again in a few minutes."
				};
				Alerts.handle(promise, error, undefined, scall, ecall);
				
				return promise;
		};

		// Get all links that have the given search as a tag.
		this.search = function(tag, scall, ecall) {
				tag = encodeURIComponent(tag);
				var promise = $http.get("/rest/link/?search=" + tag);
				var error = {
						type: "warning",
						strong: "Warning!",
						message: "Unable to retrieve links. Try again in a few minutes."
				};
				Alerts.handle(promise, error, undefined, scall, ecall);
				
				return promise;
		};

		// Get a specific link.
		this.get = function(key, scall, ecall) {
				var promise = $http.get("/rest/link/" + key + "/");
				var error = {
						type: "warning",
						strong: "Warning!",
						message: "Unable to retrieve the link. Try again in a few minutes."
				};
				Alerts.handle(promise, error, undefined, scall, ecall);
				
				return promise;
		};
}
