# \SessionsApi

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Sessions**](SessionsApi.md#Sessions) | **Get** /sessions | 



## Sessions

> SessionsList Sessions(ctx).ResourceId(resourceId).UserId(userId).Execute()





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
    resourceId := TODO // string | The ID of the resource.
    userId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the user you wish to query sessions for. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionsApi.Sessions(context.Background()).ResourceId(resourceId).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.Sessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Sessions`: SessionsList
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.Sessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceId** | [**string**](string.md) | The ID of the resource. | 
 **userId** | **string** | The ID of the user you wish to query sessions for. | 

### Return type

[**SessionsList**](SessionsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

