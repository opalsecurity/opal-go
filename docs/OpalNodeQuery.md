# OpalNodeQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Query** | Pointer to [**OpalNodeQueryBody**](OpalNodeQueryBody.md) |  | [optional] 
**First** | Pointer to **int32** | Maximum number of results to return. Defaults to 200. | [optional] 
**After** | Pointer to **string** | Cursor from a previous response to fetch the next page of results. | [optional] 

## Methods

### NewOpalNodeQuery

`func NewOpalNodeQuery(type_ string, ) *OpalNodeQuery`

NewOpalNodeQuery instantiates a new OpalNodeQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalNodeQueryWithDefaults

`func NewOpalNodeQueryWithDefaults() *OpalNodeQuery`

NewOpalNodeQueryWithDefaults instantiates a new OpalNodeQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OpalNodeQuery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpalNodeQuery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpalNodeQuery) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *OpalNodeQuery) GetQuery() OpalNodeQueryBody`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *OpalNodeQuery) GetQueryOk() (*OpalNodeQueryBody, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *OpalNodeQuery) SetQuery(v OpalNodeQueryBody)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *OpalNodeQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetFirst

`func (o *OpalNodeQuery) GetFirst() int32`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *OpalNodeQuery) GetFirstOk() (*int32, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *OpalNodeQuery) SetFirst(v int32)`

SetFirst sets First field to given value.

### HasFirst

`func (o *OpalNodeQuery) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetAfter

`func (o *OpalNodeQuery) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *OpalNodeQuery) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *OpalNodeQuery) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *OpalNodeQuery) HasAfter() bool`

HasAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


