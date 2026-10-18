# \RoleAssignmentsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoleAssignments**](RoleAssignmentsAPI.md#AddRoleAssignments) | **Post** /role-assignments | Add role assignments



## AddRoleAssignments

> AddRoleAssignmentsResponse AddRoleAssignments(ctx).AddRoleAssignmentsRequest(addRoleAssignmentsRequest).Execute()

Add role assignments



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
	addRoleAssignmentsRequest := *openapiclient.NewAddRoleAssignmentsRequest([]openapiclient.AddRoleAssignmentInput{*openapiclient.NewAddRoleAssignmentInput("1b978423-db0a-4037-a4cf-f79c60cb67b3", openapiclient.EntityTypeEnum("GROUP"), "29827fb8-f2dd-4e80-9576-28e31e9934ac", openapiclient.EntityTypeEnum("GROUP"))}) // AddRoleAssignmentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentsAPI.AddRoleAssignments(context.Background()).AddRoleAssignmentsRequest(addRoleAssignmentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentsAPI.AddRoleAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRoleAssignments`: AddRoleAssignmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentsAPI.AddRoleAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addRoleAssignmentsRequest** | [**AddRoleAssignmentsRequest**](AddRoleAssignmentsRequest.md) |  | 

### Return type

[**AddRoleAssignmentsResponse**](AddRoleAssignmentsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

