# UpdateAppInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The updated name of the app. | [optional] 
**Description** | Pointer to **string** | The updated description of the app. | [optional] 
**AdminOwnerId** | Pointer to **string** | The updated ID of the owner of the app. | [optional] 
**Visibility** | Pointer to [**VisibilityTypeEnum**](VisibilityTypeEnum.md) |  | [optional] 
**VisibilityGroupIds** | Pointer to **[]string** | The IDs of groups that can see this app when visibility is &#x60;LIMITED&#x60;. | [optional] 
**ImportVisibility** | Pointer to [**VisibilityTypeEnum**](VisibilityTypeEnum.md) |  | [optional] 
**CustomConnector** | Pointer to [**UpdateCustomConnectorInfo**](UpdateCustomConnectorInfo.md) | Updates Custom Connector configuration. Only valid for &#x60;CUSTOM_CONNECTOR&#x60; apps. | [optional] 

## Methods

### NewUpdateAppInfo

`func NewUpdateAppInfo() *UpdateAppInfo`

NewUpdateAppInfo instantiates a new UpdateAppInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppInfoWithDefaults

`func NewUpdateAppInfoWithDefaults() *UpdateAppInfo`

NewUpdateAppInfoWithDefaults instantiates a new UpdateAppInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAppInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAppInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAppInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAppInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAppInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAppInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAppInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAppInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdminOwnerId

`func (o *UpdateAppInfo) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *UpdateAppInfo) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *UpdateAppInfo) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.

### HasAdminOwnerId

`func (o *UpdateAppInfo) HasAdminOwnerId() bool`

HasAdminOwnerId returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateAppInfo) GetVisibility() VisibilityTypeEnum`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateAppInfo) GetVisibilityOk() (*VisibilityTypeEnum, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateAppInfo) SetVisibility(v VisibilityTypeEnum)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateAppInfo) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetVisibilityGroupIds

`func (o *UpdateAppInfo) GetVisibilityGroupIds() []string`

GetVisibilityGroupIds returns the VisibilityGroupIds field if non-nil, zero value otherwise.

### GetVisibilityGroupIdsOk

`func (o *UpdateAppInfo) GetVisibilityGroupIdsOk() (*[]string, bool)`

GetVisibilityGroupIdsOk returns a tuple with the VisibilityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityGroupIds

`func (o *UpdateAppInfo) SetVisibilityGroupIds(v []string)`

SetVisibilityGroupIds sets VisibilityGroupIds field to given value.

### HasVisibilityGroupIds

`func (o *UpdateAppInfo) HasVisibilityGroupIds() bool`

HasVisibilityGroupIds returns a boolean if a field has been set.

### GetImportVisibility

`func (o *UpdateAppInfo) GetImportVisibility() VisibilityTypeEnum`

GetImportVisibility returns the ImportVisibility field if non-nil, zero value otherwise.

### GetImportVisibilityOk

`func (o *UpdateAppInfo) GetImportVisibilityOk() (*VisibilityTypeEnum, bool)`

GetImportVisibilityOk returns a tuple with the ImportVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportVisibility

`func (o *UpdateAppInfo) SetImportVisibility(v VisibilityTypeEnum)`

SetImportVisibility sets ImportVisibility field to given value.

### HasImportVisibility

`func (o *UpdateAppInfo) HasImportVisibility() bool`

HasImportVisibility returns a boolean if a field has been set.

### GetCustomConnector

`func (o *UpdateAppInfo) GetCustomConnector() UpdateCustomConnectorInfo`

GetCustomConnector returns the CustomConnector field if non-nil, zero value otherwise.

### GetCustomConnectorOk

`func (o *UpdateAppInfo) GetCustomConnectorOk() (*UpdateCustomConnectorInfo, bool)`

GetCustomConnectorOk returns a tuple with the CustomConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConnector

`func (o *UpdateAppInfo) SetCustomConnector(v UpdateCustomConnectorInfo)`

SetCustomConnector sets CustomConnector field to given value.

### HasCustomConnector

`func (o *UpdateAppInfo) HasCustomConnector() bool`

HasCustomConnector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


