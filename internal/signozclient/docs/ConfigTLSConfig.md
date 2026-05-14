# ConfigTLSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ca** | Pointer to **string** |  | [optional] 
**CaFile** | Pointer to **string** |  | [optional] 
**CaRef** | Pointer to **string** |  | [optional] 
**Cert** | Pointer to **string** |  | [optional] 
**CertFile** | Pointer to **string** |  | [optional] 
**CertRef** | Pointer to **string** |  | [optional] 
**InsecureSkipVerify** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**KeyFile** | Pointer to **string** |  | [optional] 
**KeyRef** | Pointer to **string** |  | [optional] 
**MaxVersion** | Pointer to **int32** |  | [optional] 
**MinVersion** | Pointer to **int32** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigTLSConfig

`func NewConfigTLSConfig() *ConfigTLSConfig`

NewConfigTLSConfig instantiates a new ConfigTLSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigTLSConfigWithDefaults

`func NewConfigTLSConfigWithDefaults() *ConfigTLSConfig`

NewConfigTLSConfigWithDefaults instantiates a new ConfigTLSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCa

`func (o *ConfigTLSConfig) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *ConfigTLSConfig) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *ConfigTLSConfig) SetCa(v string)`

SetCa sets Ca field to given value.

### HasCa

`func (o *ConfigTLSConfig) HasCa() bool`

HasCa returns a boolean if a field has been set.

### GetCaFile

`func (o *ConfigTLSConfig) GetCaFile() string`

GetCaFile returns the CaFile field if non-nil, zero value otherwise.

### GetCaFileOk

`func (o *ConfigTLSConfig) GetCaFileOk() (*string, bool)`

GetCaFileOk returns a tuple with the CaFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaFile

`func (o *ConfigTLSConfig) SetCaFile(v string)`

SetCaFile sets CaFile field to given value.

### HasCaFile

`func (o *ConfigTLSConfig) HasCaFile() bool`

HasCaFile returns a boolean if a field has been set.

### GetCaRef

`func (o *ConfigTLSConfig) GetCaRef() string`

GetCaRef returns the CaRef field if non-nil, zero value otherwise.

### GetCaRefOk

`func (o *ConfigTLSConfig) GetCaRefOk() (*string, bool)`

GetCaRefOk returns a tuple with the CaRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaRef

`func (o *ConfigTLSConfig) SetCaRef(v string)`

SetCaRef sets CaRef field to given value.

### HasCaRef

`func (o *ConfigTLSConfig) HasCaRef() bool`

HasCaRef returns a boolean if a field has been set.

### GetCert

`func (o *ConfigTLSConfig) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *ConfigTLSConfig) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *ConfigTLSConfig) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *ConfigTLSConfig) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetCertFile

`func (o *ConfigTLSConfig) GetCertFile() string`

GetCertFile returns the CertFile field if non-nil, zero value otherwise.

### GetCertFileOk

`func (o *ConfigTLSConfig) GetCertFileOk() (*string, bool)`

GetCertFileOk returns a tuple with the CertFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertFile

`func (o *ConfigTLSConfig) SetCertFile(v string)`

SetCertFile sets CertFile field to given value.

### HasCertFile

`func (o *ConfigTLSConfig) HasCertFile() bool`

HasCertFile returns a boolean if a field has been set.

### GetCertRef

`func (o *ConfigTLSConfig) GetCertRef() string`

GetCertRef returns the CertRef field if non-nil, zero value otherwise.

### GetCertRefOk

`func (o *ConfigTLSConfig) GetCertRefOk() (*string, bool)`

GetCertRefOk returns a tuple with the CertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertRef

`func (o *ConfigTLSConfig) SetCertRef(v string)`

SetCertRef sets CertRef field to given value.

### HasCertRef

`func (o *ConfigTLSConfig) HasCertRef() bool`

HasCertRef returns a boolean if a field has been set.

### GetInsecureSkipVerify

`func (o *ConfigTLSConfig) GetInsecureSkipVerify() bool`

GetInsecureSkipVerify returns the InsecureSkipVerify field if non-nil, zero value otherwise.

### GetInsecureSkipVerifyOk

`func (o *ConfigTLSConfig) GetInsecureSkipVerifyOk() (*bool, bool)`

GetInsecureSkipVerifyOk returns a tuple with the InsecureSkipVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipVerify

`func (o *ConfigTLSConfig) SetInsecureSkipVerify(v bool)`

SetInsecureSkipVerify sets InsecureSkipVerify field to given value.

### HasInsecureSkipVerify

`func (o *ConfigTLSConfig) HasInsecureSkipVerify() bool`

HasInsecureSkipVerify returns a boolean if a field has been set.

### GetKey

`func (o *ConfigTLSConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ConfigTLSConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ConfigTLSConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ConfigTLSConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKeyFile

`func (o *ConfigTLSConfig) GetKeyFile() string`

GetKeyFile returns the KeyFile field if non-nil, zero value otherwise.

### GetKeyFileOk

`func (o *ConfigTLSConfig) GetKeyFileOk() (*string, bool)`

GetKeyFileOk returns a tuple with the KeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFile

`func (o *ConfigTLSConfig) SetKeyFile(v string)`

SetKeyFile sets KeyFile field to given value.

### HasKeyFile

`func (o *ConfigTLSConfig) HasKeyFile() bool`

HasKeyFile returns a boolean if a field has been set.

### GetKeyRef

`func (o *ConfigTLSConfig) GetKeyRef() string`

GetKeyRef returns the KeyRef field if non-nil, zero value otherwise.

### GetKeyRefOk

`func (o *ConfigTLSConfig) GetKeyRefOk() (*string, bool)`

GetKeyRefOk returns a tuple with the KeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRef

`func (o *ConfigTLSConfig) SetKeyRef(v string)`

SetKeyRef sets KeyRef field to given value.

### HasKeyRef

`func (o *ConfigTLSConfig) HasKeyRef() bool`

HasKeyRef returns a boolean if a field has been set.

### GetMaxVersion

`func (o *ConfigTLSConfig) GetMaxVersion() int32`

GetMaxVersion returns the MaxVersion field if non-nil, zero value otherwise.

### GetMaxVersionOk

`func (o *ConfigTLSConfig) GetMaxVersionOk() (*int32, bool)`

GetMaxVersionOk returns a tuple with the MaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersion

`func (o *ConfigTLSConfig) SetMaxVersion(v int32)`

SetMaxVersion sets MaxVersion field to given value.

### HasMaxVersion

`func (o *ConfigTLSConfig) HasMaxVersion() bool`

HasMaxVersion returns a boolean if a field has been set.

### GetMinVersion

`func (o *ConfigTLSConfig) GetMinVersion() int32`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *ConfigTLSConfig) GetMinVersionOk() (*int32, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *ConfigTLSConfig) SetMinVersion(v int32)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *ConfigTLSConfig) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetServerName

`func (o *ConfigTLSConfig) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ConfigTLSConfig) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ConfigTLSConfig) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ConfigTLSConfig) HasServerName() bool`

HasServerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


