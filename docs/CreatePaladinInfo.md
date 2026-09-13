# CreatePaladinInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Paladin. | 
**OwnerId** | **string** | The ID of the owner of the Paladin. | 
**MonitorMode** | Pointer to **bool** | When true, the Paladin reasons about requests but takes no action. Defaults to true. | [optional] [default to true]
**AdminViewOnly** | Pointer to **bool** | When true, recommendations are visible only to admins. Only meaningful when monitor_mode is true. Defaults to false. | [optional] [default to false]
**EnabledConnectors** | Pointer to [**[]PaladinConnector**](PaladinConnector.md) | The connectors the Paladin is allowed to use. | [optional] 
**Instructions** | Pointer to **string** | The free-form instructions that guide the Paladin&#39;s decisions. Optional; if omitted the Paladin is created without instructions. | [optional] 

## Methods

### NewCreatePaladinInfo

`func NewCreatePaladinInfo(name string, ownerId string, ) *CreatePaladinInfo`

NewCreatePaladinInfo instantiates a new CreatePaladinInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaladinInfoWithDefaults

`func NewCreatePaladinInfoWithDefaults() *CreatePaladinInfo`

NewCreatePaladinInfoWithDefaults instantiates a new CreatePaladinInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePaladinInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePaladinInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePaladinInfo) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *CreatePaladinInfo) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CreatePaladinInfo) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CreatePaladinInfo) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetMonitorMode

`func (o *CreatePaladinInfo) GetMonitorMode() bool`

GetMonitorMode returns the MonitorMode field if non-nil, zero value otherwise.

### GetMonitorModeOk

`func (o *CreatePaladinInfo) GetMonitorModeOk() (*bool, bool)`

GetMonitorModeOk returns a tuple with the MonitorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorMode

`func (o *CreatePaladinInfo) SetMonitorMode(v bool)`

SetMonitorMode sets MonitorMode field to given value.

### HasMonitorMode

`func (o *CreatePaladinInfo) HasMonitorMode() bool`

HasMonitorMode returns a boolean if a field has been set.

### GetAdminViewOnly

`func (o *CreatePaladinInfo) GetAdminViewOnly() bool`

GetAdminViewOnly returns the AdminViewOnly field if non-nil, zero value otherwise.

### GetAdminViewOnlyOk

`func (o *CreatePaladinInfo) GetAdminViewOnlyOk() (*bool, bool)`

GetAdminViewOnlyOk returns a tuple with the AdminViewOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminViewOnly

`func (o *CreatePaladinInfo) SetAdminViewOnly(v bool)`

SetAdminViewOnly sets AdminViewOnly field to given value.

### HasAdminViewOnly

`func (o *CreatePaladinInfo) HasAdminViewOnly() bool`

HasAdminViewOnly returns a boolean if a field has been set.

### GetEnabledConnectors

`func (o *CreatePaladinInfo) GetEnabledConnectors() []PaladinConnector`

GetEnabledConnectors returns the EnabledConnectors field if non-nil, zero value otherwise.

### GetEnabledConnectorsOk

`func (o *CreatePaladinInfo) GetEnabledConnectorsOk() (*[]PaladinConnector, bool)`

GetEnabledConnectorsOk returns a tuple with the EnabledConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledConnectors

`func (o *CreatePaladinInfo) SetEnabledConnectors(v []PaladinConnector)`

SetEnabledConnectors sets EnabledConnectors field to given value.

### HasEnabledConnectors

`func (o *CreatePaladinInfo) HasEnabledConnectors() bool`

HasEnabledConnectors returns a boolean if a field has been set.

### GetInstructions

`func (o *CreatePaladinInfo) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CreatePaladinInfo) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CreatePaladinInfo) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *CreatePaladinInfo) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


