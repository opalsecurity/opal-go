# RunOpalQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Query** | Pointer to [**OpalNodeQueryBody**](OpalNodeQueryBody.md) |  | [optional] 
**First** | Pointer to **int32** | Maximum number of results to return. Defaults to 200. | [optional] 
**After** | Pointer to **string** | Cursor from a previous response to fetch the next page of results. | [optional] 

## Methods

### NewRunOpalQueryRequest

`func NewRunOpalQueryRequest(type_ string, ) *RunOpalQueryRequest`

NewRunOpalQueryRequest instantiates a new RunOpalQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunOpalQueryRequestWithDefaults

`func NewRunOpalQueryRequestWithDefaults() *RunOpalQueryRequest`

NewRunOpalQueryRequestWithDefaults instantiates a new RunOpalQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RunOpalQueryRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RunOpalQueryRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RunOpalQueryRequest) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *RunOpalQueryRequest) GetQuery() OpalNodeQueryBody`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *RunOpalQueryRequest) GetQueryOk() (*OpalNodeQueryBody, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *RunOpalQueryRequest) SetQuery(v OpalNodeQueryBody)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *RunOpalQueryRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetFirst

`func (o *RunOpalQueryRequest) GetFirst() int32`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *RunOpalQueryRequest) GetFirstOk() (*int32, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *RunOpalQueryRequest) SetFirst(v int32)`

SetFirst sets First field to given value.

### HasFirst

`func (o *RunOpalQueryRequest) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetAfter

`func (o *RunOpalQueryRequest) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *RunOpalQueryRequest) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *RunOpalQueryRequest) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *RunOpalQueryRequest) HasAfter() bool`

HasAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


