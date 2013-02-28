// rest is the module for all of our RESTful services.
var rest = angular.module('rest', []);

// Add all the services.
rest.service('Alerts', ['$rootScope', AlertsService]);
rest.service('User', ['$http', 'Alerts', UserService]);
rest.service('Lists', ['$http', 'Alerts', ListsService]);
rest.service('Links', ['$http', 'Alerts', LinksService]);
rest.service('Recipes', ['$http', 'Alerts', RecipesService]);
