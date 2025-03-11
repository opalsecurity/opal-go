# AddGroupResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevelRemoteId** | Pointer to **string** | The remote ID of the access level to grant to this user. If omitted, the default access level remote ID value (empty string) is used. | [optional] 
**DurationMinutes** | Pointer to **int32** | The duration for which the resource can be accessed (in minutes). Use 0 to set to indefinite. | [optional] 

## Methods

### NewAddGroupResourceRequest

`func NewAddGroupResourceRequest() *AddGroupResourceRequest`

NewAddGroupResourceRequest instantiates a new AddGroupResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroupResourceRequestWithDefaults

`func NewAddGroupResourceRequestWithDefaults() *AddGroupResourceRequest`

NewAddGroupResourceRequestWithDefaults instantiates a new AddGroupResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevelRemoteId

`func (o *AddGroupResourceRequest) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *AddGroupResourceRequest) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *AddGroupResourceRequest) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *AddGroupResourceRequest) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.

### GetDurationMinutes

`func (o *AddGroupResourceRequest) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *AddGroupResourceRequest) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *AddGroupResourceRequest) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.

### HasDurationMinutes

`func (o *AddGroupResourceRequest) HasDurationMinutes() bool`

HasDurationMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


