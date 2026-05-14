# ConfigWebhookConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**MaxAlerts** | Pointer to **int32** |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UrlFile** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigWebhookConfig

`func NewConfigWebhookConfig() *ConfigWebhookConfig`

NewConfigWebhookConfig instantiates a new ConfigWebhookConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWebhookConfigWithDefaults

`func NewConfigWebhookConfigWithDefaults() *ConfigWebhookConfig`

NewConfigWebhookConfigWithDefaults instantiates a new ConfigWebhookConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpConfig

`func (o *ConfigWebhookConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigWebhookConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigWebhookConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigWebhookConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMaxAlerts

`func (o *ConfigWebhookConfig) GetMaxAlerts() int32`

GetMaxAlerts returns the MaxAlerts field if non-nil, zero value otherwise.

### GetMaxAlertsOk

`func (o *ConfigWebhookConfig) GetMaxAlertsOk() (*int32, bool)`

GetMaxAlertsOk returns a tuple with the MaxAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlerts

`func (o *ConfigWebhookConfig) SetMaxAlerts(v int32)`

SetMaxAlerts sets MaxAlerts field to given value.

### HasMaxAlerts

`func (o *ConfigWebhookConfig) HasMaxAlerts() bool`

HasMaxAlerts returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigWebhookConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigWebhookConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigWebhookConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigWebhookConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetTimeout

`func (o *ConfigWebhookConfig) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ConfigWebhookConfig) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ConfigWebhookConfig) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ConfigWebhookConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *ConfigWebhookConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConfigWebhookConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConfigWebhookConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConfigWebhookConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUrlFile

`func (o *ConfigWebhookConfig) GetUrlFile() string`

GetUrlFile returns the UrlFile field if non-nil, zero value otherwise.

### GetUrlFileOk

`func (o *ConfigWebhookConfig) GetUrlFileOk() (*string, bool)`

GetUrlFileOk returns a tuple with the UrlFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFile

`func (o *ConfigWebhookConfig) SetUrlFile(v string)`

SetUrlFile sets UrlFile field to given value.

### HasUrlFile

`func (o *ConfigWebhookConfig) HasUrlFile() bool`

HasUrlFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


