# CloudintegrationtypesAzureTelemetryCollectionStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**CloudintegrationtypesAzureLogsCollectionStrategy**](CloudintegrationtypesAzureLogsCollectionStrategy.md) |  | [optional] 
**Metrics** | Pointer to **map[string]interface{}** |  | [optional] 
**ResourceProvider** | **string** |  | 
**ResourceType** | **string** |  | 

## Methods

### NewCloudintegrationtypesAzureTelemetryCollectionStrategy

`func NewCloudintegrationtypesAzureTelemetryCollectionStrategy(resourceProvider string, resourceType string, ) *CloudintegrationtypesAzureTelemetryCollectionStrategy`

NewCloudintegrationtypesAzureTelemetryCollectionStrategy instantiates a new CloudintegrationtypesAzureTelemetryCollectionStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudintegrationtypesAzureTelemetryCollectionStrategyWithDefaults

`func NewCloudintegrationtypesAzureTelemetryCollectionStrategyWithDefaults() *CloudintegrationtypesAzureTelemetryCollectionStrategy`

NewCloudintegrationtypesAzureTelemetryCollectionStrategyWithDefaults instantiates a new CloudintegrationtypesAzureTelemetryCollectionStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *CloudintegrationtypesAzureTelemetryCollectionStrategy) GetLogs() CloudintegrationtypesAzureLogsCollectionStrategy`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *CloudintegrationtypesAzureTelemetryCollectionStrategy) GetLogsOk() (*CloudintegrationtypesAzureLogsCollectionStrategy, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *CloudintegrationtypesAzureTelemetryCollectionStrategy) SetLogs(v CloudintegrationtypesAzureLogsCollectionStrategy)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *CloudintegrationtypesAzureTelemetryCollectionStrategy) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMetrics

`func (o *CloudintegrationtypesAzureTelemetryCollectionStrategy) GetMetrics() map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CloudintegrationtypesAzureTelemetryCollectionStrategy) GetMetricsOk() (*map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CloudintegrationtypesAzureTelemetryCollectionStrategy) SetMetrics(v map[string]interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *CloudintegrationtypesAzureTelemetryCollectionStrategy) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetResourceProvider

`func (o *CloudintegrationtypesAzureTelemetryCollectionStrategy) GetResourceProvider() string`

GetResourceProvider returns the ResourceProvider field if non-nil, zero value otherwise.

### GetResourceProviderOk

`func (o *CloudintegrationtypesAzureTelemetryCollectionStrategy) GetResourceProviderOk() (*string, bool)`

GetResourceProviderOk returns a tuple with the ResourceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProvider

`func (o *CloudintegrationtypesAzureTelemetryCollectionStrategy) SetResourceProvider(v string)`

SetResourceProvider sets ResourceProvider field to given value.


### GetResourceType

`func (o *CloudintegrationtypesAzureTelemetryCollectionStrategy) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CloudintegrationtypesAzureTelemetryCollectionStrategy) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CloudintegrationtypesAzureTelemetryCollectionStrategy) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


