# CreateResourceCustomAccessLevelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevel** | [**ResourceAccessLevel**](ResourceAccessLevel.md) |  | 
**Policy** | Pointer to **string** | The policy document. | [optional] 
**RequestableByDefault** | Pointer to **bool** | Whether the role is requestable. Defaults to false. | [optional] 
**StackableSensitivityIndex** | Pointer to **int32** | The sensitivity index. Null to leave unranked. | [optional] 

## Methods

### NewCreateResourceCustomAccessLevelInfo

`func NewCreateResourceCustomAccessLevelInfo(accessLevel ResourceAccessLevel, ) *CreateResourceCustomAccessLevelInfo`

NewCreateResourceCustomAccessLevelInfo instantiates a new CreateResourceCustomAccessLevelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResourceCustomAccessLevelInfoWithDefaults

`func NewCreateResourceCustomAccessLevelInfoWithDefaults() *CreateResourceCustomAccessLevelInfo`

NewCreateResourceCustomAccessLevelInfoWithDefaults instantiates a new CreateResourceCustomAccessLevelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevel

`func (o *CreateResourceCustomAccessLevelInfo) GetAccessLevel() ResourceAccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *CreateResourceCustomAccessLevelInfo) GetAccessLevelOk() (*ResourceAccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *CreateResourceCustomAccessLevelInfo) SetAccessLevel(v ResourceAccessLevel)`

SetAccessLevel sets AccessLevel field to given value.


### GetPolicy

`func (o *CreateResourceCustomAccessLevelInfo) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CreateResourceCustomAccessLevelInfo) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CreateResourceCustomAccessLevelInfo) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *CreateResourceCustomAccessLevelInfo) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetRequestableByDefault

`func (o *CreateResourceCustomAccessLevelInfo) GetRequestableByDefault() bool`

GetRequestableByDefault returns the RequestableByDefault field if non-nil, zero value otherwise.

### GetRequestableByDefaultOk

`func (o *CreateResourceCustomAccessLevelInfo) GetRequestableByDefaultOk() (*bool, bool)`

GetRequestableByDefaultOk returns a tuple with the RequestableByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestableByDefault

`func (o *CreateResourceCustomAccessLevelInfo) SetRequestableByDefault(v bool)`

SetRequestableByDefault sets RequestableByDefault field to given value.

### HasRequestableByDefault

`func (o *CreateResourceCustomAccessLevelInfo) HasRequestableByDefault() bool`

HasRequestableByDefault returns a boolean if a field has been set.

### GetStackableSensitivityIndex

`func (o *CreateResourceCustomAccessLevelInfo) GetStackableSensitivityIndex() int32`

GetStackableSensitivityIndex returns the StackableSensitivityIndex field if non-nil, zero value otherwise.

### GetStackableSensitivityIndexOk

`func (o *CreateResourceCustomAccessLevelInfo) GetStackableSensitivityIndexOk() (*int32, bool)`

GetStackableSensitivityIndexOk returns a tuple with the StackableSensitivityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackableSensitivityIndex

`func (o *CreateResourceCustomAccessLevelInfo) SetStackableSensitivityIndex(v int32)`

SetStackableSensitivityIndex sets StackableSensitivityIndex field to given value.

### HasStackableSensitivityIndex

`func (o *CreateResourceCustomAccessLevelInfo) HasStackableSensitivityIndex() bool`

HasStackableSensitivityIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


