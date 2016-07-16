var app = angular.module('BSDShiftApp');

app.controller('ResultController',['$scope', '$stateParams', 'ResultService', function($scope, $stateParams, ResultService){
  $scope.result = [];
  $scope.getResult = function(){
    $scope.result = ResultService.getResult();
  };

  $scope.getResultByDay = function(){
    $scope.result[0] = ResultService.getResultByDay($stateParams.day);
  };
}]);
