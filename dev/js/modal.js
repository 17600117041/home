/*
 Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
 source code is governed by a BSD-style license that can be found in
 the LICENSE file.
 */

// This is used to auto-focus the items then they are switched
// from a span to a text box. 
function Modal() {
		return {
				restrict: 'A',
				transclude: true,
				scope: {
						id: '@modalId',
						title: '@modalTitle',
						message: '@modalMessage',
						data: "=modalData",
						form: "=modalForm",
						okbtncls: '@modalOkButtonClass',
						okbtntxt: '@modalOkButtonText',
						okbtnico: '@modalOkButtonIcon',
						onok: '&modalOnOk'

				},
				templateUrl: 'modal.html',
				link: function (scope, element, attrs) {

				}
				
		};
}
