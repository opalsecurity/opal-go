# OpalQueryResultEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | [**OpalQueryResultNode**](OpalQueryResultNode.md) |  | 
**Cursor** | **string** | Opaque cursor for this entity, used for pagination. | 

## Methods

### NewOpalQueryResultEdge

`func NewOpalQueryResultEdge(node OpalQueryResultNode, cursor string, ) *OpalQueryResultEdge`

NewOpalQueryResultEdge instantiates a new OpalQueryResultEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalQueryResultEdgeWithDefaults

`func NewOpalQueryResultEdgeWithDefaults() *OpalQueryResultEdge`

NewOpalQueryResultEdgeWithDefaults instantiates a new OpalQueryResultEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *OpalQueryResultEdge) GetNode() OpalQueryResultNode`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *OpalQueryResultEdge) GetNodeOk() (*OpalQueryResultNode, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *OpalQueryResultEdge) SetNode(v OpalQueryResultNode)`

SetNode sets Node field to given value.


### GetCursor

`func (o *OpalQueryResultEdge) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *OpalQueryResultEdge) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *OpalQueryResultEdge) SetCursor(v string)`

SetCursor sets Cursor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


