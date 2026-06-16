# AccessRelationshipFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAccessibleBy** | Pointer to [**AccessEntityFilters**](AccessEntityFilters.md) | Inbound-edge filter. The returned node must be accessible by at least one entity matching this filter. | [optional] 
**HasAccessTo** | Pointer to [**AccessEntityFilters**](AccessEntityFilters.md) | Outbound-edge filter. The returned node must have access to at least one entity matching this filter. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


