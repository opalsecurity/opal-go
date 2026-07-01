# RequestEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | [**Request**](Request.md) |  | 
**Cursor** | **string** | The cursor for this request edge | 

## Methods

### NewRequestEdge

`func NewRequestEdge(node Request, cursor string, ) *RequestEdge`

NewRequestEdge instantiates a new RequestEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestEdgeWithDefaults

`func NewRequestEdgeWithDefaults() *RequestEdge`

NewRequestEdgeWithDefaults instantiates a new RequestEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *RequestEdge) GetNode() Request`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *RequestEdge) GetNodeOk() (*Request, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *RequestEdge) SetNode(v Request)`

SetNode sets Node field to given value.


### GetCursor

`func (o *RequestEdge) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *RequestEdge) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *RequestEdge) SetCursor(v string)`

SetCursor sets Cursor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


