# McpConfigSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Source** | Pointer to **string** | &#x60;user&#x60;, &#x60;project&#x60;, &#x60;plugin&#x60; or &#x60;managed&#x60;. | [optional] 

## Methods

### NewMcpConfigSkill

`func NewMcpConfigSkill(name string, ) *McpConfigSkill`

NewMcpConfigSkill instantiates a new McpConfigSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpConfigSkillWithDefaults

`func NewMcpConfigSkillWithDefaults() *McpConfigSkill`

NewMcpConfigSkillWithDefaults instantiates a new McpConfigSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *McpConfigSkill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpConfigSkill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpConfigSkill) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *McpConfigSkill) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *McpConfigSkill) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *McpConfigSkill) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *McpConfigSkill) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


