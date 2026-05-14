# CreateDowntimeSchedule201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**RuletypesPlannedMaintenance**](RuletypesPlannedMaintenance.md) |  | 
**Status** | **string** |  | 

## Methods

### NewCreateDowntimeSchedule201Response

`func NewCreateDowntimeSchedule201Response(data RuletypesPlannedMaintenance, status string, ) *CreateDowntimeSchedule201Response`

NewCreateDowntimeSchedule201Response instantiates a new CreateDowntimeSchedule201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDowntimeSchedule201ResponseWithDefaults

`func NewCreateDowntimeSchedule201ResponseWithDefaults() *CreateDowntimeSchedule201Response`

NewCreateDowntimeSchedule201ResponseWithDefaults instantiates a new CreateDowntimeSchedule201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateDowntimeSchedule201Response) GetData() RuletypesPlannedMaintenance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateDowntimeSchedule201Response) GetDataOk() (*RuletypesPlannedMaintenance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateDowntimeSchedule201Response) SetData(v RuletypesPlannedMaintenance)`

SetData sets Data field to given value.


### GetStatus

`func (o *CreateDowntimeSchedule201Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateDowntimeSchedule201Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateDowntimeSchedule201Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


