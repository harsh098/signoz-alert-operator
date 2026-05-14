# CloudintegrationtypesGettableAgentCheckIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**CloudAccountId** | **string** |  | 
**CloudIntegrationId** | **string** |  | 
**IntegrationConfig** | [**NullableCloudintegrationtypesIntegrationConfig**](CloudintegrationtypesIntegrationConfig.md) |  | 
**IntegrationConfig** | [**CloudintegrationtypesProviderIntegrationConfig**](CloudintegrationtypesProviderIntegrationConfig.md) |  | 
**ProviderAccountId** | **string** |  | 
**RemovedAt** | **NullableTime** |  | 
**RemovedAt** | **NullableTime** |  | 

## Methods

### NewCloudintegrationtypesGettableAgentCheckIn

`func NewCloudintegrationtypesGettableAgentCheckIn(accountId string, cloudAccountId string, cloudIntegrationId string, integrationConfig NullableCloudintegrationtypesIntegrationConfig, integrationConfig CloudintegrationtypesProviderIntegrationConfig, providerAccountId string, removedAt NullableTime, removedAt NullableTime, ) *CloudintegrationtypesGettableAgentCheckIn`

NewCloudintegrationtypesGettableAgentCheckIn instantiates a new CloudintegrationtypesGettableAgentCheckIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudintegrationtypesGettableAgentCheckInWithDefaults

`func NewCloudintegrationtypesGettableAgentCheckInWithDefaults() *CloudintegrationtypesGettableAgentCheckIn`

NewCloudintegrationtypesGettableAgentCheckInWithDefaults instantiates a new CloudintegrationtypesGettableAgentCheckIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CloudintegrationtypesGettableAgentCheckIn) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCloudAccountId

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetCloudAccountId() string`

GetCloudAccountId returns the CloudAccountId field if non-nil, zero value otherwise.

### GetCloudAccountIdOk

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetCloudAccountIdOk() (*string, bool)`

GetCloudAccountIdOk returns a tuple with the CloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountId

`func (o *CloudintegrationtypesGettableAgentCheckIn) SetCloudAccountId(v string)`

SetCloudAccountId sets CloudAccountId field to given value.


### GetCloudIntegrationId

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetCloudIntegrationId() string`

GetCloudIntegrationId returns the CloudIntegrationId field if non-nil, zero value otherwise.

### GetCloudIntegrationIdOk

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetCloudIntegrationIdOk() (*string, bool)`

GetCloudIntegrationIdOk returns a tuple with the CloudIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudIntegrationId

`func (o *CloudintegrationtypesGettableAgentCheckIn) SetCloudIntegrationId(v string)`

SetCloudIntegrationId sets CloudIntegrationId field to given value.


### GetIntegrationConfig

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetIntegrationConfig() CloudintegrationtypesIntegrationConfig`

GetIntegrationConfig returns the IntegrationConfig field if non-nil, zero value otherwise.

### GetIntegrationConfigOk

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetIntegrationConfigOk() (*CloudintegrationtypesIntegrationConfig, bool)`

GetIntegrationConfigOk returns a tuple with the IntegrationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationConfig

`func (o *CloudintegrationtypesGettableAgentCheckIn) SetIntegrationConfig(v CloudintegrationtypesIntegrationConfig)`

SetIntegrationConfig sets IntegrationConfig field to given value.


### SetIntegrationConfigNil

`func (o *CloudintegrationtypesGettableAgentCheckIn) SetIntegrationConfigNil(b bool)`

 SetIntegrationConfigNil sets the value for IntegrationConfig to be an explicit nil

### UnsetIntegrationConfig
`func (o *CloudintegrationtypesGettableAgentCheckIn) UnsetIntegrationConfig()`

UnsetIntegrationConfig ensures that no value is present for IntegrationConfig, not even an explicit nil
### GetIntegrationConfig

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetIntegrationConfig() CloudintegrationtypesProviderIntegrationConfig`

GetIntegrationConfig returns the IntegrationConfig field if non-nil, zero value otherwise.

### GetIntegrationConfigOk

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetIntegrationConfigOk() (*CloudintegrationtypesProviderIntegrationConfig, bool)`

GetIntegrationConfigOk returns a tuple with the IntegrationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationConfig

`func (o *CloudintegrationtypesGettableAgentCheckIn) SetIntegrationConfig(v CloudintegrationtypesProviderIntegrationConfig)`

SetIntegrationConfig sets IntegrationConfig field to given value.


### GetProviderAccountId

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *CloudintegrationtypesGettableAgentCheckIn) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.


### GetRemovedAt

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetRemovedAt() time.Time`

GetRemovedAt returns the RemovedAt field if non-nil, zero value otherwise.

### GetRemovedAtOk

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetRemovedAtOk() (*time.Time, bool)`

GetRemovedAtOk returns a tuple with the RemovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedAt

`func (o *CloudintegrationtypesGettableAgentCheckIn) SetRemovedAt(v time.Time)`

SetRemovedAt sets RemovedAt field to given value.


### SetRemovedAtNil

`func (o *CloudintegrationtypesGettableAgentCheckIn) SetRemovedAtNil(b bool)`

 SetRemovedAtNil sets the value for RemovedAt to be an explicit nil

### UnsetRemovedAt
`func (o *CloudintegrationtypesGettableAgentCheckIn) UnsetRemovedAt()`

UnsetRemovedAt ensures that no value is present for RemovedAt, not even an explicit nil
### GetRemovedAt

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetRemovedAt() time.Time`

GetRemovedAt returns the RemovedAt field if non-nil, zero value otherwise.

### GetRemovedAtOk

`func (o *CloudintegrationtypesGettableAgentCheckIn) GetRemovedAtOk() (*time.Time, bool)`

GetRemovedAtOk returns a tuple with the RemovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedAt

`func (o *CloudintegrationtypesGettableAgentCheckIn) SetRemovedAt(v time.Time)`

SetRemovedAt sets RemovedAt field to given value.


### SetRemovedAtNil

`func (o *CloudintegrationtypesGettableAgentCheckIn) SetRemovedAtNil(b bool)`

 SetRemovedAtNil sets the value for RemovedAt to be an explicit nil

### UnsetRemovedAt
`func (o *CloudintegrationtypesGettableAgentCheckIn) UnsetRemovedAt()`

UnsetRemovedAt ensures that no value is present for RemovedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


