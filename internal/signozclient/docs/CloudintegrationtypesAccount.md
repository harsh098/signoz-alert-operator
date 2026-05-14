# CloudintegrationtypesAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentReport** | [**NullableCloudintegrationtypesAgentReport**](CloudintegrationtypesAgentReport.md) |  | 
**Config** | [**CloudintegrationtypesAccountConfig**](CloudintegrationtypesAccountConfig.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | **string** |  | 
**OrgId** | **string** |  | 
**Provider** | **string** |  | 
**ProviderAccountId** | **NullableString** |  | 
**RemovedAt** | **NullableTime** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCloudintegrationtypesAccount

`func NewCloudintegrationtypesAccount(agentReport NullableCloudintegrationtypesAgentReport, config CloudintegrationtypesAccountConfig, id string, orgId string, provider string, providerAccountId NullableString, removedAt NullableTime, ) *CloudintegrationtypesAccount`

NewCloudintegrationtypesAccount instantiates a new CloudintegrationtypesAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudintegrationtypesAccountWithDefaults

`func NewCloudintegrationtypesAccountWithDefaults() *CloudintegrationtypesAccount`

NewCloudintegrationtypesAccountWithDefaults instantiates a new CloudintegrationtypesAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentReport

`func (o *CloudintegrationtypesAccount) GetAgentReport() CloudintegrationtypesAgentReport`

GetAgentReport returns the AgentReport field if non-nil, zero value otherwise.

### GetAgentReportOk

`func (o *CloudintegrationtypesAccount) GetAgentReportOk() (*CloudintegrationtypesAgentReport, bool)`

GetAgentReportOk returns a tuple with the AgentReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentReport

`func (o *CloudintegrationtypesAccount) SetAgentReport(v CloudintegrationtypesAgentReport)`

SetAgentReport sets AgentReport field to given value.


### SetAgentReportNil

`func (o *CloudintegrationtypesAccount) SetAgentReportNil(b bool)`

 SetAgentReportNil sets the value for AgentReport to be an explicit nil

### UnsetAgentReport
`func (o *CloudintegrationtypesAccount) UnsetAgentReport()`

UnsetAgentReport ensures that no value is present for AgentReport, not even an explicit nil
### GetConfig

`func (o *CloudintegrationtypesAccount) GetConfig() CloudintegrationtypesAccountConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CloudintegrationtypesAccount) GetConfigOk() (*CloudintegrationtypesAccountConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CloudintegrationtypesAccount) SetConfig(v CloudintegrationtypesAccountConfig)`

SetConfig sets Config field to given value.


### GetCreatedAt

`func (o *CloudintegrationtypesAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudintegrationtypesAccount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudintegrationtypesAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudintegrationtypesAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CloudintegrationtypesAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudintegrationtypesAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudintegrationtypesAccount) SetId(v string)`

SetId sets Id field to given value.


### GetOrgId

`func (o *CloudintegrationtypesAccount) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CloudintegrationtypesAccount) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CloudintegrationtypesAccount) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetProvider

`func (o *CloudintegrationtypesAccount) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudintegrationtypesAccount) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudintegrationtypesAccount) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProviderAccountId

`func (o *CloudintegrationtypesAccount) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *CloudintegrationtypesAccount) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *CloudintegrationtypesAccount) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.


### SetProviderAccountIdNil

`func (o *CloudintegrationtypesAccount) SetProviderAccountIdNil(b bool)`

 SetProviderAccountIdNil sets the value for ProviderAccountId to be an explicit nil

### UnsetProviderAccountId
`func (o *CloudintegrationtypesAccount) UnsetProviderAccountId()`

UnsetProviderAccountId ensures that no value is present for ProviderAccountId, not even an explicit nil
### GetRemovedAt

`func (o *CloudintegrationtypesAccount) GetRemovedAt() time.Time`

GetRemovedAt returns the RemovedAt field if non-nil, zero value otherwise.

### GetRemovedAtOk

`func (o *CloudintegrationtypesAccount) GetRemovedAtOk() (*time.Time, bool)`

GetRemovedAtOk returns a tuple with the RemovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedAt

`func (o *CloudintegrationtypesAccount) SetRemovedAt(v time.Time)`

SetRemovedAt sets RemovedAt field to given value.


### SetRemovedAtNil

`func (o *CloudintegrationtypesAccount) SetRemovedAtNil(b bool)`

 SetRemovedAtNil sets the value for RemovedAt to be an explicit nil

### UnsetRemovedAt
`func (o *CloudintegrationtypesAccount) UnsetRemovedAt()`

UnsetRemovedAt ensures that no value is present for RemovedAt, not even an explicit nil
### GetUpdatedAt

`func (o *CloudintegrationtypesAccount) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudintegrationtypesAccount) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudintegrationtypesAccount) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudintegrationtypesAccount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


