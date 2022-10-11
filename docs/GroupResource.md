# GroupResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The ID of the group. | 
**ResourceId** | **string** | The ID of the resource. | 
**AccessLevel** | [**ResourceAccessLevel**](ResourceAccessLevel.md) |  | 

## Methods

### NewGroupResource

`func NewGroupResource(groupId string, resourceId string, accessLevel ResourceAccessLevel, ) *GroupResource`

NewGroupResource instantiates a new GroupResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupResourceWithDefaults

`func NewGroupResourceWithDefaults() *GroupResource`

NewGroupResourceWithDefaults instantiates a new GroupResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupResource) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupResource) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupResource) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetResourceId

`func (o *GroupResource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *GroupResource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *GroupResource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetAccessLevel

`func (o *GroupResource) GetAccessLevel() ResourceAccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *GroupResource) GetAccessLevelOk() (*ResourceAccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *GroupResource) SetAccessLevel(v ResourceAccessLevel)`

SetAccessLevel sets AccessLevel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


