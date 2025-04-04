# Go API client for opal

The Opal API is a RESTful API that allows you to interact with the Opal Security platform programmatically.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Generator version: 7.12.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.opal.dev/](https://www.opal.dev/)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import opal "github.com/opalsecurity/opal-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `opal.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), opal.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `opal.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), opal.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `opal.ContextOperationServerIndices` and `opal.ContextOperationServerVariables` context maps.

```go
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
*AccessRulesAPI* | [**GetAccessRule**](docs/AccessRulesAPI.md#getaccessrule) | **Get** /access-rules/{access_rule_id} | 
*AccessRulesAPI* | [**UpdateAccessRule**](docs/AccessRulesAPI.md#updateaccessrule) | **Put** /access-rules/{access_rule_id} | 
*AppsAPI* | [**GetApp**](docs/AppsAPI.md#getapp) | **Get** /apps/{app_id} | 
*AppsAPI* | [**GetApps**](docs/AppsAPI.md#getapps) | **Get** /apps | 
*AppsAPI* | [**GetSyncErrors**](docs/AppsAPI.md#getsyncerrors) | **Get** /sync_errors | 
*BundlesAPI* | [**AddBundleGroup**](docs/BundlesAPI.md#addbundlegroup) | **Post** /bundles/{bundle_id}/groups | 
*BundlesAPI* | [**AddBundleResource**](docs/BundlesAPI.md#addbundleresource) | **Post** /bundles/{bundle_id}/resources | 
*BundlesAPI* | [**CreateBundle**](docs/BundlesAPI.md#createbundle) | **Post** /bundles | 
*BundlesAPI* | [**DeleteBundle**](docs/BundlesAPI.md#deletebundle) | **Delete** /bundles/{bundle_id} | 
*BundlesAPI* | [**GetBundle**](docs/BundlesAPI.md#getbundle) | **Get** /bundles/{bundle_id} | 
*BundlesAPI* | [**GetBundleGroups**](docs/BundlesAPI.md#getbundlegroups) | **Get** /bundles/{bundle_id}/groups | 
*BundlesAPI* | [**GetBundleResources**](docs/BundlesAPI.md#getbundleresources) | **Get** /bundles/{bundle_id}/resources | 
*BundlesAPI* | [**GetBundleVisibility**](docs/BundlesAPI.md#getbundlevisibility) | **Get** /bundles/{bundle_id}/visibility | 
*BundlesAPI* | [**GetBundles**](docs/BundlesAPI.md#getbundles) | **Get** /bundles | 
*BundlesAPI* | [**RemoveBundleGroup**](docs/BundlesAPI.md#removebundlegroup) | **Delete** /bundles/{bundle_id}/groups/{group_id} | 
*BundlesAPI* | [**RemoveBundleResource**](docs/BundlesAPI.md#removebundleresource) | **Delete** /bundles/{bundle_id}/resources/{resource_id} | 
*BundlesAPI* | [**SetBundleVisibility**](docs/BundlesAPI.md#setbundlevisibility) | **Put** /bundles/{bundle_id}/visibility | 
*BundlesAPI* | [**UpdateBundle**](docs/BundlesAPI.md#updatebundle) | **Put** /bundles/{bundle_id} | 
*ConfigurationTemplatesAPI* | [**CreateConfigurationTemplate**](docs/ConfigurationTemplatesAPI.md#createconfigurationtemplate) | **Post** /configuration-templates | 
*ConfigurationTemplatesAPI* | [**DeleteConfigurationTemplate**](docs/ConfigurationTemplatesAPI.md#deleteconfigurationtemplate) | **Delete** /configuration-templates/{configuration_template_id} | 
*ConfigurationTemplatesAPI* | [**GetConfigurationTemplates**](docs/ConfigurationTemplatesAPI.md#getconfigurationtemplates) | **Get** /configuration-templates | 
*ConfigurationTemplatesAPI* | [**UpdateConfigurationTemplate**](docs/ConfigurationTemplatesAPI.md#updateconfigurationtemplate) | **Put** /configuration-templates | 
*EventsAPI* | [**Events**](docs/EventsAPI.md#events) | **Get** /events | 
*GroupBindingsAPI* | [**CreateGroupBinding**](docs/GroupBindingsAPI.md#creategroupbinding) | **Post** /group-bindings | 
*GroupBindingsAPI* | [**DeleteGroupBinding**](docs/GroupBindingsAPI.md#deletegroupbinding) | **Delete** /group-bindings/{group_binding_id} | 
*GroupBindingsAPI* | [**GetGroupBinding**](docs/GroupBindingsAPI.md#getgroupbinding) | **Get** /group-bindings/{group_binding_id} | 
*GroupBindingsAPI* | [**GetGroupBindings**](docs/GroupBindingsAPI.md#getgroupbindings) | **Get** /group-bindings | 
*GroupBindingsAPI* | [**UpdateGroupBindings**](docs/GroupBindingsAPI.md#updategroupbindings) | **Put** /group-bindings | 
*GroupsAPI* | [**AddGroupContainingGroup**](docs/GroupsAPI.md#addgroupcontaininggroup) | **Post** /groups/{group_id}/containing-groups | 
*GroupsAPI* | [**AddGroupResource**](docs/GroupsAPI.md#addgroupresource) | **Post** /groups/{group_id}/resources/{resource_id} | 
*GroupsAPI* | [**AddGroupUser**](docs/GroupsAPI.md#addgroupuser) | **Post** /groups/{group_id}/users/{user_id} | 
*GroupsAPI* | [**CreateGroup**](docs/GroupsAPI.md#creategroup) | **Post** /groups | 
*GroupsAPI* | [**DeleteGroup**](docs/GroupsAPI.md#deletegroup) | **Delete** /groups/{group_id} | 
*GroupsAPI* | [**DeleteGroupUser**](docs/GroupsAPI.md#deletegroupuser) | **Delete** /groups/{group_id}/users/{user_id} | 
*GroupsAPI* | [**GetGroup**](docs/GroupsAPI.md#getgroup) | **Get** /groups/{group_id} | 
*GroupsAPI* | [**GetGroupContainingGroup**](docs/GroupsAPI.md#getgroupcontaininggroup) | **Get** /groups/{group_id}/containing-groups/{containing_group_id} | 
*GroupsAPI* | [**GetGroupContainingGroups**](docs/GroupsAPI.md#getgroupcontaininggroups) | **Get** /groups/{group_id}/containing-groups | 
*GroupsAPI* | [**GetGroupMessageChannels**](docs/GroupsAPI.md#getgroupmessagechannels) | **Get** /groups/{group_id}/message-channels | 
*GroupsAPI* | [**GetGroupOnCallSchedules**](docs/GroupsAPI.md#getgrouponcallschedules) | **Get** /groups/{group_id}/on-call-schedules | 
*GroupsAPI* | [**GetGroupResources**](docs/GroupsAPI.md#getgroupresources) | **Get** /groups/{group_id}/resources | 
*GroupsAPI* | [**GetGroupReviewerStages**](docs/GroupsAPI.md#getgroupreviewerstages) | **Get** /groups/{group_id}/reviewer-stages | 
*GroupsAPI* | [**GetGroupReviewers**](docs/GroupsAPI.md#getgroupreviewers) | **Get** /groups/{group_id}/reviewers | 
*GroupsAPI* | [**GetGroupTags**](docs/GroupsAPI.md#getgrouptags) | **Get** /groups/{group_id}/tags | 
*GroupsAPI* | [**GetGroupUsers**](docs/GroupsAPI.md#getgroupusers) | **Get** /groups/{group_id}/users | 
*GroupsAPI* | [**GetGroupVisibility**](docs/GroupsAPI.md#getgroupvisibility) | **Get** /groups/{group_id}/visibility | 
*GroupsAPI* | [**GetGroups**](docs/GroupsAPI.md#getgroups) | **Get** /groups | 
*GroupsAPI* | [**RemoveGroupContainingGroup**](docs/GroupsAPI.md#removegroupcontaininggroup) | **Delete** /groups/{group_id}/containing-groups/{containing_group_id} | 
*GroupsAPI* | [**SetGroupMessageChannels**](docs/GroupsAPI.md#setgroupmessagechannels) | **Put** /groups/{group_id}/message-channels | 
*GroupsAPI* | [**SetGroupOnCallSchedules**](docs/GroupsAPI.md#setgrouponcallschedules) | **Put** /groups/{group_id}/on-call-schedules | 
*GroupsAPI* | [**SetGroupResources**](docs/GroupsAPI.md#setgroupresources) | **Put** /groups/{group_id}/resources | 
*GroupsAPI* | [**SetGroupReviewerStages**](docs/GroupsAPI.md#setgroupreviewerstages) | **Put** /groups/{group_id}/reviewer-stages | 
*GroupsAPI* | [**SetGroupReviewers**](docs/GroupsAPI.md#setgroupreviewers) | **Put** /groups/{group_id}/reviewers | 
*GroupsAPI* | [**SetGroupVisibility**](docs/GroupsAPI.md#setgroupvisibility) | **Put** /groups/{group_id}/visibility | 
*GroupsAPI* | [**UpdateGroups**](docs/GroupsAPI.md#updategroups) | **Put** /groups | 
*IdpGroupMappingsAPI* | [**DeleteIdpGroupMappings**](docs/IdpGroupMappingsAPI.md#deleteidpgroupmappings) | **Delete** /idp-group-mappings/{app_resource_id}/{group_id}/ | 
*IdpGroupMappingsAPI* | [**GetIdpGroupMappings**](docs/IdpGroupMappingsAPI.md#getidpgroupmappings) | **Get** /idp-group-mappings/{app_resource_id} | 
*IdpGroupMappingsAPI* | [**UpdateIdpGroupMappings**](docs/IdpGroupMappingsAPI.md#updateidpgroupmappings) | **Put** /idp-group-mappings/{app_resource_id} | 
*MessageChannelsAPI* | [**CreateMessageChannel**](docs/MessageChannelsAPI.md#createmessagechannel) | **Post** /message-channels | 
*MessageChannelsAPI* | [**GetMessageChannel**](docs/MessageChannelsAPI.md#getmessagechannel) | **Get** /message-channels/{message_channel_id} | 
*MessageChannelsAPI* | [**GetMessageChannels**](docs/MessageChannelsAPI.md#getmessagechannels) | **Get** /message-channels | 
*NonHumanIdentitiesAPI* | [**GetNhis**](docs/NonHumanIdentitiesAPI.md#getnhis) | **Get** /non-human-identities | 
*OnCallSchedulesAPI* | [**CreateOnCallSchedule**](docs/OnCallSchedulesAPI.md#createoncallschedule) | **Post** /on-call-schedules | 
*OnCallSchedulesAPI* | [**GetOnCallSchedule**](docs/OnCallSchedulesAPI.md#getoncallschedule) | **Get** /on-call-schedules/{on_call_schedule_id} | 
*OnCallSchedulesAPI* | [**GetOnCallSchedules**](docs/OnCallSchedulesAPI.md#getoncallschedules) | **Get** /on-call-schedules | 
*OwnersAPI* | [**CreateOwner**](docs/OwnersAPI.md#createowner) | **Post** /owners | 
*OwnersAPI* | [**DeleteOwner**](docs/OwnersAPI.md#deleteowner) | **Delete** /owners/{owner_id} | 
*OwnersAPI* | [**GetOwner**](docs/OwnersAPI.md#getowner) | **Get** /owners/{owner_id} | 
*OwnersAPI* | [**GetOwnerFromName**](docs/OwnersAPI.md#getownerfromname) | **Get** /owners/name/{owner_name} | 
*OwnersAPI* | [**GetOwnerUsers**](docs/OwnersAPI.md#getownerusers) | **Get** /owners/{owner_id}/users | 
*OwnersAPI* | [**GetOwners**](docs/OwnersAPI.md#getowners) | **Get** /owners | 
*OwnersAPI* | [**SetOwnerUsers**](docs/OwnersAPI.md#setownerusers) | **Put** /owners/{owner_id}/users | 
*OwnersAPI* | [**UpdateOwners**](docs/OwnersAPI.md#updateowners) | **Put** /owners | 
*RequestsAPI* | [**CreateRequest**](docs/RequestsAPI.md#createrequest) | **Post** /requests | 
*RequestsAPI* | [**GetRequests**](docs/RequestsAPI.md#getrequests) | **Get** /requests | 
*ResourcesAPI* | [**AddResourceNhi**](docs/ResourcesAPI.md#addresourcenhi) | **Post** /resources/{resource_id}/non-human-identities/{non_human_identity_id} | 
*ResourcesAPI* | [**AddResourceUser**](docs/ResourcesAPI.md#addresourceuser) | **Post** /resources/{resource_id}/users/{user_id} | 
*ResourcesAPI* | [**CreateResource**](docs/ResourcesAPI.md#createresource) | **Post** /resources | 
*ResourcesAPI* | [**DeleteResource**](docs/ResourcesAPI.md#deleteresource) | **Delete** /resources/{resource_id} | 
*ResourcesAPI* | [**DeleteResourceNhi**](docs/ResourcesAPI.md#deleteresourcenhi) | **Delete** /resources/{resource_id}/non-human-identities/{non_human_identity_id} | 
*ResourcesAPI* | [**DeleteResourceUser**](docs/ResourcesAPI.md#deleteresourceuser) | **Delete** /resources/{resource_id}/users/{user_id} | 
*ResourcesAPI* | [**GetResource**](docs/ResourcesAPI.md#getresource) | **Get** /resources/{resource_id} | 
*ResourcesAPI* | [**GetResourceMessageChannels**](docs/ResourcesAPI.md#getresourcemessagechannels) | **Get** /resources/{resource_id}/message-channels | 
*ResourcesAPI* | [**GetResourceNhis**](docs/ResourcesAPI.md#getresourcenhis) | **Get** /resources/{resource_id}/non-human-identities | 
*ResourcesAPI* | [**GetResourceReviewerStages**](docs/ResourcesAPI.md#getresourcereviewerstages) | **Get** /resources/{resource_id}/reviewer-stages | 
*ResourcesAPI* | [**GetResourceReviewers**](docs/ResourcesAPI.md#getresourcereviewers) | **Get** /resources/{resource_id}/reviewers | 
*ResourcesAPI* | [**GetResourceTags**](docs/ResourcesAPI.md#getresourcetags) | **Get** /resources/{resource_id}/tags | 
*ResourcesAPI* | [**GetResourceUsers**](docs/ResourcesAPI.md#getresourceusers) | **Get** /resources/{resource_id}/users | 
*ResourcesAPI* | [**GetResourceVisibility**](docs/ResourcesAPI.md#getresourcevisibility) | **Get** /resources/{resource_id}/visibility | 
*ResourcesAPI* | [**GetResources**](docs/ResourcesAPI.md#getresources) | **Get** /resources | 
*ResourcesAPI* | [**ResourceUserAccessStatusRetrieve**](docs/ResourcesAPI.md#resourceuseraccessstatusretrieve) | **Get** /resource-user-access-status/{resource_id}/{user_id} | 
*ResourcesAPI* | [**SetResourceMessageChannels**](docs/ResourcesAPI.md#setresourcemessagechannels) | **Put** /resources/{resource_id}/message-channels | 
*ResourcesAPI* | [**SetResourceReviewerStages**](docs/ResourcesAPI.md#setresourcereviewerstages) | **Put** /resources/{resource_id}/reviewer-stages | 
*ResourcesAPI* | [**SetResourceReviewers**](docs/ResourcesAPI.md#setresourcereviewers) | **Put** /resources/{resource_id}/reviewers | 
*ResourcesAPI* | [**SetResourceVisibility**](docs/ResourcesAPI.md#setresourcevisibility) | **Put** /resources/{resource_id}/visibility | 
*ResourcesAPI* | [**UpdateResourceUser**](docs/ResourcesAPI.md#updateresourceuser) | **Put** /resources/{resource_id}/users/{user_id} | 
*ResourcesAPI* | [**UpdateResources**](docs/ResourcesAPI.md#updateresources) | **Put** /resources | 
*SessionsAPI* | [**Sessions**](docs/SessionsAPI.md#sessions) | **Get** /sessions | 
*TagsAPI* | [**AddGroupTag**](docs/TagsAPI.md#addgrouptag) | **Post** /tags/{tag_id}/groups/{group_id} | 
*TagsAPI* | [**AddResourceTag**](docs/TagsAPI.md#addresourcetag) | **Post** /tags/{tag_id}/resources/{resource_id} | 
*TagsAPI* | [**AddUserTag**](docs/TagsAPI.md#addusertag) | **Post** /tags/{tag_id}/users/{user_id} | 
*TagsAPI* | [**CreateTag**](docs/TagsAPI.md#createtag) | **Post** /tag | 
*TagsAPI* | [**DeleteTagByID**](docs/TagsAPI.md#deletetagbyid) | **Delete** /tag/{tag_id} | 
*TagsAPI* | [**GetTag**](docs/TagsAPI.md#gettag) | **Get** /tag | 
*TagsAPI* | [**GetTagByID**](docs/TagsAPI.md#gettagbyid) | **Get** /tag/{tag_id} | 
*TagsAPI* | [**GetTags**](docs/TagsAPI.md#gettags) | **Get** /tags | 
*TagsAPI* | [**RemoveGroupTag**](docs/TagsAPI.md#removegrouptag) | **Delete** /tags/{tag_id}/groups/{group_id} | 
*TagsAPI* | [**RemoveResourceTag**](docs/TagsAPI.md#removeresourcetag) | **Delete** /tags/{tag_id}/resources/{resource_id} | 
*TagsAPI* | [**RemoveUserTag**](docs/TagsAPI.md#removeusertag) | **Delete** /tags/{tag_id}/users/{user_id} | 
*UarsAPI* | [**CreateUar**](docs/UarsAPI.md#createuar) | **Post** /uar | 
*UarsAPI* | [**GetUARs**](docs/UarsAPI.md#getuars) | **Get** /uars | 
*UarsAPI* | [**GetUar**](docs/UarsAPI.md#getuar) | **Get** /uar/{uar_id} | 
*UsersAPI* | [**GetUserTags**](docs/UsersAPI.md#getusertags) | **Get** /users/{user_id}/tags | 
*UsersAPI* | [**GetUsers**](docs/UsersAPI.md#getusers) | **Get** /users | 
*UsersAPI* | [**User**](docs/UsersAPI.md#user) | **Get** /user | 


## Documentation For Models

 - [Access](docs/Access.md)
 - [AccessList](docs/AccessList.md)
 - [AccessRuleCondition](docs/AccessRuleCondition.md)
 - [AddBundleGroupRequest](docs/AddBundleGroupRequest.md)
 - [AddBundleResourceRequest](docs/AddBundleResourceRequest.md)
 - [AddGroupResourceRequest](docs/AddGroupResourceRequest.md)
 - [AddGroupUserRequest](docs/AddGroupUserRequest.md)
 - [AddResourceNhiRequest](docs/AddResourceNhiRequest.md)
 - [AddResourceUserRequest](docs/AddResourceUserRequest.md)
 - [App](docs/App.md)
 - [AppTypeEnum](docs/AppTypeEnum.md)
 - [AppValidation](docs/AppValidation.md)
 - [AppValidationSeverityEnum](docs/AppValidationSeverityEnum.md)
 - [AppValidationStatusEnum](docs/AppValidationStatusEnum.md)
 - [AppsList](docs/AppsList.md)
 - [AwsPermissionSetMetadata](docs/AwsPermissionSetMetadata.md)
 - [AwsPermissionSetMetadataAwsPermissionSet](docs/AwsPermissionSetMetadataAwsPermissionSet.md)
 - [Bundle](docs/Bundle.md)
 - [BundleGroup](docs/BundleGroup.md)
 - [BundleResource](docs/BundleResource.md)
 - [Condition](docs/Condition.md)
 - [ConfigurationTemplate](docs/ConfigurationTemplate.md)
 - [CreateBundleInfo](docs/CreateBundleInfo.md)
 - [CreateConfigurationTemplateInfo](docs/CreateConfigurationTemplateInfo.md)
 - [CreateGroupBindingInfo](docs/CreateGroupBindingInfo.md)
 - [CreateGroupBindingInfoGroupsInner](docs/CreateGroupBindingInfoGroupsInner.md)
 - [CreateGroupInfo](docs/CreateGroupInfo.md)
 - [CreateMessageChannelInfo](docs/CreateMessageChannelInfo.md)
 - [CreateOnCallScheduleInfo](docs/CreateOnCallScheduleInfo.md)
 - [CreateOwnerInfo](docs/CreateOwnerInfo.md)
 - [CreateRequest200Response](docs/CreateRequest200Response.md)
 - [CreateRequestConfigurationInfoList](docs/CreateRequestConfigurationInfoList.md)
 - [CreateRequestInfo](docs/CreateRequestInfo.md)
 - [CreateRequestInfoCustomMetadataInner](docs/CreateRequestInfoCustomMetadataInner.md)
 - [CreateRequestInfoGroupsInner](docs/CreateRequestInfoGroupsInner.md)
 - [CreateRequestInfoResourcesInner](docs/CreateRequestInfoResourcesInner.md)
 - [CreateRequestInfoSupportTicket](docs/CreateRequestInfoSupportTicket.md)
 - [CreateResourceInfo](docs/CreateResourceInfo.md)
 - [CreateTagInfo](docs/CreateTagInfo.md)
 - [CreateUARInfo](docs/CreateUARInfo.md)
 - [EntityTypeEnum](docs/EntityTypeEnum.md)
 - [Event](docs/Event.md)
 - [Group](docs/Group.md)
 - [GroupAccessLevel](docs/GroupAccessLevel.md)
 - [GroupBinding](docs/GroupBinding.md)
 - [GroupBindingGroup](docs/GroupBindingGroup.md)
 - [GroupContainingGroup](docs/GroupContainingGroup.md)
 - [GroupContainingGroupList](docs/GroupContainingGroupList.md)
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
 - [GroupWithAccessLevel](docs/GroupWithAccessLevel.md)
 - [IdpGroupMapping](docs/IdpGroupMapping.md)
 - [IdpGroupMappingList](docs/IdpGroupMappingList.md)
 - [MessageChannel](docs/MessageChannel.md)
 - [MessageChannelIDList](docs/MessageChannelIDList.md)
 - [MessageChannelList](docs/MessageChannelList.md)
 - [MessageChannelProviderEnum](docs/MessageChannelProviderEnum.md)
 - [OnCallSchedule](docs/OnCallSchedule.md)
 - [OnCallScheduleIDList](docs/OnCallScheduleIDList.md)
 - [OnCallScheduleList](docs/OnCallScheduleList.md)
 - [OnCallScheduleProviderEnum](docs/OnCallScheduleProviderEnum.md)
 - [Owner](docs/Owner.md)
 - [PaginatedBundleGroupList](docs/PaginatedBundleGroupList.md)
 - [PaginatedBundleList](docs/PaginatedBundleList.md)
 - [PaginatedBundleResourceList](docs/PaginatedBundleResourceList.md)
 - [PaginatedConfigurationTemplateList](docs/PaginatedConfigurationTemplateList.md)
 - [PaginatedEventList](docs/PaginatedEventList.md)
 - [PaginatedGroupBindingsList](docs/PaginatedGroupBindingsList.md)
 - [PaginatedGroupsList](docs/PaginatedGroupsList.md)
 - [PaginatedOwnersList](docs/PaginatedOwnersList.md)
 - [PaginatedResourcesList](docs/PaginatedResourcesList.md)
 - [PaginatedTagsList](docs/PaginatedTagsList.md)
 - [PaginatedUARsList](docs/PaginatedUARsList.md)
 - [PaginatedUsersList](docs/PaginatedUsersList.md)
 - [PropagationStatus](docs/PropagationStatus.md)
 - [PropagationStatusEnum](docs/PropagationStatusEnum.md)
 - [Request](docs/Request.md)
 - [RequestConfiguration](docs/RequestConfiguration.md)
 - [RequestCustomFieldResponse](docs/RequestCustomFieldResponse.md)
 - [RequestCustomFieldResponseFieldValue](docs/RequestCustomFieldResponseFieldValue.md)
 - [RequestList](docs/RequestList.md)
 - [RequestStatusEnum](docs/RequestStatusEnum.md)
 - [RequestTemplateCustomFieldTypeEnum](docs/RequestTemplateCustomFieldTypeEnum.md)
 - [RequestedItem](docs/RequestedItem.md)
 - [Resource](docs/Resource.md)
 - [ResourceAccessLevel](docs/ResourceAccessLevel.md)
 - [ResourceAccessUser](docs/ResourceAccessUser.md)
 - [ResourceAccessUserList](docs/ResourceAccessUserList.md)
 - [ResourceNHI](docs/ResourceNHI.md)
 - [ResourceRemoteInfo](docs/ResourceRemoteInfo.md)
 - [ResourceRemoteInfoAwsAccount](docs/ResourceRemoteInfoAwsAccount.md)
 - [ResourceRemoteInfoAwsEc2Instance](docs/ResourceRemoteInfoAwsEc2Instance.md)
 - [ResourceRemoteInfoAwsEksCluster](docs/ResourceRemoteInfoAwsEksCluster.md)
 - [ResourceRemoteInfoAwsIamRole](docs/ResourceRemoteInfoAwsIamRole.md)
 - [ResourceRemoteInfoAwsPermissionSet](docs/ResourceRemoteInfoAwsPermissionSet.md)
 - [ResourceRemoteInfoAwsRdsInstance](docs/ResourceRemoteInfoAwsRdsInstance.md)
 - [ResourceRemoteInfoGcpBigQueryDataset](docs/ResourceRemoteInfoGcpBigQueryDataset.md)
 - [ResourceRemoteInfoGcpBigQueryTable](docs/ResourceRemoteInfoGcpBigQueryTable.md)
 - [ResourceRemoteInfoGcpBucket](docs/ResourceRemoteInfoGcpBucket.md)
 - [ResourceRemoteInfoGcpComputeInstance](docs/ResourceRemoteInfoGcpComputeInstance.md)
 - [ResourceRemoteInfoGcpFolder](docs/ResourceRemoteInfoGcpFolder.md)
 - [ResourceRemoteInfoGcpGkeCluster](docs/ResourceRemoteInfoGcpGkeCluster.md)
 - [ResourceRemoteInfoGcpOrganization](docs/ResourceRemoteInfoGcpOrganization.md)
 - [ResourceRemoteInfoGcpProject](docs/ResourceRemoteInfoGcpProject.md)
 - [ResourceRemoteInfoGcpServiceAccount](docs/ResourceRemoteInfoGcpServiceAccount.md)
 - [ResourceRemoteInfoGcpSqlInstance](docs/ResourceRemoteInfoGcpSqlInstance.md)
 - [ResourceRemoteInfoGithubRepo](docs/ResourceRemoteInfoGithubRepo.md)
 - [ResourceRemoteInfoGitlabProject](docs/ResourceRemoteInfoGitlabProject.md)
 - [ResourceRemoteInfoOktaApp](docs/ResourceRemoteInfoOktaApp.md)
 - [ResourceRemoteInfoOktaCustomRole](docs/ResourceRemoteInfoOktaCustomRole.md)
 - [ResourceRemoteInfoOktaStandardRole](docs/ResourceRemoteInfoOktaStandardRole.md)
 - [ResourceRemoteInfoPagerdutyRole](docs/ResourceRemoteInfoPagerdutyRole.md)
 - [ResourceRemoteInfoSalesforcePermissionSet](docs/ResourceRemoteInfoSalesforcePermissionSet.md)
 - [ResourceRemoteInfoSalesforceProfile](docs/ResourceRemoteInfoSalesforceProfile.md)
 - [ResourceRemoteInfoSalesforceRole](docs/ResourceRemoteInfoSalesforceRole.md)
 - [ResourceRemoteInfoTeleportRole](docs/ResourceRemoteInfoTeleportRole.md)
 - [ResourceTypeEnum](docs/ResourceTypeEnum.md)
 - [ResourceUser](docs/ResourceUser.md)
 - [ResourceUserAccessStatus](docs/ResourceUserAccessStatus.md)
 - [ResourceUserAccessStatusEnum](docs/ResourceUserAccessStatusEnum.md)
 - [ResourceWithAccessLevel](docs/ResourceWithAccessLevel.md)
 - [ReviewerIDList](docs/ReviewerIDList.md)
 - [ReviewerStage](docs/ReviewerStage.md)
 - [ReviewerStageList](docs/ReviewerStageList.md)
 - [RiskSensitivityEnum](docs/RiskSensitivityEnum.md)
 - [RuleClauses](docs/RuleClauses.md)
 - [RuleConjunction](docs/RuleConjunction.md)
 - [RuleDisjunction](docs/RuleDisjunction.md)
 - [Session](docs/Session.md)
 - [SessionsList](docs/SessionsList.md)
 - [SubEvent](docs/SubEvent.md)
 - [SyncError](docs/SyncError.md)
 - [SyncErrorList](docs/SyncErrorList.md)
 - [Tag](docs/Tag.md)
 - [TagFilter](docs/TagFilter.md)
 - [TagSelector](docs/TagSelector.md)
 - [TagsList](docs/TagsList.md)
 - [TicketPropagationConfiguration](docs/TicketPropagationConfiguration.md)
 - [TicketingProviderEnum](docs/TicketingProviderEnum.md)
 - [UAR](docs/UAR.md)
 - [UARReviewerAssignmentPolicyEnum](docs/UARReviewerAssignmentPolicyEnum.md)
 - [UARScope](docs/UARScope.md)
 - [UpdateConfigurationTemplateInfo](docs/UpdateConfigurationTemplateInfo.md)
 - [UpdateGroupBindingInfo](docs/UpdateGroupBindingInfo.md)
 - [UpdateGroupBindingInfoList](docs/UpdateGroupBindingInfoList.md)
 - [UpdateGroupInfo](docs/UpdateGroupInfo.md)
 - [UpdateGroupInfoList](docs/UpdateGroupInfoList.md)
 - [UpdateGroupResourcesInfo](docs/UpdateGroupResourcesInfo.md)
 - [UpdateIdpGroupMappingsRequest](docs/UpdateIdpGroupMappingsRequest.md)
 - [UpdateIdpGroupMappingsRequestMappingsInner](docs/UpdateIdpGroupMappingsRequestMappingsInner.md)
 - [UpdateOwnerInfo](docs/UpdateOwnerInfo.md)
 - [UpdateOwnerInfoList](docs/UpdateOwnerInfoList.md)
 - [UpdateResourceInfo](docs/UpdateResourceInfo.md)
 - [UpdateResourceInfoList](docs/UpdateResourceInfoList.md)
 - [UpdateResourceUserRequest](docs/UpdateResourceUserRequest.md)
 - [User](docs/User.md)
 - [UserHrIdpStatusEnum](docs/UserHrIdpStatusEnum.md)
 - [UserIDList](docs/UserIDList.md)
 - [UserList](docs/UserList.md)
 - [VisibilityInfo](docs/VisibilityInfo.md)
 - [VisibilityTypeEnum](docs/VisibilityTypeEnum.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), opal.ContextAccessToken, "BEARER_TOKEN_STRING")
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

