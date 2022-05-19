# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The ID of the resource. | 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**Description** | Pointer to **string** | A description of the resource. | [optional] 
**OwnerTeamId** | Pointer to **string** | The ID of the owning team of the resource. | [optional] 
**Visibility** | Pointer to [**VisibilityEnum**](VisibilityEnum.md) |  | [optional] 
**MaxDuration** | Pointer to **int32** | The maximum duration access to the resource can be requested for (in minutes). | [optional] 
**RequireManagerApproval** | Pointer to **bool** | A bool representing whether or not access requests to the resource require manager approval. | [optional] 
**RequireSupportTicket** | Pointer to **bool** | A bool representing whether or not access requests to the resource require a support ticket. | [optional] 

## Methods

### NewResource

`func NewResource(resourceId string, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *Resource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Resource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Resource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetName

`func (o *Resource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Resource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Resource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Resource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Resource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Resource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Resource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Resource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwnerTeamId

`func (o *Resource) GetOwnerTeamId() string`

GetOwnerTeamId returns the OwnerTeamId field if non-nil, zero value otherwise.

### GetOwnerTeamIdOk

`func (o *Resource) GetOwnerTeamIdOk() (*string, bool)`

GetOwnerTeamIdOk returns a tuple with the OwnerTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerTeamId

`func (o *Resource) SetOwnerTeamId(v string)`

SetOwnerTeamId sets OwnerTeamId field to given value.

### HasOwnerTeamId

`func (o *Resource) HasOwnerTeamId() bool`

HasOwnerTeamId returns a boolean if a field has been set.

### GetVisibility

`func (o *Resource) GetVisibility() VisibilityEnum`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Resource) GetVisibilityOk() (*VisibilityEnum, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Resource) SetVisibility(v VisibilityEnum)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Resource) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetMaxDuration

`func (o *Resource) GetMaxDuration() int32`

GetMaxDuration returns the MaxDuration field if non-nil, zero value otherwise.

### GetMaxDurationOk

`func (o *Resource) GetMaxDurationOk() (*int32, bool)`

GetMaxDurationOk returns a tuple with the MaxDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDuration

`func (o *Resource) SetMaxDuration(v int32)`

SetMaxDuration sets MaxDuration field to given value.

### HasMaxDuration

`func (o *Resource) HasMaxDuration() bool`

HasMaxDuration returns a boolean if a field has been set.

### GetRequireManagerApproval

`func (o *Resource) GetRequireManagerApproval() bool`

GetRequireManagerApproval returns the RequireManagerApproval field if non-nil, zero value otherwise.

### GetRequireManagerApprovalOk

`func (o *Resource) GetRequireManagerApprovalOk() (*bool, bool)`

GetRequireManagerApprovalOk returns a tuple with the RequireManagerApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireManagerApproval

`func (o *Resource) SetRequireManagerApproval(v bool)`

SetRequireManagerApproval sets RequireManagerApproval field to given value.

### HasRequireManagerApproval

`func (o *Resource) HasRequireManagerApproval() bool`

HasRequireManagerApproval returns a boolean if a field has been set.

### GetRequireSupportTicket

`func (o *Resource) GetRequireSupportTicket() bool`

GetRequireSupportTicket returns the RequireSupportTicket field if non-nil, zero value otherwise.

### GetRequireSupportTicketOk

`func (o *Resource) GetRequireSupportTicketOk() (*bool, bool)`

GetRequireSupportTicketOk returns a tuple with the RequireSupportTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSupportTicket

`func (o *Resource) SetRequireSupportTicket(v bool)`

SetRequireSupportTicket sets RequireSupportTicket field to given value.

### HasRequireSupportTicket

`func (o *Resource) HasRequireSupportTicket() bool`

HasRequireSupportTicket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


