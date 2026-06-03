# EntityTagFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The tag key to filter by. | 
**Value** | Pointer to **string** | The tag value to filter by. If omitted, matches any value for the given key. | [optional] 
**ConnectionId** | Pointer to **string** | If specified, filters by tags associated with this connection. | [optional] 

## Methods

### NewEntityTagFilter

`func NewEntityTagFilter(key string, ) *EntityTagFilter`

NewEntityTagFilter instantiates a new EntityTagFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityTagFilterWithDefaults

`func NewEntityTagFilterWithDefaults() *EntityTagFilter`

NewEntityTagFilterWithDefaults instantiates a new EntityTagFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *EntityTagFilter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EntityTagFilter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EntityTagFilter) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *EntityTagFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntityTagFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntityTagFilter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EntityTagFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetConnectionId

`func (o *EntityTagFilter) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *EntityTagFilter) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *EntityTagFilter) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *EntityTagFilter) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


