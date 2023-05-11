# UpdateSdV1PipelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**SdV1PipelineMetadata**](SdV1PipelineMetadata.md) |  | [optional] 
**Spec** | [**UpdateCmkV2ClusterRequestAllOfSpec**](UpdateCmkV2ClusterRequestAllOfSpec.md) |  | 
**Status** | Pointer to [**SdV1PipelineStatus**](SdV1PipelineStatus.md) |  | [optional] 

## Methods

### NewUpdateSdV1PipelineRequest

`func NewUpdateSdV1PipelineRequest(spec UpdateCmkV2ClusterRequestAllOfSpec, ) *UpdateSdV1PipelineRequest`

NewUpdateSdV1PipelineRequest instantiates a new UpdateSdV1PipelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSdV1PipelineRequestWithDefaults

`func NewUpdateSdV1PipelineRequestWithDefaults() *UpdateSdV1PipelineRequest`

NewUpdateSdV1PipelineRequestWithDefaults instantiates a new UpdateSdV1PipelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *UpdateSdV1PipelineRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UpdateSdV1PipelineRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UpdateSdV1PipelineRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *UpdateSdV1PipelineRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *UpdateSdV1PipelineRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UpdateSdV1PipelineRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UpdateSdV1PipelineRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *UpdateSdV1PipelineRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *UpdateSdV1PipelineRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateSdV1PipelineRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateSdV1PipelineRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateSdV1PipelineRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateSdV1PipelineRequest) GetMetadata() SdV1PipelineMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateSdV1PipelineRequest) GetMetadataOk() (*SdV1PipelineMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateSdV1PipelineRequest) SetMetadata(v SdV1PipelineMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateSdV1PipelineRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *UpdateSdV1PipelineRequest) GetSpec() UpdateCmkV2ClusterRequestAllOfSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UpdateSdV1PipelineRequest) GetSpecOk() (*UpdateCmkV2ClusterRequestAllOfSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UpdateSdV1PipelineRequest) SetSpec(v UpdateCmkV2ClusterRequestAllOfSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *UpdateSdV1PipelineRequest) GetStatus() SdV1PipelineStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateSdV1PipelineRequest) GetStatusOk() (*SdV1PipelineStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateSdV1PipelineRequest) SetStatus(v SdV1PipelineStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateSdV1PipelineRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


