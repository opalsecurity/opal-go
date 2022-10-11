# CreateMessageChannelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThirdPartyProvider** | [**MessageChannelProviderEnum**](MessageChannelProviderEnum.md) |  | 
**RemoteId** | **string** | The remote ID of the message channel | 

## Methods

### NewCreateMessageChannelInfo

`func NewCreateMessageChannelInfo(thirdPartyProvider MessageChannelProviderEnum, remoteId string, ) *CreateMessageChannelInfo`

NewCreateMessageChannelInfo instantiates a new CreateMessageChannelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMessageChannelInfoWithDefaults

`func NewCreateMessageChannelInfoWithDefaults() *CreateMessageChannelInfo`

NewCreateMessageChannelInfoWithDefaults instantiates a new CreateMessageChannelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThirdPartyProvider

`func (o *CreateMessageChannelInfo) GetThirdPartyProvider() MessageChannelProviderEnum`

GetThirdPartyProvider returns the ThirdPartyProvider field if non-nil, zero value otherwise.

### GetThirdPartyProviderOk

`func (o *CreateMessageChannelInfo) GetThirdPartyProviderOk() (*MessageChannelProviderEnum, bool)`

GetThirdPartyProviderOk returns a tuple with the ThirdPartyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyProvider

`func (o *CreateMessageChannelInfo) SetThirdPartyProvider(v MessageChannelProviderEnum)`

SetThirdPartyProvider sets ThirdPartyProvider field to given value.


### GetRemoteId

`func (o *CreateMessageChannelInfo) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *CreateMessageChannelInfo) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *CreateMessageChannelInfo) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


