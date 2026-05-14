# ServiceaccounttypesServiceAccountWithRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Email** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**OrgId** | **string** |  | 
**ServiceAccountRoles** | [**[]ServiceaccounttypesServiceAccountRole**](ServiceaccounttypesServiceAccountRole.md) |  | 
**Status** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewServiceaccounttypesServiceAccountWithRoles

`func NewServiceaccounttypesServiceAccountWithRoles(email string, id string, name string, orgId string, serviceAccountRoles []ServiceaccounttypesServiceAccountRole, status string, ) *ServiceaccounttypesServiceAccountWithRoles`

NewServiceaccounttypesServiceAccountWithRoles instantiates a new ServiceaccounttypesServiceAccountWithRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceaccounttypesServiceAccountWithRolesWithDefaults

`func NewServiceaccounttypesServiceAccountWithRolesWithDefaults() *ServiceaccounttypesServiceAccountWithRoles`

NewServiceaccounttypesServiceAccountWithRolesWithDefaults instantiates a new ServiceaccounttypesServiceAccountWithRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceaccounttypesServiceAccountWithRoles) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceaccounttypesServiceAccountWithRoles) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ServiceaccounttypesServiceAccountWithRoles) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetId

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceaccounttypesServiceAccountWithRoles) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceaccounttypesServiceAccountWithRoles) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ServiceaccounttypesServiceAccountWithRoles) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetServiceAccountRoles

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetServiceAccountRoles() []ServiceaccounttypesServiceAccountRole`

GetServiceAccountRoles returns the ServiceAccountRoles field if non-nil, zero value otherwise.

### GetServiceAccountRolesOk

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetServiceAccountRolesOk() (*[]ServiceaccounttypesServiceAccountRole, bool)`

GetServiceAccountRolesOk returns a tuple with the ServiceAccountRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountRoles

`func (o *ServiceaccounttypesServiceAccountWithRoles) SetServiceAccountRoles(v []ServiceaccounttypesServiceAccountRole)`

SetServiceAccountRoles sets ServiceAccountRoles field to given value.


### SetServiceAccountRolesNil

`func (o *ServiceaccounttypesServiceAccountWithRoles) SetServiceAccountRolesNil(b bool)`

 SetServiceAccountRolesNil sets the value for ServiceAccountRoles to be an explicit nil

### UnsetServiceAccountRoles
`func (o *ServiceaccounttypesServiceAccountWithRoles) UnsetServiceAccountRoles()`

UnsetServiceAccountRoles ensures that no value is present for ServiceAccountRoles, not even an explicit nil
### GetStatus

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceaccounttypesServiceAccountWithRoles) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceaccounttypesServiceAccountWithRoles) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceaccounttypesServiceAccountWithRoles) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceaccounttypesServiceAccountWithRoles) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


