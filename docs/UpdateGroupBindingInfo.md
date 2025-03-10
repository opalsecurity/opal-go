# UpdateGroupBindingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupBindingId** | **string** | The ID of the group binding. | 
**SourceGroupId** | **string** | The ID of the source group. | 
**Groups** | [**[]CreateGroupBindingInfoGroupsInner**](CreateGroupBindingInfoGroupsInner.md) | The list of groups. | 

## Methods

### NewUpdateGroupBindingInfo

`func NewUpdateGroupBindingInfo(groupBindingId string, sourceGroupId string, groups []CreateGroupBindingInfoGroupsInner, ) *UpdateGroupBindingInfo`

NewUpdateGroupBindingInfo instantiates a new UpdateGroupBindingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupBindingInfoWithDefaults

`func NewUpdateGroupBindingInfoWithDefaults() *UpdateGroupBindingInfo`

NewUpdateGroupBindingInfoWithDefaults instantiates a new UpdateGroupBindingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupBindingId

`func (o *UpdateGroupBindingInfo) GetGroupBindingId() string`

GetGroupBindingId returns the GroupBindingId field if non-nil, zero value otherwise.

### GetGroupBindingIdOk

`func (o *UpdateGroupBindingInfo) GetGroupBindingIdOk() (*string, bool)`

GetGroupBindingIdOk returns a tuple with the GroupBindingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBindingId

`func (o *UpdateGroupBindingInfo) SetGroupBindingId(v string)`

SetGroupBindingId sets GroupBindingId field to given value.


### GetSourceGroupId

`func (o *UpdateGroupBindingInfo) GetSourceGroupId() string`

GetSourceGroupId returns the SourceGroupId field if non-nil, zero value otherwise.

### GetSourceGroupIdOk

`func (o *UpdateGroupBindingInfo) GetSourceGroupIdOk() (*string, bool)`

GetSourceGroupIdOk returns a tuple with the SourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroupId

`func (o *UpdateGroupBindingInfo) SetSourceGroupId(v string)`

SetSourceGroupId sets SourceGroupId field to given value.


### GetGroups

`func (o *UpdateGroupBindingInfo) GetGroups() []CreateGroupBindingInfoGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UpdateGroupBindingInfo) GetGroupsOk() (*[]CreateGroupBindingInfoGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UpdateGroupBindingInfo) SetGroups(v []CreateGroupBindingInfoGroupsInner)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


