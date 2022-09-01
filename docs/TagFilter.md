# TagFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the tag. | 
**Value** | Pointer to **string** | The value of the tag. | [optional] 

## Methods

### NewTagFilter

`func NewTagFilter(key string, ) *TagFilter`

NewTagFilter instantiates a new TagFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagFilterWithDefaults

`func NewTagFilterWithDefaults() *TagFilter`

NewTagFilterWithDefaults instantiates a new TagFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TagFilter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TagFilter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TagFilter) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *TagFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TagFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TagFilter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TagFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


