# WebhookApiKeyCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the credential. | 
**Name** | **string** | The name of the API key. | 
**Value** | **string** | The value of the API key. | 
**Location** | [**WebhookApiKeyLocationEnum**](WebhookApiKeyLocationEnum.md) |  | 

## Methods

### NewWebhookApiKeyCredential

`func NewWebhookApiKeyCredential(id string, name string, value string, location WebhookApiKeyLocationEnum, ) *WebhookApiKeyCredential`

NewWebhookApiKeyCredential instantiates a new WebhookApiKeyCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookApiKeyCredentialWithDefaults

`func NewWebhookApiKeyCredentialWithDefaults() *WebhookApiKeyCredential`

NewWebhookApiKeyCredentialWithDefaults instantiates a new WebhookApiKeyCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookApiKeyCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookApiKeyCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookApiKeyCredential) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WebhookApiKeyCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookApiKeyCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookApiKeyCredential) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *WebhookApiKeyCredential) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WebhookApiKeyCredential) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WebhookApiKeyCredential) SetValue(v string)`

SetValue sets Value field to given value.


### GetLocation

`func (o *WebhookApiKeyCredential) GetLocation() WebhookApiKeyLocationEnum`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WebhookApiKeyCredential) GetLocationOk() (*WebhookApiKeyLocationEnum, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WebhookApiKeyCredential) SetLocation(v WebhookApiKeyLocationEnum)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


