# RequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestTemplateId** | **string** | The ID of the request template. | 
**Name** | **string** | The name of the request template. | 
**CustomFields** | Pointer to [**[]RequestTemplateCustomField**](RequestTemplateCustomField.md) | The fields on this template, in the order they are shown. | [optional] 

## Methods

### NewRequestTemplate

`func NewRequestTemplate(requestTemplateId string, name string, ) *RequestTemplate`

NewRequestTemplate instantiates a new RequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTemplateWithDefaults

`func NewRequestTemplateWithDefaults() *RequestTemplate`

NewRequestTemplateWithDefaults instantiates a new RequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestTemplateId

`func (o *RequestTemplate) GetRequestTemplateId() string`

GetRequestTemplateId returns the RequestTemplateId field if non-nil, zero value otherwise.

### GetRequestTemplateIdOk

`func (o *RequestTemplate) GetRequestTemplateIdOk() (*string, bool)`

GetRequestTemplateIdOk returns a tuple with the RequestTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTemplateId

`func (o *RequestTemplate) SetRequestTemplateId(v string)`

SetRequestTemplateId sets RequestTemplateId field to given value.


### GetName

`func (o *RequestTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetCustomFields

`func (o *RequestTemplate) GetCustomFields() []RequestTemplateCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RequestTemplate) GetCustomFieldsOk() (*[]RequestTemplateCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RequestTemplate) SetCustomFields(v []RequestTemplateCustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RequestTemplate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


