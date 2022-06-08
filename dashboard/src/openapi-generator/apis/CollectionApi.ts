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


import * as runtime from '../runtime';
import {
    CollectionBulkNotAvailableResponseBody,
    CollectionBulkNotAvailableResponseBodyFromJSON,
    CollectionBulkNotAvailableResponseBodyToJSON,
    CollectionBulkNotValidResponseBody,
    CollectionBulkNotValidResponseBodyFromJSON,
    CollectionBulkNotValidResponseBodyToJSON,
    CollectionBulkRequestBody,
    CollectionBulkRequestBodyFromJSON,
    CollectionBulkRequestBodyToJSON,
    CollectionBulkResponseBody,
    CollectionBulkResponseBodyFromJSON,
    CollectionBulkResponseBodyToJSON,
    CollectionBulkStatusResponseBody,
    CollectionBulkStatusResponseBodyFromJSON,
    CollectionBulkStatusResponseBodyToJSON,
    CollectionCancelNotFoundResponseBody,
    CollectionCancelNotFoundResponseBodyFromJSON,
    CollectionCancelNotFoundResponseBodyToJSON,
    CollectionCancelNotRunningResponseBody,
    CollectionCancelNotRunningResponseBodyFromJSON,
    CollectionCancelNotRunningResponseBodyToJSON,
    CollectionDeleteNotFoundResponseBody,
    CollectionDeleteNotFoundResponseBodyFromJSON,
    CollectionDeleteNotFoundResponseBodyToJSON,
    CollectionDownloadNotFoundResponseBody,
    CollectionDownloadNotFoundResponseBodyFromJSON,
    CollectionDownloadNotFoundResponseBodyToJSON,
    CollectionListResponseBody,
    CollectionListResponseBodyFromJSON,
    CollectionListResponseBodyToJSON,
    CollectionMonitorResponseBody,
    CollectionMonitorResponseBodyFromJSON,
    CollectionMonitorResponseBodyToJSON,
    CollectionPreservationActionsNotFoundResponseBody,
    CollectionPreservationActionsNotFoundResponseBodyFromJSON,
    CollectionPreservationActionsNotFoundResponseBodyToJSON,
    CollectionPreservationActionsResponseBody,
    CollectionPreservationActionsResponseBodyFromJSON,
    CollectionPreservationActionsResponseBodyToJSON,
    CollectionRetryNotFoundResponseBody,
    CollectionRetryNotFoundResponseBodyFromJSON,
    CollectionRetryNotFoundResponseBodyToJSON,
    CollectionRetryNotRunningResponseBody,
    CollectionRetryNotRunningResponseBodyFromJSON,
    CollectionRetryNotRunningResponseBodyToJSON,
    CollectionShowNotFoundResponseBody,
    CollectionShowNotFoundResponseBodyFromJSON,
    CollectionShowNotFoundResponseBodyToJSON,
    CollectionShowResponseBody,
    CollectionShowResponseBodyFromJSON,
    CollectionShowResponseBodyToJSON,
    CollectionWorkflowNotFoundResponseBody,
    CollectionWorkflowNotFoundResponseBodyFromJSON,
    CollectionWorkflowNotFoundResponseBodyToJSON,
    CollectionWorkflowResponseBody,
    CollectionWorkflowResponseBodyFromJSON,
    CollectionWorkflowResponseBodyToJSON,
} from '../models';

export interface CollectionBulkRequest {
    bulkRequestBody: CollectionBulkRequestBody;
}

export interface CollectionCancelRequest {
    id: number;
}

export interface CollectionDeleteRequest {
    id: number;
}

export interface CollectionDownloadRequest {
    id: number;
}

export interface CollectionListRequest {
    name?: string;
    aipId?: string;
    earliestCreatedTime?: Date;
    latestCreatedTime?: Date;
    status?: CollectionListStatusEnum;
    cursor?: string;
}

export interface CollectionPreservationActionsRequest {
    id: number;
}

export interface CollectionRetryRequest {
    id: number;
}

export interface CollectionShowRequest {
    id: number;
}

export interface CollectionWorkflowRequest {
    id: number;
}

/**
 * CollectionApi - interface
 * 
 * @export
 * @interface CollectionApiInterface
 */
export interface CollectionApiInterface {
    /**
     * Bulk operations (retry, cancel...).
     * @summary bulk collection
     * @param {CollectionBulkRequestBody} bulkRequestBody 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CollectionApiInterface
     */
    collectionBulkRaw(requestParameters: CollectionBulkRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<CollectionBulkResponseBody>>;

    /**
     * Bulk operations (retry, cancel...).
     * bulk collection
     */
    collectionBulk(requestParameters: CollectionBulkRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<CollectionBulkResponseBody>;

    /**
     * Retrieve status of current bulk operation.
     * @summary bulk_status collection
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CollectionApiInterface
     */
    collectionBulkStatusRaw(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<CollectionBulkStatusResponseBody>>;

    /**
     * Retrieve status of current bulk operation.
     * bulk_status collection
     */
    collectionBulkStatus(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<CollectionBulkStatusResponseBody>;

    /**
     * Cancel collection processing by ID
     * @summary cancel collection
     * @param {number} id Identifier of collection to remove
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CollectionApiInterface
     */
    collectionCancelRaw(requestParameters: CollectionCancelRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Cancel collection processing by ID
     * cancel collection
     */
    collectionCancel(requestParameters: CollectionCancelRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void>;

    /**
     * Delete collection by ID
     * @summary delete collection
     * @param {number} id Identifier of collection to delete
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CollectionApiInterface
     */
    collectionDeleteRaw(requestParameters: CollectionDeleteRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Delete collection by ID
     * delete collection
     */
    collectionDelete(requestParameters: CollectionDeleteRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void>;

    /**
     * Download collection by ID
     * @summary download collection
     * @param {number} id Identifier of collection to look up
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CollectionApiInterface
     */
    collectionDownloadRaw(requestParameters: CollectionDownloadRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<string>>;

    /**
     * Download collection by ID
     * download collection
     */
    collectionDownload(requestParameters: CollectionDownloadRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<string>;

    /**
     * List all stored collections
     * @summary list collection
     * @param {string} [name] 
     * @param {string} [aipId] 
     * @param {Date} [earliestCreatedTime] 
     * @param {Date} [latestCreatedTime] 
     * @param {'new' | 'in progress' | 'done' | 'error' | 'unknown' | 'queued' | 'pending' | 'abandoned'} [status] 
     * @param {string} [cursor] Pagination cursor
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CollectionApiInterface
     */
    collectionListRaw(requestParameters: CollectionListRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<CollectionListResponseBody>>;

    /**
     * List all stored collections
     * list collection
     */
    collectionList(requestParameters: CollectionListRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<CollectionListResponseBody>;

    /**
     * 
     * @summary monitor collection
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CollectionApiInterface
     */
    collectionMonitorRaw(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * monitor collection
     */
    collectionMonitor(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void>;

    /**
     * List all preservation actions by ID
     * @summary preservation-actions collection
     * @param {number} id Identifier of collection to look up
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CollectionApiInterface
     */
    collectionPreservationActionsRaw(requestParameters: CollectionPreservationActionsRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<CollectionPreservationActionsResponseBody>>;

    /**
     * List all preservation actions by ID
     * preservation-actions collection
     */
    collectionPreservationActions(requestParameters: CollectionPreservationActionsRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<CollectionPreservationActionsResponseBody>;

    /**
     * Retry collection processing by ID
     * @summary retry collection
     * @param {number} id Identifier of collection to retry
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CollectionApiInterface
     */
    collectionRetryRaw(requestParameters: CollectionRetryRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Retry collection processing by ID
     * retry collection
     */
    collectionRetry(requestParameters: CollectionRetryRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void>;

    /**
     * Show collection by ID
     * @summary show collection
     * @param {number} id Identifier of collection to show
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CollectionApiInterface
     */
    collectionShowRaw(requestParameters: CollectionShowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<CollectionShowResponseBody>>;

    /**
     * Show collection by ID
     * show collection
     */
    collectionShow(requestParameters: CollectionShowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<CollectionShowResponseBody>;

    /**
     * Retrieve workflow status by ID
     * @summary workflow collection
     * @param {number} id Identifier of collection to look up
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CollectionApiInterface
     */
    collectionWorkflowRaw(requestParameters: CollectionWorkflowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<CollectionWorkflowResponseBody>>;

    /**
     * Retrieve workflow status by ID
     * workflow collection
     */
    collectionWorkflow(requestParameters: CollectionWorkflowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<CollectionWorkflowResponseBody>;

}

/**
 * 
 */
export class CollectionApi extends runtime.BaseAPI implements CollectionApiInterface {

    /**
     * Bulk operations (retry, cancel...).
     * bulk collection
     */
    async collectionBulkRaw(requestParameters: CollectionBulkRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<CollectionBulkResponseBody>> {
        if (requestParameters.bulkRequestBody === null || requestParameters.bulkRequestBody === undefined) {
            throw new runtime.RequiredError('bulkRequestBody','Required parameter requestParameters.bulkRequestBody was null or undefined when calling collectionBulk.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/collection/bulk`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CollectionBulkRequestBodyToJSON(requestParameters.bulkRequestBody),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CollectionBulkResponseBodyFromJSON(jsonValue));
    }

    /**
     * Bulk operations (retry, cancel...).
     * bulk collection
     */
    async collectionBulk(requestParameters: CollectionBulkRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<CollectionBulkResponseBody> {
        const response = await this.collectionBulkRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieve status of current bulk operation.
     * bulk_status collection
     */
    async collectionBulkStatusRaw(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<CollectionBulkStatusResponseBody>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/bulk`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CollectionBulkStatusResponseBodyFromJSON(jsonValue));
    }

    /**
     * Retrieve status of current bulk operation.
     * bulk_status collection
     */
    async collectionBulkStatus(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<CollectionBulkStatusResponseBody> {
        const response = await this.collectionBulkStatusRaw(initOverrides);
        return await response.value();
    }

    /**
     * Cancel collection processing by ID
     * cancel collection
     */
    async collectionCancelRaw(requestParameters: CollectionCancelRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling collectionCancel.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/{id}/cancel`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Cancel collection processing by ID
     * cancel collection
     */
    async collectionCancel(requestParameters: CollectionCancelRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void> {
        await this.collectionCancelRaw(requestParameters, initOverrides);
    }

    /**
     * Delete collection by ID
     * delete collection
     */
    async collectionDeleteRaw(requestParameters: CollectionDeleteRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling collectionDelete.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Delete collection by ID
     * delete collection
     */
    async collectionDelete(requestParameters: CollectionDeleteRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void> {
        await this.collectionDeleteRaw(requestParameters, initOverrides);
    }

    /**
     * Download collection by ID
     * download collection
     */
    async collectionDownloadRaw(requestParameters: CollectionDownloadRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<string>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling collectionDownload.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/{id}/download`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.TextApiResponse(response) as any;
    }

    /**
     * Download collection by ID
     * download collection
     */
    async collectionDownload(requestParameters: CollectionDownloadRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<string> {
        const response = await this.collectionDownloadRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List all stored collections
     * list collection
     */
    async collectionListRaw(requestParameters: CollectionListRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<CollectionListResponseBody>> {
        const queryParameters: any = {};

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        if (requestParameters.aipId !== undefined) {
            queryParameters['aip_id'] = requestParameters.aipId;
        }

        if (requestParameters.earliestCreatedTime !== undefined) {
            queryParameters['earliest_created_time'] = (requestParameters.earliestCreatedTime as any).toISOString();
        }

        if (requestParameters.latestCreatedTime !== undefined) {
            queryParameters['latest_created_time'] = (requestParameters.latestCreatedTime as any).toISOString();
        }

        if (requestParameters.status !== undefined) {
            queryParameters['status'] = requestParameters.status;
        }

        if (requestParameters.cursor !== undefined) {
            queryParameters['cursor'] = requestParameters.cursor;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CollectionListResponseBodyFromJSON(jsonValue));
    }

    /**
     * List all stored collections
     * list collection
     */
    async collectionList(requestParameters: CollectionListRequest = {}, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<CollectionListResponseBody> {
        const response = await this.collectionListRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * monitor collection
     */
    async collectionMonitorRaw(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/monitor`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * monitor collection
     */
    async collectionMonitor(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void> {
        await this.collectionMonitorRaw(initOverrides);
    }

    /**
     * List all preservation actions by ID
     * preservation-actions collection
     */
    async collectionPreservationActionsRaw(requestParameters: CollectionPreservationActionsRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<CollectionPreservationActionsResponseBody>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling collectionPreservationActions.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/{id}/preservation-actions`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CollectionPreservationActionsResponseBodyFromJSON(jsonValue));
    }

    /**
     * List all preservation actions by ID
     * preservation-actions collection
     */
    async collectionPreservationActions(requestParameters: CollectionPreservationActionsRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<CollectionPreservationActionsResponseBody> {
        const response = await this.collectionPreservationActionsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retry collection processing by ID
     * retry collection
     */
    async collectionRetryRaw(requestParameters: CollectionRetryRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling collectionRetry.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/{id}/retry`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Retry collection processing by ID
     * retry collection
     */
    async collectionRetry(requestParameters: CollectionRetryRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void> {
        await this.collectionRetryRaw(requestParameters, initOverrides);
    }

    /**
     * Show collection by ID
     * show collection
     */
    async collectionShowRaw(requestParameters: CollectionShowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<CollectionShowResponseBody>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling collectionShow.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CollectionShowResponseBodyFromJSON(jsonValue));
    }

    /**
     * Show collection by ID
     * show collection
     */
    async collectionShow(requestParameters: CollectionShowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<CollectionShowResponseBody> {
        const response = await this.collectionShowRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieve workflow status by ID
     * workflow collection
     */
    async collectionWorkflowRaw(requestParameters: CollectionWorkflowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<CollectionWorkflowResponseBody>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling collectionWorkflow.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/{id}/workflow`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CollectionWorkflowResponseBodyFromJSON(jsonValue));
    }

    /**
     * Retrieve workflow status by ID
     * workflow collection
     */
    async collectionWorkflow(requestParameters: CollectionWorkflowRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<CollectionWorkflowResponseBody> {
        const response = await this.collectionWorkflowRaw(requestParameters, initOverrides);
        return await response.value();
    }

}

/**
 * @export
 */
export const CollectionListStatusEnum = {
    New: 'new',
    InProgress: 'in progress',
    Done: 'done',
    Error: 'error',
    Unknown: 'unknown',
    Queued: 'queued',
    Pending: 'pending',
    Abandoned: 'abandoned'
} as const;
export type CollectionListStatusEnum = typeof CollectionListStatusEnum[keyof typeof CollectionListStatusEnum];