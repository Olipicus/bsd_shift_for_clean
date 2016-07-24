var app = angular.module('BSDShiftApp');

app.controller('AssignController', ['$scope', '$state', '$stateParams', '$timeout', '$websocket', 'AppConfig', function($scope, $state, $stateParams, $timeout, $websocket, appConfig){
  var transport = new Thrift.Transport(appConfig.api_url);
  var protocol  = new Thrift.Protocol(transport);
  var client = new MemberServiceClient(protocol);

  var ws = $websocket(appConfig.ws_url);



  $scope.checkState = function(){
    try{
      $scope.member = client.getMember($stateParams.id);
    }
    catch(e){
      $scope.message = "เจ้าเป็นใครอ่ะ ข้าไม่รู้จัก";
    }
  };

  $scope.assignDay = function(){
    if($scope.member.day != '' && typeof($scope.member.day) != 'undefined') {
      $timeout(function(){ $scope.message = "คุณอยู่ในที่ ๆ ควรอยู่ อยู่แล้ว"; }, 800)
      .then($timeout(function(){$state.go('resultByDay',{"day" : $scope.member.day});},2000));
    }
    else {
      $timeout(function(){ $scope.message = "อืมมม"; }, 500)
      .then($timeout(function(){ $scope.message = "ให้ไปอยู่ไหนดีล่ะ"; }, 2000))
      .then($timeout(function(){ $scope.message = $scope.member.message; }, 3000))
      .then($timeout(function(){ $scope.message = "อ้าา รู้แล้ว"; }, 4000))
      .then($timeout(function(){ $scope.result = client.assignDay($stateParams.id); ws.send("update"); $state.go('resultByDay',{"day" : $scope.result[0].day}); }, 5000));
    }
  };
}]);
