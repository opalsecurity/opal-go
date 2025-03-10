# Bundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the bundle. | 
**Name** | Pointer to **string** | The name of the bundle. | [optional] 
**Description** | Pointer to **string** | The description of the bundle. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The creation timestamp of the bundle, in ISO 8601 format | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The last updated timestamp of the bundle, in ISO 8601 format | [optional] 
**AdminOwnerId** | **string** | The ID of the owner of the bundle. | 
**TotalNumItems** | Pointer to **int32** | The total number of items in the bundle. | [optional] 
**TotalNumResources** | Pointer to **int32** | The total number of resources in the bundle. | [optional] 
**TotalNumGroups** | Pointer to **int32** | The total number of groups in the bundle. | [optional] 

## Methods

### NewBundle

`func NewBundle(id string, adminOwnerId string, ) *Bundle`

NewBundle instantiates a new Bundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleWithDefaults

`func NewBundleWithDefaults() *Bundle`

NewBundleWithDefaults instantiates a new Bundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bundle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bundle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bundle) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Bundle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bundle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bundle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Bundle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Bundle) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Bundle) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Bundle) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Bundle) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Bundle) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Bundle) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Bundle) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Bundle) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Bundle) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Bundle) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Bundle) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Bundle) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAdminOwnerId

`func (o *Bundle) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *Bundle) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *Bundle) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.


### GetTotalNumItems

`func (o *Bundle) GetTotalNumItems() int32`

GetTotalNumItems returns the TotalNumItems field if non-nil, zero value otherwise.

### GetTotalNumItemsOk

`func (o *Bundle) GetTotalNumItemsOk() (*int32, bool)`

GetTotalNumItemsOk returns a tuple with the TotalNumItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumItems

`func (o *Bundle) SetTotalNumItems(v int32)`

SetTotalNumItems sets TotalNumItems field to given value.

### HasTotalNumItems

`func (o *Bundle) HasTotalNumItems() bool`

HasTotalNumItems returns a boolean if a field has been set.

### GetTotalNumResources

`func (o *Bundle) GetTotalNumResources() int32`

GetTotalNumResources returns the TotalNumResources field if non-nil, zero value otherwise.

### GetTotalNumResourcesOk

`func (o *Bundle) GetTotalNumResourcesOk() (*int32, bool)`

GetTotalNumResourcesOk returns a tuple with the TotalNumResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumResources

`func (o *Bundle) SetTotalNumResources(v int32)`

SetTotalNumResources sets TotalNumResources field to given value.

### HasTotalNumResources

`func (o *Bundle) HasTotalNumResources() bool`

HasTotalNumResources returns a boolean if a field has been set.

### GetTotalNumGroups

`func (o *Bundle) GetTotalNumGroups() int32`

GetTotalNumGroups returns the TotalNumGroups field if non-nil, zero value otherwise.

### GetTotalNumGroupsOk

`func (o *Bundle) GetTotalNumGroupsOk() (*int32, bool)`

GetTotalNumGroupsOk returns a tuple with the TotalNumGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumGroups

`func (o *Bundle) SetTotalNumGroups(v int32)`

SetTotalNumGroups sets TotalNumGroups field to given value.

### HasTotalNumGroups

`func (o *Bundle) HasTotalNumGroups() bool`

HasTotalNumGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


