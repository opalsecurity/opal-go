# \OpalQueriesAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunOpalQuery**](OpalQueriesAPI.md#RunOpalQuery) | **Post** /queries/run | Run an ad-hoc OpalQuery



## RunOpalQuery

> OpalNodeQueryResults RunOpalQuery(ctx).Body(body).Execute()

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
	body := OpalNodeQuery(987) // OpalNodeQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpalQueriesAPI.RunOpalQuery(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpalQueriesAPI.RunOpalQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunOpalQuery`: OpalNodeQueryResults
	fmt.Fprintf(os.Stdout, "Response from `OpalQueriesAPI.RunOpalQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunOpalQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **OpalNodeQuery** |  | 

### Return type

[**OpalNodeQueryResults**](OpalNodeQueryResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

