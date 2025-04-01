# AccessRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of the access rule. | 
**RuleClauses** | [**RuleClauses**](RuleClauses.md) |  | 

## Methods

### NewAccessRuleCondition

`func NewAccessRuleCondition(status string, ruleClauses RuleClauses, ) *AccessRuleCondition`

NewAccessRuleCondition instantiates a new AccessRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRuleConditionWithDefaults

`func NewAccessRuleConditionWithDefaults() *AccessRuleCondition`

NewAccessRuleConditionWithDefaults instantiates a new AccessRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AccessRuleCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessRuleCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessRuleCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRuleClauses

`func (o *AccessRuleCondition) GetRuleClauses() RuleClauses`

GetRuleClauses returns the RuleClauses field if non-nil, zero value otherwise.

### GetRuleClausesOk

`func (o *AccessRuleCondition) GetRuleClausesOk() (*RuleClauses, bool)`

GetRuleClausesOk returns a tuple with the RuleClauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleClauses

`func (o *AccessRuleCondition) SetRuleClauses(v RuleClauses)`

SetRuleClauses sets RuleClauses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


