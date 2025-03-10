# GroupBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupBindingId** | **string** | The ID of the group binding. | 
**CreatedById** | **string** | The ID of the user that created the group binding. | 
**CreatedAt** | **time.Time** | The date the group binding was created. | 
**SourceGroupId** | **string** | The ID of the source group. | 
**Groups** | [**[]GroupBindingGroup**](GroupBindingGroup.md) | The list of groups. | 

## Methods

### NewGroupBinding

`func NewGroupBinding(groupBindingId string, createdById string, createdAt time.Time, sourceGroupId string, groups []GroupBindingGroup, ) *GroupBinding`

NewGroupBinding instantiates a new GroupBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupBindingWithDefaults

`func NewGroupBindingWithDefaults() *GroupBinding`

NewGroupBindingWithDefaults instantiates a new GroupBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupBindingId

`func (o *GroupBinding) GetGroupBindingId() string`

GetGroupBindingId returns the GroupBindingId field if non-nil, zero value otherwise.

### GetGroupBindingIdOk

`func (o *GroupBinding) GetGroupBindingIdOk() (*string, bool)`

GetGroupBindingIdOk returns a tuple with the GroupBindingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBindingId

`func (o *GroupBinding) SetGroupBindingId(v string)`

SetGroupBindingId sets GroupBindingId field to given value.


### GetCreatedById

`func (o *GroupBinding) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *GroupBinding) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *GroupBinding) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetCreatedAt

`func (o *GroupBinding) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupBinding) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupBinding) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSourceGroupId

`func (o *GroupBinding) GetSourceGroupId() string`

GetSourceGroupId returns the SourceGroupId field if non-nil, zero value otherwise.

### GetSourceGroupIdOk

`func (o *GroupBinding) GetSourceGroupIdOk() (*string, bool)`

GetSourceGroupIdOk returns a tuple with the SourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroupId

`func (o *GroupBinding) SetSourceGroupId(v string)`

SetSourceGroupId sets SourceGroupId field to given value.


### GetGroups

`func (o *GroupBinding) GetGroups() []GroupBindingGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GroupBinding) GetGroupsOk() (*[]GroupBindingGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GroupBinding) SetGroups(v []GroupBindingGroup)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


