# RequestItemStages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedRoleName** | Pointer to **string** | The name of the requested role | [optional] 
**RequestedItemName** | **string** | The name of the requested item | 
**Stages** | [**[]RequestStage**](RequestStage.md) | The stages of review for this request | 

## Methods

### NewRequestItemStages

`func NewRequestItemStages(requestedItemName string, stages []RequestStage, ) *RequestItemStages`

NewRequestItemStages instantiates a new RequestItemStages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestItemStagesWithDefaults

`func NewRequestItemStagesWithDefaults() *RequestItemStages`

NewRequestItemStagesWithDefaults instantiates a new RequestItemStages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedRoleName

`func (o *RequestItemStages) GetRequestedRoleName() string`

GetRequestedRoleName returns the RequestedRoleName field if non-nil, zero value otherwise.

### GetRequestedRoleNameOk

`func (o *RequestItemStages) GetRequestedRoleNameOk() (*string, bool)`

GetRequestedRoleNameOk returns a tuple with the RequestedRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedRoleName

`func (o *RequestItemStages) SetRequestedRoleName(v string)`

SetRequestedRoleName sets RequestedRoleName field to given value.

### HasRequestedRoleName

`func (o *RequestItemStages) HasRequestedRoleName() bool`

HasRequestedRoleName returns a boolean if a field has been set.

### GetRequestedItemName

`func (o *RequestItemStages) GetRequestedItemName() string`

GetRequestedItemName returns the RequestedItemName field if non-nil, zero value otherwise.

### GetRequestedItemNameOk

`func (o *RequestItemStages) GetRequestedItemNameOk() (*string, bool)`

GetRequestedItemNameOk returns a tuple with the RequestedItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedItemName

`func (o *RequestItemStages) SetRequestedItemName(v string)`

SetRequestedItemName sets RequestedItemName field to given value.


### GetStages

`func (o *RequestItemStages) GetStages() []RequestStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *RequestItemStages) GetStagesOk() (*[]RequestStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *RequestItemStages) SetStages(v []RequestStage)`

SetStages sets Stages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


