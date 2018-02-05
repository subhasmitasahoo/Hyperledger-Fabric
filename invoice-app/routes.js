//SPDX-License-Identifier: Apache-2.0

var invoice = require('./controller.js');

module.exports = function(app){

  app.get('/get_invoice/:id', function(req, res){
    invoice.get_invoice(req, res);
  });
  app.get('/add_invoice/:invoice', function(req, res){
    invoice.add_invoice(req, res);
  });
  app.get('/get_all_invoices', function(req, res){
    invoice.get_all_invoices(req, res);
  });
}
