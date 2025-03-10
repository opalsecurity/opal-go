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


// OwnersAPIService OwnersAPI service
type OwnersAPIService service

type ApiCreateOwnerRequest struct {
	ctx context.Context
	ApiService *OwnersAPIService
	createOwnerInfo *CreateOwnerInfo
}

func (r ApiCreateOwnerRequest) CreateOwnerInfo(createOwnerInfo CreateOwnerInfo) ApiCreateOwnerRequest {
	r.createOwnerInfo = &createOwnerInfo
	return r
}

func (r ApiCreateOwnerRequest) Execute() (*Owner, *http.Response, error) {
	return r.ApiService.CreateOwnerExecute(r)
}

/*
CreateOwner Method for CreateOwner

Creates an owner.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateOwnerRequest
*/
func (a *OwnersAPIService) CreateOwner(ctx context.Context) ApiCreateOwnerRequest {
	return ApiCreateOwnerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Owner
func (a *OwnersAPIService) CreateOwnerExecute(r ApiCreateOwnerRequest) (*Owner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Owner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OwnersAPIService.CreateOwner")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/owners"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOwnerInfo == nil {
		return localVarReturnValue, nil, reportError("createOwnerInfo is required and must be specified")
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
	localVarPostBody = r.createOwnerInfo
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

type ApiDeleteOwnerRequest struct {
	ctx context.Context
	ApiService *OwnersAPIService
	ownerId string
}

func (r ApiDeleteOwnerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOwnerExecute(r)
}

/*
DeleteOwner Method for DeleteOwner

Deletes an owner.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerId The ID of the owner.
 @return ApiDeleteOwnerRequest
*/
func (a *OwnersAPIService) DeleteOwner(ctx context.Context, ownerId string) ApiDeleteOwnerRequest {
	return ApiDeleteOwnerRequest{
		ApiService: a,
		ctx: ctx,
		ownerId: ownerId,
	}
}

// Execute executes the request
func (a *OwnersAPIService) DeleteOwnerExecute(r ApiDeleteOwnerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OwnersAPIService.DeleteOwner")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/owners/{owner_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"owner_id"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)

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

type ApiGetOwnerRequest struct {
	ctx context.Context
	ApiService *OwnersAPIService
	ownerId string
}

func (r ApiGetOwnerRequest) Execute() (*Owner, *http.Response, error) {
	return r.ApiService.GetOwnerExecute(r)
}

/*
GetOwner Method for GetOwner

Returns an `Owner` object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerId The ID of the owner.
 @return ApiGetOwnerRequest
*/
func (a *OwnersAPIService) GetOwner(ctx context.Context, ownerId string) ApiGetOwnerRequest {
	return ApiGetOwnerRequest{
		ApiService: a,
		ctx: ctx,
		ownerId: ownerId,
	}
}

// Execute executes the request
//  @return Owner
func (a *OwnersAPIService) GetOwnerExecute(r ApiGetOwnerRequest) (*Owner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Owner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OwnersAPIService.GetOwner")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/owners/{owner_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"owner_id"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)

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

type ApiGetOwnerFromNameRequest struct {
	ctx context.Context
	ApiService *OwnersAPIService
	ownerName string
}

func (r ApiGetOwnerFromNameRequest) Execute() (*Owner, *http.Response, error) {
	return r.ApiService.GetOwnerFromNameExecute(r)
}

/*
GetOwnerFromName Method for GetOwnerFromName

Returns an `Owner` object. Does not support owners with `/` in their name, use /owners?name=... instead.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerName The name of the owner.
 @return ApiGetOwnerFromNameRequest
*/
func (a *OwnersAPIService) GetOwnerFromName(ctx context.Context, ownerName string) ApiGetOwnerFromNameRequest {
	return ApiGetOwnerFromNameRequest{
		ApiService: a,
		ctx: ctx,
		ownerName: ownerName,
	}
}

// Execute executes the request
//  @return Owner
func (a *OwnersAPIService) GetOwnerFromNameExecute(r ApiGetOwnerFromNameRequest) (*Owner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Owner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OwnersAPIService.GetOwnerFromName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/owners/name/{owner_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"owner_name"+"}", url.PathEscape(parameterValueToString(r.ownerName, "ownerName")), -1)

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

type ApiGetOwnerUsersRequest struct {
	ctx context.Context
	ApiService *OwnersAPIService
	ownerId string
}

func (r ApiGetOwnerUsersRequest) Execute() (*UserList, *http.Response, error) {
	return r.ApiService.GetOwnerUsersExecute(r)
}

/*
GetOwnerUsers Method for GetOwnerUsers

Gets the list of users for this owner, in escalation priority order if applicable.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerId The ID of the owner.
 @return ApiGetOwnerUsersRequest
*/
func (a *OwnersAPIService) GetOwnerUsers(ctx context.Context, ownerId string) ApiGetOwnerUsersRequest {
	return ApiGetOwnerUsersRequest{
		ApiService: a,
		ctx: ctx,
		ownerId: ownerId,
	}
}

// Execute executes the request
//  @return UserList
func (a *OwnersAPIService) GetOwnerUsersExecute(r ApiGetOwnerUsersRequest) (*UserList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OwnersAPIService.GetOwnerUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/owners/{owner_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"owner_id"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)

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

type ApiGetOwnersRequest struct {
	ctx context.Context
	ApiService *OwnersAPIService
	cursor *string
	pageSize *int32
	name *string
}

// The pagination cursor value.
func (r ApiGetOwnersRequest) Cursor(cursor string) ApiGetOwnersRequest {
	r.cursor = &cursor
	return r
}

// Number of results to return per page. Default is 200.
func (r ApiGetOwnersRequest) PageSize(pageSize int32) ApiGetOwnersRequest {
	r.pageSize = &pageSize
	return r
}

// Owner name to filter by.
func (r ApiGetOwnersRequest) Name(name string) ApiGetOwnersRequest {
	r.name = &name
	return r
}

func (r ApiGetOwnersRequest) Execute() (*PaginatedOwnersList, *http.Response, error) {
	return r.ApiService.GetOwnersExecute(r)
}

/*
GetOwners Method for GetOwners

Returns a list of `Owner` objects.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetOwnersRequest
*/
func (a *OwnersAPIService) GetOwners(ctx context.Context) ApiGetOwnersRequest {
	return ApiGetOwnersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedOwnersList
func (a *OwnersAPIService) GetOwnersExecute(r ApiGetOwnersRequest) (*PaginatedOwnersList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedOwnersList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OwnersAPIService.GetOwners")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/owners"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
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

type ApiSetOwnerUsersRequest struct {
	ctx context.Context
	ApiService *OwnersAPIService
	ownerId string
	userIDList *UserIDList
}

func (r ApiSetOwnerUsersRequest) UserIDList(userIDList UserIDList) ApiSetOwnerUsersRequest {
	r.userIDList = &userIDList
	return r
}

func (r ApiSetOwnerUsersRequest) Execute() (*UserList, *http.Response, error) {
	return r.ApiService.SetOwnerUsersExecute(r)
}

/*
SetOwnerUsers Method for SetOwnerUsers

Sets the list of users for this owner. If escalation is enabled, the order of this list is the escalation priority order of the users. If the owner has a source group, adding or removing users from this list won't be possible.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerId The ID of the owner.
 @return ApiSetOwnerUsersRequest
*/
func (a *OwnersAPIService) SetOwnerUsers(ctx context.Context, ownerId string) ApiSetOwnerUsersRequest {
	return ApiSetOwnerUsersRequest{
		ApiService: a,
		ctx: ctx,
		ownerId: ownerId,
	}
}

// Execute executes the request
//  @return UserList
func (a *OwnersAPIService) SetOwnerUsersExecute(r ApiSetOwnerUsersRequest) (*UserList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OwnersAPIService.SetOwnerUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/owners/{owner_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"owner_id"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userIDList == nil {
		return localVarReturnValue, nil, reportError("userIDList is required and must be specified")
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
	localVarPostBody = r.userIDList
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

type ApiUpdateOwnersRequest struct {
	ctx context.Context
	ApiService *OwnersAPIService
	updateOwnerInfoList *UpdateOwnerInfoList
}

// Owners to be updated
func (r ApiUpdateOwnersRequest) UpdateOwnerInfoList(updateOwnerInfoList UpdateOwnerInfoList) ApiUpdateOwnersRequest {
	r.updateOwnerInfoList = &updateOwnerInfoList
	return r
}

func (r ApiUpdateOwnersRequest) Execute() (*UpdateOwnerInfoList, *http.Response, error) {
	return r.ApiService.UpdateOwnersExecute(r)
}

/*
UpdateOwners Method for UpdateOwners

Bulk updates a list of owners.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateOwnersRequest
*/
func (a *OwnersAPIService) UpdateOwners(ctx context.Context) ApiUpdateOwnersRequest {
	return ApiUpdateOwnersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UpdateOwnerInfoList
func (a *OwnersAPIService) UpdateOwnersExecute(r ApiUpdateOwnersRequest) (*UpdateOwnerInfoList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateOwnerInfoList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OwnersAPIService.UpdateOwners")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/owners"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateOwnerInfoList == nil {
		return localVarReturnValue, nil, reportError("updateOwnerInfoList is required and must be specified")
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
	localVarPostBody = r.updateOwnerInfoList
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
