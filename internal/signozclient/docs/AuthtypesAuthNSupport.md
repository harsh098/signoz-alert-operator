# AuthtypesAuthNSupport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Callback** | Pointer to [**[]AuthtypesCallbackAuthNSupport**](AuthtypesCallbackAuthNSupport.md) |  | [optional] 
**Password** | Pointer to [**[]AuthtypesPasswordAuthNSupport**](AuthtypesPasswordAuthNSupport.md) |  | [optional] 

## Methods

### NewAuthtypesAuthNSupport

`func NewAuthtypesAuthNSupport() *AuthtypesAuthNSupport`

NewAuthtypesAuthNSupport instantiates a new AuthtypesAuthNSupport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthtypesAuthNSupportWithDefaults

`func NewAuthtypesAuthNSupportWithDefaults() *AuthtypesAuthNSupport`

NewAuthtypesAuthNSupportWithDefaults instantiates a new AuthtypesAuthNSupport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallback

`func (o *AuthtypesAuthNSupport) GetCallback() []AuthtypesCallbackAuthNSupport`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *AuthtypesAuthNSupport) GetCallbackOk() (*[]AuthtypesCallbackAuthNSupport, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *AuthtypesAuthNSupport) SetCallback(v []AuthtypesCallbackAuthNSupport)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *AuthtypesAuthNSupport) HasCallback() bool`

HasCallback returns a boolean if a field has been set.

### SetCallbackNil

`func (o *AuthtypesAuthNSupport) SetCallbackNil(b bool)`

 SetCallbackNil sets the value for Callback to be an explicit nil

### UnsetCallback
`func (o *AuthtypesAuthNSupport) UnsetCallback()`

UnsetCallback ensures that no value is present for Callback, not even an explicit nil
### GetPassword

`func (o *AuthtypesAuthNSupport) GetPassword() []AuthtypesPasswordAuthNSupport`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthtypesAuthNSupport) GetPasswordOk() (*[]AuthtypesPasswordAuthNSupport, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthtypesAuthNSupport) SetPassword(v []AuthtypesPasswordAuthNSupport)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AuthtypesAuthNSupport) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AuthtypesAuthNSupport) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AuthtypesAuthNSupport) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


