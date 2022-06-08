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
 * Collection not found
 * @export
 * @interface CollectionDeleteNotFoundResponseBody
 */
export interface CollectionDeleteNotFoundResponseBody {
    /**
     * Identifier of missing collection
     * @type {number}
     * @memberof CollectionDeleteNotFoundResponseBody
     */
    id: number;
    /**
     * Message of error
     * @type {string}
     * @memberof CollectionDeleteNotFoundResponseBody
     */
    message: string;
}

export function CollectionDeleteNotFoundResponseBodyFromJSON(json: any): CollectionDeleteNotFoundResponseBody {
    return CollectionDeleteNotFoundResponseBodyFromJSONTyped(json, false);
}

export function CollectionDeleteNotFoundResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): CollectionDeleteNotFoundResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'message': json['message'],
    };
}

export function CollectionDeleteNotFoundResponseBodyToJSON(value?: CollectionDeleteNotFoundResponseBody | null): any {
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
