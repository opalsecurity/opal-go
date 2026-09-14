# OpalAccessPathQueryResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Edges** | [**[]OpalAccessPathResultEdge**](OpalAccessPathResultEdge.md) | List of matched access paths. | 
**PageInfo** | [**PageInfo**](PageInfo.md) |  | 
**TotalCount** | Pointer to **int32** | Exact total number of matching paths when includeCount was true on the request; otherwise null. | [optional] 

## Methods

### NewOpalAccessPathQueryResults

`func NewOpalAccessPathQueryResults(type_ string, edges []OpalAccessPathResultEdge, pageInfo PageInfo, ) *OpalAccessPathQueryResults`

NewOpalAccessPathQueryResults instantiates a new OpalAccessPathQueryResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalAccessPathQueryResultsWithDefaults

`func NewOpalAccessPathQueryResultsWithDefaults() *OpalAccessPathQueryResults`

NewOpalAccessPathQueryResultsWithDefaults instantiates a new OpalAccessPathQueryResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OpalAccessPathQueryResults) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpalAccessPathQueryResults) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpalAccessPathQueryResults) SetType(v string)`

SetType sets Type field to given value.


### GetEdges

`func (o *OpalAccessPathQueryResults) GetEdges() []OpalAccessPathResultEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *OpalAccessPathQueryResults) GetEdgesOk() (*[]OpalAccessPathResultEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *OpalAccessPathQueryResults) SetEdges(v []OpalAccessPathResultEdge)`

SetEdges sets Edges field to given value.


### GetPageInfo

`func (o *OpalAccessPathQueryResults) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *OpalAccessPathQueryResults) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *OpalAccessPathQueryResults) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.


### GetTotalCount

`func (o *OpalAccessPathQueryResults) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OpalAccessPathQueryResults) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OpalAccessPathQueryResults) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OpalAccessPathQueryResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


