# McpConfigPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | &#x60;managed&#x60;, &#x60;user&#x60;, &#x60;project&#x60; or &#x60;local&#x60;. | [optional] 
**Kind** | **string** | &#x60;allow&#x60;, &#x60;ask&#x60;, &#x60;deny&#x60; or &#x60;default_mode&#x60;. | 
**Value** | **string** | The rule, such as &#x60;Bash(npm run test *)&#x60;, or the mode, such as &#x60;acceptEdits&#x60;. | 

## Methods

### NewMcpConfigPermission

`func NewMcpConfigPermission(kind string, value string, ) *McpConfigPermission`

NewMcpConfigPermission instantiates a new McpConfigPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpConfigPermissionWithDefaults

`func NewMcpConfigPermissionWithDefaults() *McpConfigPermission`

NewMcpConfigPermissionWithDefaults instantiates a new McpConfigPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *McpConfigPermission) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *McpConfigPermission) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *McpConfigPermission) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *McpConfigPermission) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetKind

`func (o *McpConfigPermission) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *McpConfigPermission) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *McpConfigPermission) SetKind(v string)`

SetKind sets Kind field to given value.


### GetValue

`func (o *McpConfigPermission) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *McpConfigPermission) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *McpConfigPermission) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


