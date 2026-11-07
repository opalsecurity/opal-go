# ResourceConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationTemplateId** | **string** | The ID of the configuration template that the resource is linked to, or &#x60;null&#x60; if the resource is not linked to a configuration template. Only the ID is exposed here; the template&#39;s own fields are managed through the &#x60;/configuration-templates&#x60; endpoints. | 

## Methods

### NewResourceConfigurationTemplate

`func NewResourceConfigurationTemplate(configurationTemplateId string, ) *ResourceConfigurationTemplate`

NewResourceConfigurationTemplate instantiates a new ResourceConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConfigurationTemplateWithDefaults

`func NewResourceConfigurationTemplateWithDefaults() *ResourceConfigurationTemplate`

NewResourceConfigurationTemplateWithDefaults instantiates a new ResourceConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationTemplateId

`func (o *ResourceConfigurationTemplate) GetConfigurationTemplateId() string`

GetConfigurationTemplateId returns the ConfigurationTemplateId field if non-nil, zero value otherwise.

### GetConfigurationTemplateIdOk

`func (o *ResourceConfigurationTemplate) GetConfigurationTemplateIdOk() (*string, bool)`

GetConfigurationTemplateIdOk returns a tuple with the ConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTemplateId

`func (o *ResourceConfigurationTemplate) SetConfigurationTemplateId(v string)`

SetConfigurationTemplateId sets ConfigurationTemplateId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


