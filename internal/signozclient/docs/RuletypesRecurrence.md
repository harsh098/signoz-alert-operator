# RuletypesRecurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **string** |  | 
**EndTime** | Pointer to **NullableTime** |  | [optional] 
**RepeatOn** | Pointer to [**[]RuletypesRepeatOn**](RuletypesRepeatOn.md) |  | [optional] 
**RepeatType** | [**RuletypesRepeatType**](RuletypesRepeatType.md) |  | 
**StartTime** | **time.Time** |  | 

## Methods

### NewRuletypesRecurrence

`func NewRuletypesRecurrence(duration string, repeatType RuletypesRepeatType, startTime time.Time, ) *RuletypesRecurrence`

NewRuletypesRecurrence instantiates a new RuletypesRecurrence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuletypesRecurrenceWithDefaults

`func NewRuletypesRecurrenceWithDefaults() *RuletypesRecurrence`

NewRuletypesRecurrenceWithDefaults instantiates a new RuletypesRecurrence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *RuletypesRecurrence) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *RuletypesRecurrence) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *RuletypesRecurrence) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetEndTime

`func (o *RuletypesRecurrence) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RuletypesRecurrence) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RuletypesRecurrence) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RuletypesRecurrence) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *RuletypesRecurrence) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *RuletypesRecurrence) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetRepeatOn

`func (o *RuletypesRecurrence) GetRepeatOn() []RuletypesRepeatOn`

GetRepeatOn returns the RepeatOn field if non-nil, zero value otherwise.

### GetRepeatOnOk

`func (o *RuletypesRecurrence) GetRepeatOnOk() (*[]RuletypesRepeatOn, bool)`

GetRepeatOnOk returns a tuple with the RepeatOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatOn

`func (o *RuletypesRecurrence) SetRepeatOn(v []RuletypesRepeatOn)`

SetRepeatOn sets RepeatOn field to given value.

### HasRepeatOn

`func (o *RuletypesRecurrence) HasRepeatOn() bool`

HasRepeatOn returns a boolean if a field has been set.

### SetRepeatOnNil

`func (o *RuletypesRecurrence) SetRepeatOnNil(b bool)`

 SetRepeatOnNil sets the value for RepeatOn to be an explicit nil

### UnsetRepeatOn
`func (o *RuletypesRecurrence) UnsetRepeatOn()`

UnsetRepeatOn ensures that no value is present for RepeatOn, not even an explicit nil
### GetRepeatType

`func (o *RuletypesRecurrence) GetRepeatType() RuletypesRepeatType`

GetRepeatType returns the RepeatType field if non-nil, zero value otherwise.

### GetRepeatTypeOk

`func (o *RuletypesRecurrence) GetRepeatTypeOk() (*RuletypesRepeatType, bool)`

GetRepeatTypeOk returns a tuple with the RepeatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatType

`func (o *RuletypesRecurrence) SetRepeatType(v RuletypesRepeatType)`

SetRepeatType sets RepeatType field to given value.


### GetStartTime

`func (o *RuletypesRecurrence) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RuletypesRecurrence) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RuletypesRecurrence) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


