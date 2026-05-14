# GetRuleHistoryStats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**RulestatehistorytypesGettableRuleStateHistoryStats**](RulestatehistorytypesGettableRuleStateHistoryStats.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetRuleHistoryStats200Response

`func NewGetRuleHistoryStats200Response(data RulestatehistorytypesGettableRuleStateHistoryStats, status string, ) *GetRuleHistoryStats200Response`

NewGetRuleHistoryStats200Response instantiates a new GetRuleHistoryStats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRuleHistoryStats200ResponseWithDefaults

`func NewGetRuleHistoryStats200ResponseWithDefaults() *GetRuleHistoryStats200Response`

NewGetRuleHistoryStats200ResponseWithDefaults instantiates a new GetRuleHistoryStats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRuleHistoryStats200Response) GetData() RulestatehistorytypesGettableRuleStateHistoryStats`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRuleHistoryStats200Response) GetDataOk() (*RulestatehistorytypesGettableRuleStateHistoryStats, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRuleHistoryStats200Response) SetData(v RulestatehistorytypesGettableRuleStateHistoryStats)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetRuleHistoryStats200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRuleHistoryStats200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRuleHistoryStats200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


