# \IdpUserAttributeMappingsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdpUserAttributeMapping**](IdpUserAttributeMappingsAPI.md#CreateIdpUserAttributeMapping) | **Post** /apps/{app_id}/user-attribute-mappings | Create IDP user attribute mapping
[**DeleteIdpUserAttributeMapping**](IdpUserAttributeMappingsAPI.md#DeleteIdpUserAttributeMapping) | **Delete** /apps/{app_id}/user-attribute-mappings/{mapping_id} | Delete IDP user attribute mapping
[**GetIdpUserAttributeMapping**](IdpUserAttributeMappingsAPI.md#GetIdpUserAttributeMapping) | **Get** /apps/{app_id}/user-attribute-mappings/{mapping_id} | Get IDP user attribute mapping
[**GetIdpUserAttributeMappings**](IdpUserAttributeMappingsAPI.md#GetIdpUserAttributeMappings) | **Get** /apps/{app_id}/user-attribute-mappings | List IDP user attribute mappings



## CreateIdpUserAttributeMapping

> IdpUserAttributeMapping CreateIdpUserAttributeMapping(ctx, appId).CreateIdpUserAttributeMappingInfo(createIdpUserAttributeMappingInfo).Execute()

Create IDP user attribute mapping



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
	appId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the identity provider or HRIS app.
	createIdpUserAttributeMappingInfo := *openapiclient.NewCreateIdpUserAttributeMappingInfo("managerEmail", openapiclient.IdpUserAttributeUseAsEnum("CUSTOM")) // CreateIdpUserAttributeMappingInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdpUserAttributeMappingsAPI.CreateIdpUserAttributeMapping(context.Background(), appId).CreateIdpUserAttributeMappingInfo(createIdpUserAttributeMappingInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdpUserAttributeMappingsAPI.CreateIdpUserAttributeMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdpUserAttributeMapping`: IdpUserAttributeMapping
	fmt.Fprintf(os.Stdout, "Response from `IdpUserAttributeMappingsAPI.CreateIdpUserAttributeMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of the identity provider or HRIS app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdpUserAttributeMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIdpUserAttributeMappingInfo** | [**CreateIdpUserAttributeMappingInfo**](CreateIdpUserAttributeMappingInfo.md) |  | 

### Return type

[**IdpUserAttributeMapping**](IdpUserAttributeMapping.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdpUserAttributeMapping

> DeleteIdpUserAttributeMapping(ctx, appId, mappingId).Execute()

Delete IDP user attribute mapping



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
	appId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the identity provider or HRIS app.
	mappingId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the user attribute mapping.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdpUserAttributeMappingsAPI.DeleteIdpUserAttributeMapping(context.Background(), appId, mappingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdpUserAttributeMappingsAPI.DeleteIdpUserAttributeMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of the identity provider or HRIS app. | 
**mappingId** | **string** | The ID of the user attribute mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdpUserAttributeMappingRequest struct via the builder pattern


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


## GetIdpUserAttributeMapping

> IdpUserAttributeMapping GetIdpUserAttributeMapping(ctx, appId, mappingId).Execute()

Get IDP user attribute mapping



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
	appId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the identity provider or HRIS app.
	mappingId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the user attribute mapping.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdpUserAttributeMappingsAPI.GetIdpUserAttributeMapping(context.Background(), appId, mappingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdpUserAttributeMappingsAPI.GetIdpUserAttributeMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdpUserAttributeMapping`: IdpUserAttributeMapping
	fmt.Fprintf(os.Stdout, "Response from `IdpUserAttributeMappingsAPI.GetIdpUserAttributeMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of the identity provider or HRIS app. | 
**mappingId** | **string** | The ID of the user attribute mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpUserAttributeMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdpUserAttributeMapping**](IdpUserAttributeMapping.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpUserAttributeMappings

> IdpUserAttributeMappingList GetIdpUserAttributeMappings(ctx, appId).Execute()

List IDP user attribute mappings



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
	appId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the identity provider or HRIS app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdpUserAttributeMappingsAPI.GetIdpUserAttributeMappings(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdpUserAttributeMappingsAPI.GetIdpUserAttributeMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdpUserAttributeMappings`: IdpUserAttributeMappingList
	fmt.Fprintf(os.Stdout, "Response from `IdpUserAttributeMappingsAPI.GetIdpUserAttributeMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of the identity provider or HRIS app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpUserAttributeMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdpUserAttributeMappingList**](IdpUserAttributeMappingList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

