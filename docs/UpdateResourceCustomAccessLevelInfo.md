# UpdateResourceCustomAccessLevelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevelName** | Pointer to **string** | The new human-readable name. | [optional] 
**Policy** | Pointer to **string** | The new policy document. | [optional] 
**RequestableByDefault** | Pointer to **bool** | Whether the role is requestable. | [optional] 
**StackableSensitivityIndex** | Pointer to **int32** | The new sensitivity index. | [optional] 
**ClearStackableSensitivityIndex** | Pointer to **bool** | Set to true to remove from the hierarchy. | [optional] 

## Methods

### NewUpdateResourceCustomAccessLevelInfo

`func NewUpdateResourceCustomAccessLevelInfo() *UpdateResourceCustomAccessLevelInfo`

NewUpdateResourceCustomAccessLevelInfo instantiates a new UpdateResourceCustomAccessLevelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceCustomAccessLevelInfoWithDefaults

`func NewUpdateResourceCustomAccessLevelInfoWithDefaults() *UpdateResourceCustomAccessLevelInfo`

NewUpdateResourceCustomAccessLevelInfoWithDefaults instantiates a new UpdateResourceCustomAccessLevelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevelName

`func (o *UpdateResourceCustomAccessLevelInfo) GetAccessLevelName() string`

GetAccessLevelName returns the AccessLevelName field if non-nil, zero value otherwise.

### GetAccessLevelNameOk

`func (o *UpdateResourceCustomAccessLevelInfo) GetAccessLevelNameOk() (*string, bool)`

GetAccessLevelNameOk returns a tuple with the AccessLevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelName

`func (o *UpdateResourceCustomAccessLevelInfo) SetAccessLevelName(v string)`

SetAccessLevelName sets AccessLevelName field to given value.

### HasAccessLevelName

`func (o *UpdateResourceCustomAccessLevelInfo) HasAccessLevelName() bool`

HasAccessLevelName returns a boolean if a field has been set.

### GetPolicy

`func (o *UpdateResourceCustomAccessLevelInfo) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpdateResourceCustomAccessLevelInfo) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UpdateResourceCustomAccessLevelInfo) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *UpdateResourceCustomAccessLevelInfo) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetRequestableByDefault

`func (o *UpdateResourceCustomAccessLevelInfo) GetRequestableByDefault() bool`

GetRequestableByDefault returns the RequestableByDefault field if non-nil, zero value otherwise.

### GetRequestableByDefaultOk

`func (o *UpdateResourceCustomAccessLevelInfo) GetRequestableByDefaultOk() (*bool, bool)`

GetRequestableByDefaultOk returns a tuple with the RequestableByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestableByDefault

`func (o *UpdateResourceCustomAccessLevelInfo) SetRequestableByDefault(v bool)`

SetRequestableByDefault sets RequestableByDefault field to given value.

### HasRequestableByDefault

`func (o *UpdateResourceCustomAccessLevelInfo) HasRequestableByDefault() bool`

HasRequestableByDefault returns a boolean if a field has been set.

### GetStackableSensitivityIndex

`func (o *UpdateResourceCustomAccessLevelInfo) GetStackableSensitivityIndex() int32`

GetStackableSensitivityIndex returns the StackableSensitivityIndex field if non-nil, zero value otherwise.

### GetStackableSensitivityIndexOk

`func (o *UpdateResourceCustomAccessLevelInfo) GetStackableSensitivityIndexOk() (*int32, bool)`

GetStackableSensitivityIndexOk returns a tuple with the StackableSensitivityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackableSensitivityIndex

`func (o *UpdateResourceCustomAccessLevelInfo) SetStackableSensitivityIndex(v int32)`

SetStackableSensitivityIndex sets StackableSensitivityIndex field to given value.

### HasStackableSensitivityIndex

`func (o *UpdateResourceCustomAccessLevelInfo) HasStackableSensitivityIndex() bool`

HasStackableSensitivityIndex returns a boolean if a field has been set.

### GetClearStackableSensitivityIndex

`func (o *UpdateResourceCustomAccessLevelInfo) GetClearStackableSensitivityIndex() bool`

GetClearStackableSensitivityIndex returns the ClearStackableSensitivityIndex field if non-nil, zero value otherwise.

### GetClearStackableSensitivityIndexOk

`func (o *UpdateResourceCustomAccessLevelInfo) GetClearStackableSensitivityIndexOk() (*bool, bool)`

GetClearStackableSensitivityIndexOk returns a tuple with the ClearStackableSensitivityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearStackableSensitivityIndex

`func (o *UpdateResourceCustomAccessLevelInfo) SetClearStackableSensitivityIndex(v bool)`

SetClearStackableSensitivityIndex sets ClearStackableSensitivityIndex field to given value.

### HasClearStackableSensitivityIndex

`func (o *UpdateResourceCustomAccessLevelInfo) HasClearStackableSensitivityIndex() bool`

HasClearStackableSensitivityIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


