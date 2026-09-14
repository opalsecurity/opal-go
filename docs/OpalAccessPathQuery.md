# OpalAccessPathQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Query** | Pointer to [**OpalAccessPathQueryBody**](OpalAccessPathQueryBody.md) |  | [optional] 
**First** | Pointer to **int32** | Maximum number of results to return. Defaults to 200. | [optional] 
**After** | Pointer to **string** | Opaque cursor from a previous ACCESS_PATH response to fetch the next page of results. | [optional] 
**IncludeCount** | Pointer to **bool** | When true, populate totalCount in the response. Defaults to false. | [optional] 

## Methods

### NewOpalAccessPathQuery

`func NewOpalAccessPathQuery(type_ string, ) *OpalAccessPathQuery`

NewOpalAccessPathQuery instantiates a new OpalAccessPathQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalAccessPathQueryWithDefaults

`func NewOpalAccessPathQueryWithDefaults() *OpalAccessPathQuery`

NewOpalAccessPathQueryWithDefaults instantiates a new OpalAccessPathQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OpalAccessPathQuery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpalAccessPathQuery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpalAccessPathQuery) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *OpalAccessPathQuery) GetQuery() OpalAccessPathQueryBody`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *OpalAccessPathQuery) GetQueryOk() (*OpalAccessPathQueryBody, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *OpalAccessPathQuery) SetQuery(v OpalAccessPathQueryBody)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *OpalAccessPathQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetFirst

`func (o *OpalAccessPathQuery) GetFirst() int32`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *OpalAccessPathQuery) GetFirstOk() (*int32, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *OpalAccessPathQuery) SetFirst(v int32)`

SetFirst sets First field to given value.

### HasFirst

`func (o *OpalAccessPathQuery) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetAfter

`func (o *OpalAccessPathQuery) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *OpalAccessPathQuery) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *OpalAccessPathQuery) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *OpalAccessPathQuery) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetIncludeCount

`func (o *OpalAccessPathQuery) GetIncludeCount() bool`

GetIncludeCount returns the IncludeCount field if non-nil, zero value otherwise.

### GetIncludeCountOk

`func (o *OpalAccessPathQuery) GetIncludeCountOk() (*bool, bool)`

GetIncludeCountOk returns a tuple with the IncludeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCount

`func (o *OpalAccessPathQuery) SetIncludeCount(v bool)`

SetIncludeCount sets IncludeCount field to given value.

### HasIncludeCount

`func (o *OpalAccessPathQuery) HasIncludeCount() bool`

HasIncludeCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


