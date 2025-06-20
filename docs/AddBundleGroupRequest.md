# AddBundleGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The ID of the group to add. | 
**AccessLevelRemoteId** | Pointer to **string** | The remote ID of the access level to grant to this user. Required if the group being added requires an access level. If omitted, the default access level remote ID value (empty string) is used. | [optional] 
**AccessLevelName** | Pointer to **string** | The name of the access level to grant to this user. If omitted, the default access level name value (empty string) is used. | [optional] 

## Methods

### NewAddBundleGroupRequest

`func NewAddBundleGroupRequest(groupId string, ) *AddBundleGroupRequest`

NewAddBundleGroupRequest instantiates a new AddBundleGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBundleGroupRequestWithDefaults

`func NewAddBundleGroupRequestWithDefaults() *AddBundleGroupRequest`

NewAddBundleGroupRequestWithDefaults instantiates a new AddBundleGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *AddBundleGroupRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AddBundleGroupRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AddBundleGroupRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetAccessLevelRemoteId

`func (o *AddBundleGroupRequest) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *AddBundleGroupRequest) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *AddBundleGroupRequest) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *AddBundleGroupRequest) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.

### GetAccessLevelName

`func (o *AddBundleGroupRequest) GetAccessLevelName() string`

GetAccessLevelName returns the AccessLevelName field if non-nil, zero value otherwise.

### GetAccessLevelNameOk

`func (o *AddBundleGroupRequest) GetAccessLevelNameOk() (*string, bool)`

GetAccessLevelNameOk returns a tuple with the AccessLevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelName

`func (o *AddBundleGroupRequest) SetAccessLevelName(v string)`

SetAccessLevelName sets AccessLevelName field to given value.

### HasAccessLevelName

`func (o *AddBundleGroupRequest) HasAccessLevelName() bool`

HasAccessLevelName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


