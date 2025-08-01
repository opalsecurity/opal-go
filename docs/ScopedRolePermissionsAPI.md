# \ScopedRolePermissionsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetResourceScopedRolePermissions**](ScopedRolePermissionsAPI.md#GetResourceScopedRolePermissions) | **Get** /resources/{resource_id}/scoped-role-permissions | 
[**SetResourceScopedRolePermissions**](ScopedRolePermissionsAPI.md#SetResourceScopedRolePermissions) | **Put** /resources/{resource_id}/scoped-role-permissions | 



## GetResourceScopedRolePermissions

> ScopedRolePermissionList GetResourceScopedRolePermissions(ctx, resourceId).Execute()





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
	resourceId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the resource whose scoped role permissions belong to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopedRolePermissionsAPI.GetResourceScopedRolePermissions(context.Background(), resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopedRolePermissionsAPI.GetResourceScopedRolePermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceScopedRolePermissions`: ScopedRolePermissionList
	fmt.Fprintf(os.Stdout, "Response from `ScopedRolePermissionsAPI.GetResourceScopedRolePermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource whose scoped role permissions belong to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceScopedRolePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScopedRolePermissionList**](ScopedRolePermissionList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetResourceScopedRolePermissions

> ScopedRolePermissionList SetResourceScopedRolePermissions(ctx, resourceId).ScopedRolePermissionList(scopedRolePermissionList).Execute()





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
	resourceId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the resource whose scoped role permissions belong to. Must be of OPAL_SCOPED_ROLE resource type.
	scopedRolePermissionList := *openapiclient.NewScopedRolePermissionList([]openapiclient.ScopedRolePermission{*openapiclient.NewScopedRolePermission(openapiclient.RolePermissionTargetTypeEnum("RESOURCE"), openapiclient.RolePermissionNameEnum("READ"), false)}) // ScopedRolePermissionList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopedRolePermissionsAPI.SetResourceScopedRolePermissions(context.Background(), resourceId).ScopedRolePermissionList(scopedRolePermissionList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopedRolePermissionsAPI.SetResourceScopedRolePermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetResourceScopedRolePermissions`: ScopedRolePermissionList
	fmt.Fprintf(os.Stdout, "Response from `ScopedRolePermissionsAPI.SetResourceScopedRolePermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource whose scoped role permissions belong to. Must be of OPAL_SCOPED_ROLE resource type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetResourceScopedRolePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scopedRolePermissionList** | [**ScopedRolePermissionList**](ScopedRolePermissionList.md) |  | 

### Return type

[**ScopedRolePermissionList**](ScopedRolePermissionList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

