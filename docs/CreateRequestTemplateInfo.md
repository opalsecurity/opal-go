# CreateRequestTemplateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the request template. | 
**CustomFields** | Pointer to [**[]RequestTemplateCustomFieldInput**](RequestTemplateCustomFieldInput.md) | The fields to put on the template, in the order they are shown. | [optional] 

## Methods

### NewCreateRequestTemplateInfo

`func NewCreateRequestTemplateInfo(name string, ) *CreateRequestTemplateInfo`

NewCreateRequestTemplateInfo instantiates a new CreateRequestTemplateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequestTemplateInfoWithDefaults

`func NewCreateRequestTemplateInfoWithDefaults() *CreateRequestTemplateInfo`

NewCreateRequestTemplateInfoWithDefaults instantiates a new CreateRequestTemplateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRequestTemplateInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRequestTemplateInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRequestTemplateInfo) SetName(v string)`

SetName sets Name field to given value.


### GetCustomFields

`func (o *CreateRequestTemplateInfo) GetCustomFields() []RequestTemplateCustomFieldInput`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateRequestTemplateInfo) GetCustomFieldsOk() (*[]RequestTemplateCustomFieldInput, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateRequestTemplateInfo) SetCustomFields(v []RequestTemplateCustomFieldInput)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateRequestTemplateInfo) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


