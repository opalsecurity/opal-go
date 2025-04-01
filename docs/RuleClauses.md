# RuleClauses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**When** | [**RuleConjunction**](RuleConjunction.md) |  | 
**Unless** | Pointer to [**RuleConjunction**](RuleConjunction.md) |  | [optional] 

## Methods

### NewRuleClauses

`func NewRuleClauses(when RuleConjunction, ) *RuleClauses`

NewRuleClauses instantiates a new RuleClauses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleClausesWithDefaults

`func NewRuleClausesWithDefaults() *RuleClauses`

NewRuleClausesWithDefaults instantiates a new RuleClauses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhen

`func (o *RuleClauses) GetWhen() RuleConjunction`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *RuleClauses) GetWhenOk() (*RuleConjunction, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *RuleClauses) SetWhen(v RuleConjunction)`

SetWhen sets When field to given value.


### GetUnless

`func (o *RuleClauses) GetUnless() RuleConjunction`

GetUnless returns the Unless field if non-nil, zero value otherwise.

### GetUnlessOk

`func (o *RuleClauses) GetUnlessOk() (*RuleConjunction, bool)`

GetUnlessOk returns a tuple with the Unless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnless

`func (o *RuleClauses) SetUnless(v RuleConjunction)`

SetUnless sets Unless field to given value.

### HasUnless

`func (o *RuleClauses) HasUnless() bool`

HasUnless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


