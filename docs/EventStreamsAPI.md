# \EventStreamsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEventStream**](EventStreamsAPI.md#CreateEventStream) | **Post** /event-streams | Create event stream
[**DeleteEventStream**](EventStreamsAPI.md#DeleteEventStream) | **Delete** /event-streams/{event_stream_id} | Delete event stream
[**GetEventStreams**](EventStreamsAPI.md#GetEventStreams) | **Get** /event-streams | Get event streams
[**UpdateEventStream**](EventStreamsAPI.md#UpdateEventStream) | **Put** /event-streams/{event_stream_id} | Update event stream



## CreateEventStream

> EventStream CreateEventStream(ctx).CreateEventStreamInfo(createEventStreamInfo).Execute()

Create event stream



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
	createEventStreamInfo := *openapiclient.NewCreateEventStreamInfo("Name_example", openapiclient.EventStreamConnectionTypeEnum("WEBHOOK")) // CreateEventStreamInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventStreamsAPI.CreateEventStream(context.Background()).CreateEventStreamInfo(createEventStreamInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventStreamsAPI.CreateEventStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEventStream`: EventStream
	fmt.Fprintf(os.Stdout, "Response from `EventStreamsAPI.CreateEventStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEventStreamInfo** | [**CreateEventStreamInfo**](CreateEventStreamInfo.md) |  | 

### Return type

[**EventStream**](EventStream.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEventStream

> DeleteEventStream(ctx, eventStreamId).Execute()

Delete event stream



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
	eventStreamId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the event stream.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventStreamsAPI.DeleteEventStream(context.Background(), eventStreamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventStreamsAPI.DeleteEventStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventStreamId** | **string** | The ID of the event stream. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventStreamRequest struct via the builder pattern


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


## GetEventStreams

> EventStreamList GetEventStreams(ctx).Execute()

Get event streams



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventStreamsAPI.GetEventStreams(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventStreamsAPI.GetEventStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventStreams`: EventStreamList
	fmt.Fprintf(os.Stdout, "Response from `EventStreamsAPI.GetEventStreams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventStreamsRequest struct via the builder pattern


### Return type

[**EventStreamList**](EventStreamList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEventStream

> EventStream UpdateEventStream(ctx, eventStreamId).UpdateEventStreamInfo(updateEventStreamInfo).Execute()

Update event stream



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
	eventStreamId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the event stream.
	updateEventStreamInfo := *openapiclient.NewUpdateEventStreamInfo() // UpdateEventStreamInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventStreamsAPI.UpdateEventStream(context.Background(), eventStreamId).UpdateEventStreamInfo(updateEventStreamInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventStreamsAPI.UpdateEventStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEventStream`: EventStream
	fmt.Fprintf(os.Stdout, "Response from `EventStreamsAPI.UpdateEventStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventStreamId** | **string** | The ID of the event stream. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEventStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEventStreamInfo** | [**UpdateEventStreamInfo**](UpdateEventStreamInfo.md) |  | 

### Return type

[**EventStream**](EventStream.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

