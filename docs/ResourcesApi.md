# \ResourcesApi

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddResourceUser**](ResourcesApi.md#AddResourceUser) | **Post** /resources/{resource_id}/users/{user_id} | 
[**CreateResource**](ResourcesApi.md#CreateResource) | **Post** /resources | 
[**DeleteResource**](ResourcesApi.md#DeleteResource) | **Delete** /resources/{resource_id} | 
[**DeleteResourceUser**](ResourcesApi.md#DeleteResourceUser) | **Delete** /resources/{resource_id}/users/{user_id} | 
[**GetResource**](ResourcesApi.md#GetResource) | **Get** /resources/{resource_id} | 
[**GetResourceMessageChannels**](ResourcesApi.md#GetResourceMessageChannels) | **Get** /resources/{resource_id}/message-channels | 
[**GetResourceReviewerStages**](ResourcesApi.md#GetResourceReviewerStages) | **Get** /resources/{resource_id}/reviewer-stages | 
[**GetResourceReviewers**](ResourcesApi.md#GetResourceReviewers) | **Get** /resources/{resource_id}/reviewers | 
[**GetResourceTags**](ResourcesApi.md#GetResourceTags) | **Get** /resources/{resource_id}/tags | 
[**GetResourceUsers**](ResourcesApi.md#GetResourceUsers) | **Get** /resources/{resource_id}/users | 
[**GetResourceVisibility**](ResourcesApi.md#GetResourceVisibility) | **Get** /resources/{resource_id}/visibility | 
[**GetResources**](ResourcesApi.md#GetResources) | **Get** /resources | 
[**ResourceUserAccessStatusRetrieve**](ResourcesApi.md#ResourceUserAccessStatusRetrieve) | **Get** /resource-user-access-status/{resource_id}/{user_id} | 
[**SetResourceMessageChannels**](ResourcesApi.md#SetResourceMessageChannels) | **Put** /resources/{resource_id}/message-channels | 
[**SetResourceReviewerStages**](ResourcesApi.md#SetResourceReviewerStages) | **Put** /resources/{resource_id}/reviewer-stages | 
[**SetResourceReviewers**](ResourcesApi.md#SetResourceReviewers) | **Put** /resources/{resource_id}/reviewers | 
[**SetResourceVisibility**](ResourcesApi.md#SetResourceVisibility) | **Put** /resources/{resource_id}/visibility | 
[**UpdateResources**](ResourcesApi.md#UpdateResources) | **Put** /resources | 



## AddResourceUser

> ResourceUser AddResourceUser(ctx, resourceId, userId).DurationMinutes(durationMinutes).AccessLevelRemoteId(accessLevelRemoteId).Execute()





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
    resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.
    userId := "f92aa855-cea9-4814-b9d8-f2a60d3e4a06" // string | The ID of the user to add.
    durationMinutes := int32(60) // int32 | The duration for which the resource can be accessed (in minutes). Use 0 to set to indefinite.
    accessLevelRemoteId := "arn:aws:iam::590304332660:role/AdministratorAccess" // string | The remote ID of the access level to grant to this user. If omitted, the default access level remote ID value (empty string) is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.AddResourceUser(context.Background(), resourceId, userId).DurationMinutes(durationMinutes).AccessLevelRemoteId(accessLevelRemoteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.AddResourceUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddResourceUser`: ResourceUser
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.AddResourceUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 
**userId** | **string** | The ID of the user to add. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddResourceUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **durationMinutes** | **int32** | The duration for which the resource can be accessed (in minutes). Use 0 to set to indefinite. | 
 **accessLevelRemoteId** | **string** | The remote ID of the access level to grant to this user. If omitted, the default access level remote ID value (empty string) is used. | 

### Return type

[**ResourceUser**](ResourceUser.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateResource

> Resource CreateResource(ctx).CreateResourceInfo(createResourceInfo).Execute()





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
    createResourceInfo := *openapiclient.NewCreateResourceInfo("mongo-db-prod", openapiclient.ResourceTypeEnum("AWS_IAM_ROLE"), "f454d283-ca87-4a8a-bdbb-df212eca5353") // CreateResourceInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.CreateResource(context.Background()).CreateResourceInfo(createResourceInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.CreateResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResource`: Resource
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.CreateResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createResourceInfo** | [**CreateResourceInfo**](CreateResourceInfo.md) |  | 

### Return type

[**Resource**](Resource.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResource

> DeleteResource(ctx, resourceId).Execute()





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
    resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResourcesApi.DeleteResource(context.Background(), resourceId).Execute()
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


## DeleteResourceUser

> DeleteResourceUser(ctx, resourceId, userId).AccessLevelRemoteId(accessLevelRemoteId).Execute()





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
    resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.
    userId := "f92aa855-cea9-4814-b9d8-f2a60d3e4a06" // string | The ID of a user to remove from this resource.
    accessLevelRemoteId := "arn:aws:iam::590304332660:role/AdministratorAccess" // string | The remote ID of the access level for which this user has direct access. If omitted, the default access level remote ID value (empty string) is assumed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResourcesApi.DeleteResourceUser(context.Background(), resourceId, userId).AccessLevelRemoteId(accessLevelRemoteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.DeleteResourceUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 
**userId** | **string** | The ID of a user to remove from this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessLevelRemoteId** | **string** | The remote ID of the access level for which this user has direct access. If omitted, the default access level remote ID value (empty string) is assumed. | 

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


## GetResource

> Resource GetResource(ctx, resourceId).Execute()





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
    resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.GetResource(context.Background(), resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.GetResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResource`: Resource
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.GetResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Resource**](Resource.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    openapiclient "github.com/opalsecurity/opal-go"
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


## GetResourceReviewerStages

> []ReviewerStage GetResourceReviewerStages(ctx, resourceId).Execute()





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
    resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.GetResourceReviewerStages(context.Background(), resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.GetResourceReviewerStages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceReviewerStages`: []ReviewerStage
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.GetResourceReviewerStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceReviewerStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ReviewerStage**](ReviewerStage.md)

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
    openapiclient "github.com/opalsecurity/opal-go"
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
    openapiclient "github.com/opalsecurity/opal-go"
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


## GetResourceUsers

> ResourceAccessUserList GetResourceUsers(ctx, resourceId).Limit(limit).Execute()





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
    resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.
    limit := int32(200) // int32 | Limit the number of results returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.GetResourceUsers(context.Background(), resourceId).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.GetResourceUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceUsers`: ResourceAccessUserList
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.GetResourceUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 

### Return type

[**ResourceAccessUserList**](ResourceAccessUserList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceVisibility

> VisibilityInfo GetResourceVisibility(ctx, resourceId).Execute()





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
    resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.GetResourceVisibility(context.Background(), resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.GetResourceVisibility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceVisibility`: VisibilityInfo
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.GetResourceVisibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VisibilityInfo**](VisibilityInfo.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResources

> PaginatedResourcesList GetResources(ctx).Cursor(cursor).PageSize(pageSize).ResourceTypeFilter(resourceTypeFilter).ResourceIds(resourceIds).ResourceName(resourceName).ParentResourceId(parentResourceId).Execute()





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
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    pageSize := int32(200) // int32 | Number of results to return per page. Default is 200. (optional)
    resourceTypeFilter := openapiclient.ResourceTypeEnum("AWS_IAM_ROLE") // ResourceTypeEnum | The resource type to filter by. (optional)
    resourceIds := []string{"1b978423-db0a-4037-a4cf-f79c60cb67b3"} // []string | The resource ids to filter by. (optional)
    resourceName := "example-name" // string | Resource name. (optional)
    parentResourceId := "["4baf8423-db0a-4037-a4cf-f79c60cb67a5"]" // string | The parent resource id to filter by. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.GetResources(context.Background()).Cursor(cursor).PageSize(pageSize).ResourceTypeFilter(resourceTypeFilter).ResourceIds(resourceIds).ResourceName(resourceName).ParentResourceId(parentResourceId).Execute()
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
 **resourceIds** | **[]string** | The resource ids to filter by. | 
 **resourceName** | **string** | Resource name. | 
 **parentResourceId** | **string** | The parent resource id to filter by. | 

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
    openapiclient "github.com/opalsecurity/opal-go"
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


## SetResourceMessageChannels

> []string SetResourceMessageChannels(ctx, resourceId).MessageChannelIDList(messageChannelIDList).Execute()





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


## SetResourceReviewerStages

> []ReviewerStage SetResourceReviewerStages(ctx, resourceId).ReviewerStageList(reviewerStageList).Execute()





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
    resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.
    reviewerStageList := *openapiclient.NewReviewerStageList([]openapiclient.ReviewerStage{*openapiclient.NewReviewerStage(false, "AND", []string{"OwnerIds_example"})}) // ReviewerStageList | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.SetResourceReviewerStages(context.Background(), resourceId).ReviewerStageList(reviewerStageList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.SetResourceReviewerStages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetResourceReviewerStages`: []ReviewerStage
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.SetResourceReviewerStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetResourceReviewerStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewerStageList** | [**ReviewerStageList**](ReviewerStageList.md) |  | 

### Return type

[**[]ReviewerStage**](ReviewerStage.md)

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
    openapiclient "github.com/opalsecurity/opal-go"
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


## SetResourceVisibility

> VisibilityInfo SetResourceVisibility(ctx, resourceId).VisibilityInfo(visibilityInfo).Execute()





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
    resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.
    visibilityInfo := *openapiclient.NewVisibilityInfo(openapiclient.VisibilityTypeEnum("GLOBAL")) // VisibilityInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.SetResourceVisibility(context.Background(), resourceId).VisibilityInfo(visibilityInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.SetResourceVisibility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetResourceVisibility`: VisibilityInfo
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.SetResourceVisibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetResourceVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **visibilityInfo** | [**VisibilityInfo**](VisibilityInfo.md) |  | 

### Return type

[**VisibilityInfo**](VisibilityInfo.md)

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
    openapiclient "github.com/opalsecurity/opal-go"
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

