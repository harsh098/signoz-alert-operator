# AlertmanagertypesDeprecatedGettableAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**EndsAt** | Pointer to **time.Time** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**GeneratorURL** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Receivers** | Pointer to **[]string** |  | [optional] 
**StartsAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**TypesAlertStatus**](TypesAlertStatus.md) |  | [optional] 

## Methods

### NewAlertmanagertypesDeprecatedGettableAlert

`func NewAlertmanagertypesDeprecatedGettableAlert() *AlertmanagertypesDeprecatedGettableAlert`

NewAlertmanagertypesDeprecatedGettableAlert instantiates a new AlertmanagertypesDeprecatedGettableAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertmanagertypesDeprecatedGettableAlertWithDefaults

`func NewAlertmanagertypesDeprecatedGettableAlertWithDefaults() *AlertmanagertypesDeprecatedGettableAlert`

NewAlertmanagertypesDeprecatedGettableAlertWithDefaults instantiates a new AlertmanagertypesDeprecatedGettableAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *AlertmanagertypesDeprecatedGettableAlert) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *AlertmanagertypesDeprecatedGettableAlert) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetEndsAt

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *AlertmanagertypesDeprecatedGettableAlert) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *AlertmanagertypesDeprecatedGettableAlert) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetFingerprint

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *AlertmanagertypesDeprecatedGettableAlert) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *AlertmanagertypesDeprecatedGettableAlert) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetGeneratorURL

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetGeneratorURL() string`

GetGeneratorURL returns the GeneratorURL field if non-nil, zero value otherwise.

### GetGeneratorURLOk

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetGeneratorURLOk() (*string, bool)`

GetGeneratorURLOk returns a tuple with the GeneratorURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorURL

`func (o *AlertmanagertypesDeprecatedGettableAlert) SetGeneratorURL(v string)`

SetGeneratorURL sets GeneratorURL field to given value.

### HasGeneratorURL

`func (o *AlertmanagertypesDeprecatedGettableAlert) HasGeneratorURL() bool`

HasGeneratorURL returns a boolean if a field has been set.

### GetLabels

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AlertmanagertypesDeprecatedGettableAlert) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AlertmanagertypesDeprecatedGettableAlert) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetReceivers

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetReceivers() []string`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetReceiversOk() (*[]string, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *AlertmanagertypesDeprecatedGettableAlert) SetReceivers(v []string)`

SetReceivers sets Receivers field to given value.

### HasReceivers

`func (o *AlertmanagertypesDeprecatedGettableAlert) HasReceivers() bool`

HasReceivers returns a boolean if a field has been set.

### SetReceiversNil

`func (o *AlertmanagertypesDeprecatedGettableAlert) SetReceiversNil(b bool)`

 SetReceiversNil sets the value for Receivers to be an explicit nil

### UnsetReceivers
`func (o *AlertmanagertypesDeprecatedGettableAlert) UnsetReceivers()`

UnsetReceivers ensures that no value is present for Receivers, not even an explicit nil
### GetStartsAt

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *AlertmanagertypesDeprecatedGettableAlert) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *AlertmanagertypesDeprecatedGettableAlert) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetStatus

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetStatus() TypesAlertStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertmanagertypesDeprecatedGettableAlert) GetStatusOk() (*TypesAlertStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertmanagertypesDeprecatedGettableAlert) SetStatus(v TypesAlertStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertmanagertypesDeprecatedGettableAlert) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


