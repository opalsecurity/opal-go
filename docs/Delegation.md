# Delegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the delegation. | 
**DelegatorUserId** | **string** | The ID of the user delegating their access review requests. | 
**DelegateUserId** | **string** | The ID of the user being delegated to. | 
**StartTime** | **time.Time** | The start time of the delegation. | 
**EndTime** | **time.Time** | The end time of the delegation. | 
**Reason** | **string** | The reason for the delegation. | 
**CreatedAt** | **time.Time** | The creation time of the delegation. | 
**UpdatedAt** | **time.Time** | The last updated time of the delegation. | 

## Methods

### NewDelegation

`func NewDelegation(id string, delegatorUserId string, delegateUserId string, startTime time.Time, endTime time.Time, reason string, createdAt time.Time, updatedAt time.Time, ) *Delegation`

NewDelegation instantiates a new Delegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationWithDefaults

`func NewDelegationWithDefaults() *Delegation`

NewDelegationWithDefaults instantiates a new Delegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Delegation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Delegation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Delegation) SetId(v string)`

SetId sets Id field to given value.


### GetDelegatorUserId

`func (o *Delegation) GetDelegatorUserId() string`

GetDelegatorUserId returns the DelegatorUserId field if non-nil, zero value otherwise.

### GetDelegatorUserIdOk

`func (o *Delegation) GetDelegatorUserIdOk() (*string, bool)`

GetDelegatorUserIdOk returns a tuple with the DelegatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatorUserId

`func (o *Delegation) SetDelegatorUserId(v string)`

SetDelegatorUserId sets DelegatorUserId field to given value.


### GetDelegateUserId

`func (o *Delegation) GetDelegateUserId() string`

GetDelegateUserId returns the DelegateUserId field if non-nil, zero value otherwise.

### GetDelegateUserIdOk

`func (o *Delegation) GetDelegateUserIdOk() (*string, bool)`

GetDelegateUserIdOk returns a tuple with the DelegateUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateUserId

`func (o *Delegation) SetDelegateUserId(v string)`

SetDelegateUserId sets DelegateUserId field to given value.


### GetStartTime

`func (o *Delegation) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Delegation) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Delegation) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *Delegation) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Delegation) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Delegation) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetReason

`func (o *Delegation) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Delegation) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Delegation) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCreatedAt

`func (o *Delegation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Delegation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Delegation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Delegation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Delegation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Delegation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


