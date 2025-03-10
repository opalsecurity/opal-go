# AppValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the app validation. These are not unique IDs between runs. | 
**Name** | **interface{}** |  | 
**UsageReason** | Pointer to **string** | The reason for needing the validation. | [optional] 
**Details** | Pointer to **string** | Extra details regarding the validation. Could be an error message or restrictions on permissions. | [optional] 
**Severity** | [**AppValidationSeverityEnum**](AppValidationSeverityEnum.md) |  | 
**Status** | [**AppValidationStatusEnum**](AppValidationStatusEnum.md) |  | 
**UpdatedAt** | **time.Time** | The date and time the app validation was last run. | 

## Methods

### NewAppValidation

`func NewAppValidation(key string, name interface{}, severity AppValidationSeverityEnum, status AppValidationStatusEnum, updatedAt time.Time, ) *AppValidation`

NewAppValidation instantiates a new AppValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppValidationWithDefaults

`func NewAppValidationWithDefaults() *AppValidation`

NewAppValidationWithDefaults instantiates a new AppValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AppValidation) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AppValidation) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AppValidation) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *AppValidation) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppValidation) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppValidation) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *AppValidation) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AppValidation) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUsageReason

`func (o *AppValidation) GetUsageReason() string`

GetUsageReason returns the UsageReason field if non-nil, zero value otherwise.

### GetUsageReasonOk

`func (o *AppValidation) GetUsageReasonOk() (*string, bool)`

GetUsageReasonOk returns a tuple with the UsageReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageReason

`func (o *AppValidation) SetUsageReason(v string)`

SetUsageReason sets UsageReason field to given value.

### HasUsageReason

`func (o *AppValidation) HasUsageReason() bool`

HasUsageReason returns a boolean if a field has been set.

### GetDetails

`func (o *AppValidation) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AppValidation) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AppValidation) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AppValidation) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetSeverity

`func (o *AppValidation) GetSeverity() AppValidationSeverityEnum`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AppValidation) GetSeverityOk() (*AppValidationSeverityEnum, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AppValidation) SetSeverity(v AppValidationSeverityEnum)`

SetSeverity sets Severity field to given value.


### GetStatus

`func (o *AppValidation) GetStatus() AppValidationStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppValidation) GetStatusOk() (*AppValidationStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppValidation) SetStatus(v AppValidationStatusEnum)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *AppValidation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppValidation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppValidation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


