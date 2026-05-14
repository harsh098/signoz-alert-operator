# ConfigSlackConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]ConfigSlackAction**](ConfigSlackAction.md) |  | [optional] 
**ApiUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**ApiUrlFile** | Pointer to **string** |  | [optional] 
**AppToken** | Pointer to **string** |  | [optional] 
**AppTokenFile** | Pointer to **string** |  | [optional] 
**AppUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**CallbackId** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Fallback** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**[]ConfigSlackField**](ConfigSlackField.md) |  | [optional] 
**Footer** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**IconEmoji** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**LinkNames** | Pointer to **bool** |  | [optional] 
**MessageText** | Pointer to **string** |  | [optional] 
**MrkdwnIn** | Pointer to **[]string** |  | [optional] 
**Pretext** | Pointer to **string** |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**ShortFields** | Pointer to **bool** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**ThumbUrl** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TitleLink** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigSlackConfig

`func NewConfigSlackConfig() *ConfigSlackConfig`

NewConfigSlackConfig instantiates a new ConfigSlackConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSlackConfigWithDefaults

`func NewConfigSlackConfigWithDefaults() *ConfigSlackConfig`

NewConfigSlackConfigWithDefaults instantiates a new ConfigSlackConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ConfigSlackConfig) GetActions() []ConfigSlackAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ConfigSlackConfig) GetActionsOk() (*[]ConfigSlackAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ConfigSlackConfig) SetActions(v []ConfigSlackAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ConfigSlackConfig) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetApiUrl

`func (o *ConfigSlackConfig) GetApiUrl() map[string]interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ConfigSlackConfig) GetApiUrlOk() (*map[string]interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ConfigSlackConfig) SetApiUrl(v map[string]interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *ConfigSlackConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetApiUrlFile

`func (o *ConfigSlackConfig) GetApiUrlFile() string`

GetApiUrlFile returns the ApiUrlFile field if non-nil, zero value otherwise.

### GetApiUrlFileOk

`func (o *ConfigSlackConfig) GetApiUrlFileOk() (*string, bool)`

GetApiUrlFileOk returns a tuple with the ApiUrlFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrlFile

`func (o *ConfigSlackConfig) SetApiUrlFile(v string)`

SetApiUrlFile sets ApiUrlFile field to given value.

### HasApiUrlFile

`func (o *ConfigSlackConfig) HasApiUrlFile() bool`

HasApiUrlFile returns a boolean if a field has been set.

### GetAppToken

`func (o *ConfigSlackConfig) GetAppToken() string`

GetAppToken returns the AppToken field if non-nil, zero value otherwise.

### GetAppTokenOk

`func (o *ConfigSlackConfig) GetAppTokenOk() (*string, bool)`

GetAppTokenOk returns a tuple with the AppToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppToken

`func (o *ConfigSlackConfig) SetAppToken(v string)`

SetAppToken sets AppToken field to given value.

### HasAppToken

`func (o *ConfigSlackConfig) HasAppToken() bool`

HasAppToken returns a boolean if a field has been set.

### GetAppTokenFile

`func (o *ConfigSlackConfig) GetAppTokenFile() string`

GetAppTokenFile returns the AppTokenFile field if non-nil, zero value otherwise.

### GetAppTokenFileOk

`func (o *ConfigSlackConfig) GetAppTokenFileOk() (*string, bool)`

GetAppTokenFileOk returns a tuple with the AppTokenFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTokenFile

`func (o *ConfigSlackConfig) SetAppTokenFile(v string)`

SetAppTokenFile sets AppTokenFile field to given value.

### HasAppTokenFile

`func (o *ConfigSlackConfig) HasAppTokenFile() bool`

HasAppTokenFile returns a boolean if a field has been set.

### GetAppUrl

`func (o *ConfigSlackConfig) GetAppUrl() map[string]interface{}`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *ConfigSlackConfig) GetAppUrlOk() (*map[string]interface{}, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *ConfigSlackConfig) SetAppUrl(v map[string]interface{})`

SetAppUrl sets AppUrl field to given value.

### HasAppUrl

`func (o *ConfigSlackConfig) HasAppUrl() bool`

HasAppUrl returns a boolean if a field has been set.

### GetCallbackId

`func (o *ConfigSlackConfig) GetCallbackId() string`

GetCallbackId returns the CallbackId field if non-nil, zero value otherwise.

### GetCallbackIdOk

`func (o *ConfigSlackConfig) GetCallbackIdOk() (*string, bool)`

GetCallbackIdOk returns a tuple with the CallbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackId

`func (o *ConfigSlackConfig) SetCallbackId(v string)`

SetCallbackId sets CallbackId field to given value.

### HasCallbackId

`func (o *ConfigSlackConfig) HasCallbackId() bool`

HasCallbackId returns a boolean if a field has been set.

### GetChannel

`func (o *ConfigSlackConfig) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ConfigSlackConfig) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ConfigSlackConfig) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ConfigSlackConfig) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetColor

`func (o *ConfigSlackConfig) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ConfigSlackConfig) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ConfigSlackConfig) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ConfigSlackConfig) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetFallback

`func (o *ConfigSlackConfig) GetFallback() string`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *ConfigSlackConfig) GetFallbackOk() (*string, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *ConfigSlackConfig) SetFallback(v string)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *ConfigSlackConfig) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetFields

`func (o *ConfigSlackConfig) GetFields() []ConfigSlackField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ConfigSlackConfig) GetFieldsOk() (*[]ConfigSlackField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ConfigSlackConfig) SetFields(v []ConfigSlackField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ConfigSlackConfig) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFooter

`func (o *ConfigSlackConfig) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *ConfigSlackConfig) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *ConfigSlackConfig) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *ConfigSlackConfig) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetHttpConfig

`func (o *ConfigSlackConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigSlackConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigSlackConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigSlackConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetIconEmoji

`func (o *ConfigSlackConfig) GetIconEmoji() string`

GetIconEmoji returns the IconEmoji field if non-nil, zero value otherwise.

### GetIconEmojiOk

`func (o *ConfigSlackConfig) GetIconEmojiOk() (*string, bool)`

GetIconEmojiOk returns a tuple with the IconEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconEmoji

`func (o *ConfigSlackConfig) SetIconEmoji(v string)`

SetIconEmoji sets IconEmoji field to given value.

### HasIconEmoji

`func (o *ConfigSlackConfig) HasIconEmoji() bool`

HasIconEmoji returns a boolean if a field has been set.

### GetIconUrl

`func (o *ConfigSlackConfig) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ConfigSlackConfig) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ConfigSlackConfig) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *ConfigSlackConfig) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *ConfigSlackConfig) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ConfigSlackConfig) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ConfigSlackConfig) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ConfigSlackConfig) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetLinkNames

`func (o *ConfigSlackConfig) GetLinkNames() bool`

GetLinkNames returns the LinkNames field if non-nil, zero value otherwise.

### GetLinkNamesOk

`func (o *ConfigSlackConfig) GetLinkNamesOk() (*bool, bool)`

GetLinkNamesOk returns a tuple with the LinkNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNames

`func (o *ConfigSlackConfig) SetLinkNames(v bool)`

SetLinkNames sets LinkNames field to given value.

### HasLinkNames

`func (o *ConfigSlackConfig) HasLinkNames() bool`

HasLinkNames returns a boolean if a field has been set.

### GetMessageText

`func (o *ConfigSlackConfig) GetMessageText() string`

GetMessageText returns the MessageText field if non-nil, zero value otherwise.

### GetMessageTextOk

`func (o *ConfigSlackConfig) GetMessageTextOk() (*string, bool)`

GetMessageTextOk returns a tuple with the MessageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageText

`func (o *ConfigSlackConfig) SetMessageText(v string)`

SetMessageText sets MessageText field to given value.

### HasMessageText

`func (o *ConfigSlackConfig) HasMessageText() bool`

HasMessageText returns a boolean if a field has been set.

### GetMrkdwnIn

`func (o *ConfigSlackConfig) GetMrkdwnIn() []string`

GetMrkdwnIn returns the MrkdwnIn field if non-nil, zero value otherwise.

### GetMrkdwnInOk

`func (o *ConfigSlackConfig) GetMrkdwnInOk() (*[]string, bool)`

GetMrkdwnInOk returns a tuple with the MrkdwnIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrkdwnIn

`func (o *ConfigSlackConfig) SetMrkdwnIn(v []string)`

SetMrkdwnIn sets MrkdwnIn field to given value.

### HasMrkdwnIn

`func (o *ConfigSlackConfig) HasMrkdwnIn() bool`

HasMrkdwnIn returns a boolean if a field has been set.

### GetPretext

`func (o *ConfigSlackConfig) GetPretext() string`

GetPretext returns the Pretext field if non-nil, zero value otherwise.

### GetPretextOk

`func (o *ConfigSlackConfig) GetPretextOk() (*string, bool)`

GetPretextOk returns a tuple with the Pretext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPretext

`func (o *ConfigSlackConfig) SetPretext(v string)`

SetPretext sets Pretext field to given value.

### HasPretext

`func (o *ConfigSlackConfig) HasPretext() bool`

HasPretext returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigSlackConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigSlackConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigSlackConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigSlackConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetShortFields

`func (o *ConfigSlackConfig) GetShortFields() bool`

GetShortFields returns the ShortFields field if non-nil, zero value otherwise.

### GetShortFieldsOk

`func (o *ConfigSlackConfig) GetShortFieldsOk() (*bool, bool)`

GetShortFieldsOk returns a tuple with the ShortFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortFields

`func (o *ConfigSlackConfig) SetShortFields(v bool)`

SetShortFields sets ShortFields field to given value.

### HasShortFields

`func (o *ConfigSlackConfig) HasShortFields() bool`

HasShortFields returns a boolean if a field has been set.

### GetText

`func (o *ConfigSlackConfig) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ConfigSlackConfig) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ConfigSlackConfig) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ConfigSlackConfig) HasText() bool`

HasText returns a boolean if a field has been set.

### GetThumbUrl

`func (o *ConfigSlackConfig) GetThumbUrl() string`

GetThumbUrl returns the ThumbUrl field if non-nil, zero value otherwise.

### GetThumbUrlOk

`func (o *ConfigSlackConfig) GetThumbUrlOk() (*string, bool)`

GetThumbUrlOk returns a tuple with the ThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbUrl

`func (o *ConfigSlackConfig) SetThumbUrl(v string)`

SetThumbUrl sets ThumbUrl field to given value.

### HasThumbUrl

`func (o *ConfigSlackConfig) HasThumbUrl() bool`

HasThumbUrl returns a boolean if a field has been set.

### GetTimeout

`func (o *ConfigSlackConfig) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ConfigSlackConfig) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ConfigSlackConfig) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ConfigSlackConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTitle

`func (o *ConfigSlackConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConfigSlackConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConfigSlackConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConfigSlackConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleLink

`func (o *ConfigSlackConfig) GetTitleLink() string`

GetTitleLink returns the TitleLink field if non-nil, zero value otherwise.

### GetTitleLinkOk

`func (o *ConfigSlackConfig) GetTitleLinkOk() (*string, bool)`

GetTitleLinkOk returns a tuple with the TitleLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleLink

`func (o *ConfigSlackConfig) SetTitleLink(v string)`

SetTitleLink sets TitleLink field to given value.

### HasTitleLink

`func (o *ConfigSlackConfig) HasTitleLink() bool`

HasTitleLink returns a boolean if a field has been set.

### GetUsername

`func (o *ConfigSlackConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ConfigSlackConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ConfigSlackConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ConfigSlackConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


