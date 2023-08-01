# Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the request. | 
**CreatedAt** | **time.Time** | The date and time the request was created. | 
**UpdatedAt** | **time.Time** | The date and time the request was last updated. | 
**RequesterId** | **string** | The unique identifier of the user who created the request. | 
**TargetUserId** | **string** | The unique identifier of the user who is the target of the request. | 
**Status** | [**RequestStatusEnum**](RequestStatusEnum.md) |  | 
**Reason** | **string** | The reason for the request. | 
**DurationMinutes** | Pointer to **int32** | The duration of the request in minutes. | [optional] 

## Methods

### NewRequest

`func NewRequest(id string, createdAt time.Time, updatedAt time.Time, requesterId string, targetUserId string, status RequestStatusEnum, reason string, ) *Request`

NewRequest instantiates a new Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestWithDefaults

`func NewRequestWithDefaults() *Request`

NewRequestWithDefaults instantiates a new Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Request) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Request) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Request) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Request) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Request) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Request) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Request) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Request) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Request) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRequesterId

`func (o *Request) GetRequesterId() string`

GetRequesterId returns the RequesterId field if non-nil, zero value otherwise.

### GetRequesterIdOk

`func (o *Request) GetRequesterIdOk() (*string, bool)`

GetRequesterIdOk returns a tuple with the RequesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterId

`func (o *Request) SetRequesterId(v string)`

SetRequesterId sets RequesterId field to given value.


### GetTargetUserId

`func (o *Request) GetTargetUserId() string`

GetTargetUserId returns the TargetUserId field if non-nil, zero value otherwise.

### GetTargetUserIdOk

`func (o *Request) GetTargetUserIdOk() (*string, bool)`

GetTargetUserIdOk returns a tuple with the TargetUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUserId

`func (o *Request) SetTargetUserId(v string)`

SetTargetUserId sets TargetUserId field to given value.


### GetStatus

`func (o *Request) GetStatus() RequestStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Request) GetStatusOk() (*RequestStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Request) SetStatus(v RequestStatusEnum)`

SetStatus sets Status field to given value.


### GetReason

`func (o *Request) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Request) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Request) SetReason(v string)`

SetReason sets Reason field to given value.


### GetDurationMinutes

`func (o *Request) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *Request) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *Request) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.

### HasDurationMinutes

`func (o *Request) HasDurationMinutes() bool`

HasDurationMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


