# UARScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupVisibility** | Pointer to **string** | Specifies what users can see during an Access Review | [optional] 
**Users** | Pointer to **[]string** | The access review will only include the following users. If any users are selected, any entity filters will be applied to only the entities that the selected users have access to. | [optional] 
**FilterOperator** | Pointer to **string** | Specifies whether entities must match all (AND) or any (OR) of the filters. | [optional] 
**Entities** | Pointer to **[]string** | This access review will include resources and groups with ids in the given strings. | [optional] 
**Apps** | Pointer to **[]string** | This access review will include items in the specified applications | [optional] 
**Admins** | Pointer to **[]string** | This access review will include resources and groups who are owned by one of the owners corresponding to the given IDs. | [optional] 
**GroupTypes** | Pointer to [**[]GroupTypeEnum**](GroupTypeEnum.md) | This access review will include items of the specified group types | [optional] 
**ResourceTypes** | Pointer to [**[]ResourceTypeEnum**](ResourceTypeEnum.md) | This access review will include items of the specified resource types | [optional] 
**IncludeGroupBindings** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to [**[]TagFilter**](TagFilter.md) | This access review will include resources and groups who are tagged with one of the given tags. | [optional] 
**Names** | Pointer to **[]string** | This access review will include resources and groups whose name contains one of the given strings. | [optional] 

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

### GetGroupVisibility

`func (o *UARScope) GetGroupVisibility() string`

GetGroupVisibility returns the GroupVisibility field if non-nil, zero value otherwise.

### GetGroupVisibilityOk

`func (o *UARScope) GetGroupVisibilityOk() (*string, bool)`

GetGroupVisibilityOk returns a tuple with the GroupVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupVisibility

`func (o *UARScope) SetGroupVisibility(v string)`

SetGroupVisibility sets GroupVisibility field to given value.

### HasGroupVisibility

`func (o *UARScope) HasGroupVisibility() bool`

HasGroupVisibility returns a boolean if a field has been set.

### GetUsers

`func (o *UARScope) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UARScope) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UARScope) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UARScope) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetFilterOperator

`func (o *UARScope) GetFilterOperator() string`

GetFilterOperator returns the FilterOperator field if non-nil, zero value otherwise.

### GetFilterOperatorOk

`func (o *UARScope) GetFilterOperatorOk() (*string, bool)`

GetFilterOperatorOk returns a tuple with the FilterOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterOperator

`func (o *UARScope) SetFilterOperator(v string)`

SetFilterOperator sets FilterOperator field to given value.

### HasFilterOperator

`func (o *UARScope) HasFilterOperator() bool`

HasFilterOperator returns a boolean if a field has been set.

### GetEntities

`func (o *UARScope) GetEntities() []string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *UARScope) GetEntitiesOk() (*[]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *UARScope) SetEntities(v []string)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *UARScope) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApps

`func (o *UARScope) GetApps() []string`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *UARScope) GetAppsOk() (*[]string, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *UARScope) SetApps(v []string)`

SetApps sets Apps field to given value.

### HasApps

`func (o *UARScope) HasApps() bool`

HasApps returns a boolean if a field has been set.

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

### GetGroupTypes

`func (o *UARScope) GetGroupTypes() []GroupTypeEnum`

GetGroupTypes returns the GroupTypes field if non-nil, zero value otherwise.

### GetGroupTypesOk

`func (o *UARScope) GetGroupTypesOk() (*[]GroupTypeEnum, bool)`

GetGroupTypesOk returns a tuple with the GroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypes

`func (o *UARScope) SetGroupTypes(v []GroupTypeEnum)`

SetGroupTypes sets GroupTypes field to given value.

### HasGroupTypes

`func (o *UARScope) HasGroupTypes() bool`

HasGroupTypes returns a boolean if a field has been set.

### GetResourceTypes

`func (o *UARScope) GetResourceTypes() []ResourceTypeEnum`

GetResourceTypes returns the ResourceTypes field if non-nil, zero value otherwise.

### GetResourceTypesOk

`func (o *UARScope) GetResourceTypesOk() (*[]ResourceTypeEnum, bool)`

GetResourceTypesOk returns a tuple with the ResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypes

`func (o *UARScope) SetResourceTypes(v []ResourceTypeEnum)`

SetResourceTypes sets ResourceTypes field to given value.

### HasResourceTypes

`func (o *UARScope) HasResourceTypes() bool`

HasResourceTypes returns a boolean if a field has been set.

### GetIncludeGroupBindings

`func (o *UARScope) GetIncludeGroupBindings() bool`

GetIncludeGroupBindings returns the IncludeGroupBindings field if non-nil, zero value otherwise.

### GetIncludeGroupBindingsOk

`func (o *UARScope) GetIncludeGroupBindingsOk() (*bool, bool)`

GetIncludeGroupBindingsOk returns a tuple with the IncludeGroupBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeGroupBindings

`func (o *UARScope) SetIncludeGroupBindings(v bool)`

SetIncludeGroupBindings sets IncludeGroupBindings field to given value.

### HasIncludeGroupBindings

`func (o *UARScope) HasIncludeGroupBindings() bool`

HasIncludeGroupBindings returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


