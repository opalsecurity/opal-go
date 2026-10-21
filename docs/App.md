# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | The ID of the app. | 
**Name** | **string** | The name of the app. | 
**Description** | **string** | A description of the app. | 
**AdminOwnerId** | **string** | The ID of the owner of the app. | 
**AppType** | [**AppTypeEnum**](AppTypeEnum.md) |  | 
**Visibility** | Pointer to [**VisibilityTypeEnum**](VisibilityTypeEnum.md) |  | [optional] 
**VisibilityGroupIds** | Pointer to **[]string** | The IDs of groups that can see this app when visibility is &#x60;LIMITED&#x60;. | [optional] 
**ImportVisibility** | Pointer to [**VisibilityTypeEnum**](VisibilityTypeEnum.md) | The visibility level of items imported from this app. | [optional] 
**Validations** | Pointer to [**[]AppValidation**](AppValidation.md) | Validation checks of an apps&#39; configuration and permissions. | [optional] 
**CustomConnector** | Pointer to [**CustomConnectorAppConfig**](CustomConnectorAppConfig.md) |  | [optional] 

## Methods

### NewApp

`func NewApp(appId string, name string, description string, adminOwnerId string, appType AppTypeEnum, ) *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *App) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *App) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *App) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetName

`func (o *App) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *App) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *App) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *App) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *App) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *App) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAdminOwnerId

`func (o *App) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *App) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *App) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.


### GetAppType

`func (o *App) GetAppType() AppTypeEnum`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *App) GetAppTypeOk() (*AppTypeEnum, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *App) SetAppType(v AppTypeEnum)`

SetAppType sets AppType field to given value.


### GetVisibility

`func (o *App) GetVisibility() VisibilityTypeEnum`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *App) GetVisibilityOk() (*VisibilityTypeEnum, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *App) SetVisibility(v VisibilityTypeEnum)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *App) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetVisibilityGroupIds

`func (o *App) GetVisibilityGroupIds() []string`

GetVisibilityGroupIds returns the VisibilityGroupIds field if non-nil, zero value otherwise.

### GetVisibilityGroupIdsOk

`func (o *App) GetVisibilityGroupIdsOk() (*[]string, bool)`

GetVisibilityGroupIdsOk returns a tuple with the VisibilityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityGroupIds

`func (o *App) SetVisibilityGroupIds(v []string)`

SetVisibilityGroupIds sets VisibilityGroupIds field to given value.

### HasVisibilityGroupIds

`func (o *App) HasVisibilityGroupIds() bool`

HasVisibilityGroupIds returns a boolean if a field has been set.

### GetImportVisibility

`func (o *App) GetImportVisibility() VisibilityTypeEnum`

GetImportVisibility returns the ImportVisibility field if non-nil, zero value otherwise.

### GetImportVisibilityOk

`func (o *App) GetImportVisibilityOk() (*VisibilityTypeEnum, bool)`

GetImportVisibilityOk returns a tuple with the ImportVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportVisibility

`func (o *App) SetImportVisibility(v VisibilityTypeEnum)`

SetImportVisibility sets ImportVisibility field to given value.

### HasImportVisibility

`func (o *App) HasImportVisibility() bool`

HasImportVisibility returns a boolean if a field has been set.

### GetValidations

`func (o *App) GetValidations() []AppValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *App) GetValidationsOk() (*[]AppValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *App) SetValidations(v []AppValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *App) HasValidations() bool`

HasValidations returns a boolean if a field has been set.

### GetCustomConnector

`func (o *App) GetCustomConnector() CustomConnectorAppConfig`

GetCustomConnector returns the CustomConnector field if non-nil, zero value otherwise.

### GetCustomConnectorOk

`func (o *App) GetCustomConnectorOk() (*CustomConnectorAppConfig, bool)`

GetCustomConnectorOk returns a tuple with the CustomConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConnector

`func (o *App) SetCustomConnector(v CustomConnectorAppConfig)`

SetCustomConnector sets CustomConnector field to given value.

### HasCustomConnector

`func (o *App) HasCustomConnector() bool`

HasCustomConnector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


