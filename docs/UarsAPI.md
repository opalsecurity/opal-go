# \UarsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUar**](UarsAPI.md#CreateUar) | **Post** /uar | 
[**GetUARs**](UarsAPI.md#GetUARs) | **Get** /uars | 
[**GetUar**](UarsAPI.md#GetUar) | **Get** /uar/{uar_id} | 



## CreateUar

> UAR CreateUar(ctx).CreateUARInfo(createUARInfo).Execute()





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
	createUARInfo := *openapiclient.NewCreateUARInfo("Monthly UAR (July)", openapiclient.UARReviewerAssignmentPolicyEnum("MANUALLY"), false, time.Now(), "America/Los_Angeles", false) // CreateUARInfo | The settings of the UAR.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UarsAPI.CreateUar(context.Background()).CreateUARInfo(createUARInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UarsAPI.CreateUar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUar`: UAR
	fmt.Fprintf(os.Stdout, "Response from `UarsAPI.CreateUar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUARInfo** | [**CreateUARInfo**](CreateUARInfo.md) | The settings of the UAR. | 

### Return type

[**UAR**](UAR.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUARs

> PaginatedUARsList GetUARs(ctx).Cursor(cursor).PageSize(pageSize).Execute()





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
	resp, r, err := apiClient.UarsAPI.GetUARs(context.Background()).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UarsAPI.GetUARs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUARs`: PaginatedUARsList
	fmt.Fprintf(os.Stdout, "Response from `UarsAPI.GetUARs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUARsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. Default is 200. | 

### Return type

[**PaginatedUARsList**](PaginatedUARsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUar

> UAR GetUar(ctx, uarId).Execute()





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
	uarId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the UAR.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UarsAPI.GetUar(context.Background(), uarId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UarsAPI.GetUar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUar`: UAR
	fmt.Fprintf(os.Stdout, "Response from `UarsAPI.GetUar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uarId** | **string** | The ID of the UAR. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UAR**](UAR.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

