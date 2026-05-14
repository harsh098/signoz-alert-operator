# CloudintegrationtypesService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | [**CloudintegrationtypesAssets**](CloudintegrationtypesAssets.md) |  | 
**CloudIntegrationService** | [**NullableCloudintegrationtypesCloudIntegrationService**](CloudintegrationtypesCloudIntegrationService.md) |  | 
**DataCollected** | [**CloudintegrationtypesDataCollected**](CloudintegrationtypesDataCollected.md) |  | 
**Icon** | **string** |  | 
**Id** | **string** |  | 
**Overview** | **string** |  | 
**SupportedSignals** | [**CloudintegrationtypesSupportedSignals**](CloudintegrationtypesSupportedSignals.md) |  | 
**TelemetryCollectionStrategy** | [**CloudintegrationtypesTelemetryCollectionStrategy**](CloudintegrationtypesTelemetryCollectionStrategy.md) |  | 
**Title** | **string** |  | 

## Methods

### NewCloudintegrationtypesService

`func NewCloudintegrationtypesService(assets CloudintegrationtypesAssets, cloudIntegrationService NullableCloudintegrationtypesCloudIntegrationService, dataCollected CloudintegrationtypesDataCollected, icon string, id string, overview string, supportedSignals CloudintegrationtypesSupportedSignals, telemetryCollectionStrategy CloudintegrationtypesTelemetryCollectionStrategy, title string, ) *CloudintegrationtypesService`

NewCloudintegrationtypesService instantiates a new CloudintegrationtypesService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudintegrationtypesServiceWithDefaults

`func NewCloudintegrationtypesServiceWithDefaults() *CloudintegrationtypesService`

NewCloudintegrationtypesServiceWithDefaults instantiates a new CloudintegrationtypesService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *CloudintegrationtypesService) GetAssets() CloudintegrationtypesAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *CloudintegrationtypesService) GetAssetsOk() (*CloudintegrationtypesAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *CloudintegrationtypesService) SetAssets(v CloudintegrationtypesAssets)`

SetAssets sets Assets field to given value.


### GetCloudIntegrationService

`func (o *CloudintegrationtypesService) GetCloudIntegrationService() CloudintegrationtypesCloudIntegrationService`

GetCloudIntegrationService returns the CloudIntegrationService field if non-nil, zero value otherwise.

### GetCloudIntegrationServiceOk

`func (o *CloudintegrationtypesService) GetCloudIntegrationServiceOk() (*CloudintegrationtypesCloudIntegrationService, bool)`

GetCloudIntegrationServiceOk returns a tuple with the CloudIntegrationService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudIntegrationService

`func (o *CloudintegrationtypesService) SetCloudIntegrationService(v CloudintegrationtypesCloudIntegrationService)`

SetCloudIntegrationService sets CloudIntegrationService field to given value.


### SetCloudIntegrationServiceNil

`func (o *CloudintegrationtypesService) SetCloudIntegrationServiceNil(b bool)`

 SetCloudIntegrationServiceNil sets the value for CloudIntegrationService to be an explicit nil

### UnsetCloudIntegrationService
`func (o *CloudintegrationtypesService) UnsetCloudIntegrationService()`

UnsetCloudIntegrationService ensures that no value is present for CloudIntegrationService, not even an explicit nil
### GetDataCollected

`func (o *CloudintegrationtypesService) GetDataCollected() CloudintegrationtypesDataCollected`

GetDataCollected returns the DataCollected field if non-nil, zero value otherwise.

### GetDataCollectedOk

`func (o *CloudintegrationtypesService) GetDataCollectedOk() (*CloudintegrationtypesDataCollected, bool)`

GetDataCollectedOk returns a tuple with the DataCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollected

`func (o *CloudintegrationtypesService) SetDataCollected(v CloudintegrationtypesDataCollected)`

SetDataCollected sets DataCollected field to given value.


### GetIcon

`func (o *CloudintegrationtypesService) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CloudintegrationtypesService) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CloudintegrationtypesService) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetId

`func (o *CloudintegrationtypesService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudintegrationtypesService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudintegrationtypesService) SetId(v string)`

SetId sets Id field to given value.


### GetOverview

`func (o *CloudintegrationtypesService) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *CloudintegrationtypesService) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *CloudintegrationtypesService) SetOverview(v string)`

SetOverview sets Overview field to given value.


### GetSupportedSignals

`func (o *CloudintegrationtypesService) GetSupportedSignals() CloudintegrationtypesSupportedSignals`

GetSupportedSignals returns the SupportedSignals field if non-nil, zero value otherwise.

### GetSupportedSignalsOk

`func (o *CloudintegrationtypesService) GetSupportedSignalsOk() (*CloudintegrationtypesSupportedSignals, bool)`

GetSupportedSignalsOk returns a tuple with the SupportedSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSignals

`func (o *CloudintegrationtypesService) SetSupportedSignals(v CloudintegrationtypesSupportedSignals)`

SetSupportedSignals sets SupportedSignals field to given value.


### GetTelemetryCollectionStrategy

`func (o *CloudintegrationtypesService) GetTelemetryCollectionStrategy() CloudintegrationtypesTelemetryCollectionStrategy`

GetTelemetryCollectionStrategy returns the TelemetryCollectionStrategy field if non-nil, zero value otherwise.

### GetTelemetryCollectionStrategyOk

`func (o *CloudintegrationtypesService) GetTelemetryCollectionStrategyOk() (*CloudintegrationtypesTelemetryCollectionStrategy, bool)`

GetTelemetryCollectionStrategyOk returns a tuple with the TelemetryCollectionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryCollectionStrategy

`func (o *CloudintegrationtypesService) SetTelemetryCollectionStrategy(v CloudintegrationtypesTelemetryCollectionStrategy)`

SetTelemetryCollectionStrategy sets TelemetryCollectionStrategy field to given value.


### GetTitle

`func (o *CloudintegrationtypesService) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudintegrationtypesService) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudintegrationtypesService) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


