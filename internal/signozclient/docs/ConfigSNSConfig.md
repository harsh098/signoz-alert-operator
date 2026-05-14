# ConfigSNSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiUrl** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**Sigv4** | Pointer to **map[string]interface{}** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**TargetArn** | Pointer to **string** |  | [optional] 
**TopicArn** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigSNSConfig

`func NewConfigSNSConfig() *ConfigSNSConfig`

NewConfigSNSConfig instantiates a new ConfigSNSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSNSConfigWithDefaults

`func NewConfigSNSConfigWithDefaults() *ConfigSNSConfig`

NewConfigSNSConfigWithDefaults instantiates a new ConfigSNSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiUrl

`func (o *ConfigSNSConfig) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ConfigSNSConfig) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ConfigSNSConfig) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *ConfigSNSConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetAttributes

`func (o *ConfigSNSConfig) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConfigSNSConfig) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConfigSNSConfig) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ConfigSNSConfig) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetHttpConfig

`func (o *ConfigSNSConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigSNSConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigSNSConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigSNSConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessage

`func (o *ConfigSNSConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConfigSNSConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConfigSNSConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConfigSNSConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ConfigSNSConfig) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ConfigSNSConfig) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ConfigSNSConfig) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ConfigSNSConfig) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigSNSConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigSNSConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigSNSConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigSNSConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetSigv4

`func (o *ConfigSNSConfig) GetSigv4() map[string]interface{}`

GetSigv4 returns the Sigv4 field if non-nil, zero value otherwise.

### GetSigv4Ok

`func (o *ConfigSNSConfig) GetSigv4Ok() (*map[string]interface{}, bool)`

GetSigv4Ok returns a tuple with the Sigv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigv4

`func (o *ConfigSNSConfig) SetSigv4(v map[string]interface{})`

SetSigv4 sets Sigv4 field to given value.

### HasSigv4

`func (o *ConfigSNSConfig) HasSigv4() bool`

HasSigv4 returns a boolean if a field has been set.

### GetSubject

`func (o *ConfigSNSConfig) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ConfigSNSConfig) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ConfigSNSConfig) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ConfigSNSConfig) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTargetArn

`func (o *ConfigSNSConfig) GetTargetArn() string`

GetTargetArn returns the TargetArn field if non-nil, zero value otherwise.

### GetTargetArnOk

`func (o *ConfigSNSConfig) GetTargetArnOk() (*string, bool)`

GetTargetArnOk returns a tuple with the TargetArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetArn

`func (o *ConfigSNSConfig) SetTargetArn(v string)`

SetTargetArn sets TargetArn field to given value.

### HasTargetArn

`func (o *ConfigSNSConfig) HasTargetArn() bool`

HasTargetArn returns a boolean if a field has been set.

### GetTopicArn

`func (o *ConfigSNSConfig) GetTopicArn() string`

GetTopicArn returns the TopicArn field if non-nil, zero value otherwise.

### GetTopicArnOk

`func (o *ConfigSNSConfig) GetTopicArnOk() (*string, bool)`

GetTopicArnOk returns a tuple with the TopicArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicArn

`func (o *ConfigSNSConfig) SetTopicArn(v string)`

SetTopicArn sets TopicArn field to given value.

### HasTopicArn

`func (o *ConfigSNSConfig) HasTopicArn() bool`

HasTopicArn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


