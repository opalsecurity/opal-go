# PaginatedCampaignItemsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** | The cursor with which to continue pagination if additional result pages exist. | [optional] 
**Previous** | Pointer to **string** | The cursor used to obtain the current result page. | [optional] 
**Results** | [**[]CampaignItem**](CampaignItem.md) |  | 
**TotalCount** | **int64** | Total number of items matching the filter (across all pages). | 

## Methods

### NewPaginatedCampaignItemsList

`func NewPaginatedCampaignItemsList(results []CampaignItem, totalCount int64, ) *PaginatedCampaignItemsList`

NewPaginatedCampaignItemsList instantiates a new PaginatedCampaignItemsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedCampaignItemsListWithDefaults

`func NewPaginatedCampaignItemsListWithDefaults() *PaginatedCampaignItemsList`

NewPaginatedCampaignItemsListWithDefaults instantiates a new PaginatedCampaignItemsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginatedCampaignItemsList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedCampaignItemsList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedCampaignItemsList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedCampaignItemsList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *PaginatedCampaignItemsList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedCampaignItemsList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedCampaignItemsList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedCampaignItemsList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedCampaignItemsList) GetResults() []CampaignItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedCampaignItemsList) GetResultsOk() (*[]CampaignItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedCampaignItemsList) SetResults(v []CampaignItem)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *PaginatedCampaignItemsList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedCampaignItemsList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedCampaignItemsList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


