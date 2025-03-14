# \IdpGroupMappingsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIdpGroupMappings**](IdpGroupMappingsAPI.md#DeleteIdpGroupMappings) | **Delete** /idp-group-mappings/{app_resource_id}/{group_id}/ | 
[**GetIdpGroupMappings**](IdpGroupMappingsAPI.md#GetIdpGroupMappings) | **Get** /idp-group-mappings/{app_resource_id} | 
[**UpdateIdpGroupMappings**](IdpGroupMappingsAPI.md#UpdateIdpGroupMappings) | **Put** /idp-group-mappings/{app_resource_id} | 



## DeleteIdpGroupMappings

> DeleteIdpGroupMappings(ctx, appResourceId, groupId).Execute()





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
	appResourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the Okta app.
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdpGroupMappingsAPI.DeleteIdpGroupMappings(context.Background(), appResourceId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdpGroupMappingsAPI.DeleteIdpGroupMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appResourceId** | **string** | The ID of the Okta app. | 
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdpGroupMappingsRequest struct via the builder pattern


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


## GetIdpGroupMappings

> IdpGroupMappingList GetIdpGroupMappings(ctx, appResourceId).Execute()





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
	appResourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the Okta app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdpGroupMappingsAPI.GetIdpGroupMappings(context.Background(), appResourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdpGroupMappingsAPI.GetIdpGroupMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdpGroupMappings`: IdpGroupMappingList
	fmt.Fprintf(os.Stdout, "Response from `IdpGroupMappingsAPI.GetIdpGroupMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appResourceId** | **string** | The ID of the Okta app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpGroupMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdpGroupMappingList**](IdpGroupMappingList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdpGroupMappings

> UpdateIdpGroupMappings(ctx, appResourceId).UpdateIdpGroupMappingsRequest(updateIdpGroupMappingsRequest).Execute()





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
	appResourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the Okta app.
	updateIdpGroupMappingsRequest := *openapiclient.NewUpdateIdpGroupMappingsRequest([]openapiclient.UpdateIdpGroupMappingsRequestMappingsInner{*openapiclient.NewUpdateIdpGroupMappingsRequestMappingsInner()}) // UpdateIdpGroupMappingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdpGroupMappingsAPI.UpdateIdpGroupMappings(context.Background(), appResourceId).UpdateIdpGroupMappingsRequest(updateIdpGroupMappingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdpGroupMappingsAPI.UpdateIdpGroupMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appResourceId** | **string** | The ID of the Okta app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdpGroupMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIdpGroupMappingsRequest** | [**UpdateIdpGroupMappingsRequest**](UpdateIdpGroupMappingsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

