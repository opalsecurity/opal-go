# UserTeamFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strings** | Pointer to **[]string** | Exact team names to match (OR). | [optional] 
**Unset** | Pointer to **bool** | When true, also match users with no team. | [optional] 
**Not** | Pointer to **bool** | Invert the match within the user domain. | [optional] 

## Methods

### NewUserTeamFilter

`func NewUserTeamFilter() *UserTeamFilter`

NewUserTeamFilter instantiates a new UserTeamFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTeamFilterWithDefaults

`func NewUserTeamFilterWithDefaults() *UserTeamFilter`

NewUserTeamFilterWithDefaults instantiates a new UserTeamFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrings

`func (o *UserTeamFilter) GetStrings() []string`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *UserTeamFilter) GetStringsOk() (*[]string, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *UserTeamFilter) SetStrings(v []string)`

SetStrings sets Strings field to given value.

### HasStrings

`func (o *UserTeamFilter) HasStrings() bool`

HasStrings returns a boolean if a field has been set.

### GetUnset

`func (o *UserTeamFilter) GetUnset() bool`

GetUnset returns the Unset field if non-nil, zero value otherwise.

### GetUnsetOk

`func (o *UserTeamFilter) GetUnsetOk() (*bool, bool)`

GetUnsetOk returns a tuple with the Unset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnset

`func (o *UserTeamFilter) SetUnset(v bool)`

SetUnset sets Unset field to given value.

### HasUnset

`func (o *UserTeamFilter) HasUnset() bool`

HasUnset returns a boolean if a field has been set.

### GetNot

`func (o *UserTeamFilter) GetNot() bool`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *UserTeamFilter) GetNotOk() (*bool, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *UserTeamFilter) SetNot(v bool)`

SetNot sets Not field to given value.

### HasNot

`func (o *UserTeamFilter) HasNot() bool`

HasNot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


