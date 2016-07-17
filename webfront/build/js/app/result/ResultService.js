var app = angular.module('BSDShiftApp');

app.service('ResultService',['AppConfig', function(appConfig){
  var transport = new Thrift.Transport(appConfig.api_url);
  var protocol  = new Thrift.Protocol(transport);
  var client = new MemberServiceClient(protocol);
  this.getResult = function() {
    return client.getResults();
  };

  this.getResultByDay = function(day) {
    console.log(client.getResultByDay(day));
    return client.getResultByDay(day);
  };
}]);
