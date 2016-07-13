var app = angular.module('BSDShiftApp',['ui.router']);

app.config(function($stateProvider, $urlRouterProvider){
  $urlRouterProvider.otherwise("/");
  $stateProvider
    .state('main', {
      url : '/',
      templateUrl : '/js/app/result/template/result.tpl.html'
    })
    .state('assign', {
      url : '/assign/:id',
      controller : 'AssignController',
      templateUrl : '/js/app/assign/template/assign.tpl.html'
    });
});

app.constant("AppConfig", {
    "api_url" : "http://127.0.0.1:8081/"
});
