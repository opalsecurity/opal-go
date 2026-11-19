# AddRoleAssignmentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The ID of the entity (resource or group) being accessed. | 
**EntityType** | [**EntityTypeEnum**](EntityTypeEnum.md) | The type of the entity being accessed. Must be &#x60;RESOURCE&#x60; or &#x60;GROUP&#x60;. | 
**PrincipalId** | **string** | The ID of the principal (user, group, or resource) being granted access. | 
**PrincipalType** | [**EntityTypeEnum**](EntityTypeEnum.md) | The type of the principal being granted access. Must be &#x60;RESOURCE&#x60;, &#x60;GROUP&#x60;, or &#x60;USER&#x60;. | 
**AccessLevelRemoteId** | Pointer to **string** | Remote ID of the access level to grant. Omit for entities that only have a single default access level. | [optional] 
**DurationMinutes** | Pointer to **int32** | Duration of the grant in minutes. Omit or send &#x60;0&#x60; for an indefinite grant. | [optional] 

## Methods

### NewAddRoleAssignmentInput

`func NewAddRoleAssignmentInput(entityId string, entityType EntityTypeEnum, principalId string, principalType EntityTypeEnum, ) *AddRoleAssignmentInput`

NewAddRoleAssignmentInput instantiates a new AddRoleAssignmentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRoleAssignmentInputWithDefaults

`func NewAddRoleAssignmentInputWithDefaults() *AddRoleAssignmentInput`

NewAddRoleAssignmentInputWithDefaults instantiates a new AddRoleAssignmentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *AddRoleAssignmentInput) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AddRoleAssignmentInput) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AddRoleAssignmentInput) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *AddRoleAssignmentInput) GetEntityType() EntityTypeEnum`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *AddRoleAssignmentInput) GetEntityTypeOk() (*EntityTypeEnum, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *AddRoleAssignmentInput) SetEntityType(v EntityTypeEnum)`

SetEntityType sets EntityType field to given value.


### GetPrincipalId

`func (o *AddRoleAssignmentInput) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *AddRoleAssignmentInput) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *AddRoleAssignmentInput) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetPrincipalType

`func (o *AddRoleAssignmentInput) GetPrincipalType() EntityTypeEnum`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *AddRoleAssignmentInput) GetPrincipalTypeOk() (*EntityTypeEnum, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *AddRoleAssignmentInput) SetPrincipalType(v EntityTypeEnum)`

SetPrincipalType sets PrincipalType field to given value.


### GetAccessLevelRemoteId

`func (o *AddRoleAssignmentInput) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *AddRoleAssignmentInput) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *AddRoleAssignmentInput) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *AddRoleAssignmentInput) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.

### GetDurationMinutes

`func (o *AddRoleAssignmentInput) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *AddRoleAssignmentInput) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *AddRoleAssignmentInput) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.

### HasDurationMinutes

`func (o *AddRoleAssignmentInput) HasDurationMinutes() bool`

HasDurationMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


