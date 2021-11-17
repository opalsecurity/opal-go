# \ResourcesApi

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResourceUserAccessStatusRetrieve**](ResourcesApi.md#ResourceUserAccessStatusRetrieve) | **Get** /resource-user-access-status/{resource_id}/{user_id} | 
[**ResourceUsersList**](ResourcesApi.md#ResourceUsersList) | **Get** /resource-users | 



## ResourceUserAccessStatusRetrieve

> ResourceUserAccessStatus ResourceUserAccessStatusRetrieve(ctx, resourceId, userId).AccessLevelRemoteId(accessLevelRemoteId).Cursor(cursor).PageSize(pageSize).Execute()





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
    resourceId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the resource.
    userId := "29827fb8-f2dd-4e80-9576-28e31e9934ac" // string | The ID of the user.
    accessLevelRemoteId := "arn:aws:iam::590304332660:role/AdministratorAccess" // string | The remote ID of the access level that you wish to query for the resource. If omitted, the default access level remote ID value (empty string) is used. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    pageSize := int32(200) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.ResourceUserAccessStatusRetrieve(context.Background(), resourceId, userId).AccessLevelRemoteId(accessLevelRemoteId).Cursor(cursor).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.ResourceUserAccessStatusRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResourceUserAccessStatusRetrieve`: ResourceUserAccessStatus
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.ResourceUserAccessStatusRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 
**userId** | **string** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceUserAccessStatusRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessLevelRemoteId** | **string** | The remote ID of the access level that you wish to query for the resource. If omitted, the default access level remote ID value (empty string) is used. | 
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**ResourceUserAccessStatus**](ResourceUserAccessStatus.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceUsersList

> PaginatedResourceUserList ResourceUsersList(ctx).ResourceId(resourceId).AccessLevelRemoteId(accessLevelRemoteId).Cursor(cursor).PageSize(pageSize).Execute()





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
    resourceId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the resource.
    accessLevelRemoteId := "arn:aws:iam::590304332660:role/AdministratorAccess" // string | The remote ID of the access level that you wish to query for the resource. If omitted, the default access level remote ID value (empty string) is used. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    pageSize := int32(200) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.ResourceUsersList(context.Background()).ResourceId(resourceId).AccessLevelRemoteId(accessLevelRemoteId).Cursor(cursor).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.ResourceUsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResourceUsersList`: PaginatedResourceUserList
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.ResourceUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResourceUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceId** | **string** | The ID of the resource. | 
 **accessLevelRemoteId** | **string** | The remote ID of the access level that you wish to query for the resource. If omitted, the default access level remote ID value (empty string) is used. | 
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedResourceUserList**](PaginatedResourceUserList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

