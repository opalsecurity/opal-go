# OpalAccessPathQueryBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalFilter** | Pointer to [**AccessEntityFilters**](AccessEntityFilters.md) |  | [optional] 
**EntitlementFilter** | Pointer to [**AccessEntityFilters**](AccessEntityFilters.md) |  | [optional] 
**AccessLevelRemoteIds** | Pointer to **[]string** | Filter by access-level remote IDs on the terminal edge. | [optional] 
**AccessLevelNames** | Pointer to **[]string** | Filter by access-level display names on the terminal edge. | [optional] 
**EdgeFilter** | Pointer to [**OpalAccessPathEdgeFilter**](OpalAccessPathEdgeFilter.md) |  | [optional] 

## Methods

### NewOpalAccessPathQueryBody

`func NewOpalAccessPathQueryBody() *OpalAccessPathQueryBody`

NewOpalAccessPathQueryBody instantiates a new OpalAccessPathQueryBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalAccessPathQueryBodyWithDefaults

`func NewOpalAccessPathQueryBodyWithDefaults() *OpalAccessPathQueryBody`

NewOpalAccessPathQueryBodyWithDefaults instantiates a new OpalAccessPathQueryBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalFilter

`func (o *OpalAccessPathQueryBody) GetPrincipalFilter() AccessEntityFilters`

GetPrincipalFilter returns the PrincipalFilter field if non-nil, zero value otherwise.

### GetPrincipalFilterOk

`func (o *OpalAccessPathQueryBody) GetPrincipalFilterOk() (*AccessEntityFilters, bool)`

GetPrincipalFilterOk returns a tuple with the PrincipalFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalFilter

`func (o *OpalAccessPathQueryBody) SetPrincipalFilter(v AccessEntityFilters)`

SetPrincipalFilter sets PrincipalFilter field to given value.

### HasPrincipalFilter

`func (o *OpalAccessPathQueryBody) HasPrincipalFilter() bool`

HasPrincipalFilter returns a boolean if a field has been set.

### GetEntitlementFilter

`func (o *OpalAccessPathQueryBody) GetEntitlementFilter() AccessEntityFilters`

GetEntitlementFilter returns the EntitlementFilter field if non-nil, zero value otherwise.

### GetEntitlementFilterOk

`func (o *OpalAccessPathQueryBody) GetEntitlementFilterOk() (*AccessEntityFilters, bool)`

GetEntitlementFilterOk returns a tuple with the EntitlementFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementFilter

`func (o *OpalAccessPathQueryBody) SetEntitlementFilter(v AccessEntityFilters)`

SetEntitlementFilter sets EntitlementFilter field to given value.

### HasEntitlementFilter

`func (o *OpalAccessPathQueryBody) HasEntitlementFilter() bool`

HasEntitlementFilter returns a boolean if a field has been set.

### GetAccessLevelRemoteIds

`func (o *OpalAccessPathQueryBody) GetAccessLevelRemoteIds() []string`

GetAccessLevelRemoteIds returns the AccessLevelRemoteIds field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdsOk

`func (o *OpalAccessPathQueryBody) GetAccessLevelRemoteIdsOk() (*[]string, bool)`

GetAccessLevelRemoteIdsOk returns a tuple with the AccessLevelRemoteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteIds

`func (o *OpalAccessPathQueryBody) SetAccessLevelRemoteIds(v []string)`

SetAccessLevelRemoteIds sets AccessLevelRemoteIds field to given value.

### HasAccessLevelRemoteIds

`func (o *OpalAccessPathQueryBody) HasAccessLevelRemoteIds() bool`

HasAccessLevelRemoteIds returns a boolean if a field has been set.

### GetAccessLevelNames

`func (o *OpalAccessPathQueryBody) GetAccessLevelNames() []string`

GetAccessLevelNames returns the AccessLevelNames field if non-nil, zero value otherwise.

### GetAccessLevelNamesOk

`func (o *OpalAccessPathQueryBody) GetAccessLevelNamesOk() (*[]string, bool)`

GetAccessLevelNamesOk returns a tuple with the AccessLevelNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelNames

`func (o *OpalAccessPathQueryBody) SetAccessLevelNames(v []string)`

SetAccessLevelNames sets AccessLevelNames field to given value.

### HasAccessLevelNames

`func (o *OpalAccessPathQueryBody) HasAccessLevelNames() bool`

HasAccessLevelNames returns a boolean if a field has been set.

### GetEdgeFilter

`func (o *OpalAccessPathQueryBody) GetEdgeFilter() OpalAccessPathEdgeFilter`

GetEdgeFilter returns the EdgeFilter field if non-nil, zero value otherwise.

### GetEdgeFilterOk

`func (o *OpalAccessPathQueryBody) GetEdgeFilterOk() (*OpalAccessPathEdgeFilter, bool)`

GetEdgeFilterOk returns a tuple with the EdgeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFilter

`func (o *OpalAccessPathQueryBody) SetEdgeFilter(v OpalAccessPathEdgeFilter)`

SetEdgeFilter sets EdgeFilter field to given value.

### HasEdgeFilter

`func (o *OpalAccessPathQueryBody) HasEdgeFilter() bool`

HasEdgeFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


