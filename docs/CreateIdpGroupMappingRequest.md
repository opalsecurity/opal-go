# CreateIdpGroupMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Optional alias for the group mapping | [optional] 
**HiddenFromEndUser** | Pointer to **bool** | Whether this mapping should be hidden from end users. - **New mappings**: If not provided, defaults to &#x60;false&#x60; - **Existing mappings**: If not provided, existing value is preserved (no change) - **Explicit values**: If provided, value is updated to the specified boolean  | [optional] 

## Methods

### NewCreateIdpGroupMappingRequest

`func NewCreateIdpGroupMappingRequest() *CreateIdpGroupMappingRequest`

NewCreateIdpGroupMappingRequest instantiates a new CreateIdpGroupMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIdpGroupMappingRequestWithDefaults

`func NewCreateIdpGroupMappingRequestWithDefaults() *CreateIdpGroupMappingRequest`

NewCreateIdpGroupMappingRequestWithDefaults instantiates a new CreateIdpGroupMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *CreateIdpGroupMappingRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CreateIdpGroupMappingRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CreateIdpGroupMappingRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *CreateIdpGroupMappingRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetHiddenFromEndUser

`func (o *CreateIdpGroupMappingRequest) GetHiddenFromEndUser() bool`

GetHiddenFromEndUser returns the HiddenFromEndUser field if non-nil, zero value otherwise.

### GetHiddenFromEndUserOk

`func (o *CreateIdpGroupMappingRequest) GetHiddenFromEndUserOk() (*bool, bool)`

GetHiddenFromEndUserOk returns a tuple with the HiddenFromEndUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromEndUser

`func (o *CreateIdpGroupMappingRequest) SetHiddenFromEndUser(v bool)`

SetHiddenFromEndUser sets HiddenFromEndUser field to given value.

### HasHiddenFromEndUser

`func (o *CreateIdpGroupMappingRequest) HasHiddenFromEndUser() bool`

HasHiddenFromEndUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


