# BatchRequestabilityCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | The ID of the user the access would be for. | 
**RequesterId** | Pointer to **string** | Optional ID of the person filling in the form when different from user_id. Adds the on-behalf check evaluated as that requester, following EnforceRequestOnBehalfOfUser.  | [optional] 
**Ids** | **[]string** | The asset IDs to check. Duplicates are dropped. | 

## Methods

### NewBatchRequestabilityCheckRequest

`func NewBatchRequestabilityCheckRequest(userId string, ids []string, ) *BatchRequestabilityCheckRequest`

NewBatchRequestabilityCheckRequest instantiates a new BatchRequestabilityCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRequestabilityCheckRequestWithDefaults

`func NewBatchRequestabilityCheckRequestWithDefaults() *BatchRequestabilityCheckRequest`

NewBatchRequestabilityCheckRequestWithDefaults instantiates a new BatchRequestabilityCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *BatchRequestabilityCheckRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BatchRequestabilityCheckRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BatchRequestabilityCheckRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetRequesterId

`func (o *BatchRequestabilityCheckRequest) GetRequesterId() string`

GetRequesterId returns the RequesterId field if non-nil, zero value otherwise.

### GetRequesterIdOk

`func (o *BatchRequestabilityCheckRequest) GetRequesterIdOk() (*string, bool)`

GetRequesterIdOk returns a tuple with the RequesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterId

`func (o *BatchRequestabilityCheckRequest) SetRequesterId(v string)`

SetRequesterId sets RequesterId field to given value.

### HasRequesterId

`func (o *BatchRequestabilityCheckRequest) HasRequesterId() bool`

HasRequesterId returns a boolean if a field has been set.

### GetIds

`func (o *BatchRequestabilityCheckRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BatchRequestabilityCheckRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BatchRequestabilityCheckRequest) SetIds(v []string)`

SetIds sets Ids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


