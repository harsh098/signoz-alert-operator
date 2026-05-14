# RuletypesPostablePlannedMaintenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertIds** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Schedule** | [**RuletypesSchedule**](RuletypesSchedule.md) |  | 

## Methods

### NewRuletypesPostablePlannedMaintenance

`func NewRuletypesPostablePlannedMaintenance(name string, schedule RuletypesSchedule, ) *RuletypesPostablePlannedMaintenance`

NewRuletypesPostablePlannedMaintenance instantiates a new RuletypesPostablePlannedMaintenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuletypesPostablePlannedMaintenanceWithDefaults

`func NewRuletypesPostablePlannedMaintenanceWithDefaults() *RuletypesPostablePlannedMaintenance`

NewRuletypesPostablePlannedMaintenanceWithDefaults instantiates a new RuletypesPostablePlannedMaintenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertIds

`func (o *RuletypesPostablePlannedMaintenance) GetAlertIds() []string`

GetAlertIds returns the AlertIds field if non-nil, zero value otherwise.

### GetAlertIdsOk

`func (o *RuletypesPostablePlannedMaintenance) GetAlertIdsOk() (*[]string, bool)`

GetAlertIdsOk returns a tuple with the AlertIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertIds

`func (o *RuletypesPostablePlannedMaintenance) SetAlertIds(v []string)`

SetAlertIds sets AlertIds field to given value.

### HasAlertIds

`func (o *RuletypesPostablePlannedMaintenance) HasAlertIds() bool`

HasAlertIds returns a boolean if a field has been set.

### SetAlertIdsNil

`func (o *RuletypesPostablePlannedMaintenance) SetAlertIdsNil(b bool)`

 SetAlertIdsNil sets the value for AlertIds to be an explicit nil

### UnsetAlertIds
`func (o *RuletypesPostablePlannedMaintenance) UnsetAlertIds()`

UnsetAlertIds ensures that no value is present for AlertIds, not even an explicit nil
### GetDescription

`func (o *RuletypesPostablePlannedMaintenance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuletypesPostablePlannedMaintenance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuletypesPostablePlannedMaintenance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuletypesPostablePlannedMaintenance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *RuletypesPostablePlannedMaintenance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuletypesPostablePlannedMaintenance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuletypesPostablePlannedMaintenance) SetName(v string)`

SetName sets Name field to given value.


### GetSchedule

`func (o *RuletypesPostablePlannedMaintenance) GetSchedule() RuletypesSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *RuletypesPostablePlannedMaintenance) GetScheduleOk() (*RuletypesSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *RuletypesPostablePlannedMaintenance) SetSchedule(v RuletypesSchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


