# \RequestTemplatesAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRequestTemplate**](RequestTemplatesAPI.md#CreateRequestTemplate) | **Post** /request-templates | 
[**DeleteRequestTemplate**](RequestTemplatesAPI.md#DeleteRequestTemplate) | **Delete** /request-templates/{request_template_id} | 
[**GetRequestTemplate**](RequestTemplatesAPI.md#GetRequestTemplate) | **Get** /request-templates/{request_template_id} | 
[**GetRequestTemplates**](RequestTemplatesAPI.md#GetRequestTemplates) | **Get** /request-templates | 
[**UpdateRequestTemplate**](RequestTemplatesAPI.md#UpdateRequestTemplate) | **Put** /request-templates | 



## CreateRequestTemplate

> RequestTemplate CreateRequestTemplate(ctx).CreateRequestTemplateInfo(createRequestTemplateInfo).Execute()





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
	createRequestTemplateInfo := *openapiclient.NewCreateRequestTemplateInfo("Production access questions") // CreateRequestTemplateInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestTemplatesAPI.CreateRequestTemplate(context.Background()).CreateRequestTemplateInfo(createRequestTemplateInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestTemplatesAPI.CreateRequestTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRequestTemplate`: RequestTemplate
	fmt.Fprintf(os.Stdout, "Response from `RequestTemplatesAPI.CreateRequestTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRequestTemplateInfo** | [**CreateRequestTemplateInfo**](CreateRequestTemplateInfo.md) |  | 

### Return type

[**RequestTemplate**](RequestTemplate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRequestTemplate

> DeleteRequestTemplate(ctx, requestTemplateId).Execute()





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
	requestTemplateId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the request template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestTemplatesAPI.DeleteRequestTemplate(context.Background(), requestTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestTemplatesAPI.DeleteRequestTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestTemplateId** | **string** | The ID of the request template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequestTemplateRequest struct via the builder pattern


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


## GetRequestTemplate

> RequestTemplate GetRequestTemplate(ctx, requestTemplateId).Execute()





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
	requestTemplateId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the request template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestTemplatesAPI.GetRequestTemplate(context.Background(), requestTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestTemplatesAPI.GetRequestTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequestTemplate`: RequestTemplate
	fmt.Fprintf(os.Stdout, "Response from `RequestTemplatesAPI.GetRequestTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestTemplateId** | **string** | The ID of the request template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestTemplate**](RequestTemplate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestTemplates

> PaginatedRequestTemplateList GetRequestTemplates(ctx).Execute()





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
	resp, r, err := apiClient.RequestTemplatesAPI.GetRequestTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestTemplatesAPI.GetRequestTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequestTemplates`: PaginatedRequestTemplateList
	fmt.Fprintf(os.Stdout, "Response from `RequestTemplatesAPI.GetRequestTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestTemplatesRequest struct via the builder pattern


### Return type

[**PaginatedRequestTemplateList**](PaginatedRequestTemplateList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRequestTemplate

> RequestTemplate UpdateRequestTemplate(ctx).UpdateRequestTemplateInfo(updateRequestTemplateInfo).Execute()





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
	updateRequestTemplateInfo := *openapiclient.NewUpdateRequestTemplateInfo("7c86c85d-0651-43e2-a748-d69d658418e8") // UpdateRequestTemplateInfo | Request template to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestTemplatesAPI.UpdateRequestTemplate(context.Background()).UpdateRequestTemplateInfo(updateRequestTemplateInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestTemplatesAPI.UpdateRequestTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRequestTemplate`: RequestTemplate
	fmt.Fprintf(os.Stdout, "Response from `RequestTemplatesAPI.UpdateRequestTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequestTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRequestTemplateInfo** | [**UpdateRequestTemplateInfo**](UpdateRequestTemplateInfo.md) | Request template to be updated | 

### Return type

[**RequestTemplate**](RequestTemplate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

