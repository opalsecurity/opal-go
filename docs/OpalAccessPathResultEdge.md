# OpalAccessPathResultEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | [**OpalAccessPathResultNode**](OpalAccessPathResultNode.md) |  | 
**Cursor** | **string** | Opaque cursor for this path, used for pagination. | 

## Methods

### NewOpalAccessPathResultEdge

`func NewOpalAccessPathResultEdge(node OpalAccessPathResultNode, cursor string, ) *OpalAccessPathResultEdge`

NewOpalAccessPathResultEdge instantiates a new OpalAccessPathResultEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalAccessPathResultEdgeWithDefaults

`func NewOpalAccessPathResultEdgeWithDefaults() *OpalAccessPathResultEdge`

NewOpalAccessPathResultEdgeWithDefaults instantiates a new OpalAccessPathResultEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *OpalAccessPathResultEdge) GetNode() OpalAccessPathResultNode`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *OpalAccessPathResultEdge) GetNodeOk() (*OpalAccessPathResultNode, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *OpalAccessPathResultEdge) SetNode(v OpalAccessPathResultNode)`

SetNode sets Node field to given value.


### GetCursor

`func (o *OpalAccessPathResultEdge) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *OpalAccessPathResultEdge) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *OpalAccessPathResultEdge) SetCursor(v string)`

SetCursor sets Cursor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


