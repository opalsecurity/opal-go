# RequestTemplateCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The label shown to the requester. &#x60;CALLOUT&#x60; fields are display-only, so their name is an internal identifier and is never displayed. | 
**Description** | Pointer to **string** | Helper text shown beneath the field. | [optional] 
**Type** | [**RequestTemplateCustomFieldTypeEnum**](RequestTemplateCustomFieldTypeEnum.md) |  | 
**Required** | Pointer to **bool** | Whether the requester must answer. Always false for &#x60;CALLOUT&#x60; fields, which collect no answer. | [optional] 
**Metadata** | Pointer to [**RequestTemplateCustomFieldMetadata**](RequestTemplateCustomFieldMetadata.md) |  | [optional] 

## Methods

### NewRequestTemplateCustomField

`func NewRequestTemplateCustomField(name string, type_ RequestTemplateCustomFieldTypeEnum, ) *RequestTemplateCustomField`

NewRequestTemplateCustomField instantiates a new RequestTemplateCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTemplateCustomFieldWithDefaults

`func NewRequestTemplateCustomFieldWithDefaults() *RequestTemplateCustomField`

NewRequestTemplateCustomFieldWithDefaults instantiates a new RequestTemplateCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RequestTemplateCustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestTemplateCustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestTemplateCustomField) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RequestTemplateCustomField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestTemplateCustomField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestTemplateCustomField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestTemplateCustomField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *RequestTemplateCustomField) GetType() RequestTemplateCustomFieldTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestTemplateCustomField) GetTypeOk() (*RequestTemplateCustomFieldTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestTemplateCustomField) SetType(v RequestTemplateCustomFieldTypeEnum)`

SetType sets Type field to given value.


### GetRequired

`func (o *RequestTemplateCustomField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *RequestTemplateCustomField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *RequestTemplateCustomField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *RequestTemplateCustomField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetMetadata

`func (o *RequestTemplateCustomField) GetMetadata() RequestTemplateCustomFieldMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RequestTemplateCustomField) GetMetadataOk() (*RequestTemplateCustomFieldMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RequestTemplateCustomField) SetMetadata(v RequestTemplateCustomFieldMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RequestTemplateCustomField) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


