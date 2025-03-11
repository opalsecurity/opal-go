# \GroupBindingsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupBinding**](GroupBindingsAPI.md#CreateGroupBinding) | **Post** /group-bindings | 
[**DeleteGroupBinding**](GroupBindingsAPI.md#DeleteGroupBinding) | **Delete** /group-bindings/{group_binding_id} | 
[**GetGroupBinding**](GroupBindingsAPI.md#GetGroupBinding) | **Get** /group-bindings/{group_binding_id} | 
[**GetGroupBindings**](GroupBindingsAPI.md#GetGroupBindings) | **Get** /group-bindings | 
[**UpdateGroupBindings**](GroupBindingsAPI.md#UpdateGroupBindings) | **Put** /group-bindings | 



## CreateGroupBinding

> GroupBinding CreateGroupBinding(ctx).CreateGroupBindingInfo(createGroupBindingInfo).Execute()





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
	createGroupBindingInfo := *openapiclient.NewCreateGroupBindingInfo("f454d283-ca87-4a8a-bdbb-df212eca5353", []openapiclient.CreateGroupBindingInfoGroupsInner{*openapiclient.NewCreateGroupBindingInfoGroupsInner("f454d283-ca87-4a8a-bdbb-df212eca5353")}) // CreateGroupBindingInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupBindingsAPI.CreateGroupBinding(context.Background()).CreateGroupBindingInfo(createGroupBindingInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupBindingsAPI.CreateGroupBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupBinding`: GroupBinding
	fmt.Fprintf(os.Stdout, "Response from `GroupBindingsAPI.CreateGroupBinding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGroupBindingInfo** | [**CreateGroupBindingInfo**](CreateGroupBindingInfo.md) |  | 

### Return type

[**GroupBinding**](GroupBinding.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupBinding

> DeleteGroupBinding(ctx, groupBindingId).Execute()





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
	groupBindingId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group binding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupBindingsAPI.DeleteGroupBinding(context.Background(), groupBindingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupBindingsAPI.DeleteGroupBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupBindingId** | **string** | The ID of the group binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupBindingRequest struct via the builder pattern


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


## GetGroupBinding

> GroupBinding GetGroupBinding(ctx, groupBindingId).Execute()





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
	groupBindingId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the group binding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupBindingsAPI.GetGroupBinding(context.Background(), groupBindingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupBindingsAPI.GetGroupBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupBinding`: GroupBinding
	fmt.Fprintf(os.Stdout, "Response from `GroupBindingsAPI.GetGroupBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupBindingId** | **string** | The ID of the group binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupBinding**](GroupBinding.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupBindings

> PaginatedGroupBindingsList GetGroupBindings(ctx).Cursor(cursor).PageSize(pageSize).Execute()





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
	cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
	pageSize := int32(200) // int32 | Number of results to return per page. Default is 200. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupBindingsAPI.GetGroupBindings(context.Background()).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupBindingsAPI.GetGroupBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupBindings`: PaginatedGroupBindingsList
	fmt.Fprintf(os.Stdout, "Response from `GroupBindingsAPI.GetGroupBindings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. Default is 200. | 

### Return type

[**PaginatedGroupBindingsList**](PaginatedGroupBindingsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupBindings

> UpdateGroupBindings(ctx).UpdateGroupBindingInfoList(updateGroupBindingInfoList).Execute()





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
	updateGroupBindingInfoList := *openapiclient.NewUpdateGroupBindingInfoList([]openapiclient.UpdateGroupBindingInfo{*openapiclient.NewUpdateGroupBindingInfo("0ae19dbf-324d-4216-999c-574d46182c7e", "f454d283-ca87-4a8a-bdbb-df212eca5353", []openapiclient.CreateGroupBindingInfoGroupsInner{*openapiclient.NewCreateGroupBindingInfoGroupsInner("f454d283-ca87-4a8a-bdbb-df212eca5353")})}) // UpdateGroupBindingInfoList | Group bindings to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupBindingsAPI.UpdateGroupBindings(context.Background()).UpdateGroupBindingInfoList(updateGroupBindingInfoList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupBindingsAPI.UpdateGroupBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateGroupBindingInfoList** | [**UpdateGroupBindingInfoList**](UpdateGroupBindingInfoList.md) | Group bindings to be updated | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

