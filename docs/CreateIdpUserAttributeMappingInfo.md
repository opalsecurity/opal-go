# CreateIdpUserAttributeMappingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The attribute key in the identity provider or HRIS. | 
**UseAs** | [**IdpUserAttributeUseAsEnum**](IdpUserAttributeUseAsEnum.md) |  | 

## Methods

### NewCreateIdpUserAttributeMappingInfo

`func NewCreateIdpUserAttributeMappingInfo(key string, useAs IdpUserAttributeUseAsEnum, ) *CreateIdpUserAttributeMappingInfo`

NewCreateIdpUserAttributeMappingInfo instantiates a new CreateIdpUserAttributeMappingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIdpUserAttributeMappingInfoWithDefaults

`func NewCreateIdpUserAttributeMappingInfoWithDefaults() *CreateIdpUserAttributeMappingInfo`

NewCreateIdpUserAttributeMappingInfoWithDefaults instantiates a new CreateIdpUserAttributeMappingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CreateIdpUserAttributeMappingInfo) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateIdpUserAttributeMappingInfo) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateIdpUserAttributeMappingInfo) SetKey(v string)`

SetKey sets Key field to given value.


### GetUseAs

`func (o *CreateIdpUserAttributeMappingInfo) GetUseAs() IdpUserAttributeUseAsEnum`

GetUseAs returns the UseAs field if non-nil, zero value otherwise.

### GetUseAsOk

`func (o *CreateIdpUserAttributeMappingInfo) GetUseAsOk() (*IdpUserAttributeUseAsEnum, bool)`

GetUseAsOk returns a tuple with the UseAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAs

`func (o *CreateIdpUserAttributeMappingInfo) SetUseAs(v IdpUserAttributeUseAsEnum)`

SetUseAs sets UseAs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


