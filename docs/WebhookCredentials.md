# WebhookCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | [**WebhookAuthTypeEnum**](WebhookAuthTypeEnum.md) |  | 
**ApiKeyCredentials** | Pointer to [**[]WebhookApiKeyCredential**](WebhookApiKeyCredential.md) | API key credentials, present when auth_type is API_KEY. | [optional] 
**HmacCredential1** | Pointer to [**WebhookHmacCredential**](WebhookHmacCredential.md) | Primary HMAC credential, present when auth_type is HMAC. | [optional] 
**HmacCredential2** | Pointer to [**WebhookHmacCredential**](WebhookHmacCredential.md) | Secondary HMAC credential for rotation, present when auth_type is HMAC. | [optional] 

## Methods

### NewWebhookCredentials

`func NewWebhookCredentials(authType WebhookAuthTypeEnum, ) *WebhookCredentials`

NewWebhookCredentials instantiates a new WebhookCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookCredentialsWithDefaults

`func NewWebhookCredentialsWithDefaults() *WebhookCredentials`

NewWebhookCredentialsWithDefaults instantiates a new WebhookCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *WebhookCredentials) GetAuthType() WebhookAuthTypeEnum`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *WebhookCredentials) GetAuthTypeOk() (*WebhookAuthTypeEnum, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *WebhookCredentials) SetAuthType(v WebhookAuthTypeEnum)`

SetAuthType sets AuthType field to given value.


### GetApiKeyCredentials

`func (o *WebhookCredentials) GetApiKeyCredentials() []WebhookApiKeyCredential`

GetApiKeyCredentials returns the ApiKeyCredentials field if non-nil, zero value otherwise.

### GetApiKeyCredentialsOk

`func (o *WebhookCredentials) GetApiKeyCredentialsOk() (*[]WebhookApiKeyCredential, bool)`

GetApiKeyCredentialsOk returns a tuple with the ApiKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyCredentials

`func (o *WebhookCredentials) SetApiKeyCredentials(v []WebhookApiKeyCredential)`

SetApiKeyCredentials sets ApiKeyCredentials field to given value.

### HasApiKeyCredentials

`func (o *WebhookCredentials) HasApiKeyCredentials() bool`

HasApiKeyCredentials returns a boolean if a field has been set.

### GetHmacCredential1

`func (o *WebhookCredentials) GetHmacCredential1() WebhookHmacCredential`

GetHmacCredential1 returns the HmacCredential1 field if non-nil, zero value otherwise.

### GetHmacCredential1Ok

`func (o *WebhookCredentials) GetHmacCredential1Ok() (*WebhookHmacCredential, bool)`

GetHmacCredential1Ok returns a tuple with the HmacCredential1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacCredential1

`func (o *WebhookCredentials) SetHmacCredential1(v WebhookHmacCredential)`

SetHmacCredential1 sets HmacCredential1 field to given value.

### HasHmacCredential1

`func (o *WebhookCredentials) HasHmacCredential1() bool`

HasHmacCredential1 returns a boolean if a field has been set.

### GetHmacCredential2

`func (o *WebhookCredentials) GetHmacCredential2() WebhookHmacCredential`

GetHmacCredential2 returns the HmacCredential2 field if non-nil, zero value otherwise.

### GetHmacCredential2Ok

`func (o *WebhookCredentials) GetHmacCredential2Ok() (*WebhookHmacCredential, bool)`

GetHmacCredential2Ok returns a tuple with the HmacCredential2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacCredential2

`func (o *WebhookCredentials) SetHmacCredential2(v WebhookHmacCredential)`

SetHmacCredential2 sets HmacCredential2 field to given value.

### HasHmacCredential2

`func (o *WebhookCredentials) HasHmacCredential2() bool`

HasHmacCredential2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


