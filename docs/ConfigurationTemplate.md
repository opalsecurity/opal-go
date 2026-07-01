# ConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationTemplateId** | Pointer to **string** | The ID of the configuration template. | [optional] 
**Name** | Pointer to **string** | The name of the configuration template. | [optional] 
**AdminOwnerId** | Pointer to **string** | The ID of the owner of the configuration template. | [optional] 
**Visibility** | Pointer to [**VisibilityInfo**](VisibilityInfo.md) | The visibility info of the configuration template. | [optional] 
**LinkedAuditMessageChannelIds** | Pointer to **[]string** | The IDs of the audit message channels linked to the configuration template. | [optional] 
**RequestConfigurationId** | Pointer to **string** | The ID of the request configuration linked to the configuration template. | [optional] 
**MemberOncallScheduleIds** | Pointer to **[]string** | The IDs of the on-call schedules linked to the configuration template. | [optional] 
**BreakGlassUserIds** | Pointer to **[]string** | The IDs of the break glass users linked to the configuration template. | [optional] 
**RequireMfaToApprove** | Pointer to **bool** | A bool representing whether or not to require MFA for reviewers to approve requests for this configuration template. | [optional] 
**RequireMfaToConnect** | Pointer to **bool** | A bool representing whether or not to require MFA to connect to resources associated with this configuration template. | [optional] 
**TicketPropagation** | Pointer to [**TicketPropagationConfiguration**](TicketPropagationConfiguration.md) |  | [optional] 
**CustomRequestNotification** | Pointer to **string** | Custom request notification sent upon request approval for this configuration template. | [optional] 

## Methods

### NewConfigurationTemplate

`func NewConfigurationTemplate() *ConfigurationTemplate`

NewConfigurationTemplate instantiates a new ConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationTemplateWithDefaults

`func NewConfigurationTemplateWithDefaults() *ConfigurationTemplate`

NewConfigurationTemplateWithDefaults instantiates a new ConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationTemplateId

`func (o *ConfigurationTemplate) GetConfigurationTemplateId() string`

GetConfigurationTemplateId returns the ConfigurationTemplateId field if non-nil, zero value otherwise.

### GetConfigurationTemplateIdOk

`func (o *ConfigurationTemplate) GetConfigurationTemplateIdOk() (*string, bool)`

GetConfigurationTemplateIdOk returns a tuple with the ConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTemplateId

`func (o *ConfigurationTemplate) SetConfigurationTemplateId(v string)`

SetConfigurationTemplateId sets ConfigurationTemplateId field to given value.

### HasConfigurationTemplateId

`func (o *ConfigurationTemplate) HasConfigurationTemplateId() bool`

HasConfigurationTemplateId returns a boolean if a field has been set.

### GetName

`func (o *ConfigurationTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigurationTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAdminOwnerId

`func (o *ConfigurationTemplate) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *ConfigurationTemplate) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *ConfigurationTemplate) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.

### HasAdminOwnerId

`func (o *ConfigurationTemplate) HasAdminOwnerId() bool`

HasAdminOwnerId returns a boolean if a field has been set.

### GetVisibility

`func (o *ConfigurationTemplate) GetVisibility() VisibilityInfo`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ConfigurationTemplate) GetVisibilityOk() (*VisibilityInfo, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ConfigurationTemplate) SetVisibility(v VisibilityInfo)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ConfigurationTemplate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLinkedAuditMessageChannelIds

`func (o *ConfigurationTemplate) GetLinkedAuditMessageChannelIds() []string`

GetLinkedAuditMessageChannelIds returns the LinkedAuditMessageChannelIds field if non-nil, zero value otherwise.

### GetLinkedAuditMessageChannelIdsOk

`func (o *ConfigurationTemplate) GetLinkedAuditMessageChannelIdsOk() (*[]string, bool)`

GetLinkedAuditMessageChannelIdsOk returns a tuple with the LinkedAuditMessageChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAuditMessageChannelIds

`func (o *ConfigurationTemplate) SetLinkedAuditMessageChannelIds(v []string)`

SetLinkedAuditMessageChannelIds sets LinkedAuditMessageChannelIds field to given value.

### HasLinkedAuditMessageChannelIds

`func (o *ConfigurationTemplate) HasLinkedAuditMessageChannelIds() bool`

HasLinkedAuditMessageChannelIds returns a boolean if a field has been set.

### GetRequestConfigurationId

`func (o *ConfigurationTemplate) GetRequestConfigurationId() string`

GetRequestConfigurationId returns the RequestConfigurationId field if non-nil, zero value otherwise.

### GetRequestConfigurationIdOk

`func (o *ConfigurationTemplate) GetRequestConfigurationIdOk() (*string, bool)`

GetRequestConfigurationIdOk returns a tuple with the RequestConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfigurationId

`func (o *ConfigurationTemplate) SetRequestConfigurationId(v string)`

SetRequestConfigurationId sets RequestConfigurationId field to given value.

### HasRequestConfigurationId

`func (o *ConfigurationTemplate) HasRequestConfigurationId() bool`

HasRequestConfigurationId returns a boolean if a field has been set.

### GetMemberOncallScheduleIds

`func (o *ConfigurationTemplate) GetMemberOncallScheduleIds() []string`

GetMemberOncallScheduleIds returns the MemberOncallScheduleIds field if non-nil, zero value otherwise.

### GetMemberOncallScheduleIdsOk

`func (o *ConfigurationTemplate) GetMemberOncallScheduleIdsOk() (*[]string, bool)`

GetMemberOncallScheduleIdsOk returns a tuple with the MemberOncallScheduleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOncallScheduleIds

`func (o *ConfigurationTemplate) SetMemberOncallScheduleIds(v []string)`

SetMemberOncallScheduleIds sets MemberOncallScheduleIds field to given value.

### HasMemberOncallScheduleIds

`func (o *ConfigurationTemplate) HasMemberOncallScheduleIds() bool`

HasMemberOncallScheduleIds returns a boolean if a field has been set.

### GetBreakGlassUserIds

`func (o *ConfigurationTemplate) GetBreakGlassUserIds() []string`

GetBreakGlassUserIds returns the BreakGlassUserIds field if non-nil, zero value otherwise.

### GetBreakGlassUserIdsOk

`func (o *ConfigurationTemplate) GetBreakGlassUserIdsOk() (*[]string, bool)`

GetBreakGlassUserIdsOk returns a tuple with the BreakGlassUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakGlassUserIds

`func (o *ConfigurationTemplate) SetBreakGlassUserIds(v []string)`

SetBreakGlassUserIds sets BreakGlassUserIds field to given value.

### HasBreakGlassUserIds

`func (o *ConfigurationTemplate) HasBreakGlassUserIds() bool`

HasBreakGlassUserIds returns a boolean if a field has been set.

### GetRequireMfaToApprove

`func (o *ConfigurationTemplate) GetRequireMfaToApprove() bool`

GetRequireMfaToApprove returns the RequireMfaToApprove field if non-nil, zero value otherwise.

### GetRequireMfaToApproveOk

`func (o *ConfigurationTemplate) GetRequireMfaToApproveOk() (*bool, bool)`

GetRequireMfaToApproveOk returns a tuple with the RequireMfaToApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToApprove

`func (o *ConfigurationTemplate) SetRequireMfaToApprove(v bool)`

SetRequireMfaToApprove sets RequireMfaToApprove field to given value.

### HasRequireMfaToApprove

`func (o *ConfigurationTemplate) HasRequireMfaToApprove() bool`

HasRequireMfaToApprove returns a boolean if a field has been set.

### GetRequireMfaToConnect

`func (o *ConfigurationTemplate) GetRequireMfaToConnect() bool`

GetRequireMfaToConnect returns the RequireMfaToConnect field if non-nil, zero value otherwise.

### GetRequireMfaToConnectOk

`func (o *ConfigurationTemplate) GetRequireMfaToConnectOk() (*bool, bool)`

GetRequireMfaToConnectOk returns a tuple with the RequireMfaToConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToConnect

`func (o *ConfigurationTemplate) SetRequireMfaToConnect(v bool)`

SetRequireMfaToConnect sets RequireMfaToConnect field to given value.

### HasRequireMfaToConnect

`func (o *ConfigurationTemplate) HasRequireMfaToConnect() bool`

HasRequireMfaToConnect returns a boolean if a field has been set.

### GetTicketPropagation

`func (o *ConfigurationTemplate) GetTicketPropagation() TicketPropagationConfiguration`

GetTicketPropagation returns the TicketPropagation field if non-nil, zero value otherwise.

### GetTicketPropagationOk

`func (o *ConfigurationTemplate) GetTicketPropagationOk() (*TicketPropagationConfiguration, bool)`

GetTicketPropagationOk returns a tuple with the TicketPropagation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketPropagation

`func (o *ConfigurationTemplate) SetTicketPropagation(v TicketPropagationConfiguration)`

SetTicketPropagation sets TicketPropagation field to given value.

### HasTicketPropagation

`func (o *ConfigurationTemplate) HasTicketPropagation() bool`

HasTicketPropagation returns a boolean if a field has been set.

### GetCustomRequestNotification

`func (o *ConfigurationTemplate) GetCustomRequestNotification() string`

GetCustomRequestNotification returns the CustomRequestNotification field if non-nil, zero value otherwise.

### GetCustomRequestNotificationOk

`func (o *ConfigurationTemplate) GetCustomRequestNotificationOk() (*string, bool)`

GetCustomRequestNotificationOk returns a tuple with the CustomRequestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRequestNotification

`func (o *ConfigurationTemplate) SetCustomRequestNotification(v string)`

SetCustomRequestNotification sets CustomRequestNotification field to given value.

### HasCustomRequestNotification

`func (o *ConfigurationTemplate) HasCustomRequestNotification() bool`

HasCustomRequestNotification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


