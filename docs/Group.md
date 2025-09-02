# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The ID of the group. | 
**AppId** | Pointer to **string** | The ID of the group&#39;s app. | [optional] 
**Name** | Pointer to **string** | The name of the group. | [optional] 
**Description** | Pointer to **string** | A description of the group. | [optional] 
**AdminOwnerId** | Pointer to **string** | The ID of the owner of the group. | [optional] 
**GroupLeaderUserIds** | Pointer to **[]string** | A list of User IDs for the group leaders of the group | [optional] 
**RemoteId** | Pointer to **string** | The ID of the remote. | [optional] 
**RemoteName** | Pointer to **string** | The name of the remote. | [optional] 
**GroupType** | Pointer to [**GroupTypeEnum**](GroupTypeEnum.md) |  | [optional] 
**MaxDuration** | Pointer to **int32** | The maximum duration for which the group can be requested (in minutes). | [optional] 
**RecommendedDuration** | Pointer to **int32** | The recommended duration for which the group should be requested (in minutes). -1 represents an indefinite duration. | [optional] 
**RequireManagerApproval** | Pointer to **bool** | A bool representing whether or not access requests to the group require manager approval. | [optional] 
**RequireSupportTicket** | Pointer to **bool** | A bool representing whether or not access requests to the group require an access ticket. | [optional] 
**RequireMfaToApprove** | Pointer to **bool** | A bool representing whether or not to require MFA for reviewers to approve requests for this group. | [optional] 
**RequireMfaToRequest** | Pointer to **bool** | A bool representing whether or not to require MFA for requesting access to this group. | [optional] 
**AutoApproval** | Pointer to **bool** | A bool representing whether or not to automatically approve requests to this group. | [optional] 
**RequestTemplateId** | Pointer to **string** | The ID of the associated request template. | [optional] 
**ConfigurationTemplateId** | Pointer to **string** | The ID of the associated configuration template. | [optional] 
**GroupBindingId** | Pointer to **string** | The ID of the associated group binding. | [optional] 
**IsRequestable** | Pointer to **bool** | A bool representing whether or not to allow access requests to this group. | [optional] 
**RequestConfigurations** | Pointer to [**[]RequestConfiguration**](RequestConfiguration.md) | A list of request configurations for this group. | [optional] 
**RequestConfigurationList** | Pointer to [**[]RequestConfiguration**](RequestConfiguration.md) | A list of request configurations for this group. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**Metadata** | Pointer to **string** | JSON metadata about the remote group. Only set for items linked to remote systems. See [this guide](https://docs.opal.dev/reference/end-system-objects) for details. | [optional] 
**RemoteInfo** | Pointer to [**GroupRemoteInfo**](GroupRemoteInfo.md) |  | [optional] 
**CustomRequestNotification** | Pointer to **string** | Custom request notification sent to the requester when the request is approved. | [optional] 
**RiskSensitivity** | Pointer to [**RiskSensitivityEnum**](RiskSensitivityEnum.md) | The risk sensitivity level for the group. When an override is set, this field will match that. | [optional] [readonly] 
**RiskSensitivityOverride** | Pointer to [**RiskSensitivityEnum**](RiskSensitivityEnum.md) |  | [optional] 
**LastSuccessfulSync** | Pointer to [**SyncTask**](SyncTask.md) |  | [optional] 

## Methods

### NewGroup

`func NewGroup(groupId string, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *Group) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Group) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Group) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetAppId

`func (o *Group) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Group) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Group) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Group) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Group) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Group) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdminOwnerId

`func (o *Group) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *Group) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *Group) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.

### HasAdminOwnerId

`func (o *Group) HasAdminOwnerId() bool`

HasAdminOwnerId returns a boolean if a field has been set.

### GetGroupLeaderUserIds

`func (o *Group) GetGroupLeaderUserIds() []string`

GetGroupLeaderUserIds returns the GroupLeaderUserIds field if non-nil, zero value otherwise.

### GetGroupLeaderUserIdsOk

`func (o *Group) GetGroupLeaderUserIdsOk() (*[]string, bool)`

GetGroupLeaderUserIdsOk returns a tuple with the GroupLeaderUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLeaderUserIds

`func (o *Group) SetGroupLeaderUserIds(v []string)`

SetGroupLeaderUserIds sets GroupLeaderUserIds field to given value.

### HasGroupLeaderUserIds

`func (o *Group) HasGroupLeaderUserIds() bool`

HasGroupLeaderUserIds returns a boolean if a field has been set.

### GetRemoteId

`func (o *Group) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Group) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Group) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Group) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetRemoteName

`func (o *Group) GetRemoteName() string`

GetRemoteName returns the RemoteName field if non-nil, zero value otherwise.

### GetRemoteNameOk

`func (o *Group) GetRemoteNameOk() (*string, bool)`

GetRemoteNameOk returns a tuple with the RemoteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteName

`func (o *Group) SetRemoteName(v string)`

SetRemoteName sets RemoteName field to given value.

### HasRemoteName

`func (o *Group) HasRemoteName() bool`

HasRemoteName returns a boolean if a field has been set.

### GetGroupType

`func (o *Group) GetGroupType() GroupTypeEnum`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *Group) GetGroupTypeOk() (*GroupTypeEnum, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *Group) SetGroupType(v GroupTypeEnum)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *Group) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetMaxDuration

`func (o *Group) GetMaxDuration() int32`

GetMaxDuration returns the MaxDuration field if non-nil, zero value otherwise.

### GetMaxDurationOk

`func (o *Group) GetMaxDurationOk() (*int32, bool)`

GetMaxDurationOk returns a tuple with the MaxDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDuration

`func (o *Group) SetMaxDuration(v int32)`

SetMaxDuration sets MaxDuration field to given value.

### HasMaxDuration

`func (o *Group) HasMaxDuration() bool`

HasMaxDuration returns a boolean if a field has been set.

### GetRecommendedDuration

`func (o *Group) GetRecommendedDuration() int32`

GetRecommendedDuration returns the RecommendedDuration field if non-nil, zero value otherwise.

### GetRecommendedDurationOk

`func (o *Group) GetRecommendedDurationOk() (*int32, bool)`

GetRecommendedDurationOk returns a tuple with the RecommendedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedDuration

`func (o *Group) SetRecommendedDuration(v int32)`

SetRecommendedDuration sets RecommendedDuration field to given value.

### HasRecommendedDuration

`func (o *Group) HasRecommendedDuration() bool`

HasRecommendedDuration returns a boolean if a field has been set.

### GetRequireManagerApproval

`func (o *Group) GetRequireManagerApproval() bool`

GetRequireManagerApproval returns the RequireManagerApproval field if non-nil, zero value otherwise.

### GetRequireManagerApprovalOk

`func (o *Group) GetRequireManagerApprovalOk() (*bool, bool)`

GetRequireManagerApprovalOk returns a tuple with the RequireManagerApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireManagerApproval

`func (o *Group) SetRequireManagerApproval(v bool)`

SetRequireManagerApproval sets RequireManagerApproval field to given value.

### HasRequireManagerApproval

`func (o *Group) HasRequireManagerApproval() bool`

HasRequireManagerApproval returns a boolean if a field has been set.

### GetRequireSupportTicket

`func (o *Group) GetRequireSupportTicket() bool`

GetRequireSupportTicket returns the RequireSupportTicket field if non-nil, zero value otherwise.

### GetRequireSupportTicketOk

`func (o *Group) GetRequireSupportTicketOk() (*bool, bool)`

GetRequireSupportTicketOk returns a tuple with the RequireSupportTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSupportTicket

`func (o *Group) SetRequireSupportTicket(v bool)`

SetRequireSupportTicket sets RequireSupportTicket field to given value.

### HasRequireSupportTicket

`func (o *Group) HasRequireSupportTicket() bool`

HasRequireSupportTicket returns a boolean if a field has been set.

### GetRequireMfaToApprove

`func (o *Group) GetRequireMfaToApprove() bool`

GetRequireMfaToApprove returns the RequireMfaToApprove field if non-nil, zero value otherwise.

### GetRequireMfaToApproveOk

`func (o *Group) GetRequireMfaToApproveOk() (*bool, bool)`

GetRequireMfaToApproveOk returns a tuple with the RequireMfaToApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToApprove

`func (o *Group) SetRequireMfaToApprove(v bool)`

SetRequireMfaToApprove sets RequireMfaToApprove field to given value.

### HasRequireMfaToApprove

`func (o *Group) HasRequireMfaToApprove() bool`

HasRequireMfaToApprove returns a boolean if a field has been set.

### GetRequireMfaToRequest

`func (o *Group) GetRequireMfaToRequest() bool`

GetRequireMfaToRequest returns the RequireMfaToRequest field if non-nil, zero value otherwise.

### GetRequireMfaToRequestOk

`func (o *Group) GetRequireMfaToRequestOk() (*bool, bool)`

GetRequireMfaToRequestOk returns a tuple with the RequireMfaToRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToRequest

`func (o *Group) SetRequireMfaToRequest(v bool)`

SetRequireMfaToRequest sets RequireMfaToRequest field to given value.

### HasRequireMfaToRequest

`func (o *Group) HasRequireMfaToRequest() bool`

HasRequireMfaToRequest returns a boolean if a field has been set.

### GetAutoApproval

`func (o *Group) GetAutoApproval() bool`

GetAutoApproval returns the AutoApproval field if non-nil, zero value otherwise.

### GetAutoApprovalOk

`func (o *Group) GetAutoApprovalOk() (*bool, bool)`

GetAutoApprovalOk returns a tuple with the AutoApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproval

`func (o *Group) SetAutoApproval(v bool)`

SetAutoApproval sets AutoApproval field to given value.

### HasAutoApproval

`func (o *Group) HasAutoApproval() bool`

HasAutoApproval returns a boolean if a field has been set.

### GetRequestTemplateId

`func (o *Group) GetRequestTemplateId() string`

GetRequestTemplateId returns the RequestTemplateId field if non-nil, zero value otherwise.

### GetRequestTemplateIdOk

`func (o *Group) GetRequestTemplateIdOk() (*string, bool)`

GetRequestTemplateIdOk returns a tuple with the RequestTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTemplateId

`func (o *Group) SetRequestTemplateId(v string)`

SetRequestTemplateId sets RequestTemplateId field to given value.

### HasRequestTemplateId

`func (o *Group) HasRequestTemplateId() bool`

HasRequestTemplateId returns a boolean if a field has been set.

### GetConfigurationTemplateId

`func (o *Group) GetConfigurationTemplateId() string`

GetConfigurationTemplateId returns the ConfigurationTemplateId field if non-nil, zero value otherwise.

### GetConfigurationTemplateIdOk

`func (o *Group) GetConfigurationTemplateIdOk() (*string, bool)`

GetConfigurationTemplateIdOk returns a tuple with the ConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTemplateId

`func (o *Group) SetConfigurationTemplateId(v string)`

SetConfigurationTemplateId sets ConfigurationTemplateId field to given value.

### HasConfigurationTemplateId

`func (o *Group) HasConfigurationTemplateId() bool`

HasConfigurationTemplateId returns a boolean if a field has been set.

### GetGroupBindingId

`func (o *Group) GetGroupBindingId() string`

GetGroupBindingId returns the GroupBindingId field if non-nil, zero value otherwise.

### GetGroupBindingIdOk

`func (o *Group) GetGroupBindingIdOk() (*string, bool)`

GetGroupBindingIdOk returns a tuple with the GroupBindingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBindingId

`func (o *Group) SetGroupBindingId(v string)`

SetGroupBindingId sets GroupBindingId field to given value.

### HasGroupBindingId

`func (o *Group) HasGroupBindingId() bool`

HasGroupBindingId returns a boolean if a field has been set.

### GetIsRequestable

`func (o *Group) GetIsRequestable() bool`

GetIsRequestable returns the IsRequestable field if non-nil, zero value otherwise.

### GetIsRequestableOk

`func (o *Group) GetIsRequestableOk() (*bool, bool)`

GetIsRequestableOk returns a tuple with the IsRequestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequestable

`func (o *Group) SetIsRequestable(v bool)`

SetIsRequestable sets IsRequestable field to given value.

### HasIsRequestable

`func (o *Group) HasIsRequestable() bool`

HasIsRequestable returns a boolean if a field has been set.

### GetRequestConfigurations

`func (o *Group) GetRequestConfigurations() []RequestConfiguration`

GetRequestConfigurations returns the RequestConfigurations field if non-nil, zero value otherwise.

### GetRequestConfigurationsOk

`func (o *Group) GetRequestConfigurationsOk() (*[]RequestConfiguration, bool)`

GetRequestConfigurationsOk returns a tuple with the RequestConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfigurations

`func (o *Group) SetRequestConfigurations(v []RequestConfiguration)`

SetRequestConfigurations sets RequestConfigurations field to given value.

### HasRequestConfigurations

`func (o *Group) HasRequestConfigurations() bool`

HasRequestConfigurations returns a boolean if a field has been set.

### GetRequestConfigurationList

`func (o *Group) GetRequestConfigurationList() []RequestConfiguration`

GetRequestConfigurationList returns the RequestConfigurationList field if non-nil, zero value otherwise.

### GetRequestConfigurationListOk

`func (o *Group) GetRequestConfigurationListOk() (*[]RequestConfiguration, bool)`

GetRequestConfigurationListOk returns a tuple with the RequestConfigurationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfigurationList

`func (o *Group) SetRequestConfigurationList(v []RequestConfiguration)`

SetRequestConfigurationList sets RequestConfigurationList field to given value.

### HasRequestConfigurationList

`func (o *Group) HasRequestConfigurationList() bool`

HasRequestConfigurationList returns a boolean if a field has been set.

### GetMetadata

`func (o *Group) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Group) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Group) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Group) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRemoteInfo

`func (o *Group) GetRemoteInfo() GroupRemoteInfo`

GetRemoteInfo returns the RemoteInfo field if non-nil, zero value otherwise.

### GetRemoteInfoOk

`func (o *Group) GetRemoteInfoOk() (*GroupRemoteInfo, bool)`

GetRemoteInfoOk returns a tuple with the RemoteInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteInfo

`func (o *Group) SetRemoteInfo(v GroupRemoteInfo)`

SetRemoteInfo sets RemoteInfo field to given value.

### HasRemoteInfo

`func (o *Group) HasRemoteInfo() bool`

HasRemoteInfo returns a boolean if a field has been set.

### GetCustomRequestNotification

`func (o *Group) GetCustomRequestNotification() string`

GetCustomRequestNotification returns the CustomRequestNotification field if non-nil, zero value otherwise.

### GetCustomRequestNotificationOk

`func (o *Group) GetCustomRequestNotificationOk() (*string, bool)`

GetCustomRequestNotificationOk returns a tuple with the CustomRequestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRequestNotification

`func (o *Group) SetCustomRequestNotification(v string)`

SetCustomRequestNotification sets CustomRequestNotification field to given value.

### HasCustomRequestNotification

`func (o *Group) HasCustomRequestNotification() bool`

HasCustomRequestNotification returns a boolean if a field has been set.

### GetRiskSensitivity

`func (o *Group) GetRiskSensitivity() RiskSensitivityEnum`

GetRiskSensitivity returns the RiskSensitivity field if non-nil, zero value otherwise.

### GetRiskSensitivityOk

`func (o *Group) GetRiskSensitivityOk() (*RiskSensitivityEnum, bool)`

GetRiskSensitivityOk returns a tuple with the RiskSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskSensitivity

`func (o *Group) SetRiskSensitivity(v RiskSensitivityEnum)`

SetRiskSensitivity sets RiskSensitivity field to given value.

### HasRiskSensitivity

`func (o *Group) HasRiskSensitivity() bool`

HasRiskSensitivity returns a boolean if a field has been set.

### GetRiskSensitivityOverride

`func (o *Group) GetRiskSensitivityOverride() RiskSensitivityEnum`

GetRiskSensitivityOverride returns the RiskSensitivityOverride field if non-nil, zero value otherwise.

### GetRiskSensitivityOverrideOk

`func (o *Group) GetRiskSensitivityOverrideOk() (*RiskSensitivityEnum, bool)`

GetRiskSensitivityOverrideOk returns a tuple with the RiskSensitivityOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskSensitivityOverride

`func (o *Group) SetRiskSensitivityOverride(v RiskSensitivityEnum)`

SetRiskSensitivityOverride sets RiskSensitivityOverride field to given value.

### HasRiskSensitivityOverride

`func (o *Group) HasRiskSensitivityOverride() bool`

HasRiskSensitivityOverride returns a boolean if a field has been set.

### GetLastSuccessfulSync

`func (o *Group) GetLastSuccessfulSync() SyncTask`

GetLastSuccessfulSync returns the LastSuccessfulSync field if non-nil, zero value otherwise.

### GetLastSuccessfulSyncOk

`func (o *Group) GetLastSuccessfulSyncOk() (*SyncTask, bool)`

GetLastSuccessfulSyncOk returns a tuple with the LastSuccessfulSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulSync

`func (o *Group) SetLastSuccessfulSync(v SyncTask)`

SetLastSuccessfulSync sets LastSuccessfulSync field to given value.

### HasLastSuccessfulSync

`func (o *Group) HasLastSuccessfulSync() bool`

HasLastSuccessfulSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


