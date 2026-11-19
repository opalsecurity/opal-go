# AgentActivityClientSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportsToolCalls** | **bool** | Whether the client reports tool calls. When false its hook sends nothing and lets every tool call run, and Opal stores nothing the client still sends. | 
**ConversationContent** | Pointer to **string** | Who decides whether the client&#39;s conversation content (prompts, what each tool returned, the model&#39;s replies) is sent and stored. &#x60;always&#x60; sends it from every machine; &#x60;never&#x60; sends none, even from a machine installed with --content, and Opal stores none; &#x60;person_decides&#x60; leaves it to whether the machine was installed with --content. For codex it is the prompts and replies the Codex hooks send, unless the Codex kit was rendered with --no-prompts. Absent for a client that sends no conversation content. | [optional] 
**PromptSecrets** | Pointer to **string** | What the client&#39;s hook does with a prompt it finds a secret in, on the machine, before the prompt goes on. &#x60;block&#x60; (the client&#39;s mode enforces and the organization enforces client policy) stops a prompt that holds a credential recognized by its shape, and flags one that holds only a secret recognized by its name. &#x60;flag&#x60; (any other mode that reports) lets the prompt go on. Either way the hook reports the prompt with &#x60;opal_prompt_check&#x60;. Absent for a client that reports nothing (Off), and from an Opal that predates the check; the hook then does not look. | [optional] 

## Methods

### NewAgentActivityClientSetting

`func NewAgentActivityClientSetting(reportsToolCalls bool, ) *AgentActivityClientSetting`

NewAgentActivityClientSetting instantiates a new AgentActivityClientSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityClientSettingWithDefaults

`func NewAgentActivityClientSettingWithDefaults() *AgentActivityClientSetting`

NewAgentActivityClientSettingWithDefaults instantiates a new AgentActivityClientSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportsToolCalls

`func (o *AgentActivityClientSetting) GetReportsToolCalls() bool`

GetReportsToolCalls returns the ReportsToolCalls field if non-nil, zero value otherwise.

### GetReportsToolCallsOk

`func (o *AgentActivityClientSetting) GetReportsToolCallsOk() (*bool, bool)`

GetReportsToolCallsOk returns a tuple with the ReportsToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsToolCalls

`func (o *AgentActivityClientSetting) SetReportsToolCalls(v bool)`

SetReportsToolCalls sets ReportsToolCalls field to given value.


### GetConversationContent

`func (o *AgentActivityClientSetting) GetConversationContent() string`

GetConversationContent returns the ConversationContent field if non-nil, zero value otherwise.

### GetConversationContentOk

`func (o *AgentActivityClientSetting) GetConversationContentOk() (*string, bool)`

GetConversationContentOk returns a tuple with the ConversationContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationContent

`func (o *AgentActivityClientSetting) SetConversationContent(v string)`

SetConversationContent sets ConversationContent field to given value.

### HasConversationContent

`func (o *AgentActivityClientSetting) HasConversationContent() bool`

HasConversationContent returns a boolean if a field has been set.

### GetPromptSecrets

`func (o *AgentActivityClientSetting) GetPromptSecrets() string`

GetPromptSecrets returns the PromptSecrets field if non-nil, zero value otherwise.

### GetPromptSecretsOk

`func (o *AgentActivityClientSetting) GetPromptSecretsOk() (*string, bool)`

GetPromptSecretsOk returns a tuple with the PromptSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptSecrets

`func (o *AgentActivityClientSetting) SetPromptSecrets(v string)`

SetPromptSecrets sets PromptSecrets field to given value.

### HasPromptSecrets

`func (o *AgentActivityClientSetting) HasPromptSecrets() bool`

HasPromptSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


