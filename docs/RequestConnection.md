# RequestConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Edges** | [**[]RequestEdge**](RequestEdge.md) |  | 
**PageInfo** | [**PageInfo**](PageInfo.md) |  | 
**TotalCount** | **int32** | The total number of items available | 

## Methods

### NewRequestConnection

`func NewRequestConnection(edges []RequestEdge, pageInfo PageInfo, totalCount int32, ) *RequestConnection`

NewRequestConnection instantiates a new RequestConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestConnectionWithDefaults

`func NewRequestConnectionWithDefaults() *RequestConnection`

NewRequestConnectionWithDefaults instantiates a new RequestConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdges

`func (o *RequestConnection) GetEdges() []RequestEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *RequestConnection) GetEdgesOk() (*[]RequestEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *RequestConnection) SetEdges(v []RequestEdge)`

SetEdges sets Edges field to given value.


### GetPageInfo

`func (o *RequestConnection) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *RequestConnection) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *RequestConnection) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.


### GetTotalCount

`func (o *RequestConnection) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *RequestConnection) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *RequestConnection) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


