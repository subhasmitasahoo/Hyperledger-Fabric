// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();

	$scope.getAllInvoices = function(){

		appFactory.getAllInvoices(function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			$scope.all_invoices = array;
		});
	}

	$scope.getInvoice = function(){

		var id = $scope.invoice_id;

		appFactory.getInvoice(id, function(data){
			$scope.one_invoice = data;
			console.log(data);
			if ($scope.query_invoice == "Could not locate Invoice"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.createInvoice = function(){

		appFactory.createInvoice($scope.invoice, function(data){
			$scope.create_invoice = data;
			$("#success_create").show();
		});
	}

});

// Angular Factory
app.factory('appFactory', function($http){

	var factory = {};

    factory.getAllInvoices = function(callback){

    	$http.get('/get_all_invoices/').success(function(output){
			callback(output)
		});
	}

	factory.getInvoice = function(id, callback){
    	$http.get('/get_invoice/'+id).success(function(output){
			callback(output)
		});
	}

	factory.createInvoice = function(data, callback){
		var invoice = data.id+"_"+data.sgstn+"_"+data.sstate+"_"+data.invoiceno+"_"+data.date+"_"+data.cgstn+"_"+data.cname+"_";
		invoice += data.billadd+"_"+data.shipadd+"_"+data.taxableamount+"_"+data.totalTax+"_"+data.invoicetotal;
    	$http.get('/add_invoice/'+invoice).success(function(output){
			callback(output)
		});
	}
	return factory;
});
