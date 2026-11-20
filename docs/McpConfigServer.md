# McpConfigServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The server&#39;s name in the configuration. | 
**Scope** | Pointer to **string** | &#x60;user&#x60;, &#x60;project&#x60;, &#x60;local&#x60;, &#x60;managed&#x60; or &#x60;plugin&#x60;. | [optional] 
**Transport** | Pointer to **string** | &#x60;http&#x60;, &#x60;sse&#x60; or &#x60;stdio&#x60;. | [optional] 
**Url** | Pointer to **string** | The server&#39;s URL, for an &#x60;http&#x60; or &#x60;sse&#x60; server. | [optional] 
**CommandHash** | Pointer to **string** | For a &#x60;stdio&#x60; server, a SHA-256 of its command&#39;s shape, in hex. | [optional] 

## Methods

### NewMcpConfigServer

`func NewMcpConfigServer(name string, ) *McpConfigServer`

NewMcpConfigServer instantiates a new McpConfigServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpConfigServerWithDefaults

`func NewMcpConfigServerWithDefaults() *McpConfigServer`

NewMcpConfigServerWithDefaults instantiates a new McpConfigServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *McpConfigServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpConfigServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpConfigServer) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *McpConfigServer) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *McpConfigServer) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *McpConfigServer) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *McpConfigServer) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTransport

`func (o *McpConfigServer) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *McpConfigServer) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *McpConfigServer) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *McpConfigServer) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetUrl

`func (o *McpConfigServer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *McpConfigServer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *McpConfigServer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *McpConfigServer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCommandHash

`func (o *McpConfigServer) GetCommandHash() string`

GetCommandHash returns the CommandHash field if non-nil, zero value otherwise.

### GetCommandHashOk

`func (o *McpConfigServer) GetCommandHashOk() (*string, bool)`

GetCommandHashOk returns a tuple with the CommandHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandHash

`func (o *McpConfigServer) SetCommandHash(v string)`

SetCommandHash sets CommandHash field to given value.

### HasCommandHash

`func (o *McpConfigServer) HasCommandHash() bool`

HasCommandHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


