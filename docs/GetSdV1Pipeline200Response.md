# GetSdV1Pipeline200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**SdV1PipelineMetadata**](SdV1PipelineMetadata.md) |  | [optional] 
**Spec** | [**ListSdV1Pipelines200ResponseAllOfDataInnerSpec**](ListSdV1Pipelines200ResponseAllOfDataInnerSpec.md) |  | 
**Status** | Pointer to [**SdV1PipelineStatus**](SdV1PipelineStatus.md) |  | [optional] 

## Methods

### NewGetSdV1Pipeline200Response

`func NewGetSdV1Pipeline200Response(apiVersion string, kind string, id string, spec ListSdV1Pipelines200ResponseAllOfDataInnerSpec, ) *GetSdV1Pipeline200Response`

NewGetSdV1Pipeline200Response instantiates a new GetSdV1Pipeline200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSdV1Pipeline200ResponseWithDefaults

`func NewGetSdV1Pipeline200ResponseWithDefaults() *GetSdV1Pipeline200Response`

NewGetSdV1Pipeline200ResponseWithDefaults instantiates a new GetSdV1Pipeline200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetSdV1Pipeline200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetSdV1Pipeline200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetSdV1Pipeline200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetSdV1Pipeline200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetSdV1Pipeline200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetSdV1Pipeline200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetSdV1Pipeline200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSdV1Pipeline200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSdV1Pipeline200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetSdV1Pipeline200Response) GetMetadata() SdV1PipelineMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetSdV1Pipeline200Response) GetMetadataOk() (*SdV1PipelineMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetSdV1Pipeline200Response) SetMetadata(v SdV1PipelineMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetSdV1Pipeline200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *GetSdV1Pipeline200Response) GetSpec() ListSdV1Pipelines200ResponseAllOfDataInnerSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *GetSdV1Pipeline200Response) GetSpecOk() (*ListSdV1Pipelines200ResponseAllOfDataInnerSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *GetSdV1Pipeline200Response) SetSpec(v ListSdV1Pipelines200ResponseAllOfDataInnerSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *GetSdV1Pipeline200Response) GetStatus() SdV1PipelineStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSdV1Pipeline200Response) GetStatusOk() (*SdV1PipelineStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSdV1Pipeline200Response) SetStatus(v SdV1PipelineStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSdV1Pipeline200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


