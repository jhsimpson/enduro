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
} from './EnduroStoredCollectionResponseBody';

/**
 * EnduroStored-CollectionCollectionResponseBody is the result type for an array of EnduroStored-CollectionResponseBody (default view)
 * @export
 * @interface EnduroStoredCollectionResponseBodyCollection
 */
export interface EnduroStoredCollectionResponseBodyCollection extends Array<EnduroStoredCollectionResponseBody> {
}

export function EnduroStoredCollectionResponseBodyCollectionFromJSON(json: any): EnduroStoredCollectionResponseBodyCollection {
    return EnduroStoredCollectionResponseBodyCollectionFromJSONTyped(json, false);
}

export function EnduroStoredCollectionResponseBodyCollectionFromJSONTyped(json: any, ignoreDiscriminator: boolean): EnduroStoredCollectionResponseBodyCollection {
    return json;
}

export function EnduroStoredCollectionResponseBodyCollectionToJSON(value?: EnduroStoredCollectionResponseBodyCollection | null): any {
    return value;
}
