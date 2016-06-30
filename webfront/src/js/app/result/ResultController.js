var app = angular.module('BSDShiftApp');

app.controller('ResultController',['$scope', function($scope){
  $scope.result = [
    {
      day : 'Monday',
      color : "yellow",
      member : [
        {
          name : 'Tong'
        },
        {
          name : 'Notkung'
        },
        {
          name : 'Chai'
        },
        {
          name : 'Kit'
        }
      ]
    },
    {
      day : 'Tuesday',
      color : "pink",
      member : [
        {
          name : 'Tong'
        },
        {
          name : 'Notkung'
        },
        {
          name : 'Chai'
        },
        {
          name : 'Kit'
        }
      ]
    },
    {
      day : 'Wednesday',
      color : 'green',
      member : [
        {
          name : 'Tong'
        },
        {
          name : 'Notkung'
        },
        {
          name : 'Chai'
        },
        {
          name : 'Kit'
        }
      ]
    },
    {
      day : 'Thursday',
      color : 'orange',
      member : [
        {
          name : 'Tong'
        },
        {
          name : 'Notkung'
        },
        {
          name : 'Chai'
        },
        {
          name : 'Kit'
        }
      ]
    },
    {
      day : 'Friday',
      color : 'blue',
      member : [
        {
          name : 'Tong'
        },
        {
          name : 'Notkung'
        },
        {
          name : 'Chai'
        },
        {
          name : 'Kit'
        }
      ]
    }
  ];


}]);
