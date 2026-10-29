# PaginatedRoleAssignmentsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** | The cursor with which to continue pagination if additional result pages exist. | [optional] 
**Previous** | Pointer to **string** | The cursor used to obtain the current result page. | [optional] 
**Results** | [**[]RoleAssignment**](RoleAssignment.md) |  | 

## Methods

### NewPaginatedRoleAssignmentsList

`func NewPaginatedRoleAssignmentsList(results []RoleAssignment, ) *PaginatedRoleAssignmentsList`

NewPaginatedRoleAssignmentsList instantiates a new PaginatedRoleAssignmentsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedRoleAssignmentsListWithDefaults

`func NewPaginatedRoleAssignmentsListWithDefaults() *PaginatedRoleAssignmentsList`

NewPaginatedRoleAssignmentsListWithDefaults instantiates a new PaginatedRoleAssignmentsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginatedRoleAssignmentsList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedRoleAssignmentsList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedRoleAssignmentsList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedRoleAssignmentsList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *PaginatedRoleAssignmentsList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedRoleAssignmentsList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedRoleAssignmentsList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedRoleAssignmentsList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedRoleAssignmentsList) GetResults() []RoleAssignment`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedRoleAssignmentsList) GetResultsOk() (*[]RoleAssignment, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedRoleAssignmentsList) SetResults(v []RoleAssignment)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


