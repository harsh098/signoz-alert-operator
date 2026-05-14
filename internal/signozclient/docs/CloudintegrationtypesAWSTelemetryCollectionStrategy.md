# CloudintegrationtypesAWSTelemetryCollectionStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**CloudintegrationtypesAWSLogsCollectionStrategy**](CloudintegrationtypesAWSLogsCollectionStrategy.md) |  | [optional] 
**Metrics** | Pointer to [**CloudintegrationtypesAWSMetricsCollectionStrategy**](CloudintegrationtypesAWSMetricsCollectionStrategy.md) |  | [optional] 
**S3Buckets** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewCloudintegrationtypesAWSTelemetryCollectionStrategy

`func NewCloudintegrationtypesAWSTelemetryCollectionStrategy() *CloudintegrationtypesAWSTelemetryCollectionStrategy`

NewCloudintegrationtypesAWSTelemetryCollectionStrategy instantiates a new CloudintegrationtypesAWSTelemetryCollectionStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudintegrationtypesAWSTelemetryCollectionStrategyWithDefaults

`func NewCloudintegrationtypesAWSTelemetryCollectionStrategyWithDefaults() *CloudintegrationtypesAWSTelemetryCollectionStrategy`

NewCloudintegrationtypesAWSTelemetryCollectionStrategyWithDefaults instantiates a new CloudintegrationtypesAWSTelemetryCollectionStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *CloudintegrationtypesAWSTelemetryCollectionStrategy) GetLogs() CloudintegrationtypesAWSLogsCollectionStrategy`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *CloudintegrationtypesAWSTelemetryCollectionStrategy) GetLogsOk() (*CloudintegrationtypesAWSLogsCollectionStrategy, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *CloudintegrationtypesAWSTelemetryCollectionStrategy) SetLogs(v CloudintegrationtypesAWSLogsCollectionStrategy)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *CloudintegrationtypesAWSTelemetryCollectionStrategy) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMetrics

`func (o *CloudintegrationtypesAWSTelemetryCollectionStrategy) GetMetrics() CloudintegrationtypesAWSMetricsCollectionStrategy`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CloudintegrationtypesAWSTelemetryCollectionStrategy) GetMetricsOk() (*CloudintegrationtypesAWSMetricsCollectionStrategy, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CloudintegrationtypesAWSTelemetryCollectionStrategy) SetMetrics(v CloudintegrationtypesAWSMetricsCollectionStrategy)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *CloudintegrationtypesAWSTelemetryCollectionStrategy) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetS3Buckets

`func (o *CloudintegrationtypesAWSTelemetryCollectionStrategy) GetS3Buckets() map[string][]string`

GetS3Buckets returns the S3Buckets field if non-nil, zero value otherwise.

### GetS3BucketsOk

`func (o *CloudintegrationtypesAWSTelemetryCollectionStrategy) GetS3BucketsOk() (*map[string][]string, bool)`

GetS3BucketsOk returns a tuple with the S3Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Buckets

`func (o *CloudintegrationtypesAWSTelemetryCollectionStrategy) SetS3Buckets(v map[string][]string)`

SetS3Buckets sets S3Buckets field to given value.

### HasS3Buckets

`func (o *CloudintegrationtypesAWSTelemetryCollectionStrategy) HasS3Buckets() bool`

HasS3Buckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


