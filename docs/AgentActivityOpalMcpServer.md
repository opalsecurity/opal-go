# AgentActivityOpalMcpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The server&#39;s URL, for an &#x60;http&#x60; or &#x60;sse&#x60; server. | [optional] 
**Transport** | Pointer to **string** | &#x60;http&#x60;, &#x60;sse&#x60; or &#x60;stdio&#x60;. | [optional] 
**CommandHash** | Pointer to **string** | For a &#x60;stdio&#x60; server, a SHA-256 of its command&#39;s shape, in hex. Arguments that can hold a secret are replaced before hashing, and the command itself is never sent. | [optional] 
**Scope** | Pointer to **string** | The configuration the server was found in, such as &#x60;user&#x60;, &#x60;project&#x60;, &#x60;local&#x60;, &#x60;managed&#x60; or &#x60;plugin&#x60;. | [optional] 

## Methods

### NewAgentActivityOpalMcpServer

`func NewAgentActivityOpalMcpServer() *AgentActivityOpalMcpServer`

NewAgentActivityOpalMcpServer instantiates a new AgentActivityOpalMcpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityOpalMcpServerWithDefaults

`func NewAgentActivityOpalMcpServerWithDefaults() *AgentActivityOpalMcpServer`

NewAgentActivityOpalMcpServerWithDefaults instantiates a new AgentActivityOpalMcpServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AgentActivityOpalMcpServer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AgentActivityOpalMcpServer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AgentActivityOpalMcpServer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AgentActivityOpalMcpServer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTransport

`func (o *AgentActivityOpalMcpServer) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *AgentActivityOpalMcpServer) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *AgentActivityOpalMcpServer) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *AgentActivityOpalMcpServer) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetCommandHash

`func (o *AgentActivityOpalMcpServer) GetCommandHash() string`

GetCommandHash returns the CommandHash field if non-nil, zero value otherwise.

### GetCommandHashOk

`func (o *AgentActivityOpalMcpServer) GetCommandHashOk() (*string, bool)`

GetCommandHashOk returns a tuple with the CommandHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandHash

`func (o *AgentActivityOpalMcpServer) SetCommandHash(v string)`

SetCommandHash sets CommandHash field to given value.

### HasCommandHash

`func (o *AgentActivityOpalMcpServer) HasCommandHash() bool`

HasCommandHash returns a boolean if a field has been set.

### GetScope

`func (o *AgentActivityOpalMcpServer) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AgentActivityOpalMcpServer) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AgentActivityOpalMcpServer) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AgentActivityOpalMcpServer) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


