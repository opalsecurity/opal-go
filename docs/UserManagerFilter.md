# UserManagerFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | Pointer to **[]string** | Manager user IDs to match (OR). | [optional] 
**Unset** | Pointer to **bool** | When true, also match users with no manager. | [optional] 
**Not** | Pointer to **bool** | Invert the match within the user domain. | [optional] 

## Methods

### NewUserManagerFilter

`func NewUserManagerFilter() *UserManagerFilter`

NewUserManagerFilter instantiates a new UserManagerFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserManagerFilterWithDefaults

`func NewUserManagerFilterWithDefaults() *UserManagerFilter`

NewUserManagerFilterWithDefaults instantiates a new UserManagerFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *UserManagerFilter) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *UserManagerFilter) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *UserManagerFilter) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *UserManagerFilter) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetUnset

`func (o *UserManagerFilter) GetUnset() bool`

GetUnset returns the Unset field if non-nil, zero value otherwise.

### GetUnsetOk

`func (o *UserManagerFilter) GetUnsetOk() (*bool, bool)`

GetUnsetOk returns a tuple with the Unset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnset

`func (o *UserManagerFilter) SetUnset(v bool)`

SetUnset sets Unset field to given value.

### HasUnset

`func (o *UserManagerFilter) HasUnset() bool`

HasUnset returns a boolean if a field has been set.

### GetNot

`func (o *UserManagerFilter) GetNot() bool`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *UserManagerFilter) GetNotOk() (*bool, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *UserManagerFilter) SetNot(v bool)`

SetNot sets Not field to given value.

### HasNot

`func (o *UserManagerFilter) HasNot() bool`

HasNot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


