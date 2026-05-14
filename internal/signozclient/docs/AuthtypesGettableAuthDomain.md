# AuthtypesGettableAuthDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthNProviderInfo** | Pointer to [**AuthtypesAuthNProviderInfo**](AuthtypesAuthNProviderInfo.md) |  | [optional] 
**Config** | Pointer to [**AuthtypesAuthDomainConfig**](AuthtypesAuthDomainConfig.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAuthtypesGettableAuthDomain

`func NewAuthtypesGettableAuthDomain(id string, ) *AuthtypesGettableAuthDomain`

NewAuthtypesGettableAuthDomain instantiates a new AuthtypesGettableAuthDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthtypesGettableAuthDomainWithDefaults

`func NewAuthtypesGettableAuthDomainWithDefaults() *AuthtypesGettableAuthDomain`

NewAuthtypesGettableAuthDomainWithDefaults instantiates a new AuthtypesGettableAuthDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthNProviderInfo

`func (o *AuthtypesGettableAuthDomain) GetAuthNProviderInfo() AuthtypesAuthNProviderInfo`

GetAuthNProviderInfo returns the AuthNProviderInfo field if non-nil, zero value otherwise.

### GetAuthNProviderInfoOk

`func (o *AuthtypesGettableAuthDomain) GetAuthNProviderInfoOk() (*AuthtypesAuthNProviderInfo, bool)`

GetAuthNProviderInfoOk returns a tuple with the AuthNProviderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthNProviderInfo

`func (o *AuthtypesGettableAuthDomain) SetAuthNProviderInfo(v AuthtypesAuthNProviderInfo)`

SetAuthNProviderInfo sets AuthNProviderInfo field to given value.

### HasAuthNProviderInfo

`func (o *AuthtypesGettableAuthDomain) HasAuthNProviderInfo() bool`

HasAuthNProviderInfo returns a boolean if a field has been set.

### GetConfig

`func (o *AuthtypesGettableAuthDomain) GetConfig() AuthtypesAuthDomainConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AuthtypesGettableAuthDomain) GetConfigOk() (*AuthtypesAuthDomainConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AuthtypesGettableAuthDomain) SetConfig(v AuthtypesAuthDomainConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AuthtypesGettableAuthDomain) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthtypesGettableAuthDomain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthtypesGettableAuthDomain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthtypesGettableAuthDomain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthtypesGettableAuthDomain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *AuthtypesGettableAuthDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthtypesGettableAuthDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthtypesGettableAuthDomain) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AuthtypesGettableAuthDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthtypesGettableAuthDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthtypesGettableAuthDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthtypesGettableAuthDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *AuthtypesGettableAuthDomain) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AuthtypesGettableAuthDomain) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AuthtypesGettableAuthDomain) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AuthtypesGettableAuthDomain) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthtypesGettableAuthDomain) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthtypesGettableAuthDomain) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthtypesGettableAuthDomain) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthtypesGettableAuthDomain) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


