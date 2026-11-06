# RemoveRoleAssignmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskIds** | **[]string** | Opaque identifiers for the push tasks enqueued to propagate the revocation to the end system. Encoded identically to the &#x60;PushTaskId&#x60; scalar in the GraphQL API (base64 of the underlying protobuf). One ID per deferred revocation (today, resource/user grants are always deferred); empty when every revocation completed inline (resource/group, group/user, group/group). | 
**Entries** | [**[]RemoveRoleAssignmentsEntry**](RemoveRoleAssignmentsEntry.md) | One entry per role assignment that was removed inline. Shorter than the request when some assignments did not exist (idempotent), and empty when every revocation in the batch was deferred to an async propagation task — in that case &#x60;task_ids&#x60; carries one ID per deferred revocation and &#x60;entries&#x60; is empty. Resource/user grants are currently always deferred; the other &#x60;(entity, principal)&#x60; combinations complete inline and populate &#x60;entries&#x60;. | 

## Methods

### NewRemoveRoleAssignmentsResponse

`func NewRemoveRoleAssignmentsResponse(taskIds []string, entries []RemoveRoleAssignmentsEntry, ) *RemoveRoleAssignmentsResponse`

NewRemoveRoleAssignmentsResponse instantiates a new RemoveRoleAssignmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveRoleAssignmentsResponseWithDefaults

`func NewRemoveRoleAssignmentsResponseWithDefaults() *RemoveRoleAssignmentsResponse`

NewRemoveRoleAssignmentsResponseWithDefaults instantiates a new RemoveRoleAssignmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskIds

`func (o *RemoveRoleAssignmentsResponse) GetTaskIds() []string`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *RemoveRoleAssignmentsResponse) GetTaskIdsOk() (*[]string, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *RemoveRoleAssignmentsResponse) SetTaskIds(v []string)`

SetTaskIds sets TaskIds field to given value.


### GetEntries

`func (o *RemoveRoleAssignmentsResponse) GetEntries() []RemoveRoleAssignmentsEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *RemoveRoleAssignmentsResponse) GetEntriesOk() (*[]RemoveRoleAssignmentsEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *RemoveRoleAssignmentsResponse) SetEntries(v []RemoveRoleAssignmentsEntry)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


