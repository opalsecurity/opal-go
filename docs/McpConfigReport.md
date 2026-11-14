# McpConfigReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | [**[]McpConfigServer**](McpConfigServer.md) |  | 
**Reporter** | Pointer to [**McpConfigReporter**](McpConfigReporter.md) |  | [optional] 
**Plugins** | Pointer to [**[]McpConfigPlugin**](McpConfigPlugin.md) |  | [optional] 
**Skills** | Pointer to [**[]McpConfigSkill**](McpConfigSkill.md) |  | [optional] 
**Subagents** | Pointer to [**[]McpConfigSubagent**](McpConfigSubagent.md) |  | [optional] 
**Permissions** | Pointer to [**[]McpConfigPermission**](McpConfigPermission.md) |  | [optional] 

## Methods

### NewMcpConfigReport

`func NewMcpConfigReport(servers []McpConfigServer, ) *McpConfigReport`

NewMcpConfigReport instantiates a new McpConfigReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpConfigReportWithDefaults

`func NewMcpConfigReportWithDefaults() *McpConfigReport`

NewMcpConfigReportWithDefaults instantiates a new McpConfigReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *McpConfigReport) GetServers() []McpConfigServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *McpConfigReport) GetServersOk() (*[]McpConfigServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *McpConfigReport) SetServers(v []McpConfigServer)`

SetServers sets Servers field to given value.


### GetReporter

`func (o *McpConfigReport) GetReporter() McpConfigReporter`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *McpConfigReport) GetReporterOk() (*McpConfigReporter, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *McpConfigReport) SetReporter(v McpConfigReporter)`

SetReporter sets Reporter field to given value.

### HasReporter

`func (o *McpConfigReport) HasReporter() bool`

HasReporter returns a boolean if a field has been set.

### GetPlugins

`func (o *McpConfigReport) GetPlugins() []McpConfigPlugin`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *McpConfigReport) GetPluginsOk() (*[]McpConfigPlugin, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *McpConfigReport) SetPlugins(v []McpConfigPlugin)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *McpConfigReport) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetSkills

`func (o *McpConfigReport) GetSkills() []McpConfigSkill`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *McpConfigReport) GetSkillsOk() (*[]McpConfigSkill, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *McpConfigReport) SetSkills(v []McpConfigSkill)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *McpConfigReport) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetSubagents

`func (o *McpConfigReport) GetSubagents() []McpConfigSubagent`

GetSubagents returns the Subagents field if non-nil, zero value otherwise.

### GetSubagentsOk

`func (o *McpConfigReport) GetSubagentsOk() (*[]McpConfigSubagent, bool)`

GetSubagentsOk returns a tuple with the Subagents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubagents

`func (o *McpConfigReport) SetSubagents(v []McpConfigSubagent)`

SetSubagents sets Subagents field to given value.

### HasSubagents

`func (o *McpConfigReport) HasSubagents() bool`

HasSubagents returns a boolean if a field has been set.

### GetPermissions

`func (o *McpConfigReport) GetPermissions() []McpConfigPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *McpConfigReport) GetPermissionsOk() (*[]McpConfigPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *McpConfigReport) SetPermissions(v []McpConfigPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *McpConfigReport) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


