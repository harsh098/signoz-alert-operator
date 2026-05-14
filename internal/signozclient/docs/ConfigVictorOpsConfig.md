# ConfigVictorOpsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** |  | [optional] 
**ApiKeyFile** | Pointer to **string** |  | [optional] 
**ApiUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]string** |  | [optional] 
**EntityDisplayName** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**MessageType** | Pointer to **string** |  | [optional] 
**MonitoringTool** | Pointer to **string** |  | [optional] 
**RoutingKey** | Pointer to **string** |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**StateMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigVictorOpsConfig

`func NewConfigVictorOpsConfig() *ConfigVictorOpsConfig`

NewConfigVictorOpsConfig instantiates a new ConfigVictorOpsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigVictorOpsConfigWithDefaults

`func NewConfigVictorOpsConfigWithDefaults() *ConfigVictorOpsConfig`

NewConfigVictorOpsConfigWithDefaults instantiates a new ConfigVictorOpsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *ConfigVictorOpsConfig) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ConfigVictorOpsConfig) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ConfigVictorOpsConfig) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ConfigVictorOpsConfig) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApiKeyFile

`func (o *ConfigVictorOpsConfig) GetApiKeyFile() string`

GetApiKeyFile returns the ApiKeyFile field if non-nil, zero value otherwise.

### GetApiKeyFileOk

`func (o *ConfigVictorOpsConfig) GetApiKeyFileOk() (*string, bool)`

GetApiKeyFileOk returns a tuple with the ApiKeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyFile

`func (o *ConfigVictorOpsConfig) SetApiKeyFile(v string)`

SetApiKeyFile sets ApiKeyFile field to given value.

### HasApiKeyFile

`func (o *ConfigVictorOpsConfig) HasApiKeyFile() bool`

HasApiKeyFile returns a boolean if a field has been set.

### GetApiUrl

`func (o *ConfigVictorOpsConfig) GetApiUrl() map[string]interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ConfigVictorOpsConfig) GetApiUrlOk() (*map[string]interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ConfigVictorOpsConfig) SetApiUrl(v map[string]interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *ConfigVictorOpsConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetCustomFields

`func (o *ConfigVictorOpsConfig) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ConfigVictorOpsConfig) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ConfigVictorOpsConfig) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ConfigVictorOpsConfig) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEntityDisplayName

`func (o *ConfigVictorOpsConfig) GetEntityDisplayName() string`

GetEntityDisplayName returns the EntityDisplayName field if non-nil, zero value otherwise.

### GetEntityDisplayNameOk

`func (o *ConfigVictorOpsConfig) GetEntityDisplayNameOk() (*string, bool)`

GetEntityDisplayNameOk returns a tuple with the EntityDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityDisplayName

`func (o *ConfigVictorOpsConfig) SetEntityDisplayName(v string)`

SetEntityDisplayName sets EntityDisplayName field to given value.

### HasEntityDisplayName

`func (o *ConfigVictorOpsConfig) HasEntityDisplayName() bool`

HasEntityDisplayName returns a boolean if a field has been set.

### GetHttpConfig

`func (o *ConfigVictorOpsConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigVictorOpsConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigVictorOpsConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigVictorOpsConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessageType

`func (o *ConfigVictorOpsConfig) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ConfigVictorOpsConfig) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ConfigVictorOpsConfig) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ConfigVictorOpsConfig) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetMonitoringTool

`func (o *ConfigVictorOpsConfig) GetMonitoringTool() string`

GetMonitoringTool returns the MonitoringTool field if non-nil, zero value otherwise.

### GetMonitoringToolOk

`func (o *ConfigVictorOpsConfig) GetMonitoringToolOk() (*string, bool)`

GetMonitoringToolOk returns a tuple with the MonitoringTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringTool

`func (o *ConfigVictorOpsConfig) SetMonitoringTool(v string)`

SetMonitoringTool sets MonitoringTool field to given value.

### HasMonitoringTool

`func (o *ConfigVictorOpsConfig) HasMonitoringTool() bool`

HasMonitoringTool returns a boolean if a field has been set.

### GetRoutingKey

`func (o *ConfigVictorOpsConfig) GetRoutingKey() string`

GetRoutingKey returns the RoutingKey field if non-nil, zero value otherwise.

### GetRoutingKeyOk

`func (o *ConfigVictorOpsConfig) GetRoutingKeyOk() (*string, bool)`

GetRoutingKeyOk returns a tuple with the RoutingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKey

`func (o *ConfigVictorOpsConfig) SetRoutingKey(v string)`

SetRoutingKey sets RoutingKey field to given value.

### HasRoutingKey

`func (o *ConfigVictorOpsConfig) HasRoutingKey() bool`

HasRoutingKey returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigVictorOpsConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigVictorOpsConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigVictorOpsConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigVictorOpsConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetStateMessage

`func (o *ConfigVictorOpsConfig) GetStateMessage() string`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *ConfigVictorOpsConfig) GetStateMessageOk() (*string, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *ConfigVictorOpsConfig) SetStateMessage(v string)`

SetStateMessage sets StateMessage field to given value.

### HasStateMessage

`func (o *ConfigVictorOpsConfig) HasStateMessage() bool`

HasStateMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


