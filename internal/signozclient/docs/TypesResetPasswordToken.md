# TypesResetPasswordToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Id** | **string** |  | 
**PasswordId** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewTypesResetPasswordToken

`func NewTypesResetPasswordToken(id string, ) *TypesResetPasswordToken`

NewTypesResetPasswordToken instantiates a new TypesResetPasswordToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesResetPasswordTokenWithDefaults

`func NewTypesResetPasswordTokenWithDefaults() *TypesResetPasswordToken`

NewTypesResetPasswordTokenWithDefaults instantiates a new TypesResetPasswordToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *TypesResetPasswordToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TypesResetPasswordToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TypesResetPasswordToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TypesResetPasswordToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *TypesResetPasswordToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypesResetPasswordToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypesResetPasswordToken) SetId(v string)`

SetId sets Id field to given value.


### GetPasswordId

`func (o *TypesResetPasswordToken) GetPasswordId() string`

GetPasswordId returns the PasswordId field if non-nil, zero value otherwise.

### GetPasswordIdOk

`func (o *TypesResetPasswordToken) GetPasswordIdOk() (*string, bool)`

GetPasswordIdOk returns a tuple with the PasswordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordId

`func (o *TypesResetPasswordToken) SetPasswordId(v string)`

SetPasswordId sets PasswordId field to given value.

### HasPasswordId

`func (o *TypesResetPasswordToken) HasPasswordId() bool`

HasPasswordId returns a boolean if a field has been set.

### GetToken

`func (o *TypesResetPasswordToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TypesResetPasswordToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TypesResetPasswordToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TypesResetPasswordToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


