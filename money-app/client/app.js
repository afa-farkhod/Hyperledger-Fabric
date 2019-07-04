// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();

	$scope.queryAllMoney = function(){

		appFactory.queryAllMoney(function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			$scope.all_money = array;
		});
	}

	$scope.queryMoney = function(){

		var id = $scope.money_id;

		appFactory.queryMoney(id, function(data){
			$scope.query_money = data;

			if ($scope.query_money == "Could not locate money"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.recordMoney = function(){

		appFactory.recordMoney($scope.money, function(data){
			$scope.create_money = data;
			$("#success_create").show();
		});
	}

	$scope.changeHolder = function(){

		appFactory.changeHolder($scope.holder, function(data){
			$scope.change_holder = data;
			if ($scope.change_holder == "Error: no money found"){
				$("#error_holder").show();
				$("#success_holder").hide();
			} else{
				$("#success_holder").show();
				$("#error_holder").hide();
			}
		});
	}

});

// Angular Factory
app.factory('appFactory', function($http){

	var factory = {};

    factory.queryAllMoney = function(callback){

    	$http.get('/get_all_money/').success(function(output){
			callback(output)
		});
	}

	factory.queryMoney = function(id, callback){
    	$http.get('/get_money/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordMoney = function(data, callback){

		data.serial = data.serial;

		var money = data.id + "-" + data.serial + "-" + data.date + "-" + data.holder + "-" + data.bank;

    	$http.get('/add_money/'+money).success(function(output){
			callback(output)
		});
	}

	factory.changeHolder = function(data, callback){

		var holder = data.id + "-" + data.name;

    	$http.get('/change_holder/'+holder).success(function(output){
			callback(output)
		});
	}

	return factory;
});
