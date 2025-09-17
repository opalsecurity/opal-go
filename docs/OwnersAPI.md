# \OwnersAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOwner**](OwnersAPI.md#CreateOwner) | **Post** /owners | 
[**DeleteOwner**](OwnersAPI.md#DeleteOwner) | **Delete** /owners/{owner_id} | 
[**GetOwner**](OwnersAPI.md#GetOwner) | **Get** /owners/{owner_id} | Get owner by ID
[**GetOwnerFromName**](OwnersAPI.md#GetOwnerFromName) | **Get** /owners/name/{owner_name} | 
[**GetOwnerUsers**](OwnersAPI.md#GetOwnerUsers) | **Get** /owners/{owner_id}/users | 
[**GetOwners**](OwnersAPI.md#GetOwners) | **Get** /owners | Get owners
[**SetOwnerUsers**](OwnersAPI.md#SetOwnerUsers) | **Put** /owners/{owner_id}/users | 
[**UpdateOwners**](OwnersAPI.md#UpdateOwners) | **Put** /owners | 



## CreateOwner

> Owner CreateOwner(ctx).CreateOwnerInfo(createOwnerInfo).Execute()





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
	createOwnerInfo := *openapiclient.NewCreateOwnerInfo("API Owner", []string{"UserIds_example"}) // CreateOwnerInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnersAPI.CreateOwner(context.Background()).CreateOwnerInfo(createOwnerInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnersAPI.CreateOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOwner`: Owner
	fmt.Fprintf(os.Stdout, "Response from `OwnersAPI.CreateOwner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOwnerInfo** | [**CreateOwnerInfo**](CreateOwnerInfo.md) |  | 

### Return type

[**Owner**](Owner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOwner

> DeleteOwner(ctx, ownerId).Execute()





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
	ownerId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OwnersAPI.DeleteOwner(context.Background(), ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnersAPI.DeleteOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerId** | **string** | The ID of the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOwnerRequest struct via the builder pattern


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


## GetOwner

> Owner GetOwner(ctx, ownerId).Execute()

Get owner by ID



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
	ownerId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnersAPI.GetOwner(context.Background(), ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnersAPI.GetOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwner`: Owner
	fmt.Fprintf(os.Stdout, "Response from `OwnersAPI.GetOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerId** | **string** | The ID of the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Owner**](Owner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnerFromName

> Owner GetOwnerFromName(ctx, ownerName).Execute()





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
	ownerName := "MyOwner" // string | The name of the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnersAPI.GetOwnerFromName(context.Background(), ownerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnersAPI.GetOwnerFromName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnerFromName`: Owner
	fmt.Fprintf(os.Stdout, "Response from `OwnersAPI.GetOwnerFromName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerName** | **string** | The name of the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnerFromNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Owner**](Owner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnerUsers

> UserList GetOwnerUsers(ctx, ownerId).Execute()





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
	ownerId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnersAPI.GetOwnerUsers(context.Background(), ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnersAPI.GetOwnerUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnerUsers`: UserList
	fmt.Fprintf(os.Stdout, "Response from `OwnersAPI.GetOwnerUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerId** | **string** | The ID of the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnerUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserList**](UserList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwners

> PaginatedOwnersList GetOwners(ctx).Cursor(cursor).PageSize(pageSize).Name(name).Execute()

Get owners



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
	name := "200" // string | Owner name to filter by. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnersAPI.GetOwners(context.Background()).Cursor(cursor).PageSize(pageSize).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnersAPI.GetOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwners`: PaginatedOwnersList
	fmt.Fprintf(os.Stdout, "Response from `OwnersAPI.GetOwners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. Default is 200. | 
 **name** | **string** | Owner name to filter by. | 

### Return type

[**PaginatedOwnersList**](PaginatedOwnersList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOwnerUsers

> UserList SetOwnerUsers(ctx, ownerId).UserIDList(userIDList).Execute()





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
	ownerId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the owner.
	userIDList := *openapiclient.NewUserIDList([]string{"UserIds_example"}) // UserIDList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnersAPI.SetOwnerUsers(context.Background(), ownerId).UserIDList(userIDList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnersAPI.SetOwnerUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOwnerUsers`: UserList
	fmt.Fprintf(os.Stdout, "Response from `OwnersAPI.SetOwnerUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerId** | **string** | The ID of the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOwnerUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userIDList** | [**UserIDList**](UserIDList.md) |  | 

### Return type

[**UserList**](UserList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOwners

> UpdateOwnerInfoList UpdateOwners(ctx).UpdateOwnerInfoList(updateOwnerInfoList).Execute()





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
	updateOwnerInfoList := *openapiclient.NewUpdateOwnerInfoList([]openapiclient.UpdateOwnerInfo{*openapiclient.NewUpdateOwnerInfo("f454d283-ca87-4a8a-bdbb-df212eca5353")}) // UpdateOwnerInfoList | Owners to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnersAPI.UpdateOwners(context.Background()).UpdateOwnerInfoList(updateOwnerInfoList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnersAPI.UpdateOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOwners`: UpdateOwnerInfoList
	fmt.Fprintf(os.Stdout, "Response from `OwnersAPI.UpdateOwners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOwnerInfoList** | [**UpdateOwnerInfoList**](UpdateOwnerInfoList.md) | Owners to be updated | 

### Return type

[**UpdateOwnerInfoList**](UpdateOwnerInfoList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

