# PaginatedRemoteUsersList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** | The cursor with which to continue pagination if additional result pages exist. | [optional] 
**Previous** | Pointer to **string** | The cursor used to obtain the current result page. | [optional] 
**Results** | [**[]RemoteUser**](RemoteUser.md) |  | 

## Methods

### NewPaginatedRemoteUsersList

`func NewPaginatedRemoteUsersList(results []RemoteUser, ) *PaginatedRemoteUsersList`

NewPaginatedRemoteUsersList instantiates a new PaginatedRemoteUsersList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedRemoteUsersListWithDefaults

`func NewPaginatedRemoteUsersListWithDefaults() *PaginatedRemoteUsersList`

NewPaginatedRemoteUsersListWithDefaults instantiates a new PaginatedRemoteUsersList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginatedRemoteUsersList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedRemoteUsersList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedRemoteUsersList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedRemoteUsersList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *PaginatedRemoteUsersList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedRemoteUsersList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedRemoteUsersList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedRemoteUsersList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedRemoteUsersList) GetResults() []RemoteUser`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedRemoteUsersList) GetResultsOk() (*[]RemoteUser, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedRemoteUsersList) SetResults(v []RemoteUser)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


