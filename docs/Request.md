# Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the request. | 
**CreatedAt** | **time.Time** | The date and time the request was created. | 
**UpdatedAt** | **time.Time** | The date and time the request was last updated. | 
**RequesterId** | **string** | The unique identifier of the user who created the request. | 
**TargetUserId** | Pointer to **string** | The unique identifier of the user who is the target of the request. | [optional] 
**TargetGroupId** | Pointer to **string** | The unique identifier of the group who is the target of the request. | [optional] 
**Status** | [**RequestStatusEnum**](RequestStatusEnum.md) | The status of the request. | 
**Reason** | **string** | The reason for the request. | 
**DurationMinutes** | Pointer to **int32** | The duration of the request in minutes. | [optional] 
**RequestedItemsList** | Pointer to [**[]RequestedItem**](RequestedItem.md) | The list of targets for the request. | [optional] 
**CustomFieldsResponses** | Pointer to [**[]RequestCustomFieldResponse**](RequestCustomFieldResponse.md) | The responses given to the custom fields associated to the request | [optional] 

## Methods

### NewRequest

`func NewRequest(id string, createdAt time.Time, updatedAt time.Time, requesterId string, status RequestStatusEnum, reason string, ) *Request`

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

### HasTargetUserId

`func (o *Request) HasTargetUserId() bool`

HasTargetUserId returns a boolean if a field has been set.

### GetTargetGroupId

`func (o *Request) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *Request) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *Request) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *Request) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

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

### GetRequestedItemsList

`func (o *Request) GetRequestedItemsList() []RequestedItem`

GetRequestedItemsList returns the RequestedItemsList field if non-nil, zero value otherwise.

### GetRequestedItemsListOk

`func (o *Request) GetRequestedItemsListOk() (*[]RequestedItem, bool)`

GetRequestedItemsListOk returns a tuple with the RequestedItemsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedItemsList

`func (o *Request) SetRequestedItemsList(v []RequestedItem)`

SetRequestedItemsList sets RequestedItemsList field to given value.

### HasRequestedItemsList

`func (o *Request) HasRequestedItemsList() bool`

HasRequestedItemsList returns a boolean if a field has been set.

### GetCustomFieldsResponses

`func (o *Request) GetCustomFieldsResponses() []RequestCustomFieldResponse`

GetCustomFieldsResponses returns the CustomFieldsResponses field if non-nil, zero value otherwise.

### GetCustomFieldsResponsesOk

`func (o *Request) GetCustomFieldsResponsesOk() (*[]RequestCustomFieldResponse, bool)`

GetCustomFieldsResponsesOk returns a tuple with the CustomFieldsResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldsResponses

`func (o *Request) SetCustomFieldsResponses(v []RequestCustomFieldResponse)`

SetCustomFieldsResponses sets CustomFieldsResponses field to given value.

### HasCustomFieldsResponses

`func (o *Request) HasCustomFieldsResponses() bool`

HasCustomFieldsResponses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


