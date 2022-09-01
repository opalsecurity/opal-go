# \TagsApi

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupTag**](TagsApi.md#AddGroupTag) | **Post** /tags/{tag_id}/groups/{group_id} | 
[**AddResourceTag**](TagsApi.md#AddResourceTag) | **Post** /tags/{tag_id}/resources/{resource_id} | 
[**AddUserTag**](TagsApi.md#AddUserTag) | **Post** /tags/{tag_id}/users/{user_id} | 
[**CreateTag**](TagsApi.md#CreateTag) | **Post** /tag | 
[**GetTag**](TagsApi.md#GetTag) | **Get** /tag | 
[**GetTags**](TagsApi.md#GetTags) | **Get** /tags | 
[**RemoveGroupTag**](TagsApi.md#RemoveGroupTag) | **Delete** /tags/{tag_id}/groups/{group_id} | 
[**RemoveResourceTag**](TagsApi.md#RemoveResourceTag) | **Delete** /tags/{tag_id}/resources/{resource_id} | 
[**RemoveUserTag**](TagsApi.md#RemoveUserTag) | **Delete** /tags/{tag_id}/users/{user_id} | 



## AddGroupTag

> AddGroupTag(ctx, tagId, groupId).Execute()





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
    tagId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the tag to apply.
    groupId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the group to apply the tag to.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.AddGroupTag(context.Background(), tagId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.AddGroupTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | The ID of the tag to apply. | 
**groupId** | **string** | The ID of the group to apply the tag to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupTagRequest struct via the builder pattern


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


## AddResourceTag

> AddResourceTag(ctx, tagId, resourceId).Execute()





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
    tagId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the tag to apply.
    resourceId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the resource to apply the tag to.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.AddResourceTag(context.Background(), tagId, resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.AddResourceTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | The ID of the tag to apply. | 
**resourceId** | **string** | The ID of the resource to apply the tag to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddResourceTagRequest struct via the builder pattern


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


## AddUserTag

> AddUserTag(ctx, tagId, userId).Execute()





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
    tagId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the tag to apply.
    userId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the user to apply the tag to.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.AddUserTag(context.Background(), tagId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.AddUserTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | The ID of the tag to apply. | 
**userId** | **string** | The ID of the user to apply the tag to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserTagRequest struct via the builder pattern


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


## CreateTag

> Tag CreateTag(ctx).TagKey(tagKey).TagValue(tagValue).AdminOwnerId(adminOwnerId).Execute()





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
    tagKey := "api-scope" // string | The key of the tag to create.
    tagValue := "production" // string | The value of the tag to create. (optional)
    adminOwnerId := "f92aa855-cea9-4814-b9d8-f2a60d3e4a06" // string | The ID of the owner that manages the tag. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.CreateTag(context.Background()).TagKey(tagKey).TagValue(tagValue).AdminOwnerId(adminOwnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.CreateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTag`: Tag
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagKey** | **string** | The key of the tag to create. | 
 **tagValue** | **string** | The value of the tag to create. | 
 **adminOwnerId** | **string** | The ID of the owner that manages the tag. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTag

> Tag GetTag(ctx).TagKey(tagKey).TagValue(tagValue).Execute()





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
    tagKey := "api-scope" // string | The key of the tag to get.
    tagValue := "production" // string | The value of the tag to get. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.GetTag(context.Background()).TagKey(tagKey).TagValue(tagValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTag`: Tag
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagKey** | **string** | The key of the tag to get. | 
 **tagValue** | **string** | The value of the tag to get. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> PaginatedTagsList GetTags(ctx).Cursor(cursor).PageSize(pageSize).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.GetTags(context.Background()).Cursor(cursor).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTags`: PaginatedTagsList
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. Default is 200. | 

### Return type

[**PaginatedTagsList**](PaginatedTagsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveGroupTag

> RemoveGroupTag(ctx, tagId, groupId).Execute()





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
    tagId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the tag to remove.
    groupId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the group to remove the tag from.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.RemoveGroupTag(context.Background(), tagId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.RemoveGroupTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | The ID of the tag to remove. | 
**groupId** | **string** | The ID of the group to remove the tag from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupTagRequest struct via the builder pattern


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


## RemoveResourceTag

> RemoveResourceTag(ctx, tagId, resourceId).Execute()





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
    tagId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the tag to remove.
    resourceId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the resource to remove the tag from.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.RemoveResourceTag(context.Background(), tagId, resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.RemoveResourceTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | The ID of the tag to remove. | 
**resourceId** | **string** | The ID of the resource to remove the tag from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveResourceTagRequest struct via the builder pattern


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


## RemoveUserTag

> RemoveUserTag(ctx, tagId, userId).Execute()





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
    tagId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the tag to remove.
    userId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the user to remove the tag from.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.RemoveUserTag(context.Background(), tagId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.RemoveUserTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | The ID of the tag to remove. | 
**userId** | **string** | The ID of the user to remove the tag from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserTagRequest struct via the builder pattern


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

