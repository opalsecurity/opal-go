# CustomConnectorAppConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The identifier of the Custom Connector. | 
**BaseUrl** | **string** | The base URL of the Custom Connector. | 
**TlsMode** | **bool** | Whether TLS verification is enabled for the Custom Connector. | 
**TlsCaCertContent** | Pointer to **string** | Optional PEM-encoded CA certificate content for TLS. | [optional] 
**SupportsGroups** | **bool** | Whether the Custom Connector supports groups. | 
**SupportsNestedResources** | **bool** | Whether the Custom Connector supports nested resources. | 
**SupportsNestedGroups** | **bool** | Whether the Custom Connector supports nested groups. | 
**SupportsEventIngestion** | **bool** | Whether the Custom Connector supports event ingestion. | 
**SupportsIdp** | **bool** | Whether the Custom Connector can be connected as an upstream user directory. | 

## Methods

### NewCustomConnectorAppConfig

`func NewCustomConnectorAppConfig(identifier string, baseUrl string, tlsMode bool, supportsGroups bool, supportsNestedResources bool, supportsNestedGroups bool, supportsEventIngestion bool, supportsIdp bool, ) *CustomConnectorAppConfig`

NewCustomConnectorAppConfig instantiates a new CustomConnectorAppConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomConnectorAppConfigWithDefaults

`func NewCustomConnectorAppConfigWithDefaults() *CustomConnectorAppConfig`

NewCustomConnectorAppConfigWithDefaults instantiates a new CustomConnectorAppConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *CustomConnectorAppConfig) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CustomConnectorAppConfig) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CustomConnectorAppConfig) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetBaseUrl

`func (o *CustomConnectorAppConfig) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *CustomConnectorAppConfig) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *CustomConnectorAppConfig) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetTlsMode

`func (o *CustomConnectorAppConfig) GetTlsMode() bool`

GetTlsMode returns the TlsMode field if non-nil, zero value otherwise.

### GetTlsModeOk

`func (o *CustomConnectorAppConfig) GetTlsModeOk() (*bool, bool)`

GetTlsModeOk returns a tuple with the TlsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMode

`func (o *CustomConnectorAppConfig) SetTlsMode(v bool)`

SetTlsMode sets TlsMode field to given value.


### GetTlsCaCertContent

`func (o *CustomConnectorAppConfig) GetTlsCaCertContent() string`

GetTlsCaCertContent returns the TlsCaCertContent field if non-nil, zero value otherwise.

### GetTlsCaCertContentOk

`func (o *CustomConnectorAppConfig) GetTlsCaCertContentOk() (*string, bool)`

GetTlsCaCertContentOk returns a tuple with the TlsCaCertContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCaCertContent

`func (o *CustomConnectorAppConfig) SetTlsCaCertContent(v string)`

SetTlsCaCertContent sets TlsCaCertContent field to given value.

### HasTlsCaCertContent

`func (o *CustomConnectorAppConfig) HasTlsCaCertContent() bool`

HasTlsCaCertContent returns a boolean if a field has been set.

### GetSupportsGroups

`func (o *CustomConnectorAppConfig) GetSupportsGroups() bool`

GetSupportsGroups returns the SupportsGroups field if non-nil, zero value otherwise.

### GetSupportsGroupsOk

`func (o *CustomConnectorAppConfig) GetSupportsGroupsOk() (*bool, bool)`

GetSupportsGroupsOk returns a tuple with the SupportsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsGroups

`func (o *CustomConnectorAppConfig) SetSupportsGroups(v bool)`

SetSupportsGroups sets SupportsGroups field to given value.


### GetSupportsNestedResources

`func (o *CustomConnectorAppConfig) GetSupportsNestedResources() bool`

GetSupportsNestedResources returns the SupportsNestedResources field if non-nil, zero value otherwise.

### GetSupportsNestedResourcesOk

`func (o *CustomConnectorAppConfig) GetSupportsNestedResourcesOk() (*bool, bool)`

GetSupportsNestedResourcesOk returns a tuple with the SupportsNestedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsNestedResources

`func (o *CustomConnectorAppConfig) SetSupportsNestedResources(v bool)`

SetSupportsNestedResources sets SupportsNestedResources field to given value.


### GetSupportsNestedGroups

`func (o *CustomConnectorAppConfig) GetSupportsNestedGroups() bool`

GetSupportsNestedGroups returns the SupportsNestedGroups field if non-nil, zero value otherwise.

### GetSupportsNestedGroupsOk

`func (o *CustomConnectorAppConfig) GetSupportsNestedGroupsOk() (*bool, bool)`

GetSupportsNestedGroupsOk returns a tuple with the SupportsNestedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsNestedGroups

`func (o *CustomConnectorAppConfig) SetSupportsNestedGroups(v bool)`

SetSupportsNestedGroups sets SupportsNestedGroups field to given value.


### GetSupportsEventIngestion

`func (o *CustomConnectorAppConfig) GetSupportsEventIngestion() bool`

GetSupportsEventIngestion returns the SupportsEventIngestion field if non-nil, zero value otherwise.

### GetSupportsEventIngestionOk

`func (o *CustomConnectorAppConfig) GetSupportsEventIngestionOk() (*bool, bool)`

GetSupportsEventIngestionOk returns a tuple with the SupportsEventIngestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsEventIngestion

`func (o *CustomConnectorAppConfig) SetSupportsEventIngestion(v bool)`

SetSupportsEventIngestion sets SupportsEventIngestion field to given value.


### GetSupportsIdp

`func (o *CustomConnectorAppConfig) GetSupportsIdp() bool`

GetSupportsIdp returns the SupportsIdp field if non-nil, zero value otherwise.

### GetSupportsIdpOk

`func (o *CustomConnectorAppConfig) GetSupportsIdpOk() (*bool, bool)`

GetSupportsIdpOk returns a tuple with the SupportsIdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsIdp

`func (o *CustomConnectorAppConfig) SetSupportsIdp(v bool)`

SetSupportsIdp sets SupportsIdp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


