# UserAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | [**UserAccessEntityTypeEnum**](UserAccessEntityTypeEnum.md) |  | 
**EntityId** | **string** | The ID of the resource or group. | 
**EntityName** | **string** | The name of the resource or group. | 
**Description** | **string** | The description of the resource or group. | 
**AccessLevel** | [**UserAccessLevel**](UserAccessLevel.md) |  | 
**ExpirationDate** | Pointer to **time.Time** | The day and time the user&#39;s access will expire. If the user has more than one way to access the entity, this is the latest expiration. Null if the access does not expire. | [optional] 
**HasDirectAccess** | **bool** | The user has direct access to this entity (vs. indirectly, like through a group). | 
**NumAccessPaths** | **int32** | The number of ways in which the user has access to this entity (directly and indirectly). | 
**PropagationStatus** | Pointer to [**PropagationStatus**](PropagationStatus.md) |  | [optional] 

## Methods

### NewUserAccess

`func NewUserAccess(entityType UserAccessEntityTypeEnum, entityId string, entityName string, description string, accessLevel UserAccessLevel, hasDirectAccess bool, numAccessPaths int32, ) *UserAccess`

NewUserAccess instantiates a new UserAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccessWithDefaults

`func NewUserAccessWithDefaults() *UserAccess`

NewUserAccessWithDefaults instantiates a new UserAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *UserAccess) GetEntityType() UserAccessEntityTypeEnum`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *UserAccess) GetEntityTypeOk() (*UserAccessEntityTypeEnum, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *UserAccess) SetEntityType(v UserAccessEntityTypeEnum)`

SetEntityType sets EntityType field to given value.


### GetEntityId

`func (o *UserAccess) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *UserAccess) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *UserAccess) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityName

`func (o *UserAccess) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *UserAccess) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *UserAccess) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.


### GetDescription

`func (o *UserAccess) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserAccess) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserAccess) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAccessLevel

`func (o *UserAccess) GetAccessLevel() UserAccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *UserAccess) GetAccessLevelOk() (*UserAccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *UserAccess) SetAccessLevel(v UserAccessLevel)`

SetAccessLevel sets AccessLevel field to given value.


### GetExpirationDate

`func (o *UserAccess) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *UserAccess) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *UserAccess) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *UserAccess) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetHasDirectAccess

`func (o *UserAccess) GetHasDirectAccess() bool`

GetHasDirectAccess returns the HasDirectAccess field if non-nil, zero value otherwise.

### GetHasDirectAccessOk

`func (o *UserAccess) GetHasDirectAccessOk() (*bool, bool)`

GetHasDirectAccessOk returns a tuple with the HasDirectAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDirectAccess

`func (o *UserAccess) SetHasDirectAccess(v bool)`

SetHasDirectAccess sets HasDirectAccess field to given value.


### GetNumAccessPaths

`func (o *UserAccess) GetNumAccessPaths() int32`

GetNumAccessPaths returns the NumAccessPaths field if non-nil, zero value otherwise.

### GetNumAccessPathsOk

`func (o *UserAccess) GetNumAccessPathsOk() (*int32, bool)`

GetNumAccessPathsOk returns a tuple with the NumAccessPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAccessPaths

`func (o *UserAccess) SetNumAccessPaths(v int32)`

SetNumAccessPaths sets NumAccessPaths field to given value.


### GetPropagationStatus

`func (o *UserAccess) GetPropagationStatus() PropagationStatus`

GetPropagationStatus returns the PropagationStatus field if non-nil, zero value otherwise.

### GetPropagationStatusOk

`func (o *UserAccess) GetPropagationStatusOk() (*PropagationStatus, bool)`

GetPropagationStatusOk returns a tuple with the PropagationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationStatus

`func (o *UserAccess) SetPropagationStatus(v PropagationStatus)`

SetPropagationStatus sets PropagationStatus field to given value.

### HasPropagationStatus

`func (o *UserAccess) HasPropagationStatus() bool`

HasPropagationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


