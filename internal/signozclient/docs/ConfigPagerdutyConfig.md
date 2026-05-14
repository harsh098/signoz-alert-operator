# ConfigPagerdutyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** |  | [optional] 
**Client** | Pointer to **string** |  | [optional] 
**ClientUrl** | Pointer to **string** |  | [optional] 
**Component** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **map[string]interface{}** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**Images** | Pointer to [**[]ConfigPagerdutyImage**](ConfigPagerdutyImage.md) |  | [optional] 
**Links** | Pointer to [**[]ConfigPagerdutyLink**](ConfigPagerdutyLink.md) |  | [optional] 
**RoutingKey** | Pointer to **string** |  | [optional] 
**RoutingKeyFile** | Pointer to **string** |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**ServiceKey** | Pointer to **string** |  | [optional] 
**ServiceKeyFile** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 
**Url** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConfigPagerdutyConfig

`func NewConfigPagerdutyConfig() *ConfigPagerdutyConfig`

NewConfigPagerdutyConfig instantiates a new ConfigPagerdutyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigPagerdutyConfigWithDefaults

`func NewConfigPagerdutyConfigWithDefaults() *ConfigPagerdutyConfig`

NewConfigPagerdutyConfigWithDefaults instantiates a new ConfigPagerdutyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *ConfigPagerdutyConfig) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *ConfigPagerdutyConfig) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *ConfigPagerdutyConfig) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *ConfigPagerdutyConfig) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetClient

`func (o *ConfigPagerdutyConfig) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ConfigPagerdutyConfig) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ConfigPagerdutyConfig) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *ConfigPagerdutyConfig) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetClientUrl

`func (o *ConfigPagerdutyConfig) GetClientUrl() string`

GetClientUrl returns the ClientUrl field if non-nil, zero value otherwise.

### GetClientUrlOk

`func (o *ConfigPagerdutyConfig) GetClientUrlOk() (*string, bool)`

GetClientUrlOk returns a tuple with the ClientUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUrl

`func (o *ConfigPagerdutyConfig) SetClientUrl(v string)`

SetClientUrl sets ClientUrl field to given value.

### HasClientUrl

`func (o *ConfigPagerdutyConfig) HasClientUrl() bool`

HasClientUrl returns a boolean if a field has been set.

### GetComponent

`func (o *ConfigPagerdutyConfig) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ConfigPagerdutyConfig) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ConfigPagerdutyConfig) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ConfigPagerdutyConfig) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigPagerdutyConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigPagerdutyConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigPagerdutyConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigPagerdutyConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *ConfigPagerdutyConfig) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ConfigPagerdutyConfig) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ConfigPagerdutyConfig) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ConfigPagerdutyConfig) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetGroup

`func (o *ConfigPagerdutyConfig) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ConfigPagerdutyConfig) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ConfigPagerdutyConfig) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ConfigPagerdutyConfig) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHttpConfig

`func (o *ConfigPagerdutyConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigPagerdutyConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigPagerdutyConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigPagerdutyConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetImages

`func (o *ConfigPagerdutyConfig) GetImages() []ConfigPagerdutyImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ConfigPagerdutyConfig) GetImagesOk() (*[]ConfigPagerdutyImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ConfigPagerdutyConfig) SetImages(v []ConfigPagerdutyImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *ConfigPagerdutyConfig) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetLinks

`func (o *ConfigPagerdutyConfig) GetLinks() []ConfigPagerdutyLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConfigPagerdutyConfig) GetLinksOk() (*[]ConfigPagerdutyLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConfigPagerdutyConfig) SetLinks(v []ConfigPagerdutyLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConfigPagerdutyConfig) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRoutingKey

`func (o *ConfigPagerdutyConfig) GetRoutingKey() string`

GetRoutingKey returns the RoutingKey field if non-nil, zero value otherwise.

### GetRoutingKeyOk

`func (o *ConfigPagerdutyConfig) GetRoutingKeyOk() (*string, bool)`

GetRoutingKeyOk returns a tuple with the RoutingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKey

`func (o *ConfigPagerdutyConfig) SetRoutingKey(v string)`

SetRoutingKey sets RoutingKey field to given value.

### HasRoutingKey

`func (o *ConfigPagerdutyConfig) HasRoutingKey() bool`

HasRoutingKey returns a boolean if a field has been set.

### GetRoutingKeyFile

`func (o *ConfigPagerdutyConfig) GetRoutingKeyFile() string`

GetRoutingKeyFile returns the RoutingKeyFile field if non-nil, zero value otherwise.

### GetRoutingKeyFileOk

`func (o *ConfigPagerdutyConfig) GetRoutingKeyFileOk() (*string, bool)`

GetRoutingKeyFileOk returns a tuple with the RoutingKeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKeyFile

`func (o *ConfigPagerdutyConfig) SetRoutingKeyFile(v string)`

SetRoutingKeyFile sets RoutingKeyFile field to given value.

### HasRoutingKeyFile

`func (o *ConfigPagerdutyConfig) HasRoutingKeyFile() bool`

HasRoutingKeyFile returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigPagerdutyConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigPagerdutyConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigPagerdutyConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigPagerdutyConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetServiceKey

`func (o *ConfigPagerdutyConfig) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *ConfigPagerdutyConfig) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *ConfigPagerdutyConfig) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *ConfigPagerdutyConfig) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetServiceKeyFile

`func (o *ConfigPagerdutyConfig) GetServiceKeyFile() string`

GetServiceKeyFile returns the ServiceKeyFile field if non-nil, zero value otherwise.

### GetServiceKeyFileOk

`func (o *ConfigPagerdutyConfig) GetServiceKeyFileOk() (*string, bool)`

GetServiceKeyFileOk returns a tuple with the ServiceKeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKeyFile

`func (o *ConfigPagerdutyConfig) SetServiceKeyFile(v string)`

SetServiceKeyFile sets ServiceKeyFile field to given value.

### HasServiceKeyFile

`func (o *ConfigPagerdutyConfig) HasServiceKeyFile() bool`

HasServiceKeyFile returns a boolean if a field has been set.

### GetSeverity

`func (o *ConfigPagerdutyConfig) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ConfigPagerdutyConfig) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ConfigPagerdutyConfig) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ConfigPagerdutyConfig) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSource

`func (o *ConfigPagerdutyConfig) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConfigPagerdutyConfig) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConfigPagerdutyConfig) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConfigPagerdutyConfig) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTimeout

`func (o *ConfigPagerdutyConfig) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ConfigPagerdutyConfig) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ConfigPagerdutyConfig) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ConfigPagerdutyConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *ConfigPagerdutyConfig) GetUrl() map[string]interface{}`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConfigPagerdutyConfig) GetUrlOk() (*map[string]interface{}, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConfigPagerdutyConfig) SetUrl(v map[string]interface{})`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConfigPagerdutyConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


