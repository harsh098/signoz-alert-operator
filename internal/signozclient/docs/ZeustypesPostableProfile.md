# ZeustypesPostableProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingObservabilityTool** | **string** |  | 
**HasExistingObservabilityTool** | **bool** |  | 
**LogsScalePerDayInGb** | **int64** |  | 
**NumberOfHosts** | **int64** |  | 
**NumberOfServices** | **int64** |  | 
**ReasonsForInterestInSignoz** | **[]string** |  | 
**TimelineForMigratingToSignoz** | **string** |  | 
**UsesOtel** | **bool** |  | 
**WhereDidYouDiscoverSignoz** | **string** |  | 

## Methods

### NewZeustypesPostableProfile

`func NewZeustypesPostableProfile(existingObservabilityTool string, hasExistingObservabilityTool bool, logsScalePerDayInGb int64, numberOfHosts int64, numberOfServices int64, reasonsForInterestInSignoz []string, timelineForMigratingToSignoz string, usesOtel bool, whereDidYouDiscoverSignoz string, ) *ZeustypesPostableProfile`

NewZeustypesPostableProfile instantiates a new ZeustypesPostableProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZeustypesPostableProfileWithDefaults

`func NewZeustypesPostableProfileWithDefaults() *ZeustypesPostableProfile`

NewZeustypesPostableProfileWithDefaults instantiates a new ZeustypesPostableProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistingObservabilityTool

`func (o *ZeustypesPostableProfile) GetExistingObservabilityTool() string`

GetExistingObservabilityTool returns the ExistingObservabilityTool field if non-nil, zero value otherwise.

### GetExistingObservabilityToolOk

`func (o *ZeustypesPostableProfile) GetExistingObservabilityToolOk() (*string, bool)`

GetExistingObservabilityToolOk returns a tuple with the ExistingObservabilityTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingObservabilityTool

`func (o *ZeustypesPostableProfile) SetExistingObservabilityTool(v string)`

SetExistingObservabilityTool sets ExistingObservabilityTool field to given value.


### GetHasExistingObservabilityTool

`func (o *ZeustypesPostableProfile) GetHasExistingObservabilityTool() bool`

GetHasExistingObservabilityTool returns the HasExistingObservabilityTool field if non-nil, zero value otherwise.

### GetHasExistingObservabilityToolOk

`func (o *ZeustypesPostableProfile) GetHasExistingObservabilityToolOk() (*bool, bool)`

GetHasExistingObservabilityToolOk returns a tuple with the HasExistingObservabilityTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExistingObservabilityTool

`func (o *ZeustypesPostableProfile) SetHasExistingObservabilityTool(v bool)`

SetHasExistingObservabilityTool sets HasExistingObservabilityTool field to given value.


### GetLogsScalePerDayInGb

`func (o *ZeustypesPostableProfile) GetLogsScalePerDayInGb() int64`

GetLogsScalePerDayInGb returns the LogsScalePerDayInGb field if non-nil, zero value otherwise.

### GetLogsScalePerDayInGbOk

`func (o *ZeustypesPostableProfile) GetLogsScalePerDayInGbOk() (*int64, bool)`

GetLogsScalePerDayInGbOk returns a tuple with the LogsScalePerDayInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsScalePerDayInGb

`func (o *ZeustypesPostableProfile) SetLogsScalePerDayInGb(v int64)`

SetLogsScalePerDayInGb sets LogsScalePerDayInGb field to given value.


### GetNumberOfHosts

`func (o *ZeustypesPostableProfile) GetNumberOfHosts() int64`

GetNumberOfHosts returns the NumberOfHosts field if non-nil, zero value otherwise.

### GetNumberOfHostsOk

`func (o *ZeustypesPostableProfile) GetNumberOfHostsOk() (*int64, bool)`

GetNumberOfHostsOk returns a tuple with the NumberOfHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfHosts

`func (o *ZeustypesPostableProfile) SetNumberOfHosts(v int64)`

SetNumberOfHosts sets NumberOfHosts field to given value.


### GetNumberOfServices

`func (o *ZeustypesPostableProfile) GetNumberOfServices() int64`

GetNumberOfServices returns the NumberOfServices field if non-nil, zero value otherwise.

### GetNumberOfServicesOk

`func (o *ZeustypesPostableProfile) GetNumberOfServicesOk() (*int64, bool)`

GetNumberOfServicesOk returns a tuple with the NumberOfServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfServices

`func (o *ZeustypesPostableProfile) SetNumberOfServices(v int64)`

SetNumberOfServices sets NumberOfServices field to given value.


### GetReasonsForInterestInSignoz

`func (o *ZeustypesPostableProfile) GetReasonsForInterestInSignoz() []string`

GetReasonsForInterestInSignoz returns the ReasonsForInterestInSignoz field if non-nil, zero value otherwise.

### GetReasonsForInterestInSignozOk

`func (o *ZeustypesPostableProfile) GetReasonsForInterestInSignozOk() (*[]string, bool)`

GetReasonsForInterestInSignozOk returns a tuple with the ReasonsForInterestInSignoz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonsForInterestInSignoz

`func (o *ZeustypesPostableProfile) SetReasonsForInterestInSignoz(v []string)`

SetReasonsForInterestInSignoz sets ReasonsForInterestInSignoz field to given value.


### SetReasonsForInterestInSignozNil

`func (o *ZeustypesPostableProfile) SetReasonsForInterestInSignozNil(b bool)`

 SetReasonsForInterestInSignozNil sets the value for ReasonsForInterestInSignoz to be an explicit nil

### UnsetReasonsForInterestInSignoz
`func (o *ZeustypesPostableProfile) UnsetReasonsForInterestInSignoz()`

UnsetReasonsForInterestInSignoz ensures that no value is present for ReasonsForInterestInSignoz, not even an explicit nil
### GetTimelineForMigratingToSignoz

`func (o *ZeustypesPostableProfile) GetTimelineForMigratingToSignoz() string`

GetTimelineForMigratingToSignoz returns the TimelineForMigratingToSignoz field if non-nil, zero value otherwise.

### GetTimelineForMigratingToSignozOk

`func (o *ZeustypesPostableProfile) GetTimelineForMigratingToSignozOk() (*string, bool)`

GetTimelineForMigratingToSignozOk returns a tuple with the TimelineForMigratingToSignoz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineForMigratingToSignoz

`func (o *ZeustypesPostableProfile) SetTimelineForMigratingToSignoz(v string)`

SetTimelineForMigratingToSignoz sets TimelineForMigratingToSignoz field to given value.


### GetUsesOtel

`func (o *ZeustypesPostableProfile) GetUsesOtel() bool`

GetUsesOtel returns the UsesOtel field if non-nil, zero value otherwise.

### GetUsesOtelOk

`func (o *ZeustypesPostableProfile) GetUsesOtelOk() (*bool, bool)`

GetUsesOtelOk returns a tuple with the UsesOtel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesOtel

`func (o *ZeustypesPostableProfile) SetUsesOtel(v bool)`

SetUsesOtel sets UsesOtel field to given value.


### GetWhereDidYouDiscoverSignoz

`func (o *ZeustypesPostableProfile) GetWhereDidYouDiscoverSignoz() string`

GetWhereDidYouDiscoverSignoz returns the WhereDidYouDiscoverSignoz field if non-nil, zero value otherwise.

### GetWhereDidYouDiscoverSignozOk

`func (o *ZeustypesPostableProfile) GetWhereDidYouDiscoverSignozOk() (*string, bool)`

GetWhereDidYouDiscoverSignozOk returns a tuple with the WhereDidYouDiscoverSignoz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhereDidYouDiscoverSignoz

`func (o *ZeustypesPostableProfile) SetWhereDidYouDiscoverSignoz(v string)`

SetWhereDidYouDiscoverSignoz sets WhereDidYouDiscoverSignoz field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


