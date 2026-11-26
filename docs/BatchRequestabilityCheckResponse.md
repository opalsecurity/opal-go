# BatchRequestabilityCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | The ID of the user that was checked. | 
**Results** | [**[]BatchRequestabilityCheckResult**](BatchRequestabilityCheckResult.md) |  | 

## Methods

### NewBatchRequestabilityCheckResponse

`func NewBatchRequestabilityCheckResponse(userId string, results []BatchRequestabilityCheckResult, ) *BatchRequestabilityCheckResponse`

NewBatchRequestabilityCheckResponse instantiates a new BatchRequestabilityCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRequestabilityCheckResponseWithDefaults

`func NewBatchRequestabilityCheckResponseWithDefaults() *BatchRequestabilityCheckResponse`

NewBatchRequestabilityCheckResponseWithDefaults instantiates a new BatchRequestabilityCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *BatchRequestabilityCheckResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BatchRequestabilityCheckResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BatchRequestabilityCheckResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetResults

`func (o *BatchRequestabilityCheckResponse) GetResults() []BatchRequestabilityCheckResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchRequestabilityCheckResponse) GetResultsOk() (*[]BatchRequestabilityCheckResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchRequestabilityCheckResponse) SetResults(v []BatchRequestabilityCheckResult)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


