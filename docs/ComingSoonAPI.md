# \ComingSoonAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBundle**](ComingSoonAPI.md#CreateBundle) | **Post** /experimental/bundles | 
[**DeleteBundle**](ComingSoonAPI.md#DeleteBundle) | **Delete** /experimental/bundles/{bundle_id} | 
[**GetBundle**](ComingSoonAPI.md#GetBundle) | **Get** /experimental/bundles/{bundle_id} | 
[**GetBundles**](ComingSoonAPI.md#GetBundles) | **Get** /experimental/bundles | 



## CreateBundle

> Bundle CreateBundle(ctx).CreateBundleInfo(createBundleInfo).Execute()





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
	createBundleInfo := *openapiclient.NewCreateBundleInfo("Test Bundle", "7c86c85d-0651-43e2-a748-d69d658418e8") // CreateBundleInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComingSoonAPI.CreateBundle(context.Background()).CreateBundleInfo(createBundleInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComingSoonAPI.CreateBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBundle`: Bundle
	fmt.Fprintf(os.Stdout, "Response from `ComingSoonAPI.CreateBundle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBundleInfo** | [**CreateBundleInfo**](CreateBundleInfo.md) |  | 

### Return type

[**Bundle**](Bundle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBundle

> DeleteBundle(ctx, bundleId).Execute()





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
	bundleId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the bundle.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComingSoonAPI.DeleteBundle(context.Background(), bundleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComingSoonAPI.DeleteBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The ID of the bundle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBundleRequest struct via the builder pattern


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


## GetBundle

> Bundle GetBundle(ctx, bundleId).Execute()





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
	bundleId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the bundle.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComingSoonAPI.GetBundle(context.Background(), bundleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComingSoonAPI.GetBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBundle`: Bundle
	fmt.Fprintf(os.Stdout, "Response from `ComingSoonAPI.GetBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The ID of the bundle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Bundle**](Bundle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBundles

> PaginatedBundleList GetBundles(ctx).PageSize(pageSize).Cursor(cursor).Contains(contains).Execute()





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
	pageSize := int32(200) // int32 | The maximum number of bundles to return from the beginning of the list. Default is 200, max is 1000. (optional)
	cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | A cursor indicating where to start fetching items after a specific point. (optional)
	contains := "Engineering" // string | A filter for the bundle name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComingSoonAPI.GetBundles(context.Background()).PageSize(pageSize).Cursor(cursor).Contains(contains).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComingSoonAPI.GetBundles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBundles`: PaginatedBundleList
	fmt.Fprintf(os.Stdout, "Response from `ComingSoonAPI.GetBundles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The maximum number of bundles to return from the beginning of the list. Default is 200, max is 1000. | 
 **cursor** | **string** | A cursor indicating where to start fetching items after a specific point. | 
 **contains** | **string** | A filter for the bundle name. | 

### Return type

[**PaginatedBundleList**](PaginatedBundleList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

