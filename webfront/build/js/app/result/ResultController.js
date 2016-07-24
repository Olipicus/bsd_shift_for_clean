var app = angular.module('BSDShiftApp');

app.controller('ResultController',['$scope', '$websocket', 'ResultService', function($scope, $websocket, ResultService){
  $scope.result = [];
  var ws = $websocket(ResultService.AppConfig.ws_url);

  ws.onMessage(function(message) {
    if(message.data == "update"){
      $scope.getResult();
    }
  });

  $scope.getResult = function(){
    $scope.result = ResultService.getResult();
    $scope.waitList = ResultService.getNotAssign();
    console.log($scope.waitList);
  };

}]);
