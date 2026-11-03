# CreateAppInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the app. | 
**Description** | **string** | A description of the app. | 
**AdminOwnerId** | **string** | The ID of the owner of the app. | 
**AppType** | [**AppTypeEnum**](AppTypeEnum.md) | The type of the app. Must be &#x60;CUSTOM&#x60; (Push-only App) or &#x60;CUSTOM_CONNECTOR&#x60;. | 
**Visibility** | Pointer to [**VisibilityTypeEnum**](VisibilityTypeEnum.md) | The visibility of the app. Defaults to &#x60;GLOBAL&#x60; when omitted. | [optional] 
**VisibilityGroupIds** | Pointer to **[]string** | The IDs of groups that can see this app when visibility is &#x60;LIMITED&#x60;. Must be empty when visibility is &#x60;GLOBAL&#x60;; the server returns &#x60;400&#x60; otherwise. | [optional] 
**ImportVisibility** | Pointer to [**VisibilityTypeEnum**](VisibilityTypeEnum.md) | The visibility of imported items. Defaults to &#x60;GLOBAL&#x60; when omitted. | [optional] 
**CustomConnector** | Pointer to [**CreateCustomConnectorInfo**](CreateCustomConnectorInfo.md) | Required when &#x60;app_type&#x60; is &#x60;CUSTOM_CONNECTOR&#x60;. Must not be set when &#x60;app_type&#x60; is &#x60;CUSTOM&#x60;. | [optional] 

## Methods

### NewCreateAppInfo

`func NewCreateAppInfo(name string, description string, adminOwnerId string, appType AppTypeEnum, ) *CreateAppInfo`

NewCreateAppInfo instantiates a new CreateAppInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppInfoWithDefaults

`func NewCreateAppInfoWithDefaults() *CreateAppInfo`

NewCreateAppInfoWithDefaults instantiates a new CreateAppInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAppInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAppInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAppInfo) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAppInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAppInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAppInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAdminOwnerId

`func (o *CreateAppInfo) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *CreateAppInfo) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *CreateAppInfo) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.


### GetAppType

`func (o *CreateAppInfo) GetAppType() AppTypeEnum`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *CreateAppInfo) GetAppTypeOk() (*AppTypeEnum, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *CreateAppInfo) SetAppType(v AppTypeEnum)`

SetAppType sets AppType field to given value.


### GetVisibility

`func (o *CreateAppInfo) GetVisibility() VisibilityTypeEnum`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateAppInfo) GetVisibilityOk() (*VisibilityTypeEnum, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateAppInfo) SetVisibility(v VisibilityTypeEnum)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateAppInfo) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetVisibilityGroupIds

`func (o *CreateAppInfo) GetVisibilityGroupIds() []string`

GetVisibilityGroupIds returns the VisibilityGroupIds field if non-nil, zero value otherwise.

### GetVisibilityGroupIdsOk

`func (o *CreateAppInfo) GetVisibilityGroupIdsOk() (*[]string, bool)`

GetVisibilityGroupIdsOk returns a tuple with the VisibilityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityGroupIds

`func (o *CreateAppInfo) SetVisibilityGroupIds(v []string)`

SetVisibilityGroupIds sets VisibilityGroupIds field to given value.

### HasVisibilityGroupIds

`func (o *CreateAppInfo) HasVisibilityGroupIds() bool`

HasVisibilityGroupIds returns a boolean if a field has been set.

### GetImportVisibility

`func (o *CreateAppInfo) GetImportVisibility() VisibilityTypeEnum`

GetImportVisibility returns the ImportVisibility field if non-nil, zero value otherwise.

### GetImportVisibilityOk

`func (o *CreateAppInfo) GetImportVisibilityOk() (*VisibilityTypeEnum, bool)`

GetImportVisibilityOk returns a tuple with the ImportVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportVisibility

`func (o *CreateAppInfo) SetImportVisibility(v VisibilityTypeEnum)`

SetImportVisibility sets ImportVisibility field to given value.

### HasImportVisibility

`func (o *CreateAppInfo) HasImportVisibility() bool`

HasImportVisibility returns a boolean if a field has been set.

### GetCustomConnector

`func (o *CreateAppInfo) GetCustomConnector() CreateCustomConnectorInfo`

GetCustomConnector returns the CustomConnector field if non-nil, zero value otherwise.

### GetCustomConnectorOk

`func (o *CreateAppInfo) GetCustomConnectorOk() (*CreateCustomConnectorInfo, bool)`

GetCustomConnectorOk returns a tuple with the CustomConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConnector

`func (o *CreateAppInfo) SetCustomConnector(v CreateCustomConnectorInfo)`

SetCustomConnector sets CustomConnector field to given value.

### HasCustomConnector

`func (o *CreateAppInfo) HasCustomConnector() bool`

HasCustomConnector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


