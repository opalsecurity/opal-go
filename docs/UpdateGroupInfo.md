# UpdateGroupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The ID of the group. | 
**Name** | Pointer to **string** | The name of the group. | [optional] 
**Description** | Pointer to **string** | A description of the group. | [optional] 
**AdminOwnerId** | Pointer to **string** | The ID of the owner of the group. | [optional] 
**MaxDuration** | Pointer to **int32** | The maximum duration for which the group can be requested (in minutes). Use -1 to set to indefinite. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**RecommendedDuration** | Pointer to **int32** | The recommended duration for which the group should be requested (in minutes). Will be the default value in a request. Use -1 to set to indefinite and 0 to unset. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**RequireManagerApproval** | Pointer to **bool** | A bool representing whether or not access requests to the group require manager approval. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**RequireSupportTicket** | Pointer to **bool** | A bool representing whether or not access requests to the group require an access ticket. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**FolderId** | Pointer to **string** | The ID of the folder that the group is located in. | [optional] 
**RequireMfaToApprove** | Pointer to **bool** | A bool representing whether or not to require MFA for reviewers to approve requests for this group. | [optional] 
**RequireMfaToRequest** | Pointer to **bool** | A bool representing whether or not to require MFA for requesting access to this group. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**AutoApproval** | Pointer to **bool** | A bool representing whether or not to automatically approve requests to this group. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**ConfigurationTemplateId** | Pointer to **string** | The ID of the associated configuration template. | [optional] 
**RequestTemplateId** | Pointer to **string** | The ID of the associated request template. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**IsRequestable** | Pointer to **bool** | A bool representing whether or not to allow access requests to this group. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**GroupLeaderUserIds** | Pointer to **[]string** | A list of User IDs for the group leaders of the group | [optional] 
**ExtensionsDurationInMinutes** | Pointer to **int32** | The duration for which access can be extended (in minutes). | [optional] 
**RequestConfigurations** | Pointer to [**[]RequestConfiguration**](RequestConfiguration.md) | The request configuration list of the configuration template. If not provided, the default request configuration will be used. | [optional] 
**RequestConfigurationList** | Pointer to [**CreateRequestConfigurationInfoList**](CreateRequestConfigurationInfoList.md) | The request configuration list of the configuration template. If not provided, the default request configuration will be used. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**CustomRequestNotification** | Pointer to **string** | Custom request notification sent to the requester when the request is approved. | [optional] 
**RiskSensitivityOverride** | Pointer to [**RiskSensitivityEnum**](RiskSensitivityEnum.md) |  | [optional] 

## Methods

### NewUpdateGroupInfo

`func NewUpdateGroupInfo(groupId string, ) *UpdateGroupInfo`

NewUpdateGroupInfo instantiates a new UpdateGroupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupInfoWithDefaults

`func NewUpdateGroupInfoWithDefaults() *UpdateGroupInfo`

NewUpdateGroupInfoWithDefaults instantiates a new UpdateGroupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *UpdateGroupInfo) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UpdateGroupInfo) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UpdateGroupInfo) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetName

`func (o *UpdateGroupInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGroupInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGroupInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGroupInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateGroupInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGroupInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGroupInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGroupInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdminOwnerId

`func (o *UpdateGroupInfo) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *UpdateGroupInfo) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *UpdateGroupInfo) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.

### HasAdminOwnerId

`func (o *UpdateGroupInfo) HasAdminOwnerId() bool`

HasAdminOwnerId returns a boolean if a field has been set.

### GetMaxDuration

`func (o *UpdateGroupInfo) GetMaxDuration() int32`

GetMaxDuration returns the MaxDuration field if non-nil, zero value otherwise.

### GetMaxDurationOk

`func (o *UpdateGroupInfo) GetMaxDurationOk() (*int32, bool)`

GetMaxDurationOk returns a tuple with the MaxDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDuration

`func (o *UpdateGroupInfo) SetMaxDuration(v int32)`

SetMaxDuration sets MaxDuration field to given value.

### HasMaxDuration

`func (o *UpdateGroupInfo) HasMaxDuration() bool`

HasMaxDuration returns a boolean if a field has been set.

### GetRecommendedDuration

`func (o *UpdateGroupInfo) GetRecommendedDuration() int32`

GetRecommendedDuration returns the RecommendedDuration field if non-nil, zero value otherwise.

### GetRecommendedDurationOk

`func (o *UpdateGroupInfo) GetRecommendedDurationOk() (*int32, bool)`

GetRecommendedDurationOk returns a tuple with the RecommendedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedDuration

`func (o *UpdateGroupInfo) SetRecommendedDuration(v int32)`

SetRecommendedDuration sets RecommendedDuration field to given value.

### HasRecommendedDuration

`func (o *UpdateGroupInfo) HasRecommendedDuration() bool`

HasRecommendedDuration returns a boolean if a field has been set.

### GetRequireManagerApproval

`func (o *UpdateGroupInfo) GetRequireManagerApproval() bool`

GetRequireManagerApproval returns the RequireManagerApproval field if non-nil, zero value otherwise.

### GetRequireManagerApprovalOk

`func (o *UpdateGroupInfo) GetRequireManagerApprovalOk() (*bool, bool)`

GetRequireManagerApprovalOk returns a tuple with the RequireManagerApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireManagerApproval

`func (o *UpdateGroupInfo) SetRequireManagerApproval(v bool)`

SetRequireManagerApproval sets RequireManagerApproval field to given value.

### HasRequireManagerApproval

`func (o *UpdateGroupInfo) HasRequireManagerApproval() bool`

HasRequireManagerApproval returns a boolean if a field has been set.

### GetRequireSupportTicket

`func (o *UpdateGroupInfo) GetRequireSupportTicket() bool`

GetRequireSupportTicket returns the RequireSupportTicket field if non-nil, zero value otherwise.

### GetRequireSupportTicketOk

`func (o *UpdateGroupInfo) GetRequireSupportTicketOk() (*bool, bool)`

GetRequireSupportTicketOk returns a tuple with the RequireSupportTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSupportTicket

`func (o *UpdateGroupInfo) SetRequireSupportTicket(v bool)`

SetRequireSupportTicket sets RequireSupportTicket field to given value.

### HasRequireSupportTicket

`func (o *UpdateGroupInfo) HasRequireSupportTicket() bool`

HasRequireSupportTicket returns a boolean if a field has been set.

### GetFolderId

`func (o *UpdateGroupInfo) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *UpdateGroupInfo) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *UpdateGroupInfo) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *UpdateGroupInfo) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetRequireMfaToApprove

`func (o *UpdateGroupInfo) GetRequireMfaToApprove() bool`

GetRequireMfaToApprove returns the RequireMfaToApprove field if non-nil, zero value otherwise.

### GetRequireMfaToApproveOk

`func (o *UpdateGroupInfo) GetRequireMfaToApproveOk() (*bool, bool)`

GetRequireMfaToApproveOk returns a tuple with the RequireMfaToApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToApprove

`func (o *UpdateGroupInfo) SetRequireMfaToApprove(v bool)`

SetRequireMfaToApprove sets RequireMfaToApprove field to given value.

### HasRequireMfaToApprove

`func (o *UpdateGroupInfo) HasRequireMfaToApprove() bool`

HasRequireMfaToApprove returns a boolean if a field has been set.

### GetRequireMfaToRequest

`func (o *UpdateGroupInfo) GetRequireMfaToRequest() bool`

GetRequireMfaToRequest returns the RequireMfaToRequest field if non-nil, zero value otherwise.

### GetRequireMfaToRequestOk

`func (o *UpdateGroupInfo) GetRequireMfaToRequestOk() (*bool, bool)`

GetRequireMfaToRequestOk returns a tuple with the RequireMfaToRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToRequest

`func (o *UpdateGroupInfo) SetRequireMfaToRequest(v bool)`

SetRequireMfaToRequest sets RequireMfaToRequest field to given value.

### HasRequireMfaToRequest

`func (o *UpdateGroupInfo) HasRequireMfaToRequest() bool`

HasRequireMfaToRequest returns a boolean if a field has been set.

### GetAutoApproval

`func (o *UpdateGroupInfo) GetAutoApproval() bool`

GetAutoApproval returns the AutoApproval field if non-nil, zero value otherwise.

### GetAutoApprovalOk

`func (o *UpdateGroupInfo) GetAutoApprovalOk() (*bool, bool)`

GetAutoApprovalOk returns a tuple with the AutoApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproval

`func (o *UpdateGroupInfo) SetAutoApproval(v bool)`

SetAutoApproval sets AutoApproval field to given value.

### HasAutoApproval

`func (o *UpdateGroupInfo) HasAutoApproval() bool`

HasAutoApproval returns a boolean if a field has been set.

### GetConfigurationTemplateId

`func (o *UpdateGroupInfo) GetConfigurationTemplateId() string`

GetConfigurationTemplateId returns the ConfigurationTemplateId field if non-nil, zero value otherwise.

### GetConfigurationTemplateIdOk

`func (o *UpdateGroupInfo) GetConfigurationTemplateIdOk() (*string, bool)`

GetConfigurationTemplateIdOk returns a tuple with the ConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTemplateId

`func (o *UpdateGroupInfo) SetConfigurationTemplateId(v string)`

SetConfigurationTemplateId sets ConfigurationTemplateId field to given value.

### HasConfigurationTemplateId

`func (o *UpdateGroupInfo) HasConfigurationTemplateId() bool`

HasConfigurationTemplateId returns a boolean if a field has been set.

### GetRequestTemplateId

`func (o *UpdateGroupInfo) GetRequestTemplateId() string`

GetRequestTemplateId returns the RequestTemplateId field if non-nil, zero value otherwise.

### GetRequestTemplateIdOk

`func (o *UpdateGroupInfo) GetRequestTemplateIdOk() (*string, bool)`

GetRequestTemplateIdOk returns a tuple with the RequestTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTemplateId

`func (o *UpdateGroupInfo) SetRequestTemplateId(v string)`

SetRequestTemplateId sets RequestTemplateId field to given value.

### HasRequestTemplateId

`func (o *UpdateGroupInfo) HasRequestTemplateId() bool`

HasRequestTemplateId returns a boolean if a field has been set.

### GetIsRequestable

`func (o *UpdateGroupInfo) GetIsRequestable() bool`

GetIsRequestable returns the IsRequestable field if non-nil, zero value otherwise.

### GetIsRequestableOk

`func (o *UpdateGroupInfo) GetIsRequestableOk() (*bool, bool)`

GetIsRequestableOk returns a tuple with the IsRequestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequestable

`func (o *UpdateGroupInfo) SetIsRequestable(v bool)`

SetIsRequestable sets IsRequestable field to given value.

### HasIsRequestable

`func (o *UpdateGroupInfo) HasIsRequestable() bool`

HasIsRequestable returns a boolean if a field has been set.

### GetGroupLeaderUserIds

`func (o *UpdateGroupInfo) GetGroupLeaderUserIds() []string`

GetGroupLeaderUserIds returns the GroupLeaderUserIds field if non-nil, zero value otherwise.

### GetGroupLeaderUserIdsOk

`func (o *UpdateGroupInfo) GetGroupLeaderUserIdsOk() (*[]string, bool)`

GetGroupLeaderUserIdsOk returns a tuple with the GroupLeaderUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLeaderUserIds

`func (o *UpdateGroupInfo) SetGroupLeaderUserIds(v []string)`

SetGroupLeaderUserIds sets GroupLeaderUserIds field to given value.

### HasGroupLeaderUserIds

`func (o *UpdateGroupInfo) HasGroupLeaderUserIds() bool`

HasGroupLeaderUserIds returns a boolean if a field has been set.

### GetExtensionsDurationInMinutes

`func (o *UpdateGroupInfo) GetExtensionsDurationInMinutes() int32`

GetExtensionsDurationInMinutes returns the ExtensionsDurationInMinutes field if non-nil, zero value otherwise.

### GetExtensionsDurationInMinutesOk

`func (o *UpdateGroupInfo) GetExtensionsDurationInMinutesOk() (*int32, bool)`

GetExtensionsDurationInMinutesOk returns a tuple with the ExtensionsDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionsDurationInMinutes

`func (o *UpdateGroupInfo) SetExtensionsDurationInMinutes(v int32)`

SetExtensionsDurationInMinutes sets ExtensionsDurationInMinutes field to given value.

### HasExtensionsDurationInMinutes

`func (o *UpdateGroupInfo) HasExtensionsDurationInMinutes() bool`

HasExtensionsDurationInMinutes returns a boolean if a field has been set.

### GetRequestConfigurations

`func (o *UpdateGroupInfo) GetRequestConfigurations() []RequestConfiguration`

GetRequestConfigurations returns the RequestConfigurations field if non-nil, zero value otherwise.

### GetRequestConfigurationsOk

`func (o *UpdateGroupInfo) GetRequestConfigurationsOk() (*[]RequestConfiguration, bool)`

GetRequestConfigurationsOk returns a tuple with the RequestConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfigurations

`func (o *UpdateGroupInfo) SetRequestConfigurations(v []RequestConfiguration)`

SetRequestConfigurations sets RequestConfigurations field to given value.

### HasRequestConfigurations

`func (o *UpdateGroupInfo) HasRequestConfigurations() bool`

HasRequestConfigurations returns a boolean if a field has been set.

### GetRequestConfigurationList

`func (o *UpdateGroupInfo) GetRequestConfigurationList() CreateRequestConfigurationInfoList`

GetRequestConfigurationList returns the RequestConfigurationList field if non-nil, zero value otherwise.

### GetRequestConfigurationListOk

`func (o *UpdateGroupInfo) GetRequestConfigurationListOk() (*CreateRequestConfigurationInfoList, bool)`

GetRequestConfigurationListOk returns a tuple with the RequestConfigurationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfigurationList

`func (o *UpdateGroupInfo) SetRequestConfigurationList(v CreateRequestConfigurationInfoList)`

SetRequestConfigurationList sets RequestConfigurationList field to given value.

### HasRequestConfigurationList

`func (o *UpdateGroupInfo) HasRequestConfigurationList() bool`

HasRequestConfigurationList returns a boolean if a field has been set.

### GetCustomRequestNotification

`func (o *UpdateGroupInfo) GetCustomRequestNotification() string`

GetCustomRequestNotification returns the CustomRequestNotification field if non-nil, zero value otherwise.

### GetCustomRequestNotificationOk

`func (o *UpdateGroupInfo) GetCustomRequestNotificationOk() (*string, bool)`

GetCustomRequestNotificationOk returns a tuple with the CustomRequestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRequestNotification

`func (o *UpdateGroupInfo) SetCustomRequestNotification(v string)`

SetCustomRequestNotification sets CustomRequestNotification field to given value.

### HasCustomRequestNotification

`func (o *UpdateGroupInfo) HasCustomRequestNotification() bool`

HasCustomRequestNotification returns a boolean if a field has been set.

### GetRiskSensitivityOverride

`func (o *UpdateGroupInfo) GetRiskSensitivityOverride() RiskSensitivityEnum`

GetRiskSensitivityOverride returns the RiskSensitivityOverride field if non-nil, zero value otherwise.

### GetRiskSensitivityOverrideOk

`func (o *UpdateGroupInfo) GetRiskSensitivityOverrideOk() (*RiskSensitivityEnum, bool)`

GetRiskSensitivityOverrideOk returns a tuple with the RiskSensitivityOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskSensitivityOverride

`func (o *UpdateGroupInfo) SetRiskSensitivityOverride(v RiskSensitivityEnum)`

SetRiskSensitivityOverride sets RiskSensitivityOverride field to given value.

### HasRiskSensitivityOverride

`func (o *UpdateGroupInfo) HasRiskSensitivityOverride() bool`

HasRiskSensitivityOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


