var app = angular.module('BSDShiftApp');

app.controller('AssignController', ['$scope', '$state', '$stateParams', '$timeout', 'AppConfig', function($scope, $state, $stateParams, $timeout, appConfig){
  var transport = new Thrift.Transport(appConfig.api_url);
  var protocol  = new Thrift.Protocol(transport);
  var client = new MemberServiceClient(protocol);

  $scope.checkState = function(){
    $scope.member = client.getMember($stateParams.id);
    console.log(JSON.stringify($scope.member));
    if($scope.member.day != '' && typeof($scope.member.day) != 'undifined') {
      $scope.result = client.assignDay($stateParams.id);
      $scope.message =$scope.result[0].day;
    }
  };

  $scope.assignDay = function(){
    var setMessage = function(message){ $scope.message = message; };

    $timeout(function(){ $scope.message = "อืมมม"; }, 500)
    .then($timeout(function(){ $scope.message = "ให้ไปอยู่ไหนดีล่ะ"; }, 2000))
    .then($timeout(function(){ $scope.message = "ชั่ว ๆ แบบนี้"; }, 3000))
    .then($timeout(function(){ $scope.message = "อ้าา รู้แล้ว"; }, 4000))
    .then($timeout(function(){ $scope.result = client.assignDay($stateParams.id); $state.go('resultByDay',{"day" : $scope.result[0].day}); }, 5000));


  };
}]);
