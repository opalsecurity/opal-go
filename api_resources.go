/*
Opal API

Your Home For Developer Resources.

API version: 1.0
Contact: hello@opal.dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opal

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// ResourcesApiService ResourcesApi service
type ResourcesApiService service

type ApiDeleteResourceRequest struct {
	ctx context.Context
	ApiService *ResourcesApiService
	resourceId string
}


func (r ApiDeleteResourceRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteResourceExecute(r)
}

/*
DeleteResource Method for DeleteResource

Deletes a resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceId The ID of the resource.
 @return ApiDeleteResourceRequest
*/
func (a *ResourcesApiService) DeleteResource(ctx context.Context, resourceId string) ApiDeleteResourceRequest {
	return ApiDeleteResourceRequest{
		ApiService: a,
		ctx: ctx,
		resourceId: resourceId,
	}
}

// Execute executes the request
func (a *ResourcesApiService) DeleteResourceExecute(r ApiDeleteResourceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesApiService.DeleteResource")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/{resource_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"resource_id"+"}", url.PathEscape(parameterToString(r.resourceId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetResourceMessageChannelsRequest struct {
	ctx context.Context
	ApiService *ResourcesApiService
	resourceId string
}


func (r ApiGetResourceMessageChannelsRequest) Execute() (*MessageChannelList, *http.Response, error) {
	return r.ApiService.GetResourceMessageChannelsExecute(r)
}

/*
GetResourceMessageChannels Method for GetResourceMessageChannels

Gets the list of message channels attached to a resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceId The ID of the resource.
 @return ApiGetResourceMessageChannelsRequest
*/
func (a *ResourcesApiService) GetResourceMessageChannels(ctx context.Context, resourceId string) ApiGetResourceMessageChannelsRequest {
	return ApiGetResourceMessageChannelsRequest{
		ApiService: a,
		ctx: ctx,
		resourceId: resourceId,
	}
}

// Execute executes the request
//  @return MessageChannelList
func (a *ResourcesApiService) GetResourceMessageChannelsExecute(r ApiGetResourceMessageChannelsRequest) (*MessageChannelList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessageChannelList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesApiService.GetResourceMessageChannels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/{resource_id}/message-channels"
	localVarPath = strings.Replace(localVarPath, "{"+"resource_id"+"}", url.PathEscape(parameterToString(r.resourceId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetResourceReviewersRequest struct {
	ctx context.Context
	ApiService *ResourcesApiService
	resourceId string
}


func (r ApiGetResourceReviewersRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.GetResourceReviewersExecute(r)
}

/*
GetResourceReviewers Method for GetResourceReviewers

Gets the list of team/user IDs of the reviewers for a resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceId The ID of the resource.
 @return ApiGetResourceReviewersRequest
*/
func (a *ResourcesApiService) GetResourceReviewers(ctx context.Context, resourceId string) ApiGetResourceReviewersRequest {
	return ApiGetResourceReviewersRequest{
		ApiService: a,
		ctx: ctx,
		resourceId: resourceId,
	}
}

// Execute executes the request
//  @return []string
func (a *ResourcesApiService) GetResourceReviewersExecute(r ApiGetResourceReviewersRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesApiService.GetResourceReviewers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/{resource_id}/reviewers"
	localVarPath = strings.Replace(localVarPath, "{"+"resource_id"+"}", url.PathEscape(parameterToString(r.resourceId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetResourceTagsRequest struct {
	ctx context.Context
	ApiService *ResourcesApiService
	resourceId string
}


func (r ApiGetResourceTagsRequest) Execute() (*TagsList, *http.Response, error) {
	return r.ApiService.GetResourceTagsExecute(r)
}

/*
GetResourceTags Method for GetResourceTags

Returns all tags applied to the resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceId The ID of the resource whose tags to return.
 @return ApiGetResourceTagsRequest
*/
func (a *ResourcesApiService) GetResourceTags(ctx context.Context, resourceId string) ApiGetResourceTagsRequest {
	return ApiGetResourceTagsRequest{
		ApiService: a,
		ctx: ctx,
		resourceId: resourceId,
	}
}

// Execute executes the request
//  @return TagsList
func (a *ResourcesApiService) GetResourceTagsExecute(r ApiGetResourceTagsRequest) (*TagsList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TagsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesApiService.GetResourceTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/{resource_id}/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"resource_id"+"}", url.PathEscape(parameterToString(r.resourceId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetResourcesRequest struct {
	ctx context.Context
	ApiService *ResourcesApiService
	cursor *string
	pageSize *int32
	resourceTypeFilter *ResourceTypeEnum
}

// The pagination cursor value.
func (r ApiGetResourcesRequest) Cursor(cursor string) ApiGetResourcesRequest {
	r.cursor = &cursor
	return r
}
// Number of results to return per page. Default is 200.
func (r ApiGetResourcesRequest) PageSize(pageSize int32) ApiGetResourcesRequest {
	r.pageSize = &pageSize
	return r
}
// The resource type to filter by.
func (r ApiGetResourcesRequest) ResourceTypeFilter(resourceTypeFilter ResourceTypeEnum) ApiGetResourcesRequest {
	r.resourceTypeFilter = &resourceTypeFilter
	return r
}

func (r ApiGetResourcesRequest) Execute() (*PaginatedResourcesList, *http.Response, error) {
	return r.ApiService.GetResourcesExecute(r)
}

/*
GetResources Method for GetResources

Returns a list of resources for your organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetResourcesRequest
*/
func (a *ResourcesApiService) GetResources(ctx context.Context) ApiGetResourcesRequest {
	return ApiGetResourcesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedResourcesList
func (a *ResourcesApiService) GetResourcesExecute(r ApiGetResourcesRequest) (*PaginatedResourcesList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedResourcesList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesApiService.GetResources")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.resourceTypeFilter != nil {
		localVarQueryParams.Add("resource_type_filter", parameterToString(*r.resourceTypeFilter, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiResourceUserAccessStatusRetrieveRequest struct {
	ctx context.Context
	ApiService *ResourcesApiService
	resourceId string
	userId string
	accessLevelRemoteId *string
	cursor *string
	pageSize *int32
}

// The remote ID of the access level that you wish to query for the resource. If omitted, the default access level remote ID value (empty string) is used.
func (r ApiResourceUserAccessStatusRetrieveRequest) AccessLevelRemoteId(accessLevelRemoteId string) ApiResourceUserAccessStatusRetrieveRequest {
	r.accessLevelRemoteId = &accessLevelRemoteId
	return r
}
// The pagination cursor value.
func (r ApiResourceUserAccessStatusRetrieveRequest) Cursor(cursor string) ApiResourceUserAccessStatusRetrieveRequest {
	r.cursor = &cursor
	return r
}
// Number of results to return per page. Default is 200.
func (r ApiResourceUserAccessStatusRetrieveRequest) PageSize(pageSize int32) ApiResourceUserAccessStatusRetrieveRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiResourceUserAccessStatusRetrieveRequest) Execute() (*ResourceUserAccessStatus, *http.Response, error) {
	return r.ApiService.ResourceUserAccessStatusRetrieveExecute(r)
}

/*
ResourceUserAccessStatusRetrieve Method for ResourceUserAccessStatusRetrieve

Get user's access status to a resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceId The ID of the resource.
 @param userId The ID of the user.
 @return ApiResourceUserAccessStatusRetrieveRequest
*/
func (a *ResourcesApiService) ResourceUserAccessStatusRetrieve(ctx context.Context, resourceId string, userId string) ApiResourceUserAccessStatusRetrieveRequest {
	return ApiResourceUserAccessStatusRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		resourceId: resourceId,
		userId: userId,
	}
}

// Execute executes the request
//  @return ResourceUserAccessStatus
func (a *ResourcesApiService) ResourceUserAccessStatusRetrieveExecute(r ApiResourceUserAccessStatusRetrieveRequest) (*ResourceUserAccessStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResourceUserAccessStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesApiService.ResourceUserAccessStatusRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resource-user-access-status/{resource_id}/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"resource_id"+"}", url.PathEscape(parameterToString(r.resourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accessLevelRemoteId != nil {
		localVarQueryParams.Add("access_level_remote_id", parameterToString(*r.accessLevelRemoteId, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiResourceUsersRequest struct {
	ctx context.Context
	ApiService *ResourcesApiService
	resourceId *string
	accessLevelRemoteId *string
	cursor *string
	pageSize *int32
}

// The ID of the resource.
func (r ApiResourceUsersRequest) ResourceId(resourceId string) ApiResourceUsersRequest {
	r.resourceId = &resourceId
	return r
}
// The remote ID of the access level that you wish to query for the resource. If omitted, the default access level remote ID value (empty string) is used.
func (r ApiResourceUsersRequest) AccessLevelRemoteId(accessLevelRemoteId string) ApiResourceUsersRequest {
	r.accessLevelRemoteId = &accessLevelRemoteId
	return r
}
// The pagination cursor value.
func (r ApiResourceUsersRequest) Cursor(cursor string) ApiResourceUsersRequest {
	r.cursor = &cursor
	return r
}
// Number of results to return per page. Default is 200.
func (r ApiResourceUsersRequest) PageSize(pageSize int32) ApiResourceUsersRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiResourceUsersRequest) Execute() (*PaginatedResourceUserList, *http.Response, error) {
	return r.ApiService.ResourceUsersExecute(r)
}

/*
ResourceUsers Method for ResourceUsers

Returns a list of `ResourceUser` objects.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiResourceUsersRequest
*/
func (a *ResourcesApiService) ResourceUsers(ctx context.Context) ApiResourceUsersRequest {
	return ApiResourceUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedResourceUserList
func (a *ResourcesApiService) ResourceUsersExecute(r ApiResourceUsersRequest) (*PaginatedResourceUserList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedResourceUserList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesApiService.ResourceUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resource-users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.resourceId == nil {
		return localVarReturnValue, nil, reportError("resourceId is required and must be specified")
	}

	localVarQueryParams.Add("resource_id", parameterToString(*r.resourceId, ""))
	if r.accessLevelRemoteId != nil {
		localVarQueryParams.Add("access_level_remote_id", parameterToString(*r.accessLevelRemoteId, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiSetResourceMessageChannelsRequest struct {
	ctx context.Context
	ApiService *ResourcesApiService
	resourceId string
	messageChannelIDList *MessageChannelIDList
}

func (r ApiSetResourceMessageChannelsRequest) MessageChannelIDList(messageChannelIDList MessageChannelIDList) ApiSetResourceMessageChannelsRequest {
	r.messageChannelIDList = &messageChannelIDList
	return r
}

func (r ApiSetResourceMessageChannelsRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.SetResourceMessageChannelsExecute(r)
}

/*
SetResourceMessageChannels Method for SetResourceMessageChannels

Sets the list of message channels attached to a resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceId The ID of the resource.
 @return ApiSetResourceMessageChannelsRequest
*/
func (a *ResourcesApiService) SetResourceMessageChannels(ctx context.Context, resourceId string) ApiSetResourceMessageChannelsRequest {
	return ApiSetResourceMessageChannelsRequest{
		ApiService: a,
		ctx: ctx,
		resourceId: resourceId,
	}
}

// Execute executes the request
//  @return []string
func (a *ResourcesApiService) SetResourceMessageChannelsExecute(r ApiSetResourceMessageChannelsRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesApiService.SetResourceMessageChannels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/{resource_id}/message-channels"
	localVarPath = strings.Replace(localVarPath, "{"+"resource_id"+"}", url.PathEscape(parameterToString(r.resourceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.messageChannelIDList == nil {
		return localVarReturnValue, nil, reportError("messageChannelIDList is required and must be specified")
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
	localVarPostBody = r.messageChannelIDList
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiSetResourceReviewersRequest struct {
	ctx context.Context
	ApiService *ResourcesApiService
	resourceId string
	reviewerIDList *ReviewerIDList
}

func (r ApiSetResourceReviewersRequest) ReviewerIDList(reviewerIDList ReviewerIDList) ApiSetResourceReviewersRequest {
	r.reviewerIDList = &reviewerIDList
	return r
}

func (r ApiSetResourceReviewersRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.SetResourceReviewersExecute(r)
}

/*
SetResourceReviewers Method for SetResourceReviewers

Sets the list of reviewers for a resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceId The ID of the resource.
 @return ApiSetResourceReviewersRequest
*/
func (a *ResourcesApiService) SetResourceReviewers(ctx context.Context, resourceId string) ApiSetResourceReviewersRequest {
	return ApiSetResourceReviewersRequest{
		ApiService: a,
		ctx: ctx,
		resourceId: resourceId,
	}
}

// Execute executes the request
//  @return []string
func (a *ResourcesApiService) SetResourceReviewersExecute(r ApiSetResourceReviewersRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesApiService.SetResourceReviewers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/{resource_id}/reviewers"
	localVarPath = strings.Replace(localVarPath, "{"+"resource_id"+"}", url.PathEscape(parameterToString(r.resourceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.reviewerIDList == nil {
		return localVarReturnValue, nil, reportError("reviewerIDList is required and must be specified")
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
	localVarPostBody = r.reviewerIDList
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiUpdateResourcesRequest struct {
	ctx context.Context
	ApiService *ResourcesApiService
	updateResourceInfoList *UpdateResourceInfoList
}

// Resources to be updated
func (r ApiUpdateResourcesRequest) UpdateResourceInfoList(updateResourceInfoList UpdateResourceInfoList) ApiUpdateResourcesRequest {
	r.updateResourceInfoList = &updateResourceInfoList
	return r
}

func (r ApiUpdateResourcesRequest) Execute() (*UpdateResourceInfoList, *http.Response, error) {
	return r.ApiService.UpdateResourcesExecute(r)
}

/*
UpdateResources Method for UpdateResources

Bulk updates a list of resources.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateResourcesRequest
*/
func (a *ResourcesApiService) UpdateResources(ctx context.Context) ApiUpdateResourcesRequest {
	return ApiUpdateResourcesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UpdateResourceInfoList
func (a *ResourcesApiService) UpdateResourcesExecute(r ApiUpdateResourcesRequest) (*UpdateResourceInfoList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateResourceInfoList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesApiService.UpdateResources")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateResourceInfoList == nil {
		return localVarReturnValue, nil, reportError("updateResourceInfoList is required and must be specified")
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
	localVarPostBody = r.updateResourceInfoList
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
