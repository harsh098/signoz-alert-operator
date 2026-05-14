# AuthtypesUserWithRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**IsRoot** | Pointer to **bool** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UserRoles** | Pointer to [**[]AuthtypesUserRole**](AuthtypesUserRole.md) |  | [optional] 

## Methods

### NewAuthtypesUserWithRoles

`func NewAuthtypesUserWithRoles(id string, ) *AuthtypesUserWithRoles`

NewAuthtypesUserWithRoles instantiates a new AuthtypesUserWithRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthtypesUserWithRolesWithDefaults

`func NewAuthtypesUserWithRolesWithDefaults() *AuthtypesUserWithRoles`

NewAuthtypesUserWithRolesWithDefaults instantiates a new AuthtypesUserWithRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AuthtypesUserWithRoles) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthtypesUserWithRoles) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthtypesUserWithRoles) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthtypesUserWithRoles) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *AuthtypesUserWithRoles) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AuthtypesUserWithRoles) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AuthtypesUserWithRoles) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AuthtypesUserWithRoles) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *AuthtypesUserWithRoles) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthtypesUserWithRoles) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthtypesUserWithRoles) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthtypesUserWithRoles) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *AuthtypesUserWithRoles) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthtypesUserWithRoles) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthtypesUserWithRoles) SetId(v string)`

SetId sets Id field to given value.


### GetIsRoot

`func (o *AuthtypesUserWithRoles) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *AuthtypesUserWithRoles) GetIsRootOk() (*bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoot

`func (o *AuthtypesUserWithRoles) SetIsRoot(v bool)`

SetIsRoot sets IsRoot field to given value.

### HasIsRoot

`func (o *AuthtypesUserWithRoles) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.

### GetOrgId

`func (o *AuthtypesUserWithRoles) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AuthtypesUserWithRoles) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AuthtypesUserWithRoles) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AuthtypesUserWithRoles) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetStatus

`func (o *AuthtypesUserWithRoles) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthtypesUserWithRoles) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthtypesUserWithRoles) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthtypesUserWithRoles) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthtypesUserWithRoles) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthtypesUserWithRoles) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthtypesUserWithRoles) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthtypesUserWithRoles) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserRoles

`func (o *AuthtypesUserWithRoles) GetUserRoles() []AuthtypesUserRole`

GetUserRoles returns the UserRoles field if non-nil, zero value otherwise.

### GetUserRolesOk

`func (o *AuthtypesUserWithRoles) GetUserRolesOk() (*[]AuthtypesUserRole, bool)`

GetUserRolesOk returns a tuple with the UserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRoles

`func (o *AuthtypesUserWithRoles) SetUserRoles(v []AuthtypesUserRole)`

SetUserRoles sets UserRoles field to given value.

### HasUserRoles

`func (o *AuthtypesUserWithRoles) HasUserRoles() bool`

HasUserRoles returns a boolean if a field has been set.

### SetUserRolesNil

`func (o *AuthtypesUserWithRoles) SetUserRolesNil(b bool)`

 SetUserRolesNil sets the value for UserRoles to be an explicit nil

### UnsetUserRoles
`func (o *AuthtypesUserWithRoles) UnsetUserRoles()`

UnsetUserRoles ensures that no value is present for UserRoles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


