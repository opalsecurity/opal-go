# RequestStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stage** | **int32** | The stage number | 
**Operator** | [**ReviewStageOperator**](ReviewStageOperator.md) |  | 
**Reviewers** | [**[]RequestReviewer**](RequestReviewer.md) | The reviewers for this stage | 

## Methods

### NewRequestStage

`func NewRequestStage(stage int32, operator ReviewStageOperator, reviewers []RequestReviewer, ) *RequestStage`

NewRequestStage instantiates a new RequestStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestStageWithDefaults

`func NewRequestStageWithDefaults() *RequestStage`

NewRequestStageWithDefaults instantiates a new RequestStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStage

`func (o *RequestStage) GetStage() int32`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *RequestStage) GetStageOk() (*int32, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *RequestStage) SetStage(v int32)`

SetStage sets Stage field to given value.


### GetOperator

`func (o *RequestStage) GetOperator() ReviewStageOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RequestStage) GetOperatorOk() (*ReviewStageOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RequestStage) SetOperator(v ReviewStageOperator)`

SetOperator sets Operator field to given value.


### GetReviewers

`func (o *RequestStage) GetReviewers() []RequestReviewer`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *RequestStage) GetReviewersOk() (*[]RequestReviewer, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *RequestStage) SetReviewers(v []RequestReviewer)`

SetReviewers sets Reviewers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


