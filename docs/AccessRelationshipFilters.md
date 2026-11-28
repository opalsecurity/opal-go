# AccessRelationshipFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAccessibleBy** | Pointer to [**AccessEntityFilters**](AccessEntityFilters.md) | Inbound-edge filter. The returned node must be accessible by at least one entity matching this filter. | [optional] 
**HasAccessTo** | Pointer to [**AccessEntityFilters**](AccessEntityFilters.md) | Outbound-edge filter. The returned node must have access to at least one entity matching this filter. | [optional] 
**DirectOnly** | Pointer to **bool** | When true, &#x60;hasAccessTo&#x60; / &#x60;isAccessibleBy&#x60; match only direct role assignments (no nested group membership expansion).  | [optional] 
**AllOf** | Pointer to [**[]AccessRelationshipFilters**](AccessRelationshipFilters.md) | Nested access filters that must all match (logical AND).  | [optional] 
**AnyOf** | Pointer to [**[]AccessRelationshipFilters**](AccessRelationshipFilters.md) | Nested access filters where at least one must match (logical OR).  | [optional] 
**Not** | Pointer to **map[string]interface{}** | Excludes nodes matching the embedded access filter (logical NOT).  | [optional] 

## Methods

### NewAccessRelationshipFilters

`func NewAccessRelationshipFilters() *AccessRelationshipFilters`

NewAccessRelationshipFilters instantiates a new AccessRelationshipFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRelationshipFiltersWithDefaults

`func NewAccessRelationshipFiltersWithDefaults() *AccessRelationshipFilters`

NewAccessRelationshipFiltersWithDefaults instantiates a new AccessRelationshipFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAccessibleBy

`func (o *AccessRelationshipFilters) GetIsAccessibleBy() AccessEntityFilters`

GetIsAccessibleBy returns the IsAccessibleBy field if non-nil, zero value otherwise.

### GetIsAccessibleByOk

`func (o *AccessRelationshipFilters) GetIsAccessibleByOk() (*AccessEntityFilters, bool)`

GetIsAccessibleByOk returns a tuple with the IsAccessibleBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessibleBy

`func (o *AccessRelationshipFilters) SetIsAccessibleBy(v AccessEntityFilters)`

SetIsAccessibleBy sets IsAccessibleBy field to given value.

### HasIsAccessibleBy

`func (o *AccessRelationshipFilters) HasIsAccessibleBy() bool`

HasIsAccessibleBy returns a boolean if a field has been set.

### GetHasAccessTo

`func (o *AccessRelationshipFilters) GetHasAccessTo() AccessEntityFilters`

GetHasAccessTo returns the HasAccessTo field if non-nil, zero value otherwise.

### GetHasAccessToOk

`func (o *AccessRelationshipFilters) GetHasAccessToOk() (*AccessEntityFilters, bool)`

GetHasAccessToOk returns a tuple with the HasAccessTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessTo

`func (o *AccessRelationshipFilters) SetHasAccessTo(v AccessEntityFilters)`

SetHasAccessTo sets HasAccessTo field to given value.

### HasHasAccessTo

`func (o *AccessRelationshipFilters) HasHasAccessTo() bool`

HasHasAccessTo returns a boolean if a field has been set.

### GetDirectOnly

`func (o *AccessRelationshipFilters) GetDirectOnly() bool`

GetDirectOnly returns the DirectOnly field if non-nil, zero value otherwise.

### GetDirectOnlyOk

`func (o *AccessRelationshipFilters) GetDirectOnlyOk() (*bool, bool)`

GetDirectOnlyOk returns a tuple with the DirectOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectOnly

`func (o *AccessRelationshipFilters) SetDirectOnly(v bool)`

SetDirectOnly sets DirectOnly field to given value.

### HasDirectOnly

`func (o *AccessRelationshipFilters) HasDirectOnly() bool`

HasDirectOnly returns a boolean if a field has been set.

### GetAllOf

`func (o *AccessRelationshipFilters) GetAllOf() []AccessRelationshipFilters`

GetAllOf returns the AllOf field if non-nil, zero value otherwise.

### GetAllOfOk

`func (o *AccessRelationshipFilters) GetAllOfOk() (*[]AccessRelationshipFilters, bool)`

GetAllOfOk returns a tuple with the AllOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllOf

`func (o *AccessRelationshipFilters) SetAllOf(v []AccessRelationshipFilters)`

SetAllOf sets AllOf field to given value.

### HasAllOf

`func (o *AccessRelationshipFilters) HasAllOf() bool`

HasAllOf returns a boolean if a field has been set.

### GetAnyOf

`func (o *AccessRelationshipFilters) GetAnyOf() []AccessRelationshipFilters`

GetAnyOf returns the AnyOf field if non-nil, zero value otherwise.

### GetAnyOfOk

`func (o *AccessRelationshipFilters) GetAnyOfOk() (*[]AccessRelationshipFilters, bool)`

GetAnyOfOk returns a tuple with the AnyOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyOf

`func (o *AccessRelationshipFilters) SetAnyOf(v []AccessRelationshipFilters)`

SetAnyOf sets AnyOf field to given value.

### HasAnyOf

`func (o *AccessRelationshipFilters) HasAnyOf() bool`

HasAnyOf returns a boolean if a field has been set.

### GetNot

`func (o *AccessRelationshipFilters) GetNot() map[string]interface{}`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *AccessRelationshipFilters) GetNotOk() (*map[string]interface{}, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *AccessRelationshipFilters) SetNot(v map[string]interface{})`

SetNot sets Not field to given value.

### HasNot

`func (o *AccessRelationshipFilters) HasNot() bool`

HasNot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


