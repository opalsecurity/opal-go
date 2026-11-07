# UpdatePaladinInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Paladin. | 
**MonitorMode** | Pointer to **bool** | When true, the Paladin reasons about requests but takes no action. If omitted, the existing value is preserved. | [optional] 
**AdminViewOnly** | Pointer to **bool** | When true, recommendations are visible only to admins. Only meaningful when monitor_mode is true. If omitted, the existing value is preserved. | [optional] 
**EnabledConnectors** | Pointer to [**[]PaladinConnector**](PaladinConnector.md) | The connectors the Paladin is allowed to use. If omitted, the existing connectors are preserved. | [optional] 
**Instructions** | Pointer to **string** | The free-form instructions that guide the Paladin&#39;s decisions. If omitted, the existing instructions are preserved. | [optional] 

## Methods

### NewUpdatePaladinInfo

`func NewUpdatePaladinInfo(name string, ) *UpdatePaladinInfo`

NewUpdatePaladinInfo instantiates a new UpdatePaladinInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaladinInfoWithDefaults

`func NewUpdatePaladinInfoWithDefaults() *UpdatePaladinInfo`

NewUpdatePaladinInfoWithDefaults instantiates a new UpdatePaladinInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePaladinInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePaladinInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePaladinInfo) SetName(v string)`

SetName sets Name field to given value.


### GetMonitorMode

`func (o *UpdatePaladinInfo) GetMonitorMode() bool`

GetMonitorMode returns the MonitorMode field if non-nil, zero value otherwise.

### GetMonitorModeOk

`func (o *UpdatePaladinInfo) GetMonitorModeOk() (*bool, bool)`

GetMonitorModeOk returns a tuple with the MonitorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorMode

`func (o *UpdatePaladinInfo) SetMonitorMode(v bool)`

SetMonitorMode sets MonitorMode field to given value.

### HasMonitorMode

`func (o *UpdatePaladinInfo) HasMonitorMode() bool`

HasMonitorMode returns a boolean if a field has been set.

### GetAdminViewOnly

`func (o *UpdatePaladinInfo) GetAdminViewOnly() bool`

GetAdminViewOnly returns the AdminViewOnly field if non-nil, zero value otherwise.

### GetAdminViewOnlyOk

`func (o *UpdatePaladinInfo) GetAdminViewOnlyOk() (*bool, bool)`

GetAdminViewOnlyOk returns a tuple with the AdminViewOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminViewOnly

`func (o *UpdatePaladinInfo) SetAdminViewOnly(v bool)`

SetAdminViewOnly sets AdminViewOnly field to given value.

### HasAdminViewOnly

`func (o *UpdatePaladinInfo) HasAdminViewOnly() bool`

HasAdminViewOnly returns a boolean if a field has been set.

### GetEnabledConnectors

`func (o *UpdatePaladinInfo) GetEnabledConnectors() []PaladinConnector`

GetEnabledConnectors returns the EnabledConnectors field if non-nil, zero value otherwise.

### GetEnabledConnectorsOk

`func (o *UpdatePaladinInfo) GetEnabledConnectorsOk() (*[]PaladinConnector, bool)`

GetEnabledConnectorsOk returns a tuple with the EnabledConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledConnectors

`func (o *UpdatePaladinInfo) SetEnabledConnectors(v []PaladinConnector)`

SetEnabledConnectors sets EnabledConnectors field to given value.

### HasEnabledConnectors

`func (o *UpdatePaladinInfo) HasEnabledConnectors() bool`

HasEnabledConnectors returns a boolean if a field has been set.

### GetInstructions

`func (o *UpdatePaladinInfo) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *UpdatePaladinInfo) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *UpdatePaladinInfo) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *UpdatePaladinInfo) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


