# CreateGroupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the remote group. | 
**Description** | Pointer to **string** | A description of the remote group. | [optional] 
**GroupType** | [**GroupTypeEnum**](GroupTypeEnum.md) |  | 
**AppId** | **string** | The ID of the app for the group. | 
**RemoteInfo** | Pointer to [**GroupRemoteInfo**](GroupRemoteInfo.md) |  | [optional] 
**RemoteGroupId** | Pointer to **string** | Deprecated - use remote_info instead. The ID of the group on the remote system. Include only for items linked to remote systems. See [this guide](https://docs.opal.dev/reference/end-system-objects) for details on how to specify this field. | [optional] 
**Metadata** | Pointer to **string** | Deprecated - use remote_info instead.  JSON metadata about the remote group. Include only for items linked to remote systems. See [this guide](https://docs.opal.dev/reference/end-system-objects) for details on how to specify this field. The required format is dependent on group_type and should have the following schema: &lt;style type&#x3D;\&quot;text/css\&quot;&gt; code {max-height:300px !important} &lt;/style&gt; &#x60;&#x60;&#x60;json {   \&quot;$schema\&quot;: \&quot;http://json-schema.org/draft-04/schema#\&quot;,   \&quot;title\&quot;: \&quot;Group Metadata\&quot;,   \&quot;properties\&quot;: {     \&quot;ad_group\&quot;: {       \&quot;properties\&quot;: {         \&quot;object_guid\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;object_guid\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;Active Directory Group\&quot;     },     \&quot;duo_group\&quot;: {       \&quot;properties\&quot;: {         \&quot;group_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;group_id\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;Duo Group\&quot;     },     \&quot;git_hub_team\&quot;: {       \&quot;properties\&quot;: {         \&quot;org_name\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;team_slug\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;org_name\&quot;, \&quot;team_slug\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;GitHub Team\&quot;     },     \&quot;google_groups_group\&quot;: {       \&quot;properties\&quot;: {         \&quot;group_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;group_id\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;Google Groups Group\&quot;     },     \&quot;ldap_group\&quot;: {       \&quot;properties\&quot;: {         \&quot;group_uid\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;group_uid\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;LDAP Group\&quot;     },     \&quot;okta_directory_group\&quot;: {       \&quot;properties\&quot;: {         \&quot;group_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;group_id\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;Okta Directory Group\&quot;     }   },   \&quot;additionalProperties\&quot;: false,   \&quot;minProperties\&quot;: 1,   \&quot;maxProperties\&quot;: 1,   \&quot;type\&quot;: \&quot;object\&quot; } &#x60;&#x60;&#x60; | [optional] 
**CustomRequestNotification** | Pointer to **string** | Custom request notification sent upon request approval. | [optional] 
**RiskSensitivityOverride** | Pointer to [**RiskSensitivityEnum**](RiskSensitivityEnum.md) |  | [optional] 

## Methods

### NewCreateGroupInfo

`func NewCreateGroupInfo(name string, groupType GroupTypeEnum, appId string, ) *CreateGroupInfo`

NewCreateGroupInfo instantiates a new CreateGroupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupInfoWithDefaults

`func NewCreateGroupInfoWithDefaults() *CreateGroupInfo`

NewCreateGroupInfoWithDefaults instantiates a new CreateGroupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGroupInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroupInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroupInfo) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateGroupInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGroupInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGroupInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGroupInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupType

`func (o *CreateGroupInfo) GetGroupType() GroupTypeEnum`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *CreateGroupInfo) GetGroupTypeOk() (*GroupTypeEnum, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *CreateGroupInfo) SetGroupType(v GroupTypeEnum)`

SetGroupType sets GroupType field to given value.


### GetAppId

`func (o *CreateGroupInfo) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CreateGroupInfo) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CreateGroupInfo) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetRemoteInfo

`func (o *CreateGroupInfo) GetRemoteInfo() GroupRemoteInfo`

GetRemoteInfo returns the RemoteInfo field if non-nil, zero value otherwise.

### GetRemoteInfoOk

`func (o *CreateGroupInfo) GetRemoteInfoOk() (*GroupRemoteInfo, bool)`

GetRemoteInfoOk returns a tuple with the RemoteInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteInfo

`func (o *CreateGroupInfo) SetRemoteInfo(v GroupRemoteInfo)`

SetRemoteInfo sets RemoteInfo field to given value.

### HasRemoteInfo

`func (o *CreateGroupInfo) HasRemoteInfo() bool`

HasRemoteInfo returns a boolean if a field has been set.

### GetRemoteGroupId

`func (o *CreateGroupInfo) GetRemoteGroupId() string`

GetRemoteGroupId returns the RemoteGroupId field if non-nil, zero value otherwise.

### GetRemoteGroupIdOk

`func (o *CreateGroupInfo) GetRemoteGroupIdOk() (*string, bool)`

GetRemoteGroupIdOk returns a tuple with the RemoteGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGroupId

`func (o *CreateGroupInfo) SetRemoteGroupId(v string)`

SetRemoteGroupId sets RemoteGroupId field to given value.

### HasRemoteGroupId

`func (o *CreateGroupInfo) HasRemoteGroupId() bool`

HasRemoteGroupId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateGroupInfo) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateGroupInfo) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateGroupInfo) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateGroupInfo) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCustomRequestNotification

`func (o *CreateGroupInfo) GetCustomRequestNotification() string`

GetCustomRequestNotification returns the CustomRequestNotification field if non-nil, zero value otherwise.

### GetCustomRequestNotificationOk

`func (o *CreateGroupInfo) GetCustomRequestNotificationOk() (*string, bool)`

GetCustomRequestNotificationOk returns a tuple with the CustomRequestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRequestNotification

`func (o *CreateGroupInfo) SetCustomRequestNotification(v string)`

SetCustomRequestNotification sets CustomRequestNotification field to given value.

### HasCustomRequestNotification

`func (o *CreateGroupInfo) HasCustomRequestNotification() bool`

HasCustomRequestNotification returns a boolean if a field has been set.

### GetRiskSensitivityOverride

`func (o *CreateGroupInfo) GetRiskSensitivityOverride() RiskSensitivityEnum`

GetRiskSensitivityOverride returns the RiskSensitivityOverride field if non-nil, zero value otherwise.

### GetRiskSensitivityOverrideOk

`func (o *CreateGroupInfo) GetRiskSensitivityOverrideOk() (*RiskSensitivityEnum, bool)`

GetRiskSensitivityOverrideOk returns a tuple with the RiskSensitivityOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskSensitivityOverride

`func (o *CreateGroupInfo) SetRiskSensitivityOverride(v RiskSensitivityEnum)`

SetRiskSensitivityOverride sets RiskSensitivityOverride field to given value.

### HasRiskSensitivityOverride

`func (o *CreateGroupInfo) HasRiskSensitivityOverride() bool`

HasRiskSensitivityOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


