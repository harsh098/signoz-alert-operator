# GetRuleHistoryTopContributors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RulestatehistorytypesGettableRuleStateHistoryContributor**](RulestatehistorytypesGettableRuleStateHistoryContributor.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetRuleHistoryTopContributors200Response

`func NewGetRuleHistoryTopContributors200Response(data []RulestatehistorytypesGettableRuleStateHistoryContributor, status string, ) *GetRuleHistoryTopContributors200Response`

NewGetRuleHistoryTopContributors200Response instantiates a new GetRuleHistoryTopContributors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRuleHistoryTopContributors200ResponseWithDefaults

`func NewGetRuleHistoryTopContributors200ResponseWithDefaults() *GetRuleHistoryTopContributors200Response`

NewGetRuleHistoryTopContributors200ResponseWithDefaults instantiates a new GetRuleHistoryTopContributors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRuleHistoryTopContributors200Response) GetData() []RulestatehistorytypesGettableRuleStateHistoryContributor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRuleHistoryTopContributors200Response) GetDataOk() (*[]RulestatehistorytypesGettableRuleStateHistoryContributor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRuleHistoryTopContributors200Response) SetData(v []RulestatehistorytypesGettableRuleStateHistoryContributor)`

SetData sets Data field to given value.


### SetDataNil

`func (o *GetRuleHistoryTopContributors200Response) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GetRuleHistoryTopContributors200Response) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetStatus

`func (o *GetRuleHistoryTopContributors200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRuleHistoryTopContributors200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRuleHistoryTopContributors200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


