# ConfigReceiver

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
**Name** | Pointer to **string** |  | [optional] 
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

### NewConfigReceiver

`func NewConfigReceiver() *ConfigReceiver`

NewConfigReceiver instantiates a new ConfigReceiver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigReceiverWithDefaults

`func NewConfigReceiverWithDefaults() *ConfigReceiver`

NewConfigReceiverWithDefaults instantiates a new ConfigReceiver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscordConfigs

`func (o *ConfigReceiver) GetDiscordConfigs() []ConfigDiscordConfig`

GetDiscordConfigs returns the DiscordConfigs field if non-nil, zero value otherwise.

### GetDiscordConfigsOk

`func (o *ConfigReceiver) GetDiscordConfigsOk() (*[]ConfigDiscordConfig, bool)`

GetDiscordConfigsOk returns a tuple with the DiscordConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordConfigs

`func (o *ConfigReceiver) SetDiscordConfigs(v []ConfigDiscordConfig)`

SetDiscordConfigs sets DiscordConfigs field to given value.

### HasDiscordConfigs

`func (o *ConfigReceiver) HasDiscordConfigs() bool`

HasDiscordConfigs returns a boolean if a field has been set.

### GetEmailConfigs

`func (o *ConfigReceiver) GetEmailConfigs() []ConfigEmailConfig`

GetEmailConfigs returns the EmailConfigs field if non-nil, zero value otherwise.

### GetEmailConfigsOk

`func (o *ConfigReceiver) GetEmailConfigsOk() (*[]ConfigEmailConfig, bool)`

GetEmailConfigsOk returns a tuple with the EmailConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfigs

`func (o *ConfigReceiver) SetEmailConfigs(v []ConfigEmailConfig)`

SetEmailConfigs sets EmailConfigs field to given value.

### HasEmailConfigs

`func (o *ConfigReceiver) HasEmailConfigs() bool`

HasEmailConfigs returns a boolean if a field has been set.

### GetIncidentioConfigs

`func (o *ConfigReceiver) GetIncidentioConfigs() []ConfigIncidentioConfig`

GetIncidentioConfigs returns the IncidentioConfigs field if non-nil, zero value otherwise.

### GetIncidentioConfigsOk

`func (o *ConfigReceiver) GetIncidentioConfigsOk() (*[]ConfigIncidentioConfig, bool)`

GetIncidentioConfigsOk returns a tuple with the IncidentioConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentioConfigs

`func (o *ConfigReceiver) SetIncidentioConfigs(v []ConfigIncidentioConfig)`

SetIncidentioConfigs sets IncidentioConfigs field to given value.

### HasIncidentioConfigs

`func (o *ConfigReceiver) HasIncidentioConfigs() bool`

HasIncidentioConfigs returns a boolean if a field has been set.

### GetJiraConfigs

`func (o *ConfigReceiver) GetJiraConfigs() []ConfigJiraConfig`

GetJiraConfigs returns the JiraConfigs field if non-nil, zero value otherwise.

### GetJiraConfigsOk

`func (o *ConfigReceiver) GetJiraConfigsOk() (*[]ConfigJiraConfig, bool)`

GetJiraConfigsOk returns a tuple with the JiraConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraConfigs

`func (o *ConfigReceiver) SetJiraConfigs(v []ConfigJiraConfig)`

SetJiraConfigs sets JiraConfigs field to given value.

### HasJiraConfigs

`func (o *ConfigReceiver) HasJiraConfigs() bool`

HasJiraConfigs returns a boolean if a field has been set.

### GetMattermostConfigs

`func (o *ConfigReceiver) GetMattermostConfigs() []ConfigMattermostConfig`

GetMattermostConfigs returns the MattermostConfigs field if non-nil, zero value otherwise.

### GetMattermostConfigsOk

`func (o *ConfigReceiver) GetMattermostConfigsOk() (*[]ConfigMattermostConfig, bool)`

GetMattermostConfigsOk returns a tuple with the MattermostConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMattermostConfigs

`func (o *ConfigReceiver) SetMattermostConfigs(v []ConfigMattermostConfig)`

SetMattermostConfigs sets MattermostConfigs field to given value.

### HasMattermostConfigs

`func (o *ConfigReceiver) HasMattermostConfigs() bool`

HasMattermostConfigs returns a boolean if a field has been set.

### GetMsteamsConfigs

`func (o *ConfigReceiver) GetMsteamsConfigs() []ConfigMSTeamsConfig`

GetMsteamsConfigs returns the MsteamsConfigs field if non-nil, zero value otherwise.

### GetMsteamsConfigsOk

`func (o *ConfigReceiver) GetMsteamsConfigsOk() (*[]ConfigMSTeamsConfig, bool)`

GetMsteamsConfigsOk returns a tuple with the MsteamsConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsteamsConfigs

`func (o *ConfigReceiver) SetMsteamsConfigs(v []ConfigMSTeamsConfig)`

SetMsteamsConfigs sets MsteamsConfigs field to given value.

### HasMsteamsConfigs

`func (o *ConfigReceiver) HasMsteamsConfigs() bool`

HasMsteamsConfigs returns a boolean if a field has been set.

### GetMsteamsv2Configs

`func (o *ConfigReceiver) GetMsteamsv2Configs() []ConfigMSTeamsV2Config`

GetMsteamsv2Configs returns the Msteamsv2Configs field if non-nil, zero value otherwise.

### GetMsteamsv2ConfigsOk

`func (o *ConfigReceiver) GetMsteamsv2ConfigsOk() (*[]ConfigMSTeamsV2Config, bool)`

GetMsteamsv2ConfigsOk returns a tuple with the Msteamsv2Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsteamsv2Configs

`func (o *ConfigReceiver) SetMsteamsv2Configs(v []ConfigMSTeamsV2Config)`

SetMsteamsv2Configs sets Msteamsv2Configs field to given value.

### HasMsteamsv2Configs

`func (o *ConfigReceiver) HasMsteamsv2Configs() bool`

HasMsteamsv2Configs returns a boolean if a field has been set.

### GetName

`func (o *ConfigReceiver) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigReceiver) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigReceiver) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigReceiver) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpsgenieConfigs

`func (o *ConfigReceiver) GetOpsgenieConfigs() []ConfigOpsGenieConfig`

GetOpsgenieConfigs returns the OpsgenieConfigs field if non-nil, zero value otherwise.

### GetOpsgenieConfigsOk

`func (o *ConfigReceiver) GetOpsgenieConfigsOk() (*[]ConfigOpsGenieConfig, bool)`

GetOpsgenieConfigsOk returns a tuple with the OpsgenieConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsgenieConfigs

`func (o *ConfigReceiver) SetOpsgenieConfigs(v []ConfigOpsGenieConfig)`

SetOpsgenieConfigs sets OpsgenieConfigs field to given value.

### HasOpsgenieConfigs

`func (o *ConfigReceiver) HasOpsgenieConfigs() bool`

HasOpsgenieConfigs returns a boolean if a field has been set.

### GetPagerdutyConfigs

`func (o *ConfigReceiver) GetPagerdutyConfigs() []ConfigPagerdutyConfig`

GetPagerdutyConfigs returns the PagerdutyConfigs field if non-nil, zero value otherwise.

### GetPagerdutyConfigsOk

`func (o *ConfigReceiver) GetPagerdutyConfigsOk() (*[]ConfigPagerdutyConfig, bool)`

GetPagerdutyConfigsOk returns a tuple with the PagerdutyConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerdutyConfigs

`func (o *ConfigReceiver) SetPagerdutyConfigs(v []ConfigPagerdutyConfig)`

SetPagerdutyConfigs sets PagerdutyConfigs field to given value.

### HasPagerdutyConfigs

`func (o *ConfigReceiver) HasPagerdutyConfigs() bool`

HasPagerdutyConfigs returns a boolean if a field has been set.

### GetPushoverConfigs

`func (o *ConfigReceiver) GetPushoverConfigs() []ConfigPushoverConfig`

GetPushoverConfigs returns the PushoverConfigs field if non-nil, zero value otherwise.

### GetPushoverConfigsOk

`func (o *ConfigReceiver) GetPushoverConfigsOk() (*[]ConfigPushoverConfig, bool)`

GetPushoverConfigsOk returns a tuple with the PushoverConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushoverConfigs

`func (o *ConfigReceiver) SetPushoverConfigs(v []ConfigPushoverConfig)`

SetPushoverConfigs sets PushoverConfigs field to given value.

### HasPushoverConfigs

`func (o *ConfigReceiver) HasPushoverConfigs() bool`

HasPushoverConfigs returns a boolean if a field has been set.

### GetRocketchatConfigs

`func (o *ConfigReceiver) GetRocketchatConfigs() []ConfigRocketchatConfig`

GetRocketchatConfigs returns the RocketchatConfigs field if non-nil, zero value otherwise.

### GetRocketchatConfigsOk

`func (o *ConfigReceiver) GetRocketchatConfigsOk() (*[]ConfigRocketchatConfig, bool)`

GetRocketchatConfigsOk returns a tuple with the RocketchatConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRocketchatConfigs

`func (o *ConfigReceiver) SetRocketchatConfigs(v []ConfigRocketchatConfig)`

SetRocketchatConfigs sets RocketchatConfigs field to given value.

### HasRocketchatConfigs

`func (o *ConfigReceiver) HasRocketchatConfigs() bool`

HasRocketchatConfigs returns a boolean if a field has been set.

### GetSlackConfigs

`func (o *ConfigReceiver) GetSlackConfigs() []ConfigSlackConfig`

GetSlackConfigs returns the SlackConfigs field if non-nil, zero value otherwise.

### GetSlackConfigsOk

`func (o *ConfigReceiver) GetSlackConfigsOk() (*[]ConfigSlackConfig, bool)`

GetSlackConfigsOk returns a tuple with the SlackConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConfigs

`func (o *ConfigReceiver) SetSlackConfigs(v []ConfigSlackConfig)`

SetSlackConfigs sets SlackConfigs field to given value.

### HasSlackConfigs

`func (o *ConfigReceiver) HasSlackConfigs() bool`

HasSlackConfigs returns a boolean if a field has been set.

### GetSnsConfigs

`func (o *ConfigReceiver) GetSnsConfigs() []ConfigSNSConfig`

GetSnsConfigs returns the SnsConfigs field if non-nil, zero value otherwise.

### GetSnsConfigsOk

`func (o *ConfigReceiver) GetSnsConfigsOk() (*[]ConfigSNSConfig, bool)`

GetSnsConfigsOk returns a tuple with the SnsConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnsConfigs

`func (o *ConfigReceiver) SetSnsConfigs(v []ConfigSNSConfig)`

SetSnsConfigs sets SnsConfigs field to given value.

### HasSnsConfigs

`func (o *ConfigReceiver) HasSnsConfigs() bool`

HasSnsConfigs returns a boolean if a field has been set.

### GetTelegramConfigs

`func (o *ConfigReceiver) GetTelegramConfigs() []ConfigTelegramConfig`

GetTelegramConfigs returns the TelegramConfigs field if non-nil, zero value otherwise.

### GetTelegramConfigsOk

`func (o *ConfigReceiver) GetTelegramConfigsOk() (*[]ConfigTelegramConfig, bool)`

GetTelegramConfigsOk returns a tuple with the TelegramConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramConfigs

`func (o *ConfigReceiver) SetTelegramConfigs(v []ConfigTelegramConfig)`

SetTelegramConfigs sets TelegramConfigs field to given value.

### HasTelegramConfigs

`func (o *ConfigReceiver) HasTelegramConfigs() bool`

HasTelegramConfigs returns a boolean if a field has been set.

### GetVictoropsConfigs

`func (o *ConfigReceiver) GetVictoropsConfigs() []ConfigVictorOpsConfig`

GetVictoropsConfigs returns the VictoropsConfigs field if non-nil, zero value otherwise.

### GetVictoropsConfigsOk

`func (o *ConfigReceiver) GetVictoropsConfigsOk() (*[]ConfigVictorOpsConfig, bool)`

GetVictoropsConfigsOk returns a tuple with the VictoropsConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVictoropsConfigs

`func (o *ConfigReceiver) SetVictoropsConfigs(v []ConfigVictorOpsConfig)`

SetVictoropsConfigs sets VictoropsConfigs field to given value.

### HasVictoropsConfigs

`func (o *ConfigReceiver) HasVictoropsConfigs() bool`

HasVictoropsConfigs returns a boolean if a field has been set.

### GetWebexConfigs

`func (o *ConfigReceiver) GetWebexConfigs() []ConfigWebexConfig`

GetWebexConfigs returns the WebexConfigs field if non-nil, zero value otherwise.

### GetWebexConfigsOk

`func (o *ConfigReceiver) GetWebexConfigsOk() (*[]ConfigWebexConfig, bool)`

GetWebexConfigsOk returns a tuple with the WebexConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebexConfigs

`func (o *ConfigReceiver) SetWebexConfigs(v []ConfigWebexConfig)`

SetWebexConfigs sets WebexConfigs field to given value.

### HasWebexConfigs

`func (o *ConfigReceiver) HasWebexConfigs() bool`

HasWebexConfigs returns a boolean if a field has been set.

### GetWebhookConfigs

`func (o *ConfigReceiver) GetWebhookConfigs() []ConfigWebhookConfig`

GetWebhookConfigs returns the WebhookConfigs field if non-nil, zero value otherwise.

### GetWebhookConfigsOk

`func (o *ConfigReceiver) GetWebhookConfigsOk() (*[]ConfigWebhookConfig, bool)`

GetWebhookConfigsOk returns a tuple with the WebhookConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookConfigs

`func (o *ConfigReceiver) SetWebhookConfigs(v []ConfigWebhookConfig)`

SetWebhookConfigs sets WebhookConfigs field to given value.

### HasWebhookConfigs

`func (o *ConfigReceiver) HasWebhookConfigs() bool`

HasWebhookConfigs returns a boolean if a field has been set.

### GetWechatConfigs

`func (o *ConfigReceiver) GetWechatConfigs() []ConfigWechatConfig`

GetWechatConfigs returns the WechatConfigs field if non-nil, zero value otherwise.

### GetWechatConfigsOk

`func (o *ConfigReceiver) GetWechatConfigsOk() (*[]ConfigWechatConfig, bool)`

GetWechatConfigsOk returns a tuple with the WechatConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWechatConfigs

`func (o *ConfigReceiver) SetWechatConfigs(v []ConfigWechatConfig)`

SetWechatConfigs sets WechatConfigs field to given value.

### HasWechatConfigs

`func (o *ConfigReceiver) HasWechatConfigs() bool`

HasWechatConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


