# OpalAccessQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Query** | Pointer to [**OpalAccessPathQueryBody**](OpalAccessPathQueryBody.md) |  | [optional] 
**First** | Pointer to **int32** | Maximum number of results to return. Defaults to 200. | [optional] 
**After** | Pointer to **string** | Opaque cursor from a previous ACCESS response to fetch the next page of results. | [optional] 

## Methods

### NewOpalAccessQuery

`func NewOpalAccessQuery(type_ string, ) *OpalAccessQuery`

NewOpalAccessQuery instantiates a new OpalAccessQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalAccessQueryWithDefaults

`func NewOpalAccessQueryWithDefaults() *OpalAccessQuery`

NewOpalAccessQueryWithDefaults instantiates a new OpalAccessQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OpalAccessQuery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpalAccessQuery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpalAccessQuery) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *OpalAccessQuery) GetQuery() OpalAccessPathQueryBody`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *OpalAccessQuery) GetQueryOk() (*OpalAccessPathQueryBody, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *OpalAccessQuery) SetQuery(v OpalAccessPathQueryBody)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *OpalAccessQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetFirst

`func (o *OpalAccessQuery) GetFirst() int32`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *OpalAccessQuery) GetFirstOk() (*int32, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *OpalAccessQuery) SetFirst(v int32)`

SetFirst sets First field to given value.

### HasFirst

`func (o *OpalAccessQuery) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetAfter

`func (o *OpalAccessQuery) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *OpalAccessQuery) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *OpalAccessQuery) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *OpalAccessQuery) HasAfter() bool`

HasAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


