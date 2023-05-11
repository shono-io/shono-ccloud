# GetSrcmV2Cluster200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**SrcmV2ClusterMetadata**](SrcmV2ClusterMetadata.md) |  | [optional] 
**Spec** | [**ListSrcmV2Clusters200ResponseAllOfDataInnerSpec**](ListSrcmV2Clusters200ResponseAllOfDataInnerSpec.md) |  | 
**Status** | [**SrcmV2ClusterStatus**](SrcmV2ClusterStatus.md) |  | 

## Methods

### NewGetSrcmV2Cluster200Response

`func NewGetSrcmV2Cluster200Response(apiVersion string, kind string, id string, spec ListSrcmV2Clusters200ResponseAllOfDataInnerSpec, status SrcmV2ClusterStatus, ) *GetSrcmV2Cluster200Response`

NewGetSrcmV2Cluster200Response instantiates a new GetSrcmV2Cluster200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSrcmV2Cluster200ResponseWithDefaults

`func NewGetSrcmV2Cluster200ResponseWithDefaults() *GetSrcmV2Cluster200Response`

NewGetSrcmV2Cluster200ResponseWithDefaults instantiates a new GetSrcmV2Cluster200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetSrcmV2Cluster200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetSrcmV2Cluster200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetSrcmV2Cluster200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetSrcmV2Cluster200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetSrcmV2Cluster200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetSrcmV2Cluster200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetSrcmV2Cluster200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSrcmV2Cluster200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSrcmV2Cluster200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetSrcmV2Cluster200Response) GetMetadata() SrcmV2ClusterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetSrcmV2Cluster200Response) GetMetadataOk() (*SrcmV2ClusterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetSrcmV2Cluster200Response) SetMetadata(v SrcmV2ClusterMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetSrcmV2Cluster200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *GetSrcmV2Cluster200Response) GetSpec() ListSrcmV2Clusters200ResponseAllOfDataInnerSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *GetSrcmV2Cluster200Response) GetSpecOk() (*ListSrcmV2Clusters200ResponseAllOfDataInnerSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *GetSrcmV2Cluster200Response) SetSpec(v ListSrcmV2Clusters200ResponseAllOfDataInnerSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *GetSrcmV2Cluster200Response) GetStatus() SrcmV2ClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSrcmV2Cluster200Response) GetStatusOk() (*SrcmV2ClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSrcmV2Cluster200Response) SetStatus(v SrcmV2ClusterStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


