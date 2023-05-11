# UpdateSrcmV2ClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**SrcmV2ClusterMetadata**](SrcmV2ClusterMetadata.md) |  | [optional] 
**Spec** | [**UpdateCmkV2ClusterRequestAllOfSpec**](UpdateCmkV2ClusterRequestAllOfSpec.md) |  | 
**Status** | Pointer to [**SrcmV2ClusterStatus**](SrcmV2ClusterStatus.md) |  | [optional] 

## Methods

### NewUpdateSrcmV2ClusterRequest

`func NewUpdateSrcmV2ClusterRequest(spec UpdateCmkV2ClusterRequestAllOfSpec, ) *UpdateSrcmV2ClusterRequest`

NewUpdateSrcmV2ClusterRequest instantiates a new UpdateSrcmV2ClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSrcmV2ClusterRequestWithDefaults

`func NewUpdateSrcmV2ClusterRequestWithDefaults() *UpdateSrcmV2ClusterRequest`

NewUpdateSrcmV2ClusterRequestWithDefaults instantiates a new UpdateSrcmV2ClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *UpdateSrcmV2ClusterRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UpdateSrcmV2ClusterRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UpdateSrcmV2ClusterRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *UpdateSrcmV2ClusterRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *UpdateSrcmV2ClusterRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UpdateSrcmV2ClusterRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UpdateSrcmV2ClusterRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *UpdateSrcmV2ClusterRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *UpdateSrcmV2ClusterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateSrcmV2ClusterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateSrcmV2ClusterRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateSrcmV2ClusterRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateSrcmV2ClusterRequest) GetMetadata() SrcmV2ClusterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateSrcmV2ClusterRequest) GetMetadataOk() (*SrcmV2ClusterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateSrcmV2ClusterRequest) SetMetadata(v SrcmV2ClusterMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateSrcmV2ClusterRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *UpdateSrcmV2ClusterRequest) GetSpec() UpdateCmkV2ClusterRequestAllOfSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UpdateSrcmV2ClusterRequest) GetSpecOk() (*UpdateCmkV2ClusterRequestAllOfSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UpdateSrcmV2ClusterRequest) SetSpec(v UpdateCmkV2ClusterRequestAllOfSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *UpdateSrcmV2ClusterRequest) GetStatus() SrcmV2ClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateSrcmV2ClusterRequest) GetStatusOk() (*SrcmV2ClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateSrcmV2ClusterRequest) SetStatus(v SrcmV2ClusterStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateSrcmV2ClusterRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


