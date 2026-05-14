# RuletypesSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **time.Time** |  | [optional] 
**Recurrence** | Pointer to [**RuletypesRecurrence**](RuletypesRecurrence.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**Timezone** | **string** |  | 

## Methods

### NewRuletypesSchedule

`func NewRuletypesSchedule(timezone string, ) *RuletypesSchedule`

NewRuletypesSchedule instantiates a new RuletypesSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuletypesScheduleWithDefaults

`func NewRuletypesScheduleWithDefaults() *RuletypesSchedule`

NewRuletypesScheduleWithDefaults instantiates a new RuletypesSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *RuletypesSchedule) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RuletypesSchedule) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RuletypesSchedule) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RuletypesSchedule) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetRecurrence

`func (o *RuletypesSchedule) GetRecurrence() RuletypesRecurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *RuletypesSchedule) GetRecurrenceOk() (*RuletypesRecurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *RuletypesSchedule) SetRecurrence(v RuletypesRecurrence)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *RuletypesSchedule) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### GetStartTime

`func (o *RuletypesSchedule) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RuletypesSchedule) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RuletypesSchedule) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RuletypesSchedule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTimezone

`func (o *RuletypesSchedule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *RuletypesSchedule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *RuletypesSchedule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


