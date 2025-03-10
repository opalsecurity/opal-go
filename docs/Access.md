# Access

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalId** | **string** | The ID of the principal with access. | 
**PrincipalType** | [**EntityTypeEnum**](EntityTypeEnum.md) |  | 
**EntityId** | **string** | The ID of the entity being accessed. | 
**EntityType** | [**EntityTypeEnum**](EntityTypeEnum.md) |  | 
**AccessLevel** | Pointer to [**ResourceAccessLevel**](ResourceAccessLevel.md) |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** | The day and time the principal&#39;s access will expire. | [optional] 
**HasDirectAccess** | **bool** | The principal has direct access to this entity (vs. inherited access). | 
**NumAccessPaths** | **int32** | The number of ways in which the principal has access to this entity (directly and inherited). | 

## Methods

### NewAccess

`func NewAccess(principalId string, principalType EntityTypeEnum, entityId string, entityType EntityTypeEnum, hasDirectAccess bool, numAccessPaths int32, ) *Access`

NewAccess instantiates a new Access object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessWithDefaults

`func NewAccessWithDefaults() *Access`

NewAccessWithDefaults instantiates a new Access object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalId

`func (o *Access) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *Access) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *Access) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetPrincipalType

`func (o *Access) GetPrincipalType() EntityTypeEnum`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *Access) GetPrincipalTypeOk() (*EntityTypeEnum, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *Access) SetPrincipalType(v EntityTypeEnum)`

SetPrincipalType sets PrincipalType field to given value.


### GetEntityId

`func (o *Access) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Access) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Access) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *Access) GetEntityType() EntityTypeEnum`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Access) GetEntityTypeOk() (*EntityTypeEnum, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Access) SetEntityType(v EntityTypeEnum)`

SetEntityType sets EntityType field to given value.


### GetAccessLevel

`func (o *Access) GetAccessLevel() ResourceAccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *Access) GetAccessLevelOk() (*ResourceAccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *Access) SetAccessLevel(v ResourceAccessLevel)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *Access) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetExpirationDate

`func (o *Access) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *Access) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *Access) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *Access) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetHasDirectAccess

`func (o *Access) GetHasDirectAccess() bool`

GetHasDirectAccess returns the HasDirectAccess field if non-nil, zero value otherwise.

### GetHasDirectAccessOk

`func (o *Access) GetHasDirectAccessOk() (*bool, bool)`

GetHasDirectAccessOk returns a tuple with the HasDirectAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDirectAccess

`func (o *Access) SetHasDirectAccess(v bool)`

SetHasDirectAccess sets HasDirectAccess field to given value.


### GetNumAccessPaths

`func (o *Access) GetNumAccessPaths() int32`

GetNumAccessPaths returns the NumAccessPaths field if non-nil, zero value otherwise.

### GetNumAccessPathsOk

`func (o *Access) GetNumAccessPathsOk() (*int32, bool)`

GetNumAccessPathsOk returns a tuple with the NumAccessPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAccessPaths

`func (o *Access) SetNumAccessPaths(v int32)`

SetNumAccessPaths sets NumAccessPaths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


