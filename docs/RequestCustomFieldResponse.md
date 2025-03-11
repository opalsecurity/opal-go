# RequestCustomFieldResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** |  | 
**FieldType** | [**RequestTemplateCustomFieldTypeEnum**](RequestTemplateCustomFieldTypeEnum.md) |  | 
**FieldValue** | [**RequestCustomFieldResponseFieldValue**](RequestCustomFieldResponseFieldValue.md) |  | 

## Methods

### NewRequestCustomFieldResponse

`func NewRequestCustomFieldResponse(fieldName string, fieldType RequestTemplateCustomFieldTypeEnum, fieldValue RequestCustomFieldResponseFieldValue, ) *RequestCustomFieldResponse`

NewRequestCustomFieldResponse instantiates a new RequestCustomFieldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCustomFieldResponseWithDefaults

`func NewRequestCustomFieldResponseWithDefaults() *RequestCustomFieldResponse`

NewRequestCustomFieldResponseWithDefaults instantiates a new RequestCustomFieldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *RequestCustomFieldResponse) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *RequestCustomFieldResponse) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *RequestCustomFieldResponse) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetFieldType

`func (o *RequestCustomFieldResponse) GetFieldType() RequestTemplateCustomFieldTypeEnum`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *RequestCustomFieldResponse) GetFieldTypeOk() (*RequestTemplateCustomFieldTypeEnum, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *RequestCustomFieldResponse) SetFieldType(v RequestTemplateCustomFieldTypeEnum)`

SetFieldType sets FieldType field to given value.


### GetFieldValue

`func (o *RequestCustomFieldResponse) GetFieldValue() RequestCustomFieldResponseFieldValue`

GetFieldValue returns the FieldValue field if non-nil, zero value otherwise.

### GetFieldValueOk

`func (o *RequestCustomFieldResponse) GetFieldValueOk() (*RequestCustomFieldResponseFieldValue, bool)`

GetFieldValueOk returns a tuple with the FieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValue

`func (o *RequestCustomFieldResponse) SetFieldValue(v RequestCustomFieldResponseFieldValue)`

SetFieldValue sets FieldValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


