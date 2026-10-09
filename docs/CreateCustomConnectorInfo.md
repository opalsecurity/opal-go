# CreateCustomConnectorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The identifier of the Custom Connector. | 
**BaseUrl** | **string** | The base URL of the Custom Connector. | 
**SigningSecret** | **string** | The signing secret used to authenticate requests to the Custom Connector. Write-only; never returned by the API. | 
**TlsMode** | Pointer to **bool** | Whether TLS verification is enabled. Defaults to &#x60;true&#x60; when omitted. | [optional] 
**TlsCaCertContent** | Pointer to **string** | Optional PEM-encoded CA certificate content for TLS. | [optional] 
**SupportsGroups** | Pointer to **bool** | Whether the Custom Connector supports groups. Defaults to &#x60;false&#x60;. | [optional] 
**SupportsNestedResources** | Pointer to **bool** | Whether the Custom Connector supports nested resources. Defaults to &#x60;false&#x60;. | [optional] 
**SupportsNestedGroups** | Pointer to **bool** | Whether the Custom Connector supports nested groups. Defaults to &#x60;false&#x60;. | [optional] 
**SupportsEventIngestion** | Pointer to **bool** | Whether the Custom Connector supports event ingestion. Defaults to &#x60;false&#x60;. | [optional] 

## Methods

### NewCreateCustomConnectorInfo

`func NewCreateCustomConnectorInfo(identifier string, baseUrl string, signingSecret string, ) *CreateCustomConnectorInfo`

NewCreateCustomConnectorInfo instantiates a new CreateCustomConnectorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomConnectorInfoWithDefaults

`func NewCreateCustomConnectorInfoWithDefaults() *CreateCustomConnectorInfo`

NewCreateCustomConnectorInfoWithDefaults instantiates a new CreateCustomConnectorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *CreateCustomConnectorInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateCustomConnectorInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateCustomConnectorInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetBaseUrl

`func (o *CreateCustomConnectorInfo) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *CreateCustomConnectorInfo) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *CreateCustomConnectorInfo) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetSigningSecret

`func (o *CreateCustomConnectorInfo) GetSigningSecret() string`

GetSigningSecret returns the SigningSecret field if non-nil, zero value otherwise.

### GetSigningSecretOk

`func (o *CreateCustomConnectorInfo) GetSigningSecretOk() (*string, bool)`

GetSigningSecretOk returns a tuple with the SigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecret

`func (o *CreateCustomConnectorInfo) SetSigningSecret(v string)`

SetSigningSecret sets SigningSecret field to given value.


### GetTlsMode

`func (o *CreateCustomConnectorInfo) GetTlsMode() bool`

GetTlsMode returns the TlsMode field if non-nil, zero value otherwise.

### GetTlsModeOk

`func (o *CreateCustomConnectorInfo) GetTlsModeOk() (*bool, bool)`

GetTlsModeOk returns a tuple with the TlsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMode

`func (o *CreateCustomConnectorInfo) SetTlsMode(v bool)`

SetTlsMode sets TlsMode field to given value.

### HasTlsMode

`func (o *CreateCustomConnectorInfo) HasTlsMode() bool`

HasTlsMode returns a boolean if a field has been set.

### GetTlsCaCertContent

`func (o *CreateCustomConnectorInfo) GetTlsCaCertContent() string`

GetTlsCaCertContent returns the TlsCaCertContent field if non-nil, zero value otherwise.

### GetTlsCaCertContentOk

`func (o *CreateCustomConnectorInfo) GetTlsCaCertContentOk() (*string, bool)`

GetTlsCaCertContentOk returns a tuple with the TlsCaCertContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCaCertContent

`func (o *CreateCustomConnectorInfo) SetTlsCaCertContent(v string)`

SetTlsCaCertContent sets TlsCaCertContent field to given value.

### HasTlsCaCertContent

`func (o *CreateCustomConnectorInfo) HasTlsCaCertContent() bool`

HasTlsCaCertContent returns a boolean if a field has been set.

### GetSupportsGroups

`func (o *CreateCustomConnectorInfo) GetSupportsGroups() bool`

GetSupportsGroups returns the SupportsGroups field if non-nil, zero value otherwise.

### GetSupportsGroupsOk

`func (o *CreateCustomConnectorInfo) GetSupportsGroupsOk() (*bool, bool)`

GetSupportsGroupsOk returns a tuple with the SupportsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsGroups

`func (o *CreateCustomConnectorInfo) SetSupportsGroups(v bool)`

SetSupportsGroups sets SupportsGroups field to given value.

### HasSupportsGroups

`func (o *CreateCustomConnectorInfo) HasSupportsGroups() bool`

HasSupportsGroups returns a boolean if a field has been set.

### GetSupportsNestedResources

`func (o *CreateCustomConnectorInfo) GetSupportsNestedResources() bool`

GetSupportsNestedResources returns the SupportsNestedResources field if non-nil, zero value otherwise.

### GetSupportsNestedResourcesOk

`func (o *CreateCustomConnectorInfo) GetSupportsNestedResourcesOk() (*bool, bool)`

GetSupportsNestedResourcesOk returns a tuple with the SupportsNestedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsNestedResources

`func (o *CreateCustomConnectorInfo) SetSupportsNestedResources(v bool)`

SetSupportsNestedResources sets SupportsNestedResources field to given value.

### HasSupportsNestedResources

`func (o *CreateCustomConnectorInfo) HasSupportsNestedResources() bool`

HasSupportsNestedResources returns a boolean if a field has been set.

### GetSupportsNestedGroups

`func (o *CreateCustomConnectorInfo) GetSupportsNestedGroups() bool`

GetSupportsNestedGroups returns the SupportsNestedGroups field if non-nil, zero value otherwise.

### GetSupportsNestedGroupsOk

`func (o *CreateCustomConnectorInfo) GetSupportsNestedGroupsOk() (*bool, bool)`

GetSupportsNestedGroupsOk returns a tuple with the SupportsNestedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsNestedGroups

`func (o *CreateCustomConnectorInfo) SetSupportsNestedGroups(v bool)`

SetSupportsNestedGroups sets SupportsNestedGroups field to given value.

### HasSupportsNestedGroups

`func (o *CreateCustomConnectorInfo) HasSupportsNestedGroups() bool`

HasSupportsNestedGroups returns a boolean if a field has been set.

### GetSupportsEventIngestion

`func (o *CreateCustomConnectorInfo) GetSupportsEventIngestion() bool`

GetSupportsEventIngestion returns the SupportsEventIngestion field if non-nil, zero value otherwise.

### GetSupportsEventIngestionOk

`func (o *CreateCustomConnectorInfo) GetSupportsEventIngestionOk() (*bool, bool)`

GetSupportsEventIngestionOk returns a tuple with the SupportsEventIngestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsEventIngestion

`func (o *CreateCustomConnectorInfo) SetSupportsEventIngestion(v bool)`

SetSupportsEventIngestion sets SupportsEventIngestion field to given value.

### HasSupportsEventIngestion

`func (o *CreateCustomConnectorInfo) HasSupportsEventIngestion() bool`

HasSupportsEventIngestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


