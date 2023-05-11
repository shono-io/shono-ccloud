# CreateCmkV2Cluster202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**CmkV2ClusterMetadata**](CmkV2ClusterMetadata.md) |  | [optional] 
**Spec** | [**ListCmkV2Clusters200ResponseAllOfDataInnerSpec**](ListCmkV2Clusters200ResponseAllOfDataInnerSpec.md) |  | 
**Status** | [**CmkV2ClusterStatus**](CmkV2ClusterStatus.md) |  | 

## Methods

### NewCreateCmkV2Cluster202Response

`func NewCreateCmkV2Cluster202Response(spec ListCmkV2Clusters200ResponseAllOfDataInnerSpec, status CmkV2ClusterStatus, ) *CreateCmkV2Cluster202Response`

NewCreateCmkV2Cluster202Response instantiates a new CreateCmkV2Cluster202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCmkV2Cluster202ResponseWithDefaults

`func NewCreateCmkV2Cluster202ResponseWithDefaults() *CreateCmkV2Cluster202Response`

NewCreateCmkV2Cluster202ResponseWithDefaults instantiates a new CreateCmkV2Cluster202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateCmkV2Cluster202Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateCmkV2Cluster202Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateCmkV2Cluster202Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateCmkV2Cluster202Response) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CreateCmkV2Cluster202Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateCmkV2Cluster202Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateCmkV2Cluster202Response) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateCmkV2Cluster202Response) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CreateCmkV2Cluster202Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCmkV2Cluster202Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCmkV2Cluster202Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateCmkV2Cluster202Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateCmkV2Cluster202Response) GetMetadata() CmkV2ClusterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateCmkV2Cluster202Response) GetMetadataOk() (*CmkV2ClusterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateCmkV2Cluster202Response) SetMetadata(v CmkV2ClusterMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateCmkV2Cluster202Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *CreateCmkV2Cluster202Response) GetSpec() ListCmkV2Clusters200ResponseAllOfDataInnerSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreateCmkV2Cluster202Response) GetSpecOk() (*ListCmkV2Clusters200ResponseAllOfDataInnerSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreateCmkV2Cluster202Response) SetSpec(v ListCmkV2Clusters200ResponseAllOfDataInnerSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *CreateCmkV2Cluster202Response) GetStatus() CmkV2ClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateCmkV2Cluster202Response) GetStatusOk() (*CmkV2ClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateCmkV2Cluster202Response) SetStatus(v CmkV2ClusterStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


