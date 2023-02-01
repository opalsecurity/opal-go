# CreateOnCallScheduleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThirdPartyProvider** | [**OnCallScheduleProviderEnum**](OnCallScheduleProviderEnum.md) |  | 
**RemoteId** | **string** | The remote ID of the on call schedule | 

## Methods

### NewCreateOnCallScheduleInfo

`func NewCreateOnCallScheduleInfo(thirdPartyProvider OnCallScheduleProviderEnum, remoteId string, ) *CreateOnCallScheduleInfo`

NewCreateOnCallScheduleInfo instantiates a new CreateOnCallScheduleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOnCallScheduleInfoWithDefaults

`func NewCreateOnCallScheduleInfoWithDefaults() *CreateOnCallScheduleInfo`

NewCreateOnCallScheduleInfoWithDefaults instantiates a new CreateOnCallScheduleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThirdPartyProvider

`func (o *CreateOnCallScheduleInfo) GetThirdPartyProvider() OnCallScheduleProviderEnum`

GetThirdPartyProvider returns the ThirdPartyProvider field if non-nil, zero value otherwise.

### GetThirdPartyProviderOk

`func (o *CreateOnCallScheduleInfo) GetThirdPartyProviderOk() (*OnCallScheduleProviderEnum, bool)`

GetThirdPartyProviderOk returns a tuple with the ThirdPartyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyProvider

`func (o *CreateOnCallScheduleInfo) SetThirdPartyProvider(v OnCallScheduleProviderEnum)`

SetThirdPartyProvider sets ThirdPartyProvider field to given value.


### GetRemoteId

`func (o *CreateOnCallScheduleInfo) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *CreateOnCallScheduleInfo) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *CreateOnCallScheduleInfo) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


