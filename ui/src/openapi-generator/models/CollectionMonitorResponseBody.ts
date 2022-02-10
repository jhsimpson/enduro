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
import {
    EnduroStoredCollectionResponseBody,
    EnduroStoredCollectionResponseBodyFromJSON,
    EnduroStoredCollectionResponseBodyFromJSONTyped,
    EnduroStoredCollectionResponseBodyToJSON,
} from './';

/**
 * MonitorResponseBody result type (default view)
 * @export
 * @interface CollectionMonitorResponseBody
 */
export interface CollectionMonitorResponseBody {
    /**
     * Identifier of collection
     * @type {number}
     * @memberof CollectionMonitorResponseBody
     */
    id: number;
    /**
     * 
     * @type {EnduroStoredCollectionResponseBody}
     * @memberof CollectionMonitorResponseBody
     */
    item?: EnduroStoredCollectionResponseBody;
    /**
     * Type of the event
     * @type {string}
     * @memberof CollectionMonitorResponseBody
     */
    type: string;
}

export function CollectionMonitorResponseBodyFromJSON(json: any): CollectionMonitorResponseBody {
    return CollectionMonitorResponseBodyFromJSONTyped(json, false);
}

export function CollectionMonitorResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): CollectionMonitorResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'item': !exists(json, 'item') ? undefined : EnduroStoredCollectionResponseBodyFromJSON(json['item']),
        'type': json['type'],
    };
}

export function CollectionMonitorResponseBodyToJSON(value?: CollectionMonitorResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'item': EnduroStoredCollectionResponseBodyToJSON(value.item),
        'type': value.type,
    };
}

