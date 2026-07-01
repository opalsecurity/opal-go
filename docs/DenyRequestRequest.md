# DenyRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** | Comment for the denial | 
**Level** | Pointer to [**RequestApprovalEnum**](RequestApprovalEnum.md) |  | [optional] 

## Methods

### NewDenyRequestRequest

`func NewDenyRequestRequest(comment string, ) *DenyRequestRequest`

NewDenyRequestRequest instantiates a new DenyRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDenyRequestRequestWithDefaults

`func NewDenyRequestRequestWithDefaults() *DenyRequestRequest`

NewDenyRequestRequestWithDefaults instantiates a new DenyRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *DenyRequestRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DenyRequestRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DenyRequestRequest) SetComment(v string)`

SetComment sets Comment field to given value.


### GetLevel

`func (o *DenyRequestRequest) GetLevel() RequestApprovalEnum`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *DenyRequestRequest) GetLevelOk() (*RequestApprovalEnum, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *DenyRequestRequest) SetLevel(v RequestApprovalEnum)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *DenyRequestRequest) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


