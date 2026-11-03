# AgentActivityMcpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The server&#39;s name. | [optional] 
**Source** | Pointer to **string** | Where the server&#39;s definition came from, such as &#x60;plugin&#x60;, &#x60;sdk&#x60;, &#x60;user&#x60; or &#x60;project&#x60;. Not stored. | [optional] 

## Methods

### NewAgentActivityMcpServer

`func NewAgentActivityMcpServer() *AgentActivityMcpServer`

NewAgentActivityMcpServer instantiates a new AgentActivityMcpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityMcpServerWithDefaults

`func NewAgentActivityMcpServerWithDefaults() *AgentActivityMcpServer`

NewAgentActivityMcpServerWithDefaults instantiates a new AgentActivityMcpServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AgentActivityMcpServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentActivityMcpServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentActivityMcpServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentActivityMcpServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSource

`func (o *AgentActivityMcpServer) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AgentActivityMcpServer) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AgentActivityMcpServer) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AgentActivityMcpServer) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


