# PaginatedPermissionUserList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **NullableString** | The cursor with which to continue pagination if additional result pages exist. | [optional] 
**Previous** | Pointer to **NullableString** | The cursor used to obtain the current result page. | [optional] 
**Results** | Pointer to [**[]PermissionUser**](PermissionUser.md) |  | [optional] 

## Methods

### NewPaginatedPermissionUserList

`func NewPaginatedPermissionUserList() *PaginatedPermissionUserList`

NewPaginatedPermissionUserList instantiates a new PaginatedPermissionUserList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedPermissionUserListWithDefaults

`func NewPaginatedPermissionUserListWithDefaults() *PaginatedPermissionUserList`

NewPaginatedPermissionUserListWithDefaults instantiates a new PaginatedPermissionUserList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginatedPermissionUserList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedPermissionUserList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedPermissionUserList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedPermissionUserList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedPermissionUserList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedPermissionUserList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedPermissionUserList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedPermissionUserList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedPermissionUserList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedPermissionUserList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedPermissionUserList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedPermissionUserList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedPermissionUserList) GetResults() []PermissionUser`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedPermissionUserList) GetResultsOk() (*[]PermissionUser, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedPermissionUserList) SetResults(v []PermissionUser)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedPermissionUserList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


