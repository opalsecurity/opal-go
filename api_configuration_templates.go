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


// ConfigurationTemplatesAPIService ConfigurationTemplatesAPI service
type ConfigurationTemplatesAPIService service

type ApiCreateConfigurationTemplateRequest struct {
	ctx context.Context
	ApiService *ConfigurationTemplatesAPIService
	createConfigurationTemplateInfo *CreateConfigurationTemplateInfo
}

func (r ApiCreateConfigurationTemplateRequest) CreateConfigurationTemplateInfo(createConfigurationTemplateInfo CreateConfigurationTemplateInfo) ApiCreateConfigurationTemplateRequest {
	r.createConfigurationTemplateInfo = &createConfigurationTemplateInfo
	return r
}

func (r ApiCreateConfigurationTemplateRequest) Execute() (*ConfigurationTemplate, *http.Response, error) {
	return r.ApiService.CreateConfigurationTemplateExecute(r)
}

/*
CreateConfigurationTemplate Method for CreateConfigurationTemplate

Creates a configuration template.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateConfigurationTemplateRequest
*/
func (a *ConfigurationTemplatesAPIService) CreateConfigurationTemplate(ctx context.Context) ApiCreateConfigurationTemplateRequest {
	return ApiCreateConfigurationTemplateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConfigurationTemplate
func (a *ConfigurationTemplatesAPIService) CreateConfigurationTemplateExecute(r ApiCreateConfigurationTemplateRequest) (*ConfigurationTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigurationTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationTemplatesAPIService.CreateConfigurationTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configuration-templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createConfigurationTemplateInfo == nil {
		return localVarReturnValue, nil, reportError("createConfigurationTemplateInfo is required and must be specified")
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
	localVarPostBody = r.createConfigurationTemplateInfo
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

type ApiDeleteConfigurationTemplateRequest struct {
	ctx context.Context
	ApiService *ConfigurationTemplatesAPIService
	configurationTemplateId string
}

func (r ApiDeleteConfigurationTemplateRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteConfigurationTemplateExecute(r)
}

/*
DeleteConfigurationTemplate Method for DeleteConfigurationTemplate

Deletes a configuration template.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configurationTemplateId The ID of the configuration template.
 @return ApiDeleteConfigurationTemplateRequest
*/
func (a *ConfigurationTemplatesAPIService) DeleteConfigurationTemplate(ctx context.Context, configurationTemplateId string) ApiDeleteConfigurationTemplateRequest {
	return ApiDeleteConfigurationTemplateRequest{
		ApiService: a,
		ctx: ctx,
		configurationTemplateId: configurationTemplateId,
	}
}

// Execute executes the request
func (a *ConfigurationTemplatesAPIService) DeleteConfigurationTemplateExecute(r ApiDeleteConfigurationTemplateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationTemplatesAPIService.DeleteConfigurationTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configuration-templates/{configuration_template_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"configuration_template_id"+"}", url.PathEscape(parameterValueToString(r.configurationTemplateId, "configurationTemplateId")), -1)

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

type ApiGetConfigurationTemplatesRequest struct {
	ctx context.Context
	ApiService *ConfigurationTemplatesAPIService
}

func (r ApiGetConfigurationTemplatesRequest) Execute() (*PaginatedConfigurationTemplateList, *http.Response, error) {
	return r.ApiService.GetConfigurationTemplatesExecute(r)
}

/*
GetConfigurationTemplates Method for GetConfigurationTemplates

Returns a list of `ConfigurationTemplate` objects.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetConfigurationTemplatesRequest
*/
func (a *ConfigurationTemplatesAPIService) GetConfigurationTemplates(ctx context.Context) ApiGetConfigurationTemplatesRequest {
	return ApiGetConfigurationTemplatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedConfigurationTemplateList
func (a *ConfigurationTemplatesAPIService) GetConfigurationTemplatesExecute(r ApiGetConfigurationTemplatesRequest) (*PaginatedConfigurationTemplateList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedConfigurationTemplateList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationTemplatesAPIService.GetConfigurationTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configuration-templates"

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

type ApiUpdateConfigurationTemplateRequest struct {
	ctx context.Context
	ApiService *ConfigurationTemplatesAPIService
	updateConfigurationTemplateInfo *UpdateConfigurationTemplateInfo
}

// Configuration template to be updated
func (r ApiUpdateConfigurationTemplateRequest) UpdateConfigurationTemplateInfo(updateConfigurationTemplateInfo UpdateConfigurationTemplateInfo) ApiUpdateConfigurationTemplateRequest {
	r.updateConfigurationTemplateInfo = &updateConfigurationTemplateInfo
	return r
}

func (r ApiUpdateConfigurationTemplateRequest) Execute() (*ConfigurationTemplate, *http.Response, error) {
	return r.ApiService.UpdateConfigurationTemplateExecute(r)
}

/*
UpdateConfigurationTemplate Method for UpdateConfigurationTemplate

Update a configuration template.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateConfigurationTemplateRequest
*/
func (a *ConfigurationTemplatesAPIService) UpdateConfigurationTemplate(ctx context.Context) ApiUpdateConfigurationTemplateRequest {
	return ApiUpdateConfigurationTemplateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConfigurationTemplate
func (a *ConfigurationTemplatesAPIService) UpdateConfigurationTemplateExecute(r ApiUpdateConfigurationTemplateRequest) (*ConfigurationTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigurationTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationTemplatesAPIService.UpdateConfigurationTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configuration-templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateConfigurationTemplateInfo == nil {
		return localVarReturnValue, nil, reportError("updateConfigurationTemplateInfo is required and must be specified")
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
	localVarPostBody = r.updateConfigurationTemplateInfo
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
