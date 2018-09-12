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
    instance = new FirstApiSdk.AccessTokenResponse();
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

  describe('AccessTokenResponse', function() {
    it('should create an instance of AccessTokenResponse', function() {
      // uncomment below and update the code to test AccessTokenResponse
      //var instane = new FirstApiSdk.AccessTokenResponse();
      //expect(instance).to.be.a(FirstApiSdk.AccessTokenResponse);
    });

    it('should have the property accessToken (base name: "accessToken")', function() {
      // uncomment below and update the code to test the property accessToken
      //var instane = new FirstApiSdk.AccessTokenResponse();
      //expect(instance).to.be();
    });

    it('should have the property clientRequestId (base name: "clientRequestId")', function() {
      // uncomment below and update the code to test the property clientRequestId
      //var instane = new FirstApiSdk.AccessTokenResponse();
      //expect(instance).to.be();
    });

    it('should have the property apiTraceId (base name: "apiTraceId")', function() {
      // uncomment below and update the code to test the property apiTraceId
      //var instane = new FirstApiSdk.AccessTokenResponse();
      //expect(instance).to.be();
    });

    it('should have the property requestStatus (base name: "requestStatus")', function() {
      // uncomment below and update the code to test the property requestStatus
      //var instane = new FirstApiSdk.AccessTokenResponse();
      //expect(instance).to.be();
    });

  });

}));