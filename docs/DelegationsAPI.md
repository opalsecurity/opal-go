# \DelegationsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDelegation**](DelegationsAPI.md#CreateDelegation) | **Post** /delegations | 
[**DeleteDelegation**](DelegationsAPI.md#DeleteDelegation) | **Delete** /delegations/{delegation_id} | 
[**GetDelegation**](DelegationsAPI.md#GetDelegation) | **Get** /delegations/{delegation_id} | 
[**GetDelegations**](DelegationsAPI.md#GetDelegations) | **Get** /delegations | 



## CreateDelegation

> Delegation CreateDelegation(ctx).CreateDelegationRequest(createDelegationRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/opalsecurity/opal-go"
)

func main() {
	createDelegationRequest := *openapiclient.NewCreateDelegationRequest("123e4567-e89b-12d3-a456-426614174000", "7c86c85d-0651-43e2-a748-d69d658418e8", time.Now(), time.Now(), "I need to be out of the office") // CreateDelegationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegationsAPI.CreateDelegation(context.Background()).CreateDelegationRequest(createDelegationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationsAPI.CreateDelegation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDelegation`: Delegation
	fmt.Fprintf(os.Stdout, "Response from `DelegationsAPI.CreateDelegation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDelegationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDelegationRequest** | [**CreateDelegationRequest**](CreateDelegationRequest.md) |  | 

### Return type

[**Delegation**](Delegation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDelegation

> DeleteDelegation(ctx, delegationId).Execute()





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
	delegationId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the delegation to remove.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DelegationsAPI.DeleteDelegation(context.Background(), delegationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationsAPI.DeleteDelegation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegationId** | **string** | The ID of the delegation to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDelegationRequest struct via the builder pattern


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


## GetDelegation

> Delegation GetDelegation(ctx, delegationId).Execute()





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
	delegationId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the delegation to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegationsAPI.GetDelegation(context.Background(), delegationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationsAPI.GetDelegation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDelegation`: Delegation
	fmt.Fprintf(os.Stdout, "Response from `DelegationsAPI.GetDelegation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegationId** | **string** | The ID of the delegation to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDelegationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Delegation**](Delegation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDelegations

> PaginatedDelegationsList GetDelegations(ctx).DelegatorUserId(delegatorUserId).DelegateUserId(delegateUserId).Cursor(cursor).PageSize(pageSize).Execute()





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
	delegatorUserId := "29827fb8-f2dd-4e80-9576-28e31e9934ac" // string | The delegator user ID to filter delegations by the user delegating their access review requests. (optional)
	delegateUserId := "29827fb8-f2dd-4e80-9576-28e31e9934ac" // string | The delegate user ID to filter delegations by the user being delegated to. (optional)
	cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | A cursor to indicate where to start fetching results. (optional)
	pageSize := int32(200) // int32 | The maximum number of results to return per page. The default is 200. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegationsAPI.GetDelegations(context.Background()).DelegatorUserId(delegatorUserId).DelegateUserId(delegateUserId).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationsAPI.GetDelegations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDelegations`: PaginatedDelegationsList
	fmt.Fprintf(os.Stdout, "Response from `DelegationsAPI.GetDelegations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDelegationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delegatorUserId** | **string** | The delegator user ID to filter delegations by the user delegating their access review requests. | 
 **delegateUserId** | **string** | The delegate user ID to filter delegations by the user being delegated to. | 
 **cursor** | **string** | A cursor to indicate where to start fetching results. | 
 **pageSize** | **int32** | The maximum number of results to return per page. The default is 200. | 

### Return type

[**PaginatedDelegationsList**](PaginatedDelegationsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

