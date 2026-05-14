# CloudintegrationtypesOldAWSCollectionStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsLogs** | Pointer to [**CloudintegrationtypesOldAWSLogsStrategy**](CloudintegrationtypesOldAWSLogsStrategy.md) |  | [optional] 
**AwsMetrics** | Pointer to [**CloudintegrationtypesOldAWSMetricsStrategy**](CloudintegrationtypesOldAWSMetricsStrategy.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**S3Buckets** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewCloudintegrationtypesOldAWSCollectionStrategy

`func NewCloudintegrationtypesOldAWSCollectionStrategy() *CloudintegrationtypesOldAWSCollectionStrategy`

NewCloudintegrationtypesOldAWSCollectionStrategy instantiates a new CloudintegrationtypesOldAWSCollectionStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudintegrationtypesOldAWSCollectionStrategyWithDefaults

`func NewCloudintegrationtypesOldAWSCollectionStrategyWithDefaults() *CloudintegrationtypesOldAWSCollectionStrategy`

NewCloudintegrationtypesOldAWSCollectionStrategyWithDefaults instantiates a new CloudintegrationtypesOldAWSCollectionStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsLogs

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) GetAwsLogs() CloudintegrationtypesOldAWSLogsStrategy`

GetAwsLogs returns the AwsLogs field if non-nil, zero value otherwise.

### GetAwsLogsOk

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) GetAwsLogsOk() (*CloudintegrationtypesOldAWSLogsStrategy, bool)`

GetAwsLogsOk returns a tuple with the AwsLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsLogs

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) SetAwsLogs(v CloudintegrationtypesOldAWSLogsStrategy)`

SetAwsLogs sets AwsLogs field to given value.

### HasAwsLogs

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) HasAwsLogs() bool`

HasAwsLogs returns a boolean if a field has been set.

### GetAwsMetrics

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) GetAwsMetrics() CloudintegrationtypesOldAWSMetricsStrategy`

GetAwsMetrics returns the AwsMetrics field if non-nil, zero value otherwise.

### GetAwsMetricsOk

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) GetAwsMetricsOk() (*CloudintegrationtypesOldAWSMetricsStrategy, bool)`

GetAwsMetricsOk returns a tuple with the AwsMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsMetrics

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) SetAwsMetrics(v CloudintegrationtypesOldAWSMetricsStrategy)`

SetAwsMetrics sets AwsMetrics field to given value.

### HasAwsMetrics

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) HasAwsMetrics() bool`

HasAwsMetrics returns a boolean if a field has been set.

### GetProvider

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetS3Buckets

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) GetS3Buckets() map[string][]string`

GetS3Buckets returns the S3Buckets field if non-nil, zero value otherwise.

### GetS3BucketsOk

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) GetS3BucketsOk() (*map[string][]string, bool)`

GetS3BucketsOk returns a tuple with the S3Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Buckets

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) SetS3Buckets(v map[string][]string)`

SetS3Buckets sets S3Buckets field to given value.

### HasS3Buckets

`func (o *CloudintegrationtypesOldAWSCollectionStrategy) HasS3Buckets() bool`

HasS3Buckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


