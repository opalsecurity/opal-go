# CreateRequestInfoCustomMetadataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | [**RequestTemplateCustomFieldTypeEnum**](RequestTemplateCustomFieldTypeEnum.md) |  | 
**Value** | **string** |  | 

## Methods

### NewCreateRequestInfoCustomMetadataInner

`func NewCreateRequestInfoCustomMetadataInner(name string, type_ RequestTemplateCustomFieldTypeEnum, value string, ) *CreateRequestInfoCustomMetadataInner`

NewCreateRequestInfoCustomMetadataInner instantiates a new CreateRequestInfoCustomMetadataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequestInfoCustomMetadataInnerWithDefaults

`func NewCreateRequestInfoCustomMetadataInnerWithDefaults() *CreateRequestInfoCustomMetadataInner`

NewCreateRequestInfoCustomMetadataInnerWithDefaults instantiates a new CreateRequestInfoCustomMetadataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRequestInfoCustomMetadataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRequestInfoCustomMetadataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRequestInfoCustomMetadataInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateRequestInfoCustomMetadataInner) GetType() RequestTemplateCustomFieldTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRequestInfoCustomMetadataInner) GetTypeOk() (*RequestTemplateCustomFieldTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRequestInfoCustomMetadataInner) SetType(v RequestTemplateCustomFieldTypeEnum)`

SetType sets Type field to given value.


### GetValue

`func (o *CreateRequestInfoCustomMetadataInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateRequestInfoCustomMetadataInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateRequestInfoCustomMetadataInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


