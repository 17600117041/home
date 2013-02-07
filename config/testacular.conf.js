basePath = '../';

files = [
  JASMINE,
  JASMINE_ADAPTER,
  'test/lib/jquery.js',
  'test/lib/lib/angular/angular.js',
  'test/lib/lib/angular/angular-*.js',
  'test/lib/angular/angular-mocks.js',
  'dev/js/**/*.js',
  'test/unit/**/*.js'
];

autoWatch = true;

browsers = ['chromium'];

junitReporter = {
  outputFile: 'test_out/unit.xml',
  suite: 'unit'
};
