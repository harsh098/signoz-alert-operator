# TracedetailtypesWaterfallSpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**DbName** | Pointer to **string** |  | [optional] 
**DbOperation** | Pointer to **string** |  | [optional] 
**DurationNano** | Pointer to **int32** |  | [optional] 
**Events** | Pointer to [**[]TracedetailtypesEvent**](TracedetailtypesEvent.md) |  | [optional] 
**ExternalHttpMethod** | Pointer to **string** |  | [optional] 
**ExternalHttpUrl** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to **int32** |  | [optional] 
**HasChildren** | Pointer to **bool** |  | [optional] 
**HasError** | Pointer to **bool** |  | [optional] 
**HttpHost** | Pointer to **string** |  | [optional] 
**HttpMethod** | Pointer to **string** |  | [optional] 
**HttpUrl** | Pointer to **string** |  | [optional] 
**IsRemote** | Pointer to **string** |  | [optional] 
**KindString** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentSpanId** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **map[string]string** |  | [optional] 
**ResponseStatusCode** | Pointer to **string** |  | [optional] 
**SpanId** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**StatusCodeString** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**SubTreeNodeCount** | Pointer to **int32** |  | [optional] 
**TimeUnix** | Pointer to **int32** |  | [optional] 
**TraceId** | Pointer to **string** |  | [optional] 
**TraceState** | Pointer to **string** |  | [optional] 

## Methods

### NewTracedetailtypesWaterfallSpan

`func NewTracedetailtypesWaterfallSpan() *TracedetailtypesWaterfallSpan`

NewTracedetailtypesWaterfallSpan instantiates a new TracedetailtypesWaterfallSpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTracedetailtypesWaterfallSpanWithDefaults

`func NewTracedetailtypesWaterfallSpanWithDefaults() *TracedetailtypesWaterfallSpan`

NewTracedetailtypesWaterfallSpanWithDefaults instantiates a new TracedetailtypesWaterfallSpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *TracedetailtypesWaterfallSpan) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TracedetailtypesWaterfallSpan) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TracedetailtypesWaterfallSpan) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TracedetailtypesWaterfallSpan) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TracedetailtypesWaterfallSpan) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TracedetailtypesWaterfallSpan) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetDbName

`func (o *TracedetailtypesWaterfallSpan) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *TracedetailtypesWaterfallSpan) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *TracedetailtypesWaterfallSpan) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *TracedetailtypesWaterfallSpan) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDbOperation

`func (o *TracedetailtypesWaterfallSpan) GetDbOperation() string`

GetDbOperation returns the DbOperation field if non-nil, zero value otherwise.

### GetDbOperationOk

`func (o *TracedetailtypesWaterfallSpan) GetDbOperationOk() (*string, bool)`

GetDbOperationOk returns a tuple with the DbOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbOperation

`func (o *TracedetailtypesWaterfallSpan) SetDbOperation(v string)`

SetDbOperation sets DbOperation field to given value.

### HasDbOperation

`func (o *TracedetailtypesWaterfallSpan) HasDbOperation() bool`

HasDbOperation returns a boolean if a field has been set.

### GetDurationNano

`func (o *TracedetailtypesWaterfallSpan) GetDurationNano() int32`

GetDurationNano returns the DurationNano field if non-nil, zero value otherwise.

### GetDurationNanoOk

`func (o *TracedetailtypesWaterfallSpan) GetDurationNanoOk() (*int32, bool)`

GetDurationNanoOk returns a tuple with the DurationNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationNano

`func (o *TracedetailtypesWaterfallSpan) SetDurationNano(v int32)`

SetDurationNano sets DurationNano field to given value.

### HasDurationNano

`func (o *TracedetailtypesWaterfallSpan) HasDurationNano() bool`

HasDurationNano returns a boolean if a field has been set.

### GetEvents

`func (o *TracedetailtypesWaterfallSpan) GetEvents() []TracedetailtypesEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TracedetailtypesWaterfallSpan) GetEventsOk() (*[]TracedetailtypesEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TracedetailtypesWaterfallSpan) SetEvents(v []TracedetailtypesEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *TracedetailtypesWaterfallSpan) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *TracedetailtypesWaterfallSpan) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *TracedetailtypesWaterfallSpan) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetExternalHttpMethod

`func (o *TracedetailtypesWaterfallSpan) GetExternalHttpMethod() string`

GetExternalHttpMethod returns the ExternalHttpMethod field if non-nil, zero value otherwise.

### GetExternalHttpMethodOk

`func (o *TracedetailtypesWaterfallSpan) GetExternalHttpMethodOk() (*string, bool)`

GetExternalHttpMethodOk returns a tuple with the ExternalHttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHttpMethod

`func (o *TracedetailtypesWaterfallSpan) SetExternalHttpMethod(v string)`

SetExternalHttpMethod sets ExternalHttpMethod field to given value.

### HasExternalHttpMethod

`func (o *TracedetailtypesWaterfallSpan) HasExternalHttpMethod() bool`

HasExternalHttpMethod returns a boolean if a field has been set.

### GetExternalHttpUrl

`func (o *TracedetailtypesWaterfallSpan) GetExternalHttpUrl() string`

GetExternalHttpUrl returns the ExternalHttpUrl field if non-nil, zero value otherwise.

### GetExternalHttpUrlOk

`func (o *TracedetailtypesWaterfallSpan) GetExternalHttpUrlOk() (*string, bool)`

GetExternalHttpUrlOk returns a tuple with the ExternalHttpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHttpUrl

`func (o *TracedetailtypesWaterfallSpan) SetExternalHttpUrl(v string)`

SetExternalHttpUrl sets ExternalHttpUrl field to given value.

### HasExternalHttpUrl

`func (o *TracedetailtypesWaterfallSpan) HasExternalHttpUrl() bool`

HasExternalHttpUrl returns a boolean if a field has been set.

### GetFlags

`func (o *TracedetailtypesWaterfallSpan) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *TracedetailtypesWaterfallSpan) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *TracedetailtypesWaterfallSpan) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *TracedetailtypesWaterfallSpan) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetHasChildren

`func (o *TracedetailtypesWaterfallSpan) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *TracedetailtypesWaterfallSpan) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *TracedetailtypesWaterfallSpan) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.

### HasHasChildren

`func (o *TracedetailtypesWaterfallSpan) HasHasChildren() bool`

HasHasChildren returns a boolean if a field has been set.

### GetHasError

`func (o *TracedetailtypesWaterfallSpan) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *TracedetailtypesWaterfallSpan) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *TracedetailtypesWaterfallSpan) SetHasError(v bool)`

SetHasError sets HasError field to given value.

### HasHasError

`func (o *TracedetailtypesWaterfallSpan) HasHasError() bool`

HasHasError returns a boolean if a field has been set.

### GetHttpHost

`func (o *TracedetailtypesWaterfallSpan) GetHttpHost() string`

GetHttpHost returns the HttpHost field if non-nil, zero value otherwise.

### GetHttpHostOk

`func (o *TracedetailtypesWaterfallSpan) GetHttpHostOk() (*string, bool)`

GetHttpHostOk returns a tuple with the HttpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHost

`func (o *TracedetailtypesWaterfallSpan) SetHttpHost(v string)`

SetHttpHost sets HttpHost field to given value.

### HasHttpHost

`func (o *TracedetailtypesWaterfallSpan) HasHttpHost() bool`

HasHttpHost returns a boolean if a field has been set.

### GetHttpMethod

`func (o *TracedetailtypesWaterfallSpan) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *TracedetailtypesWaterfallSpan) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *TracedetailtypesWaterfallSpan) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *TracedetailtypesWaterfallSpan) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetHttpUrl

`func (o *TracedetailtypesWaterfallSpan) GetHttpUrl() string`

GetHttpUrl returns the HttpUrl field if non-nil, zero value otherwise.

### GetHttpUrlOk

`func (o *TracedetailtypesWaterfallSpan) GetHttpUrlOk() (*string, bool)`

GetHttpUrlOk returns a tuple with the HttpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpUrl

`func (o *TracedetailtypesWaterfallSpan) SetHttpUrl(v string)`

SetHttpUrl sets HttpUrl field to given value.

### HasHttpUrl

`func (o *TracedetailtypesWaterfallSpan) HasHttpUrl() bool`

HasHttpUrl returns a boolean if a field has been set.

### GetIsRemote

`func (o *TracedetailtypesWaterfallSpan) GetIsRemote() string`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *TracedetailtypesWaterfallSpan) GetIsRemoteOk() (*string, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *TracedetailtypesWaterfallSpan) SetIsRemote(v string)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *TracedetailtypesWaterfallSpan) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetKindString

`func (o *TracedetailtypesWaterfallSpan) GetKindString() string`

GetKindString returns the KindString field if non-nil, zero value otherwise.

### GetKindStringOk

`func (o *TracedetailtypesWaterfallSpan) GetKindStringOk() (*string, bool)`

GetKindStringOk returns a tuple with the KindString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKindString

`func (o *TracedetailtypesWaterfallSpan) SetKindString(v string)`

SetKindString sets KindString field to given value.

### HasKindString

`func (o *TracedetailtypesWaterfallSpan) HasKindString() bool`

HasKindString returns a boolean if a field has been set.

### GetLevel

`func (o *TracedetailtypesWaterfallSpan) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *TracedetailtypesWaterfallSpan) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *TracedetailtypesWaterfallSpan) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *TracedetailtypesWaterfallSpan) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetName

`func (o *TracedetailtypesWaterfallSpan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TracedetailtypesWaterfallSpan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TracedetailtypesWaterfallSpan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TracedetailtypesWaterfallSpan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentSpanId

`func (o *TracedetailtypesWaterfallSpan) GetParentSpanId() string`

GetParentSpanId returns the ParentSpanId field if non-nil, zero value otherwise.

### GetParentSpanIdOk

`func (o *TracedetailtypesWaterfallSpan) GetParentSpanIdOk() (*string, bool)`

GetParentSpanIdOk returns a tuple with the ParentSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpanId

`func (o *TracedetailtypesWaterfallSpan) SetParentSpanId(v string)`

SetParentSpanId sets ParentSpanId field to given value.

### HasParentSpanId

`func (o *TracedetailtypesWaterfallSpan) HasParentSpanId() bool`

HasParentSpanId returns a boolean if a field has been set.

### GetResource

`func (o *TracedetailtypesWaterfallSpan) GetResource() map[string]string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *TracedetailtypesWaterfallSpan) GetResourceOk() (*map[string]string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *TracedetailtypesWaterfallSpan) SetResource(v map[string]string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *TracedetailtypesWaterfallSpan) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *TracedetailtypesWaterfallSpan) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *TracedetailtypesWaterfallSpan) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetResponseStatusCode

`func (o *TracedetailtypesWaterfallSpan) GetResponseStatusCode() string`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *TracedetailtypesWaterfallSpan) GetResponseStatusCodeOk() (*string, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *TracedetailtypesWaterfallSpan) SetResponseStatusCode(v string)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *TracedetailtypesWaterfallSpan) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetSpanId

`func (o *TracedetailtypesWaterfallSpan) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *TracedetailtypesWaterfallSpan) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *TracedetailtypesWaterfallSpan) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *TracedetailtypesWaterfallSpan) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetStatusCode

`func (o *TracedetailtypesWaterfallSpan) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *TracedetailtypesWaterfallSpan) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *TracedetailtypesWaterfallSpan) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *TracedetailtypesWaterfallSpan) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetStatusCodeString

`func (o *TracedetailtypesWaterfallSpan) GetStatusCodeString() string`

GetStatusCodeString returns the StatusCodeString field if non-nil, zero value otherwise.

### GetStatusCodeStringOk

`func (o *TracedetailtypesWaterfallSpan) GetStatusCodeStringOk() (*string, bool)`

GetStatusCodeStringOk returns a tuple with the StatusCodeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeString

`func (o *TracedetailtypesWaterfallSpan) SetStatusCodeString(v string)`

SetStatusCodeString sets StatusCodeString field to given value.

### HasStatusCodeString

`func (o *TracedetailtypesWaterfallSpan) HasStatusCodeString() bool`

HasStatusCodeString returns a boolean if a field has been set.

### GetStatusMessage

`func (o *TracedetailtypesWaterfallSpan) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *TracedetailtypesWaterfallSpan) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *TracedetailtypesWaterfallSpan) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *TracedetailtypesWaterfallSpan) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetSubTreeNodeCount

`func (o *TracedetailtypesWaterfallSpan) GetSubTreeNodeCount() int32`

GetSubTreeNodeCount returns the SubTreeNodeCount field if non-nil, zero value otherwise.

### GetSubTreeNodeCountOk

`func (o *TracedetailtypesWaterfallSpan) GetSubTreeNodeCountOk() (*int32, bool)`

GetSubTreeNodeCountOk returns a tuple with the SubTreeNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTreeNodeCount

`func (o *TracedetailtypesWaterfallSpan) SetSubTreeNodeCount(v int32)`

SetSubTreeNodeCount sets SubTreeNodeCount field to given value.

### HasSubTreeNodeCount

`func (o *TracedetailtypesWaterfallSpan) HasSubTreeNodeCount() bool`

HasSubTreeNodeCount returns a boolean if a field has been set.

### GetTimeUnix

`func (o *TracedetailtypesWaterfallSpan) GetTimeUnix() int32`

GetTimeUnix returns the TimeUnix field if non-nil, zero value otherwise.

### GetTimeUnixOk

`func (o *TracedetailtypesWaterfallSpan) GetTimeUnixOk() (*int32, bool)`

GetTimeUnixOk returns a tuple with the TimeUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnix

`func (o *TracedetailtypesWaterfallSpan) SetTimeUnix(v int32)`

SetTimeUnix sets TimeUnix field to given value.

### HasTimeUnix

`func (o *TracedetailtypesWaterfallSpan) HasTimeUnix() bool`

HasTimeUnix returns a boolean if a field has been set.

### GetTraceId

`func (o *TracedetailtypesWaterfallSpan) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *TracedetailtypesWaterfallSpan) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *TracedetailtypesWaterfallSpan) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *TracedetailtypesWaterfallSpan) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetTraceState

`func (o *TracedetailtypesWaterfallSpan) GetTraceState() string`

GetTraceState returns the TraceState field if non-nil, zero value otherwise.

### GetTraceStateOk

`func (o *TracedetailtypesWaterfallSpan) GetTraceStateOk() (*string, bool)`

GetTraceStateOk returns a tuple with the TraceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceState

`func (o *TracedetailtypesWaterfallSpan) SetTraceState(v string)`

SetTraceState sets TraceState field to given value.

### HasTraceState

`func (o *TracedetailtypesWaterfallSpan) HasTraceState() bool`

HasTraceState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


