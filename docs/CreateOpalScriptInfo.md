# CreateOpalScriptInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the OpalScript. | 
**ScriptType** | [**OpalScriptTypeEnum**](OpalScriptTypeEnum.md) |  | 
**Script** | **string** | The source of the OpalScript. | 
**OwnerId** | Pointer to **string** | The owner of the service user Opal provisions for a &#x60;REQUEST_REVIEW&#x60; script. Required for that script type and ignored for the others, which have no reviewer identity. | [optional] 

## Methods

### NewCreateOpalScriptInfo

`func NewCreateOpalScriptInfo(name string, scriptType OpalScriptTypeEnum, script string, ) *CreateOpalScriptInfo`

NewCreateOpalScriptInfo instantiates a new CreateOpalScriptInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOpalScriptInfoWithDefaults

`func NewCreateOpalScriptInfoWithDefaults() *CreateOpalScriptInfo`

NewCreateOpalScriptInfoWithDefaults instantiates a new CreateOpalScriptInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOpalScriptInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOpalScriptInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOpalScriptInfo) SetName(v string)`

SetName sets Name field to given value.


### GetScriptType

`func (o *CreateOpalScriptInfo) GetScriptType() OpalScriptTypeEnum`

GetScriptType returns the ScriptType field if non-nil, zero value otherwise.

### GetScriptTypeOk

`func (o *CreateOpalScriptInfo) GetScriptTypeOk() (*OpalScriptTypeEnum, bool)`

GetScriptTypeOk returns a tuple with the ScriptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptType

`func (o *CreateOpalScriptInfo) SetScriptType(v OpalScriptTypeEnum)`

SetScriptType sets ScriptType field to given value.


### GetScript

`func (o *CreateOpalScriptInfo) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *CreateOpalScriptInfo) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *CreateOpalScriptInfo) SetScript(v string)`

SetScript sets Script field to given value.


### GetOwnerId

`func (o *CreateOpalScriptInfo) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CreateOpalScriptInfo) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CreateOpalScriptInfo) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *CreateOpalScriptInfo) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


