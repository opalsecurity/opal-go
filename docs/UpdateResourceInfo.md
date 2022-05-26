# UpdateResourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The ID of the resource. | 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**Description** | Pointer to **string** | A description of the resource. | [optional] 
**OwnerTeamId** | Pointer to **string** | The ID of the owning team of the resource. | [optional] 
**Visibility** | Pointer to [**VisibilityEnum**](VisibilityEnum.md) |  | [optional] 
**MaxDuration** | Pointer to **int32** | The maximum duration access to the resource can be requested for (in minutes). Use 0 to set to indefinite. | [optional] 
**RequireManagerApproval** | Pointer to **bool** | A bool representing whether or not access requests to the resource require manager approval. | [optional] 
**RequireSupportTicket** | Pointer to **bool** | A bool representing whether or not access requests to the resource require a support ticket. | [optional] 
**FolderId** | Pointer to **string** | The ID of the folder that the resource is located in. | [optional] 

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

### GetOwnerTeamId

`func (o *UpdateResourceInfo) GetOwnerTeamId() string`

GetOwnerTeamId returns the OwnerTeamId field if non-nil, zero value otherwise.

### GetOwnerTeamIdOk

`func (o *UpdateResourceInfo) GetOwnerTeamIdOk() (*string, bool)`

GetOwnerTeamIdOk returns a tuple with the OwnerTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerTeamId

`func (o *UpdateResourceInfo) SetOwnerTeamId(v string)`

SetOwnerTeamId sets OwnerTeamId field to given value.

### HasOwnerTeamId

`func (o *UpdateResourceInfo) HasOwnerTeamId() bool`

HasOwnerTeamId returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateResourceInfo) GetVisibility() VisibilityEnum`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateResourceInfo) GetVisibilityOk() (*VisibilityEnum, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateResourceInfo) SetVisibility(v VisibilityEnum)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateResourceInfo) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


