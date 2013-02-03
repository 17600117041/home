// UserCtrl is the controller for part of the site that lists all of
// the lists.
function UserCtrl($scope, User) {
		$scope.user = User.get();
}
