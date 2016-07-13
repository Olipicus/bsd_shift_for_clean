var app = angular.module('BSDShiftApp');

app.service('ResultService',['AppConfig', function(appConfig){
  this.getResult = function() {
    var transport = new Thrift.Transport(appConfig.api_url);
    var protocol  = new Thrift.Protocol(transport);
    var client = new MemberServiceClient(protocol);
    return client.getResults();
  };
}]);
