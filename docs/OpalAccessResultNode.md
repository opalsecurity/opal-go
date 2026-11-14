# OpalAccessResultNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalId** | **string** | The principal entity ID. | 
**EntitlementId** | **string** | The entitlement entity ID. | 
**AccessLevelRemoteId** | Pointer to **string** | Remote ID of the access level. | [optional] 
**AccessLevelName** | Pointer to **string** | Display name of the access level. | [optional] 
**EffectiveExpiration** | Pointer to **time.Time** | The expiration in effect across all paths (the latest-expiring; a permanent path wins). Null means access never expires. | [optional] 
**IsDirect** | **bool** | True when any path is a direct grant. | 
**Paths** | [**[]OpalAccessPathDetail**](OpalAccessPathDetail.md) | Every access path reaching this grant. | 

## Methods

### NewOpalAccessResultNode

`func NewOpalAccessResultNode(principalId string, entitlementId string, isDirect bool, paths []OpalAccessPathDetail, ) *OpalAccessResultNode`

NewOpalAccessResultNode instantiates a new OpalAccessResultNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalAccessResultNodeWithDefaults

`func NewOpalAccessResultNodeWithDefaults() *OpalAccessResultNode`

NewOpalAccessResultNodeWithDefaults instantiates a new OpalAccessResultNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalId

`func (o *OpalAccessResultNode) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *OpalAccessResultNode) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *OpalAccessResultNode) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetEntitlementId

`func (o *OpalAccessResultNode) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *OpalAccessResultNode) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *OpalAccessResultNode) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetAccessLevelRemoteId

`func (o *OpalAccessResultNode) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *OpalAccessResultNode) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *OpalAccessResultNode) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *OpalAccessResultNode) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.

### GetAccessLevelName

`func (o *OpalAccessResultNode) GetAccessLevelName() string`

GetAccessLevelName returns the AccessLevelName field if non-nil, zero value otherwise.

### GetAccessLevelNameOk

`func (o *OpalAccessResultNode) GetAccessLevelNameOk() (*string, bool)`

GetAccessLevelNameOk returns a tuple with the AccessLevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelName

`func (o *OpalAccessResultNode) SetAccessLevelName(v string)`

SetAccessLevelName sets AccessLevelName field to given value.

### HasAccessLevelName

`func (o *OpalAccessResultNode) HasAccessLevelName() bool`

HasAccessLevelName returns a boolean if a field has been set.

### GetEffectiveExpiration

`func (o *OpalAccessResultNode) GetEffectiveExpiration() time.Time`

GetEffectiveExpiration returns the EffectiveExpiration field if non-nil, zero value otherwise.

### GetEffectiveExpirationOk

`func (o *OpalAccessResultNode) GetEffectiveExpirationOk() (*time.Time, bool)`

GetEffectiveExpirationOk returns a tuple with the EffectiveExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveExpiration

`func (o *OpalAccessResultNode) SetEffectiveExpiration(v time.Time)`

SetEffectiveExpiration sets EffectiveExpiration field to given value.

### HasEffectiveExpiration

`func (o *OpalAccessResultNode) HasEffectiveExpiration() bool`

HasEffectiveExpiration returns a boolean if a field has been set.

### GetIsDirect

`func (o *OpalAccessResultNode) GetIsDirect() bool`

GetIsDirect returns the IsDirect field if non-nil, zero value otherwise.

### GetIsDirectOk

`func (o *OpalAccessResultNode) GetIsDirectOk() (*bool, bool)`

GetIsDirectOk returns a tuple with the IsDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirect

`func (o *OpalAccessResultNode) SetIsDirect(v bool)`

SetIsDirect sets IsDirect field to given value.


### GetPaths

`func (o *OpalAccessResultNode) GetPaths() []OpalAccessPathDetail`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *OpalAccessResultNode) GetPathsOk() (*[]OpalAccessPathDetail, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *OpalAccessResultNode) SetPaths(v []OpalAccessPathDetail)`

SetPaths sets Paths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


