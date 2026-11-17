# RoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the role assignment. | 
**EntityId** | **string** | The ID of the entity (resource or group) being accessed. | 
**EntityType** | [**EntityTypeEnum**](EntityTypeEnum.md) |  | 
**PrincipalId** | **string** | The ID of the principal being granted access. | 
**PrincipalType** | [**EntityTypeEnum**](EntityTypeEnum.md) |  | 
**AccessLevel** | Pointer to [**ResourceAccessLevel**](ResourceAccessLevel.md) | The level of access granted. Null when the entity does not define access levels. | [optional] 
**Expiration** | Pointer to **time.Time** | When this access expires. Null when the grant is indefinite. | [optional] 
**GrantedAt** | Pointer to **time.Time** | When this access was first granted. | [optional] 

## Methods

### NewRoleAssignment

`func NewRoleAssignment(id string, entityId string, entityType EntityTypeEnum, principalId string, principalType EntityTypeEnum, ) *RoleAssignment`

NewRoleAssignment instantiates a new RoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAssignmentWithDefaults

`func NewRoleAssignmentWithDefaults() *RoleAssignment`

NewRoleAssignmentWithDefaults instantiates a new RoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoleAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleAssignment) SetId(v string)`

SetId sets Id field to given value.


### GetEntityId

`func (o *RoleAssignment) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *RoleAssignment) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *RoleAssignment) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *RoleAssignment) GetEntityType() EntityTypeEnum`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *RoleAssignment) GetEntityTypeOk() (*EntityTypeEnum, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *RoleAssignment) SetEntityType(v EntityTypeEnum)`

SetEntityType sets EntityType field to given value.


### GetPrincipalId

`func (o *RoleAssignment) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *RoleAssignment) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *RoleAssignment) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetPrincipalType

`func (o *RoleAssignment) GetPrincipalType() EntityTypeEnum`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *RoleAssignment) GetPrincipalTypeOk() (*EntityTypeEnum, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *RoleAssignment) SetPrincipalType(v EntityTypeEnum)`

SetPrincipalType sets PrincipalType field to given value.


### GetAccessLevel

`func (o *RoleAssignment) GetAccessLevel() ResourceAccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *RoleAssignment) GetAccessLevelOk() (*ResourceAccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *RoleAssignment) SetAccessLevel(v ResourceAccessLevel)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *RoleAssignment) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetExpiration

`func (o *RoleAssignment) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *RoleAssignment) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *RoleAssignment) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *RoleAssignment) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetGrantedAt

`func (o *RoleAssignment) GetGrantedAt() time.Time`

GetGrantedAt returns the GrantedAt field if non-nil, zero value otherwise.

### GetGrantedAtOk

`func (o *RoleAssignment) GetGrantedAtOk() (*time.Time, bool)`

GetGrantedAtOk returns a tuple with the GrantedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedAt

`func (o *RoleAssignment) SetGrantedAt(v time.Time)`

SetGrantedAt sets GrantedAt field to given value.

### HasGrantedAt

`func (o *RoleAssignment) HasGrantedAt() bool`

HasGrantedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


