# PaginatedAssignedRequestList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | [**[]Request**](Request.md) |  | 
**Cursor** | **string** | The cursor to continue pagination | 

## Methods

### NewPaginatedAssignedRequestList

`func NewPaginatedAssignedRequestList(requests []Request, cursor string, ) *PaginatedAssignedRequestList`

NewPaginatedAssignedRequestList instantiates a new PaginatedAssignedRequestList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedAssignedRequestListWithDefaults

`func NewPaginatedAssignedRequestListWithDefaults() *PaginatedAssignedRequestList`

NewPaginatedAssignedRequestListWithDefaults instantiates a new PaginatedAssignedRequestList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *PaginatedAssignedRequestList) GetRequests() []Request`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *PaginatedAssignedRequestList) GetRequestsOk() (*[]Request, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *PaginatedAssignedRequestList) SetRequests(v []Request)`

SetRequests sets Requests field to given value.


### GetCursor

`func (o *PaginatedAssignedRequestList) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *PaginatedAssignedRequestList) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *PaginatedAssignedRequestList) SetCursor(v string)`

SetCursor sets Cursor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


