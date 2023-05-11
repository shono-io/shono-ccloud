# ListSdV1Pipelines200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**SdV1PipelineListMetadata**](SdV1PipelineListMetadata.md) |  | 
**Data** | [**[]ListSdV1Pipelines200ResponseAllOfDataInner**](ListSdV1Pipelines200ResponseAllOfDataInner.md) |  | 

## Methods

### NewListSdV1Pipelines200Response

`func NewListSdV1Pipelines200Response(apiVersion string, kind string, metadata SdV1PipelineListMetadata, data []ListSdV1Pipelines200ResponseAllOfDataInner, ) *ListSdV1Pipelines200Response`

NewListSdV1Pipelines200Response instantiates a new ListSdV1Pipelines200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSdV1Pipelines200ResponseWithDefaults

`func NewListSdV1Pipelines200ResponseWithDefaults() *ListSdV1Pipelines200Response`

NewListSdV1Pipelines200ResponseWithDefaults instantiates a new ListSdV1Pipelines200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListSdV1Pipelines200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListSdV1Pipelines200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListSdV1Pipelines200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ListSdV1Pipelines200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListSdV1Pipelines200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListSdV1Pipelines200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListSdV1Pipelines200Response) GetMetadata() SdV1PipelineListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListSdV1Pipelines200Response) GetMetadataOk() (*SdV1PipelineListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListSdV1Pipelines200Response) SetMetadata(v SdV1PipelineListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListSdV1Pipelines200Response) GetData() []ListSdV1Pipelines200ResponseAllOfDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListSdV1Pipelines200Response) GetDataOk() (*[]ListSdV1Pipelines200ResponseAllOfDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListSdV1Pipelines200Response) SetData(v []ListSdV1Pipelines200ResponseAllOfDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


