var app = angular.module('BSDShiftApp',['ui.router']);

app.config(function($stateProvider, $urlRouterProvider){
  $urlRouterProvider.otherwise("/");
  $stateProvider
    .state('result', {
      url : '/result',
      templateUrl : '/js/app/result/template/result.tpl.html'
    })
    .state('assign', {
      url : '/assign/:id',
      controller : 'AssignController',
      templateUrl : '/js/app/assign/template/assign.tpl.html'
    })
    .state('resultByDay', {
      url : '/result/:day',
      controller : 'ResultController',
      templateUrl : '/js/app/result/template/result_day.tpl.html'
    })
});

app.constant("AppConfig", {
    "api_url" : "http://bsd.olipicus.com:8080/api"
});
