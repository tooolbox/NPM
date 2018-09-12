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
import Airline from './Airline';
import CarRental from './CarRental';
import Lodging from './Lodging';





/**
* The IndustrySpecificExtensions model module.
* @module model/IndustrySpecificExtensions
* @version 6.3.0
*/
export default class IndustrySpecificExtensions {
    /**
    * Constructs a new <code>IndustrySpecificExtensions</code>.
    * @alias module:model/IndustrySpecificExtensions
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>IndustrySpecificExtensions</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/IndustrySpecificExtensions} obj Optional instance to populate.
    * @return {module:model/IndustrySpecificExtensions} The populated <code>IndustrySpecificExtensions</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new IndustrySpecificExtensions();

            
            
            

            if (data.hasOwnProperty('airline')) {
                obj['airline'] = Airline.constructFromObject(data['airline']);
            }
            if (data.hasOwnProperty('lodging')) {
                obj['lodging'] = Lodging.constructFromObject(data['lodging']);
            }
            if (data.hasOwnProperty('carRental')) {
                obj['carRental'] = CarRental.constructFromObject(data['carRental']);
            }
        }
        return obj;
    }

    /**
    * @member {module:model/Airline} airline
    */
    airline = undefined;
    /**
    * @member {module:model/Lodging} lodging
    */
    lodging = undefined;
    /**
    * @member {module:model/CarRental} carRental
    */
    carRental = undefined;








}

