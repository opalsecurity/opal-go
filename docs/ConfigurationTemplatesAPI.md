# \ConfigurationTemplatesAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfigurationTemplate**](ConfigurationTemplatesAPI.md#CreateConfigurationTemplate) | **Post** /configuration-templates | 
[**DeleteConfigurationTemplate**](ConfigurationTemplatesAPI.md#DeleteConfigurationTemplate) | **Delete** /configuration-templates/{configuration_template_id} | 
[**GetConfigurationTemplates**](ConfigurationTemplatesAPI.md#GetConfigurationTemplates) | **Get** /configuration-templates | 
[**UpdateConfigurationTemplate**](ConfigurationTemplatesAPI.md#UpdateConfigurationTemplate) | **Put** /configuration-templates | 



## CreateConfigurationTemplate

> ConfigurationTemplate CreateConfigurationTemplate(ctx).CreateConfigurationTemplateInfo(createConfigurationTemplateInfo).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opalsecurity/opal-go"
)

func main() {
	createConfigurationTemplateInfo := *openapiclient.NewCreateConfigurationTemplateInfo("7c86c85d-0651-43e2-a748-d69d658418e8", *openapiclient.NewVisibilityInfo(openapiclient.VisibilityTypeEnum("GLOBAL")), false, false, "Prod AWS Template") // CreateConfigurationTemplateInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationTemplatesAPI.CreateConfigurationTemplate(context.Background()).CreateConfigurationTemplateInfo(createConfigurationTemplateInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplatesAPI.CreateConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConfigurationTemplate`: ConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplatesAPI.CreateConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConfigurationTemplateInfo** | [**CreateConfigurationTemplateInfo**](CreateConfigurationTemplateInfo.md) |  | 

### Return type

[**ConfigurationTemplate**](ConfigurationTemplate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfigurationTemplate

> DeleteConfigurationTemplate(ctx, configurationTemplateId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opalsecurity/opal-go"
)

func main() {
	configurationTemplateId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the configuration template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationTemplatesAPI.DeleteConfigurationTemplate(context.Background(), configurationTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplatesAPI.DeleteConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationTemplateId** | **string** | The ID of the configuration template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationTemplates

> PaginatedConfigurationTemplateList GetConfigurationTemplates(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opalsecurity/opal-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationTemplatesAPI.GetConfigurationTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplatesAPI.GetConfigurationTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationTemplates`: PaginatedConfigurationTemplateList
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplatesAPI.GetConfigurationTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationTemplatesRequest struct via the builder pattern


### Return type

[**PaginatedConfigurationTemplateList**](PaginatedConfigurationTemplateList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigurationTemplate

> ConfigurationTemplate UpdateConfigurationTemplate(ctx).UpdateConfigurationTemplateInfo(updateConfigurationTemplateInfo).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opalsecurity/opal-go"
)

func main() {
	updateConfigurationTemplateInfo := *openapiclient.NewUpdateConfigurationTemplateInfo("7c86c85d-0651-43e2-a748-d69d658418e8") // UpdateConfigurationTemplateInfo | Configuration template to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationTemplatesAPI.UpdateConfigurationTemplate(context.Background()).UpdateConfigurationTemplateInfo(updateConfigurationTemplateInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplatesAPI.UpdateConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConfigurationTemplate`: ConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplatesAPI.UpdateConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateConfigurationTemplateInfo** | [**UpdateConfigurationTemplateInfo**](UpdateConfigurationTemplateInfo.md) | Configuration template to be updated | 

### Return type

[**ConfigurationTemplate**](ConfigurationTemplate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

