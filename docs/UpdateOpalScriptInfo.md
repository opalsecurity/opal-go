# UpdateOpalScriptInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the OpalScript. | [optional] 
**Script** | Pointer to **string** | The source of the OpalScript. A new version is recorded only when the source differs from the current one. | [optional] 

## Methods

### NewUpdateOpalScriptInfo

`func NewUpdateOpalScriptInfo() *UpdateOpalScriptInfo`

NewUpdateOpalScriptInfo instantiates a new UpdateOpalScriptInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOpalScriptInfoWithDefaults

`func NewUpdateOpalScriptInfoWithDefaults() *UpdateOpalScriptInfo`

NewUpdateOpalScriptInfoWithDefaults instantiates a new UpdateOpalScriptInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateOpalScriptInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOpalScriptInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOpalScriptInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOpalScriptInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScript

`func (o *UpdateOpalScriptInfo) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *UpdateOpalScriptInfo) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *UpdateOpalScriptInfo) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *UpdateOpalScriptInfo) HasScript() bool`

HasScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


