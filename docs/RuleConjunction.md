# RuleConjunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clauses** | [**[]RuleDisjunction**](RuleDisjunction.md) |  | 

## Methods

### NewRuleConjunction

`func NewRuleConjunction(clauses []RuleDisjunction, ) *RuleConjunction`

NewRuleConjunction instantiates a new RuleConjunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleConjunctionWithDefaults

`func NewRuleConjunctionWithDefaults() *RuleConjunction`

NewRuleConjunctionWithDefaults instantiates a new RuleConjunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClauses

`func (o *RuleConjunction) GetClauses() []RuleDisjunction`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *RuleConjunction) GetClausesOk() (*[]RuleDisjunction, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *RuleConjunction) SetClauses(v []RuleDisjunction)`

SetClauses sets Clauses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


