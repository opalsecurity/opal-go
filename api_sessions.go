/*
 * Opal API
 *
 * Your Home For Developer Resources.
 *
 * API version: 1.0
 * Contact: hello@opal.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opal

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// SessionsApiService SessionsApi service
type SessionsApiService service

type ApiSessionsRequest struct {
	ctx _context.Context
	ApiService *SessionsApiService
	resourceId *string
	userId *string
}

func (r ApiSessionsRequest) ResourceId(resourceId string) ApiSessionsRequest {
	r.resourceId = &resourceId
	return r
}
func (r ApiSessionsRequest) UserId(userId string) ApiSessionsRequest {
	r.userId = &userId
	return r
}

func (r ApiSessionsRequest) Execute() (SessionsList, *_nethttp.Response, error) {
	return r.ApiService.SessionsExecute(r)
}

/*
 * Sessions Method for Sessions
 * Returns a list of `Session` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSessionsRequest
 */
func (a *SessionsApiService) Sessions(ctx _context.Context) ApiSessionsRequest {
	return ApiSessionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SessionsList
 */
func (a *SessionsApiService) SessionsExecute(r ApiSessionsRequest) (SessionsList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SessionsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionsApiService.Sessions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sessions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.resourceId == nil {
		return localVarReturnValue, nil, reportError("resourceId is required and must be specified")
	}

	localVarQueryParams.Add("resource_id", parameterToString(*r.resourceId, ""))
	if r.userId != nil {
		localVarQueryParams.Add("user_id", parameterToString(*r.userId, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}