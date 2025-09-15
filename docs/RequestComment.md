# RequestComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | The date and time the comment was created. | 
**RequestId** | **string** | The unique identifier of the request the comment is associated with. | 
**UserId** | **string** | The unique identifier of the user who made the comment. | 
**UserFullName** | Pointer to **string** | The user&#39;s full name. | [optional] 
**UserEmail** | Pointer to **string** | The user&#39;s email address. | [optional] 
**Comment** | **string** | The content of the comment. | 

## Methods

### NewRequestComment

`func NewRequestComment(createdAt time.Time, requestId string, userId string, comment string, ) *RequestComment`

NewRequestComment instantiates a new RequestComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCommentWithDefaults

`func NewRequestCommentWithDefaults() *RequestComment`

NewRequestCommentWithDefaults instantiates a new RequestComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RequestComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RequestComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RequestComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetRequestId

`func (o *RequestComment) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RequestComment) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RequestComment) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetUserId

`func (o *RequestComment) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RequestComment) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RequestComment) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserFullName

`func (o *RequestComment) GetUserFullName() string`

GetUserFullName returns the UserFullName field if non-nil, zero value otherwise.

### GetUserFullNameOk

`func (o *RequestComment) GetUserFullNameOk() (*string, bool)`

GetUserFullNameOk returns a tuple with the UserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFullName

`func (o *RequestComment) SetUserFullName(v string)`

SetUserFullName sets UserFullName field to given value.

### HasUserFullName

`func (o *RequestComment) HasUserFullName() bool`

HasUserFullName returns a boolean if a field has been set.

### GetUserEmail

`func (o *RequestComment) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *RequestComment) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *RequestComment) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *RequestComment) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetComment

`func (o *RequestComment) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RequestComment) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RequestComment) SetComment(v string)`

SetComment sets Comment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


