# GroupContainingGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainingGroupId** | **string** | The groupID of the containing group. | 
**DurationMinutes** | Pointer to **int32** | The updated duration for which the group can be accessed (in minutes). Use 0 for indefinite, or a negative value to revoke access. | [optional] 
**AccessLevelRemoteId** | Pointer to **string** | The updated remote ID of the access level granted to this group. | [optional] 

## Methods

### NewGroupContainingGroup

`func NewGroupContainingGroup(containingGroupId string, ) *GroupContainingGroup`

NewGroupContainingGroup instantiates a new GroupContainingGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupContainingGroupWithDefaults

`func NewGroupContainingGroupWithDefaults() *GroupContainingGroup`

NewGroupContainingGroupWithDefaults instantiates a new GroupContainingGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainingGroupId

`func (o *GroupContainingGroup) GetContainingGroupId() string`

GetContainingGroupId returns the ContainingGroupId field if non-nil, zero value otherwise.

### GetContainingGroupIdOk

`func (o *GroupContainingGroup) GetContainingGroupIdOk() (*string, bool)`

GetContainingGroupIdOk returns a tuple with the ContainingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainingGroupId

`func (o *GroupContainingGroup) SetContainingGroupId(v string)`

SetContainingGroupId sets ContainingGroupId field to given value.


### GetDurationMinutes

`func (o *GroupContainingGroup) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *GroupContainingGroup) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *GroupContainingGroup) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.

### HasDurationMinutes

`func (o *GroupContainingGroup) HasDurationMinutes() bool`

HasDurationMinutes returns a boolean if a field has been set.

### GetAccessLevelRemoteId

`func (o *GroupContainingGroup) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *GroupContainingGroup) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *GroupContainingGroup) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *GroupContainingGroup) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


