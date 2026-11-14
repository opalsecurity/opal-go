# \OpalScriptsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOpalScript**](OpalScriptsAPI.md#CreateOpalScript) | **Post** /opal-scripts | Create OpalScript
[**DeleteOpalScript**](OpalScriptsAPI.md#DeleteOpalScript) | **Delete** /opal-scripts/{opal_script_id} | Delete OpalScript
[**GetOpalScript**](OpalScriptsAPI.md#GetOpalScript) | **Get** /opal-scripts/{opal_script_id} | Get OpalScript by ID
[**GetOpalScriptFromName**](OpalScriptsAPI.md#GetOpalScriptFromName) | **Get** /opal-scripts/name/{opal_script_name} | Get OpalScript by name
[**GetOpalScripts**](OpalScriptsAPI.md#GetOpalScripts) | **Get** /opal-scripts | List OpalScripts
[**UpdateOpalScript**](OpalScriptsAPI.md#UpdateOpalScript) | **Put** /opal-scripts/{opal_script_id} | Update OpalScript



## CreateOpalScript

> OpalScript CreateOpalScript(ctx).CreateOpalScriptInfo(createOpalScriptInfo).Execute()

Create OpalScript



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
	createOpalScriptInfo := *openapiclient.NewCreateOpalScriptInfo("MFA check", openapiclient.OpalScriptTypeEnum("REQUEST_REVIEW"), "Script_example") // CreateOpalScriptInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpalScriptsAPI.CreateOpalScript(context.Background()).CreateOpalScriptInfo(createOpalScriptInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpalScriptsAPI.CreateOpalScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOpalScript`: OpalScript
	fmt.Fprintf(os.Stdout, "Response from `OpalScriptsAPI.CreateOpalScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOpalScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOpalScriptInfo** | [**CreateOpalScriptInfo**](CreateOpalScriptInfo.md) |  | 

### Return type

[**OpalScript**](OpalScript.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOpalScript

> DeleteOpalScript(ctx, opalScriptId).Execute()

Delete OpalScript



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
	opalScriptId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the OpalScript.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OpalScriptsAPI.DeleteOpalScript(context.Background(), opalScriptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpalScriptsAPI.DeleteOpalScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opalScriptId** | **string** | The ID of the OpalScript. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOpalScriptRequest struct via the builder pattern


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


## GetOpalScript

> OpalScript GetOpalScript(ctx, opalScriptId).Execute()

Get OpalScript by ID



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
	opalScriptId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the OpalScript.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpalScriptsAPI.GetOpalScript(context.Background(), opalScriptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpalScriptsAPI.GetOpalScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpalScript`: OpalScript
	fmt.Fprintf(os.Stdout, "Response from `OpalScriptsAPI.GetOpalScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opalScriptId** | **string** | The ID of the OpalScript. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpalScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpalScript**](OpalScript.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpalScriptFromName

> OpalScript GetOpalScriptFromName(ctx, opalScriptName).Execute()

Get OpalScript by name



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
	opalScriptName := "MFA check" // string | The name of the OpalScript.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpalScriptsAPI.GetOpalScriptFromName(context.Background(), opalScriptName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpalScriptsAPI.GetOpalScriptFromName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpalScriptFromName`: OpalScript
	fmt.Fprintf(os.Stdout, "Response from `OpalScriptsAPI.GetOpalScriptFromName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opalScriptName** | **string** | The name of the OpalScript. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpalScriptFromNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpalScript**](OpalScript.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpalScripts

> OpalScriptList GetOpalScripts(ctx).Execute()

List OpalScripts



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
	resp, r, err := apiClient.OpalScriptsAPI.GetOpalScripts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpalScriptsAPI.GetOpalScripts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpalScripts`: OpalScriptList
	fmt.Fprintf(os.Stdout, "Response from `OpalScriptsAPI.GetOpalScripts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpalScriptsRequest struct via the builder pattern


### Return type

[**OpalScriptList**](OpalScriptList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOpalScript

> OpalScript UpdateOpalScript(ctx, opalScriptId).UpdateOpalScriptInfo(updateOpalScriptInfo).Execute()

Update OpalScript



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
	opalScriptId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the OpalScript.
	updateOpalScriptInfo := *openapiclient.NewUpdateOpalScriptInfo() // UpdateOpalScriptInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpalScriptsAPI.UpdateOpalScript(context.Background(), opalScriptId).UpdateOpalScriptInfo(updateOpalScriptInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpalScriptsAPI.UpdateOpalScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOpalScript`: OpalScript
	fmt.Fprintf(os.Stdout, "Response from `OpalScriptsAPI.UpdateOpalScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opalScriptId** | **string** | The ID of the OpalScript. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOpalScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOpalScriptInfo** | [**UpdateOpalScriptInfo**](UpdateOpalScriptInfo.md) |  | 

### Return type

[**OpalScript**](OpalScript.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

