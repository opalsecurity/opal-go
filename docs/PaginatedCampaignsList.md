# PaginatedCampaignsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** | The cursor with which to continue pagination if additional result pages exist. | [optional] 
**Previous** | Pointer to **string** | The cursor used to obtain the current result page. | [optional] 
**Results** | [**[]Campaign**](Campaign.md) |  | 

## Methods

### NewPaginatedCampaignsList

`func NewPaginatedCampaignsList(results []Campaign, ) *PaginatedCampaignsList`

NewPaginatedCampaignsList instantiates a new PaginatedCampaignsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedCampaignsListWithDefaults

`func NewPaginatedCampaignsListWithDefaults() *PaginatedCampaignsList`

NewPaginatedCampaignsListWithDefaults instantiates a new PaginatedCampaignsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginatedCampaignsList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedCampaignsList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedCampaignsList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedCampaignsList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *PaginatedCampaignsList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedCampaignsList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedCampaignsList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedCampaignsList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedCampaignsList) GetResults() []Campaign`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedCampaignsList) GetResultsOk() (*[]Campaign, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedCampaignsList) SetResults(v []Campaign)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


