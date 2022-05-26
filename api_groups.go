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

// GroupsApiService GroupsApi service
type GroupsApiService service

type ApiConvertGroupRequest struct {
	ctx context.Context
	ApiService *GroupsApiService
	groupId string
	groupFunction *GroupFunctionEnum
	newAdminIDList *NewAdminIDList
	ownerTeamId *string
}

// The group function to convert to.
func (r ApiConvertGroupRequest) GroupFunction(groupFunction GroupFunctionEnum) ApiConvertGroupRequest {
	r.groupFunction = &groupFunction
	return r
}
func (r ApiConvertGroupRequest) NewAdminIDList(newAdminIDList NewAdminIDList) ApiConvertGroupRequest {
	r.newAdminIDList = &newAdminIDList
	return r
}
// The ID of the owning team of the group. Required when converting from Team to Group.
func (r ApiConvertGroupRequest) OwnerTeamId(ownerTeamId string) ApiConvertGroupRequest {
	r.ownerTeamId = &ownerTeamId
	return r
}

func (r ApiConvertGroupRequest) Execute() (*Group, *http.Response, error) {
	return r.ApiService.ConvertGroupExecute(r)
}

/*
ConvertGroup Method for ConvertGroup

Updates a groups function.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The ID of the group.
 @return ApiConvertGroupRequest
*/
func (a *GroupsApiService) ConvertGroup(ctx context.Context, groupId string) ApiConvertGroupRequest {
	return ApiConvertGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return Group
func (a *GroupsApiService) ConvertGroupExecute(r ApiConvertGroupRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.ConvertGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group_id}/convert"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupFunction == nil {
		return localVarReturnValue, nil, reportError("groupFunction is required and must be specified")
	}
	if r.newAdminIDList == nil {
		return localVarReturnValue, nil, reportError("newAdminIDList is required and must be specified")
	}

	localVarQueryParams.Add("group_function", parameterToString(*r.groupFunction, ""))
	if r.ownerTeamId != nil {
		localVarQueryParams.Add("owner_team_id", parameterToString(*r.ownerTeamId, ""))
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
	localVarPostBody = r.newAdminIDList
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

type ApiDeleteGroupRequest struct {
	ctx context.Context
	ApiService *GroupsApiService
	groupId string
}


func (r ApiDeleteGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGroupExecute(r)
}

/*
DeleteGroup Method for DeleteGroup

Deletes a group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The ID of the group.
 @return ApiDeleteGroupRequest
*/
func (a *GroupsApiService) DeleteGroup(ctx context.Context, groupId string) ApiDeleteGroupRequest {
	return ApiDeleteGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
func (a *GroupsApiService) DeleteGroupExecute(r ApiDeleteGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.DeleteGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

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

type ApiGetGroupMessageChannelsRequest struct {
	ctx context.Context
	ApiService *GroupsApiService
	groupId string
}


func (r ApiGetGroupMessageChannelsRequest) Execute() (*MessageChannelList, *http.Response, error) {
	return r.ApiService.GetGroupMessageChannelsExecute(r)
}

/*
GetGroupMessageChannels Method for GetGroupMessageChannels

Gets the list of message channels attached to a group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The ID of the group.
 @return ApiGetGroupMessageChannelsRequest
*/
func (a *GroupsApiService) GetGroupMessageChannels(ctx context.Context, groupId string) ApiGetGroupMessageChannelsRequest {
	return ApiGetGroupMessageChannelsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return MessageChannelList
func (a *GroupsApiService) GetGroupMessageChannelsExecute(r ApiGetGroupMessageChannelsRequest) (*MessageChannelList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessageChannelList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.GetGroupMessageChannels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group_id}/message-channels"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

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

type ApiGetGroupReviewersRequest struct {
	ctx context.Context
	ApiService *GroupsApiService
	groupId string
}


func (r ApiGetGroupReviewersRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.GetGroupReviewersExecute(r)
}

/*
GetGroupReviewers Method for GetGroupReviewers

Gets the list of team/user IDs of the reviewers for a group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The ID of the group.
 @return ApiGetGroupReviewersRequest
*/
func (a *GroupsApiService) GetGroupReviewers(ctx context.Context, groupId string) ApiGetGroupReviewersRequest {
	return ApiGetGroupReviewersRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []string
func (a *GroupsApiService) GetGroupReviewersExecute(r ApiGetGroupReviewersRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.GetGroupReviewers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group_id}/reviewers"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

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

type ApiGetGroupTagsRequest struct {
	ctx context.Context
	ApiService *GroupsApiService
	groupId string
}


func (r ApiGetGroupTagsRequest) Execute() (*TagsList, *http.Response, error) {
	return r.ApiService.GetGroupTagsExecute(r)
}

/*
GetGroupTags Method for GetGroupTags

Returns all tags applied to the group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The ID of the group whose tags to return.
 @return ApiGetGroupTagsRequest
*/
func (a *GroupsApiService) GetGroupTags(ctx context.Context, groupId string) ApiGetGroupTagsRequest {
	return ApiGetGroupTagsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return TagsList
func (a *GroupsApiService) GetGroupTagsExecute(r ApiGetGroupTagsRequest) (*TagsList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TagsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.GetGroupTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group_id}/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

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

type ApiGetGroupsRequest struct {
	ctx context.Context
	ApiService *GroupsApiService
	cursor *string
	pageSize *int32
	groupFunctionFilter *GroupFunctionEnum
	groupTypeFilter *GroupTypeEnum
}

// The pagination cursor value.
func (r ApiGetGroupsRequest) Cursor(cursor string) ApiGetGroupsRequest {
	r.cursor = &cursor
	return r
}
// Number of results to return per page. Default is 200.
func (r ApiGetGroupsRequest) PageSize(pageSize int32) ApiGetGroupsRequest {
	r.pageSize = &pageSize
	return r
}
// The group function to filter by.
func (r ApiGetGroupsRequest) GroupFunctionFilter(groupFunctionFilter GroupFunctionEnum) ApiGetGroupsRequest {
	r.groupFunctionFilter = &groupFunctionFilter
	return r
}
// The group type to filter by.
func (r ApiGetGroupsRequest) GroupTypeFilter(groupTypeFilter GroupTypeEnum) ApiGetGroupsRequest {
	r.groupTypeFilter = &groupTypeFilter
	return r
}

func (r ApiGetGroupsRequest) Execute() (*PaginatedGroupsList, *http.Response, error) {
	return r.ApiService.GetGroupsExecute(r)
}

/*
GetGroups Method for GetGroups

Returns a list of groups for your organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetGroupsRequest
*/
func (a *GroupsApiService) GetGroups(ctx context.Context) ApiGetGroupsRequest {
	return ApiGetGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedGroupsList
func (a *GroupsApiService) GetGroupsExecute(r ApiGetGroupsRequest) (*PaginatedGroupsList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedGroupsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.GetGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.groupFunctionFilter != nil {
		localVarQueryParams.Add("group_function_filter", parameterToString(*r.groupFunctionFilter, ""))
	}
	if r.groupTypeFilter != nil {
		localVarQueryParams.Add("group_type_filter", parameterToString(*r.groupTypeFilter, ""))
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

type ApiSetGroupMessageChannelsRequest struct {
	ctx context.Context
	ApiService *GroupsApiService
	groupId string
	messageChannelIDList *MessageChannelIDList
}

func (r ApiSetGroupMessageChannelsRequest) MessageChannelIDList(messageChannelIDList MessageChannelIDList) ApiSetGroupMessageChannelsRequest {
	r.messageChannelIDList = &messageChannelIDList
	return r
}

func (r ApiSetGroupMessageChannelsRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.SetGroupMessageChannelsExecute(r)
}

/*
SetGroupMessageChannels Method for SetGroupMessageChannels

Sets the list of message channels attached to a group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The ID of the group.
 @return ApiSetGroupMessageChannelsRequest
*/
func (a *GroupsApiService) SetGroupMessageChannels(ctx context.Context, groupId string) ApiSetGroupMessageChannelsRequest {
	return ApiSetGroupMessageChannelsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []string
func (a *GroupsApiService) SetGroupMessageChannelsExecute(r ApiSetGroupMessageChannelsRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.SetGroupMessageChannels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group_id}/message-channels"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

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

type ApiSetGroupReviewersRequest struct {
	ctx context.Context
	ApiService *GroupsApiService
	groupId string
	reviewerIDList *ReviewerIDList
}

func (r ApiSetGroupReviewersRequest) ReviewerIDList(reviewerIDList ReviewerIDList) ApiSetGroupReviewersRequest {
	r.reviewerIDList = &reviewerIDList
	return r
}

func (r ApiSetGroupReviewersRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.SetGroupReviewersExecute(r)
}

/*
SetGroupReviewers Method for SetGroupReviewers

Sets the list of reviewers for a group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The ID of the group.
 @return ApiSetGroupReviewersRequest
*/
func (a *GroupsApiService) SetGroupReviewers(ctx context.Context, groupId string) ApiSetGroupReviewersRequest {
	return ApiSetGroupReviewersRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []string
func (a *GroupsApiService) SetGroupReviewersExecute(r ApiSetGroupReviewersRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.SetGroupReviewers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group_id}/reviewers"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

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

type ApiUpdateGroupsRequest struct {
	ctx context.Context
	ApiService *GroupsApiService
	updateGroupInfoList *UpdateGroupInfoList
}

// Groups to be updated
func (r ApiUpdateGroupsRequest) UpdateGroupInfoList(updateGroupInfoList UpdateGroupInfoList) ApiUpdateGroupsRequest {
	r.updateGroupInfoList = &updateGroupInfoList
	return r
}

func (r ApiUpdateGroupsRequest) Execute() (*UpdateGroupInfoList, *http.Response, error) {
	return r.ApiService.UpdateGroupsExecute(r)
}

/*
UpdateGroups Method for UpdateGroups

Bulk updates a list of groups.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateGroupsRequest
*/
func (a *GroupsApiService) UpdateGroups(ctx context.Context) ApiUpdateGroupsRequest {
	return ApiUpdateGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UpdateGroupInfoList
func (a *GroupsApiService) UpdateGroupsExecute(r ApiUpdateGroupsRequest) (*UpdateGroupInfoList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateGroupInfoList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.UpdateGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateGroupInfoList == nil {
		return localVarReturnValue, nil, reportError("updateGroupInfoList is required and must be specified")
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
	localVarPostBody = r.updateGroupInfoList
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