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


// GroupBindingsAPIService GroupBindingsAPI service
type GroupBindingsAPIService service

type ApiCreateGroupBindingRequest struct {
	ctx context.Context
	ApiService *GroupBindingsAPIService
	createGroupBindingInfo *CreateGroupBindingInfo
}

func (r ApiCreateGroupBindingRequest) CreateGroupBindingInfo(createGroupBindingInfo CreateGroupBindingInfo) ApiCreateGroupBindingRequest {
	r.createGroupBindingInfo = &createGroupBindingInfo
	return r
}

func (r ApiCreateGroupBindingRequest) Execute() (*GroupBinding, *http.Response, error) {
	return r.ApiService.CreateGroupBindingExecute(r)
}

/*
CreateGroupBinding Method for CreateGroupBinding

Creates a group binding.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateGroupBindingRequest
*/
func (a *GroupBindingsAPIService) CreateGroupBinding(ctx context.Context) ApiCreateGroupBindingRequest {
	return ApiCreateGroupBindingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GroupBinding
func (a *GroupBindingsAPIService) CreateGroupBindingExecute(r ApiCreateGroupBindingRequest) (*GroupBinding, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupBinding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupBindingsAPIService.CreateGroupBinding")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/group-bindings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createGroupBindingInfo == nil {
		return localVarReturnValue, nil, reportError("createGroupBindingInfo is required and must be specified")
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
	localVarPostBody = r.createGroupBindingInfo
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

type ApiDeleteGroupBindingRequest struct {
	ctx context.Context
	ApiService *GroupBindingsAPIService
	groupBindingId string
}

func (r ApiDeleteGroupBindingRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGroupBindingExecute(r)
}

/*
DeleteGroupBinding Method for DeleteGroupBinding

Deletes a group binding.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupBindingId The ID of the group binding.
 @return ApiDeleteGroupBindingRequest
*/
func (a *GroupBindingsAPIService) DeleteGroupBinding(ctx context.Context, groupBindingId string) ApiDeleteGroupBindingRequest {
	return ApiDeleteGroupBindingRequest{
		ApiService: a,
		ctx: ctx,
		groupBindingId: groupBindingId,
	}
}

// Execute executes the request
func (a *GroupBindingsAPIService) DeleteGroupBindingExecute(r ApiDeleteGroupBindingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupBindingsAPIService.DeleteGroupBinding")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/group-bindings/{group_binding_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"group_binding_id"+"}", url.PathEscape(parameterValueToString(r.groupBindingId, "groupBindingId")), -1)

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

type ApiGetGroupBindingRequest struct {
	ctx context.Context
	ApiService *GroupBindingsAPIService
	groupBindingId string
}

func (r ApiGetGroupBindingRequest) Execute() (*GroupBinding, *http.Response, error) {
	return r.ApiService.GetGroupBindingExecute(r)
}

/*
GetGroupBinding Method for GetGroupBinding

Returns a `GroupBinding` object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupBindingId The ID of the group binding.
 @return ApiGetGroupBindingRequest
*/
func (a *GroupBindingsAPIService) GetGroupBinding(ctx context.Context, groupBindingId string) ApiGetGroupBindingRequest {
	return ApiGetGroupBindingRequest{
		ApiService: a,
		ctx: ctx,
		groupBindingId: groupBindingId,
	}
}

// Execute executes the request
//  @return GroupBinding
func (a *GroupBindingsAPIService) GetGroupBindingExecute(r ApiGetGroupBindingRequest) (*GroupBinding, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupBinding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupBindingsAPIService.GetGroupBinding")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/group-bindings/{group_binding_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"group_binding_id"+"}", url.PathEscape(parameterValueToString(r.groupBindingId, "groupBindingId")), -1)

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

type ApiGetGroupBindingsRequest struct {
	ctx context.Context
	ApiService *GroupBindingsAPIService
	cursor *string
	pageSize *int32
}

// The pagination cursor value.
func (r ApiGetGroupBindingsRequest) Cursor(cursor string) ApiGetGroupBindingsRequest {
	r.cursor = &cursor
	return r
}

// Number of results to return per page. Default is 200.
func (r ApiGetGroupBindingsRequest) PageSize(pageSize int32) ApiGetGroupBindingsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetGroupBindingsRequest) Execute() (*PaginatedGroupBindingsList, *http.Response, error) {
	return r.ApiService.GetGroupBindingsExecute(r)
}

/*
GetGroupBindings Method for GetGroupBindings

Returns a list of `GroupBinding` objects.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetGroupBindingsRequest
*/
func (a *GroupBindingsAPIService) GetGroupBindings(ctx context.Context) ApiGetGroupBindingsRequest {
	return ApiGetGroupBindingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedGroupBindingsList
func (a *GroupBindingsAPIService) GetGroupBindingsExecute(r ApiGetGroupBindingsRequest) (*PaginatedGroupBindingsList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedGroupBindingsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupBindingsAPIService.GetGroupBindings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/group-bindings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
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

type ApiUpdateGroupBindingsRequest struct {
	ctx context.Context
	ApiService *GroupBindingsAPIService
	updateGroupBindingInfoList *UpdateGroupBindingInfoList
}

// Group bindings to be updated
func (r ApiUpdateGroupBindingsRequest) UpdateGroupBindingInfoList(updateGroupBindingInfoList UpdateGroupBindingInfoList) ApiUpdateGroupBindingsRequest {
	r.updateGroupBindingInfoList = &updateGroupBindingInfoList
	return r
}

func (r ApiUpdateGroupBindingsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateGroupBindingsExecute(r)
}

/*
UpdateGroupBindings Method for UpdateGroupBindings

Bulk updates a list of group bindings.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateGroupBindingsRequest
*/
func (a *GroupBindingsAPIService) UpdateGroupBindings(ctx context.Context) ApiUpdateGroupBindingsRequest {
	return ApiUpdateGroupBindingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *GroupBindingsAPIService) UpdateGroupBindingsExecute(r ApiUpdateGroupBindingsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupBindingsAPIService.UpdateGroupBindings")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/group-bindings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateGroupBindingInfoList == nil {
		return nil, reportError("updateGroupBindingInfoList is required and must be specified")
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
	localVarPostBody = r.updateGroupBindingInfoList
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
