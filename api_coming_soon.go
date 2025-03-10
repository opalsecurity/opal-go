/*
Opal API

The Opal API is a RESTful API that allows you to interact with the Opal Security platform programmatically.

API version: 1.0
Contact: hello@opal.dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opal

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ComingSoonAPIService ComingSoonAPI service
type ComingSoonAPIService service

type ApiCreateBundleRequest struct {
	ctx context.Context
	ApiService *ComingSoonAPIService
	createBundleInfo *CreateBundleInfo
}

func (r ApiCreateBundleRequest) CreateBundleInfo(createBundleInfo CreateBundleInfo) ApiCreateBundleRequest {
	r.createBundleInfo = &createBundleInfo
	return r
}

func (r ApiCreateBundleRequest) Execute() (*Bundle, *http.Response, error) {
	return r.ApiService.CreateBundleExecute(r)
}

/*
CreateBundle Method for CreateBundle

Creates a bundle.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateBundleRequest
*/
func (a *ComingSoonAPIService) CreateBundle(ctx context.Context) ApiCreateBundleRequest {
	return ApiCreateBundleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Bundle
func (a *ComingSoonAPIService) CreateBundleExecute(r ApiCreateBundleRequest) (*Bundle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Bundle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComingSoonAPIService.CreateBundle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/experimental/bundles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createBundleInfo == nil {
		return localVarReturnValue, nil, reportError("createBundleInfo is required and must be specified")
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
	localVarPostBody = r.createBundleInfo
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

type ApiDeleteBundleRequest struct {
	ctx context.Context
	ApiService *ComingSoonAPIService
	bundleId string
}

func (r ApiDeleteBundleRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteBundleExecute(r)
}

/*
DeleteBundle Method for DeleteBundle

Deletes a bundle.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bundleId The ID of the bundle.
 @return ApiDeleteBundleRequest
*/
func (a *ComingSoonAPIService) DeleteBundle(ctx context.Context, bundleId string) ApiDeleteBundleRequest {
	return ApiDeleteBundleRequest{
		ApiService: a,
		ctx: ctx,
		bundleId: bundleId,
	}
}

// Execute executes the request
func (a *ComingSoonAPIService) DeleteBundleExecute(r ApiDeleteBundleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComingSoonAPIService.DeleteBundle")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/experimental/bundles/{bundle_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"bundle_id"+"}", url.PathEscape(parameterValueToString(r.bundleId, "bundleId")), -1)

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

type ApiGetBundleRequest struct {
	ctx context.Context
	ApiService *ComingSoonAPIService
	bundleId string
}

func (r ApiGetBundleRequest) Execute() (*Bundle, *http.Response, error) {
	return r.ApiService.GetBundleExecute(r)
}

/*
GetBundle Method for GetBundle

Returns a `Bundle` object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bundleId The ID of the bundle.
 @return ApiGetBundleRequest
*/
func (a *ComingSoonAPIService) GetBundle(ctx context.Context, bundleId string) ApiGetBundleRequest {
	return ApiGetBundleRequest{
		ApiService: a,
		ctx: ctx,
		bundleId: bundleId,
	}
}

// Execute executes the request
//  @return Bundle
func (a *ComingSoonAPIService) GetBundleExecute(r ApiGetBundleRequest) (*Bundle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Bundle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComingSoonAPIService.GetBundle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/experimental/bundles/{bundle_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"bundle_id"+"}", url.PathEscape(parameterValueToString(r.bundleId, "bundleId")), -1)

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

type ApiGetBundlesRequest struct {
	ctx context.Context
	ApiService *ComingSoonAPIService
	pageSize *int32
	cursor *string
	contains *string
}

// The maximum number of bundles to return from the beginning of the list. Default is 200, max is 1000.
func (r ApiGetBundlesRequest) PageSize(pageSize int32) ApiGetBundlesRequest {
	r.pageSize = &pageSize
	return r
}

// A cursor indicating where to start fetching items after a specific point.
func (r ApiGetBundlesRequest) Cursor(cursor string) ApiGetBundlesRequest {
	r.cursor = &cursor
	return r
}

// A filter for the bundle name.
func (r ApiGetBundlesRequest) Contains(contains string) ApiGetBundlesRequest {
	r.contains = &contains
	return r
}

func (r ApiGetBundlesRequest) Execute() (*PaginatedBundleList, *http.Response, error) {
	return r.ApiService.GetBundlesExecute(r)
}

/*
GetBundles Method for GetBundles

Returns a list of `Bundle` objects.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBundlesRequest
*/
func (a *ComingSoonAPIService) GetBundles(ctx context.Context) ApiGetBundlesRequest {
	return ApiGetBundlesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedBundleList
func (a *ComingSoonAPIService) GetBundlesExecute(r ApiGetBundlesRequest) (*PaginatedBundleList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedBundleList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComingSoonAPIService.GetBundles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/experimental/bundles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	if r.contains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "contains", r.contains, "form", "")
	}
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
