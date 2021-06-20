# \PermissionsApi

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermAccessStatusRetrieve**](PermissionsApi.md#PermAccessStatusRetrieve) | **Get** /permission-access-status/{permission_id}/{user_id} | 
[**PermissionUsersList**](PermissionsApi.md#PermissionUsersList) | **Get** /permission-users | 



## PermAccessStatusRetrieve

> PermissionAccessStatus PermAccessStatusRetrieve(ctx, permissionId, userId).Cursor(cursor).PageSize(pageSize).Execute()





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
    permissionId := TODO // string | The ID of the permission.
    userId := TODO // string | The ID of the user.
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    pageSize := int32(200) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionsApi.PermAccessStatusRetrieve(context.Background(), permissionId, userId).Cursor(cursor).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermAccessStatusRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermAccessStatusRetrieve`: PermissionAccessStatus
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermAccessStatusRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionId** | [**string**](.md) | The ID of the permission. | 
**userId** | [**string**](.md) | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermAccessStatusRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PermissionAccessStatus**](PermissionAccessStatus.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionUsersList

> PaginatedPermissionUserList PermissionUsersList(ctx).PermissionId(permissionId).Cursor(cursor).PageSize(pageSize).Execute()





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
    permissionId := TODO // string | The ID of the permission.
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    pageSize := int32(200) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionsApi.PermissionUsersList(context.Background()).PermissionId(permissionId).Cursor(cursor).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionUsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionUsersList`: PaginatedPermissionUserList
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionId** | [**string**](string.md) | The ID of the permission. | 
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedPermissionUserList**](PaginatedPermissionUserList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

