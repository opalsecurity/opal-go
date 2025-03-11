# PaginatedBundleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Previous** | Pointer to **string** | The cursor used to obtain the current result page. | [optional] 
**Next** | Pointer to **string** | The cursor with which to continue pagination if additional result pages exist. | [optional] 
**TotalCount** | Pointer to **int32** | The total number of items in the result set. | [optional] 
**Bundles** | [**[]Bundle**](Bundle.md) |  | 

## Methods

### NewPaginatedBundleList

`func NewPaginatedBundleList(bundles []Bundle, ) *PaginatedBundleList`

NewPaginatedBundleList instantiates a new PaginatedBundleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedBundleListWithDefaults

`func NewPaginatedBundleListWithDefaults() *PaginatedBundleList`

NewPaginatedBundleListWithDefaults instantiates a new PaginatedBundleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrevious

`func (o *PaginatedBundleList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedBundleList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedBundleList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedBundleList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedBundleList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedBundleList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedBundleList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedBundleList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetTotalCount

`func (o *PaginatedBundleList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedBundleList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedBundleList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedBundleList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetBundles

`func (o *PaginatedBundleList) GetBundles() []Bundle`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *PaginatedBundleList) GetBundlesOk() (*[]Bundle, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *PaginatedBundleList) SetBundles(v []Bundle)`

SetBundles sets Bundles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


