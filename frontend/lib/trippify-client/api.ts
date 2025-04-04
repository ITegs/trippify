/* tslint:disable */
/* eslint-disable */
/**
 * trippify API
 * trippify API
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from './configuration';
import type { AxiosPromise, AxiosInstance, RawAxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
import type { RequestArgs } from './base';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, BaseAPI, RequiredError, operationServerMap } from './base';

/**
 * 
 * @export
 * @enum {string}
 */

export const ErrorCode = {
    ServerError: 'server_error'
} as const;

export type ErrorCode = typeof ErrorCode[keyof typeof ErrorCode];


/**
 * 
 * @export
 * @interface Image
 */
export interface Image {
    /**
     * 
     * @type {string}
     * @memberof Image
     */
    'source': string;
    /**
     * 
     * @type {number}
     * @memberof Image
     */
    'timestamp'?: number;
}
/**
 * 
 * @export
 * @interface LoginUser
 */
export interface LoginUser {
    /**
     * 
     * @type {string}
     * @memberof LoginUser
     */
    'username': string;
    /**
     * 
     * @type {string}
     * @memberof LoginUser
     */
    'password': string;
}
/**
 * 
 * @export
 * @interface ModelError
 */
export interface ModelError {
    /**
     * A human readable error message
     * @type {string}
     * @memberof ModelError
     */
    'message': string;
    /**
     * 
     * @type {ErrorCode}
     * @memberof ModelError
     */
    'code': ErrorCode;
}


/**
 * 
 * @export
 * @interface NewSpot
 */
export interface NewSpot {
    /**
     * 
     * @type {string}
     * @memberof NewSpot
     */
    'title': string;
    /**
     * 
     * @type {number}
     * @memberof NewSpot
     */
    'latitude': number;
    /**
     * 
     * @type {number}
     * @memberof NewSpot
     */
    'longitude': number;
    /**
     * 
     * @type {string}
     * @memberof NewSpot
     */
    'date_from': string;
    /**
     * 
     * @type {string}
     * @memberof NewSpot
     */
    'description'?: string;
    /**
     * 
     * @type {Array<Image>}
     * @memberof NewSpot
     */
    'images': Array<Image>;
}
/**
 * 
 * @export
 * @interface NewTrip
 */
export interface NewTrip {
    /**
     * 
     * @type {string}
     * @memberof NewTrip
     */
    'owner_username'?: string;
    /**
     * 
     * @type {string}
     * @memberof NewTrip
     */
    'title': string;
}
/**
 * 
 * @export
 * @interface NewUser
 */
export interface NewUser {
    /**
     * 
     * @type {string}
     * @memberof NewUser
     */
    'name'?: string;
    /**
     * 
     * @type {string}
     * @memberof NewUser
     */
    'profile_pic'?: string;
    /**
     * 
     * @type {string}
     * @memberof NewUser
     */
    'username': string;
    /**
     * 
     * @type {string}
     * @memberof NewUser
     */
    'password': string;
}
/**
 * 
 * @export
 * @interface Spot
 */
export interface Spot {
    /**
     * A MongoDB ObjectId
     * @type {string}
     * @memberof Spot
     */
    '_id': string;
    /**
     * 
     * @type {string}
     * @memberof Spot
     */
    'date_to'?: string;
    /**
     * 
     * @type {string}
     * @memberof Spot
     */
    'title': string;
    /**
     * 
     * @type {number}
     * @memberof Spot
     */
    'latitude': number;
    /**
     * 
     * @type {number}
     * @memberof Spot
     */
    'longitude': number;
    /**
     * 
     * @type {string}
     * @memberof Spot
     */
    'date_from': string;
    /**
     * 
     * @type {string}
     * @memberof Spot
     */
    'description'?: string;
    /**
     * 
     * @type {Array<Image>}
     * @memberof Spot
     */
    'images': Array<Image>;
}
/**
 * 
 * @export
 * @interface Trip
 */
export interface Trip {
    /**
     * A MongoDB ObjectId
     * @type {string}
     * @memberof Trip
     */
    '_id': string;
    /**
     * 
     * @type {User}
     * @memberof Trip
     */
    'owner'?: User;
    /**
     * 
     * @type {string}
     * @memberof Trip
     */
    'start_date': string;
    /**
     * 
     * @type {string}
     * @memberof Trip
     */
    'end_date'?: string;
    /**
     * 
     * @type {Array<TripSpot>}
     * @memberof Trip
     */
    'spots'?: Array<TripSpot>;
    /**
     * 
     * @type {string}
     * @memberof Trip
     */
    'owner_username'?: string;
    /**
     * 
     * @type {string}
     * @memberof Trip
     */
    'title': string;
}
/**
 * 
 * @export
 * @interface TripSpot
 */
export interface TripSpot {
    /**
     * A MongoDB ObjectId
     * @type {string}
     * @memberof TripSpot
     */
    'spotId': string;
    /**
     * 
     * @type {number}
     * @memberof TripSpot
     */
    'latitude': number;
    /**
     * 
     * @type {number}
     * @memberof TripSpot
     */
    'longitude': number;
}
/**
 * 
 * @export
 * @interface User
 */
export interface User {
    /**
     * A MongoDB ObjectId
     * @type {string}
     * @memberof User
     */
    '_id': string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    'username'?: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    'name'?: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    'profile_pic'?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof User
     */
    'trips'?: Array<string>;
}

/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Add a spot to a given trip
         * @param {string} tripId The id of the trip
         * @param {NewSpot} newSpot 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        addSpotToTripPost: async (tripId: string, newSpot: NewSpot, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'tripId' is not null or undefined
            assertParamExists('addSpotToTripPost', 'tripId', tripId)
            // verify required parameter 'newSpot' is not null or undefined
            assertParamExists('addSpotToTripPost', 'newSpot', newSpot)
            const localVarPath = `/api/trip/{tripId}/spot/add`
                .replace(`{${"tripId"}}`, encodeURIComponent(String(tripId)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication bearerAuth required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(newSpot, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Get the firstTrip
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        firstTripGet: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/firstTrip`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Login
         * @param {LoginUser} loginUser 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        loginPost: async (loginUser: LoginUser, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'loginUser' is not null or undefined
            assertParamExists('loginPost', 'loginUser', loginUser)
            const localVarPath = `/api/login`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(loginUser, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Create a new trip
         * @param {NewTrip} newTrip 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        newTripPost: async (newTrip: NewTrip, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'newTrip' is not null or undefined
            assertParamExists('newTripPost', 'newTrip', newTrip)
            const localVarPath = `/api/newTrip`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication bearerAuth required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(newTrip, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Create a new user
         * @param {NewUser} newUser 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        newUserPost: async (newUser: NewUser, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'newUser' is not null or undefined
            assertParamExists('newUserPost', 'newUser', newUser)
            const localVarPath = `/api/newUser`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(newUser, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Get the spot with the given id
         * @param {string} spotId The id of the spot
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        spotGet: async (spotId: string, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'spotId' is not null or undefined
            assertParamExists('spotGet', 'spotId', spotId)
            const localVarPath = `/api/spot/{spotId}`
                .replace(`{${"spotId"}}`, encodeURIComponent(String(spotId)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Get the trip with the given id
         * @param {string} tripId The id of the trip
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        tripGet: async (tripId: string, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'tripId' is not null or undefined
            assertParamExists('tripGet', 'tripId', tripId)
            const localVarPath = `/api/trip/{tripId}`
                .replace(`{${"tripId"}}`, encodeURIComponent(String(tripId)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Get the user with the given username
         * @param {string} username The username of the user
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        userGet: async (username: string, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'username' is not null or undefined
            assertParamExists('userGet', 'username', username)
            const localVarPath = `/api/user/{username}`
                .replace(`{${"username"}}`, encodeURIComponent(String(username)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * DefaultApi - functional programming interface
 * @export
 */
export const DefaultApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = DefaultApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @summary Add a spot to a given trip
         * @param {string} tripId The id of the trip
         * @param {NewSpot} newSpot 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async addSpotToTripPost(tripId: string, newSpot: NewSpot, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Spot>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.addSpotToTripPost(tripId, newSpot, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.addSpotToTripPost']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Get the firstTrip
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async firstTripGet(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Trip>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.firstTripGet(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.firstTripGet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Login
         * @param {LoginUser} loginUser 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async loginPost(loginUser: LoginUser, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.loginPost(loginUser, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.loginPost']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Create a new trip
         * @param {NewTrip} newTrip 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async newTripPost(newTrip: NewTrip, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Trip>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.newTripPost(newTrip, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.newTripPost']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Create a new user
         * @param {NewUser} newUser 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async newUserPost(newUser: NewUser, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<User>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.newUserPost(newUser, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.newUserPost']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Get the spot with the given id
         * @param {string} spotId The id of the spot
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async spotGet(spotId: string, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Spot>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.spotGet(spotId, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.spotGet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Get the trip with the given id
         * @param {string} tripId The id of the trip
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async tripGet(tripId: string, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Trip>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.tripGet(tripId, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.tripGet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Get the user with the given username
         * @param {string} username The username of the user
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async userGet(username: string, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<User>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.userGet(username, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.userGet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
    }
};

/**
 * DefaultApi - factory interface
 * @export
 */
export const DefaultApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = DefaultApiFp(configuration)
    return {
        /**
         * 
         * @summary Add a spot to a given trip
         * @param {string} tripId The id of the trip
         * @param {NewSpot} newSpot 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        addSpotToTripPost(tripId: string, newSpot: NewSpot, options?: RawAxiosRequestConfig): AxiosPromise<Spot> {
            return localVarFp.addSpotToTripPost(tripId, newSpot, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get the firstTrip
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        firstTripGet(options?: RawAxiosRequestConfig): AxiosPromise<Trip> {
            return localVarFp.firstTripGet(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Login
         * @param {LoginUser} loginUser 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        loginPost(loginUser: LoginUser, options?: RawAxiosRequestConfig): AxiosPromise<string> {
            return localVarFp.loginPost(loginUser, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Create a new trip
         * @param {NewTrip} newTrip 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        newTripPost(newTrip: NewTrip, options?: RawAxiosRequestConfig): AxiosPromise<Trip> {
            return localVarFp.newTripPost(newTrip, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Create a new user
         * @param {NewUser} newUser 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        newUserPost(newUser: NewUser, options?: RawAxiosRequestConfig): AxiosPromise<User> {
            return localVarFp.newUserPost(newUser, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get the spot with the given id
         * @param {string} spotId The id of the spot
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        spotGet(spotId: string, options?: RawAxiosRequestConfig): AxiosPromise<Spot> {
            return localVarFp.spotGet(spotId, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get the trip with the given id
         * @param {string} tripId The id of the trip
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        tripGet(tripId: string, options?: RawAxiosRequestConfig): AxiosPromise<Trip> {
            return localVarFp.tripGet(tripId, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get the user with the given username
         * @param {string} username The username of the user
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        userGet(username: string, options?: RawAxiosRequestConfig): AxiosPromise<User> {
            return localVarFp.userGet(username, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * DefaultApi - object-oriented interface
 * @export
 * @class DefaultApi
 * @extends {BaseAPI}
 */
export class DefaultApi extends BaseAPI {
    /**
     * 
     * @summary Add a spot to a given trip
     * @param {string} tripId The id of the trip
     * @param {NewSpot} newSpot 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public addSpotToTripPost(tripId: string, newSpot: NewSpot, options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).addSpotToTripPost(tripId, newSpot, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get the firstTrip
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public firstTripGet(options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).firstTripGet(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Login
     * @param {LoginUser} loginUser 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public loginPost(loginUser: LoginUser, options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).loginPost(loginUser, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Create a new trip
     * @param {NewTrip} newTrip 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public newTripPost(newTrip: NewTrip, options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).newTripPost(newTrip, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Create a new user
     * @param {NewUser} newUser 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public newUserPost(newUser: NewUser, options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).newUserPost(newUser, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get the spot with the given id
     * @param {string} spotId The id of the spot
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public spotGet(spotId: string, options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).spotGet(spotId, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get the trip with the given id
     * @param {string} tripId The id of the trip
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public tripGet(tripId: string, options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).tripGet(tripId, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get the user with the given username
     * @param {string} username The username of the user
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public userGet(username: string, options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).userGet(username, options).then((request) => request(this.axios, this.basePath));
    }
}



