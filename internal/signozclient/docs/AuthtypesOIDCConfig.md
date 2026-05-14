# AuthtypesOIDCConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimMapping** | Pointer to [**AuthtypesAttributeMapping**](AuthtypesAttributeMapping.md) |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**GetUserInfo** | Pointer to **bool** |  | [optional] 
**InsecureSkipEmailVerified** | Pointer to **bool** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**IssuerAlias** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthtypesOIDCConfig

`func NewAuthtypesOIDCConfig() *AuthtypesOIDCConfig`

NewAuthtypesOIDCConfig instantiates a new AuthtypesOIDCConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthtypesOIDCConfigWithDefaults

`func NewAuthtypesOIDCConfigWithDefaults() *AuthtypesOIDCConfig`

NewAuthtypesOIDCConfigWithDefaults instantiates a new AuthtypesOIDCConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimMapping

`func (o *AuthtypesOIDCConfig) GetClaimMapping() AuthtypesAttributeMapping`

GetClaimMapping returns the ClaimMapping field if non-nil, zero value otherwise.

### GetClaimMappingOk

`func (o *AuthtypesOIDCConfig) GetClaimMappingOk() (*AuthtypesAttributeMapping, bool)`

GetClaimMappingOk returns a tuple with the ClaimMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimMapping

`func (o *AuthtypesOIDCConfig) SetClaimMapping(v AuthtypesAttributeMapping)`

SetClaimMapping sets ClaimMapping field to given value.

### HasClaimMapping

`func (o *AuthtypesOIDCConfig) HasClaimMapping() bool`

HasClaimMapping returns a boolean if a field has been set.

### GetClientId

`func (o *AuthtypesOIDCConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AuthtypesOIDCConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AuthtypesOIDCConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AuthtypesOIDCConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *AuthtypesOIDCConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AuthtypesOIDCConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AuthtypesOIDCConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AuthtypesOIDCConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetGetUserInfo

`func (o *AuthtypesOIDCConfig) GetGetUserInfo() bool`

GetGetUserInfo returns the GetUserInfo field if non-nil, zero value otherwise.

### GetGetUserInfoOk

`func (o *AuthtypesOIDCConfig) GetGetUserInfoOk() (*bool, bool)`

GetGetUserInfoOk returns a tuple with the GetUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetUserInfo

`func (o *AuthtypesOIDCConfig) SetGetUserInfo(v bool)`

SetGetUserInfo sets GetUserInfo field to given value.

### HasGetUserInfo

`func (o *AuthtypesOIDCConfig) HasGetUserInfo() bool`

HasGetUserInfo returns a boolean if a field has been set.

### GetInsecureSkipEmailVerified

`func (o *AuthtypesOIDCConfig) GetInsecureSkipEmailVerified() bool`

GetInsecureSkipEmailVerified returns the InsecureSkipEmailVerified field if non-nil, zero value otherwise.

### GetInsecureSkipEmailVerifiedOk

`func (o *AuthtypesOIDCConfig) GetInsecureSkipEmailVerifiedOk() (*bool, bool)`

GetInsecureSkipEmailVerifiedOk returns a tuple with the InsecureSkipEmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipEmailVerified

`func (o *AuthtypesOIDCConfig) SetInsecureSkipEmailVerified(v bool)`

SetInsecureSkipEmailVerified sets InsecureSkipEmailVerified field to given value.

### HasInsecureSkipEmailVerified

`func (o *AuthtypesOIDCConfig) HasInsecureSkipEmailVerified() bool`

HasInsecureSkipEmailVerified returns a boolean if a field has been set.

### GetIssuer

`func (o *AuthtypesOIDCConfig) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *AuthtypesOIDCConfig) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *AuthtypesOIDCConfig) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *AuthtypesOIDCConfig) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetIssuerAlias

`func (o *AuthtypesOIDCConfig) GetIssuerAlias() string`

GetIssuerAlias returns the IssuerAlias field if non-nil, zero value otherwise.

### GetIssuerAliasOk

`func (o *AuthtypesOIDCConfig) GetIssuerAliasOk() (*string, bool)`

GetIssuerAliasOk returns a tuple with the IssuerAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerAlias

`func (o *AuthtypesOIDCConfig) SetIssuerAlias(v string)`

SetIssuerAlias sets IssuerAlias field to given value.

### HasIssuerAlias

`func (o *AuthtypesOIDCConfig) HasIssuerAlias() bool`

HasIssuerAlias returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


