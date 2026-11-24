# BatchRequestabilityCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the asset that was checked. | 
**Requestable** | **bool** |  | 
**Reason** | [**RequestabilityReasonEnum**](RequestabilityReasonEnum.md) |  | 
**AccessLevels** | Pointer to [**[]RequestableAccessLevel**](RequestableAccessLevel.md) | Per-access-level requestability. Omitted when reason is NOT_FOUND. | [optional] 

## Methods

### NewBatchRequestabilityCheckResult

`func NewBatchRequestabilityCheckResult(id string, requestable bool, reason RequestabilityReasonEnum, ) *BatchRequestabilityCheckResult`

NewBatchRequestabilityCheckResult instantiates a new BatchRequestabilityCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRequestabilityCheckResultWithDefaults

`func NewBatchRequestabilityCheckResultWithDefaults() *BatchRequestabilityCheckResult`

NewBatchRequestabilityCheckResultWithDefaults instantiates a new BatchRequestabilityCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchRequestabilityCheckResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchRequestabilityCheckResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchRequestabilityCheckResult) SetId(v string)`

SetId sets Id field to given value.


### GetRequestable

`func (o *BatchRequestabilityCheckResult) GetRequestable() bool`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *BatchRequestabilityCheckResult) GetRequestableOk() (*bool, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *BatchRequestabilityCheckResult) SetRequestable(v bool)`

SetRequestable sets Requestable field to given value.


### GetReason

`func (o *BatchRequestabilityCheckResult) GetReason() RequestabilityReasonEnum`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BatchRequestabilityCheckResult) GetReasonOk() (*RequestabilityReasonEnum, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BatchRequestabilityCheckResult) SetReason(v RequestabilityReasonEnum)`

SetReason sets Reason field to given value.


### GetAccessLevels

`func (o *BatchRequestabilityCheckResult) GetAccessLevels() []RequestableAccessLevel`

GetAccessLevels returns the AccessLevels field if non-nil, zero value otherwise.

### GetAccessLevelsOk

`func (o *BatchRequestabilityCheckResult) GetAccessLevelsOk() (*[]RequestableAccessLevel, bool)`

GetAccessLevelsOk returns a tuple with the AccessLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevels

`func (o *BatchRequestabilityCheckResult) SetAccessLevels(v []RequestableAccessLevel)`

SetAccessLevels sets AccessLevels field to given value.

### HasAccessLevels

`func (o *BatchRequestabilityCheckResult) HasAccessLevels() bool`

HasAccessLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


