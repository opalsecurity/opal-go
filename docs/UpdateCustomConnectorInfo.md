# UpdateCustomConnectorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The identifier of the Custom Connector. | [optional] 
**BaseUrl** | Pointer to **string** | The base URL of the Custom Connector. | [optional] 
**SigningSecret** | Pointer to **string** | New signing secret. Write-only; never returned by the API. Omit to leave the existing secret unchanged. | [optional] 
**TlsMode** | Pointer to **bool** | Whether TLS verification is enabled for the Custom Connector. | [optional] 
**TlsCaCertContent** | Pointer to **string** | Optional PEM-encoded CA certificate content for TLS. | [optional] 
**SupportsGroups** | Pointer to **bool** | Whether the Custom Connector supports groups. | [optional] 
**SupportsNestedResources** | Pointer to **bool** | Whether the Custom Connector supports nested resources. | [optional] 
**SupportsNestedGroups** | Pointer to **bool** | Whether the Custom Connector supports nested groups. | [optional] 
**SupportsEventIngestion** | Pointer to **bool** | Whether the Custom Connector supports event ingestion. | [optional] 

## Methods

### NewUpdateCustomConnectorInfo

`func NewUpdateCustomConnectorInfo() *UpdateCustomConnectorInfo`

NewUpdateCustomConnectorInfo instantiates a new UpdateCustomConnectorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomConnectorInfoWithDefaults

`func NewUpdateCustomConnectorInfoWithDefaults() *UpdateCustomConnectorInfo`

NewUpdateCustomConnectorInfoWithDefaults instantiates a new UpdateCustomConnectorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *UpdateCustomConnectorInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UpdateCustomConnectorInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UpdateCustomConnectorInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *UpdateCustomConnectorInfo) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetBaseUrl

`func (o *UpdateCustomConnectorInfo) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *UpdateCustomConnectorInfo) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *UpdateCustomConnectorInfo) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *UpdateCustomConnectorInfo) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetSigningSecret

`func (o *UpdateCustomConnectorInfo) GetSigningSecret() string`

GetSigningSecret returns the SigningSecret field if non-nil, zero value otherwise.

### GetSigningSecretOk

`func (o *UpdateCustomConnectorInfo) GetSigningSecretOk() (*string, bool)`

GetSigningSecretOk returns a tuple with the SigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecret

`func (o *UpdateCustomConnectorInfo) SetSigningSecret(v string)`

SetSigningSecret sets SigningSecret field to given value.

### HasSigningSecret

`func (o *UpdateCustomConnectorInfo) HasSigningSecret() bool`

HasSigningSecret returns a boolean if a field has been set.

### GetTlsMode

`func (o *UpdateCustomConnectorInfo) GetTlsMode() bool`

GetTlsMode returns the TlsMode field if non-nil, zero value otherwise.

### GetTlsModeOk

`func (o *UpdateCustomConnectorInfo) GetTlsModeOk() (*bool, bool)`

GetTlsModeOk returns a tuple with the TlsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMode

`func (o *UpdateCustomConnectorInfo) SetTlsMode(v bool)`

SetTlsMode sets TlsMode field to given value.

### HasTlsMode

`func (o *UpdateCustomConnectorInfo) HasTlsMode() bool`

HasTlsMode returns a boolean if a field has been set.

### GetTlsCaCertContent

`func (o *UpdateCustomConnectorInfo) GetTlsCaCertContent() string`

GetTlsCaCertContent returns the TlsCaCertContent field if non-nil, zero value otherwise.

### GetTlsCaCertContentOk

`func (o *UpdateCustomConnectorInfo) GetTlsCaCertContentOk() (*string, bool)`

GetTlsCaCertContentOk returns a tuple with the TlsCaCertContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCaCertContent

`func (o *UpdateCustomConnectorInfo) SetTlsCaCertContent(v string)`

SetTlsCaCertContent sets TlsCaCertContent field to given value.

### HasTlsCaCertContent

`func (o *UpdateCustomConnectorInfo) HasTlsCaCertContent() bool`

HasTlsCaCertContent returns a boolean if a field has been set.

### GetSupportsGroups

`func (o *UpdateCustomConnectorInfo) GetSupportsGroups() bool`

GetSupportsGroups returns the SupportsGroups field if non-nil, zero value otherwise.

### GetSupportsGroupsOk

`func (o *UpdateCustomConnectorInfo) GetSupportsGroupsOk() (*bool, bool)`

GetSupportsGroupsOk returns a tuple with the SupportsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsGroups

`func (o *UpdateCustomConnectorInfo) SetSupportsGroups(v bool)`

SetSupportsGroups sets SupportsGroups field to given value.

### HasSupportsGroups

`func (o *UpdateCustomConnectorInfo) HasSupportsGroups() bool`

HasSupportsGroups returns a boolean if a field has been set.

### GetSupportsNestedResources

`func (o *UpdateCustomConnectorInfo) GetSupportsNestedResources() bool`

GetSupportsNestedResources returns the SupportsNestedResources field if non-nil, zero value otherwise.

### GetSupportsNestedResourcesOk

`func (o *UpdateCustomConnectorInfo) GetSupportsNestedResourcesOk() (*bool, bool)`

GetSupportsNestedResourcesOk returns a tuple with the SupportsNestedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsNestedResources

`func (o *UpdateCustomConnectorInfo) SetSupportsNestedResources(v bool)`

SetSupportsNestedResources sets SupportsNestedResources field to given value.

### HasSupportsNestedResources

`func (o *UpdateCustomConnectorInfo) HasSupportsNestedResources() bool`

HasSupportsNestedResources returns a boolean if a field has been set.

### GetSupportsNestedGroups

`func (o *UpdateCustomConnectorInfo) GetSupportsNestedGroups() bool`

GetSupportsNestedGroups returns the SupportsNestedGroups field if non-nil, zero value otherwise.

### GetSupportsNestedGroupsOk

`func (o *UpdateCustomConnectorInfo) GetSupportsNestedGroupsOk() (*bool, bool)`

GetSupportsNestedGroupsOk returns a tuple with the SupportsNestedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsNestedGroups

`func (o *UpdateCustomConnectorInfo) SetSupportsNestedGroups(v bool)`

SetSupportsNestedGroups sets SupportsNestedGroups field to given value.

### HasSupportsNestedGroups

`func (o *UpdateCustomConnectorInfo) HasSupportsNestedGroups() bool`

HasSupportsNestedGroups returns a boolean if a field has been set.

### GetSupportsEventIngestion

`func (o *UpdateCustomConnectorInfo) GetSupportsEventIngestion() bool`

GetSupportsEventIngestion returns the SupportsEventIngestion field if non-nil, zero value otherwise.

### GetSupportsEventIngestionOk

`func (o *UpdateCustomConnectorInfo) GetSupportsEventIngestionOk() (*bool, bool)`

GetSupportsEventIngestionOk returns a tuple with the SupportsEventIngestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsEventIngestion

`func (o *UpdateCustomConnectorInfo) SetSupportsEventIngestion(v bool)`

SetSupportsEventIngestion sets SupportsEventIngestion field to given value.

### HasSupportsEventIngestion

`func (o *UpdateCustomConnectorInfo) HasSupportsEventIngestion() bool`

HasSupportsEventIngestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


