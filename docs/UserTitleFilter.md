# UserTitleFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strings** | Pointer to **[]string** | Exact job titles to match (OR). | [optional] 
**Unset** | Pointer to **bool** | When true, also match users with no job title. | [optional] 
**Not** | Pointer to **bool** | Invert the match within the user domain. | [optional] 

## Methods

### NewUserTitleFilter

`func NewUserTitleFilter() *UserTitleFilter`

NewUserTitleFilter instantiates a new UserTitleFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTitleFilterWithDefaults

`func NewUserTitleFilterWithDefaults() *UserTitleFilter`

NewUserTitleFilterWithDefaults instantiates a new UserTitleFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrings

`func (o *UserTitleFilter) GetStrings() []string`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *UserTitleFilter) GetStringsOk() (*[]string, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *UserTitleFilter) SetStrings(v []string)`

SetStrings sets Strings field to given value.

### HasStrings

`func (o *UserTitleFilter) HasStrings() bool`

HasStrings returns a boolean if a field has been set.

### GetUnset

`func (o *UserTitleFilter) GetUnset() bool`

GetUnset returns the Unset field if non-nil, zero value otherwise.

### GetUnsetOk

`func (o *UserTitleFilter) GetUnsetOk() (*bool, bool)`

GetUnsetOk returns a tuple with the Unset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnset

`func (o *UserTitleFilter) SetUnset(v bool)`

SetUnset sets Unset field to given value.

### HasUnset

`func (o *UserTitleFilter) HasUnset() bool`

HasUnset returns a boolean if a field has been set.

### GetNot

`func (o *UserTitleFilter) GetNot() bool`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *UserTitleFilter) GetNotOk() (*bool, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *UserTitleFilter) SetNot(v bool)`

SetNot sets Not field to given value.

### HasNot

`func (o *UserTitleFilter) HasNot() bool`

HasNot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


