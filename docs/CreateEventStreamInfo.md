# CreateEventStreamInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name for the event stream. | 
**ConnectionType** | [**EventStreamConnectionTypeEnum**](EventStreamConnectionTypeEnum.md) |  | 
**WebhookUrl** | Pointer to **string** | The webhook URL. Required when connection_type is WEBHOOK. | [optional] 
**Credentials** | Pointer to [**WebhookCredentials**](WebhookCredentials.md) |  | [optional] 

## Methods

### NewCreateEventStreamInfo

`func NewCreateEventStreamInfo(name string, connectionType EventStreamConnectionTypeEnum, ) *CreateEventStreamInfo`

NewCreateEventStreamInfo instantiates a new CreateEventStreamInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEventStreamInfoWithDefaults

`func NewCreateEventStreamInfoWithDefaults() *CreateEventStreamInfo`

NewCreateEventStreamInfoWithDefaults instantiates a new CreateEventStreamInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEventStreamInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEventStreamInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEventStreamInfo) SetName(v string)`

SetName sets Name field to given value.


### GetConnectionType

`func (o *CreateEventStreamInfo) GetConnectionType() EventStreamConnectionTypeEnum`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateEventStreamInfo) GetConnectionTypeOk() (*EventStreamConnectionTypeEnum, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CreateEventStreamInfo) SetConnectionType(v EventStreamConnectionTypeEnum)`

SetConnectionType sets ConnectionType field to given value.


### GetWebhookUrl

`func (o *CreateEventStreamInfo) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateEventStreamInfo) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateEventStreamInfo) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateEventStreamInfo) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetCredentials

`func (o *CreateEventStreamInfo) GetCredentials() WebhookCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CreateEventStreamInfo) GetCredentialsOk() (*WebhookCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CreateEventStreamInfo) SetCredentials(v WebhookCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CreateEventStreamInfo) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


