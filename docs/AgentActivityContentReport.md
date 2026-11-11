# AgentActivityContentReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssistantText** | Pointer to **string** | Every text the model wrote in the turn, read from the session&#39;s transcript (&#x60;Stop&#x60;, &#x60;SubagentStop&#x60;). Preferred over &#x60;last_assistant_message&#x60;, which holds only the final message. | [optional] 
**Thinking** | Pointer to **string** | The model&#39;s thinking in the turn, when the transcript holds it as plain text. | [optional] 
**Images** | Pointer to [**AgentActivityImages**](AgentActivityImages.md) |  | [optional] 
**Truncated** | Pointer to **map[string]int32** | The fields the kit cut (&#x60;prompt&#x60;, &#x60;tool_response&#x60;, &#x60;assistant_text&#x60;, &#x60;last_assistant_message&#x60;, &#x60;thinking&#x60;), each with its size in bytes before the cut. A prompt or a tool result keeps its start; a turn&#39;s text and thinking keep their end. Sizes are stored up to 2147483647. | [optional] 

## Methods

### NewAgentActivityContentReport

`func NewAgentActivityContentReport() *AgentActivityContentReport`

NewAgentActivityContentReport instantiates a new AgentActivityContentReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityContentReportWithDefaults

`func NewAgentActivityContentReportWithDefaults() *AgentActivityContentReport`

NewAgentActivityContentReportWithDefaults instantiates a new AgentActivityContentReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistantText

`func (o *AgentActivityContentReport) GetAssistantText() string`

GetAssistantText returns the AssistantText field if non-nil, zero value otherwise.

### GetAssistantTextOk

`func (o *AgentActivityContentReport) GetAssistantTextOk() (*string, bool)`

GetAssistantTextOk returns a tuple with the AssistantText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantText

`func (o *AgentActivityContentReport) SetAssistantText(v string)`

SetAssistantText sets AssistantText field to given value.

### HasAssistantText

`func (o *AgentActivityContentReport) HasAssistantText() bool`

HasAssistantText returns a boolean if a field has been set.

### GetThinking

`func (o *AgentActivityContentReport) GetThinking() string`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *AgentActivityContentReport) GetThinkingOk() (*string, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *AgentActivityContentReport) SetThinking(v string)`

SetThinking sets Thinking field to given value.

### HasThinking

`func (o *AgentActivityContentReport) HasThinking() bool`

HasThinking returns a boolean if a field has been set.

### GetImages

`func (o *AgentActivityContentReport) GetImages() AgentActivityImages`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *AgentActivityContentReport) GetImagesOk() (*AgentActivityImages, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *AgentActivityContentReport) SetImages(v AgentActivityImages)`

SetImages sets Images field to given value.

### HasImages

`func (o *AgentActivityContentReport) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetTruncated

`func (o *AgentActivityContentReport) GetTruncated() map[string]int32`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *AgentActivityContentReport) GetTruncatedOk() (*map[string]int32, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *AgentActivityContentReport) SetTruncated(v map[string]int32)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *AgentActivityContentReport) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


