# ConfigIncidentioConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertSourceToken** | Pointer to **string** |  | [optional] 
**AlertSourceTokenFile** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**MaxAlerts** | Pointer to **int32** |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 
**Url** | Pointer to **map[string]interface{}** |  | [optional] 
**UrlFile** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigIncidentioConfig

`func NewConfigIncidentioConfig() *ConfigIncidentioConfig`

NewConfigIncidentioConfig instantiates a new ConfigIncidentioConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigIncidentioConfigWithDefaults

`func NewConfigIncidentioConfigWithDefaults() *ConfigIncidentioConfig`

NewConfigIncidentioConfigWithDefaults instantiates a new ConfigIncidentioConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertSourceToken

`func (o *ConfigIncidentioConfig) GetAlertSourceToken() string`

GetAlertSourceToken returns the AlertSourceToken field if non-nil, zero value otherwise.

### GetAlertSourceTokenOk

`func (o *ConfigIncidentioConfig) GetAlertSourceTokenOk() (*string, bool)`

GetAlertSourceTokenOk returns a tuple with the AlertSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSourceToken

`func (o *ConfigIncidentioConfig) SetAlertSourceToken(v string)`

SetAlertSourceToken sets AlertSourceToken field to given value.

### HasAlertSourceToken

`func (o *ConfigIncidentioConfig) HasAlertSourceToken() bool`

HasAlertSourceToken returns a boolean if a field has been set.

### GetAlertSourceTokenFile

`func (o *ConfigIncidentioConfig) GetAlertSourceTokenFile() string`

GetAlertSourceTokenFile returns the AlertSourceTokenFile field if non-nil, zero value otherwise.

### GetAlertSourceTokenFileOk

`func (o *ConfigIncidentioConfig) GetAlertSourceTokenFileOk() (*string, bool)`

GetAlertSourceTokenFileOk returns a tuple with the AlertSourceTokenFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSourceTokenFile

`func (o *ConfigIncidentioConfig) SetAlertSourceTokenFile(v string)`

SetAlertSourceTokenFile sets AlertSourceTokenFile field to given value.

### HasAlertSourceTokenFile

`func (o *ConfigIncidentioConfig) HasAlertSourceTokenFile() bool`

HasAlertSourceTokenFile returns a boolean if a field has been set.

### GetHttpConfig

`func (o *ConfigIncidentioConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigIncidentioConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigIncidentioConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigIncidentioConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMaxAlerts

`func (o *ConfigIncidentioConfig) GetMaxAlerts() int32`

GetMaxAlerts returns the MaxAlerts field if non-nil, zero value otherwise.

### GetMaxAlertsOk

`func (o *ConfigIncidentioConfig) GetMaxAlertsOk() (*int32, bool)`

GetMaxAlertsOk returns a tuple with the MaxAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlerts

`func (o *ConfigIncidentioConfig) SetMaxAlerts(v int32)`

SetMaxAlerts sets MaxAlerts field to given value.

### HasMaxAlerts

`func (o *ConfigIncidentioConfig) HasMaxAlerts() bool`

HasMaxAlerts returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigIncidentioConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigIncidentioConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigIncidentioConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigIncidentioConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetTimeout

`func (o *ConfigIncidentioConfig) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ConfigIncidentioConfig) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ConfigIncidentioConfig) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ConfigIncidentioConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *ConfigIncidentioConfig) GetUrl() map[string]interface{}`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConfigIncidentioConfig) GetUrlOk() (*map[string]interface{}, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConfigIncidentioConfig) SetUrl(v map[string]interface{})`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConfigIncidentioConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUrlFile

`func (o *ConfigIncidentioConfig) GetUrlFile() string`

GetUrlFile returns the UrlFile field if non-nil, zero value otherwise.

### GetUrlFileOk

`func (o *ConfigIncidentioConfig) GetUrlFileOk() (*string, bool)`

GetUrlFileOk returns a tuple with the UrlFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFile

`func (o *ConfigIncidentioConfig) SetUrlFile(v string)`

SetUrlFile sets UrlFile field to given value.

### HasUrlFile

`func (o *ConfigIncidentioConfig) HasUrlFile() bool`

HasUrlFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


