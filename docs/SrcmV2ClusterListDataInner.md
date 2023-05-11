# SrcmV2ClusterListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | [**SrcmV2ClusterMetadata**](SrcmV2ClusterMetadata.md) |  | 
**Spec** | **map[string]interface{}** |  | 
**Status** | [**SrcmV2ClusterStatus**](SrcmV2ClusterStatus.md) |  | 

## Methods

### NewSrcmV2ClusterListDataInner

`func NewSrcmV2ClusterListDataInner(id string, metadata SrcmV2ClusterMetadata, spec map[string]interface{}, status SrcmV2ClusterStatus, ) *SrcmV2ClusterListDataInner`

NewSrcmV2ClusterListDataInner instantiates a new SrcmV2ClusterListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrcmV2ClusterListDataInnerWithDefaults

`func NewSrcmV2ClusterListDataInnerWithDefaults() *SrcmV2ClusterListDataInner`

NewSrcmV2ClusterListDataInnerWithDefaults instantiates a new SrcmV2ClusterListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SrcmV2ClusterListDataInner) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SrcmV2ClusterListDataInner) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SrcmV2ClusterListDataInner) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SrcmV2ClusterListDataInner) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SrcmV2ClusterListDataInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SrcmV2ClusterListDataInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SrcmV2ClusterListDataInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SrcmV2ClusterListDataInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *SrcmV2ClusterListDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SrcmV2ClusterListDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SrcmV2ClusterListDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *SrcmV2ClusterListDataInner) GetMetadata() SrcmV2ClusterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SrcmV2ClusterListDataInner) GetMetadataOk() (*SrcmV2ClusterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SrcmV2ClusterListDataInner) SetMetadata(v SrcmV2ClusterMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *SrcmV2ClusterListDataInner) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SrcmV2ClusterListDataInner) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SrcmV2ClusterListDataInner) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *SrcmV2ClusterListDataInner) GetStatus() SrcmV2ClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SrcmV2ClusterListDataInner) GetStatusOk() (*SrcmV2ClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SrcmV2ClusterListDataInner) SetStatus(v SrcmV2ClusterStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


