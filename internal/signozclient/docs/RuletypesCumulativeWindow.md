# RuletypesCumulativeWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | **string** |  | 
**Schedule** | [**RuletypesCumulativeSchedule**](RuletypesCumulativeSchedule.md) |  | 
**Timezone** | **string** |  | 

## Methods

### NewRuletypesCumulativeWindow

`func NewRuletypesCumulativeWindow(frequency string, schedule RuletypesCumulativeSchedule, timezone string, ) *RuletypesCumulativeWindow`

NewRuletypesCumulativeWindow instantiates a new RuletypesCumulativeWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuletypesCumulativeWindowWithDefaults

`func NewRuletypesCumulativeWindowWithDefaults() *RuletypesCumulativeWindow`

NewRuletypesCumulativeWindowWithDefaults instantiates a new RuletypesCumulativeWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *RuletypesCumulativeWindow) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *RuletypesCumulativeWindow) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *RuletypesCumulativeWindow) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetSchedule

`func (o *RuletypesCumulativeWindow) GetSchedule() RuletypesCumulativeSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *RuletypesCumulativeWindow) GetScheduleOk() (*RuletypesCumulativeSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *RuletypesCumulativeWindow) SetSchedule(v RuletypesCumulativeSchedule)`

SetSchedule sets Schedule field to given value.


### GetTimezone

`func (o *RuletypesCumulativeWindow) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *RuletypesCumulativeWindow) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *RuletypesCumulativeWindow) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


