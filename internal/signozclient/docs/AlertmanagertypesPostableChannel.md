# AlertmanagertypesPostableChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscordConfigs** | Pointer to [**[]ConfigDiscordConfig**](ConfigDiscordConfig.md) |  | [optional] 
**EmailConfigs** | Pointer to [**[]ConfigEmailConfig**](ConfigEmailConfig.md) |  | [optional] 
**IncidentioConfigs** | Pointer to [**[]ConfigIncidentioConfig**](ConfigIncidentioConfig.md) |  | [optional] 
**JiraConfigs** | Pointer to [**[]ConfigJiraConfig**](ConfigJiraConfig.md) |  | [optional] 
**MattermostConfigs** | Pointer to [**[]ConfigMattermostConfig**](ConfigMattermostConfig.md) |  | [optional] 
**MsteamsConfigs** | Pointer to [**[]ConfigMSTeamsConfig**](ConfigMSTeamsConfig.md) |  | [optional] 
**Msteamsv2Configs** | Pointer to [**[]ConfigMSTeamsV2Config**](ConfigMSTeamsV2Config.md) |  | [optional] 
**Name** | **string** |  | 
**OpsgenieConfigs** | Pointer to [**[]ConfigOpsGenieConfig**](ConfigOpsGenieConfig.md) |  | [optional] 
**PagerdutyConfigs** | Pointer to [**[]ConfigPagerdutyConfig**](ConfigPagerdutyConfig.md) |  | [optional] 
**PushoverConfigs** | Pointer to [**[]ConfigPushoverConfig**](ConfigPushoverConfig.md) |  | [optional] 
**RocketchatConfigs** | Pointer to [**[]ConfigRocketchatConfig**](ConfigRocketchatConfig.md) |  | [optional] 
**SlackConfigs** | Pointer to [**[]ConfigSlackConfig**](ConfigSlackConfig.md) |  | [optional] 
**SnsConfigs** | Pointer to [**[]ConfigSNSConfig**](ConfigSNSConfig.md) |  | [optional] 
**TelegramConfigs** | Pointer to [**[]ConfigTelegramConfig**](ConfigTelegramConfig.md) |  | [optional] 
**VictoropsConfigs** | Pointer to [**[]ConfigVictorOpsConfig**](ConfigVictorOpsConfig.md) |  | [optional] 
**WebexConfigs** | Pointer to [**[]ConfigWebexConfig**](ConfigWebexConfig.md) |  | [optional] 
**WebhookConfigs** | Pointer to [**[]ConfigWebhookConfig**](ConfigWebhookConfig.md) |  | [optional] 
**WechatConfigs** | Pointer to [**[]ConfigWechatConfig**](ConfigWechatConfig.md) |  | [optional] 

## Methods

### NewAlertmanagertypesPostableChannel

`func NewAlertmanagertypesPostableChannel(name string, ) *AlertmanagertypesPostableChannel`

NewAlertmanagertypesPostableChannel instantiates a new AlertmanagertypesPostableChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertmanagertypesPostableChannelWithDefaults

`func NewAlertmanagertypesPostableChannelWithDefaults() *AlertmanagertypesPostableChannel`

NewAlertmanagertypesPostableChannelWithDefaults instantiates a new AlertmanagertypesPostableChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscordConfigs

`func (o *AlertmanagertypesPostableChannel) GetDiscordConfigs() []ConfigDiscordConfig`

GetDiscordConfigs returns the DiscordConfigs field if non-nil, zero value otherwise.

### GetDiscordConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetDiscordConfigsOk() (*[]ConfigDiscordConfig, bool)`

GetDiscordConfigsOk returns a tuple with the DiscordConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordConfigs

`func (o *AlertmanagertypesPostableChannel) SetDiscordConfigs(v []ConfigDiscordConfig)`

SetDiscordConfigs sets DiscordConfigs field to given value.

### HasDiscordConfigs

`func (o *AlertmanagertypesPostableChannel) HasDiscordConfigs() bool`

HasDiscordConfigs returns a boolean if a field has been set.

### GetEmailConfigs

`func (o *AlertmanagertypesPostableChannel) GetEmailConfigs() []ConfigEmailConfig`

GetEmailConfigs returns the EmailConfigs field if non-nil, zero value otherwise.

### GetEmailConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetEmailConfigsOk() (*[]ConfigEmailConfig, bool)`

GetEmailConfigsOk returns a tuple with the EmailConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfigs

`func (o *AlertmanagertypesPostableChannel) SetEmailConfigs(v []ConfigEmailConfig)`

SetEmailConfigs sets EmailConfigs field to given value.

### HasEmailConfigs

`func (o *AlertmanagertypesPostableChannel) HasEmailConfigs() bool`

HasEmailConfigs returns a boolean if a field has been set.

### GetIncidentioConfigs

`func (o *AlertmanagertypesPostableChannel) GetIncidentioConfigs() []ConfigIncidentioConfig`

GetIncidentioConfigs returns the IncidentioConfigs field if non-nil, zero value otherwise.

### GetIncidentioConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetIncidentioConfigsOk() (*[]ConfigIncidentioConfig, bool)`

GetIncidentioConfigsOk returns a tuple with the IncidentioConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentioConfigs

`func (o *AlertmanagertypesPostableChannel) SetIncidentioConfigs(v []ConfigIncidentioConfig)`

SetIncidentioConfigs sets IncidentioConfigs field to given value.

### HasIncidentioConfigs

`func (o *AlertmanagertypesPostableChannel) HasIncidentioConfigs() bool`

HasIncidentioConfigs returns a boolean if a field has been set.

### GetJiraConfigs

`func (o *AlertmanagertypesPostableChannel) GetJiraConfigs() []ConfigJiraConfig`

GetJiraConfigs returns the JiraConfigs field if non-nil, zero value otherwise.

### GetJiraConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetJiraConfigsOk() (*[]ConfigJiraConfig, bool)`

GetJiraConfigsOk returns a tuple with the JiraConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraConfigs

`func (o *AlertmanagertypesPostableChannel) SetJiraConfigs(v []ConfigJiraConfig)`

SetJiraConfigs sets JiraConfigs field to given value.

### HasJiraConfigs

`func (o *AlertmanagertypesPostableChannel) HasJiraConfigs() bool`

HasJiraConfigs returns a boolean if a field has been set.

### GetMattermostConfigs

`func (o *AlertmanagertypesPostableChannel) GetMattermostConfigs() []ConfigMattermostConfig`

GetMattermostConfigs returns the MattermostConfigs field if non-nil, zero value otherwise.

### GetMattermostConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetMattermostConfigsOk() (*[]ConfigMattermostConfig, bool)`

GetMattermostConfigsOk returns a tuple with the MattermostConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMattermostConfigs

`func (o *AlertmanagertypesPostableChannel) SetMattermostConfigs(v []ConfigMattermostConfig)`

SetMattermostConfigs sets MattermostConfigs field to given value.

### HasMattermostConfigs

`func (o *AlertmanagertypesPostableChannel) HasMattermostConfigs() bool`

HasMattermostConfigs returns a boolean if a field has been set.

### GetMsteamsConfigs

`func (o *AlertmanagertypesPostableChannel) GetMsteamsConfigs() []ConfigMSTeamsConfig`

GetMsteamsConfigs returns the MsteamsConfigs field if non-nil, zero value otherwise.

### GetMsteamsConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetMsteamsConfigsOk() (*[]ConfigMSTeamsConfig, bool)`

GetMsteamsConfigsOk returns a tuple with the MsteamsConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsteamsConfigs

`func (o *AlertmanagertypesPostableChannel) SetMsteamsConfigs(v []ConfigMSTeamsConfig)`

SetMsteamsConfigs sets MsteamsConfigs field to given value.

### HasMsteamsConfigs

`func (o *AlertmanagertypesPostableChannel) HasMsteamsConfigs() bool`

HasMsteamsConfigs returns a boolean if a field has been set.

### GetMsteamsv2Configs

`func (o *AlertmanagertypesPostableChannel) GetMsteamsv2Configs() []ConfigMSTeamsV2Config`

GetMsteamsv2Configs returns the Msteamsv2Configs field if non-nil, zero value otherwise.

### GetMsteamsv2ConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetMsteamsv2ConfigsOk() (*[]ConfigMSTeamsV2Config, bool)`

GetMsteamsv2ConfigsOk returns a tuple with the Msteamsv2Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsteamsv2Configs

`func (o *AlertmanagertypesPostableChannel) SetMsteamsv2Configs(v []ConfigMSTeamsV2Config)`

SetMsteamsv2Configs sets Msteamsv2Configs field to given value.

### HasMsteamsv2Configs

`func (o *AlertmanagertypesPostableChannel) HasMsteamsv2Configs() bool`

HasMsteamsv2Configs returns a boolean if a field has been set.

### GetName

`func (o *AlertmanagertypesPostableChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertmanagertypesPostableChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertmanagertypesPostableChannel) SetName(v string)`

SetName sets Name field to given value.


### GetOpsgenieConfigs

`func (o *AlertmanagertypesPostableChannel) GetOpsgenieConfigs() []ConfigOpsGenieConfig`

GetOpsgenieConfigs returns the OpsgenieConfigs field if non-nil, zero value otherwise.

### GetOpsgenieConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetOpsgenieConfigsOk() (*[]ConfigOpsGenieConfig, bool)`

GetOpsgenieConfigsOk returns a tuple with the OpsgenieConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsgenieConfigs

`func (o *AlertmanagertypesPostableChannel) SetOpsgenieConfigs(v []ConfigOpsGenieConfig)`

SetOpsgenieConfigs sets OpsgenieConfigs field to given value.

### HasOpsgenieConfigs

`func (o *AlertmanagertypesPostableChannel) HasOpsgenieConfigs() bool`

HasOpsgenieConfigs returns a boolean if a field has been set.

### GetPagerdutyConfigs

`func (o *AlertmanagertypesPostableChannel) GetPagerdutyConfigs() []ConfigPagerdutyConfig`

GetPagerdutyConfigs returns the PagerdutyConfigs field if non-nil, zero value otherwise.

### GetPagerdutyConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetPagerdutyConfigsOk() (*[]ConfigPagerdutyConfig, bool)`

GetPagerdutyConfigsOk returns a tuple with the PagerdutyConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerdutyConfigs

`func (o *AlertmanagertypesPostableChannel) SetPagerdutyConfigs(v []ConfigPagerdutyConfig)`

SetPagerdutyConfigs sets PagerdutyConfigs field to given value.

### HasPagerdutyConfigs

`func (o *AlertmanagertypesPostableChannel) HasPagerdutyConfigs() bool`

HasPagerdutyConfigs returns a boolean if a field has been set.

### GetPushoverConfigs

`func (o *AlertmanagertypesPostableChannel) GetPushoverConfigs() []ConfigPushoverConfig`

GetPushoverConfigs returns the PushoverConfigs field if non-nil, zero value otherwise.

### GetPushoverConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetPushoverConfigsOk() (*[]ConfigPushoverConfig, bool)`

GetPushoverConfigsOk returns a tuple with the PushoverConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushoverConfigs

`func (o *AlertmanagertypesPostableChannel) SetPushoverConfigs(v []ConfigPushoverConfig)`

SetPushoverConfigs sets PushoverConfigs field to given value.

### HasPushoverConfigs

`func (o *AlertmanagertypesPostableChannel) HasPushoverConfigs() bool`

HasPushoverConfigs returns a boolean if a field has been set.

### GetRocketchatConfigs

`func (o *AlertmanagertypesPostableChannel) GetRocketchatConfigs() []ConfigRocketchatConfig`

GetRocketchatConfigs returns the RocketchatConfigs field if non-nil, zero value otherwise.

### GetRocketchatConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetRocketchatConfigsOk() (*[]ConfigRocketchatConfig, bool)`

GetRocketchatConfigsOk returns a tuple with the RocketchatConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRocketchatConfigs

`func (o *AlertmanagertypesPostableChannel) SetRocketchatConfigs(v []ConfigRocketchatConfig)`

SetRocketchatConfigs sets RocketchatConfigs field to given value.

### HasRocketchatConfigs

`func (o *AlertmanagertypesPostableChannel) HasRocketchatConfigs() bool`

HasRocketchatConfigs returns a boolean if a field has been set.

### GetSlackConfigs

`func (o *AlertmanagertypesPostableChannel) GetSlackConfigs() []ConfigSlackConfig`

GetSlackConfigs returns the SlackConfigs field if non-nil, zero value otherwise.

### GetSlackConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetSlackConfigsOk() (*[]ConfigSlackConfig, bool)`

GetSlackConfigsOk returns a tuple with the SlackConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConfigs

`func (o *AlertmanagertypesPostableChannel) SetSlackConfigs(v []ConfigSlackConfig)`

SetSlackConfigs sets SlackConfigs field to given value.

### HasSlackConfigs

`func (o *AlertmanagertypesPostableChannel) HasSlackConfigs() bool`

HasSlackConfigs returns a boolean if a field has been set.

### GetSnsConfigs

`func (o *AlertmanagertypesPostableChannel) GetSnsConfigs() []ConfigSNSConfig`

GetSnsConfigs returns the SnsConfigs field if non-nil, zero value otherwise.

### GetSnsConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetSnsConfigsOk() (*[]ConfigSNSConfig, bool)`

GetSnsConfigsOk returns a tuple with the SnsConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnsConfigs

`func (o *AlertmanagertypesPostableChannel) SetSnsConfigs(v []ConfigSNSConfig)`

SetSnsConfigs sets SnsConfigs field to given value.

### HasSnsConfigs

`func (o *AlertmanagertypesPostableChannel) HasSnsConfigs() bool`

HasSnsConfigs returns a boolean if a field has been set.

### GetTelegramConfigs

`func (o *AlertmanagertypesPostableChannel) GetTelegramConfigs() []ConfigTelegramConfig`

GetTelegramConfigs returns the TelegramConfigs field if non-nil, zero value otherwise.

### GetTelegramConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetTelegramConfigsOk() (*[]ConfigTelegramConfig, bool)`

GetTelegramConfigsOk returns a tuple with the TelegramConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramConfigs

`func (o *AlertmanagertypesPostableChannel) SetTelegramConfigs(v []ConfigTelegramConfig)`

SetTelegramConfigs sets TelegramConfigs field to given value.

### HasTelegramConfigs

`func (o *AlertmanagertypesPostableChannel) HasTelegramConfigs() bool`

HasTelegramConfigs returns a boolean if a field has been set.

### GetVictoropsConfigs

`func (o *AlertmanagertypesPostableChannel) GetVictoropsConfigs() []ConfigVictorOpsConfig`

GetVictoropsConfigs returns the VictoropsConfigs field if non-nil, zero value otherwise.

### GetVictoropsConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetVictoropsConfigsOk() (*[]ConfigVictorOpsConfig, bool)`

GetVictoropsConfigsOk returns a tuple with the VictoropsConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVictoropsConfigs

`func (o *AlertmanagertypesPostableChannel) SetVictoropsConfigs(v []ConfigVictorOpsConfig)`

SetVictoropsConfigs sets VictoropsConfigs field to given value.

### HasVictoropsConfigs

`func (o *AlertmanagertypesPostableChannel) HasVictoropsConfigs() bool`

HasVictoropsConfigs returns a boolean if a field has been set.

### GetWebexConfigs

`func (o *AlertmanagertypesPostableChannel) GetWebexConfigs() []ConfigWebexConfig`

GetWebexConfigs returns the WebexConfigs field if non-nil, zero value otherwise.

### GetWebexConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetWebexConfigsOk() (*[]ConfigWebexConfig, bool)`

GetWebexConfigsOk returns a tuple with the WebexConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebexConfigs

`func (o *AlertmanagertypesPostableChannel) SetWebexConfigs(v []ConfigWebexConfig)`

SetWebexConfigs sets WebexConfigs field to given value.

### HasWebexConfigs

`func (o *AlertmanagertypesPostableChannel) HasWebexConfigs() bool`

HasWebexConfigs returns a boolean if a field has been set.

### GetWebhookConfigs

`func (o *AlertmanagertypesPostableChannel) GetWebhookConfigs() []ConfigWebhookConfig`

GetWebhookConfigs returns the WebhookConfigs field if non-nil, zero value otherwise.

### GetWebhookConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetWebhookConfigsOk() (*[]ConfigWebhookConfig, bool)`

GetWebhookConfigsOk returns a tuple with the WebhookConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookConfigs

`func (o *AlertmanagertypesPostableChannel) SetWebhookConfigs(v []ConfigWebhookConfig)`

SetWebhookConfigs sets WebhookConfigs field to given value.

### HasWebhookConfigs

`func (o *AlertmanagertypesPostableChannel) HasWebhookConfigs() bool`

HasWebhookConfigs returns a boolean if a field has been set.

### GetWechatConfigs

`func (o *AlertmanagertypesPostableChannel) GetWechatConfigs() []ConfigWechatConfig`

GetWechatConfigs returns the WechatConfigs field if non-nil, zero value otherwise.

### GetWechatConfigsOk

`func (o *AlertmanagertypesPostableChannel) GetWechatConfigsOk() (*[]ConfigWechatConfig, bool)`

GetWechatConfigsOk returns a tuple with the WechatConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWechatConfigs

`func (o *AlertmanagertypesPostableChannel) SetWechatConfigs(v []ConfigWechatConfig)`

SetWechatConfigs sets WechatConfigs field to given value.

### HasWechatConfigs

`func (o *AlertmanagertypesPostableChannel) HasWechatConfigs() bool`

HasWechatConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


