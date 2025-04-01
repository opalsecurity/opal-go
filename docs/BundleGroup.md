# BundleGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | Pointer to **string** | The ID of the bundle containing the group. | [optional] [readonly] 
**GroupId** | Pointer to **string** | The ID of the group within a bundle. | [optional] [readonly] 

## Methods

### NewBundleGroup

`func NewBundleGroup() *BundleGroup`

NewBundleGroup instantiates a new BundleGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleGroupWithDefaults

`func NewBundleGroupWithDefaults() *BundleGroup`

NewBundleGroupWithDefaults instantiates a new BundleGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *BundleGroup) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *BundleGroup) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *BundleGroup) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *BundleGroup) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetGroupId

`func (o *BundleGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BundleGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BundleGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BundleGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


