# MessageChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageChannelId** | **string** | The ID of the message channel. | 
**ThirdPartyProvider** | Pointer to [**MessageChannelProviderEnum**](MessageChannelProviderEnum.md) |  | [optional] 
**RemoteId** | Pointer to **string** | The remote ID of the message channel | [optional] 
**Name** | Pointer to **string** | The name of the message channel. | [optional] 
**MessageChannelType** | Pointer to [**MessageChannelTypeEnum**](MessageChannelTypeEnum.md) |  | [optional] 
**IsPrivate** | Pointer to **bool** | A bool representing whether or not the message channel is private. | [optional] 

## Methods

### NewMessageChannel

`func NewMessageChannel(messageChannelId string, ) *MessageChannel`

NewMessageChannel instantiates a new MessageChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageChannelWithDefaults

`func NewMessageChannelWithDefaults() *MessageChannel`

NewMessageChannelWithDefaults instantiates a new MessageChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageChannelId

`func (o *MessageChannel) GetMessageChannelId() string`

GetMessageChannelId returns the MessageChannelId field if non-nil, zero value otherwise.

### GetMessageChannelIdOk

`func (o *MessageChannel) GetMessageChannelIdOk() (*string, bool)`

GetMessageChannelIdOk returns a tuple with the MessageChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageChannelId

`func (o *MessageChannel) SetMessageChannelId(v string)`

SetMessageChannelId sets MessageChannelId field to given value.


### GetThirdPartyProvider

`func (o *MessageChannel) GetThirdPartyProvider() MessageChannelProviderEnum`

GetThirdPartyProvider returns the ThirdPartyProvider field if non-nil, zero value otherwise.

### GetThirdPartyProviderOk

`func (o *MessageChannel) GetThirdPartyProviderOk() (*MessageChannelProviderEnum, bool)`

GetThirdPartyProviderOk returns a tuple with the ThirdPartyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyProvider

`func (o *MessageChannel) SetThirdPartyProvider(v MessageChannelProviderEnum)`

SetThirdPartyProvider sets ThirdPartyProvider field to given value.

### HasThirdPartyProvider

`func (o *MessageChannel) HasThirdPartyProvider() bool`

HasThirdPartyProvider returns a boolean if a field has been set.

### GetRemoteId

`func (o *MessageChannel) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *MessageChannel) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *MessageChannel) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *MessageChannel) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetName

`func (o *MessageChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MessageChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MessageChannel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MessageChannel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMessageChannelType

`func (o *MessageChannel) GetMessageChannelType() MessageChannelTypeEnum`

GetMessageChannelType returns the MessageChannelType field if non-nil, zero value otherwise.

### GetMessageChannelTypeOk

`func (o *MessageChannel) GetMessageChannelTypeOk() (*MessageChannelTypeEnum, bool)`

GetMessageChannelTypeOk returns a tuple with the MessageChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageChannelType

`func (o *MessageChannel) SetMessageChannelType(v MessageChannelTypeEnum)`

SetMessageChannelType sets MessageChannelType field to given value.

### HasMessageChannelType

`func (o *MessageChannel) HasMessageChannelType() bool`

HasMessageChannelType returns a boolean if a field has been set.

### GetIsPrivate

`func (o *MessageChannel) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *MessageChannel) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *MessageChannel) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *MessageChannel) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


