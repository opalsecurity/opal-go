# \OpalQueriesAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunOpalQuery**](OpalQueriesAPI.md#RunOpalQuery) | **Post** /queries/run | Run an ad-hoc OpalQuery



## RunOpalQuery

> OpalQueryResults RunOpalQuery(ctx).RunOpalQueryRequest(runOpalQueryRequest).Execute()

Run an ad-hoc OpalQuery



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
	runOpalQueryRequest := *openapiclient.NewRunOpalQueryRequest("Type_example") // RunOpalQueryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpalQueriesAPI.RunOpalQuery(context.Background()).RunOpalQueryRequest(runOpalQueryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpalQueriesAPI.RunOpalQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunOpalQuery`: OpalQueryResults
	fmt.Fprintf(os.Stdout, "Response from `OpalQueriesAPI.RunOpalQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunOpalQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runOpalQueryRequest** | [**RunOpalQueryRequest**](RunOpalQueryRequest.md) |  | 

### Return type

[**OpalQueryResults**](OpalQueryResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

