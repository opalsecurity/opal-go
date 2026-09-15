# OpalAccessResultEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | [**OpalAccessResultNode**](OpalAccessResultNode.md) |  | 
**Cursor** | **string** | Opaque cursor for this grant, used for pagination. | 

## Methods

### NewOpalAccessResultEdge

`func NewOpalAccessResultEdge(node OpalAccessResultNode, cursor string, ) *OpalAccessResultEdge`

NewOpalAccessResultEdge instantiates a new OpalAccessResultEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalAccessResultEdgeWithDefaults

`func NewOpalAccessResultEdgeWithDefaults() *OpalAccessResultEdge`

NewOpalAccessResultEdgeWithDefaults instantiates a new OpalAccessResultEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *OpalAccessResultEdge) GetNode() OpalAccessResultNode`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *OpalAccessResultEdge) GetNodeOk() (*OpalAccessResultNode, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *OpalAccessResultEdge) SetNode(v OpalAccessResultNode)`

SetNode sets Node field to given value.


### GetCursor

`func (o *OpalAccessResultEdge) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *OpalAccessResultEdge) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *OpalAccessResultEdge) SetCursor(v string)`

SetCursor sets Cursor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


