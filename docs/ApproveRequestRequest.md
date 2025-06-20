# ApproveRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | The decision level for the approval | 
**Comment** | Pointer to **string** | Optional comment for the approval | [optional] 

## Methods

### NewApproveRequestRequest

`func NewApproveRequestRequest(level string, ) *ApproveRequestRequest`

NewApproveRequestRequest instantiates a new ApproveRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproveRequestRequestWithDefaults

`func NewApproveRequestRequestWithDefaults() *ApproveRequestRequest`

NewApproveRequestRequestWithDefaults instantiates a new ApproveRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *ApproveRequestRequest) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ApproveRequestRequest) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ApproveRequestRequest) SetLevel(v string)`

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


