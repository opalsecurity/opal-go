# AgentActivityOpalReported

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfiguredEmail** | Pointer to **string** | An email set by an administrator, an MDM profile or the installer. | [optional] 
**ClaudeAccountEmail** | Pointer to **string** | The email of the account Claude Code is signed in with. | [optional] 
**GitEmail** | Pointer to **string** | The machine&#39;s global &#x60;git config user.email&#x60;. | [optional] 
**OsUser** | Pointer to **string** | The operating-system username. Not used to find an agent. | [optional] 
**Hostname** | Pointer to **string** | The machine&#39;s name. The &#x60;X-Opal-Host&#x60; header wins when it names a different machine. | [optional] 

## Methods

### NewAgentActivityOpalReported

`func NewAgentActivityOpalReported() *AgentActivityOpalReported`

NewAgentActivityOpalReported instantiates a new AgentActivityOpalReported object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityOpalReportedWithDefaults

`func NewAgentActivityOpalReportedWithDefaults() *AgentActivityOpalReported`

NewAgentActivityOpalReportedWithDefaults instantiates a new AgentActivityOpalReported object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguredEmail

`func (o *AgentActivityOpalReported) GetConfiguredEmail() string`

GetConfiguredEmail returns the ConfiguredEmail field if non-nil, zero value otherwise.

### GetConfiguredEmailOk

`func (o *AgentActivityOpalReported) GetConfiguredEmailOk() (*string, bool)`

GetConfiguredEmailOk returns a tuple with the ConfiguredEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredEmail

`func (o *AgentActivityOpalReported) SetConfiguredEmail(v string)`

SetConfiguredEmail sets ConfiguredEmail field to given value.

### HasConfiguredEmail

`func (o *AgentActivityOpalReported) HasConfiguredEmail() bool`

HasConfiguredEmail returns a boolean if a field has been set.

### GetClaudeAccountEmail

`func (o *AgentActivityOpalReported) GetClaudeAccountEmail() string`

GetClaudeAccountEmail returns the ClaudeAccountEmail field if non-nil, zero value otherwise.

### GetClaudeAccountEmailOk

`func (o *AgentActivityOpalReported) GetClaudeAccountEmailOk() (*string, bool)`

GetClaudeAccountEmailOk returns a tuple with the ClaudeAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaudeAccountEmail

`func (o *AgentActivityOpalReported) SetClaudeAccountEmail(v string)`

SetClaudeAccountEmail sets ClaudeAccountEmail field to given value.

### HasClaudeAccountEmail

`func (o *AgentActivityOpalReported) HasClaudeAccountEmail() bool`

HasClaudeAccountEmail returns a boolean if a field has been set.

### GetGitEmail

`func (o *AgentActivityOpalReported) GetGitEmail() string`

GetGitEmail returns the GitEmail field if non-nil, zero value otherwise.

### GetGitEmailOk

`func (o *AgentActivityOpalReported) GetGitEmailOk() (*string, bool)`

GetGitEmailOk returns a tuple with the GitEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitEmail

`func (o *AgentActivityOpalReported) SetGitEmail(v string)`

SetGitEmail sets GitEmail field to given value.

### HasGitEmail

`func (o *AgentActivityOpalReported) HasGitEmail() bool`

HasGitEmail returns a boolean if a field has been set.

### GetOsUser

`func (o *AgentActivityOpalReported) GetOsUser() string`

GetOsUser returns the OsUser field if non-nil, zero value otherwise.

### GetOsUserOk

`func (o *AgentActivityOpalReported) GetOsUserOk() (*string, bool)`

GetOsUserOk returns a tuple with the OsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsUser

`func (o *AgentActivityOpalReported) SetOsUser(v string)`

SetOsUser sets OsUser field to given value.

### HasOsUser

`func (o *AgentActivityOpalReported) HasOsUser() bool`

HasOsUser returns a boolean if a field has been set.

### GetHostname

`func (o *AgentActivityOpalReported) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *AgentActivityOpalReported) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *AgentActivityOpalReported) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *AgentActivityOpalReported) HasHostname() bool`

HasHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


