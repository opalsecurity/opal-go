# OpalQueryResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Edges** | [**[]OpalAccessPathResultEdge**](OpalAccessPathResultEdge.md) | List of matched access paths. | 
**PageInfo** | [**PageInfo**](PageInfo.md) |  | 
**TotalCount** | Pointer to **int32** | Exact total number of matching paths when includeCount was true on the request; otherwise null. | [optional] 

## Methods

### NewOpalQueryResults

`func NewOpalQueryResults(type_ string, edges []OpalAccessPathResultEdge, pageInfo PageInfo, ) *OpalQueryResults`

NewOpalQueryResults instantiates a new OpalQueryResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalQueryResultsWithDefaults

`func NewOpalQueryResultsWithDefaults() *OpalQueryResults`

NewOpalQueryResultsWithDefaults instantiates a new OpalQueryResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OpalQueryResults) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpalQueryResults) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpalQueryResults) SetType(v string)`

SetType sets Type field to given value.


### GetEdges

`func (o *OpalQueryResults) GetEdges() []OpalAccessPathResultEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *OpalQueryResults) GetEdgesOk() (*[]OpalAccessPathResultEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *OpalQueryResults) SetEdges(v []OpalAccessPathResultEdge)`

SetEdges sets Edges field to given value.


### GetPageInfo

`func (o *OpalQueryResults) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *OpalQueryResults) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *OpalQueryResults) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.


### GetTotalCount

`func (o *OpalQueryResults) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OpalQueryResults) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OpalQueryResults) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OpalQueryResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


