# Owner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | **string** | The ID of the owner. | 
**Name** | Pointer to **string** | The name of the owner. | [optional] 
**Description** | Pointer to **string** | A description of the owner. | [optional] 
**AccessRequestEscalationPeriod** | Pointer to **int32** | The amount of time (in minutes) before the next reviewer is notified. Use 0 to remove escalation policy. | [optional] 

## Methods

### NewOwner

`func NewOwner(ownerId string, ) *Owner`

NewOwner instantiates a new Owner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerWithDefaults

`func NewOwnerWithDefaults() *Owner`

NewOwnerWithDefaults instantiates a new Owner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *Owner) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Owner) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Owner) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetName

`func (o *Owner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Owner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Owner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Owner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Owner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Owner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Owner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Owner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccessRequestEscalationPeriod

`func (o *Owner) GetAccessRequestEscalationPeriod() int32`

GetAccessRequestEscalationPeriod returns the AccessRequestEscalationPeriod field if non-nil, zero value otherwise.

### GetAccessRequestEscalationPeriodOk

`func (o *Owner) GetAccessRequestEscalationPeriodOk() (*int32, bool)`

GetAccessRequestEscalationPeriodOk returns a tuple with the AccessRequestEscalationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestEscalationPeriod

`func (o *Owner) SetAccessRequestEscalationPeriod(v int32)`

SetAccessRequestEscalationPeriod sets AccessRequestEscalationPeriod field to given value.

### HasAccessRequestEscalationPeriod

`func (o *Owner) HasAccessRequestEscalationPeriod() bool`

HasAccessRequestEscalationPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


