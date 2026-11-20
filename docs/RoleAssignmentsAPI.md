# \RoleAssignmentsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoleAssignments**](RoleAssignmentsAPI.md#AddRoleAssignments) | **Post** /role-assignments | Add role assignments
[**ListRoleAssignments**](RoleAssignmentsAPI.md#ListRoleAssignments) | **Get** /role-assignments | List role assignments
[**RemoveRoleAssignments**](RoleAssignmentsAPI.md#RemoveRoleAssignments) | **Delete** /role-assignments | Remove role assignments



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


## ListRoleAssignments

> PaginatedRoleAssignmentsList ListRoleAssignments(ctx).EntityIds(entityIds).PrincipalIds(principalIds).EntityType(entityType).PrincipalType(principalType).Cursor(cursor).PageSize(pageSize).Execute()

List role assignments



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
	entityIds := []string{"Inner_example"} // []string | Filter to grants on these entity (resource or group) IDs. Any-match. (optional)
	principalIds := []string{"Inner_example"} // []string | Filter to grants with these principal (user, group, or resource) IDs. Any-match. (optional)
	entityType := openapiclient.RoleAssignmentEntityTypeEnum("GROUP") // RoleAssignmentEntityTypeEnum | Filter by entity type. Omit to include both RESOURCE and GROUP entities. (optional)
	principalType := openapiclient.EntityTypeEnum("GROUP") // EntityTypeEnum | Filter by principal type. Omit to include USER, GROUP, and RESOURCE principals. (optional)
	cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
	pageSize := int32(200) // int32 | Number of results to return per page. Default is 200. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentsAPI.ListRoleAssignments(context.Background()).EntityIds(entityIds).PrincipalIds(principalIds).EntityType(entityType).PrincipalType(principalType).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentsAPI.ListRoleAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoleAssignments`: PaginatedRoleAssignmentsList
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentsAPI.ListRoleAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityIds** | **[]string** | Filter to grants on these entity (resource or group) IDs. Any-match. | 
 **principalIds** | **[]string** | Filter to grants with these principal (user, group, or resource) IDs. Any-match. | 
 **entityType** | [**RoleAssignmentEntityTypeEnum**](RoleAssignmentEntityTypeEnum.md) | Filter by entity type. Omit to include both RESOURCE and GROUP entities. | 
 **principalType** | [**EntityTypeEnum**](EntityTypeEnum.md) | Filter by principal type. Omit to include USER, GROUP, and RESOURCE principals. | 
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. Default is 200. | 

### Return type

[**PaginatedRoleAssignmentsList**](PaginatedRoleAssignmentsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRoleAssignments

> RemoveRoleAssignmentsResponse RemoveRoleAssignments(ctx).RemoveRoleAssignmentsRequest(removeRoleAssignmentsRequest).Execute()

Remove role assignments



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
	removeRoleAssignmentsRequest := *openapiclient.NewRemoveRoleAssignmentsRequest([]openapiclient.RemoveRoleAssignmentInput{*openapiclient.NewRemoveRoleAssignmentInput("EntityId_example", openapiclient.EntityTypeEnum("GROUP"), "PrincipalId_example", openapiclient.EntityTypeEnum("GROUP"))}) // RemoveRoleAssignmentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentsAPI.RemoveRoleAssignments(context.Background()).RemoveRoleAssignmentsRequest(removeRoleAssignmentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentsAPI.RemoveRoleAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveRoleAssignments`: RemoveRoleAssignmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentsAPI.RemoveRoleAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeRoleAssignmentsRequest** | [**RemoveRoleAssignmentsRequest**](RemoveRoleAssignmentsRequest.md) |  | 

### Return type

[**RemoveRoleAssignmentsResponse**](RemoveRoleAssignmentsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

