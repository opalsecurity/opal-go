# AgentActivityEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HookEventName** | **string** | Which hook fired, e.g. &#x60;PostToolUse&#x60;, &#x60;PreToolUse&#x60;, &#x60;PermissionDenied&#x60;. Required; an event with no name cannot be interpreted and is rejected. | 
**ToolInput** | Pointer to **map[string]interface{}** | The call&#39;s arguments, as the runtime sent them. Server-defined for an MCP tool; a fixed shape per built-in (&#x60;command&#x60; for Bash, &#x60;file_path&#x60; for Edit, &#x60;url&#x60; for WebFetch).  Stored as sent, and separately reduced to a normalized target the activity feed and policy can match on. The stored copy is the record of fact; the derived target is recomputed whenever the extractor improves.  Capped before storage. A capped payload is no longer valid JSON, so the row records that it was capped rather than leaving each reader to discover it by failing to parse. | [optional] 
**SessionId** | Pointer to **string** | The agent session this event belongs to. Opaque to Opal. | [optional] 
**PromptId** | Pointer to **string** | Groups every tool call made under one user instruction, so a run can be read back as the instruction that caused it rather than as unrelated rows. | [optional] 
**AgentId** | Pointer to **string** | Set only when the call came from an autonomous subagent rather than the session a person is typing into. Stored as &#x60;subagent_id&#x60; - Opal&#39;s own &#x60;agent_id&#x60; already means something else. | [optional] 
**AgentType** | Pointer to **string** | The subagent&#39;s type, e.g. &#x60;general-purpose&#x60;. Set alongside &#x60;agent_id&#x60;. | [optional] 
**PermissionMode** | Pointer to **string** | The session&#39;s permission mode, e.g. &#x60;default&#x60;, &#x60;acceptEdits&#x60;, &#x60;bypassPermissions&#x60;. A session running &#x60;bypassPermissions&#x60; had no approval gate in front of any tool call it made. | [optional] 
**ActorEmail** | Pointer to **string** | The identity the sender reports for this specific event. **Unverified**, the same as &#x60;X-Opal-Actor&#x60;, and used only when neither the API key nor that header identified a developer. An &#x60;http&#x60; hook cannot populate this; a &#x60;command&#x60;-type hook or a relay can (&#x60;$USER&#x60;, a hostname, &#x60;git config user.email&#x60;), which is why it is per-event rather than per-request - a relay can batch several developers. | [optional] 
**OpalReported** | Pointer to [**AgentActivityOpalReported**](AgentActivityOpalReported.md) |  | [optional] 
**DurationMs** | Pointer to **interface{}** | How long the tool call took to run, in milliseconds, as Claude Code measured it. Claude Code&#39;s documentation says it leaves out time spent in permission prompts and in &#x60;PreToolUse&#x60; hooks. It is sent with &#x60;PostToolUse&#x60; and &#x60;PostToolUseFailure&#x60; only; it is ignored on any other event. **Unverified.** Stored only when it is a number from 0 to 86400000 (24 hours), rounded to a whole millisecond; any other value is dropped and the event is stored without it. A value of the wrong type never rejects the event or the batch it came in. | [optional] 
**TranscriptPath** | Pointer to **string** | Read at ingest **only** to parse an operating-system username out of, as the last-resort attribution step, and then discarded. The path itself is never stored: it is an absolute path on an employee&#39;s laptop and Opal has no business accumulating those. | [optional] 
**ToolName** | Pointer to **string** | The tool invoked, for tool-scoped events. Empty for session-level events.  An MCP tool is named &#x60;mcp__&lt;server&gt;__&lt;tool&gt;&#x60;. The whole string is stored as sent, and the server and the tool are also stored in their own fields (see &#x60;mcp_server&#x60;). | [optional] 
**McpServer** | Pointer to [**AgentActivityMcpServer**](AgentActivityMcpServer.md) |  | [optional] 
**OpalMcpServer** | Pointer to [**AgentActivityOpalMcpServer**](AgentActivityOpalMcpServer.md) |  | [optional] 
**ToolUseId** | Pointer to **string** | The tool invocation&#39;s id, used together with &#x60;hook_event_name&#x60; to deduplicate a retried delivery. Absent on events that carry none, which are stored at-least-once unless they carry &#x60;opal_event_id&#x60;. | [optional] 
**OpalEventId** | Pointer to **string** | An id Opal&#39;s hook kit gives each event it holds for sending, so that a retried delivery of an event without a &#x60;tool_use_id&#x60; (a prompt, the end of a turn) is stored once. Used only for an event that has no &#x60;tool_use_id&#x60; and is not a tool-call event, and only when it is 8 to 64 letters, digits and dashes; otherwise ignored. | [optional] 
**Cwd** | Pointer to **string** | The working directory the agent was operating in. | [optional] 
**Timestamp** | Pointer to **time.Time** | When the client says the event occurred (RFC 3339). Clamped server-side to a window around receipt, because a client-controlled sort key is otherwise a way to pin rows to the top of every list. Defaults to receipt time when absent or unparseable. | [optional] 
**Prompt** | Pointer to **interface{}** | The text a person submitted (&#x60;UserPromptSubmit&#x60;), a string. Stored only as described above. Read whatever its type, so that a change in its shape cannot make Opal refuse the event. | [optional] 
**ToolResponse** | Pointer to **interface{}** | What the tool returned (&#x60;PostToolUse&#x60;), in the tool&#39;s own shape, or as text when the sender already converted it. Stored only as described above, as text: every image is replaced by a placeholder that names its type and size. | [optional] 
**LastAssistantMessage** | Pointer to **interface{}** | The text of the model&#39;s final message in the turn (&#x60;Stop&#x60;, &#x60;SubagentStop&#x60;), a string. Stored only as described above, and only when &#x60;opal_content.assistant_text&#x60; is absent. Read whatever its type, like &#x60;prompt&#x60;. | [optional] 
**OpalContent** | Pointer to [**AgentActivityContentReport**](AgentActivityContentReport.md) |  | [optional] 

## Methods

### NewAgentActivityEvent

`func NewAgentActivityEvent(hookEventName string, ) *AgentActivityEvent`

NewAgentActivityEvent instantiates a new AgentActivityEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityEventWithDefaults

`func NewAgentActivityEventWithDefaults() *AgentActivityEvent`

NewAgentActivityEventWithDefaults instantiates a new AgentActivityEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHookEventName

`func (o *AgentActivityEvent) GetHookEventName() string`

GetHookEventName returns the HookEventName field if non-nil, zero value otherwise.

### GetHookEventNameOk

`func (o *AgentActivityEvent) GetHookEventNameOk() (*string, bool)`

GetHookEventNameOk returns a tuple with the HookEventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookEventName

`func (o *AgentActivityEvent) SetHookEventName(v string)`

SetHookEventName sets HookEventName field to given value.


### GetToolInput

`func (o *AgentActivityEvent) GetToolInput() map[string]interface{}`

GetToolInput returns the ToolInput field if non-nil, zero value otherwise.

### GetToolInputOk

`func (o *AgentActivityEvent) GetToolInputOk() (*map[string]interface{}, bool)`

GetToolInputOk returns a tuple with the ToolInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolInput

`func (o *AgentActivityEvent) SetToolInput(v map[string]interface{})`

SetToolInput sets ToolInput field to given value.

### HasToolInput

`func (o *AgentActivityEvent) HasToolInput() bool`

HasToolInput returns a boolean if a field has been set.

### GetSessionId

`func (o *AgentActivityEvent) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *AgentActivityEvent) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *AgentActivityEvent) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *AgentActivityEvent) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetPromptId

`func (o *AgentActivityEvent) GetPromptId() string`

GetPromptId returns the PromptId field if non-nil, zero value otherwise.

### GetPromptIdOk

`func (o *AgentActivityEvent) GetPromptIdOk() (*string, bool)`

GetPromptIdOk returns a tuple with the PromptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptId

`func (o *AgentActivityEvent) SetPromptId(v string)`

SetPromptId sets PromptId field to given value.

### HasPromptId

`func (o *AgentActivityEvent) HasPromptId() bool`

HasPromptId returns a boolean if a field has been set.

### GetAgentId

`func (o *AgentActivityEvent) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AgentActivityEvent) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AgentActivityEvent) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *AgentActivityEvent) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAgentType

`func (o *AgentActivityEvent) GetAgentType() string`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *AgentActivityEvent) GetAgentTypeOk() (*string, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *AgentActivityEvent) SetAgentType(v string)`

SetAgentType sets AgentType field to given value.

### HasAgentType

`func (o *AgentActivityEvent) HasAgentType() bool`

HasAgentType returns a boolean if a field has been set.

### GetPermissionMode

`func (o *AgentActivityEvent) GetPermissionMode() string`

GetPermissionMode returns the PermissionMode field if non-nil, zero value otherwise.

### GetPermissionModeOk

`func (o *AgentActivityEvent) GetPermissionModeOk() (*string, bool)`

GetPermissionModeOk returns a tuple with the PermissionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionMode

`func (o *AgentActivityEvent) SetPermissionMode(v string)`

SetPermissionMode sets PermissionMode field to given value.

### HasPermissionMode

`func (o *AgentActivityEvent) HasPermissionMode() bool`

HasPermissionMode returns a boolean if a field has been set.

### GetActorEmail

`func (o *AgentActivityEvent) GetActorEmail() string`

GetActorEmail returns the ActorEmail field if non-nil, zero value otherwise.

### GetActorEmailOk

`func (o *AgentActivityEvent) GetActorEmailOk() (*string, bool)`

GetActorEmailOk returns a tuple with the ActorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorEmail

`func (o *AgentActivityEvent) SetActorEmail(v string)`

SetActorEmail sets ActorEmail field to given value.

### HasActorEmail

`func (o *AgentActivityEvent) HasActorEmail() bool`

HasActorEmail returns a boolean if a field has been set.

### GetOpalReported

`func (o *AgentActivityEvent) GetOpalReported() AgentActivityOpalReported`

GetOpalReported returns the OpalReported field if non-nil, zero value otherwise.

### GetOpalReportedOk

`func (o *AgentActivityEvent) GetOpalReportedOk() (*AgentActivityOpalReported, bool)`

GetOpalReportedOk returns a tuple with the OpalReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpalReported

`func (o *AgentActivityEvent) SetOpalReported(v AgentActivityOpalReported)`

SetOpalReported sets OpalReported field to given value.

### HasOpalReported

`func (o *AgentActivityEvent) HasOpalReported() bool`

HasOpalReported returns a boolean if a field has been set.

### GetDurationMs

`func (o *AgentActivityEvent) GetDurationMs() interface{}`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *AgentActivityEvent) GetDurationMsOk() (*interface{}, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *AgentActivityEvent) SetDurationMs(v interface{})`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *AgentActivityEvent) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### SetDurationMsNil

`func (o *AgentActivityEvent) SetDurationMsNil(b bool)`

 SetDurationMsNil sets the value for DurationMs to be an explicit nil

### UnsetDurationMs
`func (o *AgentActivityEvent) UnsetDurationMs()`

UnsetDurationMs ensures that no value is present for DurationMs, not even an explicit nil
### GetTranscriptPath

`func (o *AgentActivityEvent) GetTranscriptPath() string`

GetTranscriptPath returns the TranscriptPath field if non-nil, zero value otherwise.

### GetTranscriptPathOk

`func (o *AgentActivityEvent) GetTranscriptPathOk() (*string, bool)`

GetTranscriptPathOk returns a tuple with the TranscriptPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptPath

`func (o *AgentActivityEvent) SetTranscriptPath(v string)`

SetTranscriptPath sets TranscriptPath field to given value.

### HasTranscriptPath

`func (o *AgentActivityEvent) HasTranscriptPath() bool`

HasTranscriptPath returns a boolean if a field has been set.

### GetToolName

`func (o *AgentActivityEvent) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *AgentActivityEvent) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *AgentActivityEvent) SetToolName(v string)`

SetToolName sets ToolName field to given value.

### HasToolName

`func (o *AgentActivityEvent) HasToolName() bool`

HasToolName returns a boolean if a field has been set.

### GetMcpServer

`func (o *AgentActivityEvent) GetMcpServer() AgentActivityMcpServer`

GetMcpServer returns the McpServer field if non-nil, zero value otherwise.

### GetMcpServerOk

`func (o *AgentActivityEvent) GetMcpServerOk() (*AgentActivityMcpServer, bool)`

GetMcpServerOk returns a tuple with the McpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServer

`func (o *AgentActivityEvent) SetMcpServer(v AgentActivityMcpServer)`

SetMcpServer sets McpServer field to given value.

### HasMcpServer

`func (o *AgentActivityEvent) HasMcpServer() bool`

HasMcpServer returns a boolean if a field has been set.

### GetOpalMcpServer

`func (o *AgentActivityEvent) GetOpalMcpServer() AgentActivityOpalMcpServer`

GetOpalMcpServer returns the OpalMcpServer field if non-nil, zero value otherwise.

### GetOpalMcpServerOk

`func (o *AgentActivityEvent) GetOpalMcpServerOk() (*AgentActivityOpalMcpServer, bool)`

GetOpalMcpServerOk returns a tuple with the OpalMcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpalMcpServer

`func (o *AgentActivityEvent) SetOpalMcpServer(v AgentActivityOpalMcpServer)`

SetOpalMcpServer sets OpalMcpServer field to given value.

### HasOpalMcpServer

`func (o *AgentActivityEvent) HasOpalMcpServer() bool`

HasOpalMcpServer returns a boolean if a field has been set.

### GetToolUseId

`func (o *AgentActivityEvent) GetToolUseId() string`

GetToolUseId returns the ToolUseId field if non-nil, zero value otherwise.

### GetToolUseIdOk

`func (o *AgentActivityEvent) GetToolUseIdOk() (*string, bool)`

GetToolUseIdOk returns a tuple with the ToolUseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolUseId

`func (o *AgentActivityEvent) SetToolUseId(v string)`

SetToolUseId sets ToolUseId field to given value.

### HasToolUseId

`func (o *AgentActivityEvent) HasToolUseId() bool`

HasToolUseId returns a boolean if a field has been set.

### GetOpalEventId

`func (o *AgentActivityEvent) GetOpalEventId() string`

GetOpalEventId returns the OpalEventId field if non-nil, zero value otherwise.

### GetOpalEventIdOk

`func (o *AgentActivityEvent) GetOpalEventIdOk() (*string, bool)`

GetOpalEventIdOk returns a tuple with the OpalEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpalEventId

`func (o *AgentActivityEvent) SetOpalEventId(v string)`

SetOpalEventId sets OpalEventId field to given value.

### HasOpalEventId

`func (o *AgentActivityEvent) HasOpalEventId() bool`

HasOpalEventId returns a boolean if a field has been set.

### GetCwd

`func (o *AgentActivityEvent) GetCwd() string`

GetCwd returns the Cwd field if non-nil, zero value otherwise.

### GetCwdOk

`func (o *AgentActivityEvent) GetCwdOk() (*string, bool)`

GetCwdOk returns a tuple with the Cwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwd

`func (o *AgentActivityEvent) SetCwd(v string)`

SetCwd sets Cwd field to given value.

### HasCwd

`func (o *AgentActivityEvent) HasCwd() bool`

HasCwd returns a boolean if a field has been set.

### GetTimestamp

`func (o *AgentActivityEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AgentActivityEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AgentActivityEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AgentActivityEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetPrompt

`func (o *AgentActivityEvent) GetPrompt() interface{}`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AgentActivityEvent) GetPromptOk() (*interface{}, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AgentActivityEvent) SetPrompt(v interface{})`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AgentActivityEvent) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### SetPromptNil

`func (o *AgentActivityEvent) SetPromptNil(b bool)`

 SetPromptNil sets the value for Prompt to be an explicit nil

### UnsetPrompt
`func (o *AgentActivityEvent) UnsetPrompt()`

UnsetPrompt ensures that no value is present for Prompt, not even an explicit nil
### GetToolResponse

`func (o *AgentActivityEvent) GetToolResponse() interface{}`

GetToolResponse returns the ToolResponse field if non-nil, zero value otherwise.

### GetToolResponseOk

`func (o *AgentActivityEvent) GetToolResponseOk() (*interface{}, bool)`

GetToolResponseOk returns a tuple with the ToolResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolResponse

`func (o *AgentActivityEvent) SetToolResponse(v interface{})`

SetToolResponse sets ToolResponse field to given value.

### HasToolResponse

`func (o *AgentActivityEvent) HasToolResponse() bool`

HasToolResponse returns a boolean if a field has been set.

### SetToolResponseNil

`func (o *AgentActivityEvent) SetToolResponseNil(b bool)`

 SetToolResponseNil sets the value for ToolResponse to be an explicit nil

### UnsetToolResponse
`func (o *AgentActivityEvent) UnsetToolResponse()`

UnsetToolResponse ensures that no value is present for ToolResponse, not even an explicit nil
### GetLastAssistantMessage

`func (o *AgentActivityEvent) GetLastAssistantMessage() interface{}`

GetLastAssistantMessage returns the LastAssistantMessage field if non-nil, zero value otherwise.

### GetLastAssistantMessageOk

`func (o *AgentActivityEvent) GetLastAssistantMessageOk() (*interface{}, bool)`

GetLastAssistantMessageOk returns a tuple with the LastAssistantMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAssistantMessage

`func (o *AgentActivityEvent) SetLastAssistantMessage(v interface{})`

SetLastAssistantMessage sets LastAssistantMessage field to given value.

### HasLastAssistantMessage

`func (o *AgentActivityEvent) HasLastAssistantMessage() bool`

HasLastAssistantMessage returns a boolean if a field has been set.

### SetLastAssistantMessageNil

`func (o *AgentActivityEvent) SetLastAssistantMessageNil(b bool)`

 SetLastAssistantMessageNil sets the value for LastAssistantMessage to be an explicit nil

### UnsetLastAssistantMessage
`func (o *AgentActivityEvent) UnsetLastAssistantMessage()`

UnsetLastAssistantMessage ensures that no value is present for LastAssistantMessage, not even an explicit nil
### GetOpalContent

`func (o *AgentActivityEvent) GetOpalContent() AgentActivityContentReport`

GetOpalContent returns the OpalContent field if non-nil, zero value otherwise.

### GetOpalContentOk

`func (o *AgentActivityEvent) GetOpalContentOk() (*AgentActivityContentReport, bool)`

GetOpalContentOk returns a tuple with the OpalContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpalContent

`func (o *AgentActivityEvent) SetOpalContent(v AgentActivityContentReport)`

SetOpalContent sets OpalContent field to given value.

### HasOpalContent

`func (o *AgentActivityEvent) HasOpalContent() bool`

HasOpalContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


