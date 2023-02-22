# \OnCallSchedulesApi

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOnCallSchedule**](OnCallSchedulesApi.md#CreateOnCallSchedule) | **Post** /on-call-schedules | 
[**GetOnCallSchedule**](OnCallSchedulesApi.md#GetOnCallSchedule) | **Get** /on-call-schedules/{on_call_schedule_id} | 
[**GetOnCallSchedules**](OnCallSchedulesApi.md#GetOnCallSchedules) | **Get** /on-call-schedules | 



## CreateOnCallSchedule

> OnCallSchedule CreateOnCallSchedule(ctx).CreateOnCallScheduleInfo(createOnCallScheduleInfo).Execute()





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
    createOnCallScheduleInfo := *openapiclient.NewCreateOnCallScheduleInfo(openapiclient.OnCallScheduleProviderEnum("OPSGENIE"), "PNZNINN") // CreateOnCallScheduleInfo | The `OnCallSchedule` object to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnCallSchedulesApi.CreateOnCallSchedule(context.Background()).CreateOnCallScheduleInfo(createOnCallScheduleInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnCallSchedulesApi.CreateOnCallSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOnCallSchedule`: OnCallSchedule
    fmt.Fprintf(os.Stdout, "Response from `OnCallSchedulesApi.CreateOnCallSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOnCallScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOnCallScheduleInfo** | [**CreateOnCallScheduleInfo**](CreateOnCallScheduleInfo.md) | The &#x60;OnCallSchedule&#x60; object to be created. | 

### Return type

[**OnCallSchedule**](OnCallSchedule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnCallSchedule

> OnCallSchedule GetOnCallSchedule(ctx, onCallScheduleId).Execute()





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
    onCallScheduleId := "9546209c-42c2-4801-96d7-9ec42df0f59c" // string | The ID of the on_call_schedule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnCallSchedulesApi.GetOnCallSchedule(context.Background(), onCallScheduleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnCallSchedulesApi.GetOnCallSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOnCallSchedule`: OnCallSchedule
    fmt.Fprintf(os.Stdout, "Response from `OnCallSchedulesApi.GetOnCallSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onCallScheduleId** | **string** | The ID of the on_call_schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnCallScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OnCallSchedule**](OnCallSchedule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnCallSchedules

> OnCallScheduleList GetOnCallSchedules(ctx).Execute()





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
    resp, r, err := apiClient.OnCallSchedulesApi.GetOnCallSchedules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnCallSchedulesApi.GetOnCallSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOnCallSchedules`: OnCallScheduleList
    fmt.Fprintf(os.Stdout, "Response from `OnCallSchedulesApi.GetOnCallSchedules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnCallSchedulesRequest struct via the builder pattern


### Return type

[**OnCallScheduleList**](OnCallScheduleList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

