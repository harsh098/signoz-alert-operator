# InframonitoringtypesPodRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedPodCount** | **int32** |  | 
**Meta** | **map[string]interface{}** |  | 
**PendingPodCount** | **int32** |  | 
**PodAge** | **int64** |  | 
**PodCPU** | **float64** |  | 
**PodCPULimit** | **float64** |  | 
**PodCPURequest** | **float64** |  | 
**PodMemory** | **float64** |  | 
**PodMemoryLimit** | **float64** |  | 
**PodMemoryRequest** | **float64** |  | 
**PodPhase** | [**InframonitoringtypesPodPhase**](InframonitoringtypesPodPhase.md) |  | 
**PodUID** | **string** |  | 
**RunningPodCount** | **int32** |  | 
**SucceededPodCount** | **int32** |  | 
**UnknownPodCount** | **int32** |  | 

## Methods

### NewInframonitoringtypesPodRecord

`func NewInframonitoringtypesPodRecord(failedPodCount int32, meta map[string]interface{}, pendingPodCount int32, podAge int64, podCPU float64, podCPULimit float64, podCPURequest float64, podMemory float64, podMemoryLimit float64, podMemoryRequest float64, podPhase InframonitoringtypesPodPhase, podUID string, runningPodCount int32, succeededPodCount int32, unknownPodCount int32, ) *InframonitoringtypesPodRecord`

NewInframonitoringtypesPodRecord instantiates a new InframonitoringtypesPodRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInframonitoringtypesPodRecordWithDefaults

`func NewInframonitoringtypesPodRecordWithDefaults() *InframonitoringtypesPodRecord`

NewInframonitoringtypesPodRecordWithDefaults instantiates a new InframonitoringtypesPodRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedPodCount

`func (o *InframonitoringtypesPodRecord) GetFailedPodCount() int32`

GetFailedPodCount returns the FailedPodCount field if non-nil, zero value otherwise.

### GetFailedPodCountOk

`func (o *InframonitoringtypesPodRecord) GetFailedPodCountOk() (*int32, bool)`

GetFailedPodCountOk returns a tuple with the FailedPodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedPodCount

`func (o *InframonitoringtypesPodRecord) SetFailedPodCount(v int32)`

SetFailedPodCount sets FailedPodCount field to given value.


### GetMeta

`func (o *InframonitoringtypesPodRecord) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InframonitoringtypesPodRecord) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InframonitoringtypesPodRecord) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### SetMetaNil

`func (o *InframonitoringtypesPodRecord) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *InframonitoringtypesPodRecord) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetPendingPodCount

`func (o *InframonitoringtypesPodRecord) GetPendingPodCount() int32`

GetPendingPodCount returns the PendingPodCount field if non-nil, zero value otherwise.

### GetPendingPodCountOk

`func (o *InframonitoringtypesPodRecord) GetPendingPodCountOk() (*int32, bool)`

GetPendingPodCountOk returns a tuple with the PendingPodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingPodCount

`func (o *InframonitoringtypesPodRecord) SetPendingPodCount(v int32)`

SetPendingPodCount sets PendingPodCount field to given value.


### GetPodAge

`func (o *InframonitoringtypesPodRecord) GetPodAge() int64`

GetPodAge returns the PodAge field if non-nil, zero value otherwise.

### GetPodAgeOk

`func (o *InframonitoringtypesPodRecord) GetPodAgeOk() (*int64, bool)`

GetPodAgeOk returns a tuple with the PodAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodAge

`func (o *InframonitoringtypesPodRecord) SetPodAge(v int64)`

SetPodAge sets PodAge field to given value.


### GetPodCPU

`func (o *InframonitoringtypesPodRecord) GetPodCPU() float64`

GetPodCPU returns the PodCPU field if non-nil, zero value otherwise.

### GetPodCPUOk

`func (o *InframonitoringtypesPodRecord) GetPodCPUOk() (*float64, bool)`

GetPodCPUOk returns a tuple with the PodCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCPU

`func (o *InframonitoringtypesPodRecord) SetPodCPU(v float64)`

SetPodCPU sets PodCPU field to given value.


### GetPodCPULimit

`func (o *InframonitoringtypesPodRecord) GetPodCPULimit() float64`

GetPodCPULimit returns the PodCPULimit field if non-nil, zero value otherwise.

### GetPodCPULimitOk

`func (o *InframonitoringtypesPodRecord) GetPodCPULimitOk() (*float64, bool)`

GetPodCPULimitOk returns a tuple with the PodCPULimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCPULimit

`func (o *InframonitoringtypesPodRecord) SetPodCPULimit(v float64)`

SetPodCPULimit sets PodCPULimit field to given value.


### GetPodCPURequest

`func (o *InframonitoringtypesPodRecord) GetPodCPURequest() float64`

GetPodCPURequest returns the PodCPURequest field if non-nil, zero value otherwise.

### GetPodCPURequestOk

`func (o *InframonitoringtypesPodRecord) GetPodCPURequestOk() (*float64, bool)`

GetPodCPURequestOk returns a tuple with the PodCPURequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCPURequest

`func (o *InframonitoringtypesPodRecord) SetPodCPURequest(v float64)`

SetPodCPURequest sets PodCPURequest field to given value.


### GetPodMemory

`func (o *InframonitoringtypesPodRecord) GetPodMemory() float64`

GetPodMemory returns the PodMemory field if non-nil, zero value otherwise.

### GetPodMemoryOk

`func (o *InframonitoringtypesPodRecord) GetPodMemoryOk() (*float64, bool)`

GetPodMemoryOk returns a tuple with the PodMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodMemory

`func (o *InframonitoringtypesPodRecord) SetPodMemory(v float64)`

SetPodMemory sets PodMemory field to given value.


### GetPodMemoryLimit

`func (o *InframonitoringtypesPodRecord) GetPodMemoryLimit() float64`

GetPodMemoryLimit returns the PodMemoryLimit field if non-nil, zero value otherwise.

### GetPodMemoryLimitOk

`func (o *InframonitoringtypesPodRecord) GetPodMemoryLimitOk() (*float64, bool)`

GetPodMemoryLimitOk returns a tuple with the PodMemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodMemoryLimit

`func (o *InframonitoringtypesPodRecord) SetPodMemoryLimit(v float64)`

SetPodMemoryLimit sets PodMemoryLimit field to given value.


### GetPodMemoryRequest

`func (o *InframonitoringtypesPodRecord) GetPodMemoryRequest() float64`

GetPodMemoryRequest returns the PodMemoryRequest field if non-nil, zero value otherwise.

### GetPodMemoryRequestOk

`func (o *InframonitoringtypesPodRecord) GetPodMemoryRequestOk() (*float64, bool)`

GetPodMemoryRequestOk returns a tuple with the PodMemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodMemoryRequest

`func (o *InframonitoringtypesPodRecord) SetPodMemoryRequest(v float64)`

SetPodMemoryRequest sets PodMemoryRequest field to given value.


### GetPodPhase

`func (o *InframonitoringtypesPodRecord) GetPodPhase() InframonitoringtypesPodPhase`

GetPodPhase returns the PodPhase field if non-nil, zero value otherwise.

### GetPodPhaseOk

`func (o *InframonitoringtypesPodRecord) GetPodPhaseOk() (*InframonitoringtypesPodPhase, bool)`

GetPodPhaseOk returns a tuple with the PodPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPhase

`func (o *InframonitoringtypesPodRecord) SetPodPhase(v InframonitoringtypesPodPhase)`

SetPodPhase sets PodPhase field to given value.


### GetPodUID

`func (o *InframonitoringtypesPodRecord) GetPodUID() string`

GetPodUID returns the PodUID field if non-nil, zero value otherwise.

### GetPodUIDOk

`func (o *InframonitoringtypesPodRecord) GetPodUIDOk() (*string, bool)`

GetPodUIDOk returns a tuple with the PodUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodUID

`func (o *InframonitoringtypesPodRecord) SetPodUID(v string)`

SetPodUID sets PodUID field to given value.


### GetRunningPodCount

`func (o *InframonitoringtypesPodRecord) GetRunningPodCount() int32`

GetRunningPodCount returns the RunningPodCount field if non-nil, zero value otherwise.

### GetRunningPodCountOk

`func (o *InframonitoringtypesPodRecord) GetRunningPodCountOk() (*int32, bool)`

GetRunningPodCountOk returns a tuple with the RunningPodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningPodCount

`func (o *InframonitoringtypesPodRecord) SetRunningPodCount(v int32)`

SetRunningPodCount sets RunningPodCount field to given value.


### GetSucceededPodCount

`func (o *InframonitoringtypesPodRecord) GetSucceededPodCount() int32`

GetSucceededPodCount returns the SucceededPodCount field if non-nil, zero value otherwise.

### GetSucceededPodCountOk

`func (o *InframonitoringtypesPodRecord) GetSucceededPodCountOk() (*int32, bool)`

GetSucceededPodCountOk returns a tuple with the SucceededPodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededPodCount

`func (o *InframonitoringtypesPodRecord) SetSucceededPodCount(v int32)`

SetSucceededPodCount sets SucceededPodCount field to given value.


### GetUnknownPodCount

`func (o *InframonitoringtypesPodRecord) GetUnknownPodCount() int32`

GetUnknownPodCount returns the UnknownPodCount field if non-nil, zero value otherwise.

### GetUnknownPodCountOk

`func (o *InframonitoringtypesPodRecord) GetUnknownPodCountOk() (*int32, bool)`

GetUnknownPodCountOk returns a tuple with the UnknownPodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownPodCount

`func (o *InframonitoringtypesPodRecord) SetUnknownPodCount(v int32)`

SetUnknownPodCount sets UnknownPodCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


