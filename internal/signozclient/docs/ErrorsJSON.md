# ErrorsJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Errors** | Pointer to [**[]ErrorsResponseerroradditional**](ErrorsResponseerroradditional.md) |  | [optional] 
**Message** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorsJSON

`func NewErrorsJSON(code string, message string, ) *ErrorsJSON`

NewErrorsJSON instantiates a new ErrorsJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsJSONWithDefaults

`func NewErrorsJSONWithDefaults() *ErrorsJSON`

NewErrorsJSONWithDefaults instantiates a new ErrorsJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorsJSON) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorsJSON) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorsJSON) SetCode(v string)`

SetCode sets Code field to given value.


### GetErrors

`func (o *ErrorsJSON) GetErrors() []ErrorsResponseerroradditional`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsJSON) GetErrorsOk() (*[]ErrorsResponseerroradditional, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsJSON) SetErrors(v []ErrorsResponseerroradditional)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ErrorsJSON) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorsJSON) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorsJSON) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorsJSON) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetUrl

`func (o *ErrorsJSON) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ErrorsJSON) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ErrorsJSON) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ErrorsJSON) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


