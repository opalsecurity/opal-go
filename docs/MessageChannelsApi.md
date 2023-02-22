# \MessageChannelsApi

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessageChannel**](MessageChannelsApi.md#CreateMessageChannel) | **Post** /message-channels | 
[**GetMessageChannel**](MessageChannelsApi.md#GetMessageChannel) | **Get** /message-channels/{message_channel_id} | 
[**GetMessageChannels**](MessageChannelsApi.md#GetMessageChannels) | **Get** /message-channels | 



## CreateMessageChannel

> MessageChannel CreateMessageChannel(ctx).CreateMessageChannelInfo(createMessageChannelInfo).Execute()





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
    createMessageChannelInfo := *openapiclient.NewCreateMessageChannelInfo(openapiclient.MessageChannelProviderEnum("SLACK"), "C03FJR97276") // CreateMessageChannelInfo | The `MessageChannel` object to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageChannelsApi.CreateMessageChannel(context.Background()).CreateMessageChannelInfo(createMessageChannelInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageChannelsApi.CreateMessageChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMessageChannel`: MessageChannel
    fmt.Fprintf(os.Stdout, "Response from `MessageChannelsApi.CreateMessageChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMessageChannelInfo** | [**CreateMessageChannelInfo**](CreateMessageChannelInfo.md) | The &#x60;MessageChannel&#x60; object to be created. | 

### Return type

[**MessageChannel**](MessageChannel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageChannel

> MessageChannel GetMessageChannel(ctx, messageChannelId).Execute()





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
    messageChannelId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the message_channel.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageChannelsApi.GetMessageChannel(context.Background(), messageChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageChannelsApi.GetMessageChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessageChannel`: MessageChannel
    fmt.Fprintf(os.Stdout, "Response from `MessageChannelsApi.GetMessageChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageChannelId** | **string** | The ID of the message_channel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageChannel**](MessageChannel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageChannels

> MessageChannelList GetMessageChannels(ctx).Execute()





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
    resp, r, err := apiClient.MessageChannelsApi.GetMessageChannels(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageChannelsApi.GetMessageChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessageChannels`: MessageChannelList
    fmt.Fprintf(os.Stdout, "Response from `MessageChannelsApi.GetMessageChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageChannelsRequest struct via the builder pattern


### Return type

[**MessageChannelList**](MessageChannelList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

