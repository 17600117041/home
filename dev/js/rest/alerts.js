/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this source
 code is governed by a BSD-style license that can be found in the
 LICENSE file.
 */

/* This is a service that stores alert messages that need to be
 * displayed.
 */
function AlertsService($rootScope) {
		this.alerts = [];
		var self = this;
		this.add = function (type, strong, message) {
				this.alerts.unshift({
						type: type,
						strong: strong,
						message: message
				});
				window.setTimeout(function() {
						$rootScope.$apply(self.remove(self.alerts.length - 1));
				}, 15000);
		};

		this.remove = function(index) {
				this.alerts.splice(index, 1);
		};

		this.handle = function(promise, error, success, scall, ecall) {
				promise
						.success(function (data, status, headers, config) {
								if (success != undefined) {
										self.add(success.type, success.strong, 
														 success.message);
								}

								if (scall != undefined) {
										scall(data, status, headers, config);
								}
						})
						.error(function (data, status, headers, config) {
								if (error != undefined) {
										self.add(error.type, error.strong, 
														 error.message);
								}
								
								if (ecall != undefined) {
										ecall(data, status, headers, config);
								}
						});
		};
}
