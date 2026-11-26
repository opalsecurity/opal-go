# UpdateAccessRuleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the access rule. | 
**Description** | **string** | A description of the group. | 
**AdminOwnerId** | **string** | The ID of the owner of the group. | 
**Status** | **string** | The status of the access rule. | 
**RuleClauses** | Pointer to [**RuleClauses**](RuleClauses.md) | Tag-based CNF conditions. May be sent with &#x60;query&#x60; (AND). On PUT, omitting this field while sending &#x60;query&#x60; clears tags.  | [optional] 
**Query** | Pointer to [**OpalNodeQueryBody**](OpalNodeQueryBody.md) | Opal Query defining the user set (same NODE filter body as &#x60;POST /queries/run&#x60;). May be sent with &#x60;ruleClauses&#x60; (AND). On PUT, omitting this field while sending &#x60;ruleClauses&#x60; clears the query. Nested AND/OR, job title / team / manager, and has-access filters are supported.  | [optional] 

## Methods

### NewUpdateAccessRuleInfo

`func NewUpdateAccessRuleInfo(name string, description string, adminOwnerId string, status string, ) *UpdateAccessRuleInfo`

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

### HasRuleClauses

`func (o *UpdateAccessRuleInfo) HasRuleClauses() bool`

HasRuleClauses returns a boolean if a field has been set.

### GetQuery

`func (o *UpdateAccessRuleInfo) GetQuery() OpalNodeQueryBody`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *UpdateAccessRuleInfo) GetQueryOk() (*OpalNodeQueryBody, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *UpdateAccessRuleInfo) SetQuery(v OpalNodeQueryBody)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *UpdateAccessRuleInfo) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


