# UpdateConfigurationTemplateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationTemplateId** | **string** | The ID of the configuration template. | 
**Name** | Pointer to **string** | The name of the configuration template. | [optional] 
**AdminOwnerId** | Pointer to **string** | The ID of the owner of the configuration template. | [optional] 
**Visibility** | Pointer to [**VisibilityInfo**](VisibilityInfo.md) |  | [optional] 
**LinkedAuditMessageChannelIds** | Pointer to **[]string** | The IDs of the audit message channels linked to the configuration template. | [optional] 
**RequestConfigurationList** | Pointer to [**CreateRequestConfigurationInfoList**](CreateRequestConfigurationInfoList.md) |  | [optional] 
**MemberOncallScheduleIds** | Pointer to **[]string** | The IDs of the on-call schedules linked to the configuration template. | [optional] 
**BreakGlassUserIds** | Pointer to **[]string** | The IDs of the break glass users linked to the configuration template. | [optional] 
**RequireMfaToApprove** | Pointer to **bool** | A bool representing whether or not to require MFA for reviewers to approve requests for this configuration template. | [optional] 
**RequireMfaToConnect** | Pointer to **bool** | A bool representing whether or not to require MFA to connect to resources associated with this configuration template. | [optional] 

## Methods

### NewUpdateConfigurationTemplateInfo

`func NewUpdateConfigurationTemplateInfo(configurationTemplateId string, ) *UpdateConfigurationTemplateInfo`

NewUpdateConfigurationTemplateInfo instantiates a new UpdateConfigurationTemplateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConfigurationTemplateInfoWithDefaults

`func NewUpdateConfigurationTemplateInfoWithDefaults() *UpdateConfigurationTemplateInfo`

NewUpdateConfigurationTemplateInfoWithDefaults instantiates a new UpdateConfigurationTemplateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationTemplateId

`func (o *UpdateConfigurationTemplateInfo) GetConfigurationTemplateId() string`

GetConfigurationTemplateId returns the ConfigurationTemplateId field if non-nil, zero value otherwise.

### GetConfigurationTemplateIdOk

`func (o *UpdateConfigurationTemplateInfo) GetConfigurationTemplateIdOk() (*string, bool)`

GetConfigurationTemplateIdOk returns a tuple with the ConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTemplateId

`func (o *UpdateConfigurationTemplateInfo) SetConfigurationTemplateId(v string)`

SetConfigurationTemplateId sets ConfigurationTemplateId field to given value.


### GetName

`func (o *UpdateConfigurationTemplateInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateConfigurationTemplateInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateConfigurationTemplateInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateConfigurationTemplateInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAdminOwnerId

`func (o *UpdateConfigurationTemplateInfo) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *UpdateConfigurationTemplateInfo) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *UpdateConfigurationTemplateInfo) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.

### HasAdminOwnerId

`func (o *UpdateConfigurationTemplateInfo) HasAdminOwnerId() bool`

HasAdminOwnerId returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateConfigurationTemplateInfo) GetVisibility() VisibilityInfo`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateConfigurationTemplateInfo) GetVisibilityOk() (*VisibilityInfo, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateConfigurationTemplateInfo) SetVisibility(v VisibilityInfo)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateConfigurationTemplateInfo) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLinkedAuditMessageChannelIds

`func (o *UpdateConfigurationTemplateInfo) GetLinkedAuditMessageChannelIds() []string`

GetLinkedAuditMessageChannelIds returns the LinkedAuditMessageChannelIds field if non-nil, zero value otherwise.

### GetLinkedAuditMessageChannelIdsOk

`func (o *UpdateConfigurationTemplateInfo) GetLinkedAuditMessageChannelIdsOk() (*[]string, bool)`

GetLinkedAuditMessageChannelIdsOk returns a tuple with the LinkedAuditMessageChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAuditMessageChannelIds

`func (o *UpdateConfigurationTemplateInfo) SetLinkedAuditMessageChannelIds(v []string)`

SetLinkedAuditMessageChannelIds sets LinkedAuditMessageChannelIds field to given value.

### HasLinkedAuditMessageChannelIds

`func (o *UpdateConfigurationTemplateInfo) HasLinkedAuditMessageChannelIds() bool`

HasLinkedAuditMessageChannelIds returns a boolean if a field has been set.

### GetRequestConfigurationList

`func (o *UpdateConfigurationTemplateInfo) GetRequestConfigurationList() CreateRequestConfigurationInfoList`

GetRequestConfigurationList returns the RequestConfigurationList field if non-nil, zero value otherwise.

### GetRequestConfigurationListOk

`func (o *UpdateConfigurationTemplateInfo) GetRequestConfigurationListOk() (*CreateRequestConfigurationInfoList, bool)`

GetRequestConfigurationListOk returns a tuple with the RequestConfigurationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfigurationList

`func (o *UpdateConfigurationTemplateInfo) SetRequestConfigurationList(v CreateRequestConfigurationInfoList)`

SetRequestConfigurationList sets RequestConfigurationList field to given value.

### HasRequestConfigurationList

`func (o *UpdateConfigurationTemplateInfo) HasRequestConfigurationList() bool`

HasRequestConfigurationList returns a boolean if a field has been set.

### GetMemberOncallScheduleIds

`func (o *UpdateConfigurationTemplateInfo) GetMemberOncallScheduleIds() []string`

GetMemberOncallScheduleIds returns the MemberOncallScheduleIds field if non-nil, zero value otherwise.

### GetMemberOncallScheduleIdsOk

`func (o *UpdateConfigurationTemplateInfo) GetMemberOncallScheduleIdsOk() (*[]string, bool)`

GetMemberOncallScheduleIdsOk returns a tuple with the MemberOncallScheduleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOncallScheduleIds

`func (o *UpdateConfigurationTemplateInfo) SetMemberOncallScheduleIds(v []string)`

SetMemberOncallScheduleIds sets MemberOncallScheduleIds field to given value.

### HasMemberOncallScheduleIds

`func (o *UpdateConfigurationTemplateInfo) HasMemberOncallScheduleIds() bool`

HasMemberOncallScheduleIds returns a boolean if a field has been set.

### GetBreakGlassUserIds

`func (o *UpdateConfigurationTemplateInfo) GetBreakGlassUserIds() []string`

GetBreakGlassUserIds returns the BreakGlassUserIds field if non-nil, zero value otherwise.

### GetBreakGlassUserIdsOk

`func (o *UpdateConfigurationTemplateInfo) GetBreakGlassUserIdsOk() (*[]string, bool)`

GetBreakGlassUserIdsOk returns a tuple with the BreakGlassUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakGlassUserIds

`func (o *UpdateConfigurationTemplateInfo) SetBreakGlassUserIds(v []string)`

SetBreakGlassUserIds sets BreakGlassUserIds field to given value.

### HasBreakGlassUserIds

`func (o *UpdateConfigurationTemplateInfo) HasBreakGlassUserIds() bool`

HasBreakGlassUserIds returns a boolean if a field has been set.

### GetRequireMfaToApprove

`func (o *UpdateConfigurationTemplateInfo) GetRequireMfaToApprove() bool`

GetRequireMfaToApprove returns the RequireMfaToApprove field if non-nil, zero value otherwise.

### GetRequireMfaToApproveOk

`func (o *UpdateConfigurationTemplateInfo) GetRequireMfaToApproveOk() (*bool, bool)`

GetRequireMfaToApproveOk returns a tuple with the RequireMfaToApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToApprove

`func (o *UpdateConfigurationTemplateInfo) SetRequireMfaToApprove(v bool)`

SetRequireMfaToApprove sets RequireMfaToApprove field to given value.

### HasRequireMfaToApprove

`func (o *UpdateConfigurationTemplateInfo) HasRequireMfaToApprove() bool`

HasRequireMfaToApprove returns a boolean if a field has been set.

### GetRequireMfaToConnect

`func (o *UpdateConfigurationTemplateInfo) GetRequireMfaToConnect() bool`

GetRequireMfaToConnect returns the RequireMfaToConnect field if non-nil, zero value otherwise.

### GetRequireMfaToConnectOk

`func (o *UpdateConfigurationTemplateInfo) GetRequireMfaToConnectOk() (*bool, bool)`

GetRequireMfaToConnectOk returns a tuple with the RequireMfaToConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToConnect

`func (o *UpdateConfigurationTemplateInfo) SetRequireMfaToConnect(v bool)`

SetRequireMfaToConnect sets RequireMfaToConnect field to given value.

### HasRequireMfaToConnect

`func (o *UpdateConfigurationTemplateInfo) HasRequireMfaToConnect() bool`

HasRequireMfaToConnect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


