'use strict';

/* jasmine specs for controllers go here */

describe('AlertsCtrl', function(){
		var scope;
		var ctrl;
		var alerts;

		beforeEach(module('rest'));
		beforeEach(inject(function($rootScope, $controller, Alerts) {
				scope = $rootScope.$new();
				alerts = Alerts;
				ctrl = $controller(AlertsCtrl, {
						$scope: scope,
						Alerts: Alerts
						});
		}));
		
		
		it('should start with no alerts', function() {
				expect(scope.alerts.length).toBe(0);
		});


		it("should keep it's alerts in sycn with the Alerts service." , 
			 function() {
					 var len = 0;
					 // Add a message.
					 var ms = [
							 {
									 type: "info",
									 strong: "Just Saying!",
									 message: "I thought you'd like to know this info."
							 },
							 {
									 type: "warning",
									 strong: "uh oh!",
									 message: "geez"
							 },
							 {
									 type: "error",
									 strong: "Yikes!",
									 message: "no!"
							 },
							 {
									 type: "success",
									 strong: "howdy!",
									 message: "you did it."
							 }
					 ];

					 for (var x = 0; x < ms.length; x++) {
							 scope.$apply(alerts.add(ms[x].type, ms[x].strong, ms[x].message));
							 len++;
							 
							 expect(scope.alerts.length).toBe(len);
							 expect(scope.alerts[0].type).toBe(ms[x].type);
							 expect(scope.alerts[0].strong).toBe(ms[x].strong);
							 expect(scope.alerts[0].message).toBe(ms[x].message);
					 }
			 });
});
