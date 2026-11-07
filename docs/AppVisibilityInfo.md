# AppVisibilityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visibility** | [**VisibilityTypeEnum**](VisibilityTypeEnum.md) |  | 
**VisibilityGroupIds** | Pointer to **[]string** | The IDs of groups that can see this app when visibility is &#x60;LIMITED&#x60;. Must be empty when visibility is &#x60;GLOBAL&#x60;. | [optional] 
**ImportVisibility** | Pointer to [**VisibilityTypeEnum**](VisibilityTypeEnum.md) | The visibility of items imported from this app. Defaults to &#x60;GLOBAL&#x60; when omitted on a write. | [optional] 

## Methods

### NewAppVisibilityInfo

`func NewAppVisibilityInfo(visibility VisibilityTypeEnum, ) *AppVisibilityInfo`

NewAppVisibilityInfo instantiates a new AppVisibilityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppVisibilityInfoWithDefaults

`func NewAppVisibilityInfoWithDefaults() *AppVisibilityInfo`

NewAppVisibilityInfoWithDefaults instantiates a new AppVisibilityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisibility

`func (o *AppVisibilityInfo) GetVisibility() VisibilityTypeEnum`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AppVisibilityInfo) GetVisibilityOk() (*VisibilityTypeEnum, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AppVisibilityInfo) SetVisibility(v VisibilityTypeEnum)`

SetVisibility sets Visibility field to given value.


### GetVisibilityGroupIds

`func (o *AppVisibilityInfo) GetVisibilityGroupIds() []string`

GetVisibilityGroupIds returns the VisibilityGroupIds field if non-nil, zero value otherwise.

### GetVisibilityGroupIdsOk

`func (o *AppVisibilityInfo) GetVisibilityGroupIdsOk() (*[]string, bool)`

GetVisibilityGroupIdsOk returns a tuple with the VisibilityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityGroupIds

`func (o *AppVisibilityInfo) SetVisibilityGroupIds(v []string)`

SetVisibilityGroupIds sets VisibilityGroupIds field to given value.

### HasVisibilityGroupIds

`func (o *AppVisibilityInfo) HasVisibilityGroupIds() bool`

HasVisibilityGroupIds returns a boolean if a field has been set.

### GetImportVisibility

`func (o *AppVisibilityInfo) GetImportVisibility() VisibilityTypeEnum`

GetImportVisibility returns the ImportVisibility field if non-nil, zero value otherwise.

### GetImportVisibilityOk

`func (o *AppVisibilityInfo) GetImportVisibilityOk() (*VisibilityTypeEnum, bool)`

GetImportVisibilityOk returns a tuple with the ImportVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportVisibility

`func (o *AppVisibilityInfo) SetImportVisibility(v VisibilityTypeEnum)`

SetImportVisibility sets ImportVisibility field to given value.

### HasImportVisibility

`func (o *AppVisibilityInfo) HasImportVisibility() bool`

HasImportVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


