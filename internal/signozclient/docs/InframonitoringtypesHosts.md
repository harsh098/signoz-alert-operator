# InframonitoringtypesHosts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeBeforeRetention** | **bool** |  | 
**Records** | [**[]InframonitoringtypesHostRecord**](InframonitoringtypesHostRecord.md) |  | 
**RequiredMetricsCheck** | [**InframonitoringtypesRequiredMetricsCheck**](InframonitoringtypesRequiredMetricsCheck.md) |  | 
**Total** | **int32** |  | 
**Type** | [**InframonitoringtypesResponseType**](InframonitoringtypesResponseType.md) |  | 
**Warning** | Pointer to [**Querybuildertypesv5QueryWarnData**](Querybuildertypesv5QueryWarnData.md) |  | [optional] 

## Methods

### NewInframonitoringtypesHosts

`func NewInframonitoringtypesHosts(endTimeBeforeRetention bool, records []InframonitoringtypesHostRecord, requiredMetricsCheck InframonitoringtypesRequiredMetricsCheck, total int32, type_ InframonitoringtypesResponseType, ) *InframonitoringtypesHosts`

NewInframonitoringtypesHosts instantiates a new InframonitoringtypesHosts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInframonitoringtypesHostsWithDefaults

`func NewInframonitoringtypesHostsWithDefaults() *InframonitoringtypesHosts`

NewInframonitoringtypesHostsWithDefaults instantiates a new InframonitoringtypesHosts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeBeforeRetention

`func (o *InframonitoringtypesHosts) GetEndTimeBeforeRetention() bool`

GetEndTimeBeforeRetention returns the EndTimeBeforeRetention field if non-nil, zero value otherwise.

### GetEndTimeBeforeRetentionOk

`func (o *InframonitoringtypesHosts) GetEndTimeBeforeRetentionOk() (*bool, bool)`

GetEndTimeBeforeRetentionOk returns a tuple with the EndTimeBeforeRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeBeforeRetention

`func (o *InframonitoringtypesHosts) SetEndTimeBeforeRetention(v bool)`

SetEndTimeBeforeRetention sets EndTimeBeforeRetention field to given value.


### GetRecords

`func (o *InframonitoringtypesHosts) GetRecords() []InframonitoringtypesHostRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *InframonitoringtypesHosts) GetRecordsOk() (*[]InframonitoringtypesHostRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *InframonitoringtypesHosts) SetRecords(v []InframonitoringtypesHostRecord)`

SetRecords sets Records field to given value.


### SetRecordsNil

`func (o *InframonitoringtypesHosts) SetRecordsNil(b bool)`

 SetRecordsNil sets the value for Records to be an explicit nil

### UnsetRecords
`func (o *InframonitoringtypesHosts) UnsetRecords()`

UnsetRecords ensures that no value is present for Records, not even an explicit nil
### GetRequiredMetricsCheck

`func (o *InframonitoringtypesHosts) GetRequiredMetricsCheck() InframonitoringtypesRequiredMetricsCheck`

GetRequiredMetricsCheck returns the RequiredMetricsCheck field if non-nil, zero value otherwise.

### GetRequiredMetricsCheckOk

`func (o *InframonitoringtypesHosts) GetRequiredMetricsCheckOk() (*InframonitoringtypesRequiredMetricsCheck, bool)`

GetRequiredMetricsCheckOk returns a tuple with the RequiredMetricsCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredMetricsCheck

`func (o *InframonitoringtypesHosts) SetRequiredMetricsCheck(v InframonitoringtypesRequiredMetricsCheck)`

SetRequiredMetricsCheck sets RequiredMetricsCheck field to given value.


### GetTotal

`func (o *InframonitoringtypesHosts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InframonitoringtypesHosts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InframonitoringtypesHosts) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetType

`func (o *InframonitoringtypesHosts) GetType() InframonitoringtypesResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InframonitoringtypesHosts) GetTypeOk() (*InframonitoringtypesResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InframonitoringtypesHosts) SetType(v InframonitoringtypesResponseType)`

SetType sets Type field to given value.


### GetWarning

`func (o *InframonitoringtypesHosts) GetWarning() Querybuildertypesv5QueryWarnData`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *InframonitoringtypesHosts) GetWarningOk() (*Querybuildertypesv5QueryWarnData, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *InframonitoringtypesHosts) SetWarning(v Querybuildertypesv5QueryWarnData)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *InframonitoringtypesHosts) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


