var app = angular.module('BSDShiftApp');

app.service('ResultService',['AppConfig', function(appConfig){
  var transport = new Thrift.Transport(appConfig.api_url);
  var protocol  = new Thrift.Protocol(transport);
  var client = new MemberServiceClient(protocol);

  this.AppConfig = appConfig;

  this.getResult = function() {
    return client.getResults();
  };

  this.getResultByDay = function(day) {
    return client.getResultByDay(day);
  };

  this.getNotAssign = function() {
    try{
      return client.getNotAssign();
    }
    catch(e){
      console.log(e);
    }
  };
}]);
