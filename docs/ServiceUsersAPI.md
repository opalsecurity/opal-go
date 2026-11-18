# \ServiceUsersAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceUser**](ServiceUsersAPI.md#CreateServiceUser) | **Post** /service-users | Create service user



## CreateServiceUser

> User CreateServiceUser(ctx).CreateServiceUserRequest(createServiceUserRequest).Execute()

Create service user



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
	createServiceUserRequest := *openapiclient.NewCreateServiceUserRequest("paladin-prod", "7c86c85d-0651-43e2-a748-d69d658418e8") // CreateServiceUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceUsersAPI.CreateServiceUser(context.Background()).CreateServiceUserRequest(createServiceUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceUsersAPI.CreateServiceUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceUser`: User
	fmt.Fprintf(os.Stdout, "Response from `ServiceUsersAPI.CreateServiceUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServiceUserRequest** | [**CreateServiceUserRequest**](CreateServiceUserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

