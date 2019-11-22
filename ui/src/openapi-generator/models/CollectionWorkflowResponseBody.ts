// tslint:disable
// eslint-disable
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
    EnduroCollectionWorkflowHistoryResponseBodyCollection,
    EnduroCollectionWorkflowHistoryResponseBodyCollectionFromJSON,
    EnduroCollectionWorkflowHistoryResponseBodyCollectionFromJSONTyped,
    EnduroCollectionWorkflowHistoryResponseBodyCollectionToJSON,
} from './';

/**
 * WorkflowResponseBody result type (default view)
 * @export
 * @interface CollectionWorkflowResponseBody
 */
export interface CollectionWorkflowResponseBody {
    /**
     * 
     * @type {EnduroCollectionWorkflowHistoryResponseBodyCollection}
     * @memberof CollectionWorkflowResponseBody
     */
    history?: EnduroCollectionWorkflowHistoryResponseBodyCollection;
    /**
     * 
     * @type {string}
     * @memberof CollectionWorkflowResponseBody
     */
    status?: string;
}

export function CollectionWorkflowResponseBodyFromJSON(json: any): CollectionWorkflowResponseBody {
    return CollectionWorkflowResponseBodyFromJSONTyped(json, false);
}

export function CollectionWorkflowResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): CollectionWorkflowResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'history': !exists(json, 'history') ? undefined : EnduroCollectionWorkflowHistoryResponseBodyCollectionFromJSON(json['history']),
        'status': !exists(json, 'status') ? undefined : json['status'],
    };
}

export function CollectionWorkflowResponseBodyToJSON(value?: CollectionWorkflowResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'history': EnduroCollectionWorkflowHistoryResponseBodyCollectionToJSON(value.history),
        'status': value.status,
    };
}

