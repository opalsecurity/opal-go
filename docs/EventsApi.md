# \EventsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Events**](EventsAPI.md#Events) | **Get** /events | 



## Events

> PaginatedEventList Events(ctx).StartDateFilter(startDateFilter).EndDateFilter(endDateFilter).ActorFilter(actorFilter).ObjectFilter(objectFilter).EventTypeFilter(eventTypeFilter).ApiTokenFilter(apiTokenFilter).Cursor(cursor).PageSize(pageSize).Execute()





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
	actorFilter := "29827fb8-f2dd-4e80-9576-28e31e9934ac" // string | An actor filter for the events. Supply the ID of the actor. (optional)
	objectFilter := "29827fb8-f2dd-4e80-9576-28e31e9934ac" // string | An object filter for the events. Supply the ID of the object. (optional)
	eventTypeFilter := "USER_MFA_RESET" // string | An event type filter for the events. (optional)
	apiTokenFilter := "fullaccess:**************************M_g==" // string | An API filter for the events. Supply the name and preview of the API token. (optional)
	cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
	pageSize := int32(200) // int32 | Number of results to return per page. Default is 200. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.Events(context.Background()).StartDateFilter(startDateFilter).EndDateFilter(endDateFilter).ActorFilter(actorFilter).ObjectFilter(objectFilter).EventTypeFilter(eventTypeFilter).ApiTokenFilter(apiTokenFilter).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.Events``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Events`: PaginatedEventList
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.Events`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDateFilter** | **string** | A start date filter for the events. | 
 **endDateFilter** | **string** | An end date filter for the events. | 
 **actorFilter** | **string** | An actor filter for the events. Supply the ID of the actor. | 
 **objectFilter** | **string** | An object filter for the events. Supply the ID of the object. | 
 **eventTypeFilter** | **string** | An event type filter for the events. | 
 **apiTokenFilter** | **string** | An API filter for the events. Supply the name and preview of the API token. | 
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. Default is 200. | 

### Return type

[**PaginatedEventList**](PaginatedEventList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

