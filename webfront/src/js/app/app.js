var app = angular.module('BSDShiftApp',['ui.router']);

app.config(function($stateProvider, $urlRouterProvider){
  $urlRouterProvider.otherwise("/");
  $stateProvider
    .state('main', {
      url : '/',
      templateUrl : '/js/app/result/template/result.tpl.html'
    })
});
