# PaginatedBundleResourceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Previous** | Pointer to **string** | The cursor used to obtain the current result page. | [optional] 
**Next** | Pointer to **string** | The cursor with which to continue pagination if additional result pages exist. | [optional] 
**TotalCount** | Pointer to **int32** | The total number of items in the result set. | [optional] 
**BundleResources** | [**[]BundleResource**](BundleResource.md) |  | 

## Methods

### NewPaginatedBundleResourceList

`func NewPaginatedBundleResourceList(bundleResources []BundleResource, ) *PaginatedBundleResourceList`

NewPaginatedBundleResourceList instantiates a new PaginatedBundleResourceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedBundleResourceListWithDefaults

`func NewPaginatedBundleResourceListWithDefaults() *PaginatedBundleResourceList`

NewPaginatedBundleResourceListWithDefaults instantiates a new PaginatedBundleResourceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrevious

`func (o *PaginatedBundleResourceList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedBundleResourceList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedBundleResourceList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedBundleResourceList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedBundleResourceList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedBundleResourceList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedBundleResourceList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedBundleResourceList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetTotalCount

`func (o *PaginatedBundleResourceList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedBundleResourceList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedBundleResourceList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedBundleResourceList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetBundleResources

`func (o *PaginatedBundleResourceList) GetBundleResources() []BundleResource`

GetBundleResources returns the BundleResources field if non-nil, zero value otherwise.

### GetBundleResourcesOk

`func (o *PaginatedBundleResourceList) GetBundleResourcesOk() (*[]BundleResource, bool)`

GetBundleResourcesOk returns a tuple with the BundleResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleResources

`func (o *PaginatedBundleResourceList) SetBundleResources(v []BundleResource)`

SetBundleResources sets BundleResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


