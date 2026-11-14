# AgentActivityClientSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportsToolCalls** | **bool** | Whether the client reports tool calls. When false its hook sends nothing and lets every tool call run, and Opal stores nothing the client still sends. | 
**ConversationContent** | Pointer to **string** | Who decides whether the client&#39;s conversation content (prompts, what each tool returned, the model&#39;s replies) is sent and stored. &#x60;always&#x60; sends it from every machine; &#x60;never&#x60; sends none, even from a machine installed with --content, and Opal stores none; &#x60;person_decides&#x60; leaves it to whether the machine was installed with --content. Absent for a client whose hook sends no conversation content (codex). | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


