# \RequestsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRequest**](RequestsAPI.md#CreateRequest) | **Post** /requests | 
[**GetRequests**](RequestsAPI.md#GetRequests) | **Get** /requests | 



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


## GetRequests

> RequestList GetRequests(ctx).Cursor(cursor).PageSize(pageSize).ShowPendingOnly(showPendingOnly).Execute()





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
	showPendingOnly := true // bool | Boolean toggle for if it should only show pending requests. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.GetRequests(context.Background()).Cursor(cursor).PageSize(pageSize).ShowPendingOnly(showPendingOnly).Execute()
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

