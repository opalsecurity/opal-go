# AgentActivityOpalReported

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfiguredEmail** | Pointer to **string** | An email set by an administrator, an MDM profile or the installer. | [optional] 
**ClaudeAccountEmail** | Pointer to **string** | The email of the account Claude Code is signed in with. | [optional] 
**CodexAccountEmail** | Pointer to **string** | The email of the ChatGPT account Codex is signed in with. | [optional] 
**CodexAuthMode** | Pointer to **string** | How Codex is signed in, from &#x60;auth_mode&#x60; in its &#x60;auth.json&#x60;: &#x60;chatgpt&#x60; for a ChatGPT sign-in, &#x60;apikey&#x60; for an API key. In an organization that has turned on person agents, a Codex call from a person who owns no OpenAI agent creates \&quot;&lt;Name&gt;&#39;s ChatGPT\&quot; for a ChatGPT sign-in and \&quot;&lt;Name&gt;&#39;s OpenAI\&quot; for an API key. Without it, no agent is created. | [optional] 
**ChatgptAccountEmail** | Pointer to **string** | The email of the ChatGPT account the sending runtime is signed in with. | [optional] 
**CursorAccountEmail** | Pointer to **string** | The email of the account Cursor is signed in with, from the &#x60;user_email&#x60; of Cursor&#39;s hook input. Tried before the other emails here. | [optional] 
**Client** | Pointer to **string** | Which agent runtime sent the event: &#x60;claude_code&#x60;, &#x60;codex&#x60;, &#x60;cursor&#x60;, &#x60;openai_agents_python&#x60;, &#x60;openai_agents_js&#x60;, &#x60;chatgpt&#x60; or &#x60;chatgpt_workspace_agent&#x60;. Any other value is stored as &#x60;unknown&#x60;. When absent, an event carrying &#x60;cursor_version&#x60; is Cursor, one carrying &#x60;turn_id&#x60; and &#x60;model&#x60; and no &#x60;prompt_id&#x60; is Codex, and anything else is Claude Code. | [optional] 
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

### GetCodexAccountEmail

`func (o *AgentActivityOpalReported) GetCodexAccountEmail() string`

GetCodexAccountEmail returns the CodexAccountEmail field if non-nil, zero value otherwise.

### GetCodexAccountEmailOk

`func (o *AgentActivityOpalReported) GetCodexAccountEmailOk() (*string, bool)`

GetCodexAccountEmailOk returns a tuple with the CodexAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodexAccountEmail

`func (o *AgentActivityOpalReported) SetCodexAccountEmail(v string)`

SetCodexAccountEmail sets CodexAccountEmail field to given value.

### HasCodexAccountEmail

`func (o *AgentActivityOpalReported) HasCodexAccountEmail() bool`

HasCodexAccountEmail returns a boolean if a field has been set.

### GetCodexAuthMode

`func (o *AgentActivityOpalReported) GetCodexAuthMode() string`

GetCodexAuthMode returns the CodexAuthMode field if non-nil, zero value otherwise.

### GetCodexAuthModeOk

`func (o *AgentActivityOpalReported) GetCodexAuthModeOk() (*string, bool)`

GetCodexAuthModeOk returns a tuple with the CodexAuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodexAuthMode

`func (o *AgentActivityOpalReported) SetCodexAuthMode(v string)`

SetCodexAuthMode sets CodexAuthMode field to given value.

### HasCodexAuthMode

`func (o *AgentActivityOpalReported) HasCodexAuthMode() bool`

HasCodexAuthMode returns a boolean if a field has been set.

### GetChatgptAccountEmail

`func (o *AgentActivityOpalReported) GetChatgptAccountEmail() string`

GetChatgptAccountEmail returns the ChatgptAccountEmail field if non-nil, zero value otherwise.

### GetChatgptAccountEmailOk

`func (o *AgentActivityOpalReported) GetChatgptAccountEmailOk() (*string, bool)`

GetChatgptAccountEmailOk returns a tuple with the ChatgptAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatgptAccountEmail

`func (o *AgentActivityOpalReported) SetChatgptAccountEmail(v string)`

SetChatgptAccountEmail sets ChatgptAccountEmail field to given value.

### HasChatgptAccountEmail

`func (o *AgentActivityOpalReported) HasChatgptAccountEmail() bool`

HasChatgptAccountEmail returns a boolean if a field has been set.

### GetCursorAccountEmail

`func (o *AgentActivityOpalReported) GetCursorAccountEmail() string`

GetCursorAccountEmail returns the CursorAccountEmail field if non-nil, zero value otherwise.

### GetCursorAccountEmailOk

`func (o *AgentActivityOpalReported) GetCursorAccountEmailOk() (*string, bool)`

GetCursorAccountEmailOk returns a tuple with the CursorAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorAccountEmail

`func (o *AgentActivityOpalReported) SetCursorAccountEmail(v string)`

SetCursorAccountEmail sets CursorAccountEmail field to given value.

### HasCursorAccountEmail

`func (o *AgentActivityOpalReported) HasCursorAccountEmail() bool`

HasCursorAccountEmail returns a boolean if a field has been set.

### GetClient

`func (o *AgentActivityOpalReported) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *AgentActivityOpalReported) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *AgentActivityOpalReported) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *AgentActivityOpalReported) HasClient() bool`

HasClient returns a boolean if a field has been set.

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


