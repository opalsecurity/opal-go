# IdpUserAttributeMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the mapping. | 
**AppId** | **string** | The ID of the identity provider or HRIS app. | 
**Key** | **string** | The attribute key in the identity provider or HRIS. | 
**UseAs** | [**IdpUserAttributeUseAsEnum**](IdpUserAttributeUseAsEnum.md) |  | 

## Methods

### NewIdpUserAttributeMapping

`func NewIdpUserAttributeMapping(id string, appId string, key string, useAs IdpUserAttributeUseAsEnum, ) *IdpUserAttributeMapping`

NewIdpUserAttributeMapping instantiates a new IdpUserAttributeMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpUserAttributeMappingWithDefaults

`func NewIdpUserAttributeMappingWithDefaults() *IdpUserAttributeMapping`

NewIdpUserAttributeMappingWithDefaults instantiates a new IdpUserAttributeMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdpUserAttributeMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdpUserAttributeMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdpUserAttributeMapping) SetId(v string)`

SetId sets Id field to given value.


### GetAppId

`func (o *IdpUserAttributeMapping) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *IdpUserAttributeMapping) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *IdpUserAttributeMapping) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetKey

`func (o *IdpUserAttributeMapping) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IdpUserAttributeMapping) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IdpUserAttributeMapping) SetKey(v string)`

SetKey sets Key field to given value.


### GetUseAs

`func (o *IdpUserAttributeMapping) GetUseAs() IdpUserAttributeUseAsEnum`

GetUseAs returns the UseAs field if non-nil, zero value otherwise.

### GetUseAsOk

`func (o *IdpUserAttributeMapping) GetUseAsOk() (*IdpUserAttributeUseAsEnum, bool)`

GetUseAsOk returns a tuple with the UseAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAs

`func (o *IdpUserAttributeMapping) SetUseAs(v IdpUserAttributeUseAsEnum)`

SetUseAs sets UseAs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


