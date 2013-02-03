// ListsCtrl is the controller for viewing all lists.
function ListsCtrl($scope, $location, List) {
		$scope.view = function(key) {
				$location.path('/view/' + key + '/');
		};

		$scope.copymodal = function(key, name, event) {
				$scope.copyKey = key;
				$scope.copyName = "Copy of " + name;

				$('#copyModal').modal();

				// Don't let it fall through to the view.
				event.stopPropagation();
		};


		$scope.copy = function() {
				$('#copyModal').modal('hide');

				List.get($scope.copyKey, function(l) {
										 l.Name = $scope.copyName;
										 List.create(l, function(nl) {
																	 $location.path('/view/' + nl.Key + '/');
															 });
								 });

		};

		$scope.delete = function(key, event) {
				$scope.deleteKey = key;

				$('#deleteModal').modal();

				// Don't let it fall through to the view.
				event.stopPropagation();
		};

		$scope.sure = function() {
				$('#deleteModal').modal('hide');
				List.delete($scope.deleteKey, function() {
												List.getall(function (lists) {
																				$scope.lists = lists;
																		});
										});
		};

		$scope.new = function() {
				$('#newModal').modal();

		};

		$scope.back = function() {
				$('#newModal').modal('hide');
				$('#copyModal').modal('hide');
		};

		$scope.save = function() {
				$('#newModal').modal('hide');
				List.create({"Name": $scope.name}, function (l) {
											 $location.path('/view/' + l.Key + '/');
									 });
		};

		List.getall(function (lists) {
										$scope.lists = lists;
								});

}
