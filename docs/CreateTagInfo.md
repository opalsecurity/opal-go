# CreateTagInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagKey** | **string** | The key of the tag to create. | 
**TagValue** | Pointer to **string** | The value of the tag to create. | [optional] 

## Methods

### NewCreateTagInfo

`func NewCreateTagInfo(tagKey string, ) *CreateTagInfo`

NewCreateTagInfo instantiates a new CreateTagInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagInfoWithDefaults

`func NewCreateTagInfoWithDefaults() *CreateTagInfo`

NewCreateTagInfoWithDefaults instantiates a new CreateTagInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagKey

`func (o *CreateTagInfo) GetTagKey() string`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *CreateTagInfo) GetTagKeyOk() (*string, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *CreateTagInfo) SetTagKey(v string)`

SetTagKey sets TagKey field to given value.


### GetTagValue

`func (o *CreateTagInfo) GetTagValue() string`

GetTagValue returns the TagValue field if non-nil, zero value otherwise.

### GetTagValueOk

`func (o *CreateTagInfo) GetTagValueOk() (*string, bool)`

GetTagValueOk returns a tuple with the TagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValue

`func (o *CreateTagInfo) SetTagValue(v string)`

SetTagValue sets TagValue field to given value.

### HasTagValue

`func (o *CreateTagInfo) HasTagValue() bool`

HasTagValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


