# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupIds** | Pointer to **[]string** | The list of group IDs to match. | [optional] 
**RoleRemoteIds** | Pointer to **[]string** | The list of role remote IDs to match. | [optional] 

## Methods

### NewCondition

`func NewCondition() *Condition`

NewCondition instantiates a new Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionWithDefaults

`func NewConditionWithDefaults() *Condition`

NewConditionWithDefaults instantiates a new Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupIds

`func (o *Condition) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *Condition) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *Condition) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *Condition) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetRoleRemoteIds

`func (o *Condition) GetRoleRemoteIds() []string`

GetRoleRemoteIds returns the RoleRemoteIds field if non-nil, zero value otherwise.

### GetRoleRemoteIdsOk

`func (o *Condition) GetRoleRemoteIdsOk() (*[]string, bool)`

GetRoleRemoteIdsOk returns a tuple with the RoleRemoteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleRemoteIds

`func (o *Condition) SetRoleRemoteIds(v []string)`

SetRoleRemoteIds sets RoleRemoteIds field to given value.

### HasRoleRemoteIds

`func (o *Condition) HasRoleRemoteIds() bool`

HasRoleRemoteIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


