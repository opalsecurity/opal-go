# AddRoleAssignmentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleAssignments** | [**[]AddRoleAssignmentInput**](AddRoleAssignmentInput.md) | The list of role assignments to add. | 

## Methods

### NewAddRoleAssignmentsRequest

`func NewAddRoleAssignmentsRequest(roleAssignments []AddRoleAssignmentInput, ) *AddRoleAssignmentsRequest`

NewAddRoleAssignmentsRequest instantiates a new AddRoleAssignmentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRoleAssignmentsRequestWithDefaults

`func NewAddRoleAssignmentsRequestWithDefaults() *AddRoleAssignmentsRequest`

NewAddRoleAssignmentsRequestWithDefaults instantiates a new AddRoleAssignmentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleAssignments

`func (o *AddRoleAssignmentsRequest) GetRoleAssignments() []AddRoleAssignmentInput`

GetRoleAssignments returns the RoleAssignments field if non-nil, zero value otherwise.

### GetRoleAssignmentsOk

`func (o *AddRoleAssignmentsRequest) GetRoleAssignmentsOk() (*[]AddRoleAssignmentInput, bool)`

GetRoleAssignmentsOk returns a tuple with the RoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAssignments

`func (o *AddRoleAssignmentsRequest) SetRoleAssignments(v []AddRoleAssignmentInput)`

SetRoleAssignments sets RoleAssignments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


