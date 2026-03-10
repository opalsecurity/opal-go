# ReviewerStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireManagerApproval** | **bool** | Whether this reviewer stage should require manager approval. | 
**RequireAdminApproval** | Pointer to **bool** | Whether this reviewer stage should require admin approval. | [optional] 
**Operator** | **string** | The operator of the reviewer stage. Admin and manager approval are also treated as reviewers. | 
**OwnerIds** | **[]string** | The IDs of owners assigned as reviewers for this stage. | 
**ServiceUserIds** | Pointer to **[]string** | The IDs of service users assigned as reviewers for this stage. | [optional] 

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


### GetRequireAdminApproval

`func (o *ReviewerStage) GetRequireAdminApproval() bool`

GetRequireAdminApproval returns the RequireAdminApproval field if non-nil, zero value otherwise.

### GetRequireAdminApprovalOk

`func (o *ReviewerStage) GetRequireAdminApprovalOk() (*bool, bool)`

GetRequireAdminApprovalOk returns a tuple with the RequireAdminApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAdminApproval

`func (o *ReviewerStage) SetRequireAdminApproval(v bool)`

SetRequireAdminApproval sets RequireAdminApproval field to given value.

### HasRequireAdminApproval

`func (o *ReviewerStage) HasRequireAdminApproval() bool`

HasRequireAdminApproval returns a boolean if a field has been set.

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


### GetServiceUserIds

`func (o *ReviewerStage) GetServiceUserIds() []string`

GetServiceUserIds returns the ServiceUserIds field if non-nil, zero value otherwise.

### GetServiceUserIdsOk

`func (o *ReviewerStage) GetServiceUserIdsOk() (*[]string, bool)`

GetServiceUserIdsOk returns a tuple with the ServiceUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUserIds

`func (o *ReviewerStage) SetServiceUserIds(v []string)`

SetServiceUserIds sets ServiceUserIds field to given value.

### HasServiceUserIds

`func (o *ReviewerStage) HasServiceUserIds() bool`

HasServiceUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


