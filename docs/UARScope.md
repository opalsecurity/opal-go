# UARScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]TagFilter**](TagFilter.md) | This access review will include resources and groups who are tagged with one of the given tags. | [optional] 
**Names** | Pointer to **[]string** | This access review will include resources and groups whose name contains one of the given strings. | [optional] 
**Admins** | Pointer to **[]string** | This access review will include resources and groups who are owned by one of the owners corresponding to the given IDs. | [optional] 

## Methods

### NewUARScope

`func NewUARScope() *UARScope`

NewUARScope instantiates a new UARScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUARScopeWithDefaults

`func NewUARScopeWithDefaults() *UARScope`

NewUARScopeWithDefaults instantiates a new UARScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *UARScope) GetTags() []TagFilter`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UARScope) GetTagsOk() (*[]TagFilter, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UARScope) SetTags(v []TagFilter)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UARScope) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNames

`func (o *UARScope) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *UARScope) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *UARScope) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *UARScope) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetAdmins

`func (o *UARScope) GetAdmins() []string`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *UARScope) GetAdminsOk() (*[]string, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *UARScope) SetAdmins(v []string)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *UARScope) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


