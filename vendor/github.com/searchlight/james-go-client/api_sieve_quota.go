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
	"strings"
)


// SieveQuotaAPIService SieveQuotaAPI service
type SieveQuotaAPIService service

type ApiGetUserSieveQuotaRequest struct {
	ctx context.Context
	ApiService *SieveQuotaAPIService
	userEmail string
}

func (r ApiGetUserSieveQuotaRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetUserSieveQuotaExecute(r)
}

/*
GetUserSieveQuota Retrieve user sieve quota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userEmail
 @return ApiGetUserSieveQuotaRequest
*/
func (a *SieveQuotaAPIService) GetUserSieveQuota(ctx context.Context, userEmail string) ApiGetUserSieveQuotaRequest {
	return ApiGetUserSieveQuotaRequest{
		ApiService: a,
		ctx: ctx,
		userEmail: userEmail,
	}
}

// Execute executes the request
func (a *SieveQuotaAPIService) GetUserSieveQuotaExecute(r ApiGetUserSieveQuotaRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SieveQuotaAPIService.GetUserSieveQuota")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sieve/quota/users/{userEmail}"
	localVarPath = strings.Replace(localVarPath, "{"+"userEmail"+"}", url.PathEscape(parameterValueToString(r.userEmail, "userEmail")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRemoveUserSieveQuotaRequest struct {
	ctx context.Context
	ApiService *SieveQuotaAPIService
	userEmail string
}

func (r ApiRemoveUserSieveQuotaRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveUserSieveQuotaExecute(r)
}

/*
RemoveUserSieveQuota Remove user sieve quota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userEmail
 @return ApiRemoveUserSieveQuotaRequest
*/
func (a *SieveQuotaAPIService) RemoveUserSieveQuota(ctx context.Context, userEmail string) ApiRemoveUserSieveQuotaRequest {
	return ApiRemoveUserSieveQuotaRequest{
		ApiService: a,
		ctx: ctx,
		userEmail: userEmail,
	}
}

// Execute executes the request
func (a *SieveQuotaAPIService) RemoveUserSieveQuotaExecute(r ApiRemoveUserSieveQuotaRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SieveQuotaAPIService.RemoveUserSieveQuota")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sieve/quota/users/{userEmail}"
	localVarPath = strings.Replace(localVarPath, "{"+"userEmail"+"}", url.PathEscape(parameterValueToString(r.userEmail, "userEmail")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSieveQuotaDefaultDeleteRequest struct {
	ctx context.Context
	ApiService *SieveQuotaAPIService
}

func (r ApiSieveQuotaDefaultDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.SieveQuotaDefaultDeleteExecute(r)
}

/*
SieveQuotaDefaultDelete Remove global sieve quota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSieveQuotaDefaultDeleteRequest
*/
func (a *SieveQuotaAPIService) SieveQuotaDefaultDelete(ctx context.Context) ApiSieveQuotaDefaultDeleteRequest {
	return ApiSieveQuotaDefaultDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SieveQuotaAPIService) SieveQuotaDefaultDeleteExecute(r ApiSieveQuotaDefaultDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SieveQuotaAPIService.SieveQuotaDefaultDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sieve/quota/default"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSieveQuotaDefaultGetRequest struct {
	ctx context.Context
	ApiService *SieveQuotaAPIService
}

func (r ApiSieveQuotaDefaultGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SieveQuotaDefaultGetExecute(r)
}

/*
SieveQuotaDefaultGet Retrieve global sieve quota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSieveQuotaDefaultGetRequest
*/
func (a *SieveQuotaAPIService) SieveQuotaDefaultGet(ctx context.Context) ApiSieveQuotaDefaultGetRequest {
	return ApiSieveQuotaDefaultGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SieveQuotaAPIService) SieveQuotaDefaultGetExecute(r ApiSieveQuotaDefaultGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SieveQuotaAPIService.SieveQuotaDefaultGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sieve/quota/default"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateGlobalSieveQuotaRequest struct {
	ctx context.Context
	ApiService *SieveQuotaAPIService
	updateGlobalSieveQuotaRequest *UpdateGlobalSieveQuotaRequest
}

func (r ApiUpdateGlobalSieveQuotaRequest) UpdateGlobalSieveQuotaRequest(updateGlobalSieveQuotaRequest UpdateGlobalSieveQuotaRequest) ApiUpdateGlobalSieveQuotaRequest {
	r.updateGlobalSieveQuotaRequest = &updateGlobalSieveQuotaRequest
	return r
}

func (r ApiUpdateGlobalSieveQuotaRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateGlobalSieveQuotaExecute(r)
}

/*
UpdateGlobalSieveQuota Update global sieve quota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateGlobalSieveQuotaRequest
*/
func (a *SieveQuotaAPIService) UpdateGlobalSieveQuota(ctx context.Context) ApiUpdateGlobalSieveQuotaRequest {
	return ApiUpdateGlobalSieveQuotaRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SieveQuotaAPIService) UpdateGlobalSieveQuotaExecute(r ApiUpdateGlobalSieveQuotaRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SieveQuotaAPIService.UpdateGlobalSieveQuota")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sieve/quota/default"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateGlobalSieveQuotaRequest == nil {
		return nil, reportError("updateGlobalSieveQuotaRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateGlobalSieveQuotaRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateUserSieveQuotaRequest struct {
	ctx context.Context
	ApiService *SieveQuotaAPIService
	userEmail string
	updateUserSieveQuotaRequest *UpdateUserSieveQuotaRequest
}

func (r ApiUpdateUserSieveQuotaRequest) UpdateUserSieveQuotaRequest(updateUserSieveQuotaRequest UpdateUserSieveQuotaRequest) ApiUpdateUserSieveQuotaRequest {
	r.updateUserSieveQuotaRequest = &updateUserSieveQuotaRequest
	return r
}

func (r ApiUpdateUserSieveQuotaRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateUserSieveQuotaExecute(r)
}

/*
UpdateUserSieveQuota Update user sieve quota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userEmail
 @return ApiUpdateUserSieveQuotaRequest
*/
func (a *SieveQuotaAPIService) UpdateUserSieveQuota(ctx context.Context, userEmail string) ApiUpdateUserSieveQuotaRequest {
	return ApiUpdateUserSieveQuotaRequest{
		ApiService: a,
		ctx: ctx,
		userEmail: userEmail,
	}
}

// Execute executes the request
func (a *SieveQuotaAPIService) UpdateUserSieveQuotaExecute(r ApiUpdateUserSieveQuotaRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SieveQuotaAPIService.UpdateUserSieveQuota")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sieve/quota/users/{userEmail}"
	localVarPath = strings.Replace(localVarPath, "{"+"userEmail"+"}", url.PathEscape(parameterValueToString(r.userEmail, "userEmail")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateUserSieveQuotaRequest == nil {
		return nil, reportError("updateUserSieveQuotaRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateUserSieveQuotaRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}