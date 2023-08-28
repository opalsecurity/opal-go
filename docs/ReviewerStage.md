# ReviewerStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireManagerApproval** | **bool** | Whether this reviewer stage should require manager approval. | 
**Operator** | **string** | The operator of the reviewer stage. | 
**OwnerIds** | **[]string** |  | 

## Methods

### NewReviewerStage

`func NewReviewerStage(requireManagerApproval bool, operator string, ownerIds []string, ) *ReviewerStage`

NewReviewerStage instantiates a new ReviewerStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewerStageWithDefaults

`func NewReviewerStageWithDefaults() *ReviewerStage`

NewReviewerStageWithDefaults instantiates a new ReviewerStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequireManagerApproval

`func (o *ReviewerStage) GetRequireManagerApproval() bool`

GetRequireManagerApproval returns the RequireManagerApproval field if non-nil, zero value otherwise.

### GetRequireManagerApprovalOk

`func (o *ReviewerStage) GetRequireManagerApprovalOk() (*bool, bool)`

GetRequireManagerApprovalOk returns a tuple with the RequireManagerApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireManagerApproval

`func (o *ReviewerStage) SetRequireManagerApproval(v bool)`

SetRequireManagerApproval sets RequireManagerApproval field to given value.


### GetOperator

`func (o *ReviewerStage) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ReviewerStage) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ReviewerStage) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetOwnerIds

`func (o *ReviewerStage) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *ReviewerStage) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *ReviewerStage) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


