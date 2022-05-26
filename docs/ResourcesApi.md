# \ResourcesApi

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteResource**](ResourcesApi.md#DeleteResource) | **Delete** /resources/{resource_id} | 
[**GetResourceMessageChannels**](ResourcesApi.md#GetResourceMessageChannels) | **Get** /resources/{resource_id}/message-channels | 
[**GetResourceReviewers**](ResourcesApi.md#GetResourceReviewers) | **Get** /resources/{resource_id}/reviewers | 
[**GetResourceTags**](ResourcesApi.md#GetResourceTags) | **Get** /resources/{resource_id}/tags | 
[**GetResources**](ResourcesApi.md#GetResources) | **Get** /resources | 
[**ResourceUserAccessStatusRetrieve**](ResourcesApi.md#ResourceUserAccessStatusRetrieve) | **Get** /resource-user-access-status/{resource_id}/{user_id} | 
[**ResourceUsers**](ResourcesApi.md#ResourceUsers) | **Get** /resource-users | 
[**SetResourceMessageChannels**](ResourcesApi.md#SetResourceMessageChannels) | **Put** /resources/{resource_id}/message-channels | 
[**SetResourceReviewers**](ResourcesApi.md#SetResourceReviewers) | **Put** /resources/{resource_id}/reviewers | 
[**UpdateResources**](ResourcesApi.md#UpdateResources) | **Put** /resources | 



## DeleteResource

> DeleteResource(ctx, resourceId).Execute()





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
    resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.DeleteResource(context.Background(), resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.DeleteResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceRequest struct via the builder pattern


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


## GetResourceMessageChannels

> MessageChannelList GetResourceMessageChannels(ctx, resourceId).Execute()





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
    resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.GetResourceMessageChannels(context.Background(), resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.GetResourceMessageChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceMessageChannels`: MessageChannelList
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.GetResourceMessageChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceMessageChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetResourceReviewers

> []string GetResourceReviewers(ctx, resourceId).Execute()





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
    resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.GetResourceReviewers(context.Background(), resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.GetResourceReviewers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceReviewers`: []string
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.GetResourceReviewers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceReviewersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceTags

> TagsList GetResourceTags(ctx, resourceId).Execute()





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
    resourceId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the resource whose tags to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.GetResourceTags(context.Background(), resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.GetResourceTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceTags`: TagsList
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.GetResourceTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource whose tags to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagsList**](TagsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResources

> PaginatedResourcesList GetResources(ctx).Cursor(cursor).PageSize(pageSize).ResourceTypeFilter(resourceTypeFilter).Execute()





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
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    pageSize := int32(200) // int32 | Number of results to return per page. Default is 200. (optional)
    resourceTypeFilter := openapiclient.ResourceTypeEnum("AWS_IAM_ROLE") // ResourceTypeEnum | The resource type to filter by. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.GetResources(context.Background()).Cursor(cursor).PageSize(pageSize).ResourceTypeFilter(resourceTypeFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.GetResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResources`: PaginatedResourcesList
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.GetResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. Default is 200. | 
 **resourceTypeFilter** | [**ResourceTypeEnum**](ResourceTypeEnum.md) | The resource type to filter by. | 

### Return type

[**PaginatedResourcesList**](PaginatedResourcesList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceUserAccessStatusRetrieve

> ResourceUserAccessStatus ResourceUserAccessStatusRetrieve(ctx, resourceId, userId).AccessLevelRemoteId(accessLevelRemoteId).Cursor(cursor).PageSize(pageSize).Execute()





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
    resourceId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the resource.
    userId := "29827fb8-f2dd-4e80-9576-28e31e9934ac" // string | The ID of the user.
    accessLevelRemoteId := "arn:aws:iam::590304332660:role/AdministratorAccess" // string | The remote ID of the access level that you wish to query for the resource. If omitted, the default access level remote ID value (empty string) is used. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    pageSize := int32(200) // int32 | Number of results to return per page. Default is 200. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.ResourceUserAccessStatusRetrieve(context.Background(), resourceId, userId).AccessLevelRemoteId(accessLevelRemoteId).Cursor(cursor).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.ResourceUserAccessStatusRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResourceUserAccessStatusRetrieve`: ResourceUserAccessStatus
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.ResourceUserAccessStatusRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 
**userId** | **string** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceUserAccessStatusRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessLevelRemoteId** | **string** | The remote ID of the access level that you wish to query for the resource. If omitted, the default access level remote ID value (empty string) is used. | 
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. Default is 200. | 

### Return type

[**ResourceUserAccessStatus**](ResourceUserAccessStatus.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceUsers

> PaginatedResourceUserList ResourceUsers(ctx).ResourceId(resourceId).AccessLevelRemoteId(accessLevelRemoteId).Cursor(cursor).PageSize(pageSize).Execute()





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
    resourceId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the resource.
    accessLevelRemoteId := "arn:aws:iam::590304332660:role/AdministratorAccess" // string | The remote ID of the access level that you wish to query for the resource. If omitted, the default access level remote ID value (empty string) is used. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    pageSize := int32(200) // int32 | Number of results to return per page. Default is 200. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.ResourceUsers(context.Background()).ResourceId(resourceId).AccessLevelRemoteId(accessLevelRemoteId).Cursor(cursor).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.ResourceUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResourceUsers`: PaginatedResourceUserList
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.ResourceUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResourceUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceId** | **string** | The ID of the resource. | 
 **accessLevelRemoteId** | **string** | The remote ID of the access level that you wish to query for the resource. If omitted, the default access level remote ID value (empty string) is used. | 
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. Default is 200. | 

### Return type

[**PaginatedResourceUserList**](PaginatedResourceUserList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetResourceMessageChannels

> []string SetResourceMessageChannels(ctx, resourceId).MessageChannelIDList(messageChannelIDList).Execute()





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
    resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.
    messageChannelIDList := *openapiclient.NewMessageChannelIDList([]string{"MessageChannelIds_example"}) // MessageChannelIDList | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.SetResourceMessageChannels(context.Background(), resourceId).MessageChannelIDList(messageChannelIDList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.SetResourceMessageChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetResourceMessageChannels`: []string
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.SetResourceMessageChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetResourceMessageChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **messageChannelIDList** | [**MessageChannelIDList**](MessageChannelIDList.md) |  | 

### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetResourceReviewers

> []string SetResourceReviewers(ctx, resourceId).ReviewerIDList(reviewerIDList).Execute()





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
    resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.
    reviewerIDList := *openapiclient.NewReviewerIDList([]string{"ReviewerIds_example"}) // ReviewerIDList | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.SetResourceReviewers(context.Background(), resourceId).ReviewerIDList(reviewerIDList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.SetResourceReviewers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetResourceReviewers`: []string
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.SetResourceReviewers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetResourceReviewersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewerIDList** | [**ReviewerIDList**](ReviewerIDList.md) |  | 

### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResources

> UpdateResourceInfoList UpdateResources(ctx).UpdateResourceInfoList(updateResourceInfoList).Execute()





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
    updateResourceInfoList := *openapiclient.NewUpdateResourceInfoList([]openapiclient.UpdateResourceInfo{*openapiclient.NewUpdateResourceInfo("f454d283-ca87-4a8a-bdbb-df212eca5353")}) // UpdateResourceInfoList | Resources to be updated

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.UpdateResources(context.Background()).UpdateResourceInfoList(updateResourceInfoList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.UpdateResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateResources`: UpdateResourceInfoList
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.UpdateResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateResourceInfoList** | [**UpdateResourceInfoList**](UpdateResourceInfoList.md) | Resources to be updated | 

### Return type

[**UpdateResourceInfoList**](UpdateResourceInfoList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

