# ListLLMPricingRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**LlmpricingruletypesGettablePricingRules**](LlmpricingruletypesGettablePricingRules.md) |  | 
**Status** | **string** |  | 

## Methods

### NewListLLMPricingRules200Response

`func NewListLLMPricingRules200Response(data LlmpricingruletypesGettablePricingRules, status string, ) *ListLLMPricingRules200Response`

NewListLLMPricingRules200Response instantiates a new ListLLMPricingRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLLMPricingRules200ResponseWithDefaults

`func NewListLLMPricingRules200ResponseWithDefaults() *ListLLMPricingRules200Response`

NewListLLMPricingRules200ResponseWithDefaults instantiates a new ListLLMPricingRules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListLLMPricingRules200Response) GetData() LlmpricingruletypesGettablePricingRules`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListLLMPricingRules200Response) GetDataOk() (*LlmpricingruletypesGettablePricingRules, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListLLMPricingRules200Response) SetData(v LlmpricingruletypesGettablePricingRules)`

SetData sets Data field to given value.


### GetStatus

`func (o *ListLLMPricingRules200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListLLMPricingRules200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListLLMPricingRules200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


