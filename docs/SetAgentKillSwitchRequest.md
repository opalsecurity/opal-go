# SetAgentKillSwitchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | True denies agent tool calls for this organization; false lifts the denial. | 
**DurationMinutes** | Pointer to **int32** | How long the denial lasts, from now. Ignored when &#x60;active&#x60; is false. Capped server-side - a kill switch with no end is how an organization discovers months later that its engineers have been quietly working around it. | [optional] 
**Reason** | Pointer to **string** | Why. Shown verbatim to the developer whose tool call is denied, so write it for them rather than for the audit log - it is the only explanation they get. | [optional] 

## Methods

### NewSetAgentKillSwitchRequest

`func NewSetAgentKillSwitchRequest(active bool, ) *SetAgentKillSwitchRequest`

NewSetAgentKillSwitchRequest instantiates a new SetAgentKillSwitchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAgentKillSwitchRequestWithDefaults

`func NewSetAgentKillSwitchRequestWithDefaults() *SetAgentKillSwitchRequest`

NewSetAgentKillSwitchRequestWithDefaults instantiates a new SetAgentKillSwitchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *SetAgentKillSwitchRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SetAgentKillSwitchRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SetAgentKillSwitchRequest) SetActive(v bool)`

SetActive sets Active field to given value.


### GetDurationMinutes

`func (o *SetAgentKillSwitchRequest) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *SetAgentKillSwitchRequest) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *SetAgentKillSwitchRequest) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.

### HasDurationMinutes

`func (o *SetAgentKillSwitchRequest) HasDurationMinutes() bool`

HasDurationMinutes returns a boolean if a field has been set.

### GetReason

`func (o *SetAgentKillSwitchRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SetAgentKillSwitchRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SetAgentKillSwitchRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SetAgentKillSwitchRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


