# \EventsApi

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsList**](EventsApi.md#EventsList) | **Get** /events | 



## EventsList

> PaginatedEventList EventsList(ctx).StartDateFilter(startDateFilter).EndDateFilter(endDateFilter).ActorFilter(actorFilter).ObjectFilter(objectFilter).EventTypeFilter(eventTypeFilter).Cursor(cursor).PageSize(pageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    startDateFilter := "2021/11/01" // string | A start date filter for the events. (optional)
    endDateFilter := "2021-11-12" // string | An end date filter for the events. (optional)
    actorFilter := "29827fb8-f2dd-4e80-9576-28e31e9934ac" // string | An actor filter for the events. Supply the ID of the actor. (optional)
    objectFilter := "29827fb8-f2dd-4e80-9576-28e31e9934ac" // string | An object filter for the events. Supply the ID of the object. (optional)
    eventTypeFilter := "29827fb8-f2dd-4e80-9576-28e31e9934ac" // string | An event type filter for the events. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    pageSize := int32(200) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventsApi.EventsList(context.Background()).StartDateFilter(startDateFilter).EndDateFilter(endDateFilter).ActorFilter(actorFilter).ObjectFilter(objectFilter).EventTypeFilter(eventTypeFilter).Cursor(cursor).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsList`: PaginatedEventList
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDateFilter** | **string** | A start date filter for the events. | 
 **endDateFilter** | **string** | An end date filter for the events. | 
 **actorFilter** | **string** | An actor filter for the events. Supply the ID of the actor. | 
 **objectFilter** | **string** | An object filter for the events. Supply the ID of the object. | 
 **eventTypeFilter** | **string** | An event type filter for the events. | 
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. | 

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

