# AgentKillSwitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Whether agent tool calls are being denied right now. This is the field to read; it already accounts for expiry. | 
**ActivatedAt** | Pointer to **time.Time** | When the latest activation was turned on. Absent if it never has been. | [optional] 
**ActivatedByUserId** | Pointer to **string** | The admin who turned it on. Absent when that user has since been hard-deleted - the activation record outlives them deliberately. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | When the latest activation stops denying on its own. In the past for an activation that already ran out. | [optional] 
**DeactivatedAt** | Pointer to **time.Time** | When the latest activation was turned off before its expiry, either by an admin or by being superseded. Absent when it ended at &#x60;expires_at&#x60;, or has not ended. | [optional] 
**Reason** | Pointer to **string** | The admin&#39;s stated reason. Not an internal note - it is shown to the developer whose tool call is denied, in their own terminal. | [optional] 

## Methods

### NewAgentKillSwitch

`func NewAgentKillSwitch(active bool, ) *AgentKillSwitch`

NewAgentKillSwitch instantiates a new AgentKillSwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentKillSwitchWithDefaults

`func NewAgentKillSwitchWithDefaults() *AgentKillSwitch`

NewAgentKillSwitchWithDefaults instantiates a new AgentKillSwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AgentKillSwitch) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AgentKillSwitch) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AgentKillSwitch) SetActive(v bool)`

SetActive sets Active field to given value.


### GetActivatedAt

`func (o *AgentKillSwitch) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *AgentKillSwitch) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *AgentKillSwitch) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *AgentKillSwitch) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### GetActivatedByUserId

`func (o *AgentKillSwitch) GetActivatedByUserId() string`

GetActivatedByUserId returns the ActivatedByUserId field if non-nil, zero value otherwise.

### GetActivatedByUserIdOk

`func (o *AgentKillSwitch) GetActivatedByUserIdOk() (*string, bool)`

GetActivatedByUserIdOk returns a tuple with the ActivatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedByUserId

`func (o *AgentKillSwitch) SetActivatedByUserId(v string)`

SetActivatedByUserId sets ActivatedByUserId field to given value.

### HasActivatedByUserId

`func (o *AgentKillSwitch) HasActivatedByUserId() bool`

HasActivatedByUserId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AgentKillSwitch) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AgentKillSwitch) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AgentKillSwitch) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AgentKillSwitch) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetDeactivatedAt

`func (o *AgentKillSwitch) GetDeactivatedAt() time.Time`

GetDeactivatedAt returns the DeactivatedAt field if non-nil, zero value otherwise.

### GetDeactivatedAtOk

`func (o *AgentKillSwitch) GetDeactivatedAtOk() (*time.Time, bool)`

GetDeactivatedAtOk returns a tuple with the DeactivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivatedAt

`func (o *AgentKillSwitch) SetDeactivatedAt(v time.Time)`

SetDeactivatedAt sets DeactivatedAt field to given value.

### HasDeactivatedAt

`func (o *AgentKillSwitch) HasDeactivatedAt() bool`

HasDeactivatedAt returns a boolean if a field has been set.

### GetReason

`func (o *AgentKillSwitch) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AgentKillSwitch) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AgentKillSwitch) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AgentKillSwitch) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


