/* tslint:disable */
/* eslint-disable */
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * Package not found
 * @export
 * @interface PackageShowNotFoundResponseBody
 */
export interface PackageShowNotFoundResponseBody {
    /**
     * Identifier of missing package
     * @type {number}
     * @memberof PackageShowNotFoundResponseBody
     */
    id: number;
    /**
     * Message of error
     * @type {string}
     * @memberof PackageShowNotFoundResponseBody
     */
    message: string;
}

export function PackageShowNotFoundResponseBodyFromJSON(json: any): PackageShowNotFoundResponseBody {
    return PackageShowNotFoundResponseBodyFromJSONTyped(json, false);
}

export function PackageShowNotFoundResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): PackageShowNotFoundResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'message': json['message'],
    };
}

export function PackageShowNotFoundResponseBodyToJSON(value?: PackageShowNotFoundResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'message': value.message,
    };
}
