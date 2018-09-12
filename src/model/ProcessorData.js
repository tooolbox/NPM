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


import ApiClient from '../ApiClient';





/**
* The ProcessorData model module.
* @module model/ProcessorData
* @version 6.3.0
*/
export default class ProcessorData {
    /**
    * Constructs a new <code>ProcessorData</code>.
    * Model for processor data
    * @alias module:model/ProcessorData
    * @class
    * @param responseCode {String} 
    * @param responseMessage {String} 
    */

    constructor(responseCode, responseMessage) {
        

        
        

        this['responseCode'] = responseCode;this['responseMessage'] = responseMessage;

        
    }

    /**
    * Constructs a <code>ProcessorData</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/ProcessorData} obj Optional instance to populate.
    * @return {module:model/ProcessorData} The populated <code>ProcessorData</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ProcessorData();

            
            
            

            if (data.hasOwnProperty('responseCode')) {
                obj['responseCode'] = ApiClient.convertToType(data['responseCode'], 'String');
            }
            if (data.hasOwnProperty('responseMessage')) {
                obj['responseMessage'] = ApiClient.convertToType(data['responseMessage'], 'String');
            }
            if (data.hasOwnProperty('approvalCode')) {
                obj['approvalCode'] = ApiClient.convertToType(data['approvalCode'], 'String');
            }
        }
        return obj;
    }

    /**
    * @member {String} responseCode
    */
    responseCode = undefined;
    /**
    * @member {String} responseMessage
    */
    responseMessage = undefined;
    /**
    * @member {String} approvalCode
    */
    approvalCode = undefined;








}

