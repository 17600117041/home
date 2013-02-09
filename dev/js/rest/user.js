/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this source
 code is governed by a BSD-style license that can be found in the
 LICENSE file.
 */

/* This is an interface to the RESTful User service. */
function UserService($http, Alerts) {
		this.get = function(call) {
				var promise = $http.get("/rest/user/");
				var error = {
						type: "warning",
						strong: "Warning!",
						message: "Unable to retrieve user information."
				};
				Alerts.handle(promise, error, undefined, call);
				
		};
}
