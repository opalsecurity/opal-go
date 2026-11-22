# AddRoleAssignmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskIds** | **[]string** | Opaque IDs of push tasks enqueued to propagate the new assignments to end systems. Each ID is a base64-encoded, URL-safe string that mirrors the identifier returned by the GraphQL &#x60;addRoleAssignments&#x60; mutation. | 
**Entries** | [**[]AddRoleAssignmentsEntry**](AddRoleAssignmentsEntry.md) | One entry per role assignment created, in request order. | 

## Methods

### NewAddRoleAssignmentsResponse

`func NewAddRoleAssignmentsResponse(taskIds []string, entries []AddRoleAssignmentsEntry, ) *AddRoleAssignmentsResponse`

NewAddRoleAssignmentsResponse instantiates a new AddRoleAssignmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRoleAssignmentsResponseWithDefaults

`func NewAddRoleAssignmentsResponseWithDefaults() *AddRoleAssignmentsResponse`

NewAddRoleAssignmentsResponseWithDefaults instantiates a new AddRoleAssignmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskIds

`func (o *AddRoleAssignmentsResponse) GetTaskIds() []string`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *AddRoleAssignmentsResponse) GetTaskIdsOk() (*[]string, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *AddRoleAssignmentsResponse) SetTaskIds(v []string)`

SetTaskIds sets TaskIds field to given value.


### GetEntries

`func (o *AddRoleAssignmentsResponse) GetEntries() []AddRoleAssignmentsEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *AddRoleAssignmentsResponse) GetEntriesOk() (*[]AddRoleAssignmentsEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *AddRoleAssignmentsResponse) SetEntries(v []AddRoleAssignmentsEntry)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


