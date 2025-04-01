# UpdateResourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The ID of the resource. | 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**Description** | Pointer to **string** | A description of the resource. | [optional] 
**AdminOwnerId** | Pointer to **string** | The ID of the owner of the resource. | [optional] 
**MaxDuration** | Pointer to **int32** | The maximum duration for which the resource can be requested (in minutes). Use -1 to set to indefinite. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**RecommendedDuration** | Pointer to **int32** | The recommended duration for which the resource should be requested (in minutes). Will be the default value in a request. Use -1 to set to indefinite and 0 to unset. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**RequireManagerApproval** | Pointer to **bool** | A bool representing whether or not access requests to the resource require manager approval. | [optional] 
**RequireSupportTicket** | Pointer to **bool** | A bool representing whether or not access requests to the resource require an access ticket. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**FolderId** | Pointer to **string** | The ID of the folder that the resource is located in. | [optional] 
**RequireMfaToApprove** | Pointer to **bool** | A bool representing whether or not to require MFA for reviewers to approve requests for this resource. | [optional] 
**RequireMfaToRequest** | Pointer to **bool** | A bool representing whether or not to require MFA for requesting access to this resource. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**RequireMfaToConnect** | Pointer to **bool** | A bool representing whether or not to require MFA to connect to this resource. | [optional] 
**AutoApproval** | Pointer to **bool** | A bool representing whether or not to automatically approve requests to this resource. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**TicketPropagation** | Pointer to [**TicketPropagationConfiguration**](TicketPropagationConfiguration.md) |  | [optional] 
**CustomRequestNotification** | Pointer to **string** | Custom request notification sent upon request approval. | [optional] 
**RiskSensitivityOverride** | Pointer to [**RiskSensitivityEnum**](RiskSensitivityEnum.md) |  | [optional] 
**ConfigurationTemplateId** | Pointer to **string** | The ID of the associated configuration template. | [optional] 
**RequestTemplateId** | Pointer to **string** | The ID of the associated request template. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**IsRequestable** | Pointer to **bool** | A bool representing whether or not to allow access requests to this resource. Deprecated in favor of &#x60;request_configurations&#x60;. | [optional] 
**RequestConfigurations** | Pointer to [**[]RequestConfiguration**](RequestConfiguration.md) | A list of configurations for requests to this resource. If not provided, the default request configuration will be used. | [optional] 
**RequestConfigurationList** | Pointer to [**CreateRequestConfigurationInfoList**](CreateRequestConfigurationInfoList.md) |  | [optional] 

## Methods

### NewUpdateResourceInfo

`func NewUpdateResourceInfo(resourceId string, ) *UpdateResourceInfo`

NewUpdateResourceInfo instantiates a new UpdateResourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceInfoWithDefaults

`func NewUpdateResourceInfoWithDefaults() *UpdateResourceInfo`

NewUpdateResourceInfoWithDefaults instantiates a new UpdateResourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *UpdateResourceInfo) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *UpdateResourceInfo) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *UpdateResourceInfo) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetName

`func (o *UpdateResourceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateResourceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateResourceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateResourceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateResourceInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateResourceInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateResourceInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateResourceInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdminOwnerId

`func (o *UpdateResourceInfo) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *UpdateResourceInfo) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *UpdateResourceInfo) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.

### HasAdminOwnerId

`func (o *UpdateResourceInfo) HasAdminOwnerId() bool`

HasAdminOwnerId returns a boolean if a field has been set.

### GetMaxDuration

`func (o *UpdateResourceInfo) GetMaxDuration() int32`

GetMaxDuration returns the MaxDuration field if non-nil, zero value otherwise.

### GetMaxDurationOk

`func (o *UpdateResourceInfo) GetMaxDurationOk() (*int32, bool)`

GetMaxDurationOk returns a tuple with the MaxDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDuration

`func (o *UpdateResourceInfo) SetMaxDuration(v int32)`

SetMaxDuration sets MaxDuration field to given value.

### HasMaxDuration

`func (o *UpdateResourceInfo) HasMaxDuration() bool`

HasMaxDuration returns a boolean if a field has been set.

### GetRecommendedDuration

`func (o *UpdateResourceInfo) GetRecommendedDuration() int32`

GetRecommendedDuration returns the RecommendedDuration field if non-nil, zero value otherwise.

### GetRecommendedDurationOk

`func (o *UpdateResourceInfo) GetRecommendedDurationOk() (*int32, bool)`

GetRecommendedDurationOk returns a tuple with the RecommendedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedDuration

`func (o *UpdateResourceInfo) SetRecommendedDuration(v int32)`

SetRecommendedDuration sets RecommendedDuration field to given value.

### HasRecommendedDuration

`func (o *UpdateResourceInfo) HasRecommendedDuration() bool`

HasRecommendedDuration returns a boolean if a field has been set.

### GetRequireManagerApproval

`func (o *UpdateResourceInfo) GetRequireManagerApproval() bool`

GetRequireManagerApproval returns the RequireManagerApproval field if non-nil, zero value otherwise.

### GetRequireManagerApprovalOk

`func (o *UpdateResourceInfo) GetRequireManagerApprovalOk() (*bool, bool)`

GetRequireManagerApprovalOk returns a tuple with the RequireManagerApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireManagerApproval

`func (o *UpdateResourceInfo) SetRequireManagerApproval(v bool)`

SetRequireManagerApproval sets RequireManagerApproval field to given value.

### HasRequireManagerApproval

`func (o *UpdateResourceInfo) HasRequireManagerApproval() bool`

HasRequireManagerApproval returns a boolean if a field has been set.

### GetRequireSupportTicket

`func (o *UpdateResourceInfo) GetRequireSupportTicket() bool`

GetRequireSupportTicket returns the RequireSupportTicket field if non-nil, zero value otherwise.

### GetRequireSupportTicketOk

`func (o *UpdateResourceInfo) GetRequireSupportTicketOk() (*bool, bool)`

GetRequireSupportTicketOk returns a tuple with the RequireSupportTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSupportTicket

`func (o *UpdateResourceInfo) SetRequireSupportTicket(v bool)`

SetRequireSupportTicket sets RequireSupportTicket field to given value.

### HasRequireSupportTicket

`func (o *UpdateResourceInfo) HasRequireSupportTicket() bool`

HasRequireSupportTicket returns a boolean if a field has been set.

### GetFolderId

`func (o *UpdateResourceInfo) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *UpdateResourceInfo) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *UpdateResourceInfo) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *UpdateResourceInfo) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetRequireMfaToApprove

`func (o *UpdateResourceInfo) GetRequireMfaToApprove() bool`

GetRequireMfaToApprove returns the RequireMfaToApprove field if non-nil, zero value otherwise.

### GetRequireMfaToApproveOk

`func (o *UpdateResourceInfo) GetRequireMfaToApproveOk() (*bool, bool)`

GetRequireMfaToApproveOk returns a tuple with the RequireMfaToApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToApprove

`func (o *UpdateResourceInfo) SetRequireMfaToApprove(v bool)`

SetRequireMfaToApprove sets RequireMfaToApprove field to given value.

### HasRequireMfaToApprove

`func (o *UpdateResourceInfo) HasRequireMfaToApprove() bool`

HasRequireMfaToApprove returns a boolean if a field has been set.

### GetRequireMfaToRequest

`func (o *UpdateResourceInfo) GetRequireMfaToRequest() bool`

GetRequireMfaToRequest returns the RequireMfaToRequest field if non-nil, zero value otherwise.

### GetRequireMfaToRequestOk

`func (o *UpdateResourceInfo) GetRequireMfaToRequestOk() (*bool, bool)`

GetRequireMfaToRequestOk returns a tuple with the RequireMfaToRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToRequest

`func (o *UpdateResourceInfo) SetRequireMfaToRequest(v bool)`

SetRequireMfaToRequest sets RequireMfaToRequest field to given value.

### HasRequireMfaToRequest

`func (o *UpdateResourceInfo) HasRequireMfaToRequest() bool`

HasRequireMfaToRequest returns a boolean if a field has been set.

### GetRequireMfaToConnect

`func (o *UpdateResourceInfo) GetRequireMfaToConnect() bool`

GetRequireMfaToConnect returns the RequireMfaToConnect field if non-nil, zero value otherwise.

### GetRequireMfaToConnectOk

`func (o *UpdateResourceInfo) GetRequireMfaToConnectOk() (*bool, bool)`

GetRequireMfaToConnectOk returns a tuple with the RequireMfaToConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToConnect

`func (o *UpdateResourceInfo) SetRequireMfaToConnect(v bool)`

SetRequireMfaToConnect sets RequireMfaToConnect field to given value.

### HasRequireMfaToConnect

`func (o *UpdateResourceInfo) HasRequireMfaToConnect() bool`

HasRequireMfaToConnect returns a boolean if a field has been set.

### GetAutoApproval

`func (o *UpdateResourceInfo) GetAutoApproval() bool`

GetAutoApproval returns the AutoApproval field if non-nil, zero value otherwise.

### GetAutoApprovalOk

`func (o *UpdateResourceInfo) GetAutoApprovalOk() (*bool, bool)`

GetAutoApprovalOk returns a tuple with the AutoApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproval

`func (o *UpdateResourceInfo) SetAutoApproval(v bool)`

SetAutoApproval sets AutoApproval field to given value.

### HasAutoApproval

`func (o *UpdateResourceInfo) HasAutoApproval() bool`

HasAutoApproval returns a boolean if a field has been set.

### GetTicketPropagation

`func (o *UpdateResourceInfo) GetTicketPropagation() TicketPropagationConfiguration`

GetTicketPropagation returns the TicketPropagation field if non-nil, zero value otherwise.

### GetTicketPropagationOk

`func (o *UpdateResourceInfo) GetTicketPropagationOk() (*TicketPropagationConfiguration, bool)`

GetTicketPropagationOk returns a tuple with the TicketPropagation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketPropagation

`func (o *UpdateResourceInfo) SetTicketPropagation(v TicketPropagationConfiguration)`

SetTicketPropagation sets TicketPropagation field to given value.

### HasTicketPropagation

`func (o *UpdateResourceInfo) HasTicketPropagation() bool`

HasTicketPropagation returns a boolean if a field has been set.

### GetCustomRequestNotification

`func (o *UpdateResourceInfo) GetCustomRequestNotification() string`

GetCustomRequestNotification returns the CustomRequestNotification field if non-nil, zero value otherwise.

### GetCustomRequestNotificationOk

`func (o *UpdateResourceInfo) GetCustomRequestNotificationOk() (*string, bool)`

GetCustomRequestNotificationOk returns a tuple with the CustomRequestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRequestNotification

`func (o *UpdateResourceInfo) SetCustomRequestNotification(v string)`

SetCustomRequestNotification sets CustomRequestNotification field to given value.

### HasCustomRequestNotification

`func (o *UpdateResourceInfo) HasCustomRequestNotification() bool`

HasCustomRequestNotification returns a boolean if a field has been set.

### GetRiskSensitivityOverride

`func (o *UpdateResourceInfo) GetRiskSensitivityOverride() RiskSensitivityEnum`

GetRiskSensitivityOverride returns the RiskSensitivityOverride field if non-nil, zero value otherwise.

### GetRiskSensitivityOverrideOk

`func (o *UpdateResourceInfo) GetRiskSensitivityOverrideOk() (*RiskSensitivityEnum, bool)`

GetRiskSensitivityOverrideOk returns a tuple with the RiskSensitivityOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskSensitivityOverride

`func (o *UpdateResourceInfo) SetRiskSensitivityOverride(v RiskSensitivityEnum)`

SetRiskSensitivityOverride sets RiskSensitivityOverride field to given value.

### HasRiskSensitivityOverride

`func (o *UpdateResourceInfo) HasRiskSensitivityOverride() bool`

HasRiskSensitivityOverride returns a boolean if a field has been set.

### GetConfigurationTemplateId

`func (o *UpdateResourceInfo) GetConfigurationTemplateId() string`

GetConfigurationTemplateId returns the ConfigurationTemplateId field if non-nil, zero value otherwise.

### GetConfigurationTemplateIdOk

`func (o *UpdateResourceInfo) GetConfigurationTemplateIdOk() (*string, bool)`

GetConfigurationTemplateIdOk returns a tuple with the ConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTemplateId

`func (o *UpdateResourceInfo) SetConfigurationTemplateId(v string)`

SetConfigurationTemplateId sets ConfigurationTemplateId field to given value.

### HasConfigurationTemplateId

`func (o *UpdateResourceInfo) HasConfigurationTemplateId() bool`

HasConfigurationTemplateId returns a boolean if a field has been set.

### GetRequestTemplateId

`func (o *UpdateResourceInfo) GetRequestTemplateId() string`

GetRequestTemplateId returns the RequestTemplateId field if non-nil, zero value otherwise.

### GetRequestTemplateIdOk

`func (o *UpdateResourceInfo) GetRequestTemplateIdOk() (*string, bool)`

GetRequestTemplateIdOk returns a tuple with the RequestTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTemplateId

`func (o *UpdateResourceInfo) SetRequestTemplateId(v string)`

SetRequestTemplateId sets RequestTemplateId field to given value.

### HasRequestTemplateId

`func (o *UpdateResourceInfo) HasRequestTemplateId() bool`

HasRequestTemplateId returns a boolean if a field has been set.

### GetIsRequestable

`func (o *UpdateResourceInfo) GetIsRequestable() bool`

GetIsRequestable returns the IsRequestable field if non-nil, zero value otherwise.

### GetIsRequestableOk

`func (o *UpdateResourceInfo) GetIsRequestableOk() (*bool, bool)`

GetIsRequestableOk returns a tuple with the IsRequestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequestable

`func (o *UpdateResourceInfo) SetIsRequestable(v bool)`

SetIsRequestable sets IsRequestable field to given value.

### HasIsRequestable

`func (o *UpdateResourceInfo) HasIsRequestable() bool`

HasIsRequestable returns a boolean if a field has been set.

### GetRequestConfigurations

`func (o *UpdateResourceInfo) GetRequestConfigurations() []RequestConfiguration`

GetRequestConfigurations returns the RequestConfigurations field if non-nil, zero value otherwise.

### GetRequestConfigurationsOk

`func (o *UpdateResourceInfo) GetRequestConfigurationsOk() (*[]RequestConfiguration, bool)`

GetRequestConfigurationsOk returns a tuple with the RequestConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfigurations

`func (o *UpdateResourceInfo) SetRequestConfigurations(v []RequestConfiguration)`

SetRequestConfigurations sets RequestConfigurations field to given value.

### HasRequestConfigurations

`func (o *UpdateResourceInfo) HasRequestConfigurations() bool`

HasRequestConfigurations returns a boolean if a field has been set.

### GetRequestConfigurationList

`func (o *UpdateResourceInfo) GetRequestConfigurationList() CreateRequestConfigurationInfoList`

GetRequestConfigurationList returns the RequestConfigurationList field if non-nil, zero value otherwise.

### GetRequestConfigurationListOk

`func (o *UpdateResourceInfo) GetRequestConfigurationListOk() (*CreateRequestConfigurationInfoList, bool)`

GetRequestConfigurationListOk returns a tuple with the RequestConfigurationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfigurationList

`func (o *UpdateResourceInfo) SetRequestConfigurationList(v CreateRequestConfigurationInfoList)`

SetRequestConfigurationList sets RequestConfigurationList field to given value.

### HasRequestConfigurationList

`func (o *UpdateResourceInfo) HasRequestConfigurationList() bool`

HasRequestConfigurationList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


