# \IpAclAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIPZone**](IpAclAPI.md#CreateIPZone) | **Post** /ip-zones | Create IP zone
[**DeleteIPZone**](IpAclAPI.md#DeleteIPZone) | **Delete** /ip-zones/{ip_zone_id} | Delete IP zone
[**GetIPACLSettings**](IpAclAPI.md#GetIPACLSettings) | **Get** /ip-acl-settings | Get API IP allowlist settings
[**GetIPZone**](IpAclAPI.md#GetIPZone) | **Get** /ip-zones/{ip_zone_id} | Get IP zone
[**GetIPZones**](IpAclAPI.md#GetIPZones) | **Get** /ip-zones | Get IP zones
[**UpdateIPACLSettings**](IpAclAPI.md#UpdateIPACLSettings) | **Put** /ip-acl-settings | Update API IP allowlist settings
[**UpdateIPZone**](IpAclAPI.md#UpdateIPZone) | **Put** /ip-zones/{ip_zone_id} | Update IP zone



## CreateIPZone

> IpZone CreateIPZone(ctx).CreateIpZoneInfo(createIpZoneInfo).Execute()

Create IP zone



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
	createIpZoneInfo := *openapiclient.NewCreateIpZoneInfo("Office network", []string{"IpAddresses_example"}) // CreateIpZoneInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpAclAPI.CreateIPZone(context.Background()).CreateIpZoneInfo(createIpZoneInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpAclAPI.CreateIPZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIPZone`: IpZone
	fmt.Fprintf(os.Stdout, "Response from `IpAclAPI.CreateIPZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIPZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIpZoneInfo** | [**CreateIpZoneInfo**](CreateIpZoneInfo.md) |  | 

### Return type

[**IpZone**](IpZone.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIPZone

> DeleteIPZone(ctx, ipZoneId).Execute()

Delete IP zone



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
	ipZoneId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the IP zone.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IpAclAPI.DeleteIPZone(context.Background(), ipZoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpAclAPI.DeleteIPZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipZoneId** | **string** | The ID of the IP zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIPZoneRequest struct via the builder pattern


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


## GetIPACLSettings

> IpAclSettings GetIPACLSettings(ctx).Execute()

Get API IP allowlist settings



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
	resp, r, err := apiClient.IpAclAPI.GetIPACLSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpAclAPI.GetIPACLSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIPACLSettings`: IpAclSettings
	fmt.Fprintf(os.Stdout, "Response from `IpAclAPI.GetIPACLSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIPACLSettingsRequest struct via the builder pattern


### Return type

[**IpAclSettings**](IpAclSettings.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIPZone

> IpZone GetIPZone(ctx, ipZoneId).Execute()

Get IP zone



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
	ipZoneId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the IP zone.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpAclAPI.GetIPZone(context.Background(), ipZoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpAclAPI.GetIPZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIPZone`: IpZone
	fmt.Fprintf(os.Stdout, "Response from `IpAclAPI.GetIPZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipZoneId** | **string** | The ID of the IP zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIPZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IpZone**](IpZone.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIPZones

> IpZoneList GetIPZones(ctx).Execute()

Get IP zones



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
	resp, r, err := apiClient.IpAclAPI.GetIPZones(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpAclAPI.GetIPZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIPZones`: IpZoneList
	fmt.Fprintf(os.Stdout, "Response from `IpAclAPI.GetIPZones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIPZonesRequest struct via the builder pattern


### Return type

[**IpZoneList**](IpZoneList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIPACLSettings

> IpAclSettings UpdateIPACLSettings(ctx).IpAclSettings(ipAclSettings).Execute()

Update API IP allowlist settings



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
	ipAclSettings := *openapiclient.NewIpAclSettings(false) // IpAclSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpAclAPI.UpdateIPACLSettings(context.Background()).IpAclSettings(ipAclSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpAclAPI.UpdateIPACLSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIPACLSettings`: IpAclSettings
	fmt.Fprintf(os.Stdout, "Response from `IpAclAPI.UpdateIPACLSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIPACLSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipAclSettings** | [**IpAclSettings**](IpAclSettings.md) |  | 

### Return type

[**IpAclSettings**](IpAclSettings.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIPZone

> IpZone UpdateIPZone(ctx, ipZoneId).UpdateIpZoneInfo(updateIpZoneInfo).Execute()

Update IP zone



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
	ipZoneId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the IP zone.
	updateIpZoneInfo := *openapiclient.NewUpdateIpZoneInfo() // UpdateIpZoneInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpAclAPI.UpdateIPZone(context.Background(), ipZoneId).UpdateIpZoneInfo(updateIpZoneInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpAclAPI.UpdateIPZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIPZone`: IpZone
	fmt.Fprintf(os.Stdout, "Response from `IpAclAPI.UpdateIPZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipZoneId** | **string** | The ID of the IP zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIPZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIpZoneInfo** | [**UpdateIpZoneInfo**](UpdateIpZoneInfo.md) |  | 

### Return type

[**IpZone**](IpZone.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

