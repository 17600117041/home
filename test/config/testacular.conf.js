basePath = '../';

files = [
  JASMINE,
  JASMINE_ADAPTER,
  'lib/jquery.js',
  'lib/angular.js',
  'lib/angular-*.js',
  'lib/angular/angular-mocks.js',
  '../dev/js/**/*.js',
  'unit/**/*.js'
];

autoWatch = true;

browsers = ['chromium'];

junitReporter = {
  outputFile: 'test_out/unit.xml',
  suite: 'unit'
};
