# RequestableAccessLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevelRemoteId** | **string** | The remote ID of the access level. Empty for assets without access levels. | 
**AccessLevelName** | Pointer to **string** | The human-readable name of the access level. | [optional] 
**Requestable** | **bool** |  | 
**Reason** | Pointer to [**RequestabilityReasonEnum**](RequestabilityReasonEnum.md) | Present when the access level is not requestable. | [optional] 

## Methods

### NewRequestableAccessLevel

`func NewRequestableAccessLevel(accessLevelRemoteId string, requestable bool, ) *RequestableAccessLevel`

NewRequestableAccessLevel instantiates a new RequestableAccessLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestableAccessLevelWithDefaults

`func NewRequestableAccessLevelWithDefaults() *RequestableAccessLevel`

NewRequestableAccessLevelWithDefaults instantiates a new RequestableAccessLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevelRemoteId

`func (o *RequestableAccessLevel) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *RequestableAccessLevel) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *RequestableAccessLevel) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.


### GetAccessLevelName

`func (o *RequestableAccessLevel) GetAccessLevelName() string`

GetAccessLevelName returns the AccessLevelName field if non-nil, zero value otherwise.

### GetAccessLevelNameOk

`func (o *RequestableAccessLevel) GetAccessLevelNameOk() (*string, bool)`

GetAccessLevelNameOk returns a tuple with the AccessLevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelName

`func (o *RequestableAccessLevel) SetAccessLevelName(v string)`

SetAccessLevelName sets AccessLevelName field to given value.

### HasAccessLevelName

`func (o *RequestableAccessLevel) HasAccessLevelName() bool`

HasAccessLevelName returns a boolean if a field has been set.

### GetRequestable

`func (o *RequestableAccessLevel) GetRequestable() bool`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *RequestableAccessLevel) GetRequestableOk() (*bool, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *RequestableAccessLevel) SetRequestable(v bool)`

SetRequestable sets Requestable field to given value.


### GetReason

`func (o *RequestableAccessLevel) GetReason() RequestabilityReasonEnum`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RequestableAccessLevel) GetReasonOk() (*RequestabilityReasonEnum, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RequestableAccessLevel) SetReason(v RequestabilityReasonEnum)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RequestableAccessLevel) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


