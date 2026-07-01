# RequestReviewer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the reviewer | 
**FullName** | Pointer to **string** | The user&#39;s full name. | [optional] 
**Status** | **string** | The status of this reviewer&#39;s review | 

## Methods

### NewRequestReviewer

`func NewRequestReviewer(id string, status string, ) *RequestReviewer`

NewRequestReviewer instantiates a new RequestReviewer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestReviewerWithDefaults

`func NewRequestReviewerWithDefaults() *RequestReviewer`

NewRequestReviewerWithDefaults instantiates a new RequestReviewer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestReviewer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestReviewer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestReviewer) SetId(v string)`

SetId sets Id field to given value.


### GetFullName

`func (o *RequestReviewer) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *RequestReviewer) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *RequestReviewer) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *RequestReviewer) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetStatus

`func (o *RequestReviewer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestReviewer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestReviewer) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


