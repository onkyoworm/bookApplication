var app = angular.module("bookTransfer", []);

app.controller("myCtrl", function($scope){
  $scope.users = ['Bob', 'Leroy'];
});

app.controller("sendRequest", function($scope, $http){
  $scope.sendRequestHTTP = function(){
    var enrollId = $scope.enrollId;
    var newBookName = $scope.newBookName;
    $http({
      url: "http://127.0.0.1:3000/login",
      method: "POST",
      data: {enrollId: "jim", enrollSecret: "6avZQLwcUe9b"}
    }).then(function(response){});
  }
});
