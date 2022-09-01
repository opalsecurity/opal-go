# \AppsApi

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApps**](AppsApi.md#GetApps) | **Get** /apps | 



## GetApps

> AppsList GetApps(ctx).AppTypeFilter(appTypeFilter).OwnerFilter(ownerFilter).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appTypeFilter := []openapiclient.AppTypeEnum{openapiclient.AppTypeEnum("ACTIVE_DIRECTORY")} // []AppTypeEnum | A list of app types to filter by. (optional)
    ownerFilter := "29827fb8-f2dd-4e80-9576-28e31e9934ac" // string | An owner ID to filter by. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.GetApps(context.Background()).AppTypeFilter(appTypeFilter).OwnerFilter(ownerFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.GetApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApps`: AppsList
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.GetApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appTypeFilter** | [**[]AppTypeEnum**](AppTypeEnum.md) | A list of app types to filter by. | 
 **ownerFilter** | **string** | An owner ID to filter by. | 

### Return type

[**AppsList**](AppsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

