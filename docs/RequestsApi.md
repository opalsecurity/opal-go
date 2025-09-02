# \RequestsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveRequest**](RequestsAPI.md#ApproveRequest) | **Post** /requests/{id}/approve | 
[**CreateRequest**](RequestsAPI.md#CreateRequest) | **Post** /requests | 
[**GetRequest**](RequestsAPI.md#GetRequest) | **Get** /requests/{id} | 
[**GetRequests**](RequestsAPI.md#GetRequests) | **Get** /requests | 
[**GetRequestsRelay**](RequestsAPI.md#GetRequestsRelay) | **Get** /requests/relay | 



## ApproveRequest

> ApproveRequest200Response ApproveRequest(ctx, id).ApproveRequestRequest(approveRequestRequest).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the request to approve
	approveRequestRequest := *openapiclient.NewApproveRequestRequest(openapiclient.RequestApprovalEnum("REGULAR")) // ApproveRequestRequest | Approval parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.ApproveRequest(context.Background(), id).ApproveRequestRequest(approveRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.ApproveRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveRequest`: ApproveRequest200Response
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.ApproveRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the request to approve | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approveRequestRequest** | [**ApproveRequestRequest**](ApproveRequestRequest.md) | Approval parameters | 

### Return type

[**ApproveRequest200Response**](ApproveRequest200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRequest

> CreateRequest200Response CreateRequest(ctx).CreateRequestInfo(createRequestInfo).Execute()





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
	createRequestInfo := *openapiclient.NewCreateRequestInfo([]openapiclient.CreateRequestInfoResourcesInner{*openapiclient.NewCreateRequestInfoResourcesInner()}, []openapiclient.CreateRequestInfoGroupsInner{*openapiclient.NewCreateRequestInfoGroupsInner("f454d283-ca87-4a8a-bdbb-df212eca5353")}, "Reason_example", int32(123)) // CreateRequestInfo | Resources to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.CreateRequest(context.Background()).CreateRequestInfo(createRequestInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.CreateRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRequest`: CreateRequest200Response
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.CreateRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRequestInfo** | [**CreateRequestInfo**](CreateRequestInfo.md) | Resources to be updated | 

### Return type

[**CreateRequest200Response**](CreateRequest200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequest

> Request GetRequest(ctx, id).Execute()





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
	id := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.GetRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.GetRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequest`: Request
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.GetRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Request**](Request.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequests

> RequestList GetRequests(ctx).StartDateFilter(startDateFilter).EndDateFilter(endDateFilter).RequesterId(requesterId).TargetUserId(targetUserId).Cursor(cursor).PageSize(pageSize).ShowPendingOnly(showPendingOnly).Execute()





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
	startDateFilter := "2021-11-01" // string | A start date filter for the events. (optional)
	endDateFilter := "2021-11-12" // string | An end date filter for the events. (optional)
	requesterId := "37cb7e41-12ba-46da-92ff-030abe0450b1" // string | Filter requests by their requester ID. (optional)
	targetUserId := "37cb7e41-12ba-46da-92ff-030abe0450b1" // string | Filter requests by their target user ID. (optional)
	cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
	pageSize := int32(200) // int32 | Number of results to return per page. Default is 200. (optional)
	showPendingOnly := true // bool | Boolean toggle for if it should only show pending requests. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.GetRequests(context.Background()).StartDateFilter(startDateFilter).EndDateFilter(endDateFilter).RequesterId(requesterId).TargetUserId(targetUserId).Cursor(cursor).PageSize(pageSize).ShowPendingOnly(showPendingOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.GetRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequests`: RequestList
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.GetRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDateFilter** | **string** | A start date filter for the events. | 
 **endDateFilter** | **string** | An end date filter for the events. | 
 **requesterId** | **string** | Filter requests by their requester ID. | 
 **targetUserId** | **string** | Filter requests by their target user ID. | 
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. Default is 200. | 
 **showPendingOnly** | **bool** | Boolean toggle for if it should only show pending requests. | 

### Return type

[**RequestList**](RequestList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestsRelay

> RequestConnection GetRequestsRelay(ctx).First(first).After(after).Last(last).Before(before).Status(status).To(to).From(from).Execute()





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
	first := int32(10) // int32 | Number of results to return after the cursor. Use either first/after or last/before, not both. (optional)
	after := "Y3Vyc29yOnYyOpK5MjAyMS0wMS0wN1QwNzo0MToyNy4xMTlaFjYwZmM2YmJlZjk4YzE1N2ZhNjFhYjk4Nw==" // string | Cursor to fetch results after. Used with 'first' for forward pagination. (optional)
	last := int32(10) // int32 | Number of results to return before the cursor. Use either first/after or last/before, not both. (optional)
	before := "Y3Vyc29yOnYyOpK5MjAyMS0wMS0wN1QwNzo0MToyNy4xMTlaFjYwZmM2YmJlZjk4YzE1N2ZhNjFhYjk4Nw==" // string | Cursor to fetch results before. Used with 'last' for backward pagination. (optional)
	status := openapiclient.RequestStatusEnum("PENDING") // RequestStatusEnum | Filter requests by their status. (optional)
	to := "37cb7e41-12ba-46da-92ff-030abe0450b1" // string | Filter requests assigned to a specific user ID. (optional)
	from := "37cb7e41-12ba-46da-92ff-030abe0450b1" // string | Filter requests made by a specific user ID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.GetRequestsRelay(context.Background()).First(first).After(after).Last(last).Before(before).Status(status).To(to).From(from).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.GetRequestsRelay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequestsRelay`: RequestConnection
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.GetRequestsRelay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestsRelayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **first** | **int32** | Number of results to return after the cursor. Use either first/after or last/before, not both. | 
 **after** | **string** | Cursor to fetch results after. Used with &#39;first&#39; for forward pagination. | 
 **last** | **int32** | Number of results to return before the cursor. Use either first/after or last/before, not both. | 
 **before** | **string** | Cursor to fetch results before. Used with &#39;last&#39; for backward pagination. | 
 **status** | [**RequestStatusEnum**](RequestStatusEnum.md) | Filter requests by their status. | 
 **to** | **string** | Filter requests assigned to a specific user ID. | 
 **from** | **string** | Filter requests made by a specific user ID. | 

### Return type

[**RequestConnection**](RequestConnection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

