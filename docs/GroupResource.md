# GroupResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The ID of the group. | 
**ResourceId** | **string** | The ID of the resource. | 
**GroupName** | Pointer to **string** | The name of the group | [optional] 
**ResourceName** | Pointer to **string** | The name of the resource | [optional] 
**ExpirationDate** | Pointer to **time.Time** | The day and time the group&#39;s access will expire. | [optional] 
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


### GetGroupName

`func (o *GroupResource) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GroupResource) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GroupResource) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GroupResource) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetResourceName

`func (o *GroupResource) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *GroupResource) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *GroupResource) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *GroupResource) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetExpirationDate

`func (o *GroupResource) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *GroupResource) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *GroupResource) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *GroupResource) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

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


