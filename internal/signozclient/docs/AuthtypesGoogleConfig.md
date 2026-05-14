# AuthtypesGoogleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedGroups** | Pointer to **[]string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**DomainToAdminEmail** | Pointer to **map[string]string** |  | [optional] 
**FetchGroups** | Pointer to **bool** |  | [optional] 
**FetchTransitiveGroupMembership** | Pointer to **bool** |  | [optional] 
**InsecureSkipEmailVerified** | Pointer to **bool** |  | [optional] 
**RedirectURI** | Pointer to **string** |  | [optional] 
**ServiceAccountJson** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthtypesGoogleConfig

`func NewAuthtypesGoogleConfig() *AuthtypesGoogleConfig`

NewAuthtypesGoogleConfig instantiates a new AuthtypesGoogleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthtypesGoogleConfigWithDefaults

`func NewAuthtypesGoogleConfigWithDefaults() *AuthtypesGoogleConfig`

NewAuthtypesGoogleConfigWithDefaults instantiates a new AuthtypesGoogleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedGroups

`func (o *AuthtypesGoogleConfig) GetAllowedGroups() []string`

GetAllowedGroups returns the AllowedGroups field if non-nil, zero value otherwise.

### GetAllowedGroupsOk

`func (o *AuthtypesGoogleConfig) GetAllowedGroupsOk() (*[]string, bool)`

GetAllowedGroupsOk returns a tuple with the AllowedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGroups

`func (o *AuthtypesGoogleConfig) SetAllowedGroups(v []string)`

SetAllowedGroups sets AllowedGroups field to given value.

### HasAllowedGroups

`func (o *AuthtypesGoogleConfig) HasAllowedGroups() bool`

HasAllowedGroups returns a boolean if a field has been set.

### GetClientId

`func (o *AuthtypesGoogleConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AuthtypesGoogleConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AuthtypesGoogleConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AuthtypesGoogleConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *AuthtypesGoogleConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AuthtypesGoogleConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AuthtypesGoogleConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AuthtypesGoogleConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetDomainToAdminEmail

`func (o *AuthtypesGoogleConfig) GetDomainToAdminEmail() map[string]string`

GetDomainToAdminEmail returns the DomainToAdminEmail field if non-nil, zero value otherwise.

### GetDomainToAdminEmailOk

`func (o *AuthtypesGoogleConfig) GetDomainToAdminEmailOk() (*map[string]string, bool)`

GetDomainToAdminEmailOk returns a tuple with the DomainToAdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainToAdminEmail

`func (o *AuthtypesGoogleConfig) SetDomainToAdminEmail(v map[string]string)`

SetDomainToAdminEmail sets DomainToAdminEmail field to given value.

### HasDomainToAdminEmail

`func (o *AuthtypesGoogleConfig) HasDomainToAdminEmail() bool`

HasDomainToAdminEmail returns a boolean if a field has been set.

### GetFetchGroups

`func (o *AuthtypesGoogleConfig) GetFetchGroups() bool`

GetFetchGroups returns the FetchGroups field if non-nil, zero value otherwise.

### GetFetchGroupsOk

`func (o *AuthtypesGoogleConfig) GetFetchGroupsOk() (*bool, bool)`

GetFetchGroupsOk returns a tuple with the FetchGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchGroups

`func (o *AuthtypesGoogleConfig) SetFetchGroups(v bool)`

SetFetchGroups sets FetchGroups field to given value.

### HasFetchGroups

`func (o *AuthtypesGoogleConfig) HasFetchGroups() bool`

HasFetchGroups returns a boolean if a field has been set.

### GetFetchTransitiveGroupMembership

`func (o *AuthtypesGoogleConfig) GetFetchTransitiveGroupMembership() bool`

GetFetchTransitiveGroupMembership returns the FetchTransitiveGroupMembership field if non-nil, zero value otherwise.

### GetFetchTransitiveGroupMembershipOk

`func (o *AuthtypesGoogleConfig) GetFetchTransitiveGroupMembershipOk() (*bool, bool)`

GetFetchTransitiveGroupMembershipOk returns a tuple with the FetchTransitiveGroupMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchTransitiveGroupMembership

`func (o *AuthtypesGoogleConfig) SetFetchTransitiveGroupMembership(v bool)`

SetFetchTransitiveGroupMembership sets FetchTransitiveGroupMembership field to given value.

### HasFetchTransitiveGroupMembership

`func (o *AuthtypesGoogleConfig) HasFetchTransitiveGroupMembership() bool`

HasFetchTransitiveGroupMembership returns a boolean if a field has been set.

### GetInsecureSkipEmailVerified

`func (o *AuthtypesGoogleConfig) GetInsecureSkipEmailVerified() bool`

GetInsecureSkipEmailVerified returns the InsecureSkipEmailVerified field if non-nil, zero value otherwise.

### GetInsecureSkipEmailVerifiedOk

`func (o *AuthtypesGoogleConfig) GetInsecureSkipEmailVerifiedOk() (*bool, bool)`

GetInsecureSkipEmailVerifiedOk returns a tuple with the InsecureSkipEmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipEmailVerified

`func (o *AuthtypesGoogleConfig) SetInsecureSkipEmailVerified(v bool)`

SetInsecureSkipEmailVerified sets InsecureSkipEmailVerified field to given value.

### HasInsecureSkipEmailVerified

`func (o *AuthtypesGoogleConfig) HasInsecureSkipEmailVerified() bool`

HasInsecureSkipEmailVerified returns a boolean if a field has been set.

### GetRedirectURI

`func (o *AuthtypesGoogleConfig) GetRedirectURI() string`

GetRedirectURI returns the RedirectURI field if non-nil, zero value otherwise.

### GetRedirectURIOk

`func (o *AuthtypesGoogleConfig) GetRedirectURIOk() (*string, bool)`

GetRedirectURIOk returns a tuple with the RedirectURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectURI

`func (o *AuthtypesGoogleConfig) SetRedirectURI(v string)`

SetRedirectURI sets RedirectURI field to given value.

### HasRedirectURI

`func (o *AuthtypesGoogleConfig) HasRedirectURI() bool`

HasRedirectURI returns a boolean if a field has been set.

### GetServiceAccountJson

`func (o *AuthtypesGoogleConfig) GetServiceAccountJson() string`

GetServiceAccountJson returns the ServiceAccountJson field if non-nil, zero value otherwise.

### GetServiceAccountJsonOk

`func (o *AuthtypesGoogleConfig) GetServiceAccountJsonOk() (*string, bool)`

GetServiceAccountJsonOk returns a tuple with the ServiceAccountJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountJson

`func (o *AuthtypesGoogleConfig) SetServiceAccountJson(v string)`

SetServiceAccountJson sets ServiceAccountJson field to given value.

### HasServiceAccountJson

`func (o *AuthtypesGoogleConfig) HasServiceAccountJson() bool`

HasServiceAccountJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


