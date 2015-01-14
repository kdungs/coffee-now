/*jslint browser: true */
/*global angular */
'use strict';

var app = angular.module('coffeeNow', ['ionic']);

app.run(function ($ionicPlatform) {
  $ionicPlatform.ready(function () {
    if (window.cordova && window.cordova.plugins.Keyboard) {
      window.cordova.plugins.Keyboard.hideKeyboardAccessoryBar(true);
    }
    if (window.StatusBar) {
      window.StatusBar.styleDefault();
    }
  });
});

app.config(function ($stateProvider, $urlRouterProvider) {
  $urlRouterProvider.otherwise('/');
  $stateProvider
  .state('main', {
    url: '/',
    templateUrl: 'templates/main.html'
  })
  .state('request', {
    url: '/request',
    templateUrl: 'templates/request.html'
  })
  .state('settings', {
    url: '/settings',
    templateUrl: 'templates/settings.html'
  });
});
