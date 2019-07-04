//SPDX-License-Identifier: Apache-2.0

var money = require('./controller.js');

module.exports = function(app){

  app.get('/get_money/:id', function(req, res){
    money.get_money(req, res);
  });
  app.get('/add_money/:money', function(req, res){
    money.add_money(req, res);
  });
  app.get('/get_all_money', function(req, res){
    money.get_all_money(req, res);
  });
  app.get('/change_holder/:holder', function(req, res){
    money.change_holder(req, res);
  });
}
