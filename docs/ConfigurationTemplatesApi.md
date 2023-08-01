# \ConfigurationTemplatesApi

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfigurationTemplate**](ConfigurationTemplatesApi.md#CreateConfigurationTemplate) | **Post** /configuration-templates | 
[**GetConfigurationTemplates**](ConfigurationTemplatesApi.md#GetConfigurationTemplates) | **Get** /configuration-templates | 



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
    resp, r, err := apiClient.ConfigurationTemplatesApi.CreateConfigurationTemplate(context.Background()).CreateConfigurationTemplateInfo(createConfigurationTemplateInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplatesApi.CreateConfigurationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConfigurationTemplate`: ConfigurationTemplate
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplatesApi.CreateConfigurationTemplate`: %v\n", resp)
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
    resp, r, err := apiClient.ConfigurationTemplatesApi.GetConfigurationTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplatesApi.GetConfigurationTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigurationTemplates`: PaginatedConfigurationTemplateList
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplatesApi.GetConfigurationTemplates`: %v\n", resp)
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

