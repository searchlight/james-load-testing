/*
Apache JAMES Web Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// CassandraSchemaUpgradeAPIService CassandraSchemaUpgradeAPI service
type CassandraSchemaUpgradeAPIService service

type ApiGetLatestAvailableSchemaVersionRequest struct {
	ctx context.Context
	ApiService *CassandraSchemaUpgradeAPIService
}

func (r ApiGetLatestAvailableSchemaVersionRequest) Execute() (*GetLatestAvailableSchemaVersion200Response, *http.Response, error) {
	return r.ApiService.GetLatestAvailableSchemaVersionExecute(r)
}

/*
GetLatestAvailableSchemaVersion Retrieve latest available Cassandra schema version

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetLatestAvailableSchemaVersionRequest
*/
func (a *CassandraSchemaUpgradeAPIService) GetLatestAvailableSchemaVersion(ctx context.Context) ApiGetLatestAvailableSchemaVersionRequest {
	return ApiGetLatestAvailableSchemaVersionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetLatestAvailableSchemaVersion200Response
func (a *CassandraSchemaUpgradeAPIService) GetLatestAvailableSchemaVersionExecute(r ApiGetLatestAvailableSchemaVersionRequest) (*GetLatestAvailableSchemaVersion200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetLatestAvailableSchemaVersion200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CassandraSchemaUpgradeAPIService.GetLatestAvailableSchemaVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cassandra/version/latest"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSchemaVersionRequest struct {
	ctx context.Context
	ApiService *CassandraSchemaUpgradeAPIService
}

func (r ApiGetSchemaVersionRequest) Execute() (*GetSchemaVersion200Response, *http.Response, error) {
	return r.ApiService.GetSchemaVersionExecute(r)
}

/*
GetSchemaVersion Retrieve current Cassandra schema version

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSchemaVersionRequest
*/
func (a *CassandraSchemaUpgradeAPIService) GetSchemaVersion(ctx context.Context) ApiGetSchemaVersionRequest {
	return ApiGetSchemaVersionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetSchemaVersion200Response
func (a *CassandraSchemaUpgradeAPIService) GetSchemaVersionExecute(r ApiGetSchemaVersionRequest) (*GetSchemaVersion200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSchemaVersion200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CassandraSchemaUpgradeAPIService.GetSchemaVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cassandra/version"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpgradeSchemaVersionRequest struct {
	ctx context.Context
	ApiService *CassandraSchemaUpgradeAPIService
	getLatestAvailableSchemaVersion200Response *GetLatestAvailableSchemaVersion200Response
}

// The target schema version to upgrade to
func (r ApiUpgradeSchemaVersionRequest) GetLatestAvailableSchemaVersion200Response(getLatestAvailableSchemaVersion200Response GetLatestAvailableSchemaVersion200Response) ApiUpgradeSchemaVersionRequest {
	r.getLatestAvailableSchemaVersion200Response = &getLatestAvailableSchemaVersion200Response
	return r
}

func (r ApiUpgradeSchemaVersionRequest) Execute() (*UpgradeSchemaVersion200Response, *http.Response, error) {
	return r.ApiService.UpgradeSchemaVersionExecute(r)
}

/*
UpgradeSchemaVersion Upgrade to a specific Cassandra schema version

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpgradeSchemaVersionRequest
*/
func (a *CassandraSchemaUpgradeAPIService) UpgradeSchemaVersion(ctx context.Context) ApiUpgradeSchemaVersionRequest {
	return ApiUpgradeSchemaVersionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UpgradeSchemaVersion200Response
func (a *CassandraSchemaUpgradeAPIService) UpgradeSchemaVersionExecute(r ApiUpgradeSchemaVersionRequest) (*UpgradeSchemaVersion200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpgradeSchemaVersion200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CassandraSchemaUpgradeAPIService.UpgradeSchemaVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cassandra/version/upgrade"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getLatestAvailableSchemaVersion200Response == nil {
		return localVarReturnValue, nil, reportError("getLatestAvailableSchemaVersion200Response is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.getLatestAvailableSchemaVersion200Response
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpgradeToLatestSchemaVersionRequest struct {
	ctx context.Context
	ApiService *CassandraSchemaUpgradeAPIService
}

func (r ApiUpgradeToLatestSchemaVersionRequest) Execute() (*UpgradeToLatestSchemaVersion200Response, *http.Response, error) {
	return r.ApiService.UpgradeToLatestSchemaVersionExecute(r)
}

/*
UpgradeToLatestSchemaVersion Upgrade to the latest Cassandra schema version

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpgradeToLatestSchemaVersionRequest
*/
func (a *CassandraSchemaUpgradeAPIService) UpgradeToLatestSchemaVersion(ctx context.Context) ApiUpgradeToLatestSchemaVersionRequest {
	return ApiUpgradeToLatestSchemaVersionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UpgradeToLatestSchemaVersion200Response
func (a *CassandraSchemaUpgradeAPIService) UpgradeToLatestSchemaVersionExecute(r ApiUpgradeToLatestSchemaVersionRequest) (*UpgradeToLatestSchemaVersion200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpgradeToLatestSchemaVersion200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CassandraSchemaUpgradeAPIService.UpgradeToLatestSchemaVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cassandra/version/upgrade/latest"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}