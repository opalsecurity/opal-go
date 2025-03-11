# AddResourceNhiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMinutes** | **int32** | The duration for which the resource can be accessed (in minutes). Use 0 to set to indefinite. | 
**AccessLevelRemoteId** | Pointer to **string** | The remote ID of the access level to grant. If omitted, the default access level remote ID value (empty string) is used. | [optional] 

## Methods

### NewAddResourceNhiRequest

`func NewAddResourceNhiRequest(durationMinutes int32, ) *AddResourceNhiRequest`

NewAddResourceNhiRequest instantiates a new AddResourceNhiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddResourceNhiRequestWithDefaults

`func NewAddResourceNhiRequestWithDefaults() *AddResourceNhiRequest`

NewAddResourceNhiRequestWithDefaults instantiates a new AddResourceNhiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMinutes

`func (o *AddResourceNhiRequest) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *AddResourceNhiRequest) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *AddResourceNhiRequest) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.


### GetAccessLevelRemoteId

`func (o *AddResourceNhiRequest) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *AddResourceNhiRequest) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *AddResourceNhiRequest) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *AddResourceNhiRequest) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


