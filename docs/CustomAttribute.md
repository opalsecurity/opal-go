# CustomAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The name of the custom attribute. | 
**ValueType** | [**CustomAttributeValueTypeEnum**](CustomAttributeValueTypeEnum.md) |  | 
**Values** | **[]string** | The attribute&#39;s value(s), encoded as strings. | 

## Methods

### NewCustomAttribute

`func NewCustomAttribute(key string, valueType CustomAttributeValueTypeEnum, values []string, ) *CustomAttribute`

NewCustomAttribute instantiates a new CustomAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeWithDefaults

`func NewCustomAttributeWithDefaults() *CustomAttribute`

NewCustomAttributeWithDefaults instantiates a new CustomAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CustomAttribute) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomAttribute) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomAttribute) SetKey(v string)`

SetKey sets Key field to given value.


### GetValueType

`func (o *CustomAttribute) GetValueType() CustomAttributeValueTypeEnum`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *CustomAttribute) GetValueTypeOk() (*CustomAttributeValueTypeEnum, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *CustomAttribute) SetValueType(v CustomAttributeValueTypeEnum)`

SetValueType sets ValueType field to given value.


### GetValues

`func (o *CustomAttribute) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CustomAttribute) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CustomAttribute) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


