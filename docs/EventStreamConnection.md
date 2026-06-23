# EventStreamConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the connection. | 
**ConnectionType** | [**EventStreamConnectionTypeEnum**](EventStreamConnectionTypeEnum.md) |  | 
**Enabled** | **bool** | Whether the connection is enabled. | 
**WebhookUrl** | Pointer to **string** | The webhook URL, present when connection_type is WEBHOOK. | [optional] 
**Credentials** | Pointer to [**WebhookCredentials**](WebhookCredentials.md) |  | [optional] 

## Methods

### NewEventStreamConnection

`func NewEventStreamConnection(name string, connectionType EventStreamConnectionTypeEnum, enabled bool, ) *EventStreamConnection`

NewEventStreamConnection instantiates a new EventStreamConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventStreamConnectionWithDefaults

`func NewEventStreamConnectionWithDefaults() *EventStreamConnection`

NewEventStreamConnectionWithDefaults instantiates a new EventStreamConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventStreamConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventStreamConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventStreamConnection) SetName(v string)`

SetName sets Name field to given value.


### GetConnectionType

`func (o *EventStreamConnection) GetConnectionType() EventStreamConnectionTypeEnum`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *EventStreamConnection) GetConnectionTypeOk() (*EventStreamConnectionTypeEnum, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *EventStreamConnection) SetConnectionType(v EventStreamConnectionTypeEnum)`

SetConnectionType sets ConnectionType field to given value.


### GetEnabled

`func (o *EventStreamConnection) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EventStreamConnection) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EventStreamConnection) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetWebhookUrl

`func (o *EventStreamConnection) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *EventStreamConnection) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *EventStreamConnection) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *EventStreamConnection) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetCredentials

`func (o *EventStreamConnection) GetCredentials() WebhookCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *EventStreamConnection) GetCredentialsOk() (*WebhookCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *EventStreamConnection) SetCredentials(v WebhookCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *EventStreamConnection) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


