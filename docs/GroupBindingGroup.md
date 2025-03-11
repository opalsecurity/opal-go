# GroupBindingGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The ID of the group. | 
**GroupType** | [**GroupTypeEnum**](GroupTypeEnum.md) |  | 

## Methods

### NewGroupBindingGroup

`func NewGroupBindingGroup(groupId string, groupType GroupTypeEnum, ) *GroupBindingGroup`

NewGroupBindingGroup instantiates a new GroupBindingGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupBindingGroupWithDefaults

`func NewGroupBindingGroupWithDefaults() *GroupBindingGroup`

NewGroupBindingGroupWithDefaults instantiates a new GroupBindingGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupBindingGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupBindingGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupBindingGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetGroupType

`func (o *GroupBindingGroup) GetGroupType() GroupTypeEnum`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *GroupBindingGroup) GetGroupTypeOk() (*GroupTypeEnum, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *GroupBindingGroup) SetGroupType(v GroupTypeEnum)`

SetGroupType sets GroupType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


