# \PaladinAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaladin**](PaladinAPI.md#GetPaladin) | **Get** /paladin/{paladin_id} | Get Paladin by ID
[**GetPaladinFromName**](PaladinAPI.md#GetPaladinFromName) | **Get** /paladin/name/{paladin_name} | Get Paladins by name



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

