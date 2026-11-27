# ResourceCustomAccessLevelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceCustomAccessLevelId** | **string** | The unique ID of the custom access level. | 
**ResourceId** | **string** | The resource this access level belongs to. | 
**AccessLevel** | [**ResourceAccessLevel**](ResourceAccessLevel.md) |  | 
**Policy** | Pointer to **string** | The policy document for this access level. | [optional] 
**RequestableByDefault** | **bool** | Whether this role is requestable. | 
**StackableSensitivityIndex** | Pointer to **int32** | The sensitivity index in the access hierarchy. Null if unranked. | [optional] 
**MemberResourceIds** | Pointer to **[]string** | When addressed via a parent resource, lists the child resource IDs that back this aggregated entry. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewResourceCustomAccessLevelResponse

`func NewResourceCustomAccessLevelResponse(resourceCustomAccessLevelId string, resourceId string, accessLevel ResourceAccessLevel, requestableByDefault bool, ) *ResourceCustomAccessLevelResponse`

NewResourceCustomAccessLevelResponse instantiates a new ResourceCustomAccessLevelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceCustomAccessLevelResponseWithDefaults

`func NewResourceCustomAccessLevelResponseWithDefaults() *ResourceCustomAccessLevelResponse`

NewResourceCustomAccessLevelResponseWithDefaults instantiates a new ResourceCustomAccessLevelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceCustomAccessLevelId

`func (o *ResourceCustomAccessLevelResponse) GetResourceCustomAccessLevelId() string`

GetResourceCustomAccessLevelId returns the ResourceCustomAccessLevelId field if non-nil, zero value otherwise.

### GetResourceCustomAccessLevelIdOk

`func (o *ResourceCustomAccessLevelResponse) GetResourceCustomAccessLevelIdOk() (*string, bool)`

GetResourceCustomAccessLevelIdOk returns a tuple with the ResourceCustomAccessLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCustomAccessLevelId

`func (o *ResourceCustomAccessLevelResponse) SetResourceCustomAccessLevelId(v string)`

SetResourceCustomAccessLevelId sets ResourceCustomAccessLevelId field to given value.


### GetResourceId

`func (o *ResourceCustomAccessLevelResponse) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceCustomAccessLevelResponse) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceCustomAccessLevelResponse) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetAccessLevel

`func (o *ResourceCustomAccessLevelResponse) GetAccessLevel() ResourceAccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *ResourceCustomAccessLevelResponse) GetAccessLevelOk() (*ResourceAccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *ResourceCustomAccessLevelResponse) SetAccessLevel(v ResourceAccessLevel)`

SetAccessLevel sets AccessLevel field to given value.


### GetPolicy

`func (o *ResourceCustomAccessLevelResponse) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ResourceCustomAccessLevelResponse) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ResourceCustomAccessLevelResponse) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ResourceCustomAccessLevelResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetRequestableByDefault

`func (o *ResourceCustomAccessLevelResponse) GetRequestableByDefault() bool`

GetRequestableByDefault returns the RequestableByDefault field if non-nil, zero value otherwise.

### GetRequestableByDefaultOk

`func (o *ResourceCustomAccessLevelResponse) GetRequestableByDefaultOk() (*bool, bool)`

GetRequestableByDefaultOk returns a tuple with the RequestableByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestableByDefault

`func (o *ResourceCustomAccessLevelResponse) SetRequestableByDefault(v bool)`

SetRequestableByDefault sets RequestableByDefault field to given value.


### GetStackableSensitivityIndex

`func (o *ResourceCustomAccessLevelResponse) GetStackableSensitivityIndex() int32`

GetStackableSensitivityIndex returns the StackableSensitivityIndex field if non-nil, zero value otherwise.

### GetStackableSensitivityIndexOk

`func (o *ResourceCustomAccessLevelResponse) GetStackableSensitivityIndexOk() (*int32, bool)`

GetStackableSensitivityIndexOk returns a tuple with the StackableSensitivityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackableSensitivityIndex

`func (o *ResourceCustomAccessLevelResponse) SetStackableSensitivityIndex(v int32)`

SetStackableSensitivityIndex sets StackableSensitivityIndex field to given value.

### HasStackableSensitivityIndex

`func (o *ResourceCustomAccessLevelResponse) HasStackableSensitivityIndex() bool`

HasStackableSensitivityIndex returns a boolean if a field has been set.

### GetMemberResourceIds

`func (o *ResourceCustomAccessLevelResponse) GetMemberResourceIds() []string`

GetMemberResourceIds returns the MemberResourceIds field if non-nil, zero value otherwise.

### GetMemberResourceIdsOk

`func (o *ResourceCustomAccessLevelResponse) GetMemberResourceIdsOk() (*[]string, bool)`

GetMemberResourceIdsOk returns a tuple with the MemberResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberResourceIds

`func (o *ResourceCustomAccessLevelResponse) SetMemberResourceIds(v []string)`

SetMemberResourceIds sets MemberResourceIds field to given value.

### HasMemberResourceIds

`func (o *ResourceCustomAccessLevelResponse) HasMemberResourceIds() bool`

HasMemberResourceIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResourceCustomAccessLevelResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceCustomAccessLevelResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceCustomAccessLevelResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResourceCustomAccessLevelResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ResourceCustomAccessLevelResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceCustomAccessLevelResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceCustomAccessLevelResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ResourceCustomAccessLevelResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


