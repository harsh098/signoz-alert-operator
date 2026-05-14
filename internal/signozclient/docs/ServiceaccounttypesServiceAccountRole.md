# ServiceaccounttypesServiceAccountRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | **string** |  | 
**Role** | [**AuthtypesRole**](AuthtypesRole.md) |  | 
**RoleId** | **string** |  | 
**ServiceAccountId** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewServiceaccounttypesServiceAccountRole

`func NewServiceaccounttypesServiceAccountRole(id string, role AuthtypesRole, roleId string, serviceAccountId string, ) *ServiceaccounttypesServiceAccountRole`

NewServiceaccounttypesServiceAccountRole instantiates a new ServiceaccounttypesServiceAccountRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceaccounttypesServiceAccountRoleWithDefaults

`func NewServiceaccounttypesServiceAccountRoleWithDefaults() *ServiceaccounttypesServiceAccountRole`

NewServiceaccounttypesServiceAccountRoleWithDefaults instantiates a new ServiceaccounttypesServiceAccountRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceaccounttypesServiceAccountRole) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceaccounttypesServiceAccountRole) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceaccounttypesServiceAccountRole) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceaccounttypesServiceAccountRole) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *ServiceaccounttypesServiceAccountRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceaccounttypesServiceAccountRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceaccounttypesServiceAccountRole) SetId(v string)`

SetId sets Id field to given value.


### GetRole

`func (o *ServiceaccounttypesServiceAccountRole) GetRole() AuthtypesRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ServiceaccounttypesServiceAccountRole) GetRoleOk() (*AuthtypesRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ServiceaccounttypesServiceAccountRole) SetRole(v AuthtypesRole)`

SetRole sets Role field to given value.


### GetRoleId

`func (o *ServiceaccounttypesServiceAccountRole) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ServiceaccounttypesServiceAccountRole) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ServiceaccounttypesServiceAccountRole) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetServiceAccountId

`func (o *ServiceaccounttypesServiceAccountRole) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *ServiceaccounttypesServiceAccountRole) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *ServiceaccounttypesServiceAccountRole) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.


### GetUpdatedAt

`func (o *ServiceaccounttypesServiceAccountRole) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceaccounttypesServiceAccountRole) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceaccounttypesServiceAccountRole) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceaccounttypesServiceAccountRole) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


