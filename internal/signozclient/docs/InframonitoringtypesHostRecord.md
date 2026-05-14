# InframonitoringtypesHostRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveHostCount** | **int32** |  | 
**Cpu** | **float64** |  | 
**DiskUsage** | **float64** |  | 
**HostName** | **string** |  | 
**InactiveHostCount** | **int32** |  | 
**Load15** | **float64** |  | 
**Memory** | **float64** |  | 
**Meta** | **map[string]interface{}** |  | 
**Status** | [**InframonitoringtypesHostStatus**](InframonitoringtypesHostStatus.md) |  | 
**Wait** | **float64** |  | 

## Methods

### NewInframonitoringtypesHostRecord

`func NewInframonitoringtypesHostRecord(activeHostCount int32, cpu float64, diskUsage float64, hostName string, inactiveHostCount int32, load15 float64, memory float64, meta map[string]interface{}, status InframonitoringtypesHostStatus, wait float64, ) *InframonitoringtypesHostRecord`

NewInframonitoringtypesHostRecord instantiates a new InframonitoringtypesHostRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInframonitoringtypesHostRecordWithDefaults

`func NewInframonitoringtypesHostRecordWithDefaults() *InframonitoringtypesHostRecord`

NewInframonitoringtypesHostRecordWithDefaults instantiates a new InframonitoringtypesHostRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveHostCount

`func (o *InframonitoringtypesHostRecord) GetActiveHostCount() int32`

GetActiveHostCount returns the ActiveHostCount field if non-nil, zero value otherwise.

### GetActiveHostCountOk

`func (o *InframonitoringtypesHostRecord) GetActiveHostCountOk() (*int32, bool)`

GetActiveHostCountOk returns a tuple with the ActiveHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveHostCount

`func (o *InframonitoringtypesHostRecord) SetActiveHostCount(v int32)`

SetActiveHostCount sets ActiveHostCount field to given value.


### GetCpu

`func (o *InframonitoringtypesHostRecord) GetCpu() float64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *InframonitoringtypesHostRecord) GetCpuOk() (*float64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *InframonitoringtypesHostRecord) SetCpu(v float64)`

SetCpu sets Cpu field to given value.


### GetDiskUsage

`func (o *InframonitoringtypesHostRecord) GetDiskUsage() float64`

GetDiskUsage returns the DiskUsage field if non-nil, zero value otherwise.

### GetDiskUsageOk

`func (o *InframonitoringtypesHostRecord) GetDiskUsageOk() (*float64, bool)`

GetDiskUsageOk returns a tuple with the DiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsage

`func (o *InframonitoringtypesHostRecord) SetDiskUsage(v float64)`

SetDiskUsage sets DiskUsage field to given value.


### GetHostName

`func (o *InframonitoringtypesHostRecord) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *InframonitoringtypesHostRecord) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *InframonitoringtypesHostRecord) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetInactiveHostCount

`func (o *InframonitoringtypesHostRecord) GetInactiveHostCount() int32`

GetInactiveHostCount returns the InactiveHostCount field if non-nil, zero value otherwise.

### GetInactiveHostCountOk

`func (o *InframonitoringtypesHostRecord) GetInactiveHostCountOk() (*int32, bool)`

GetInactiveHostCountOk returns a tuple with the InactiveHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveHostCount

`func (o *InframonitoringtypesHostRecord) SetInactiveHostCount(v int32)`

SetInactiveHostCount sets InactiveHostCount field to given value.


### GetLoad15

`func (o *InframonitoringtypesHostRecord) GetLoad15() float64`

GetLoad15 returns the Load15 field if non-nil, zero value otherwise.

### GetLoad15Ok

`func (o *InframonitoringtypesHostRecord) GetLoad15Ok() (*float64, bool)`

GetLoad15Ok returns a tuple with the Load15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad15

`func (o *InframonitoringtypesHostRecord) SetLoad15(v float64)`

SetLoad15 sets Load15 field to given value.


### GetMemory

`func (o *InframonitoringtypesHostRecord) GetMemory() float64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *InframonitoringtypesHostRecord) GetMemoryOk() (*float64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *InframonitoringtypesHostRecord) SetMemory(v float64)`

SetMemory sets Memory field to given value.


### GetMeta

`func (o *InframonitoringtypesHostRecord) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InframonitoringtypesHostRecord) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InframonitoringtypesHostRecord) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### SetMetaNil

`func (o *InframonitoringtypesHostRecord) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *InframonitoringtypesHostRecord) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetStatus

`func (o *InframonitoringtypesHostRecord) GetStatus() InframonitoringtypesHostStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InframonitoringtypesHostRecord) GetStatusOk() (*InframonitoringtypesHostStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InframonitoringtypesHostRecord) SetStatus(v InframonitoringtypesHostStatus)`

SetStatus sets Status field to given value.


### GetWait

`func (o *InframonitoringtypesHostRecord) GetWait() float64`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *InframonitoringtypesHostRecord) GetWaitOk() (*float64, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *InframonitoringtypesHostRecord) SetWait(v float64)`

SetWait sets Wait field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


