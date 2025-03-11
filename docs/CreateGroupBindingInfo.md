# CreateGroupBindingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceGroupId** | **string** | The ID of the source group. | 
**Groups** | [**[]CreateGroupBindingInfoGroupsInner**](CreateGroupBindingInfoGroupsInner.md) | The list of groups. | 

## Methods

### NewCreateGroupBindingInfo

`func NewCreateGroupBindingInfo(sourceGroupId string, groups []CreateGroupBindingInfoGroupsInner, ) *CreateGroupBindingInfo`

NewCreateGroupBindingInfo instantiates a new CreateGroupBindingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupBindingInfoWithDefaults

`func NewCreateGroupBindingInfoWithDefaults() *CreateGroupBindingInfo`

NewCreateGroupBindingInfoWithDefaults instantiates a new CreateGroupBindingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceGroupId

`func (o *CreateGroupBindingInfo) GetSourceGroupId() string`

GetSourceGroupId returns the SourceGroupId field if non-nil, zero value otherwise.

### GetSourceGroupIdOk

`func (o *CreateGroupBindingInfo) GetSourceGroupIdOk() (*string, bool)`

GetSourceGroupIdOk returns a tuple with the SourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroupId

`func (o *CreateGroupBindingInfo) SetSourceGroupId(v string)`

SetSourceGroupId sets SourceGroupId field to given value.


### GetGroups

`func (o *CreateGroupBindingInfo) GetGroups() []CreateGroupBindingInfoGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CreateGroupBindingInfo) GetGroupsOk() (*[]CreateGroupBindingInfoGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CreateGroupBindingInfo) SetGroups(v []CreateGroupBindingInfoGroupsInner)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


