var app = angular.module('BSDShiftApp');

app.controller('ResultByDayController',['$scope', '$stateParams', '$websocket', 'ResultService', function($scope, $stateParams, $websocket, ResultService){
  $scope.result = [];

  var ws = $websocket(ResultService.AppConfig.ws_url);

  ws.onMessage(function(message) {
    if(message.data == "update"){
      $scope.getResultByDay();
    }
  });

  $scope.getResultByDay = function(){
    $scope.result[0] = ResultService.getResultByDay($stateParams.day);
  };
}]);
