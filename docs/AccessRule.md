# AccessRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessRuleId** | **string** | The ID (group ID) of the access rule. | 
**Name** | **string** | The name of the access rule. | 
**Description** | **string** | A description of the group. | 
**AdminOwnerId** | **string** | The ID of the owner of the group. | 
**Status** | **string** | The status of the access rule. | 
**RuleClauses** | [**RuleClauses**](RuleClauses.md) | Tag-based CNF conditions. Present for tag-only rules, tags-as-query rules, and hybrids (tags AND an Opal Query). Empty &#x60;when.clauses&#x60; for pure Opal Query rules.  | 
**Query** | Pointer to [**OpalNodeQueryBody**](OpalNodeQueryBody.md) | Opal Query defining the user set when this rule uses one. For hybrids, both &#x60;ruleClauses&#x60; and &#x60;query&#x60; are set (AND). Omitted for tag-only rules and when the REST dialect cannot express the stored filter. Same NODE filter body as &#x60;POST /queries/run&#x60;.  | [optional] 
**UsesOpalQuery** | Pointer to **bool** | True when the access rule stores an Opal Query (including hybrids). &#x60;query&#x60; is set when the REST dialect can express that filter; it is omitted for filters the API cannot express (e.g. risk, identity detection, agent fields). Rely on this flag, not only on whether &#x60;query&#x60; is present.  | [optional] 

## Methods

### NewAccessRule

`func NewAccessRule(accessRuleId string, name string, description string, adminOwnerId string, status string, ruleClauses RuleClauses, ) *AccessRule`

NewAccessRule instantiates a new AccessRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRuleWithDefaults

`func NewAccessRuleWithDefaults() *AccessRule`

NewAccessRuleWithDefaults instantiates a new AccessRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessRuleId

`func (o *AccessRule) GetAccessRuleId() string`

GetAccessRuleId returns the AccessRuleId field if non-nil, zero value otherwise.

### GetAccessRuleIdOk

`func (o *AccessRule) GetAccessRuleIdOk() (*string, bool)`

GetAccessRuleIdOk returns a tuple with the AccessRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRuleId

`func (o *AccessRule) SetAccessRuleId(v string)`

SetAccessRuleId sets AccessRuleId field to given value.


### GetName

`func (o *AccessRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessRule) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AccessRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessRule) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAdminOwnerId

`func (o *AccessRule) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *AccessRule) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *AccessRule) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.


### GetStatus

`func (o *AccessRule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessRule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessRule) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRuleClauses

`func (o *AccessRule) GetRuleClauses() RuleClauses`

GetRuleClauses returns the RuleClauses field if non-nil, zero value otherwise.

### GetRuleClausesOk

`func (o *AccessRule) GetRuleClausesOk() (*RuleClauses, bool)`

GetRuleClausesOk returns a tuple with the RuleClauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleClauses

`func (o *AccessRule) SetRuleClauses(v RuleClauses)`

SetRuleClauses sets RuleClauses field to given value.


### GetQuery

`func (o *AccessRule) GetQuery() OpalNodeQueryBody`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AccessRule) GetQueryOk() (*OpalNodeQueryBody, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AccessRule) SetQuery(v OpalNodeQueryBody)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *AccessRule) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetUsesOpalQuery

`func (o *AccessRule) GetUsesOpalQuery() bool`

GetUsesOpalQuery returns the UsesOpalQuery field if non-nil, zero value otherwise.

### GetUsesOpalQueryOk

`func (o *AccessRule) GetUsesOpalQueryOk() (*bool, bool)`

GetUsesOpalQueryOk returns a tuple with the UsesOpalQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesOpalQuery

`func (o *AccessRule) SetUsesOpalQuery(v bool)`

SetUsesOpalQuery sets UsesOpalQuery field to given value.

### HasUsesOpalQuery

`func (o *AccessRule) HasUsesOpalQuery() bool`

HasUsesOpalQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


