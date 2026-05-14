# AuthtypesSamlConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeMapping** | Pointer to [**AuthtypesAttributeMapping**](AuthtypesAttributeMapping.md) |  | [optional] 
**InsecureSkipAuthNRequestsSigned** | Pointer to **bool** |  | [optional] 
**SamlCert** | Pointer to **string** |  | [optional] 
**SamlEntity** | Pointer to **string** |  | [optional] 
**SamlIdp** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthtypesSamlConfig

`func NewAuthtypesSamlConfig() *AuthtypesSamlConfig`

NewAuthtypesSamlConfig instantiates a new AuthtypesSamlConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthtypesSamlConfigWithDefaults

`func NewAuthtypesSamlConfigWithDefaults() *AuthtypesSamlConfig`

NewAuthtypesSamlConfigWithDefaults instantiates a new AuthtypesSamlConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeMapping

`func (o *AuthtypesSamlConfig) GetAttributeMapping() AuthtypesAttributeMapping`

GetAttributeMapping returns the AttributeMapping field if non-nil, zero value otherwise.

### GetAttributeMappingOk

`func (o *AuthtypesSamlConfig) GetAttributeMappingOk() (*AuthtypesAttributeMapping, bool)`

GetAttributeMappingOk returns a tuple with the AttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMapping

`func (o *AuthtypesSamlConfig) SetAttributeMapping(v AuthtypesAttributeMapping)`

SetAttributeMapping sets AttributeMapping field to given value.

### HasAttributeMapping

`func (o *AuthtypesSamlConfig) HasAttributeMapping() bool`

HasAttributeMapping returns a boolean if a field has been set.

### GetInsecureSkipAuthNRequestsSigned

`func (o *AuthtypesSamlConfig) GetInsecureSkipAuthNRequestsSigned() bool`

GetInsecureSkipAuthNRequestsSigned returns the InsecureSkipAuthNRequestsSigned field if non-nil, zero value otherwise.

### GetInsecureSkipAuthNRequestsSignedOk

`func (o *AuthtypesSamlConfig) GetInsecureSkipAuthNRequestsSignedOk() (*bool, bool)`

GetInsecureSkipAuthNRequestsSignedOk returns a tuple with the InsecureSkipAuthNRequestsSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipAuthNRequestsSigned

`func (o *AuthtypesSamlConfig) SetInsecureSkipAuthNRequestsSigned(v bool)`

SetInsecureSkipAuthNRequestsSigned sets InsecureSkipAuthNRequestsSigned field to given value.

### HasInsecureSkipAuthNRequestsSigned

`func (o *AuthtypesSamlConfig) HasInsecureSkipAuthNRequestsSigned() bool`

HasInsecureSkipAuthNRequestsSigned returns a boolean if a field has been set.

### GetSamlCert

`func (o *AuthtypesSamlConfig) GetSamlCert() string`

GetSamlCert returns the SamlCert field if non-nil, zero value otherwise.

### GetSamlCertOk

`func (o *AuthtypesSamlConfig) GetSamlCertOk() (*string, bool)`

GetSamlCertOk returns a tuple with the SamlCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCert

`func (o *AuthtypesSamlConfig) SetSamlCert(v string)`

SetSamlCert sets SamlCert field to given value.

### HasSamlCert

`func (o *AuthtypesSamlConfig) HasSamlCert() bool`

HasSamlCert returns a boolean if a field has been set.

### GetSamlEntity

`func (o *AuthtypesSamlConfig) GetSamlEntity() string`

GetSamlEntity returns the SamlEntity field if non-nil, zero value otherwise.

### GetSamlEntityOk

`func (o *AuthtypesSamlConfig) GetSamlEntityOk() (*string, bool)`

GetSamlEntityOk returns a tuple with the SamlEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlEntity

`func (o *AuthtypesSamlConfig) SetSamlEntity(v string)`

SetSamlEntity sets SamlEntity field to given value.

### HasSamlEntity

`func (o *AuthtypesSamlConfig) HasSamlEntity() bool`

HasSamlEntity returns a boolean if a field has been set.

### GetSamlIdp

`func (o *AuthtypesSamlConfig) GetSamlIdp() string`

GetSamlIdp returns the SamlIdp field if non-nil, zero value otherwise.

### GetSamlIdpOk

`func (o *AuthtypesSamlConfig) GetSamlIdpOk() (*string, bool)`

GetSamlIdpOk returns a tuple with the SamlIdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlIdp

`func (o *AuthtypesSamlConfig) SetSamlIdp(v string)`

SetSamlIdp sets SamlIdp field to given value.

### HasSamlIdp

`func (o *AuthtypesSamlConfig) HasSamlIdp() bool`

HasSamlIdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


