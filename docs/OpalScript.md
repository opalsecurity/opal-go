# OpalScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpalScriptId** | **string** | The ID of the OpalScript. | 
**Name** | **string** | The name of the OpalScript. | 
**ScriptType** | [**OpalScriptTypeEnum**](OpalScriptTypeEnum.md) |  | 
**Script** | **string** | The source of the OpalScript&#39;s latest version. | 
**Version** | **int32** | The version number of the OpalScript&#39;s latest version. | 
**UpdatedAt** | Pointer to **time.Time** | The date the OpalScript was last updated. | [optional] 

## Methods

### NewOpalScript

`func NewOpalScript(opalScriptId string, name string, scriptType OpalScriptTypeEnum, script string, version int32, ) *OpalScript`

NewOpalScript instantiates a new OpalScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalScriptWithDefaults

`func NewOpalScriptWithDefaults() *OpalScript`

NewOpalScriptWithDefaults instantiates a new OpalScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpalScriptId

`func (o *OpalScript) GetOpalScriptId() string`

GetOpalScriptId returns the OpalScriptId field if non-nil, zero value otherwise.

### GetOpalScriptIdOk

`func (o *OpalScript) GetOpalScriptIdOk() (*string, bool)`

GetOpalScriptIdOk returns a tuple with the OpalScriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpalScriptId

`func (o *OpalScript) SetOpalScriptId(v string)`

SetOpalScriptId sets OpalScriptId field to given value.


### GetName

`func (o *OpalScript) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpalScript) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpalScript) SetName(v string)`

SetName sets Name field to given value.


### GetScriptType

`func (o *OpalScript) GetScriptType() OpalScriptTypeEnum`

GetScriptType returns the ScriptType field if non-nil, zero value otherwise.

### GetScriptTypeOk

`func (o *OpalScript) GetScriptTypeOk() (*OpalScriptTypeEnum, bool)`

GetScriptTypeOk returns a tuple with the ScriptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptType

`func (o *OpalScript) SetScriptType(v OpalScriptTypeEnum)`

SetScriptType sets ScriptType field to given value.


### GetScript

`func (o *OpalScript) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *OpalScript) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *OpalScript) SetScript(v string)`

SetScript sets Script field to given value.


### GetVersion

`func (o *OpalScript) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OpalScript) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OpalScript) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetUpdatedAt

`func (o *OpalScript) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OpalScript) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OpalScript) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OpalScript) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


