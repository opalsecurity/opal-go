# \AppsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApp**](AppsAPI.md#GetApp) | **Get** /apps/{app_id} | 
[**GetApps**](AppsAPI.md#GetApps) | **Get** /apps | 
[**GetSyncErrors**](AppsAPI.md#GetSyncErrors) | **Get** /sync_errors | 



## GetApp

> App GetApp(ctx, appId).Execute()





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
	appId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetApp(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApp`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**App**](App.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApps

> AppsList GetApps(ctx).AppTypeFilter(appTypeFilter).OwnerFilter(ownerFilter).Execute()





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
	appTypeFilter := []openapiclient.AppTypeEnum{openapiclient.AppTypeEnum("ACTIVE_DIRECTORY")} // []AppTypeEnum | A list of app types to filter by. (optional)
	ownerFilter := "29827fb8-f2dd-4e80-9576-28e31e9934ac" // string | An owner ID to filter by. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetApps(context.Background()).AppTypeFilter(appTypeFilter).OwnerFilter(ownerFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApps`: AppsList
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appTypeFilter** | [**[]AppTypeEnum**](AppTypeEnum.md) | A list of app types to filter by. | 
 **ownerFilter** | **string** | An owner ID to filter by. | 

### Return type

[**AppsList**](AppsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncErrors

> []SyncErrorList GetSyncErrors(ctx).AppId(appId).ResourceId(resourceId).GroupId(groupId).Execute()





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
	appId := "29827fb8-f2dd-4e80-9576-28e31e9934ac" // string | The ID of the app to list sync errors for. (optional)
	resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource to list sync errors for. (optional)
	groupId := "9546209c-42c2-4801-96d7-9ec42df0f59c" // string | The ID of the group to list sync errors for. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetSyncErrors(context.Background()).AppId(appId).ResourceId(resourceId).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetSyncErrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncErrors`: []SyncErrorList
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetSyncErrors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** | The ID of the app to list sync errors for. | 
 **resourceId** | **string** | The ID of the resource to list sync errors for. | 
 **groupId** | **string** | The ID of the group to list sync errors for. | 

### Return type

[**[]SyncErrorList**](SyncErrorList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

