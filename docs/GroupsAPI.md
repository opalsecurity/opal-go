# \GroupsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupContainingGroup**](GroupsAPI.md#AddGroupContainingGroup) | **Post** /groups/{group_id}/containing-groups | 
[**AddGroupResource**](GroupsAPI.md#AddGroupResource) | **Post** /groups/{group_id}/resources/{resource_id} | 
[**AddGroupUser**](GroupsAPI.md#AddGroupUser) | **Post** /groups/{group_id}/users/{user_id} | 
[**CreateGroup**](GroupsAPI.md#CreateGroup) | **Post** /groups | 
[**DeleteGroup**](GroupsAPI.md#DeleteGroup) | **Delete** /groups/{group_id} | 
[**DeleteGroupUser**](GroupsAPI.md#DeleteGroupUser) | **Delete** /groups/{group_id}/users/{user_id} | 
[**GetGroup**](GroupsAPI.md#GetGroup) | **Get** /groups/{group_id} | Get group by ID
[**GetGroupContainingGroup**](GroupsAPI.md#GetGroupContainingGroup) | **Get** /groups/{group_id}/containing-groups/{containing_group_id} | Get nested group by ID
[**GetGroupContainingGroups**](GroupsAPI.md#GetGroupContainingGroups) | **Get** /groups/{group_id}/containing-groups | Get nested groups
[**GetGroupMessageChannels**](GroupsAPI.md#GetGroupMessageChannels) | **Get** /groups/{group_id}/message-channels | 
[**GetGroupOnCallSchedules**](GroupsAPI.md#GetGroupOnCallSchedules) | **Get** /groups/{group_id}/on-call-schedules | 
[**GetGroupResources**](GroupsAPI.md#GetGroupResources) | **Get** /groups/{group_id}/resources | 
[**GetGroupReviewerStages**](GroupsAPI.md#GetGroupReviewerStages) | **Get** /groups/{group_id}/reviewer-stages | 
[**GetGroupReviewers**](GroupsAPI.md#GetGroupReviewers) | **Get** /groups/{group_id}/reviewers | 
[**GetGroupTags**](GroupsAPI.md#GetGroupTags) | **Get** /groups/{group_id}/tags | 
[**GetGroupUsers**](GroupsAPI.md#GetGroupUsers) | **Get** /groups/{group_id}/users | 
[**GetGroupVisibility**](GroupsAPI.md#GetGroupVisibility) | **Get** /groups/{group_id}/visibility | 
[**GetGroups**](GroupsAPI.md#GetGroups) | **Get** /groups | Get groups
[**GetUserGroups**](GroupsAPI.md#GetUserGroups) | **Get** /groups/users/{user_id} | 
[**RemoveGroupContainingGroup**](GroupsAPI.md#RemoveGroupContainingGroup) | **Delete** /groups/{group_id}/containing-groups/{containing_group_id} | 
[**SetGroupMessageChannels**](GroupsAPI.md#SetGroupMessageChannels) | **Put** /groups/{group_id}/message-channels | 
[**SetGroupOnCallSchedules**](GroupsAPI.md#SetGroupOnCallSchedules) | **Put** /groups/{group_id}/on-call-schedules | 
[**SetGroupResources**](GroupsAPI.md#SetGroupResources) | **Put** /groups/{group_id}/resources | 
[**SetGroupReviewerStages**](GroupsAPI.md#SetGroupReviewerStages) | **Put** /groups/{group_id}/reviewer-stages | 
[**SetGroupReviewers**](GroupsAPI.md#SetGroupReviewers) | **Put** /groups/{group_id}/reviewers | 
[**SetGroupVisibility**](GroupsAPI.md#SetGroupVisibility) | **Put** /groups/{group_id}/visibility | 
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


## DeleteGroupUser

> DeleteGroupUser(ctx, groupId, userId).AccessLevelRemoteId(accessLevelRemoteId).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.DeleteGroupUser(context.Background(), groupId, userId).AccessLevelRemoteId(accessLevelRemoteId).Execute()
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

> GroupContainingGroupList GetGroupContainingGroups(ctx, groupId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupContainingGroups(context.Background(), groupId).Execute()
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
	pageSize := int32(200) // int32 | Number of results to return per page. Default is 200. (optional)

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


## GetGroups

> PaginatedGroupsList GetGroups(ctx).Cursor(cursor).PageSize(pageSize).GroupTypeFilter(groupTypeFilter).GroupIds(groupIds).GroupName(groupName).Execute()

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
	groupTypeFilter := openapiclient.GroupTypeEnum("ACTIVE_DIRECTORY_GROUP") // GroupTypeEnum | The group type to filter by. (optional)
	groupIds := []string{"1b978423-db0a-4037-a4cf-f79c60cb67b3"} // []string | The group ids to filter by. (optional)
	groupName := "example-name" // string | Group name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroups(context.Background()).Cursor(cursor).PageSize(pageSize).GroupTypeFilter(groupTypeFilter).GroupIds(groupIds).GroupName(groupName).Execute()
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
 **groupTypeFilter** | [**GroupTypeEnum**](GroupTypeEnum.md) | The group type to filter by. | 
 **groupIds** | **[]string** | The group ids to filter by. | 
 **groupName** | **string** | Group name. | 

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


## RemoveGroupContainingGroup

> RemoveGroupContainingGroup(ctx, groupId, containingGroupId).Execute()





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
	r, err := apiClient.GroupsAPI.RemoveGroupContainingGroup(context.Background(), groupId, containingGroupId).Execute()
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
**groupId** | **string** | The ID of the group. | 
**containingGroupId** | **string** | The ID of the containing group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupContainingGroupRequest struct via the builder pattern


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

