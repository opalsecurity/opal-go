# UpdateEventStreamInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Updated name for the event stream. | [optional] 
**Enabled** | Pointer to **bool** | Whether the event stream should be enabled. | [optional] 
**WebhookUrl** | Pointer to **string** | Updated webhook URL. | [optional] 
**Credentials** | Pointer to [**WebhookCredentials**](WebhookCredentials.md) |  | [optional] 

## Methods

### NewUpdateEventStreamInfo

`func NewUpdateEventStreamInfo() *UpdateEventStreamInfo`

NewUpdateEventStreamInfo instantiates a new UpdateEventStreamInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEventStreamInfoWithDefaults

`func NewUpdateEventStreamInfoWithDefaults() *UpdateEventStreamInfo`

NewUpdateEventStreamInfoWithDefaults instantiates a new UpdateEventStreamInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateEventStreamInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEventStreamInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEventStreamInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateEventStreamInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateEventStreamInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateEventStreamInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateEventStreamInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateEventStreamInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *UpdateEventStreamInfo) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *UpdateEventStreamInfo) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *UpdateEventStreamInfo) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *UpdateEventStreamInfo) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetCredentials

`func (o *UpdateEventStreamInfo) GetCredentials() WebhookCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *UpdateEventStreamInfo) GetCredentialsOk() (*WebhookCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *UpdateEventStreamInfo) SetCredentials(v WebhookCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *UpdateEventStreamInfo) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


