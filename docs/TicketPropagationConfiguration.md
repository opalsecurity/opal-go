# TicketPropagationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledOnGrant** | **bool** |  | 
**EnabledOnRevocation** | **bool** |  | 
**TicketProvider** | Pointer to [**TicketingProviderEnum**](TicketingProviderEnum.md) |  | [optional] 
**TicketProjectId** | Pointer to **string** |  | [optional] 

## Methods

### NewTicketPropagationConfiguration

`func NewTicketPropagationConfiguration(enabledOnGrant bool, enabledOnRevocation bool, ) *TicketPropagationConfiguration`

NewTicketPropagationConfiguration instantiates a new TicketPropagationConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketPropagationConfigurationWithDefaults

`func NewTicketPropagationConfigurationWithDefaults() *TicketPropagationConfiguration`

NewTicketPropagationConfigurationWithDefaults instantiates a new TicketPropagationConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledOnGrant

`func (o *TicketPropagationConfiguration) GetEnabledOnGrant() bool`

GetEnabledOnGrant returns the EnabledOnGrant field if non-nil, zero value otherwise.

### GetEnabledOnGrantOk

`func (o *TicketPropagationConfiguration) GetEnabledOnGrantOk() (*bool, bool)`

GetEnabledOnGrantOk returns a tuple with the EnabledOnGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnGrant

`func (o *TicketPropagationConfiguration) SetEnabledOnGrant(v bool)`

SetEnabledOnGrant sets EnabledOnGrant field to given value.


### GetEnabledOnRevocation

`func (o *TicketPropagationConfiguration) GetEnabledOnRevocation() bool`

GetEnabledOnRevocation returns the EnabledOnRevocation field if non-nil, zero value otherwise.

### GetEnabledOnRevocationOk

`func (o *TicketPropagationConfiguration) GetEnabledOnRevocationOk() (*bool, bool)`

GetEnabledOnRevocationOk returns a tuple with the EnabledOnRevocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnRevocation

`func (o *TicketPropagationConfiguration) SetEnabledOnRevocation(v bool)`

SetEnabledOnRevocation sets EnabledOnRevocation field to given value.


### GetTicketProvider

`func (o *TicketPropagationConfiguration) GetTicketProvider() TicketingProviderEnum`

GetTicketProvider returns the TicketProvider field if non-nil, zero value otherwise.

### GetTicketProviderOk

`func (o *TicketPropagationConfiguration) GetTicketProviderOk() (*TicketingProviderEnum, bool)`

GetTicketProviderOk returns a tuple with the TicketProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketProvider

`func (o *TicketPropagationConfiguration) SetTicketProvider(v TicketingProviderEnum)`

SetTicketProvider sets TicketProvider field to given value.

### HasTicketProvider

`func (o *TicketPropagationConfiguration) HasTicketProvider() bool`

HasTicketProvider returns a boolean if a field has been set.

### GetTicketProjectId

`func (o *TicketPropagationConfiguration) GetTicketProjectId() string`

GetTicketProjectId returns the TicketProjectId field if non-nil, zero value otherwise.

### GetTicketProjectIdOk

`func (o *TicketPropagationConfiguration) GetTicketProjectIdOk() (*string, bool)`

GetTicketProjectIdOk returns a tuple with the TicketProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketProjectId

`func (o *TicketPropagationConfiguration) SetTicketProjectId(v string)`

SetTicketProjectId sets TicketProjectId field to given value.

### HasTicketProjectId

`func (o *TicketPropagationConfiguration) HasTicketProjectId() bool`

HasTicketProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


