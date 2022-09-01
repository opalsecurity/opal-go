# VisibilityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visibility** | [**VisibilityTypeEnum**](VisibilityTypeEnum.md) |  | 
**VisibilityGroupIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVisibilityInfo

`func NewVisibilityInfo(visibility VisibilityTypeEnum, ) *VisibilityInfo`

NewVisibilityInfo instantiates a new VisibilityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisibilityInfoWithDefaults

`func NewVisibilityInfoWithDefaults() *VisibilityInfo`

NewVisibilityInfoWithDefaults instantiates a new VisibilityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisibility

`func (o *VisibilityInfo) GetVisibility() VisibilityTypeEnum`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *VisibilityInfo) GetVisibilityOk() (*VisibilityTypeEnum, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *VisibilityInfo) SetVisibility(v VisibilityTypeEnum)`

SetVisibility sets Visibility field to given value.


### GetVisibilityGroupIds

`func (o *VisibilityInfo) GetVisibilityGroupIds() []string`

GetVisibilityGroupIds returns the VisibilityGroupIds field if non-nil, zero value otherwise.

### GetVisibilityGroupIdsOk

`func (o *VisibilityInfo) GetVisibilityGroupIdsOk() (*[]string, bool)`

GetVisibilityGroupIdsOk returns a tuple with the VisibilityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityGroupIds

`func (o *VisibilityInfo) SetVisibilityGroupIds(v []string)`

SetVisibilityGroupIds sets VisibilityGroupIds field to given value.

### HasVisibilityGroupIds

`func (o *VisibilityInfo) HasVisibilityGroupIds() bool`

HasVisibilityGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


