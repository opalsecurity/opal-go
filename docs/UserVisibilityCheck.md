# UserVisibilityCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visible** | **bool** |  | 
**Reason** | [**VisibilityReasonEnum**](VisibilityReasonEnum.md) |  | 

## Methods

### NewUserVisibilityCheck

`func NewUserVisibilityCheck(visible bool, reason VisibilityReasonEnum, ) *UserVisibilityCheck`

NewUserVisibilityCheck instantiates a new UserVisibilityCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserVisibilityCheckWithDefaults

`func NewUserVisibilityCheckWithDefaults() *UserVisibilityCheck`

NewUserVisibilityCheckWithDefaults instantiates a new UserVisibilityCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisible

`func (o *UserVisibilityCheck) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *UserVisibilityCheck) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *UserVisibilityCheck) SetVisible(v bool)`

SetVisible sets Visible field to given value.


### GetReason

`func (o *UserVisibilityCheck) GetReason() VisibilityReasonEnum`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UserVisibilityCheck) GetReasonOk() (*VisibilityReasonEnum, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UserVisibilityCheck) SetReason(v VisibilityReasonEnum)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


