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
)


// RequestsAPIService RequestsAPI service
type RequestsAPIService service

type ApiCreateRequestRequest struct {
	ctx context.Context
	ApiService *RequestsAPIService
	createRequestInfo *CreateRequestInfo
}

// Resources to be updated
func (r ApiCreateRequestRequest) CreateRequestInfo(createRequestInfo CreateRequestInfo) ApiCreateRequestRequest {
	r.createRequestInfo = &createRequestInfo
	return r
}

func (r ApiCreateRequestRequest) Execute() (*CreateRequest200Response, *http.Response, error) {
	return r.ApiService.CreateRequestExecute(r)
}

/*
CreateRequest Method for CreateRequest

Create an access request

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateRequestRequest
*/
func (a *RequestsAPIService) CreateRequest(ctx context.Context) ApiCreateRequestRequest {
	return ApiCreateRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateRequest200Response
func (a *RequestsAPIService) CreateRequestExecute(r ApiCreateRequestRequest) (*CreateRequest200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateRequest200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestsAPIService.CreateRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/requests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createRequestInfo == nil {
		return localVarReturnValue, nil, reportError("createRequestInfo is required and must be specified")
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
	localVarPostBody = r.createRequestInfo
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

type ApiGetRequestsRequest struct {
	ctx context.Context
	ApiService *RequestsAPIService
	cursor *string
	pageSize *int32
	showPendingOnly *bool
}

// The pagination cursor value.
func (r ApiGetRequestsRequest) Cursor(cursor string) ApiGetRequestsRequest {
	r.cursor = &cursor
	return r
}

// Number of results to return per page. Default is 200.
func (r ApiGetRequestsRequest) PageSize(pageSize int32) ApiGetRequestsRequest {
	r.pageSize = &pageSize
	return r
}

// Boolean toggle for if it should only show pending requests.
func (r ApiGetRequestsRequest) ShowPendingOnly(showPendingOnly bool) ApiGetRequestsRequest {
	r.showPendingOnly = &showPendingOnly
	return r
}

func (r ApiGetRequestsRequest) Execute() (*RequestList, *http.Response, error) {
	return r.ApiService.GetRequestsExecute(r)
}

/*
GetRequests Method for GetRequests

Returns a list of requests for your organization that is visible by the admin.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRequestsRequest
*/
func (a *RequestsAPIService) GetRequests(ctx context.Context) ApiGetRequestsRequest {
	return ApiGetRequestsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RequestList
func (a *RequestsAPIService) GetRequestsExecute(r ApiGetRequestsRequest) (*RequestList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RequestList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestsAPIService.GetRequests")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/requests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
	}
	if r.showPendingOnly != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_pending_only", r.showPendingOnly, "form", "")
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
