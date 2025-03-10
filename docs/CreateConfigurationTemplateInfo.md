# CreateConfigurationTemplateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminOwnerId** | **string** | The ID of the owner of the configuration template. | 
**Visibility** | [**VisibilityInfo**](VisibilityInfo.md) | The visibility info of the configuration template. | 
**LinkedAuditMessageChannelIds** | Pointer to **[]string** | The IDs of the audit message channels linked to the configuration template. | [optional] 
**MemberOncallScheduleIds** | Pointer to **[]string** | The IDs of the on-call schedules linked to the configuration template. | [optional] 
**BreakGlassUserIds** | Pointer to **[]string** | The IDs of the break glass users linked to the configuration template. | [optional] 
**RequireMfaToApprove** | **bool** | A bool representing whether or not to require MFA for reviewers to approve requests for this configuration template. | 
**RequireMfaToConnect** | **bool** | A bool representing whether or not to require MFA to connect to resources associated with this configuration template. | 
**Name** | **string** | The name of the configuration template. | 
**RequestConfigurations** | Pointer to [**[]RequestConfiguration**](RequestConfiguration.md) | The request configuration list of the configuration template. If not provided, the default request configuration will be used. | [optional] 
**RequestConfigurationList** | Pointer to [**CreateRequestConfigurationInfoList**](CreateRequestConfigurationInfoList.md) | The request configuration list of the configuration template. If not provided, the default request configuration will be used. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**TicketPropagation** | Pointer to [**TicketPropagationConfiguration**](TicketPropagationConfiguration.md) |  | [optional] 
**CustomRequestNotification** | Pointer to **string** | Custom request notification sent upon request approval for this configuration template. | [optional] 

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


### GetRequestConfigurations

`func (o *CreateConfigurationTemplateInfo) GetRequestConfigurations() []RequestConfiguration`

GetRequestConfigurations returns the RequestConfigurations field if non-nil, zero value otherwise.

### GetRequestConfigurationsOk

`func (o *CreateConfigurationTemplateInfo) GetRequestConfigurationsOk() (*[]RequestConfiguration, bool)`

GetRequestConfigurationsOk returns a tuple with the RequestConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfigurations

`func (o *CreateConfigurationTemplateInfo) SetRequestConfigurations(v []RequestConfiguration)`

SetRequestConfigurations sets RequestConfigurations field to given value.

### HasRequestConfigurations

`func (o *CreateConfigurationTemplateInfo) HasRequestConfigurations() bool`

HasRequestConfigurations returns a boolean if a field has been set.

### GetRequestConfigurationList

`func (o *CreateConfigurationTemplateInfo) GetRequestConfigurationList() CreateRequestConfigurationInfoList`

GetRequestConfigurationList returns the RequestConfigurationList field if non-nil, zero value otherwise.

### GetRequestConfigurationListOk

`func (o *CreateConfigurationTemplateInfo) GetRequestConfigurationListOk() (*CreateRequestConfigurationInfoList, bool)`

GetRequestConfigurationListOk returns a tuple with the RequestConfigurationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfigurationList

`func (o *CreateConfigurationTemplateInfo) SetRequestConfigurationList(v CreateRequestConfigurationInfoList)`

SetRequestConfigurationList sets RequestConfigurationList field to given value.

### HasRequestConfigurationList

`func (o *CreateConfigurationTemplateInfo) HasRequestConfigurationList() bool`

HasRequestConfigurationList returns a boolean if a field has been set.

### GetTicketPropagation

`func (o *CreateConfigurationTemplateInfo) GetTicketPropagation() TicketPropagationConfiguration`

GetTicketPropagation returns the TicketPropagation field if non-nil, zero value otherwise.

### GetTicketPropagationOk

`func (o *CreateConfigurationTemplateInfo) GetTicketPropagationOk() (*TicketPropagationConfiguration, bool)`

GetTicketPropagationOk returns a tuple with the TicketPropagation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketPropagation

`func (o *CreateConfigurationTemplateInfo) SetTicketPropagation(v TicketPropagationConfiguration)`

SetTicketPropagation sets TicketPropagation field to given value.

### HasTicketPropagation

`func (o *CreateConfigurationTemplateInfo) HasTicketPropagation() bool`

HasTicketPropagation returns a boolean if a field has been set.

### GetCustomRequestNotification

`func (o *CreateConfigurationTemplateInfo) GetCustomRequestNotification() string`

GetCustomRequestNotification returns the CustomRequestNotification field if non-nil, zero value otherwise.

### GetCustomRequestNotificationOk

`func (o *CreateConfigurationTemplateInfo) GetCustomRequestNotificationOk() (*string, bool)`

GetCustomRequestNotificationOk returns a tuple with the CustomRequestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRequestNotification

`func (o *CreateConfigurationTemplateInfo) SetCustomRequestNotification(v string)`

SetCustomRequestNotification sets CustomRequestNotification field to given value.

### HasCustomRequestNotification

`func (o *CreateConfigurationTemplateInfo) HasCustomRequestNotification() bool`

HasCustomRequestNotification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


