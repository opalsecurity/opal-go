# McpConfigSubagent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Scope** | Pointer to **string** | &#x60;user&#x60;, &#x60;project&#x60;, &#x60;plugin&#x60; or &#x60;managed&#x60;. | [optional] 
**Tools** | Pointer to **[]string** | The tools its file lists. Absent when the file lists none, which lets the subagent use every tool of the session, MCP tools included. | [optional] 

## Methods

### NewMcpConfigSubagent

`func NewMcpConfigSubagent(name string, ) *McpConfigSubagent`

NewMcpConfigSubagent instantiates a new McpConfigSubagent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpConfigSubagentWithDefaults

`func NewMcpConfigSubagentWithDefaults() *McpConfigSubagent`

NewMcpConfigSubagentWithDefaults instantiates a new McpConfigSubagent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *McpConfigSubagent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpConfigSubagent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpConfigSubagent) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *McpConfigSubagent) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *McpConfigSubagent) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *McpConfigSubagent) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *McpConfigSubagent) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTools

`func (o *McpConfigSubagent) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *McpConfigSubagent) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *McpConfigSubagent) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *McpConfigSubagent) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


