# RequestTemplateCustomFieldInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The label shown to the requester. Required for every type except &#x60;CALLOUT&#x60;, which is display-only -- its name is an internal identifier, never displayed, and is generated if omitted. | [optional] 
**Description** | Pointer to **string** | Helper text shown beneath the field. | [optional] 
**Type** | [**RequestTemplateCustomFieldTypeEnum**](RequestTemplateCustomFieldTypeEnum.md) |  | 
**Required** | Pointer to **bool** | Whether the requester must answer. Ignored for &#x60;CALLOUT&#x60; fields, which collect no answer. | [optional] 
**Metadata** | Pointer to [**RequestTemplateCustomFieldMetadata**](RequestTemplateCustomFieldMetadata.md) |  | [optional] 

## Methods

### NewRequestTemplateCustomFieldInput

`func NewRequestTemplateCustomFieldInput(type_ RequestTemplateCustomFieldTypeEnum, ) *RequestTemplateCustomFieldInput`

NewRequestTemplateCustomFieldInput instantiates a new RequestTemplateCustomFieldInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTemplateCustomFieldInputWithDefaults

`func NewRequestTemplateCustomFieldInputWithDefaults() *RequestTemplateCustomFieldInput`

NewRequestTemplateCustomFieldInputWithDefaults instantiates a new RequestTemplateCustomFieldInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RequestTemplateCustomFieldInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestTemplateCustomFieldInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestTemplateCustomFieldInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestTemplateCustomFieldInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RequestTemplateCustomFieldInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestTemplateCustomFieldInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestTemplateCustomFieldInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestTemplateCustomFieldInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *RequestTemplateCustomFieldInput) GetType() RequestTemplateCustomFieldTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestTemplateCustomFieldInput) GetTypeOk() (*RequestTemplateCustomFieldTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestTemplateCustomFieldInput) SetType(v RequestTemplateCustomFieldTypeEnum)`

SetType sets Type field to given value.


### GetRequired

`func (o *RequestTemplateCustomFieldInput) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *RequestTemplateCustomFieldInput) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *RequestTemplateCustomFieldInput) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *RequestTemplateCustomFieldInput) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetMetadata

`func (o *RequestTemplateCustomFieldInput) GetMetadata() RequestTemplateCustomFieldMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RequestTemplateCustomFieldInput) GetMetadataOk() (*RequestTemplateCustomFieldMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RequestTemplateCustomFieldInput) SetMetadata(v RequestTemplateCustomFieldMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RequestTemplateCustomFieldInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


