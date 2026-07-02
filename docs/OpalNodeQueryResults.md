# OpalNodeQueryResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Edges** | [**[]OpalQueryResultEdge**](OpalQueryResultEdge.md) | List of matched entities. | 
**PageInfo** | [**PageInfo**](PageInfo.md) |  | 

## Methods

### NewOpalNodeQueryResults

`func NewOpalNodeQueryResults(type_ string, edges []OpalQueryResultEdge, pageInfo PageInfo, ) *OpalNodeQueryResults`

NewOpalNodeQueryResults instantiates a new OpalNodeQueryResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalNodeQueryResultsWithDefaults

`func NewOpalNodeQueryResultsWithDefaults() *OpalNodeQueryResults`

NewOpalNodeQueryResultsWithDefaults instantiates a new OpalNodeQueryResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OpalNodeQueryResults) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpalNodeQueryResults) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpalNodeQueryResults) SetType(v string)`

SetType sets Type field to given value.


### GetEdges

`func (o *OpalNodeQueryResults) GetEdges() []OpalQueryResultEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *OpalNodeQueryResults) GetEdgesOk() (*[]OpalQueryResultEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *OpalNodeQueryResults) SetEdges(v []OpalQueryResultEdge)`

SetEdges sets Edges field to given value.


### GetPageInfo

`func (o *OpalNodeQueryResults) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *OpalNodeQueryResults) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *OpalNodeQueryResults) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


