# IdpGroupMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The ID of the group. | 
**Alias** | Pointer to **string** | The alias of the group. | [optional] 
**HiddenFromEndUser** | **bool** | A bool representing whether or not the group is hidden from the end user. | 

## Methods

### NewIdpGroupMapping

`func NewIdpGroupMapping(groupId string, hiddenFromEndUser bool, ) *IdpGroupMapping`

NewIdpGroupMapping instantiates a new IdpGroupMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpGroupMappingWithDefaults

`func NewIdpGroupMappingWithDefaults() *IdpGroupMapping`

NewIdpGroupMappingWithDefaults instantiates a new IdpGroupMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *IdpGroupMapping) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *IdpGroupMapping) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *IdpGroupMapping) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetAlias

`func (o *IdpGroupMapping) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *IdpGroupMapping) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *IdpGroupMapping) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *IdpGroupMapping) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetHiddenFromEndUser

`func (o *IdpGroupMapping) GetHiddenFromEndUser() bool`

GetHiddenFromEndUser returns the HiddenFromEndUser field if non-nil, zero value otherwise.

### GetHiddenFromEndUserOk

`func (o *IdpGroupMapping) GetHiddenFromEndUserOk() (*bool, bool)`

GetHiddenFromEndUserOk returns a tuple with the HiddenFromEndUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromEndUser

`func (o *IdpGroupMapping) SetHiddenFromEndUser(v bool)`

SetHiddenFromEndUser sets HiddenFromEndUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


