var rest = angular.module('rest', []);

rest.service('Alerts', ['$rootScope', AlertsService]);
rest.service('User', ['$http', 'Alerts', UserService]);
rest.service('Lists', ['$http', 'Alerts', ListsService]);
rest.service('Recipes', ['$http', 'Alerts', RecipesService]);
