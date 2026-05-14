# ConfigRocketchatConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]ConfigRocketchatAttachmentAction**](ConfigRocketchatAttachmentAction.md) |  | [optional] 
**ApiUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Emoji** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**[]ConfigRocketchatAttachmentField**](ConfigRocketchatAttachmentField.md) |  | [optional] 
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**LinkNames** | Pointer to **bool** |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**ShortFields** | Pointer to **bool** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**ThumbUrl** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TitleLink** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TokenFile** | Pointer to **string** |  | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**TokenIdFile** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigRocketchatConfig

`func NewConfigRocketchatConfig() *ConfigRocketchatConfig`

NewConfigRocketchatConfig instantiates a new ConfigRocketchatConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigRocketchatConfigWithDefaults

`func NewConfigRocketchatConfigWithDefaults() *ConfigRocketchatConfig`

NewConfigRocketchatConfigWithDefaults instantiates a new ConfigRocketchatConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ConfigRocketchatConfig) GetActions() []ConfigRocketchatAttachmentAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ConfigRocketchatConfig) GetActionsOk() (*[]ConfigRocketchatAttachmentAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ConfigRocketchatConfig) SetActions(v []ConfigRocketchatAttachmentAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ConfigRocketchatConfig) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetApiUrl

`func (o *ConfigRocketchatConfig) GetApiUrl() map[string]interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ConfigRocketchatConfig) GetApiUrlOk() (*map[string]interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ConfigRocketchatConfig) SetApiUrl(v map[string]interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *ConfigRocketchatConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetChannel

`func (o *ConfigRocketchatConfig) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ConfigRocketchatConfig) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ConfigRocketchatConfig) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ConfigRocketchatConfig) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetColor

`func (o *ConfigRocketchatConfig) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ConfigRocketchatConfig) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ConfigRocketchatConfig) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ConfigRocketchatConfig) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetEmoji

`func (o *ConfigRocketchatConfig) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *ConfigRocketchatConfig) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *ConfigRocketchatConfig) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *ConfigRocketchatConfig) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetFields

`func (o *ConfigRocketchatConfig) GetFields() []ConfigRocketchatAttachmentField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ConfigRocketchatConfig) GetFieldsOk() (*[]ConfigRocketchatAttachmentField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ConfigRocketchatConfig) SetFields(v []ConfigRocketchatAttachmentField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ConfigRocketchatConfig) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetHttpConfig

`func (o *ConfigRocketchatConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigRocketchatConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigRocketchatConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigRocketchatConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetIconUrl

`func (o *ConfigRocketchatConfig) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ConfigRocketchatConfig) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ConfigRocketchatConfig) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *ConfigRocketchatConfig) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *ConfigRocketchatConfig) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ConfigRocketchatConfig) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ConfigRocketchatConfig) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ConfigRocketchatConfig) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetLinkNames

`func (o *ConfigRocketchatConfig) GetLinkNames() bool`

GetLinkNames returns the LinkNames field if non-nil, zero value otherwise.

### GetLinkNamesOk

`func (o *ConfigRocketchatConfig) GetLinkNamesOk() (*bool, bool)`

GetLinkNamesOk returns a tuple with the LinkNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNames

`func (o *ConfigRocketchatConfig) SetLinkNames(v bool)`

SetLinkNames sets LinkNames field to given value.

### HasLinkNames

`func (o *ConfigRocketchatConfig) HasLinkNames() bool`

HasLinkNames returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigRocketchatConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigRocketchatConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigRocketchatConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigRocketchatConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetShortFields

`func (o *ConfigRocketchatConfig) GetShortFields() bool`

GetShortFields returns the ShortFields field if non-nil, zero value otherwise.

### GetShortFieldsOk

`func (o *ConfigRocketchatConfig) GetShortFieldsOk() (*bool, bool)`

GetShortFieldsOk returns a tuple with the ShortFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortFields

`func (o *ConfigRocketchatConfig) SetShortFields(v bool)`

SetShortFields sets ShortFields field to given value.

### HasShortFields

`func (o *ConfigRocketchatConfig) HasShortFields() bool`

HasShortFields returns a boolean if a field has been set.

### GetText

`func (o *ConfigRocketchatConfig) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ConfigRocketchatConfig) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ConfigRocketchatConfig) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ConfigRocketchatConfig) HasText() bool`

HasText returns a boolean if a field has been set.

### GetThumbUrl

`func (o *ConfigRocketchatConfig) GetThumbUrl() string`

GetThumbUrl returns the ThumbUrl field if non-nil, zero value otherwise.

### GetThumbUrlOk

`func (o *ConfigRocketchatConfig) GetThumbUrlOk() (*string, bool)`

GetThumbUrlOk returns a tuple with the ThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbUrl

`func (o *ConfigRocketchatConfig) SetThumbUrl(v string)`

SetThumbUrl sets ThumbUrl field to given value.

### HasThumbUrl

`func (o *ConfigRocketchatConfig) HasThumbUrl() bool`

HasThumbUrl returns a boolean if a field has been set.

### GetTitle

`func (o *ConfigRocketchatConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConfigRocketchatConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConfigRocketchatConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConfigRocketchatConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleLink

`func (o *ConfigRocketchatConfig) GetTitleLink() string`

GetTitleLink returns the TitleLink field if non-nil, zero value otherwise.

### GetTitleLinkOk

`func (o *ConfigRocketchatConfig) GetTitleLinkOk() (*string, bool)`

GetTitleLinkOk returns a tuple with the TitleLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleLink

`func (o *ConfigRocketchatConfig) SetTitleLink(v string)`

SetTitleLink sets TitleLink field to given value.

### HasTitleLink

`func (o *ConfigRocketchatConfig) HasTitleLink() bool`

HasTitleLink returns a boolean if a field has been set.

### GetToken

`func (o *ConfigRocketchatConfig) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ConfigRocketchatConfig) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ConfigRocketchatConfig) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ConfigRocketchatConfig) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenFile

`func (o *ConfigRocketchatConfig) GetTokenFile() string`

GetTokenFile returns the TokenFile field if non-nil, zero value otherwise.

### GetTokenFileOk

`func (o *ConfigRocketchatConfig) GetTokenFileOk() (*string, bool)`

GetTokenFileOk returns a tuple with the TokenFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFile

`func (o *ConfigRocketchatConfig) SetTokenFile(v string)`

SetTokenFile sets TokenFile field to given value.

### HasTokenFile

`func (o *ConfigRocketchatConfig) HasTokenFile() bool`

HasTokenFile returns a boolean if a field has been set.

### GetTokenId

`func (o *ConfigRocketchatConfig) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ConfigRocketchatConfig) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ConfigRocketchatConfig) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ConfigRocketchatConfig) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTokenIdFile

`func (o *ConfigRocketchatConfig) GetTokenIdFile() string`

GetTokenIdFile returns the TokenIdFile field if non-nil, zero value otherwise.

### GetTokenIdFileOk

`func (o *ConfigRocketchatConfig) GetTokenIdFileOk() (*string, bool)`

GetTokenIdFileOk returns a tuple with the TokenIdFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIdFile

`func (o *ConfigRocketchatConfig) SetTokenIdFile(v string)`

SetTokenIdFile sets TokenIdFile field to given value.

### HasTokenIdFile

`func (o *ConfigRocketchatConfig) HasTokenIdFile() bool`

HasTokenIdFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


