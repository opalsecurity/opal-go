# GetResourceUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ResourceUser**](ResourceUser.md) |  | 
**Cursor** | Pointer to **string** | Pagination cursor for the next page of results | [optional] 
**TotalCount** | Pointer to **int32** | Total number of results | [optional] 

## Methods

### NewGetResourceUser200Response

`func NewGetResourceUser200Response(data []ResourceUser, ) *GetResourceUser200Response`

NewGetResourceUser200Response instantiates a new GetResourceUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResourceUser200ResponseWithDefaults

`func NewGetResourceUser200ResponseWithDefaults() *GetResourceUser200Response`

NewGetResourceUser200ResponseWithDefaults instantiates a new GetResourceUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetResourceUser200Response) GetData() []ResourceUser`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetResourceUser200Response) GetDataOk() (*[]ResourceUser, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetResourceUser200Response) SetData(v []ResourceUser)`

SetData sets Data field to given value.


### GetCursor

`func (o *GetResourceUser200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetResourceUser200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetResourceUser200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetResourceUser200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetResourceUser200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetResourceUser200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetResourceUser200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetResourceUser200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


