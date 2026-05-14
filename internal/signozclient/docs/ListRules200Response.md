# ListRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RuletypesRule**](RuletypesRule.md) |  | 
**Status** | **string** |  | 

## Methods

### NewListRules200Response

`func NewListRules200Response(data []RuletypesRule, status string, ) *ListRules200Response`

NewListRules200Response instantiates a new ListRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRules200ResponseWithDefaults

`func NewListRules200ResponseWithDefaults() *ListRules200Response`

NewListRules200ResponseWithDefaults instantiates a new ListRules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListRules200Response) GetData() []RuletypesRule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListRules200Response) GetDataOk() (*[]RuletypesRule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListRules200Response) SetData(v []RuletypesRule)`

SetData sets Data field to given value.


### GetStatus

`func (o *ListRules200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListRules200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListRules200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


