var app = angular.module('BSDShiftApp');

app.controller('AssignController', ['$scope', '$state', '$stateParams', 'AppConfig', function($scope, $state, $stateParams, appConfig){
  var transport = new Thrift.Transport(appConfig.api_url);
  var protocol  = new Thrift.Protocol(transport);
  var client = new MemberServiceClient(protocol);

  $scope.checkState = function(){
    var objMember = client.getMember($stateParams.id);
    console.log(objMember);
    if(objMember.day != '' && typeof(objMember.day) != 'undifined') {
      $scope.result = client.assignDay($stateParams.id);
    }
  };

  $scope.assignDay = function(){
    $scope.result = client.assignDay($stateParams.id);
  };
}]);
