# PaginatedBundleGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Previous** | Pointer to **string** | The cursor used to obtain the current result page. | [optional] 
**Next** | Pointer to **string** | The cursor with which to continue pagination if additional result pages exist. | [optional] 
**TotalCount** | Pointer to **int32** | The total number of items in the result set. | [optional] 
**BundleGroups** | [**[]BundleGroup**](BundleGroup.md) |  | 

## Methods

### NewPaginatedBundleGroupList

`func NewPaginatedBundleGroupList(bundleGroups []BundleGroup, ) *PaginatedBundleGroupList`

NewPaginatedBundleGroupList instantiates a new PaginatedBundleGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedBundleGroupListWithDefaults

`func NewPaginatedBundleGroupListWithDefaults() *PaginatedBundleGroupList`

NewPaginatedBundleGroupListWithDefaults instantiates a new PaginatedBundleGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrevious

`func (o *PaginatedBundleGroupList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedBundleGroupList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedBundleGroupList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedBundleGroupList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedBundleGroupList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedBundleGroupList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedBundleGroupList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedBundleGroupList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetTotalCount

`func (o *PaginatedBundleGroupList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedBundleGroupList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedBundleGroupList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedBundleGroupList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetBundleGroups

`func (o *PaginatedBundleGroupList) GetBundleGroups() []BundleGroup`

GetBundleGroups returns the BundleGroups field if non-nil, zero value otherwise.

### GetBundleGroupsOk

`func (o *PaginatedBundleGroupList) GetBundleGroupsOk() (*[]BundleGroup, bool)`

GetBundleGroupsOk returns a tuple with the BundleGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleGroups

`func (o *PaginatedBundleGroupList) SetBundleGroups(v []BundleGroup)`

SetBundleGroups sets BundleGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


