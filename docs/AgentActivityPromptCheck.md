# AgentActivityPromptCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outcome** | Pointer to **string** | &#x60;blocked&#x60; when the kit stopped the prompt before it reached the model, &#x60;flagged&#x60; when it let the prompt go on. | [optional] 
**Kinds** | Pointer to **[]string** | The kinds of secret found, as Opal&#39;s redaction names them (&#x60;github_token&#x60;, &#x60;aws_access_key_id&#x60;, &#x60;env_secret&#x60;, ...). Never the secret itself. | [optional] 

## Methods

### NewAgentActivityPromptCheck

`func NewAgentActivityPromptCheck() *AgentActivityPromptCheck`

NewAgentActivityPromptCheck instantiates a new AgentActivityPromptCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityPromptCheckWithDefaults

`func NewAgentActivityPromptCheckWithDefaults() *AgentActivityPromptCheck`

NewAgentActivityPromptCheckWithDefaults instantiates a new AgentActivityPromptCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutcome

`func (o *AgentActivityPromptCheck) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AgentActivityPromptCheck) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AgentActivityPromptCheck) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *AgentActivityPromptCheck) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetKinds

`func (o *AgentActivityPromptCheck) GetKinds() []string`

GetKinds returns the Kinds field if non-nil, zero value otherwise.

### GetKindsOk

`func (o *AgentActivityPromptCheck) GetKindsOk() (*[]string, bool)`

GetKindsOk returns a tuple with the Kinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKinds

`func (o *AgentActivityPromptCheck) SetKinds(v []string)`

SetKinds sets Kinds field to given value.

### HasKinds

`func (o *AgentActivityPromptCheck) HasKinds() bool`

HasKinds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


