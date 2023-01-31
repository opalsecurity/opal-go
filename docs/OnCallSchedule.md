# OnCallSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnCallScheduleId** | Pointer to **string** | The ID of the on-call schedule. | [optional] 
**ThirdPartyProvider** | Pointer to [**OnCallScheduleProviderEnum**](OnCallScheduleProviderEnum.md) |  | [optional] 
**RemoteId** | Pointer to **string** | The remote ID of the on call schedule | [optional] 
**Name** | Pointer to **string** | The name of the on call schedule. | [optional] 

## Methods

### NewOnCallSchedule

`func NewOnCallSchedule() *OnCallSchedule`

NewOnCallSchedule instantiates a new OnCallSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnCallScheduleWithDefaults

`func NewOnCallScheduleWithDefaults() *OnCallSchedule`

NewOnCallScheduleWithDefaults instantiates a new OnCallSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnCallScheduleId

`func (o *OnCallSchedule) GetOnCallScheduleId() string`

GetOnCallScheduleId returns the OnCallScheduleId field if non-nil, zero value otherwise.

### GetOnCallScheduleIdOk

`func (o *OnCallSchedule) GetOnCallScheduleIdOk() (*string, bool)`

GetOnCallScheduleIdOk returns a tuple with the OnCallScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCallScheduleId

`func (o *OnCallSchedule) SetOnCallScheduleId(v string)`

SetOnCallScheduleId sets OnCallScheduleId field to given value.

### HasOnCallScheduleId

`func (o *OnCallSchedule) HasOnCallScheduleId() bool`

HasOnCallScheduleId returns a boolean if a field has been set.

### GetThirdPartyProvider

`func (o *OnCallSchedule) GetThirdPartyProvider() OnCallScheduleProviderEnum`

GetThirdPartyProvider returns the ThirdPartyProvider field if non-nil, zero value otherwise.

### GetThirdPartyProviderOk

`func (o *OnCallSchedule) GetThirdPartyProviderOk() (*OnCallScheduleProviderEnum, bool)`

GetThirdPartyProviderOk returns a tuple with the ThirdPartyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyProvider

`func (o *OnCallSchedule) SetThirdPartyProvider(v OnCallScheduleProviderEnum)`

SetThirdPartyProvider sets ThirdPartyProvider field to given value.

### HasThirdPartyProvider

`func (o *OnCallSchedule) HasThirdPartyProvider() bool`

HasThirdPartyProvider returns a boolean if a field has been set.

### GetRemoteId

`func (o *OnCallSchedule) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *OnCallSchedule) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *OnCallSchedule) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *OnCallSchedule) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetName

`func (o *OnCallSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OnCallSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OnCallSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OnCallSchedule) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


