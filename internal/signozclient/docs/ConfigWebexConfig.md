# ConfigWebexConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**RoomId** | Pointer to **string** |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 

## Methods

### NewConfigWebexConfig

`func NewConfigWebexConfig() *ConfigWebexConfig`

NewConfigWebexConfig instantiates a new ConfigWebexConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWebexConfigWithDefaults

`func NewConfigWebexConfigWithDefaults() *ConfigWebexConfig`

NewConfigWebexConfigWithDefaults instantiates a new ConfigWebexConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiUrl

`func (o *ConfigWebexConfig) GetApiUrl() map[string]interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ConfigWebexConfig) GetApiUrlOk() (*map[string]interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ConfigWebexConfig) SetApiUrl(v map[string]interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *ConfigWebexConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetHttpConfig

`func (o *ConfigWebexConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigWebexConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigWebexConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigWebexConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessage

`func (o *ConfigWebexConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConfigWebexConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConfigWebexConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConfigWebexConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRoomId

`func (o *ConfigWebexConfig) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *ConfigWebexConfig) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *ConfigWebexConfig) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *ConfigWebexConfig) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigWebexConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigWebexConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigWebexConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigWebexConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


