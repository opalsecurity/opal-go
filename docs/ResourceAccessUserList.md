# ResourceAccessUserList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** | The cursor with which to continue pagination if additional result pages exist. | [optional] 
**Previous** | Pointer to **string** | The cursor used to obtain the current result page. | [optional] 
**Results** | Pointer to [**[]ResourceAccessUser**](ResourceAccessUser.md) |  | [optional] 

## Methods

### NewResourceAccessUserList

`func NewResourceAccessUserList() *ResourceAccessUserList`

NewResourceAccessUserList instantiates a new ResourceAccessUserList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAccessUserListWithDefaults

`func NewResourceAccessUserListWithDefaults() *ResourceAccessUserList`

NewResourceAccessUserListWithDefaults instantiates a new ResourceAccessUserList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *ResourceAccessUserList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResourceAccessUserList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResourceAccessUserList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResourceAccessUserList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *ResourceAccessUserList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *ResourceAccessUserList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *ResourceAccessUserList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *ResourceAccessUserList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *ResourceAccessUserList) GetResults() []ResourceAccessUser`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResourceAccessUserList) GetResultsOk() (*[]ResourceAccessUser, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResourceAccessUserList) SetResults(v []ResourceAccessUser)`

SetResults sets Results field to given value.

### HasResults

`func (o *ResourceAccessUserList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


