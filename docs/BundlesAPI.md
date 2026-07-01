# \BundlesAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBundleGroup**](BundlesAPI.md#AddBundleGroup) | **Post** /bundles/{bundle_id}/groups | 
[**AddBundleResource**](BundlesAPI.md#AddBundleResource) | **Post** /bundles/{bundle_id}/resources | 
[**CreateBundle**](BundlesAPI.md#CreateBundle) | **Post** /bundles | 
[**DeleteBundle**](BundlesAPI.md#DeleteBundle) | **Delete** /bundles/{bundle_id} | 
[**GetBundle**](BundlesAPI.md#GetBundle) | **Get** /bundles/{bundle_id} | Get bundle by ID
[**GetBundleGroups**](BundlesAPI.md#GetBundleGroups) | **Get** /bundles/{bundle_id}/groups | 
[**GetBundleResources**](BundlesAPI.md#GetBundleResources) | **Get** /bundles/{bundle_id}/resources | 
[**GetBundleVisibility**](BundlesAPI.md#GetBundleVisibility) | **Get** /bundles/{bundle_id}/visibility | 
[**GetBundles**](BundlesAPI.md#GetBundles) | **Get** /bundles | Get bundles
[**RemoveBundleGroup**](BundlesAPI.md#RemoveBundleGroup) | **Delete** /bundles/{bundle_id}/groups/{group_id} | 
[**RemoveBundleResource**](BundlesAPI.md#RemoveBundleResource) | **Delete** /bundles/{bundle_id}/resources/{resource_id} | 
[**SetBundleVisibility**](BundlesAPI.md#SetBundleVisibility) | **Put** /bundles/{bundle_id}/visibility | 
[**UpdateBundle**](BundlesAPI.md#UpdateBundle) | **Put** /bundles/{bundle_id} | 



## AddBundleGroup

> BundleGroup AddBundleGroup(ctx, bundleId).AddBundleGroupRequest(addBundleGroupRequest).Execute()





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
	addBundleGroupRequest := *openapiclient.NewAddBundleGroupRequest("72e75a6f-7183-48c5-94ff-6013f213314b") // AddBundleGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundlesAPI.AddBundleGroup(context.Background(), bundleId).AddBundleGroupRequest(addBundleGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.AddBundleGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddBundleGroup`: BundleGroup
	fmt.Fprintf(os.Stdout, "Response from `BundlesAPI.AddBundleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The ID of the bundle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddBundleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addBundleGroupRequest** | [**AddBundleGroupRequest**](AddBundleGroupRequest.md) |  | 

### Return type

[**BundleGroup**](BundleGroup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddBundleResource

> BundleResource AddBundleResource(ctx, bundleId).AddBundleResourceRequest(addBundleResourceRequest).Execute()





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
	addBundleResourceRequest := *openapiclient.NewAddBundleResourceRequest("72e75a6f-7183-48c5-94ff-6013f213314b") // AddBundleResourceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundlesAPI.AddBundleResource(context.Background(), bundleId).AddBundleResourceRequest(addBundleResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.AddBundleResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddBundleResource`: BundleResource
	fmt.Fprintf(os.Stdout, "Response from `BundlesAPI.AddBundleResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The ID of the bundle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddBundleResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addBundleResourceRequest** | [**AddBundleResourceRequest**](AddBundleResourceRequest.md) |  | 

### Return type

[**BundleResource**](BundleResource.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.BundlesAPI.CreateBundle(context.Background()).CreateBundleInfo(createBundleInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.CreateBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBundle`: Bundle
	fmt.Fprintf(os.Stdout, "Response from `BundlesAPI.CreateBundle`: %v\n", resp)
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
	r, err := apiClient.BundlesAPI.DeleteBundle(context.Background(), bundleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.DeleteBundle``: %v\n", err)
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

Get bundle by ID



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
	resp, r, err := apiClient.BundlesAPI.GetBundle(context.Background(), bundleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.GetBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBundle`: Bundle
	fmt.Fprintf(os.Stdout, "Response from `BundlesAPI.GetBundle`: %v\n", resp)
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


## GetBundleGroups

> PaginatedBundleGroupList GetBundleGroups(ctx, bundleId).PageSize(pageSize).Cursor(cursor).Execute()





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
	pageSize := int32(200) // int32 | The maximum number of groups to return from the beginning of the list. Default is 200, max is 1000. (optional)
	cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | A cursor indicating where to start fetching items after a specific point. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundlesAPI.GetBundleGroups(context.Background(), bundleId).PageSize(pageSize).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.GetBundleGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBundleGroups`: PaginatedBundleGroupList
	fmt.Fprintf(os.Stdout, "Response from `BundlesAPI.GetBundleGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The ID of the bundle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBundleGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The maximum number of groups to return from the beginning of the list. Default is 200, max is 1000. | 
 **cursor** | **string** | A cursor indicating where to start fetching items after a specific point. | 

### Return type

[**PaginatedBundleGroupList**](PaginatedBundleGroupList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBundleResources

> PaginatedBundleResourceList GetBundleResources(ctx, bundleId).PageSize(pageSize).Cursor(cursor).Execute()





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
	pageSize := int32(200) // int32 | The maximum number of resources to return from the beginning of the list. Default is 200, max is 1000. (optional)
	cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | A cursor indicating where to start fetching items after a specific point. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundlesAPI.GetBundleResources(context.Background(), bundleId).PageSize(pageSize).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.GetBundleResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBundleResources`: PaginatedBundleResourceList
	fmt.Fprintf(os.Stdout, "Response from `BundlesAPI.GetBundleResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The ID of the bundle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBundleResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The maximum number of resources to return from the beginning of the list. Default is 200, max is 1000. | 
 **cursor** | **string** | A cursor indicating where to start fetching items after a specific point. | 

### Return type

[**PaginatedBundleResourceList**](PaginatedBundleResourceList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBundleVisibility

> VisibilityInfo GetBundleVisibility(ctx, bundleId).Execute()





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
	resp, r, err := apiClient.BundlesAPI.GetBundleVisibility(context.Background(), bundleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.GetBundleVisibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBundleVisibility`: VisibilityInfo
	fmt.Fprintf(os.Stdout, "Response from `BundlesAPI.GetBundleVisibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The ID of the bundle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBundleVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VisibilityInfo**](VisibilityInfo.md)

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

Get bundles



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
	resp, r, err := apiClient.BundlesAPI.GetBundles(context.Background()).PageSize(pageSize).Cursor(cursor).Contains(contains).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.GetBundles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBundles`: PaginatedBundleList
	fmt.Fprintf(os.Stdout, "Response from `BundlesAPI.GetBundles`: %v\n", resp)
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


## RemoveBundleGroup

> RemoveBundleGroup(ctx, bundleId, groupId).AccessLevelRemoteId(accessLevelRemoteId).Execute()





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
	groupId := "72e75a6f-7183-48c5-94ff-6013f213314b" // string | The ID of the group to remove.
	accessLevelRemoteId := "arn:aws:iam::590304332660:role/AdministratorAccess" // string | The remote ID of the access level to remove. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BundlesAPI.RemoveBundleGroup(context.Background(), bundleId, groupId).AccessLevelRemoteId(accessLevelRemoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.RemoveBundleGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The ID of the bundle. | 
**groupId** | **string** | The ID of the group to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveBundleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessLevelRemoteId** | **string** | The remote ID of the access level to remove. | 

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


## RemoveBundleResource

> RemoveBundleResource(ctx, bundleId, resourceId).AccessLevelRemoteId(accessLevelRemoteId).Execute()





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
	resourceId := "72e75a6f-7183-48c5-94ff-6013f213314b" // string | The ID of the resource to remove.
	accessLevelRemoteId := "arn:aws:iam::590304332660:role/AdministratorAccess" // string | The remote ID of the access level to grant. If omitted, the default access level remote ID value (empty string) is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BundlesAPI.RemoveBundleResource(context.Background(), bundleId, resourceId).AccessLevelRemoteId(accessLevelRemoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.RemoveBundleResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The ID of the bundle. | 
**resourceId** | **string** | The ID of the resource to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveBundleResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessLevelRemoteId** | **string** | The remote ID of the access level to grant. If omitted, the default access level remote ID value (empty string) is used. | 

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


## SetBundleVisibility

> SetBundleVisibility(ctx, bundleId).VisibilityInfo(visibilityInfo).Execute()





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
	visibilityInfo := *openapiclient.NewVisibilityInfo(openapiclient.VisibilityTypeEnum("GLOBAL")) // VisibilityInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BundlesAPI.SetBundleVisibility(context.Background(), bundleId).VisibilityInfo(visibilityInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.SetBundleVisibility``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetBundleVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **visibilityInfo** | [**VisibilityInfo**](VisibilityInfo.md) |  | 

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


## UpdateBundle

> Bundle UpdateBundle(ctx, bundleId).Bundle(bundle).Execute()





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
	bundleId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the bundle to be updated.
	bundle := *openapiclient.NewBundle() // Bundle | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundlesAPI.UpdateBundle(context.Background(), bundleId).Bundle(bundle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.UpdateBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBundle`: Bundle
	fmt.Fprintf(os.Stdout, "Response from `BundlesAPI.UpdateBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The ID of the bundle to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bundle** | [**Bundle**](Bundle.md) |  | 

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

