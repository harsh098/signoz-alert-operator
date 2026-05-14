# RuletypesPlannedMaintenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertIds** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Kind** | [**RuletypesMaintenanceKind**](RuletypesMaintenanceKind.md) |  | 
**Name** | **string** |  | 
**Schedule** | [**RuletypesSchedule**](RuletypesSchedule.md) |  | 
**Status** | [**RuletypesMaintenanceStatus**](RuletypesMaintenanceStatus.md) |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewRuletypesPlannedMaintenance

`func NewRuletypesPlannedMaintenance(id string, kind RuletypesMaintenanceKind, name string, schedule RuletypesSchedule, status RuletypesMaintenanceStatus, ) *RuletypesPlannedMaintenance`

NewRuletypesPlannedMaintenance instantiates a new RuletypesPlannedMaintenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuletypesPlannedMaintenanceWithDefaults

`func NewRuletypesPlannedMaintenanceWithDefaults() *RuletypesPlannedMaintenance`

NewRuletypesPlannedMaintenanceWithDefaults instantiates a new RuletypesPlannedMaintenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertIds

`func (o *RuletypesPlannedMaintenance) GetAlertIds() []string`

GetAlertIds returns the AlertIds field if non-nil, zero value otherwise.

### GetAlertIdsOk

`func (o *RuletypesPlannedMaintenance) GetAlertIdsOk() (*[]string, bool)`

GetAlertIdsOk returns a tuple with the AlertIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertIds

`func (o *RuletypesPlannedMaintenance) SetAlertIds(v []string)`

SetAlertIds sets AlertIds field to given value.

### HasAlertIds

`func (o *RuletypesPlannedMaintenance) HasAlertIds() bool`

HasAlertIds returns a boolean if a field has been set.

### SetAlertIdsNil

`func (o *RuletypesPlannedMaintenance) SetAlertIdsNil(b bool)`

 SetAlertIdsNil sets the value for AlertIds to be an explicit nil

### UnsetAlertIds
`func (o *RuletypesPlannedMaintenance) UnsetAlertIds()`

UnsetAlertIds ensures that no value is present for AlertIds, not even an explicit nil
### GetCreatedAt

`func (o *RuletypesPlannedMaintenance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RuletypesPlannedMaintenance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RuletypesPlannedMaintenance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RuletypesPlannedMaintenance) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RuletypesPlannedMaintenance) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RuletypesPlannedMaintenance) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RuletypesPlannedMaintenance) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RuletypesPlannedMaintenance) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *RuletypesPlannedMaintenance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuletypesPlannedMaintenance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuletypesPlannedMaintenance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuletypesPlannedMaintenance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *RuletypesPlannedMaintenance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuletypesPlannedMaintenance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuletypesPlannedMaintenance) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *RuletypesPlannedMaintenance) GetKind() RuletypesMaintenanceKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RuletypesPlannedMaintenance) GetKindOk() (*RuletypesMaintenanceKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RuletypesPlannedMaintenance) SetKind(v RuletypesMaintenanceKind)`

SetKind sets Kind field to given value.


### GetName

`func (o *RuletypesPlannedMaintenance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuletypesPlannedMaintenance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuletypesPlannedMaintenance) SetName(v string)`

SetName sets Name field to given value.


### GetSchedule

`func (o *RuletypesPlannedMaintenance) GetSchedule() RuletypesSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *RuletypesPlannedMaintenance) GetScheduleOk() (*RuletypesSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *RuletypesPlannedMaintenance) SetSchedule(v RuletypesSchedule)`

SetSchedule sets Schedule field to given value.


### GetStatus

`func (o *RuletypesPlannedMaintenance) GetStatus() RuletypesMaintenanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RuletypesPlannedMaintenance) GetStatusOk() (*RuletypesMaintenanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RuletypesPlannedMaintenance) SetStatus(v RuletypesMaintenanceStatus)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *RuletypesPlannedMaintenance) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RuletypesPlannedMaintenance) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RuletypesPlannedMaintenance) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RuletypesPlannedMaintenance) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RuletypesPlannedMaintenance) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RuletypesPlannedMaintenance) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RuletypesPlannedMaintenance) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RuletypesPlannedMaintenance) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


