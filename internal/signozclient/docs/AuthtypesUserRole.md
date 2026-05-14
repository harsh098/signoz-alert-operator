# AuthtypesUserRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**Role** | [**AuthtypesRole**](AuthtypesRole.md) |  | 
**RoleId** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 

## Methods

### NewAuthtypesUserRole

`func NewAuthtypesUserRole(createdAt time.Time, id string, role AuthtypesRole, roleId string, updatedAt time.Time, userId string, ) *AuthtypesUserRole`

NewAuthtypesUserRole instantiates a new AuthtypesUserRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthtypesUserRoleWithDefaults

`func NewAuthtypesUserRoleWithDefaults() *AuthtypesUserRole`

NewAuthtypesUserRoleWithDefaults instantiates a new AuthtypesUserRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AuthtypesUserRole) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthtypesUserRole) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthtypesUserRole) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *AuthtypesUserRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthtypesUserRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthtypesUserRole) SetId(v string)`

SetId sets Id field to given value.


### GetRole

`func (o *AuthtypesUserRole) GetRole() AuthtypesRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AuthtypesUserRole) GetRoleOk() (*AuthtypesRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AuthtypesUserRole) SetRole(v AuthtypesRole)`

SetRole sets Role field to given value.


### GetRoleId

`func (o *AuthtypesUserRole) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *AuthtypesUserRole) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *AuthtypesUserRole) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetUpdatedAt

`func (o *AuthtypesUserRole) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthtypesUserRole) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthtypesUserRole) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *AuthtypesUserRole) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthtypesUserRole) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthtypesUserRole) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


