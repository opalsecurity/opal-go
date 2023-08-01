# Go API client for opal

Your Home For Developer Resources.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.opal.dev/](https://www.opal.dev/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import opal "github.com/opalsecurity/opal-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), opal.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), opal.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), opal.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), opal.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.opal.dev/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AppsApi* | [**GetApp**](docs/AppsApi.md#getapp) | **Get** /apps/{app_id} | 
*AppsApi* | [**GetApps**](docs/AppsApi.md#getapps) | **Get** /apps | 
*ConfigurationTemplatesApi* | [**CreateConfigurationTemplate**](docs/ConfigurationTemplatesApi.md#createconfigurationtemplate) | **Post** /configuration-templates | 
*ConfigurationTemplatesApi* | [**GetConfigurationTemplates**](docs/ConfigurationTemplatesApi.md#getconfigurationtemplates) | **Get** /configuration-templates | 
*DefaultApi* | [**UpdateConfigurationTemplate**](docs/DefaultApi.md#updateconfigurationtemplate) | **Put** /configuration-templates | 
*EventsApi* | [**Events**](docs/EventsApi.md#events) | **Get** /events | 
*GroupsApi* | [**AddGroupResource**](docs/GroupsApi.md#addgroupresource) | **Post** /groups/{group_id}/resources/{resource_id} | 
*GroupsApi* | [**AddGroupUser**](docs/GroupsApi.md#addgroupuser) | **Post** /groups/{group_id}/users/{user_id} | 
*GroupsApi* | [**CreateGroup**](docs/GroupsApi.md#creategroup) | **Post** /groups | 
*GroupsApi* | [**DeleteGroup**](docs/GroupsApi.md#deletegroup) | **Delete** /groups/{group_id} | 
*GroupsApi* | [**DeleteGroupUser**](docs/GroupsApi.md#deletegroupuser) | **Delete** /groups/{group_id}/users/{user_id} | 
*GroupsApi* | [**GetGroupMessageChannels**](docs/GroupsApi.md#getgroupmessagechannels) | **Get** /groups/{group_id}/message-channels | 
*GroupsApi* | [**GetGroupOnCallSchedules**](docs/GroupsApi.md#getgrouponcallschedules) | **Get** /groups/{group_id}/on-call-schedules | 
*GroupsApi* | [**GetGroupResources**](docs/GroupsApi.md#getgroupresources) | **Get** /groups/{group_id}/resources | 
*GroupsApi* | [**GetGroupReviewerStages**](docs/GroupsApi.md#getgroupreviewerstages) | **Get** /groups/{group_id}/reviewer-stages | 
*GroupsApi* | [**GetGroupReviewers**](docs/GroupsApi.md#getgroupreviewers) | **Get** /groups/{group_id}/reviewers | 
*GroupsApi* | [**GetGroupTags**](docs/GroupsApi.md#getgrouptags) | **Get** /groups/{group_id}/tags | 
*GroupsApi* | [**GetGroupUsers**](docs/GroupsApi.md#getgroupusers) | **Get** /groups/{group_id}/users | 
*GroupsApi* | [**GetGroupVisibility**](docs/GroupsApi.md#getgroupvisibility) | **Get** /groups/{group_id}/visibility | 
*GroupsApi* | [**GetGroups**](docs/GroupsApi.md#getgroups) | **Get** /groups | 
*GroupsApi* | [**SetGroupMessageChannels**](docs/GroupsApi.md#setgroupmessagechannels) | **Put** /groups/{group_id}/message-channels | 
*GroupsApi* | [**SetGroupOnCallSchedules**](docs/GroupsApi.md#setgrouponcallschedules) | **Put** /groups/{group_id}/on-call-schedules | 
*GroupsApi* | [**SetGroupResources**](docs/GroupsApi.md#setgroupresources) | **Put** /groups/{group_id}/resources | 
*GroupsApi* | [**SetGroupReviewerStages**](docs/GroupsApi.md#setgroupreviewerstages) | **Put** /groups/{group_id}/reviewer-stages | 
*GroupsApi* | [**SetGroupReviewers**](docs/GroupsApi.md#setgroupreviewers) | **Put** /groups/{group_id}/reviewers | 
*GroupsApi* | [**SetGroupVisibility**](docs/GroupsApi.md#setgroupvisibility) | **Put** /groups/{group_id}/visibility | 
*GroupsApi* | [**UpdateGroups**](docs/GroupsApi.md#updategroups) | **Put** /groups | 
*MessageChannelsApi* | [**CreateMessageChannel**](docs/MessageChannelsApi.md#createmessagechannel) | **Post** /message-channels | 
*MessageChannelsApi* | [**GetMessageChannel**](docs/MessageChannelsApi.md#getmessagechannel) | **Get** /message-channels/{message_channel_id} | 
*MessageChannelsApi* | [**GetMessageChannels**](docs/MessageChannelsApi.md#getmessagechannels) | **Get** /message-channels | 
*OnCallSchedulesApi* | [**CreateOnCallSchedule**](docs/OnCallSchedulesApi.md#createoncallschedule) | **Post** /on-call-schedules | 
*OnCallSchedulesApi* | [**GetOnCallSchedule**](docs/OnCallSchedulesApi.md#getoncallschedule) | **Get** /on-call-schedules/{on_call_schedule_id} | 
*OnCallSchedulesApi* | [**GetOnCallSchedules**](docs/OnCallSchedulesApi.md#getoncallschedules) | **Get** /on-call-schedules | 
*OwnersApi* | [**CreateOwner**](docs/OwnersApi.md#createowner) | **Post** /owners | 
*OwnersApi* | [**DeleteOwner**](docs/OwnersApi.md#deleteowner) | **Delete** /owners/{owner_id} | 
*OwnersApi* | [**GetOwner**](docs/OwnersApi.md#getowner) | **Get** /owners/{owner_id} | 
*OwnersApi* | [**GetOwnerFromName**](docs/OwnersApi.md#getownerfromname) | **Get** /owners/name/{owner_name} | 
*OwnersApi* | [**GetOwnerUsers**](docs/OwnersApi.md#getownerusers) | **Get** /owners/{owner_id}/users | 
*OwnersApi* | [**GetOwners**](docs/OwnersApi.md#getowners) | **Get** /owners | 
*OwnersApi* | [**SetOwnerUsers**](docs/OwnersApi.md#setownerusers) | **Put** /owners/{owner_id}/users | 
*OwnersApi* | [**UpdateOwners**](docs/OwnersApi.md#updateowners) | **Put** /owners | 
*RequestsApi* | [**GetRequests**](docs/RequestsApi.md#getrequests) | **Get** /requests | 
*ResourcesApi* | [**AddResourceUser**](docs/ResourcesApi.md#addresourceuser) | **Post** /resources/{resource_id}/users/{user_id} | 
*ResourcesApi* | [**CreateResource**](docs/ResourcesApi.md#createresource) | **Post** /resources | 
*ResourcesApi* | [**DeleteResource**](docs/ResourcesApi.md#deleteresource) | **Delete** /resources/{resource_id} | 
*ResourcesApi* | [**DeleteResourceUser**](docs/ResourcesApi.md#deleteresourceuser) | **Delete** /resources/{resource_id}/users/{user_id} | 
*ResourcesApi* | [**GetResource**](docs/ResourcesApi.md#getresource) | **Get** /resources/{resource_id} | 
*ResourcesApi* | [**GetResourceMessageChannels**](docs/ResourcesApi.md#getresourcemessagechannels) | **Get** /resources/{resource_id}/message-channels | 
*ResourcesApi* | [**GetResourceReviewerStages**](docs/ResourcesApi.md#getresourcereviewerstages) | **Get** /resources/{resource_id}/reviewer-stages | 
*ResourcesApi* | [**GetResourceReviewers**](docs/ResourcesApi.md#getresourcereviewers) | **Get** /resources/{resource_id}/reviewers | 
*ResourcesApi* | [**GetResourceTags**](docs/ResourcesApi.md#getresourcetags) | **Get** /resources/{resource_id}/tags | 
*ResourcesApi* | [**GetResourceUsers**](docs/ResourcesApi.md#getresourceusers) | **Get** /resources/{resource_id}/users | 
*ResourcesApi* | [**GetResourceVisibility**](docs/ResourcesApi.md#getresourcevisibility) | **Get** /resources/{resource_id}/visibility | 
*ResourcesApi* | [**GetResources**](docs/ResourcesApi.md#getresources) | **Get** /resources | 
*ResourcesApi* | [**ResourceUserAccessStatusRetrieve**](docs/ResourcesApi.md#resourceuseraccessstatusretrieve) | **Get** /resource-user-access-status/{resource_id}/{user_id} | 
*ResourcesApi* | [**SetResourceMessageChannels**](docs/ResourcesApi.md#setresourcemessagechannels) | **Put** /resources/{resource_id}/message-channels | 
*ResourcesApi* | [**SetResourceReviewerStages**](docs/ResourcesApi.md#setresourcereviewerstages) | **Put** /resources/{resource_id}/reviewer-stages | 
*ResourcesApi* | [**SetResourceReviewers**](docs/ResourcesApi.md#setresourcereviewers) | **Put** /resources/{resource_id}/reviewers | 
*ResourcesApi* | [**SetResourceVisibility**](docs/ResourcesApi.md#setresourcevisibility) | **Put** /resources/{resource_id}/visibility | 
*ResourcesApi* | [**UpdateResources**](docs/ResourcesApi.md#updateresources) | **Put** /resources | 
*SessionsApi* | [**Sessions**](docs/SessionsApi.md#sessions) | **Get** /sessions | 
*TagsApi* | [**AddGroupTag**](docs/TagsApi.md#addgrouptag) | **Post** /tags/{tag_id}/groups/{group_id} | 
*TagsApi* | [**AddResourceTag**](docs/TagsApi.md#addresourcetag) | **Post** /tags/{tag_id}/resources/{resource_id} | 
*TagsApi* | [**AddUserTag**](docs/TagsApi.md#addusertag) | **Post** /tags/{tag_id}/users/{user_id} | 
*TagsApi* | [**CreateTag**](docs/TagsApi.md#createtag) | **Post** /tag | 
*TagsApi* | [**GetTag**](docs/TagsApi.md#gettag) | **Get** /tag | 
*TagsApi* | [**GetTags**](docs/TagsApi.md#gettags) | **Get** /tags | 
*TagsApi* | [**RemoveGroupTag**](docs/TagsApi.md#removegrouptag) | **Delete** /tags/{tag_id}/groups/{group_id} | 
*TagsApi* | [**RemoveResourceTag**](docs/TagsApi.md#removeresourcetag) | **Delete** /tags/{tag_id}/resources/{resource_id} | 
*TagsApi* | [**RemoveUserTag**](docs/TagsApi.md#removeusertag) | **Delete** /tags/{tag_id}/users/{user_id} | 
*UarsApi* | [**CreateUar**](docs/UarsApi.md#createuar) | **Post** /uar | 
*UarsApi* | [**GetUARs**](docs/UarsApi.md#getuars) | **Get** /uars | 
*UarsApi* | [**GetUar**](docs/UarsApi.md#getuar) | **Get** /uar/{uar_id} | 
*UsersApi* | [**GetUserTags**](docs/UsersApi.md#getusertags) | **Get** /users/{user_id}/tags | 
*UsersApi* | [**GetUsers**](docs/UsersApi.md#getusers) | **Get** /users | 
*UsersApi* | [**User**](docs/UsersApi.md#user) | **Get** /user | 


## Documentation For Models

 - [App](docs/App.md)
 - [AppTypeEnum](docs/AppTypeEnum.md)
 - [AppsList](docs/AppsList.md)
 - [AwsPermissionSetMetadata](docs/AwsPermissionSetMetadata.md)
 - [AwsPermissionSetMetadataAwsPermissionSet](docs/AwsPermissionSetMetadataAwsPermissionSet.md)
 - [ConfigurationTemplate](docs/ConfigurationTemplate.md)
 - [CreateConfigurationTemplateInfo](docs/CreateConfigurationTemplateInfo.md)
 - [CreateGroupInfo](docs/CreateGroupInfo.md)
 - [CreateMessageChannelInfo](docs/CreateMessageChannelInfo.md)
 - [CreateOnCallScheduleInfo](docs/CreateOnCallScheduleInfo.md)
 - [CreateOwnerInfo](docs/CreateOwnerInfo.md)
 - [CreateResourceInfo](docs/CreateResourceInfo.md)
 - [CreateUARInfo](docs/CreateUARInfo.md)
 - [EntityTypeEnum](docs/EntityTypeEnum.md)
 - [Event](docs/Event.md)
 - [Group](docs/Group.md)
 - [GroupAccessLevel](docs/GroupAccessLevel.md)
 - [GroupRemoteInfo](docs/GroupRemoteInfo.md)
 - [GroupRemoteInfoActiveDirectoryGroup](docs/GroupRemoteInfoActiveDirectoryGroup.md)
 - [GroupRemoteInfoAzureAdMicrosoft365Group](docs/GroupRemoteInfoAzureAdMicrosoft365Group.md)
 - [GroupRemoteInfoAzureAdSecurityGroup](docs/GroupRemoteInfoAzureAdSecurityGroup.md)
 - [GroupRemoteInfoDuoGroup](docs/GroupRemoteInfoDuoGroup.md)
 - [GroupRemoteInfoGithubTeam](docs/GroupRemoteInfoGithubTeam.md)
 - [GroupRemoteInfoGitlabGroup](docs/GroupRemoteInfoGitlabGroup.md)
 - [GroupRemoteInfoGoogleGroup](docs/GroupRemoteInfoGoogleGroup.md)
 - [GroupRemoteInfoLdapGroup](docs/GroupRemoteInfoLdapGroup.md)
 - [GroupRemoteInfoOktaGroup](docs/GroupRemoteInfoOktaGroup.md)
 - [GroupResource](docs/GroupResource.md)
 - [GroupResourceList](docs/GroupResourceList.md)
 - [GroupTypeEnum](docs/GroupTypeEnum.md)
 - [GroupUser](docs/GroupUser.md)
 - [GroupUserList](docs/GroupUserList.md)
 - [MessageChannel](docs/MessageChannel.md)
 - [MessageChannelIDList](docs/MessageChannelIDList.md)
 - [MessageChannelList](docs/MessageChannelList.md)
 - [MessageChannelProviderEnum](docs/MessageChannelProviderEnum.md)
 - [OnCallSchedule](docs/OnCallSchedule.md)
 - [OnCallScheduleIDList](docs/OnCallScheduleIDList.md)
 - [OnCallScheduleList](docs/OnCallScheduleList.md)
 - [OnCallScheduleProviderEnum](docs/OnCallScheduleProviderEnum.md)
 - [Owner](docs/Owner.md)
 - [PaginatedConfigurationTemplateList](docs/PaginatedConfigurationTemplateList.md)
 - [PaginatedEventList](docs/PaginatedEventList.md)
 - [PaginatedGroupsList](docs/PaginatedGroupsList.md)
 - [PaginatedOwnersList](docs/PaginatedOwnersList.md)
 - [PaginatedResourcesList](docs/PaginatedResourcesList.md)
 - [PaginatedTagsList](docs/PaginatedTagsList.md)
 - [PaginatedUARsList](docs/PaginatedUARsList.md)
 - [PaginatedUsersList](docs/PaginatedUsersList.md)
 - [Request](docs/Request.md)
 - [RequestList](docs/RequestList.md)
 - [RequestStatusEnum](docs/RequestStatusEnum.md)
 - [Resource](docs/Resource.md)
 - [ResourceAccessLevel](docs/ResourceAccessLevel.md)
 - [ResourceAccessUser](docs/ResourceAccessUser.md)
 - [ResourceAccessUserList](docs/ResourceAccessUserList.md)
 - [ResourceRemoteInfo](docs/ResourceRemoteInfo.md)
 - [ResourceRemoteInfoAwsAccount](docs/ResourceRemoteInfoAwsAccount.md)
 - [ResourceRemoteInfoAwsEc2Instance](docs/ResourceRemoteInfoAwsEc2Instance.md)
 - [ResourceRemoteInfoAwsEksCluster](docs/ResourceRemoteInfoAwsEksCluster.md)
 - [ResourceRemoteInfoAwsIamRole](docs/ResourceRemoteInfoAwsIamRole.md)
 - [ResourceRemoteInfoAwsPermissionSet](docs/ResourceRemoteInfoAwsPermissionSet.md)
 - [ResourceRemoteInfoAwsRdsInstance](docs/ResourceRemoteInfoAwsRdsInstance.md)
 - [ResourceRemoteInfoGithubRepo](docs/ResourceRemoteInfoGithubRepo.md)
 - [ResourceRemoteInfoGitlabProject](docs/ResourceRemoteInfoGitlabProject.md)
 - [ResourceRemoteInfoOktaApp](docs/ResourceRemoteInfoOktaApp.md)
 - [ResourceRemoteInfoOktaCustomRole](docs/ResourceRemoteInfoOktaCustomRole.md)
 - [ResourceRemoteInfoOktaStandardRole](docs/ResourceRemoteInfoOktaStandardRole.md)
 - [ResourceRemoteInfoTeleportRole](docs/ResourceRemoteInfoTeleportRole.md)
 - [ResourceTypeEnum](docs/ResourceTypeEnum.md)
 - [ResourceUser](docs/ResourceUser.md)
 - [ResourceUserAccessStatus](docs/ResourceUserAccessStatus.md)
 - [ResourceUserAccessStatusEnum](docs/ResourceUserAccessStatusEnum.md)
 - [ResourceWithAccessLevel](docs/ResourceWithAccessLevel.md)
 - [ReviewerIDList](docs/ReviewerIDList.md)
 - [ReviewerStage](docs/ReviewerStage.md)
 - [ReviewerStageList](docs/ReviewerStageList.md)
 - [Session](docs/Session.md)
 - [SessionsList](docs/SessionsList.md)
 - [SubEvent](docs/SubEvent.md)
 - [Tag](docs/Tag.md)
 - [TagFilter](docs/TagFilter.md)
 - [TagsList](docs/TagsList.md)
 - [UAR](docs/UAR.md)
 - [UARReviewerAssignmentPolicyEnum](docs/UARReviewerAssignmentPolicyEnum.md)
 - [UARScope](docs/UARScope.md)
 - [UpdateConfigurationTemplateInfo](docs/UpdateConfigurationTemplateInfo.md)
 - [UpdateGroupInfo](docs/UpdateGroupInfo.md)
 - [UpdateGroupInfoList](docs/UpdateGroupInfoList.md)
 - [UpdateGroupResourcesInfo](docs/UpdateGroupResourcesInfo.md)
 - [UpdateOwnerInfo](docs/UpdateOwnerInfo.md)
 - [UpdateOwnerInfoList](docs/UpdateOwnerInfoList.md)
 - [UpdateResourceInfo](docs/UpdateResourceInfo.md)
 - [UpdateResourceInfoList](docs/UpdateResourceInfoList.md)
 - [User](docs/User.md)
 - [UserIDList](docs/UserIDList.md)
 - [UserList](docs/UserList.md)
 - [VisibilityInfo](docs/VisibilityInfo.md)
 - [VisibilityTypeEnum](docs/VisibilityTypeEnum.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

hello@opal.dev

