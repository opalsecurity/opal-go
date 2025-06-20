# UpdateAccessRuleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the access rule. | 
**Description** | **string** | A description of the group. | 
**AdminOwnerId** | **string** | The ID of the owner of the group. | 
**Status** | **string** | The status of the access rule. | 
**RuleClauses** | [**RuleClauses**](RuleClauses.md) |  | 

## Methods

### NewUpdateAccessRuleInfo

`func NewUpdateAccessRuleInfo(name string, description string, adminOwnerId string, status string, ruleClauses RuleClauses, ) *UpdateAccessRuleInfo`

NewUpdateAccessRuleInfo instantiates a new UpdateAccessRuleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccessRuleInfoWithDefaults

`func NewUpdateAccessRuleInfoWithDefaults() *UpdateAccessRuleInfo`

NewUpdateAccessRuleInfoWithDefaults instantiates a new UpdateAccessRuleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAccessRuleInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAccessRuleInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAccessRuleInfo) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateAccessRuleInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAccessRuleInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAccessRuleInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAdminOwnerId

`func (o *UpdateAccessRuleInfo) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *UpdateAccessRuleInfo) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *UpdateAccessRuleInfo) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.


### GetStatus

`func (o *UpdateAccessRuleInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAccessRuleInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAccessRuleInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRuleClauses

`func (o *UpdateAccessRuleInfo) GetRuleClauses() RuleClauses`

GetRuleClauses returns the RuleClauses field if non-nil, zero value otherwise.

### GetRuleClausesOk

`func (o *UpdateAccessRuleInfo) GetRuleClausesOk() (*RuleClauses, bool)`

GetRuleClausesOk returns a tuple with the RuleClauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleClauses

`func (o *UpdateAccessRuleInfo) SetRuleClauses(v RuleClauses)`

SetRuleClauses sets RuleClauses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


