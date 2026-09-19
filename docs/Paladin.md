# Paladin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaladinId** | **string** | The ID of the Paladin. Use this value as a reviewer in a request configuration&#39;s service_user_ids. | 
**Name** | **string** | The name of the Paladin. | 
**OwnerId** | **string** | The ID of the owner of the Paladin. | 
**MonitorMode** | **bool** | When true, the Paladin reasons about requests but takes no action. Shown as \&quot;Monitor mode\&quot; in the UI. | 
**AdminViewOnly** | **bool** | When true, the Paladin&#39;s recommendations are visible only to admins. Only meaningful when monitor_mode is true. | 
**EnabledConnectors** | [**[]PaladinConnector**](PaladinConnector.md) | The connectors the Paladin is allowed to use. | 
**Instructions** | **string** | The free-form instructions that guide the Paladin&#39;s decisions. | 

## Methods

### NewPaladin

`func NewPaladin(paladinId string, name string, ownerId string, monitorMode bool, adminViewOnly bool, enabledConnectors []PaladinConnector, instructions string, ) *Paladin`

NewPaladin instantiates a new Paladin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaladinWithDefaults

`func NewPaladinWithDefaults() *Paladin`

NewPaladinWithDefaults instantiates a new Paladin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaladinId

`func (o *Paladin) GetPaladinId() string`

GetPaladinId returns the PaladinId field if non-nil, zero value otherwise.

### GetPaladinIdOk

`func (o *Paladin) GetPaladinIdOk() (*string, bool)`

GetPaladinIdOk returns a tuple with the PaladinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaladinId

`func (o *Paladin) SetPaladinId(v string)`

SetPaladinId sets PaladinId field to given value.


### GetName

`func (o *Paladin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Paladin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Paladin) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *Paladin) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Paladin) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Paladin) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetMonitorMode

`func (o *Paladin) GetMonitorMode() bool`

GetMonitorMode returns the MonitorMode field if non-nil, zero value otherwise.

### GetMonitorModeOk

`func (o *Paladin) GetMonitorModeOk() (*bool, bool)`

GetMonitorModeOk returns a tuple with the MonitorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorMode

`func (o *Paladin) SetMonitorMode(v bool)`

SetMonitorMode sets MonitorMode field to given value.


### GetAdminViewOnly

`func (o *Paladin) GetAdminViewOnly() bool`

GetAdminViewOnly returns the AdminViewOnly field if non-nil, zero value otherwise.

### GetAdminViewOnlyOk

`func (o *Paladin) GetAdminViewOnlyOk() (*bool, bool)`

GetAdminViewOnlyOk returns a tuple with the AdminViewOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminViewOnly

`func (o *Paladin) SetAdminViewOnly(v bool)`

SetAdminViewOnly sets AdminViewOnly field to given value.


### GetEnabledConnectors

`func (o *Paladin) GetEnabledConnectors() []PaladinConnector`

GetEnabledConnectors returns the EnabledConnectors field if non-nil, zero value otherwise.

### GetEnabledConnectorsOk

`func (o *Paladin) GetEnabledConnectorsOk() (*[]PaladinConnector, bool)`

GetEnabledConnectorsOk returns a tuple with the EnabledConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledConnectors

`func (o *Paladin) SetEnabledConnectors(v []PaladinConnector)`

SetEnabledConnectors sets EnabledConnectors field to given value.


### GetInstructions

`func (o *Paladin) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *Paladin) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *Paladin) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


