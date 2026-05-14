# AuthtypesOrgSessionContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthNSupport** | Pointer to [**AuthtypesAuthNSupport**](AuthtypesAuthNSupport.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Warning** | Pointer to [**ErrorsJSON**](ErrorsJSON.md) |  | [optional] 

## Methods

### NewAuthtypesOrgSessionContext

`func NewAuthtypesOrgSessionContext() *AuthtypesOrgSessionContext`

NewAuthtypesOrgSessionContext instantiates a new AuthtypesOrgSessionContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthtypesOrgSessionContextWithDefaults

`func NewAuthtypesOrgSessionContextWithDefaults() *AuthtypesOrgSessionContext`

NewAuthtypesOrgSessionContextWithDefaults instantiates a new AuthtypesOrgSessionContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthNSupport

`func (o *AuthtypesOrgSessionContext) GetAuthNSupport() AuthtypesAuthNSupport`

GetAuthNSupport returns the AuthNSupport field if non-nil, zero value otherwise.

### GetAuthNSupportOk

`func (o *AuthtypesOrgSessionContext) GetAuthNSupportOk() (*AuthtypesAuthNSupport, bool)`

GetAuthNSupportOk returns a tuple with the AuthNSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthNSupport

`func (o *AuthtypesOrgSessionContext) SetAuthNSupport(v AuthtypesAuthNSupport)`

SetAuthNSupport sets AuthNSupport field to given value.

### HasAuthNSupport

`func (o *AuthtypesOrgSessionContext) HasAuthNSupport() bool`

HasAuthNSupport returns a boolean if a field has been set.

### GetId

`func (o *AuthtypesOrgSessionContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthtypesOrgSessionContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthtypesOrgSessionContext) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthtypesOrgSessionContext) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthtypesOrgSessionContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthtypesOrgSessionContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthtypesOrgSessionContext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthtypesOrgSessionContext) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWarning

`func (o *AuthtypesOrgSessionContext) GetWarning() ErrorsJSON`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *AuthtypesOrgSessionContext) GetWarningOk() (*ErrorsJSON, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *AuthtypesOrgSessionContext) SetWarning(v ErrorsJSON)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *AuthtypesOrgSessionContext) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


