# ApproveRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActingUserId** | Pointer to **string** | Approve as this Opal user instead of the caller. Only for service users an admin has allowed to act for users, in organizations with the feature enabled. The user must be active in the identity provider and be a reviewer on the request&#39;s current stage and not the request&#39;s target; the decision counts as theirs. Requests that require MFA to approve, and ADMIN level, are rejected. | [optional] 
**Level** | [**RequestApprovalEnum**](RequestApprovalEnum.md) |  | 
**Comment** | Pointer to **string** | Optional comment for the approval | [optional] 

## Methods

### NewApproveRequestRequest

`func NewApproveRequestRequest(level RequestApprovalEnum, ) *ApproveRequestRequest`

NewApproveRequestRequest instantiates a new ApproveRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproveRequestRequestWithDefaults

`func NewApproveRequestRequestWithDefaults() *ApproveRequestRequest`

NewApproveRequestRequestWithDefaults instantiates a new ApproveRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActingUserId

`func (o *ApproveRequestRequest) GetActingUserId() string`

GetActingUserId returns the ActingUserId field if non-nil, zero value otherwise.

### GetActingUserIdOk

`func (o *ApproveRequestRequest) GetActingUserIdOk() (*string, bool)`

GetActingUserIdOk returns a tuple with the ActingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserId

`func (o *ApproveRequestRequest) SetActingUserId(v string)`

SetActingUserId sets ActingUserId field to given value.

### HasActingUserId

`func (o *ApproveRequestRequest) HasActingUserId() bool`

HasActingUserId returns a boolean if a field has been set.

### GetLevel

`func (o *ApproveRequestRequest) GetLevel() RequestApprovalEnum`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ApproveRequestRequest) GetLevelOk() (*RequestApprovalEnum, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ApproveRequestRequest) SetLevel(v RequestApprovalEnum)`

SetLevel sets Level field to given value.


### GetComment

`func (o *ApproveRequestRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApproveRequestRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApproveRequestRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApproveRequestRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


