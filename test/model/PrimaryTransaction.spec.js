/**
 * Payment Gateway API Specification
 * Payment Gateway API for payment processing. 
 *
 * OpenAPI spec version: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.FirstApiSdk);
  }
}(this, function(expect, FirstApiSdk) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new FirstApiSdk.PrimaryTransaction();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('PrimaryTransaction', function() {
    it('should create an instance of PrimaryTransaction', function() {
      // uncomment below and update the code to test PrimaryTransaction
      //var instane = new FirstApiSdk.PrimaryTransaction();
      //expect(instance).to.be.a(FirstApiSdk.PrimaryTransaction);
    });

    it('should have the property transactionType (base name: "transactionType")', function() {
      // uncomment below and update the code to test the property transactionType
      //var instane = new FirstApiSdk.PrimaryTransaction();
      //expect(instance).to.be();
    });

    it('should have the property storeId (base name: "storeId")', function() {
      // uncomment below and update the code to test the property storeId
      //var instane = new FirstApiSdk.PrimaryTransaction();
      //expect(instance).to.be();
    });

    it('should have the property clientTransactionId (base name: "clientTransactionId")', function() {
      // uncomment below and update the code to test the property clientTransactionId
      //var instane = new FirstApiSdk.PrimaryTransaction();
      //expect(instance).to.be();
    });

    it('should have the property amount (base name: "amount")', function() {
      // uncomment below and update the code to test the property amount
      //var instane = new FirstApiSdk.PrimaryTransaction();
      //expect(instance).to.be();
    });

    it('should have the property paymentMethod (base name: "paymentMethod")', function() {
      // uncomment below and update the code to test the property paymentMethod
      //var instane = new FirstApiSdk.PrimaryTransaction();
      //expect(instance).to.be();
    });

    it('should have the property order (base name: "order")', function() {
      // uncomment below and update the code to test the property order
      //var instane = new FirstApiSdk.PrimaryTransaction();
      //expect(instance).to.be();
    });

    it('should have the property basketItems (base name: "basketItems")', function() {
      // uncomment below and update the code to test the property basketItems
      //var instane = new FirstApiSdk.PrimaryTransaction();
      //expect(instance).to.be();
    });

    it('should have the property splitShipment (base name: "splitShipment")', function() {
      // uncomment below and update the code to test the property splitShipment
      //var instane = new FirstApiSdk.PrimaryTransaction();
      //expect(instance).to.be();
    });

    it('should have the property additionalDetails (base name: "additionalDetails")', function() {
      // uncomment below and update the code to test the property additionalDetails
      //var instane = new FirstApiSdk.PrimaryTransaction();
      //expect(instance).to.be();
    });

    it('should have the property industrySpecificExtensions (base name: "industrySpecificExtensions")', function() {
      // uncomment below and update the code to test the property industrySpecificExtensions
      //var instane = new FirstApiSdk.PrimaryTransaction();
      //expect(instance).to.be();
    });

    it('should have the property storedCredentials (base name: "storedCredentials")', function() {
      // uncomment below and update the code to test the property storedCredentials
      //var instane = new FirstApiSdk.PrimaryTransaction();
      //expect(instance).to.be();
    });

  });

}));