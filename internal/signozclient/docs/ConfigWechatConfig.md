# ConfigWechatConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional] 
**ApiSecret** | Pointer to **string** |  | [optional] 
**ApiSecretFile** | Pointer to **string** |  | [optional] 
**ApiUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**CorpId** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**MessageType** | Pointer to **string** |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**ToParty** | Pointer to **string** |  | [optional] 
**ToTag** | Pointer to **string** |  | [optional] 
**ToUser** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigWechatConfig

`func NewConfigWechatConfig() *ConfigWechatConfig`

NewConfigWechatConfig instantiates a new ConfigWechatConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWechatConfigWithDefaults

`func NewConfigWechatConfigWithDefaults() *ConfigWechatConfig`

NewConfigWechatConfigWithDefaults instantiates a new ConfigWechatConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *ConfigWechatConfig) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ConfigWechatConfig) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ConfigWechatConfig) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ConfigWechatConfig) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetApiSecret

`func (o *ConfigWechatConfig) GetApiSecret() string`

GetApiSecret returns the ApiSecret field if non-nil, zero value otherwise.

### GetApiSecretOk

`func (o *ConfigWechatConfig) GetApiSecretOk() (*string, bool)`

GetApiSecretOk returns a tuple with the ApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSecret

`func (o *ConfigWechatConfig) SetApiSecret(v string)`

SetApiSecret sets ApiSecret field to given value.

### HasApiSecret

`func (o *ConfigWechatConfig) HasApiSecret() bool`

HasApiSecret returns a boolean if a field has been set.

### GetApiSecretFile

`func (o *ConfigWechatConfig) GetApiSecretFile() string`

GetApiSecretFile returns the ApiSecretFile field if non-nil, zero value otherwise.

### GetApiSecretFileOk

`func (o *ConfigWechatConfig) GetApiSecretFileOk() (*string, bool)`

GetApiSecretFileOk returns a tuple with the ApiSecretFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSecretFile

`func (o *ConfigWechatConfig) SetApiSecretFile(v string)`

SetApiSecretFile sets ApiSecretFile field to given value.

### HasApiSecretFile

`func (o *ConfigWechatConfig) HasApiSecretFile() bool`

HasApiSecretFile returns a boolean if a field has been set.

### GetApiUrl

`func (o *ConfigWechatConfig) GetApiUrl() map[string]interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ConfigWechatConfig) GetApiUrlOk() (*map[string]interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ConfigWechatConfig) SetApiUrl(v map[string]interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *ConfigWechatConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetCorpId

`func (o *ConfigWechatConfig) GetCorpId() string`

GetCorpId returns the CorpId field if non-nil, zero value otherwise.

### GetCorpIdOk

`func (o *ConfigWechatConfig) GetCorpIdOk() (*string, bool)`

GetCorpIdOk returns a tuple with the CorpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorpId

`func (o *ConfigWechatConfig) SetCorpId(v string)`

SetCorpId sets CorpId field to given value.

### HasCorpId

`func (o *ConfigWechatConfig) HasCorpId() bool`

HasCorpId returns a boolean if a field has been set.

### GetHttpConfig

`func (o *ConfigWechatConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigWechatConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigWechatConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigWechatConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessage

`func (o *ConfigWechatConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConfigWechatConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConfigWechatConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConfigWechatConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageType

`func (o *ConfigWechatConfig) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ConfigWechatConfig) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ConfigWechatConfig) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ConfigWechatConfig) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigWechatConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigWechatConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigWechatConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigWechatConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetToParty

`func (o *ConfigWechatConfig) GetToParty() string`

GetToParty returns the ToParty field if non-nil, zero value otherwise.

### GetToPartyOk

`func (o *ConfigWechatConfig) GetToPartyOk() (*string, bool)`

GetToPartyOk returns a tuple with the ToParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToParty

`func (o *ConfigWechatConfig) SetToParty(v string)`

SetToParty sets ToParty field to given value.

### HasToParty

`func (o *ConfigWechatConfig) HasToParty() bool`

HasToParty returns a boolean if a field has been set.

### GetToTag

`func (o *ConfigWechatConfig) GetToTag() string`

GetToTag returns the ToTag field if non-nil, zero value otherwise.

### GetToTagOk

`func (o *ConfigWechatConfig) GetToTagOk() (*string, bool)`

GetToTagOk returns a tuple with the ToTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTag

`func (o *ConfigWechatConfig) SetToTag(v string)`

SetToTag sets ToTag field to given value.

### HasToTag

`func (o *ConfigWechatConfig) HasToTag() bool`

HasToTag returns a boolean if a field has been set.

### GetToUser

`func (o *ConfigWechatConfig) GetToUser() string`

GetToUser returns the ToUser field if non-nil, zero value otherwise.

### GetToUserOk

`func (o *ConfigWechatConfig) GetToUserOk() (*string, bool)`

GetToUserOk returns a tuple with the ToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUser

`func (o *ConfigWechatConfig) SetToUser(v string)`

SetToUser sets ToUser field to given value.

### HasToUser

`func (o *ConfigWechatConfig) HasToUser() bool`

HasToUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


