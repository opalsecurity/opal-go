# GroupUserList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]GroupUser**](GroupUser.md) |  | [optional] 
**Next** | Pointer to **string** | The cursor with which to continue pagination if additional result pages exist. | [optional] 
**Previous** | Pointer to **string** | The cursor used to obtain the current result page. | [optional] 

## Methods

### NewGroupUserList

`func NewGroupUserList() *GroupUserList`

NewGroupUserList instantiates a new GroupUserList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUserListWithDefaults

`func NewGroupUserListWithDefaults() *GroupUserList`

NewGroupUserListWithDefaults instantiates a new GroupUserList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GroupUserList) GetResults() []GroupUser`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GroupUserList) GetResultsOk() (*[]GroupUser, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GroupUserList) SetResults(v []GroupUser)`

SetResults sets Results field to given value.

### HasResults

`func (o *GroupUserList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetNext

`func (o *GroupUserList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GroupUserList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GroupUserList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *GroupUserList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *GroupUserList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *GroupUserList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *GroupUserList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *GroupUserList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


