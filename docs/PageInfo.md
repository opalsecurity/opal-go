# PageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | **bool** | Whether there are more items after the end cursor | 
**EndCursor** | **string** | The cursor to continue pagination forwards | 
**HasPreviousPage** | **bool** | Whether there are more items before the start cursor | 
**StartCursor** | **string** | The cursor to continue pagination backwards | 

## Methods

### NewPageInfo

`func NewPageInfo(hasNextPage bool, endCursor string, hasPreviousPage bool, startCursor string, ) *PageInfo`

NewPageInfo instantiates a new PageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageInfoWithDefaults

`func NewPageInfoWithDefaults() *PageInfo`

NewPageInfoWithDefaults instantiates a new PageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *PageInfo) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *PageInfo) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *PageInfo) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetEndCursor

`func (o *PageInfo) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *PageInfo) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *PageInfo) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.


### GetHasPreviousPage

`func (o *PageInfo) GetHasPreviousPage() bool`

GetHasPreviousPage returns the HasPreviousPage field if non-nil, zero value otherwise.

### GetHasPreviousPageOk

`func (o *PageInfo) GetHasPreviousPageOk() (*bool, bool)`

GetHasPreviousPageOk returns a tuple with the HasPreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreviousPage

`func (o *PageInfo) SetHasPreviousPage(v bool)`

SetHasPreviousPage sets HasPreviousPage field to given value.


### GetStartCursor

`func (o *PageInfo) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *PageInfo) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *PageInfo) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


