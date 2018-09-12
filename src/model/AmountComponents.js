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
* The AmountComponents model module.
* @module model/AmountComponents
* @version 6.3.0
*/
export default class AmountComponents {
    /**
    * Constructs a new <code>AmountComponents</code>.
    * @alias module:model/AmountComponents
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>AmountComponents</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/AmountComponents} obj Optional instance to populate.
    * @return {module:model/AmountComponents} The populated <code>AmountComponents</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AmountComponents();

            
            
            

            if (data.hasOwnProperty('subtotal')) {
                obj['subtotal'] = ApiClient.convertToType(data['subtotal'], 'Number');
            }
            if (data.hasOwnProperty('vatAmount')) {
                obj['vatAmount'] = ApiClient.convertToType(data['vatAmount'], 'Number');
            }
            if (data.hasOwnProperty('localTax')) {
                obj['localTax'] = ApiClient.convertToType(data['localTax'], 'Number');
            }
            if (data.hasOwnProperty('shipping')) {
                obj['shipping'] = ApiClient.convertToType(data['shipping'], 'Number');
            }
            if (data.hasOwnProperty('cashback')) {
                obj['cashback'] = ApiClient.convertToType(data['cashback'], 'Number');
            }
            if (data.hasOwnProperty('tip')) {
                obj['tip'] = ApiClient.convertToType(data['tip'], 'Number');
            }
            if (data.hasOwnProperty('convenienceFee')) {
                obj['convenienceFee'] = ApiClient.convertToType(data['convenienceFee'], 'Number');
            }
        }
        return obj;
    }

    /**
    * @member {Number} subtotal
    */
    subtotal = undefined;
    /**
    * @member {Number} vatAmount
    */
    vatAmount = undefined;
    /**
    * @member {Number} localTax
    */
    localTax = undefined;
    /**
    * @member {Number} shipping
    */
    shipping = undefined;
    /**
    * @member {Number} cashback
    */
    cashback = undefined;
    /**
    * @member {Number} tip
    */
    tip = undefined;
    /**
    * @member {Number} convenienceFee
    */
    convenienceFee = undefined;








}

