# \PaladinAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaladin**](PaladinAPI.md#CreatePaladin) | **Post** /paladin | Create Paladin
[**CreatePaladinContextSource**](PaladinAPI.md#CreatePaladinContextSource) | **Post** /paladin/{paladin_id}/context-sources | Add a Paladin context source
[**DeletePaladin**](PaladinAPI.md#DeletePaladin) | **Delete** /paladin/{paladin_id} | Delete Paladin
[**DeletePaladinContextSource**](PaladinAPI.md#DeletePaladinContextSource) | **Delete** /paladin/{paladin_id}/context-sources/{context_source_id} | Remove a Paladin context source
[**GetPaladin**](PaladinAPI.md#GetPaladin) | **Get** /paladin/{paladin_id} | Get Paladin by ID
[**GetPaladinFromName**](PaladinAPI.md#GetPaladinFromName) | **Get** /paladin/name/{paladin_name} | Get Paladins by name
[**ListPaladinContextSources**](PaladinAPI.md#ListPaladinContextSources) | **Get** /paladin/{paladin_id}/context-sources | List Paladin context sources
[**UpdatePaladin**](PaladinAPI.md#UpdatePaladin) | **Put** /paladin/{paladin_id} | Update Paladin



## CreatePaladin

> Paladin CreatePaladin(ctx).CreatePaladinInfo(createPaladinInfo).Execute()

Create Paladin



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
	createPaladinInfo := *openapiclient.NewCreatePaladinInfo("paladin-agent-1", "7c86c85d-0651-43e2-a748-d69d658418e8") // CreatePaladinInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaladinAPI.CreatePaladin(context.Background()).CreatePaladinInfo(createPaladinInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaladinAPI.CreatePaladin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaladin`: Paladin
	fmt.Fprintf(os.Stdout, "Response from `PaladinAPI.CreatePaladin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaladinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPaladinInfo** | [**CreatePaladinInfo**](CreatePaladinInfo.md) |  | 

### Return type

[**Paladin**](Paladin.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePaladinContextSource

> PaladinContextSource CreatePaladinContextSource(ctx, paladinId).CreatePaladinContextSourceInfo(createPaladinContextSourceInfo).Execute()

Add a Paladin context source



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
	paladinId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the Paladin.
	createPaladinContextSourceInfo := *openapiclient.NewCreatePaladinContextSourceInfo(openapiclient.PaladinContextSourceKind("SLACK_CHANNEL"), openapiclient.PaladinContextSourceProvider("SLACK"), "RemoteId_example") // CreatePaladinContextSourceInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaladinAPI.CreatePaladinContextSource(context.Background(), paladinId).CreatePaladinContextSourceInfo(createPaladinContextSourceInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaladinAPI.CreatePaladinContextSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaladinContextSource`: PaladinContextSource
	fmt.Fprintf(os.Stdout, "Response from `PaladinAPI.CreatePaladinContextSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paladinId** | **string** | The ID of the Paladin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaladinContextSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPaladinContextSourceInfo** | [**CreatePaladinContextSourceInfo**](CreatePaladinContextSourceInfo.md) |  | 

### Return type

[**PaladinContextSource**](PaladinContextSource.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePaladin

> DeletePaladin(ctx, paladinId).Execute()

Delete Paladin



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
	paladinId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the Paladin.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PaladinAPI.DeletePaladin(context.Background(), paladinId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaladinAPI.DeletePaladin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paladinId** | **string** | The ID of the Paladin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaladinRequest struct via the builder pattern


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


## DeletePaladinContextSource

> DeletePaladinContextSource(ctx, paladinId, contextSourceId).Execute()

Remove a Paladin context source



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
	paladinId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the Paladin.
	contextSourceId := "8a1f2c3d-4b5e-6f70-8192-a3b4c5d6e7f8" // string | The ID of the context source.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PaladinAPI.DeletePaladinContextSource(context.Background(), paladinId, contextSourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaladinAPI.DeletePaladinContextSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paladinId** | **string** | The ID of the Paladin. | 
**contextSourceId** | **string** | The ID of the context source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaladinContextSourceRequest struct via the builder pattern


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


## GetPaladin

> Paladin GetPaladin(ctx, paladinId).Execute()

Get Paladin by ID



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
	paladinId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the Paladin.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaladinAPI.GetPaladin(context.Background(), paladinId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaladinAPI.GetPaladin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaladin`: Paladin
	fmt.Fprintf(os.Stdout, "Response from `PaladinAPI.GetPaladin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paladinId** | **string** | The ID of the Paladin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaladinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Paladin**](Paladin.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaladinFromName

> PaladinList GetPaladinFromName(ctx, paladinName).Execute()

Get Paladins by name



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
	paladinName := "paladin-agent-1" // string | The name of the Paladin.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaladinAPI.GetPaladinFromName(context.Background(), paladinName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaladinAPI.GetPaladinFromName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaladinFromName`: PaladinList
	fmt.Fprintf(os.Stdout, "Response from `PaladinAPI.GetPaladinFromName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paladinName** | **string** | The name of the Paladin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaladinFromNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaladinList**](PaladinList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaladinContextSources

> PaladinContextSourceList ListPaladinContextSources(ctx, paladinId).Execute()

List Paladin context sources



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
	paladinId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the Paladin.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaladinAPI.ListPaladinContextSources(context.Background(), paladinId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaladinAPI.ListPaladinContextSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPaladinContextSources`: PaladinContextSourceList
	fmt.Fprintf(os.Stdout, "Response from `PaladinAPI.ListPaladinContextSources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paladinId** | **string** | The ID of the Paladin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPaladinContextSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaladinContextSourceList**](PaladinContextSourceList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaladin

> Paladin UpdatePaladin(ctx, paladinId).UpdatePaladinInfo(updatePaladinInfo).Execute()

Update Paladin



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
	paladinId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the Paladin.
	updatePaladinInfo := *openapiclient.NewUpdatePaladinInfo("paladin-agent-1") // UpdatePaladinInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaladinAPI.UpdatePaladin(context.Background(), paladinId).UpdatePaladinInfo(updatePaladinInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaladinAPI.UpdatePaladin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePaladin`: Paladin
	fmt.Fprintf(os.Stdout, "Response from `PaladinAPI.UpdatePaladin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paladinId** | **string** | The ID of the Paladin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaladinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePaladinInfo** | [**UpdatePaladinInfo**](UpdatePaladinInfo.md) |  | 

### Return type

[**Paladin**](Paladin.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

