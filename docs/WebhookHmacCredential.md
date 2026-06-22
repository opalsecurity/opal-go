# WebhookHmacCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the credential. | 
**Secret** | **string** | The HMAC secret value. | 
**CreatedAt** | **time.Time** | When the credential was created. | 

## Methods

### NewWebhookHmacCredential

`func NewWebhookHmacCredential(id string, secret string, createdAt time.Time, ) *WebhookHmacCredential`

NewWebhookHmacCredential instantiates a new WebhookHmacCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookHmacCredentialWithDefaults

`func NewWebhookHmacCredentialWithDefaults() *WebhookHmacCredential`

NewWebhookHmacCredentialWithDefaults instantiates a new WebhookHmacCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookHmacCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookHmacCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookHmacCredential) SetId(v string)`

SetId sets Id field to given value.


### GetSecret

`func (o *WebhookHmacCredential) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookHmacCredential) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookHmacCredential) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetCreatedAt

`func (o *WebhookHmacCredential) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookHmacCredential) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookHmacCredential) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


