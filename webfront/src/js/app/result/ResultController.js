var app = angular.module('BSDShiftApp');

app.controller('ResultController',['$scope', 'ResultService', function($scope, ResultService){
  $scope.result = [];
  $scope.getResult = function(){
    $scope.result = ResultService.getResult();
  };
}]);
