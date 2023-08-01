# CreateConfigurationTemplateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminOwnerId** | **string** | The ID of the owner of the configuration template. | 
**Visibility** | [**VisibilityInfo**](VisibilityInfo.md) |  | 
**LinkedAuditMessageChannelIds** | Pointer to **[]string** | The IDs of the audit message channels linked to the configuration template. | [optional] 
**RequestConfigurationId** | Pointer to **string** | The ID of the request configuration linked to the configuration template. | [optional] 
**MemberOncallScheduleIds** | Pointer to **[]string** | The IDs of the on-call schedules linked to the configuration template. | [optional] 
**BreakGlassUserIds** | Pointer to **[]string** | The IDs of the break glass users linked to the configuration template. | [optional] 
**RequireMfaToApprove** | **bool** | A bool representing whether or not to require MFA for reviewers to approve requests for this configuration template. | 
**RequireMfaToConnect** | **bool** | A bool representing whether or not to require MFA to connect to resources associated with this configuration template. | 
**Name** | **string** | The name of the configuration template. | 

## Methods

### NewCreateConfigurationTemplateInfo

`func NewCreateConfigurationTemplateInfo(adminOwnerId string, visibility VisibilityInfo, requireMfaToApprove bool, requireMfaToConnect bool, name string, ) *CreateConfigurationTemplateInfo`

NewCreateConfigurationTemplateInfo instantiates a new CreateConfigurationTemplateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConfigurationTemplateInfoWithDefaults

`func NewCreateConfigurationTemplateInfoWithDefaults() *CreateConfigurationTemplateInfo`

NewCreateConfigurationTemplateInfoWithDefaults instantiates a new CreateConfigurationTemplateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminOwnerId

`func (o *CreateConfigurationTemplateInfo) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *CreateConfigurationTemplateInfo) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *CreateConfigurationTemplateInfo) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.


### GetVisibility

`func (o *CreateConfigurationTemplateInfo) GetVisibility() VisibilityInfo`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateConfigurationTemplateInfo) GetVisibilityOk() (*VisibilityInfo, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateConfigurationTemplateInfo) SetVisibility(v VisibilityInfo)`

SetVisibility sets Visibility field to given value.


### GetLinkedAuditMessageChannelIds

`func (o *CreateConfigurationTemplateInfo) GetLinkedAuditMessageChannelIds() []string`

GetLinkedAuditMessageChannelIds returns the LinkedAuditMessageChannelIds field if non-nil, zero value otherwise.

### GetLinkedAuditMessageChannelIdsOk

`func (o *CreateConfigurationTemplateInfo) GetLinkedAuditMessageChannelIdsOk() (*[]string, bool)`

GetLinkedAuditMessageChannelIdsOk returns a tuple with the LinkedAuditMessageChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAuditMessageChannelIds

`func (o *CreateConfigurationTemplateInfo) SetLinkedAuditMessageChannelIds(v []string)`

SetLinkedAuditMessageChannelIds sets LinkedAuditMessageChannelIds field to given value.

### HasLinkedAuditMessageChannelIds

`func (o *CreateConfigurationTemplateInfo) HasLinkedAuditMessageChannelIds() bool`

HasLinkedAuditMessageChannelIds returns a boolean if a field has been set.

### GetRequestConfigurationId

`func (o *CreateConfigurationTemplateInfo) GetRequestConfigurationId() string`

GetRequestConfigurationId returns the RequestConfigurationId field if non-nil, zero value otherwise.

### GetRequestConfigurationIdOk

`func (o *CreateConfigurationTemplateInfo) GetRequestConfigurationIdOk() (*string, bool)`

GetRequestConfigurationIdOk returns a tuple with the RequestConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfigurationId

`func (o *CreateConfigurationTemplateInfo) SetRequestConfigurationId(v string)`

SetRequestConfigurationId sets RequestConfigurationId field to given value.

### HasRequestConfigurationId

`func (o *CreateConfigurationTemplateInfo) HasRequestConfigurationId() bool`

HasRequestConfigurationId returns a boolean if a field has been set.

### GetMemberOncallScheduleIds

`func (o *CreateConfigurationTemplateInfo) GetMemberOncallScheduleIds() []string`

GetMemberOncallScheduleIds returns the MemberOncallScheduleIds field if non-nil, zero value otherwise.

### GetMemberOncallScheduleIdsOk

`func (o *CreateConfigurationTemplateInfo) GetMemberOncallScheduleIdsOk() (*[]string, bool)`

GetMemberOncallScheduleIdsOk returns a tuple with the MemberOncallScheduleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOncallScheduleIds

`func (o *CreateConfigurationTemplateInfo) SetMemberOncallScheduleIds(v []string)`

SetMemberOncallScheduleIds sets MemberOncallScheduleIds field to given value.

### HasMemberOncallScheduleIds

`func (o *CreateConfigurationTemplateInfo) HasMemberOncallScheduleIds() bool`

HasMemberOncallScheduleIds returns a boolean if a field has been set.

### GetBreakGlassUserIds

`func (o *CreateConfigurationTemplateInfo) GetBreakGlassUserIds() []string`

GetBreakGlassUserIds returns the BreakGlassUserIds field if non-nil, zero value otherwise.

### GetBreakGlassUserIdsOk

`func (o *CreateConfigurationTemplateInfo) GetBreakGlassUserIdsOk() (*[]string, bool)`

GetBreakGlassUserIdsOk returns a tuple with the BreakGlassUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakGlassUserIds

`func (o *CreateConfigurationTemplateInfo) SetBreakGlassUserIds(v []string)`

SetBreakGlassUserIds sets BreakGlassUserIds field to given value.

### HasBreakGlassUserIds

`func (o *CreateConfigurationTemplateInfo) HasBreakGlassUserIds() bool`

HasBreakGlassUserIds returns a boolean if a field has been set.

### GetRequireMfaToApprove

`func (o *CreateConfigurationTemplateInfo) GetRequireMfaToApprove() bool`

GetRequireMfaToApprove returns the RequireMfaToApprove field if non-nil, zero value otherwise.

### GetRequireMfaToApproveOk

`func (o *CreateConfigurationTemplateInfo) GetRequireMfaToApproveOk() (*bool, bool)`

GetRequireMfaToApproveOk returns a tuple with the RequireMfaToApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToApprove

`func (o *CreateConfigurationTemplateInfo) SetRequireMfaToApprove(v bool)`

SetRequireMfaToApprove sets RequireMfaToApprove field to given value.


### GetRequireMfaToConnect

`func (o *CreateConfigurationTemplateInfo) GetRequireMfaToConnect() bool`

GetRequireMfaToConnect returns the RequireMfaToConnect field if non-nil, zero value otherwise.

### GetRequireMfaToConnectOk

`func (o *CreateConfigurationTemplateInfo) GetRequireMfaToConnectOk() (*bool, bool)`

GetRequireMfaToConnectOk returns a tuple with the RequireMfaToConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToConnect

`func (o *CreateConfigurationTemplateInfo) SetRequireMfaToConnect(v bool)`

SetRequireMfaToConnect sets RequireMfaToConnect field to given value.


### GetName

`func (o *CreateConfigurationTemplateInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateConfigurationTemplateInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateConfigurationTemplateInfo) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


