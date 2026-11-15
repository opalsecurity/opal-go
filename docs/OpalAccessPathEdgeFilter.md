# OpalAccessPathEdgeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectOnly** | Pointer to **bool** | When true, only return direct (depth-1) principal-to-entitlement edges. | [optional] 
**AccessDurationType** | Pointer to **string** | Constrain results by whether the terminal access expires. | [optional] 

## Methods

### NewOpalAccessPathEdgeFilter

`func NewOpalAccessPathEdgeFilter() *OpalAccessPathEdgeFilter`

NewOpalAccessPathEdgeFilter instantiates a new OpalAccessPathEdgeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalAccessPathEdgeFilterWithDefaults

`func NewOpalAccessPathEdgeFilterWithDefaults() *OpalAccessPathEdgeFilter`

NewOpalAccessPathEdgeFilterWithDefaults instantiates a new OpalAccessPathEdgeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectOnly

`func (o *OpalAccessPathEdgeFilter) GetDirectOnly() bool`

GetDirectOnly returns the DirectOnly field if non-nil, zero value otherwise.

### GetDirectOnlyOk

`func (o *OpalAccessPathEdgeFilter) GetDirectOnlyOk() (*bool, bool)`

GetDirectOnlyOk returns a tuple with the DirectOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectOnly

`func (o *OpalAccessPathEdgeFilter) SetDirectOnly(v bool)`

SetDirectOnly sets DirectOnly field to given value.

### HasDirectOnly

`func (o *OpalAccessPathEdgeFilter) HasDirectOnly() bool`

HasDirectOnly returns a boolean if a field has been set.

### GetAccessDurationType

`func (o *OpalAccessPathEdgeFilter) GetAccessDurationType() string`

GetAccessDurationType returns the AccessDurationType field if non-nil, zero value otherwise.

### GetAccessDurationTypeOk

`func (o *OpalAccessPathEdgeFilter) GetAccessDurationTypeOk() (*string, bool)`

GetAccessDurationTypeOk returns a tuple with the AccessDurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDurationType

`func (o *OpalAccessPathEdgeFilter) SetAccessDurationType(v string)`

SetAccessDurationType sets AccessDurationType field to given value.

### HasAccessDurationType

`func (o *OpalAccessPathEdgeFilter) HasAccessDurationType() bool`

HasAccessDurationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


