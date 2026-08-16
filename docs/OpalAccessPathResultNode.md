# OpalAccessPathResultNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalId** | **string** | The principal entity ID. | 
**EntitlementId** | **string** | The entitlement entity ID. | 
**AccessLevelRemoteId** | Pointer to **string** | Remote ID of the terminal access level. | [optional] 
**AccessLevelName** | Pointer to **string** | Display name of the terminal access level. | [optional] 
**Expiration** | Pointer to **time.Time** | Expiration of the terminal access, if any. | [optional] 
**Depth** | **int32** | Number of hops from principal to entitlement (path length - 1). | 
**Path** | **[]string** | Entity IDs along the path from principal to entitlement. | 

## Methods

### NewOpalAccessPathResultNode

`func NewOpalAccessPathResultNode(principalId string, entitlementId string, depth int32, path []string, ) *OpalAccessPathResultNode`

NewOpalAccessPathResultNode instantiates a new OpalAccessPathResultNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalAccessPathResultNodeWithDefaults

`func NewOpalAccessPathResultNodeWithDefaults() *OpalAccessPathResultNode`

NewOpalAccessPathResultNodeWithDefaults instantiates a new OpalAccessPathResultNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalId

`func (o *OpalAccessPathResultNode) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *OpalAccessPathResultNode) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *OpalAccessPathResultNode) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetEntitlementId

`func (o *OpalAccessPathResultNode) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *OpalAccessPathResultNode) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *OpalAccessPathResultNode) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetAccessLevelRemoteId

`func (o *OpalAccessPathResultNode) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *OpalAccessPathResultNode) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *OpalAccessPathResultNode) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *OpalAccessPathResultNode) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.

### GetAccessLevelName

`func (o *OpalAccessPathResultNode) GetAccessLevelName() string`

GetAccessLevelName returns the AccessLevelName field if non-nil, zero value otherwise.

### GetAccessLevelNameOk

`func (o *OpalAccessPathResultNode) GetAccessLevelNameOk() (*string, bool)`

GetAccessLevelNameOk returns a tuple with the AccessLevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelName

`func (o *OpalAccessPathResultNode) SetAccessLevelName(v string)`

SetAccessLevelName sets AccessLevelName field to given value.

### HasAccessLevelName

`func (o *OpalAccessPathResultNode) HasAccessLevelName() bool`

HasAccessLevelName returns a boolean if a field has been set.

### GetExpiration

`func (o *OpalAccessPathResultNode) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *OpalAccessPathResultNode) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *OpalAccessPathResultNode) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *OpalAccessPathResultNode) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetDepth

`func (o *OpalAccessPathResultNode) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *OpalAccessPathResultNode) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *OpalAccessPathResultNode) SetDepth(v int32)`

SetDepth sets Depth field to given value.


### GetPath

`func (o *OpalAccessPathResultNode) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OpalAccessPathResultNode) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OpalAccessPathResultNode) SetPath(v []string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


