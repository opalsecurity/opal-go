# EntityAdminFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerIDs** | **[]string** | The owner (group) UUIDs to match entities against. | 
**Not** | Pointer to **bool** | Invert the match — return resources/groups NOT owned by the given owners. | [optional] 

## Methods

### NewEntityAdminFilter

`func NewEntityAdminFilter(ownerIDs []string, ) *EntityAdminFilter`

NewEntityAdminFilter instantiates a new EntityAdminFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityAdminFilterWithDefaults

`func NewEntityAdminFilterWithDefaults() *EntityAdminFilter`

NewEntityAdminFilterWithDefaults instantiates a new EntityAdminFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerIDs

`func (o *EntityAdminFilter) GetOwnerIDs() []string`

GetOwnerIDs returns the OwnerIDs field if non-nil, zero value otherwise.

### GetOwnerIDsOk

`func (o *EntityAdminFilter) GetOwnerIDsOk() (*[]string, bool)`

GetOwnerIDsOk returns a tuple with the OwnerIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIDs

`func (o *EntityAdminFilter) SetOwnerIDs(v []string)`

SetOwnerIDs sets OwnerIDs field to given value.


### GetNot

`func (o *EntityAdminFilter) GetNot() bool`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *EntityAdminFilter) GetNotOk() (*bool, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *EntityAdminFilter) SetNot(v bool)`

SetNot sets Not field to given value.

### HasNot

`func (o *EntityAdminFilter) HasNot() bool`

HasNot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


