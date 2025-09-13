# RequestReviewerStages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevelName** | Pointer to **string** | The name of the access level requested. | [optional] 
**AccessLevelRemoteId** | Pointer to **string** | The ID of the access level requested on the remote system. | [optional] 
**ItemName** | **string** | The name of the requested item | 
**ItemId** | **string** | The ID of the resource requested. | 
**Stages** | [**[]RequestStage**](RequestStage.md) | The stages of review for this request | 

## Methods

### NewRequestReviewerStages

`func NewRequestReviewerStages(itemName string, itemId string, stages []RequestStage, ) *RequestReviewerStages`

NewRequestReviewerStages instantiates a new RequestReviewerStages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestReviewerStagesWithDefaults

`func NewRequestReviewerStagesWithDefaults() *RequestReviewerStages`

NewRequestReviewerStagesWithDefaults instantiates a new RequestReviewerStages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevelName

`func (o *RequestReviewerStages) GetAccessLevelName() string`

GetAccessLevelName returns the AccessLevelName field if non-nil, zero value otherwise.

### GetAccessLevelNameOk

`func (o *RequestReviewerStages) GetAccessLevelNameOk() (*string, bool)`

GetAccessLevelNameOk returns a tuple with the AccessLevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelName

`func (o *RequestReviewerStages) SetAccessLevelName(v string)`

SetAccessLevelName sets AccessLevelName field to given value.

### HasAccessLevelName

`func (o *RequestReviewerStages) HasAccessLevelName() bool`

HasAccessLevelName returns a boolean if a field has been set.

### GetAccessLevelRemoteId

`func (o *RequestReviewerStages) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *RequestReviewerStages) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *RequestReviewerStages) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *RequestReviewerStages) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.

### GetItemName

`func (o *RequestReviewerStages) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *RequestReviewerStages) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *RequestReviewerStages) SetItemName(v string)`

SetItemName sets ItemName field to given value.


### GetItemId

`func (o *RequestReviewerStages) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *RequestReviewerStages) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *RequestReviewerStages) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetStages

`func (o *RequestReviewerStages) GetStages() []RequestStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *RequestReviewerStages) GetStagesOk() (*[]RequestStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *RequestReviewerStages) SetStages(v []RequestStage)`

SetStages sets Stages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


