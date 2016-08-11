var app = angular.module('BSDShiftApp',['ui.router','ngWebSocket']);

app.config(function($stateProvider, $urlRouterProvider){
  $urlRouterProvider.otherwise("/");
  $stateProvider
    .state('result', {
      url : '/result',
      templateUrl : './js/app/result/template/result.tpl.html'
    })
    .state('assign', {
      url : '/assign/:id',
      controller : 'AssignController',
      templateUrl : './js/app/assign/template/assign.tpl.html'
    })
    .state('resultByDay', {
      url : '/result/:day',
      controller : 'ResultByDayController',
      templateUrl : './js/app/result/template/result_day.tpl.html'
    })
});

app.constant("AppConfig", {
    "api_url" : "http://127.0.0.1:8801/api",
    "ws_url" : "ws://127.0.0.1:8801/ws"
});
