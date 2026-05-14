# InframonitoringtypesPods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeBeforeRetention** | **bool** |  | 
**Records** | [**[]InframonitoringtypesPodRecord**](InframonitoringtypesPodRecord.md) |  | 
**RequiredMetricsCheck** | [**InframonitoringtypesRequiredMetricsCheck**](InframonitoringtypesRequiredMetricsCheck.md) |  | 
**Total** | **int32** |  | 
**Type** | [**InframonitoringtypesResponseType**](InframonitoringtypesResponseType.md) |  | 
**Warning** | Pointer to [**Querybuildertypesv5QueryWarnData**](Querybuildertypesv5QueryWarnData.md) |  | [optional] 

## Methods

### NewInframonitoringtypesPods

`func NewInframonitoringtypesPods(endTimeBeforeRetention bool, records []InframonitoringtypesPodRecord, requiredMetricsCheck InframonitoringtypesRequiredMetricsCheck, total int32, type_ InframonitoringtypesResponseType, ) *InframonitoringtypesPods`

NewInframonitoringtypesPods instantiates a new InframonitoringtypesPods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInframonitoringtypesPodsWithDefaults

`func NewInframonitoringtypesPodsWithDefaults() *InframonitoringtypesPods`

NewInframonitoringtypesPodsWithDefaults instantiates a new InframonitoringtypesPods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeBeforeRetention

`func (o *InframonitoringtypesPods) GetEndTimeBeforeRetention() bool`

GetEndTimeBeforeRetention returns the EndTimeBeforeRetention field if non-nil, zero value otherwise.

### GetEndTimeBeforeRetentionOk

`func (o *InframonitoringtypesPods) GetEndTimeBeforeRetentionOk() (*bool, bool)`

GetEndTimeBeforeRetentionOk returns a tuple with the EndTimeBeforeRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeBeforeRetention

`func (o *InframonitoringtypesPods) SetEndTimeBeforeRetention(v bool)`

SetEndTimeBeforeRetention sets EndTimeBeforeRetention field to given value.


### GetRecords

`func (o *InframonitoringtypesPods) GetRecords() []InframonitoringtypesPodRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *InframonitoringtypesPods) GetRecordsOk() (*[]InframonitoringtypesPodRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *InframonitoringtypesPods) SetRecords(v []InframonitoringtypesPodRecord)`

SetRecords sets Records field to given value.


### SetRecordsNil

`func (o *InframonitoringtypesPods) SetRecordsNil(b bool)`

 SetRecordsNil sets the value for Records to be an explicit nil

### UnsetRecords
`func (o *InframonitoringtypesPods) UnsetRecords()`

UnsetRecords ensures that no value is present for Records, not even an explicit nil
### GetRequiredMetricsCheck

`func (o *InframonitoringtypesPods) GetRequiredMetricsCheck() InframonitoringtypesRequiredMetricsCheck`

GetRequiredMetricsCheck returns the RequiredMetricsCheck field if non-nil, zero value otherwise.

### GetRequiredMetricsCheckOk

`func (o *InframonitoringtypesPods) GetRequiredMetricsCheckOk() (*InframonitoringtypesRequiredMetricsCheck, bool)`

GetRequiredMetricsCheckOk returns a tuple with the RequiredMetricsCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredMetricsCheck

`func (o *InframonitoringtypesPods) SetRequiredMetricsCheck(v InframonitoringtypesRequiredMetricsCheck)`

SetRequiredMetricsCheck sets RequiredMetricsCheck field to given value.


### GetTotal

`func (o *InframonitoringtypesPods) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InframonitoringtypesPods) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InframonitoringtypesPods) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetType

`func (o *InframonitoringtypesPods) GetType() InframonitoringtypesResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InframonitoringtypesPods) GetTypeOk() (*InframonitoringtypesResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InframonitoringtypesPods) SetType(v InframonitoringtypesResponseType)`

SetType sets Type field to given value.


### GetWarning

`func (o *InframonitoringtypesPods) GetWarning() Querybuildertypesv5QueryWarnData`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *InframonitoringtypesPods) GetWarningOk() (*Querybuildertypesv5QueryWarnData, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *InframonitoringtypesPods) SetWarning(v Querybuildertypesv5QueryWarnData)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *InframonitoringtypesPods) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


