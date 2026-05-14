# GetRuleHistoryOverallStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RulestatehistorytypesGettableRuleStateWindow**](RulestatehistorytypesGettableRuleStateWindow.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetRuleHistoryOverallStatus200Response

`func NewGetRuleHistoryOverallStatus200Response(data []RulestatehistorytypesGettableRuleStateWindow, status string, ) *GetRuleHistoryOverallStatus200Response`

NewGetRuleHistoryOverallStatus200Response instantiates a new GetRuleHistoryOverallStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRuleHistoryOverallStatus200ResponseWithDefaults

`func NewGetRuleHistoryOverallStatus200ResponseWithDefaults() *GetRuleHistoryOverallStatus200Response`

NewGetRuleHistoryOverallStatus200ResponseWithDefaults instantiates a new GetRuleHistoryOverallStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRuleHistoryOverallStatus200Response) GetData() []RulestatehistorytypesGettableRuleStateWindow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRuleHistoryOverallStatus200Response) GetDataOk() (*[]RulestatehistorytypesGettableRuleStateWindow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRuleHistoryOverallStatus200Response) SetData(v []RulestatehistorytypesGettableRuleStateWindow)`

SetData sets Data field to given value.


### SetDataNil

`func (o *GetRuleHistoryOverallStatus200Response) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GetRuleHistoryOverallStatus200Response) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetStatus

`func (o *GetRuleHistoryOverallStatus200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRuleHistoryOverallStatus200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRuleHistoryOverallStatus200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


