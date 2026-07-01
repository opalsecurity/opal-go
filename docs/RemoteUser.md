# RemoteUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | The ID of the user. | 
**RemoteId** | **string** | The ID of the remote user. | 
**ThirdPartyProvider** | [**ThirdPartyProviderEnum**](ThirdPartyProviderEnum.md) | The third party provider of the remote user. | 

## Methods

### NewRemoteUser

`func NewRemoteUser(userId string, remoteId string, thirdPartyProvider ThirdPartyProviderEnum, ) *RemoteUser`

NewRemoteUser instantiates a new RemoteUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteUserWithDefaults

`func NewRemoteUserWithDefaults() *RemoteUser`

NewRemoteUserWithDefaults instantiates a new RemoteUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *RemoteUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RemoteUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RemoteUser) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetRemoteId

`func (o *RemoteUser) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *RemoteUser) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *RemoteUser) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetThirdPartyProvider

`func (o *RemoteUser) GetThirdPartyProvider() ThirdPartyProviderEnum`

GetThirdPartyProvider returns the ThirdPartyProvider field if non-nil, zero value otherwise.

### GetThirdPartyProviderOk

`func (o *RemoteUser) GetThirdPartyProviderOk() (*ThirdPartyProviderEnum, bool)`

GetThirdPartyProviderOk returns a tuple with the ThirdPartyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyProvider

`func (o *RemoteUser) SetThirdPartyProvider(v ThirdPartyProviderEnum)`

SetThirdPartyProvider sets ThirdPartyProvider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


