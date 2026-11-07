# OpalAccessQueryResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Edges** | [**[]OpalAccessResultEdge**](OpalAccessResultEdge.md) | List of matched access grants. | 
**PageInfo** | [**PageInfo**](PageInfo.md) |  | 

## Methods

### NewOpalAccessQueryResults

`func NewOpalAccessQueryResults(type_ string, edges []OpalAccessResultEdge, pageInfo PageInfo, ) *OpalAccessQueryResults`

NewOpalAccessQueryResults instantiates a new OpalAccessQueryResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalAccessQueryResultsWithDefaults

`func NewOpalAccessQueryResultsWithDefaults() *OpalAccessQueryResults`

NewOpalAccessQueryResultsWithDefaults instantiates a new OpalAccessQueryResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OpalAccessQueryResults) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpalAccessQueryResults) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpalAccessQueryResults) SetType(v string)`

SetType sets Type field to given value.


### GetEdges

`func (o *OpalAccessQueryResults) GetEdges() []OpalAccessResultEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *OpalAccessQueryResults) GetEdgesOk() (*[]OpalAccessResultEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *OpalAccessQueryResults) SetEdges(v []OpalAccessResultEdge)`

SetEdges sets Edges field to given value.


### GetPageInfo

`func (o *OpalAccessQueryResults) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *OpalAccessQueryResults) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *OpalAccessQueryResults) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


