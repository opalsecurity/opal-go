# ScopedRolePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetIds** | Pointer to **[]string** | The IDs of the entities that this permission applies to. If empty of missing, the permission will have untargeted scope. | [optional] 
**TargetType** | [**RolePermissionTargetTypeEnum**](RolePermissionTargetTypeEnum.md) |  | 
**PermissionName** | [**RolePermissionNameEnum**](RolePermissionNameEnum.md) |  | 
**AllowAll** | **bool** |  | 

## Methods

### NewScopedRolePermission

`func NewScopedRolePermission(targetType RolePermissionTargetTypeEnum, permissionName RolePermissionNameEnum, allowAll bool, ) *ScopedRolePermission`

NewScopedRolePermission instantiates a new ScopedRolePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopedRolePermissionWithDefaults

`func NewScopedRolePermissionWithDefaults() *ScopedRolePermission`

NewScopedRolePermissionWithDefaults instantiates a new ScopedRolePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetIds

`func (o *ScopedRolePermission) GetTargetIds() []string`

GetTargetIds returns the TargetIds field if non-nil, zero value otherwise.

### GetTargetIdsOk

`func (o *ScopedRolePermission) GetTargetIdsOk() (*[]string, bool)`

GetTargetIdsOk returns a tuple with the TargetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIds

`func (o *ScopedRolePermission) SetTargetIds(v []string)`

SetTargetIds sets TargetIds field to given value.

### HasTargetIds

`func (o *ScopedRolePermission) HasTargetIds() bool`

HasTargetIds returns a boolean if a field has been set.

### GetTargetType

`func (o *ScopedRolePermission) GetTargetType() RolePermissionTargetTypeEnum`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ScopedRolePermission) GetTargetTypeOk() (*RolePermissionTargetTypeEnum, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ScopedRolePermission) SetTargetType(v RolePermissionTargetTypeEnum)`

SetTargetType sets TargetType field to given value.


### GetPermissionName

`func (o *ScopedRolePermission) GetPermissionName() RolePermissionNameEnum`

GetPermissionName returns the PermissionName field if non-nil, zero value otherwise.

### GetPermissionNameOk

`func (o *ScopedRolePermission) GetPermissionNameOk() (*RolePermissionNameEnum, bool)`

GetPermissionNameOk returns a tuple with the PermissionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionName

`func (o *ScopedRolePermission) SetPermissionName(v RolePermissionNameEnum)`

SetPermissionName sets PermissionName field to given value.


### GetAllowAll

`func (o *ScopedRolePermission) GetAllowAll() bool`

GetAllowAll returns the AllowAll field if non-nil, zero value otherwise.

### GetAllowAllOk

`func (o *ScopedRolePermission) GetAllowAllOk() (*bool, bool)`

GetAllowAllOk returns a tuple with the AllowAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAll

`func (o *ScopedRolePermission) SetAllowAll(v bool)`

SetAllowAll sets AllowAll field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


