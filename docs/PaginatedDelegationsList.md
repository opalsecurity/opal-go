# PaginatedDelegationsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Delegation**](Delegation.md) | The delegations in the result set. | [optional] 
**Next** | Pointer to **string** | The cursor with which to continue pagination if additional result pages exist. | [optional] 
**Previous** | Pointer to **string** | The cursor used to obtain the current result page. | [optional] 
**TotalCount** | Pointer to **int32** | The total number of items in the result set. | [optional] 

## Methods

### NewPaginatedDelegationsList

`func NewPaginatedDelegationsList() *PaginatedDelegationsList`

NewPaginatedDelegationsList instantiates a new PaginatedDelegationsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedDelegationsListWithDefaults

`func NewPaginatedDelegationsListWithDefaults() *PaginatedDelegationsList`

NewPaginatedDelegationsListWithDefaults instantiates a new PaginatedDelegationsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PaginatedDelegationsList) GetResults() []Delegation`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedDelegationsList) GetResultsOk() (*[]Delegation, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedDelegationsList) SetResults(v []Delegation)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedDelegationsList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedDelegationsList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedDelegationsList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedDelegationsList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedDelegationsList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *PaginatedDelegationsList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedDelegationsList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedDelegationsList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedDelegationsList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetTotalCount

`func (o *PaginatedDelegationsList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedDelegationsList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedDelegationsList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedDelegationsList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


