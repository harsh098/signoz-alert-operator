# AuthtypesAuthDomainConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GoogleAuthConfig** | Pointer to [**AuthtypesGoogleConfig**](AuthtypesGoogleConfig.md) |  | [optional] 
**OidcConfig** | Pointer to [**AuthtypesOIDCConfig**](AuthtypesOIDCConfig.md) |  | [optional] 
**RoleMapping** | Pointer to [**AuthtypesRoleMapping**](AuthtypesRoleMapping.md) |  | [optional] 
**SamlConfig** | Pointer to [**AuthtypesSamlConfig**](AuthtypesSamlConfig.md) |  | [optional] 
**SsoEnabled** | Pointer to **bool** |  | [optional] 
**SsoType** | Pointer to [**AuthtypesAuthNProvider**](AuthtypesAuthNProvider.md) |  | [optional] 
**AttributeMapping** | Pointer to [**AuthtypesAttributeMapping**](AuthtypesAttributeMapping.md) |  | [optional] 
**InsecureSkipAuthNRequestsSigned** | Pointer to **bool** |  | [optional] 
**SamlCert** | Pointer to **string** |  | [optional] 
**SamlEntity** | Pointer to **string** |  | [optional] 
**SamlIdp** | Pointer to **string** |  | [optional] 
**AllowedGroups** | Pointer to **[]string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**DomainToAdminEmail** | Pointer to **map[string]string** |  | [optional] 
**FetchGroups** | Pointer to **bool** |  | [optional] 
**FetchTransitiveGroupMembership** | Pointer to **bool** |  | [optional] 
**InsecureSkipEmailVerified** | Pointer to **bool** |  | [optional] 
**RedirectURI** | Pointer to **string** |  | [optional] 
**ServiceAccountJson** | Pointer to **string** |  | [optional] 
**ClaimMapping** | Pointer to [**AuthtypesAttributeMapping**](AuthtypesAttributeMapping.md) |  | [optional] 
**GetUserInfo** | Pointer to **bool** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**IssuerAlias** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthtypesAuthDomainConfig

`func NewAuthtypesAuthDomainConfig() *AuthtypesAuthDomainConfig`

NewAuthtypesAuthDomainConfig instantiates a new AuthtypesAuthDomainConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthtypesAuthDomainConfigWithDefaults

`func NewAuthtypesAuthDomainConfigWithDefaults() *AuthtypesAuthDomainConfig`

NewAuthtypesAuthDomainConfigWithDefaults instantiates a new AuthtypesAuthDomainConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGoogleAuthConfig

`func (o *AuthtypesAuthDomainConfig) GetGoogleAuthConfig() AuthtypesGoogleConfig`

GetGoogleAuthConfig returns the GoogleAuthConfig field if non-nil, zero value otherwise.

### GetGoogleAuthConfigOk

`func (o *AuthtypesAuthDomainConfig) GetGoogleAuthConfigOk() (*AuthtypesGoogleConfig, bool)`

GetGoogleAuthConfigOk returns a tuple with the GoogleAuthConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAuthConfig

`func (o *AuthtypesAuthDomainConfig) SetGoogleAuthConfig(v AuthtypesGoogleConfig)`

SetGoogleAuthConfig sets GoogleAuthConfig field to given value.

### HasGoogleAuthConfig

`func (o *AuthtypesAuthDomainConfig) HasGoogleAuthConfig() bool`

HasGoogleAuthConfig returns a boolean if a field has been set.

### GetOidcConfig

`func (o *AuthtypesAuthDomainConfig) GetOidcConfig() AuthtypesOIDCConfig`

GetOidcConfig returns the OidcConfig field if non-nil, zero value otherwise.

### GetOidcConfigOk

`func (o *AuthtypesAuthDomainConfig) GetOidcConfigOk() (*AuthtypesOIDCConfig, bool)`

GetOidcConfigOk returns a tuple with the OidcConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcConfig

`func (o *AuthtypesAuthDomainConfig) SetOidcConfig(v AuthtypesOIDCConfig)`

SetOidcConfig sets OidcConfig field to given value.

### HasOidcConfig

`func (o *AuthtypesAuthDomainConfig) HasOidcConfig() bool`

HasOidcConfig returns a boolean if a field has been set.

### GetRoleMapping

`func (o *AuthtypesAuthDomainConfig) GetRoleMapping() AuthtypesRoleMapping`

GetRoleMapping returns the RoleMapping field if non-nil, zero value otherwise.

### GetRoleMappingOk

`func (o *AuthtypesAuthDomainConfig) GetRoleMappingOk() (*AuthtypesRoleMapping, bool)`

GetRoleMappingOk returns a tuple with the RoleMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMapping

`func (o *AuthtypesAuthDomainConfig) SetRoleMapping(v AuthtypesRoleMapping)`

SetRoleMapping sets RoleMapping field to given value.

### HasRoleMapping

`func (o *AuthtypesAuthDomainConfig) HasRoleMapping() bool`

HasRoleMapping returns a boolean if a field has been set.

### GetSamlConfig

`func (o *AuthtypesAuthDomainConfig) GetSamlConfig() AuthtypesSamlConfig`

GetSamlConfig returns the SamlConfig field if non-nil, zero value otherwise.

### GetSamlConfigOk

`func (o *AuthtypesAuthDomainConfig) GetSamlConfigOk() (*AuthtypesSamlConfig, bool)`

GetSamlConfigOk returns a tuple with the SamlConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlConfig

`func (o *AuthtypesAuthDomainConfig) SetSamlConfig(v AuthtypesSamlConfig)`

SetSamlConfig sets SamlConfig field to given value.

### HasSamlConfig

`func (o *AuthtypesAuthDomainConfig) HasSamlConfig() bool`

HasSamlConfig returns a boolean if a field has been set.

### GetSsoEnabled

`func (o *AuthtypesAuthDomainConfig) GetSsoEnabled() bool`

GetSsoEnabled returns the SsoEnabled field if non-nil, zero value otherwise.

### GetSsoEnabledOk

`func (o *AuthtypesAuthDomainConfig) GetSsoEnabledOk() (*bool, bool)`

GetSsoEnabledOk returns a tuple with the SsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEnabled

`func (o *AuthtypesAuthDomainConfig) SetSsoEnabled(v bool)`

SetSsoEnabled sets SsoEnabled field to given value.

### HasSsoEnabled

`func (o *AuthtypesAuthDomainConfig) HasSsoEnabled() bool`

HasSsoEnabled returns a boolean if a field has been set.

### GetSsoType

`func (o *AuthtypesAuthDomainConfig) GetSsoType() AuthtypesAuthNProvider`

GetSsoType returns the SsoType field if non-nil, zero value otherwise.

### GetSsoTypeOk

`func (o *AuthtypesAuthDomainConfig) GetSsoTypeOk() (*AuthtypesAuthNProvider, bool)`

GetSsoTypeOk returns a tuple with the SsoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoType

`func (o *AuthtypesAuthDomainConfig) SetSsoType(v AuthtypesAuthNProvider)`

SetSsoType sets SsoType field to given value.

### HasSsoType

`func (o *AuthtypesAuthDomainConfig) HasSsoType() bool`

HasSsoType returns a boolean if a field has been set.

### GetAttributeMapping

`func (o *AuthtypesAuthDomainConfig) GetAttributeMapping() AuthtypesAttributeMapping`

GetAttributeMapping returns the AttributeMapping field if non-nil, zero value otherwise.

### GetAttributeMappingOk

`func (o *AuthtypesAuthDomainConfig) GetAttributeMappingOk() (*AuthtypesAttributeMapping, bool)`

GetAttributeMappingOk returns a tuple with the AttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMapping

`func (o *AuthtypesAuthDomainConfig) SetAttributeMapping(v AuthtypesAttributeMapping)`

SetAttributeMapping sets AttributeMapping field to given value.

### HasAttributeMapping

`func (o *AuthtypesAuthDomainConfig) HasAttributeMapping() bool`

HasAttributeMapping returns a boolean if a field has been set.

### GetInsecureSkipAuthNRequestsSigned

`func (o *AuthtypesAuthDomainConfig) GetInsecureSkipAuthNRequestsSigned() bool`

GetInsecureSkipAuthNRequestsSigned returns the InsecureSkipAuthNRequestsSigned field if non-nil, zero value otherwise.

### GetInsecureSkipAuthNRequestsSignedOk

`func (o *AuthtypesAuthDomainConfig) GetInsecureSkipAuthNRequestsSignedOk() (*bool, bool)`

GetInsecureSkipAuthNRequestsSignedOk returns a tuple with the InsecureSkipAuthNRequestsSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipAuthNRequestsSigned

`func (o *AuthtypesAuthDomainConfig) SetInsecureSkipAuthNRequestsSigned(v bool)`

SetInsecureSkipAuthNRequestsSigned sets InsecureSkipAuthNRequestsSigned field to given value.

### HasInsecureSkipAuthNRequestsSigned

`func (o *AuthtypesAuthDomainConfig) HasInsecureSkipAuthNRequestsSigned() bool`

HasInsecureSkipAuthNRequestsSigned returns a boolean if a field has been set.

### GetSamlCert

`func (o *AuthtypesAuthDomainConfig) GetSamlCert() string`

GetSamlCert returns the SamlCert field if non-nil, zero value otherwise.

### GetSamlCertOk

`func (o *AuthtypesAuthDomainConfig) GetSamlCertOk() (*string, bool)`

GetSamlCertOk returns a tuple with the SamlCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCert

`func (o *AuthtypesAuthDomainConfig) SetSamlCert(v string)`

SetSamlCert sets SamlCert field to given value.

### HasSamlCert

`func (o *AuthtypesAuthDomainConfig) HasSamlCert() bool`

HasSamlCert returns a boolean if a field has been set.

### GetSamlEntity

`func (o *AuthtypesAuthDomainConfig) GetSamlEntity() string`

GetSamlEntity returns the SamlEntity field if non-nil, zero value otherwise.

### GetSamlEntityOk

`func (o *AuthtypesAuthDomainConfig) GetSamlEntityOk() (*string, bool)`

GetSamlEntityOk returns a tuple with the SamlEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlEntity

`func (o *AuthtypesAuthDomainConfig) SetSamlEntity(v string)`

SetSamlEntity sets SamlEntity field to given value.

### HasSamlEntity

`func (o *AuthtypesAuthDomainConfig) HasSamlEntity() bool`

HasSamlEntity returns a boolean if a field has been set.

### GetSamlIdp

`func (o *AuthtypesAuthDomainConfig) GetSamlIdp() string`

GetSamlIdp returns the SamlIdp field if non-nil, zero value otherwise.

### GetSamlIdpOk

`func (o *AuthtypesAuthDomainConfig) GetSamlIdpOk() (*string, bool)`

GetSamlIdpOk returns a tuple with the SamlIdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlIdp

`func (o *AuthtypesAuthDomainConfig) SetSamlIdp(v string)`

SetSamlIdp sets SamlIdp field to given value.

### HasSamlIdp

`func (o *AuthtypesAuthDomainConfig) HasSamlIdp() bool`

HasSamlIdp returns a boolean if a field has been set.

### GetAllowedGroups

`func (o *AuthtypesAuthDomainConfig) GetAllowedGroups() []string`

GetAllowedGroups returns the AllowedGroups field if non-nil, zero value otherwise.

### GetAllowedGroupsOk

`func (o *AuthtypesAuthDomainConfig) GetAllowedGroupsOk() (*[]string, bool)`

GetAllowedGroupsOk returns a tuple with the AllowedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGroups

`func (o *AuthtypesAuthDomainConfig) SetAllowedGroups(v []string)`

SetAllowedGroups sets AllowedGroups field to given value.

### HasAllowedGroups

`func (o *AuthtypesAuthDomainConfig) HasAllowedGroups() bool`

HasAllowedGroups returns a boolean if a field has been set.

### GetClientId

`func (o *AuthtypesAuthDomainConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AuthtypesAuthDomainConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AuthtypesAuthDomainConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AuthtypesAuthDomainConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *AuthtypesAuthDomainConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AuthtypesAuthDomainConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AuthtypesAuthDomainConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AuthtypesAuthDomainConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetDomainToAdminEmail

`func (o *AuthtypesAuthDomainConfig) GetDomainToAdminEmail() map[string]string`

GetDomainToAdminEmail returns the DomainToAdminEmail field if non-nil, zero value otherwise.

### GetDomainToAdminEmailOk

`func (o *AuthtypesAuthDomainConfig) GetDomainToAdminEmailOk() (*map[string]string, bool)`

GetDomainToAdminEmailOk returns a tuple with the DomainToAdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainToAdminEmail

`func (o *AuthtypesAuthDomainConfig) SetDomainToAdminEmail(v map[string]string)`

SetDomainToAdminEmail sets DomainToAdminEmail field to given value.

### HasDomainToAdminEmail

`func (o *AuthtypesAuthDomainConfig) HasDomainToAdminEmail() bool`

HasDomainToAdminEmail returns a boolean if a field has been set.

### GetFetchGroups

`func (o *AuthtypesAuthDomainConfig) GetFetchGroups() bool`

GetFetchGroups returns the FetchGroups field if non-nil, zero value otherwise.

### GetFetchGroupsOk

`func (o *AuthtypesAuthDomainConfig) GetFetchGroupsOk() (*bool, bool)`

GetFetchGroupsOk returns a tuple with the FetchGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchGroups

`func (o *AuthtypesAuthDomainConfig) SetFetchGroups(v bool)`

SetFetchGroups sets FetchGroups field to given value.

### HasFetchGroups

`func (o *AuthtypesAuthDomainConfig) HasFetchGroups() bool`

HasFetchGroups returns a boolean if a field has been set.

### GetFetchTransitiveGroupMembership

`func (o *AuthtypesAuthDomainConfig) GetFetchTransitiveGroupMembership() bool`

GetFetchTransitiveGroupMembership returns the FetchTransitiveGroupMembership field if non-nil, zero value otherwise.

### GetFetchTransitiveGroupMembershipOk

`func (o *AuthtypesAuthDomainConfig) GetFetchTransitiveGroupMembershipOk() (*bool, bool)`

GetFetchTransitiveGroupMembershipOk returns a tuple with the FetchTransitiveGroupMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchTransitiveGroupMembership

`func (o *AuthtypesAuthDomainConfig) SetFetchTransitiveGroupMembership(v bool)`

SetFetchTransitiveGroupMembership sets FetchTransitiveGroupMembership field to given value.

### HasFetchTransitiveGroupMembership

`func (o *AuthtypesAuthDomainConfig) HasFetchTransitiveGroupMembership() bool`

HasFetchTransitiveGroupMembership returns a boolean if a field has been set.

### GetInsecureSkipEmailVerified

`func (o *AuthtypesAuthDomainConfig) GetInsecureSkipEmailVerified() bool`

GetInsecureSkipEmailVerified returns the InsecureSkipEmailVerified field if non-nil, zero value otherwise.

### GetInsecureSkipEmailVerifiedOk

`func (o *AuthtypesAuthDomainConfig) GetInsecureSkipEmailVerifiedOk() (*bool, bool)`

GetInsecureSkipEmailVerifiedOk returns a tuple with the InsecureSkipEmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipEmailVerified

`func (o *AuthtypesAuthDomainConfig) SetInsecureSkipEmailVerified(v bool)`

SetInsecureSkipEmailVerified sets InsecureSkipEmailVerified field to given value.

### HasInsecureSkipEmailVerified

`func (o *AuthtypesAuthDomainConfig) HasInsecureSkipEmailVerified() bool`

HasInsecureSkipEmailVerified returns a boolean if a field has been set.

### GetRedirectURI

`func (o *AuthtypesAuthDomainConfig) GetRedirectURI() string`

GetRedirectURI returns the RedirectURI field if non-nil, zero value otherwise.

### GetRedirectURIOk

`func (o *AuthtypesAuthDomainConfig) GetRedirectURIOk() (*string, bool)`

GetRedirectURIOk returns a tuple with the RedirectURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectURI

`func (o *AuthtypesAuthDomainConfig) SetRedirectURI(v string)`

SetRedirectURI sets RedirectURI field to given value.

### HasRedirectURI

`func (o *AuthtypesAuthDomainConfig) HasRedirectURI() bool`

HasRedirectURI returns a boolean if a field has been set.

### GetServiceAccountJson

`func (o *AuthtypesAuthDomainConfig) GetServiceAccountJson() string`

GetServiceAccountJson returns the ServiceAccountJson field if non-nil, zero value otherwise.

### GetServiceAccountJsonOk

`func (o *AuthtypesAuthDomainConfig) GetServiceAccountJsonOk() (*string, bool)`

GetServiceAccountJsonOk returns a tuple with the ServiceAccountJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountJson

`func (o *AuthtypesAuthDomainConfig) SetServiceAccountJson(v string)`

SetServiceAccountJson sets ServiceAccountJson field to given value.

### HasServiceAccountJson

`func (o *AuthtypesAuthDomainConfig) HasServiceAccountJson() bool`

HasServiceAccountJson returns a boolean if a field has been set.

### GetClaimMapping

`func (o *AuthtypesAuthDomainConfig) GetClaimMapping() AuthtypesAttributeMapping`

GetClaimMapping returns the ClaimMapping field if non-nil, zero value otherwise.

### GetClaimMappingOk

`func (o *AuthtypesAuthDomainConfig) GetClaimMappingOk() (*AuthtypesAttributeMapping, bool)`

GetClaimMappingOk returns a tuple with the ClaimMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimMapping

`func (o *AuthtypesAuthDomainConfig) SetClaimMapping(v AuthtypesAttributeMapping)`

SetClaimMapping sets ClaimMapping field to given value.

### HasClaimMapping

`func (o *AuthtypesAuthDomainConfig) HasClaimMapping() bool`

HasClaimMapping returns a boolean if a field has been set.

### GetGetUserInfo

`func (o *AuthtypesAuthDomainConfig) GetGetUserInfo() bool`

GetGetUserInfo returns the GetUserInfo field if non-nil, zero value otherwise.

### GetGetUserInfoOk

`func (o *AuthtypesAuthDomainConfig) GetGetUserInfoOk() (*bool, bool)`

GetGetUserInfoOk returns a tuple with the GetUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetUserInfo

`func (o *AuthtypesAuthDomainConfig) SetGetUserInfo(v bool)`

SetGetUserInfo sets GetUserInfo field to given value.

### HasGetUserInfo

`func (o *AuthtypesAuthDomainConfig) HasGetUserInfo() bool`

HasGetUserInfo returns a boolean if a field has been set.

### GetIssuer

`func (o *AuthtypesAuthDomainConfig) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *AuthtypesAuthDomainConfig) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *AuthtypesAuthDomainConfig) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *AuthtypesAuthDomainConfig) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetIssuerAlias

`func (o *AuthtypesAuthDomainConfig) GetIssuerAlias() string`

GetIssuerAlias returns the IssuerAlias field if non-nil, zero value otherwise.

### GetIssuerAliasOk

`func (o *AuthtypesAuthDomainConfig) GetIssuerAliasOk() (*string, bool)`

GetIssuerAliasOk returns a tuple with the IssuerAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerAlias

`func (o *AuthtypesAuthDomainConfig) SetIssuerAlias(v string)`

SetIssuerAlias sets IssuerAlias field to given value.

### HasIssuerAlias

`func (o *AuthtypesAuthDomainConfig) HasIssuerAlias() bool`

HasIssuerAlias returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


