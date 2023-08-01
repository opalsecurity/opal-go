# RequestList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to [**[]Request**](Request.md) | The list of requests. | [optional] 
**Cursor** | Pointer to **string** | The cursor to use in the next request to get the next page of results. | [optional] 

## Methods

### NewRequestList

`func NewRequestList() *RequestList`

NewRequestList instantiates a new RequestList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestListWithDefaults

`func NewRequestListWithDefaults() *RequestList`

NewRequestListWithDefaults instantiates a new RequestList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *RequestList) GetRequests() []Request`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *RequestList) GetRequestsOk() (*[]Request, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *RequestList) SetRequests(v []Request)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *RequestList) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetCursor

`func (o *RequestList) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *RequestList) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *RequestList) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *RequestList) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


