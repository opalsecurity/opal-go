# OpalNodeQueryBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeFilters** | Pointer to [**AccessEntityFilters**](AccessEntityFilters.md) |  | [optional] 
**AccessFilters** | Pointer to [**AccessRelationshipFilters**](AccessRelationshipFilters.md) |  | [optional] 

## Methods

### NewOpalNodeQueryBody

`func NewOpalNodeQueryBody() *OpalNodeQueryBody`

NewOpalNodeQueryBody instantiates a new OpalNodeQueryBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalNodeQueryBodyWithDefaults

`func NewOpalNodeQueryBodyWithDefaults() *OpalNodeQueryBody`

NewOpalNodeQueryBodyWithDefaults instantiates a new OpalNodeQueryBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeFilters

`func (o *OpalNodeQueryBody) GetNodeFilters() AccessEntityFilters`

GetNodeFilters returns the NodeFilters field if non-nil, zero value otherwise.

### GetNodeFiltersOk

`func (o *OpalNodeQueryBody) GetNodeFiltersOk() (*AccessEntityFilters, bool)`

GetNodeFiltersOk returns a tuple with the NodeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFilters

`func (o *OpalNodeQueryBody) SetNodeFilters(v AccessEntityFilters)`

SetNodeFilters sets NodeFilters field to given value.

### HasNodeFilters

`func (o *OpalNodeQueryBody) HasNodeFilters() bool`

HasNodeFilters returns a boolean if a field has been set.

### GetAccessFilters

`func (o *OpalNodeQueryBody) GetAccessFilters() AccessRelationshipFilters`

GetAccessFilters returns the AccessFilters field if non-nil, zero value otherwise.

### GetAccessFiltersOk

`func (o *OpalNodeQueryBody) GetAccessFiltersOk() (*AccessRelationshipFilters, bool)`

GetAccessFiltersOk returns a tuple with the AccessFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessFilters

`func (o *OpalNodeQueryBody) SetAccessFilters(v AccessRelationshipFilters)`

SetAccessFilters sets AccessFilters field to given value.

### HasAccessFilters

`func (o *OpalNodeQueryBody) HasAccessFilters() bool`

HasAccessFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


