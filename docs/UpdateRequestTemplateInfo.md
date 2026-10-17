# UpdateRequestTemplateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestTemplateId** | **string** | The ID of the request template to update. | 
**Name** | Pointer to **string** | The new name of the request template. | [optional] 
**CustomFields** | Pointer to [**[]RequestTemplateCustomFieldInput**](RequestTemplateCustomFieldInput.md) | The complete set of fields for the template. Any field not included is removed. | [optional] 

## Methods

### NewUpdateRequestTemplateInfo

`func NewUpdateRequestTemplateInfo(requestTemplateId string, ) *UpdateRequestTemplateInfo`

NewUpdateRequestTemplateInfo instantiates a new UpdateRequestTemplateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRequestTemplateInfoWithDefaults

`func NewUpdateRequestTemplateInfoWithDefaults() *UpdateRequestTemplateInfo`

NewUpdateRequestTemplateInfoWithDefaults instantiates a new UpdateRequestTemplateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestTemplateId

`func (o *UpdateRequestTemplateInfo) GetRequestTemplateId() string`

GetRequestTemplateId returns the RequestTemplateId field if non-nil, zero value otherwise.

### GetRequestTemplateIdOk

`func (o *UpdateRequestTemplateInfo) GetRequestTemplateIdOk() (*string, bool)`

GetRequestTemplateIdOk returns a tuple with the RequestTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTemplateId

`func (o *UpdateRequestTemplateInfo) SetRequestTemplateId(v string)`

SetRequestTemplateId sets RequestTemplateId field to given value.


### GetName

`func (o *UpdateRequestTemplateInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRequestTemplateInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRequestTemplateInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateRequestTemplateInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCustomFields

`func (o *UpdateRequestTemplateInfo) GetCustomFields() []RequestTemplateCustomFieldInput`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *UpdateRequestTemplateInfo) GetCustomFieldsOk() (*[]RequestTemplateCustomFieldInput, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *UpdateRequestTemplateInfo) SetCustomFields(v []RequestTemplateCustomFieldInput)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *UpdateRequestTemplateInfo) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


