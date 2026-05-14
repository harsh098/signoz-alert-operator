# RulestatehistorytypesGettableRuleStateTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]RulestatehistorytypesGettableRuleStateHistory**](RulestatehistorytypesGettableRuleStateHistory.md) |  | 
**NextCursor** | Pointer to **string** |  | [optional] 
**Total** | **int32** |  | 

## Methods

### NewRulestatehistorytypesGettableRuleStateTimeline

`func NewRulestatehistorytypesGettableRuleStateTimeline(items []RulestatehistorytypesGettableRuleStateHistory, total int32, ) *RulestatehistorytypesGettableRuleStateTimeline`

NewRulestatehistorytypesGettableRuleStateTimeline instantiates a new RulestatehistorytypesGettableRuleStateTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulestatehistorytypesGettableRuleStateTimelineWithDefaults

`func NewRulestatehistorytypesGettableRuleStateTimelineWithDefaults() *RulestatehistorytypesGettableRuleStateTimeline`

NewRulestatehistorytypesGettableRuleStateTimelineWithDefaults instantiates a new RulestatehistorytypesGettableRuleStateTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *RulestatehistorytypesGettableRuleStateTimeline) GetItems() []RulestatehistorytypesGettableRuleStateHistory`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RulestatehistorytypesGettableRuleStateTimeline) GetItemsOk() (*[]RulestatehistorytypesGettableRuleStateHistory, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RulestatehistorytypesGettableRuleStateTimeline) SetItems(v []RulestatehistorytypesGettableRuleStateHistory)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *RulestatehistorytypesGettableRuleStateTimeline) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *RulestatehistorytypesGettableRuleStateTimeline) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetNextCursor

`func (o *RulestatehistorytypesGettableRuleStateTimeline) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *RulestatehistorytypesGettableRuleStateTimeline) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *RulestatehistorytypesGettableRuleStateTimeline) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *RulestatehistorytypesGettableRuleStateTimeline) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetTotal

`func (o *RulestatehistorytypesGettableRuleStateTimeline) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RulestatehistorytypesGettableRuleStateTimeline) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RulestatehistorytypesGettableRuleStateTimeline) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


