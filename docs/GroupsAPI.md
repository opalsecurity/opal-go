# \GroupsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupContainingGroup**](GroupsAPI.md#AddGroupContainingGroup) | **Post** /groups/{group_id}/containing-groups | 
[**AddGroupResource**](GroupsAPI.md#AddGroupResource) | **Post** /groups/{group_id}/resources/{resource_id} | 
[**AddGroupUser**](GroupsAPI.md#AddGroupUser) | **Post** /groups/{group_id}/users/{user_id} | 
[**CheckGroupsRequestable**](GroupsAPI.md#CheckGroupsRequestable) | **Post** /groups/requestable/check | Check requestability of many groups
[**CreateGroup**](GroupsAPI.md#CreateGroup) | **Post** /groups | 
[**DeleteGroup**](GroupsAPI.md#DeleteGroup) | **Delete** /groups/{group_id} | 
[**DeleteGroupCustomAttributeByKey**](GroupsAPI.md#DeleteGroupCustomAttributeByKey) | **Delete** /groups/{group_id}/custom-attributes/{attribute_key} | 
[**DeleteGroupUser**](GroupsAPI.md#DeleteGroupUser) | **Delete** /groups/{group_id}/users/{user_id} | 
[**GetGroup**](GroupsAPI.md#GetGroup) | **Get** /groups/{group_id} | Get group by ID
[**GetGroupAccessLevels**](GroupsAPI.md#GetGroupAccessLevels) | **Get** /groups/{group_id}/access_levels | Get group access levels
[**GetGroupBreakGlassUsers**](GroupsAPI.md#GetGroupBreakGlassUsers) | **Get** /groups/{group_id}/break-glass-users | 
[**GetGroupConfigurationTemplate**](GroupsAPI.md#GetGroupConfigurationTemplate) | **Get** /groups/{group_id}/configuration-template | Get the ID of the configuration template linked to a group
[**GetGroupContainingGroup**](GroupsAPI.md#GetGroupContainingGroup) | **Get** /groups/{group_id}/containing-groups/{containing_group_id} | Get nested group by ID
[**GetGroupContainingGroups**](GroupsAPI.md#GetGroupContainingGroups) | **Get** /groups/{group_id}/containing-groups | Get nested groups
[**GetGroupCustomAttributes**](GroupsAPI.md#GetGroupCustomAttributes) | **Get** /groups/{group_id}/custom-attributes | 
[**GetGroupGroups**](GroupsAPI.md#GetGroupGroups) | **Get** /groups/{group_id}/groups | Get member groups
[**GetGroupMessageChannels**](GroupsAPI.md#GetGroupMessageChannels) | **Get** /groups/{group_id}/message-channels | 
[**GetGroupOnCallSchedules**](GroupsAPI.md#GetGroupOnCallSchedules) | **Get** /groups/{group_id}/on-call-schedules | 
[**GetGroupRequestableToUser**](GroupsAPI.md#GetGroupRequestableToUser) | **Get** /groups/{group_id}/requestable/users/{user_id} | 
[**GetGroupResources**](GroupsAPI.md#GetGroupResources) | **Get** /groups/{group_id}/resources | 
[**GetGroupReviewerStages**](GroupsAPI.md#GetGroupReviewerStages) | **Get** /groups/{group_id}/reviewer-stages | 
[**GetGroupReviewers**](GroupsAPI.md#GetGroupReviewers) | **Get** /groups/{group_id}/reviewers | 
[**GetGroupTags**](GroupsAPI.md#GetGroupTags) | **Get** /groups/{group_id}/tags | 
[**GetGroupUsers**](GroupsAPI.md#GetGroupUsers) | **Get** /groups/{group_id}/users | 
[**GetGroupVisibility**](GroupsAPI.md#GetGroupVisibility) | **Get** /groups/{group_id}/visibility | 
[**GetGroupVisibleToUser**](GroupsAPI.md#GetGroupVisibleToUser) | **Get** /groups/{group_id}/visible/users/{user_id} | 
[**GetGroups**](GroupsAPI.md#GetGroups) | **Get** /groups | Get groups
[**GetUserGroups**](GroupsAPI.md#GetUserGroups) | **Get** /groups/users/{user_id} | 
[**PatchGroupCustomAttributes**](GroupsAPI.md#PatchGroupCustomAttributes) | **Patch** /groups/{group_id}/custom-attributes | 
[**RemoveGroupContainingGroup**](GroupsAPI.md#RemoveGroupContainingGroup) | **Delete** /groups/{group_id}/containing-groups/{containing_group_id} | 
[**SetGroupBreakGlassUsers**](GroupsAPI.md#SetGroupBreakGlassUsers) | **Put** /groups/{group_id}/break-glass-users | 
[**SetGroupCustomAttributes**](GroupsAPI.md#SetGroupCustomAttributes) | **Put** /groups/{group_id}/custom-attributes | 
[**SetGroupMessageChannels**](GroupsAPI.md#SetGroupMessageChannels) | **Put** /groups/{group_id}/message-channels | 
[**SetGroupOnCallSchedules**](GroupsAPI.md#SetGroupOnCallSchedules) | **Put** /groups/{group_id}/on-call-schedules | 
[**SetGroupResources**](GroupsAPI.md#SetGroupResources) | **Put** /groups/{group_id}/resources | 
[**SetGroupReviewerStages**](GroupsAPI.md#SetGroupReviewerStages) | **Put** /groups/{group_id}/reviewer-stages | 
[**SetGroupReviewers**](GroupsAPI.md#SetGroupReviewers) | **Put** /groups/{group_id}/reviewers | 
[**SetGroupVisibility**](GroupsAPI.md#SetGroupVisibility) | **Put** /groups/{group_id}/visibility | 
[**UpdateGroupConfigurationTemplate**](GroupsAPI.md#UpdateGroupConfigurationTemplate) | **Put** /groups/{group_id}/configuration-template | Set the ID of the configuration template linked to a group
[**UpdateGroupUser**](GroupsAPI.md#UpdateGroupUser) | **Put** /groups/{group_id}/users/{user_id} | 
[**UpdateGroups**](GroupsAPI.md#UpdateGroups) | **Put** /groups | 



## AddGroupContainingGroup

> GroupContainingGroup AddGroupContainingGroup(ctx, groupId).GroupContainingGroup(groupContainingGroup).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	groupContainingGroup := *openapiclient.NewGroupContainingGroup("f454d283-ca87-4a8a-bdbb-df212eca5353") // GroupContainingGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AddGroupContainingGroup(context.Background(), groupId).GroupContainingGroup(groupContainingGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddGroupContainingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroupContainingGroup`: GroupContainingGroup
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AddGroupContainingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupContainingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupContainingGroup** | [**GroupContainingGroup**](GroupContainingGroup.md) |  | 

### Return type

[**GroupContainingGroup**](GroupContainingGroup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddGroupResource

> GroupResource AddGroupResource(ctx, groupId, resourceId).AccessLevelRemoteId(accessLevelRemoteId).AddGroupResourceRequest(addGroupResourceRequest).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	resourceId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the resource.
	accessLevelRemoteId := "arn:aws:iam::590304332660:role/AdministratorAccess" // string | The remote ID of the access level to grant to this user. If omitted, the default access level remote ID value (empty string) is used. (optional)
	addGroupResourceRequest := *openapiclient.NewAddGroupResourceRequest() // AddGroupResourceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AddGroupResource(context.Background(), groupId, resourceId).AccessLevelRemoteId(accessLevelRemoteId).AddGroupResourceRequest(addGroupResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddGroupResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroupResource`: GroupResource
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AddGroupResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 
**resourceId** | **string** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessLevelRemoteId** | **string** | The remote ID of the access level to grant to this user. If omitted, the default access level remote ID value (empty string) is used. | 
 **addGroupResourceRequest** | [**AddGroupResourceRequest**](AddGroupResourceRequest.md) |  | 

### Return type

[**GroupResource**](GroupResource.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddGroupUser

> GroupUser AddGroupUser(ctx, groupId, userId).DurationMinutes(durationMinutes).AccessLevelRemoteId(accessLevelRemoteId).AddGroupUserRequest(addGroupUserRequest).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	userId := "f92aa855-cea9-4814-b9d8-f2a60d3e4a06" // string | The ID of the user to add.
	durationMinutes := int32(60) // int32 | The duration for which the group can be accessed (in minutes). Use 0 to set to indefinite. (optional)
	accessLevelRemoteId := "arn:aws:iam::590304332660:role/AdministratorAccess" // string | The remote ID of the access level to grant to this user. If omitted, the default access level remote ID value (empty string) is used. (optional)
	addGroupUserRequest := *openapiclient.NewAddGroupUserRequest(int32(60)) // AddGroupUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AddGroupUser(context.Background(), groupId, userId).DurationMinutes(durationMinutes).AccessLevelRemoteId(accessLevelRemoteId).AddGroupUserRequest(addGroupUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddGroupUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroupUser`: GroupUser
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AddGroupUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 
**userId** | **string** | The ID of the user to add. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **durationMinutes** | **int32** | The duration for which the group can be accessed (in minutes). Use 0 to set to indefinite. | 
 **accessLevelRemoteId** | **string** | The remote ID of the access level to grant to this user. If omitted, the default access level remote ID value (empty string) is used. | 
 **addGroupUserRequest** | [**AddGroupUserRequest**](AddGroupUserRequest.md) |  | 

### Return type

[**GroupUser**](GroupUser.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckGroupsRequestable

> BatchRequestabilityCheckResponse CheckGroupsRequestable(ctx).BatchRequestabilityCheckRequest(batchRequestabilityCheckRequest).Execute()

Check requestability of many groups



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
	batchRequestabilityCheckRequest := *openapiclient.NewBatchRequestabilityCheckRequest("UserId_example", []string{"Ids_example"}) // BatchRequestabilityCheckRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CheckGroupsRequestable(context.Background()).BatchRequestabilityCheckRequest(batchRequestabilityCheckRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CheckGroupsRequestable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckGroupsRequestable`: BatchRequestabilityCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CheckGroupsRequestable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckGroupsRequestableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchRequestabilityCheckRequest** | [**BatchRequestabilityCheckRequest**](BatchRequestabilityCheckRequest.md) |  | 

### Return type

[**BatchRequestabilityCheckResponse**](BatchRequestabilityCheckResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> Group CreateGroup(ctx).CreateGroupInfo(createGroupInfo).Execute()





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
	createGroupInfo := *openapiclient.NewCreateGroupInfo("mongo-db-prod", openapiclient.GroupTypeEnum("ACTIVE_DIRECTORY_GROUP"), "f454d283-ca87-4a8a-bdbb-df212eca5353") // CreateGroupInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CreateGroup(context.Background()).CreateGroupInfo(createGroupInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGroupInfo** | [**CreateGroupInfo**](CreateGroupInfo.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, groupId).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.DeleteGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## DeleteGroupCustomAttributeByKey

> DeleteGroupCustomAttributeByKey(ctx, groupId, attributeKey).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67b5" // string | The ID of the group.
	attributeKey := "environment" // string | The custom attribute key to remove.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.DeleteGroupCustomAttributeByKey(context.Background(), groupId, attributeKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroupCustomAttributeByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 
**attributeKey** | **string** | The custom attribute key to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupCustomAttributeByKeyRequest struct via the builder pattern


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


## DeleteGroupUser

> DeleteGroupUser(ctx, groupId, userId).AccessLevelRemoteId(accessLevelRemoteId).DeleteGroupUserRequest(deleteGroupUserRequest).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	userId := "f92aa855-cea9-4814-b9d8-f2a60d3e4a06" // string | The ID of a user to remove from this group.
	accessLevelRemoteId := "30" // string | The remote ID of the access level for which this user has direct access. If omitted, the default access level remote ID value (empty string) is assumed. (optional)
	deleteGroupUserRequest := *openapiclient.NewDeleteGroupUserRequest() // DeleteGroupUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.DeleteGroupUser(context.Background(), groupId, userId).AccessLevelRemoteId(accessLevelRemoteId).DeleteGroupUserRequest(deleteGroupUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroupUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 
**userId** | **string** | The ID of a user to remove from this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessLevelRemoteId** | **string** | The remote ID of the access level for which this user has direct access. If omitted, the default access level remote ID value (empty string) is assumed. | 
 **deleteGroupUserRequest** | [**DeleteGroupUserRequest**](DeleteGroupUserRequest.md) |  | 

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


## GetGroup

> Group GetGroup(ctx, groupId).Execute()

Get group by ID



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
	groupId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The ID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupAccessLevels

> GroupAccessLevelList GetGroupAccessLevels(ctx, groupId).Execute()

Get group access levels



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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group whose access levels to return.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupAccessLevels(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupAccessLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupAccessLevels`: GroupAccessLevelList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupAccessLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group whose access levels to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupAccessLevelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupAccessLevelList**](GroupAccessLevelList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupBreakGlassUsers

> UserIDList GetGroupBreakGlassUsers(ctx, groupId).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupBreakGlassUsers(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupBreakGlassUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupBreakGlassUsers`: UserIDList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupBreakGlassUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupBreakGlassUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserIDList**](UserIDList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupConfigurationTemplate

> GroupConfigurationTemplate GetGroupConfigurationTemplate(ctx, groupId).Execute()

Get the ID of the configuration template linked to a group



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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group whose configuration template link is being read.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupConfigurationTemplate(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupConfigurationTemplate`: GroupConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group whose configuration template link is being read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupConfigurationTemplate**](GroupConfigurationTemplate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupContainingGroup

> GroupContainingGroup GetGroupContainingGroup(ctx, groupId, containingGroupId).Execute()

Get nested group by ID



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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	containingGroupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the containing group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupContainingGroup(context.Background(), groupId, containingGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupContainingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupContainingGroup`: GroupContainingGroup
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupContainingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 
**containingGroupId** | **string** | The ID of the containing group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupContainingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupContainingGroup**](GroupContainingGroup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupContainingGroups

> GroupContainingGroupList GetGroupContainingGroups(ctx, groupId).AccessLevelRemoteId(accessLevelRemoteId).Execute()

Get nested groups



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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	accessLevelRemoteId := "arn:aws:iam::590304332660:role/AdministratorAccess" // string | The access level's remote ID to filter by. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupContainingGroups(context.Background(), groupId).AccessLevelRemoteId(accessLevelRemoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupContainingGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupContainingGroups`: GroupContainingGroupList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupContainingGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupContainingGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessLevelRemoteId** | **string** | The access level&#39;s remote ID to filter by. | 

### Return type

[**GroupContainingGroupList**](GroupContainingGroupList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupCustomAttributes

> CustomAttributesList GetGroupCustomAttributes(ctx, groupId).Execute()





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
	groupId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the group whose custom attributes to return.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupCustomAttributes(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupCustomAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupCustomAttributes`: CustomAttributesList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupCustomAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group whose custom attributes to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupCustomAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomAttributesList**](CustomAttributesList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupGroups

> GroupMemberGroupList GetGroupGroups(ctx, groupId).Execute()

Get member groups



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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupGroups(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupGroups`: GroupMemberGroupList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupMemberGroupList**](GroupMemberGroupList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMessageChannels

> MessageChannelList GetGroupMessageChannels(ctx, groupId).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupMessageChannels(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupMessageChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupMessageChannels`: MessageChannelList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupMessageChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMessageChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageChannelList**](MessageChannelList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupOnCallSchedules

> OnCallScheduleList GetGroupOnCallSchedules(ctx, groupId).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupOnCallSchedules(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupOnCallSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupOnCallSchedules`: OnCallScheduleList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupOnCallSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupOnCallSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OnCallScheduleList**](OnCallScheduleList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupRequestableToUser

> UserRequestabilityCheck GetGroupRequestableToUser(ctx, groupId, userId).AccessLevelRemoteId(accessLevelRemoteId).RequesterId(requesterId).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	userId := "f92aa855-cea9-4814-b9d8-f2a60d3e4a06" // string | The ID of the user the access would be for.
	accessLevelRemoteId := "accessLevelRemoteId_example" // string | Optional access level remote ID to check a single level. Must be an access level of the group. Omit for groups without access levels.  (optional)
	requesterId := "29827fb8-f2dd-4e80-9576-28e31e9934ac" // string | Optional ID of the person filling in the form when different from user_id. Adds the on-behalf check evaluated as that requester, following EnforceRequestOnBehalfOfUser.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupRequestableToUser(context.Background(), groupId, userId).AccessLevelRemoteId(accessLevelRemoteId).RequesterId(requesterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupRequestableToUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupRequestableToUser`: UserRequestabilityCheck
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupRequestableToUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 
**userId** | **string** | The ID of the user the access would be for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequestableToUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessLevelRemoteId** | **string** | Optional access level remote ID to check a single level. Must be an access level of the group. Omit for groups without access levels.  | 
 **requesterId** | **string** | Optional ID of the person filling in the form when different from user_id. Adds the on-behalf check evaluated as that requester, following EnforceRequestOnBehalfOfUser.  | 

### Return type

[**UserRequestabilityCheck**](UserRequestabilityCheck.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupResources

> GroupResourceList GetGroupResources(ctx, groupId).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupResources(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupResources`: GroupResourceList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupResourceList**](GroupResourceList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupReviewerStages

> []ReviewerStage GetGroupReviewerStages(ctx, groupId).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupReviewerStages(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupReviewerStages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupReviewerStages`: []ReviewerStage
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupReviewerStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupReviewerStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ReviewerStage**](ReviewerStage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupReviewers

> []string GetGroupReviewers(ctx, groupId).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupReviewers(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupReviewers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupReviewers`: []string
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupReviewers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupReviewersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupTags

> TagsList GetGroupTags(ctx, groupId).Execute()





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
	groupId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the group whose tags to return.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupTags(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupTags`: TagsList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group whose tags to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagsList**](TagsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupUsers

> GroupUserList GetGroupUsers(ctx, groupId).Cursor(cursor).PageSize(pageSize).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
	pageSize := int32(200) // int32 | Number of results to return per page, up to 1000. When set (or when a cursor is provided), the response contains a single page of results and a `next` cursor; the default page size is 200. When both page_size and cursor are omitted, the entire group membership is returned in one response with no `next` cursor. For large groups, prefer setting page_size and following `next`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupUsers(context.Background(), groupId).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupUsers`: GroupUserList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page, up to 1000. When set (or when a cursor is provided), the response contains a single page of results and a &#x60;next&#x60; cursor; the default page size is 200. When both page_size and cursor are omitted, the entire group membership is returned in one response with no &#x60;next&#x60; cursor. For large groups, prefer setting page_size and following &#x60;next&#x60;. | 

### Return type

[**GroupUserList**](GroupUserList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupVisibility

> VisibilityInfo GetGroupVisibility(ctx, groupId).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupVisibility(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupVisibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupVisibility`: VisibilityInfo
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupVisibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VisibilityInfo**](VisibilityInfo.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupVisibleToUser

> UserVisibilityCheck GetGroupVisibleToUser(ctx, groupId, userId).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	userId := "f92aa855-cea9-4814-b9d8-f2a60d3e4a06" // string | The ID of the user whose visibility is being checked.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupVisibleToUser(context.Background(), groupId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupVisibleToUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupVisibleToUser`: UserVisibilityCheck
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupVisibleToUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 
**userId** | **string** | The ID of the user whose visibility is being checked. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupVisibleToUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserVisibilityCheck**](UserVisibilityCheck.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> PaginatedGroupsList GetGroups(ctx).Cursor(cursor).PageSize(pageSize).GroupTypeFilter(groupTypeFilter).GroupIds(groupIds).GroupName(groupName).GroupNameMatch(groupNameMatch).RemoteId(remoteId).TagIds(tagIds).Requestable(requestable).AppId(appId).RequestableByUserId(requestableByUserId).Execute()

Get groups



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
	groupTypeFilter := openapiclient.GroupTypeEnum("ACTIVE_DIRECTORY_GROUP") // GroupTypeEnum | The group type to filter by. Required when remote_id is provided. (optional)
	groupIds := []string{"1b978423-db0a-4037-a4cf-f79c60cb67b3"} // []string | The group ids to filter by. (optional)
	groupName := "example-name" // string | Filter groups by name. Matches against both `name` and `remote_name`. Matching is case-insensitive. Use `group_name_match` to choose the match type; when omitted, CONTAINS (substring) is used. (optional)
	groupNameMatch := openapiclient.NameMatchType("CONTAINS") // NameMatchType | How to apply `group_name`. Requires `group_name` to be set. Defaults to CONTAINS when `group_name` is provided without this parameter. (optional)
	remoteId := "remoteId_example" // string | Filter groups by their remote id. Returns groups whose remote id matches the provided value exactly. Requires group_type_filter. Remote ids are only unique within a connection, so groups from different connections can share a remote id and more than one group may be returned. Unmanaged groups are not returned. (optional)
	tagIds := []string{"Inner_example"} // []string | The IDs of the tags to filter by. Returns only groups that have any of these tags applied. (optional)
	requestable := true // bool | If true, only return groups that allow access requests. Does not check whether the caller is permitted to request the group. (optional)
	appId := "f454d283-ca87-4a8a-bdbb-df212eca5353" // string | Only return groups that belong to this app. (optional)
	requestableByUserId := "29827fb8-f2dd-4e80-9576-28e31e9934ac" // string | Only return groups this user can request at one or more access levels, using the same evaluation as `GET /groups/{group_id}/requestable/users/{user_id}`. Filtering for another user requires an Opal admin or a custom role with the CHECK_REQUESTABILITY permission. The response is always filtered by the caller's own READ visibility, so a non-admin CHECK_REQUESTABILITY holder sees the intersection of groups they can read and groups the target user can request. Filtering for yourself needs no permission.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroups(context.Background()).Cursor(cursor).PageSize(pageSize).GroupTypeFilter(groupTypeFilter).GroupIds(groupIds).GroupName(groupName).GroupNameMatch(groupNameMatch).RemoteId(remoteId).TagIds(tagIds).Requestable(requestable).AppId(appId).RequestableByUserId(requestableByUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroups`: PaginatedGroupsList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. Default is 200. | 
 **groupTypeFilter** | [**GroupTypeEnum**](GroupTypeEnum.md) | The group type to filter by. Required when remote_id is provided. | 
 **groupIds** | **[]string** | The group ids to filter by. | 
 **groupName** | **string** | Filter groups by name. Matches against both &#x60;name&#x60; and &#x60;remote_name&#x60;. Matching is case-insensitive. Use &#x60;group_name_match&#x60; to choose the match type; when omitted, CONTAINS (substring) is used. | 
 **groupNameMatch** | [**NameMatchType**](NameMatchType.md) | How to apply &#x60;group_name&#x60;. Requires &#x60;group_name&#x60; to be set. Defaults to CONTAINS when &#x60;group_name&#x60; is provided without this parameter. | 
 **remoteId** | **string** | Filter groups by their remote id. Returns groups whose remote id matches the provided value exactly. Requires group_type_filter. Remote ids are only unique within a connection, so groups from different connections can share a remote id and more than one group may be returned. Unmanaged groups are not returned. | 
 **tagIds** | **[]string** | The IDs of the tags to filter by. Returns only groups that have any of these tags applied. | 
 **requestable** | **bool** | If true, only return groups that allow access requests. Does not check whether the caller is permitted to request the group. | 
 **appId** | **string** | Only return groups that belong to this app. | 
 **requestableByUserId** | **string** | Only return groups this user can request at one or more access levels, using the same evaluation as &#x60;GET /groups/{group_id}/requestable/users/{user_id}&#x60;. Filtering for another user requires an Opal admin or a custom role with the CHECK_REQUESTABILITY permission. The response is always filtered by the caller&#39;s own READ visibility, so a non-admin CHECK_REQUESTABILITY holder sees the intersection of groups they can read and groups the target user can request. Filtering for yourself needs no permission.  | 

### Return type

[**PaginatedGroupsList**](PaginatedGroupsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroups

> GroupUserList GetUserGroups(ctx, userId).Cursor(cursor).PageSize(pageSize).Execute()





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
	userId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the user whose groups to return.
	cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
	pageSize := int32(200) // int32 | Number of results to return per page. Default is 200. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetUserGroups(context.Background(), userId).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetUserGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroups`: GroupUserList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetUserGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The ID of the user whose groups to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. Default is 200. | 

### Return type

[**GroupUserList**](GroupUserList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchGroupCustomAttributes

> CustomAttributesList PatchGroupCustomAttributes(ctx, groupId).CustomAttributesList(customAttributesList).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67b5" // string | The ID of the group whose custom attributes to upsert.
	customAttributesList := *openapiclient.NewCustomAttributesList([]openapiclient.CustomAttribute{*openapiclient.NewCustomAttribute("team", openapiclient.CustomAttributeValueTypeEnum("TEXT"), []string{"Values_example"})}) // CustomAttributesList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.PatchGroupCustomAttributes(context.Background(), groupId).CustomAttributesList(customAttributesList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.PatchGroupCustomAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchGroupCustomAttributes`: CustomAttributesList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.PatchGroupCustomAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group whose custom attributes to upsert. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchGroupCustomAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customAttributesList** | [**CustomAttributesList**](CustomAttributesList.md) |  | 

### Return type

[**CustomAttributesList**](CustomAttributesList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveGroupContainingGroup

> RemoveGroupContainingGroup(ctx, groupId, containingGroupId).AccessLevelRemoteId(accessLevelRemoteId).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the member group to remove.
	containingGroupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the containing group.
	accessLevelRemoteId := "arn:aws:iam::590304332660:role/AdministratorAccess" // string | The remote ID of the member group's access level to filter by. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.RemoveGroupContainingGroup(context.Background(), groupId, containingGroupId).AccessLevelRemoteId(accessLevelRemoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.RemoveGroupContainingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the member group to remove. | 
**containingGroupId** | **string** | The ID of the containing group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupContainingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessLevelRemoteId** | **string** | The remote ID of the member group&#39;s access level to filter by. | 

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


## SetGroupBreakGlassUsers

> UserIDList SetGroupBreakGlassUsers(ctx, groupId).UserIDList(userIDList).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	userIDList := *openapiclient.NewUserIDList([]string{"UserIds_example"}) // UserIDList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.SetGroupBreakGlassUsers(context.Background(), groupId).UserIDList(userIDList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.SetGroupBreakGlassUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroupBreakGlassUsers`: UserIDList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.SetGroupBreakGlassUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupBreakGlassUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userIDList** | [**UserIDList**](UserIDList.md) |  | 

### Return type

[**UserIDList**](UserIDList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroupCustomAttributes

> CustomAttributesList SetGroupCustomAttributes(ctx, groupId).CustomAttributesList(customAttributesList).Execute()





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
	groupId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the group whose custom attributes to set.
	customAttributesList := *openapiclient.NewCustomAttributesList([]openapiclient.CustomAttribute{*openapiclient.NewCustomAttribute("team", openapiclient.CustomAttributeValueTypeEnum("TEXT"), []string{"Values_example"})}) // CustomAttributesList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.SetGroupCustomAttributes(context.Background(), groupId).CustomAttributesList(customAttributesList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.SetGroupCustomAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroupCustomAttributes`: CustomAttributesList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.SetGroupCustomAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group whose custom attributes to set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupCustomAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customAttributesList** | [**CustomAttributesList**](CustomAttributesList.md) |  | 

### Return type

[**CustomAttributesList**](CustomAttributesList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroupMessageChannels

> []string SetGroupMessageChannels(ctx, groupId).MessageChannelIDList(messageChannelIDList).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	messageChannelIDList := *openapiclient.NewMessageChannelIDList([]string{"MessageChannelIds_example"}) // MessageChannelIDList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.SetGroupMessageChannels(context.Background(), groupId).MessageChannelIDList(messageChannelIDList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.SetGroupMessageChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroupMessageChannels`: []string
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.SetGroupMessageChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupMessageChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **messageChannelIDList** | [**MessageChannelIDList**](MessageChannelIDList.md) |  | 

### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroupOnCallSchedules

> []string SetGroupOnCallSchedules(ctx, groupId).OnCallScheduleIDList(onCallScheduleIDList).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	onCallScheduleIDList := *openapiclient.NewOnCallScheduleIDList([]string{"OnCallScheduleIds_example"}) // OnCallScheduleIDList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.SetGroupOnCallSchedules(context.Background(), groupId).OnCallScheduleIDList(onCallScheduleIDList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.SetGroupOnCallSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroupOnCallSchedules`: []string
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.SetGroupOnCallSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupOnCallSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onCallScheduleIDList** | [**OnCallScheduleIDList**](OnCallScheduleIDList.md) |  | 

### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroupResources

> SetGroupResources(ctx, groupId).UpdateGroupResourcesInfo(updateGroupResourcesInfo).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	updateGroupResourcesInfo := *openapiclient.NewUpdateGroupResourcesInfo([]openapiclient.ResourceWithAccessLevel{*openapiclient.NewResourceWithAccessLevel("b5a5ca27-0ea3-4d86-9199-2126d57d1fbd")}) // UpdateGroupResourcesInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.SetGroupResources(context.Background(), groupId).UpdateGroupResourcesInfo(updateGroupResourcesInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.SetGroupResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGroupResourcesInfo** | [**UpdateGroupResourcesInfo**](UpdateGroupResourcesInfo.md) |  | 

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


## SetGroupReviewerStages

> []ReviewerStage SetGroupReviewerStages(ctx, groupId).ReviewerStageList(reviewerStageList).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	reviewerStageList := *openapiclient.NewReviewerStageList([]openapiclient.ReviewerStage{*openapiclient.NewReviewerStage(false, "AND", []string{"OwnerIds_example"})}) // ReviewerStageList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.SetGroupReviewerStages(context.Background(), groupId).ReviewerStageList(reviewerStageList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.SetGroupReviewerStages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroupReviewerStages`: []ReviewerStage
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.SetGroupReviewerStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupReviewerStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewerStageList** | [**ReviewerStageList**](ReviewerStageList.md) |  | 

### Return type

[**[]ReviewerStage**](ReviewerStage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroupReviewers

> []string SetGroupReviewers(ctx, groupId).ReviewerIDList(reviewerIDList).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	reviewerIDList := *openapiclient.NewReviewerIDList([]string{"ReviewerIds_example"}) // ReviewerIDList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.SetGroupReviewers(context.Background(), groupId).ReviewerIDList(reviewerIDList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.SetGroupReviewers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroupReviewers`: []string
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.SetGroupReviewers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupReviewersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewerIDList** | [**ReviewerIDList**](ReviewerIDList.md) |  | 

### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroupVisibility

> VisibilityInfo SetGroupVisibility(ctx, groupId).VisibilityInfo(visibilityInfo).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	visibilityInfo := *openapiclient.NewVisibilityInfo(openapiclient.VisibilityTypeEnum("GLOBAL")) // VisibilityInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.SetGroupVisibility(context.Background(), groupId).VisibilityInfo(visibilityInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.SetGroupVisibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroupVisibility`: VisibilityInfo
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.SetGroupVisibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **visibilityInfo** | [**VisibilityInfo**](VisibilityInfo.md) |  | 

### Return type

[**VisibilityInfo**](VisibilityInfo.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupConfigurationTemplate

> GroupConfigurationTemplate UpdateGroupConfigurationTemplate(ctx, groupId).GroupConfigurationTemplate(groupConfigurationTemplate).Execute()

Set the ID of the configuration template linked to a group



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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group whose configuration template link is being set.
	groupConfigurationTemplate := *openapiclient.NewGroupConfigurationTemplate("06851574-e50d-40ca-8c78-f72ae6ab4304") // GroupConfigurationTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UpdateGroupConfigurationTemplate(context.Background(), groupId).GroupConfigurationTemplate(groupConfigurationTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroupConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroupConfigurationTemplate`: GroupConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroupConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group whose configuration template link is being set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupConfigurationTemplate** | [**GroupConfigurationTemplate**](GroupConfigurationTemplate.md) |  | 

### Return type

[**GroupConfigurationTemplate**](GroupConfigurationTemplate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupUser

> GroupUser UpdateGroupUser(ctx, groupId, userId).UpdateGroupUserRequest(updateGroupUserRequest).Execute()





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
	groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
	userId := "f92aa855-cea9-4814-b9d8-f2a60d3e4a06" // string | The ID of the user whose access is being updated.
	updateGroupUserRequest := *openapiclient.NewUpdateGroupUserRequest(int32(120)) // UpdateGroupUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UpdateGroupUser(context.Background(), groupId, userId).UpdateGroupUserRequest(updateGroupUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroupUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroupUser`: GroupUser
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroupUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 
**userId** | **string** | The ID of the user whose access is being updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGroupUserRequest** | [**UpdateGroupUserRequest**](UpdateGroupUserRequest.md) |  | 

### Return type

[**GroupUser**](GroupUser.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroups

> UpdateGroupInfoList UpdateGroups(ctx).UpdateGroupInfoList(updateGroupInfoList).Execute()





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
	updateGroupInfoList := *openapiclient.NewUpdateGroupInfoList([]openapiclient.UpdateGroupInfo{*openapiclient.NewUpdateGroupInfo("f454d283-ca87-4a87-bdbb-df212eca5353")}) // UpdateGroupInfoList | Groups to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UpdateGroups(context.Background()).UpdateGroupInfoList(updateGroupInfoList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroups`: UpdateGroupInfoList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateGroupInfoList** | [**UpdateGroupInfoList**](UpdateGroupInfoList.md) | Groups to be updated | 

### Return type

[**UpdateGroupInfoList**](UpdateGroupInfoList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

