# McpConfigPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Marketplace** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** | The installed version, when the kit could read it. | [optional] 
**Scope** | Pointer to **string** | The settings source that enables it: &#x60;managed&#x60;, &#x60;user&#x60;, &#x60;project&#x60; or &#x60;local&#x60;. | [optional] 

## Methods

### NewMcpConfigPlugin

`func NewMcpConfigPlugin(name string, ) *McpConfigPlugin`

NewMcpConfigPlugin instantiates a new McpConfigPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpConfigPluginWithDefaults

`func NewMcpConfigPluginWithDefaults() *McpConfigPlugin`

NewMcpConfigPluginWithDefaults instantiates a new McpConfigPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *McpConfigPlugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpConfigPlugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpConfigPlugin) SetName(v string)`

SetName sets Name field to given value.


### GetMarketplace

`func (o *McpConfigPlugin) GetMarketplace() string`

GetMarketplace returns the Marketplace field if non-nil, zero value otherwise.

### GetMarketplaceOk

`func (o *McpConfigPlugin) GetMarketplaceOk() (*string, bool)`

GetMarketplaceOk returns a tuple with the Marketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplace

`func (o *McpConfigPlugin) SetMarketplace(v string)`

SetMarketplace sets Marketplace field to given value.

### HasMarketplace

`func (o *McpConfigPlugin) HasMarketplace() bool`

HasMarketplace returns a boolean if a field has been set.

### GetVersion

`func (o *McpConfigPlugin) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *McpConfigPlugin) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *McpConfigPlugin) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *McpConfigPlugin) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetScope

`func (o *McpConfigPlugin) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *McpConfigPlugin) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *McpConfigPlugin) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *McpConfigPlugin) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


