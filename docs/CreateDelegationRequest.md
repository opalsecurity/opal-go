# CreateDelegationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegatorUserId** | **string** | The ID of the user delegating their access review requests. | 
**DelegateUserId** | **string** | The ID of the user being delegated to. | 
**StartTime** | **time.Time** | The start time of the delegation. | 
**EndTime** | **time.Time** | The end time of the delegation. | 
**Reason** | **string** | The reason for the delegation. | 

## Methods

### NewCreateDelegationRequest

`func NewCreateDelegationRequest(delegatorUserId string, delegateUserId string, startTime time.Time, endTime time.Time, reason string, ) *CreateDelegationRequest`

NewCreateDelegationRequest instantiates a new CreateDelegationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDelegationRequestWithDefaults

`func NewCreateDelegationRequestWithDefaults() *CreateDelegationRequest`

NewCreateDelegationRequestWithDefaults instantiates a new CreateDelegationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegatorUserId

`func (o *CreateDelegationRequest) GetDelegatorUserId() string`

GetDelegatorUserId returns the DelegatorUserId field if non-nil, zero value otherwise.

### GetDelegatorUserIdOk

`func (o *CreateDelegationRequest) GetDelegatorUserIdOk() (*string, bool)`

GetDelegatorUserIdOk returns a tuple with the DelegatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatorUserId

`func (o *CreateDelegationRequest) SetDelegatorUserId(v string)`

SetDelegatorUserId sets DelegatorUserId field to given value.


### GetDelegateUserId

`func (o *CreateDelegationRequest) GetDelegateUserId() string`

GetDelegateUserId returns the DelegateUserId field if non-nil, zero value otherwise.

### GetDelegateUserIdOk

`func (o *CreateDelegationRequest) GetDelegateUserIdOk() (*string, bool)`

GetDelegateUserIdOk returns a tuple with the DelegateUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateUserId

`func (o *CreateDelegationRequest) SetDelegateUserId(v string)`

SetDelegateUserId sets DelegateUserId field to given value.


### GetStartTime

`func (o *CreateDelegationRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateDelegationRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateDelegationRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *CreateDelegationRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CreateDelegationRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CreateDelegationRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetReason

`func (o *CreateDelegationRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateDelegationRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateDelegationRequest) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


