# \DefaultApi

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateConfigurationTemplate**](DefaultApi.md#UpdateConfigurationTemplate) | **Put** /configuration-templates | 



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
    resp, r, err := apiClient.DefaultApi.UpdateConfigurationTemplate(context.Background()).UpdateConfigurationTemplateInfo(updateConfigurationTemplateInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConfigurationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigurationTemplate`: ConfigurationTemplate
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConfigurationTemplate`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

