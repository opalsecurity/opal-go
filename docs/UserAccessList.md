# UserAccessList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]UserAccess**](UserAccess.md) |  | 
**Next** | Pointer to **string** | The cursor with which to continue pagination if additional result pages exist. | [optional] 
**Previous** | Pointer to **string** | The cursor used to obtain the current result page. | [optional] 

## Methods

### NewUserAccessList

`func NewUserAccessList(results []UserAccess, ) *UserAccessList`

NewUserAccessList instantiates a new UserAccessList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccessListWithDefaults

`func NewUserAccessListWithDefaults() *UserAccessList`

NewUserAccessListWithDefaults instantiates a new UserAccessList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *UserAccessList) GetResults() []UserAccess`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UserAccessList) GetResultsOk() (*[]UserAccess, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UserAccessList) SetResults(v []UserAccess)`

SetResults sets Results field to given value.


### GetNext

`func (o *UserAccessList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *UserAccessList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *UserAccessList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *UserAccessList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *UserAccessList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *UserAccessList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *UserAccessList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *UserAccessList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


