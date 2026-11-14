# CreateTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | [**Token**](Token.md) |  | 
**SignedToken** | **string** | The secret token string. Store this securely; it is only returned once at creation time.  | 

## Methods

### NewCreateTokenResponse

`func NewCreateTokenResponse(token Token, signedToken string, ) *CreateTokenResponse`

NewCreateTokenResponse instantiates a new CreateTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokenResponseWithDefaults

`func NewCreateTokenResponseWithDefaults() *CreateTokenResponse`

NewCreateTokenResponseWithDefaults instantiates a new CreateTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *CreateTokenResponse) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateTokenResponse) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateTokenResponse) SetToken(v Token)`

SetToken sets Token field to given value.


### GetSignedToken

`func (o *CreateTokenResponse) GetSignedToken() string`

GetSignedToken returns the SignedToken field if non-nil, zero value otherwise.

### GetSignedTokenOk

`func (o *CreateTokenResponse) GetSignedTokenOk() (*string, bool)`

GetSignedTokenOk returns a tuple with the SignedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedToken

`func (o *CreateTokenResponse) SetSignedToken(v string)`

SetSignedToken sets SignedToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


