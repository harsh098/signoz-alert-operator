# GlobaltypesConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiAssistantUrl** | **NullableString** |  | 
**ExternalUrl** | **string** |  | 
**IdentN** | Pointer to [**GlobaltypesIdentNConfig**](GlobaltypesIdentNConfig.md) |  | [optional] 
**IngestionUrl** | **string** |  | 
**McpUrl** | **NullableString** |  | 

## Methods

### NewGlobaltypesConfig

`func NewGlobaltypesConfig(aiAssistantUrl NullableString, externalUrl string, ingestionUrl string, mcpUrl NullableString, ) *GlobaltypesConfig`

NewGlobaltypesConfig instantiates a new GlobaltypesConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobaltypesConfigWithDefaults

`func NewGlobaltypesConfigWithDefaults() *GlobaltypesConfig`

NewGlobaltypesConfigWithDefaults instantiates a new GlobaltypesConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiAssistantUrl

`func (o *GlobaltypesConfig) GetAiAssistantUrl() string`

GetAiAssistantUrl returns the AiAssistantUrl field if non-nil, zero value otherwise.

### GetAiAssistantUrlOk

`func (o *GlobaltypesConfig) GetAiAssistantUrlOk() (*string, bool)`

GetAiAssistantUrlOk returns a tuple with the AiAssistantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiAssistantUrl

`func (o *GlobaltypesConfig) SetAiAssistantUrl(v string)`

SetAiAssistantUrl sets AiAssistantUrl field to given value.


### SetAiAssistantUrlNil

`func (o *GlobaltypesConfig) SetAiAssistantUrlNil(b bool)`

 SetAiAssistantUrlNil sets the value for AiAssistantUrl to be an explicit nil

### UnsetAiAssistantUrl
`func (o *GlobaltypesConfig) UnsetAiAssistantUrl()`

UnsetAiAssistantUrl ensures that no value is present for AiAssistantUrl, not even an explicit nil
### GetExternalUrl

`func (o *GlobaltypesConfig) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *GlobaltypesConfig) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *GlobaltypesConfig) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.


### GetIdentN

`func (o *GlobaltypesConfig) GetIdentN() GlobaltypesIdentNConfig`

GetIdentN returns the IdentN field if non-nil, zero value otherwise.

### GetIdentNOk

`func (o *GlobaltypesConfig) GetIdentNOk() (*GlobaltypesIdentNConfig, bool)`

GetIdentNOk returns a tuple with the IdentN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentN

`func (o *GlobaltypesConfig) SetIdentN(v GlobaltypesIdentNConfig)`

SetIdentN sets IdentN field to given value.

### HasIdentN

`func (o *GlobaltypesConfig) HasIdentN() bool`

HasIdentN returns a boolean if a field has been set.

### GetIngestionUrl

`func (o *GlobaltypesConfig) GetIngestionUrl() string`

GetIngestionUrl returns the IngestionUrl field if non-nil, zero value otherwise.

### GetIngestionUrlOk

`func (o *GlobaltypesConfig) GetIngestionUrlOk() (*string, bool)`

GetIngestionUrlOk returns a tuple with the IngestionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionUrl

`func (o *GlobaltypesConfig) SetIngestionUrl(v string)`

SetIngestionUrl sets IngestionUrl field to given value.


### GetMcpUrl

`func (o *GlobaltypesConfig) GetMcpUrl() string`

GetMcpUrl returns the McpUrl field if non-nil, zero value otherwise.

### GetMcpUrlOk

`func (o *GlobaltypesConfig) GetMcpUrlOk() (*string, bool)`

GetMcpUrlOk returns a tuple with the McpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpUrl

`func (o *GlobaltypesConfig) SetMcpUrl(v string)`

SetMcpUrl sets McpUrl field to given value.


### SetMcpUrlNil

`func (o *GlobaltypesConfig) SetMcpUrlNil(b bool)`

 SetMcpUrlNil sets the value for McpUrl to be an explicit nil

### UnsetMcpUrl
`func (o *GlobaltypesConfig) UnsetMcpUrl()`

UnsetMcpUrl ensures that no value is present for McpUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


