# GetKsqldbcmV2Cluster200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**KsqldbcmV2ClusterMetadata**](KsqldbcmV2ClusterMetadata.md) |  | [optional] 
**Spec** | [**ListKsqldbcmV2Clusters200ResponseAllOfDataInnerSpec**](ListKsqldbcmV2Clusters200ResponseAllOfDataInnerSpec.md) |  | 
**Status** | [**KsqldbcmV2ClusterStatus**](KsqldbcmV2ClusterStatus.md) |  | 

## Methods

### NewGetKsqldbcmV2Cluster200Response

`func NewGetKsqldbcmV2Cluster200Response(apiVersion string, kind string, id string, spec ListKsqldbcmV2Clusters200ResponseAllOfDataInnerSpec, status KsqldbcmV2ClusterStatus, ) *GetKsqldbcmV2Cluster200Response`

NewGetKsqldbcmV2Cluster200Response instantiates a new GetKsqldbcmV2Cluster200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKsqldbcmV2Cluster200ResponseWithDefaults

`func NewGetKsqldbcmV2Cluster200ResponseWithDefaults() *GetKsqldbcmV2Cluster200Response`

NewGetKsqldbcmV2Cluster200ResponseWithDefaults instantiates a new GetKsqldbcmV2Cluster200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetKsqldbcmV2Cluster200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetKsqldbcmV2Cluster200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetKsqldbcmV2Cluster200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetKsqldbcmV2Cluster200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetKsqldbcmV2Cluster200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetKsqldbcmV2Cluster200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetKsqldbcmV2Cluster200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetKsqldbcmV2Cluster200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetKsqldbcmV2Cluster200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetKsqldbcmV2Cluster200Response) GetMetadata() KsqldbcmV2ClusterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetKsqldbcmV2Cluster200Response) GetMetadataOk() (*KsqldbcmV2ClusterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetKsqldbcmV2Cluster200Response) SetMetadata(v KsqldbcmV2ClusterMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetKsqldbcmV2Cluster200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *GetKsqldbcmV2Cluster200Response) GetSpec() ListKsqldbcmV2Clusters200ResponseAllOfDataInnerSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *GetKsqldbcmV2Cluster200Response) GetSpecOk() (*ListKsqldbcmV2Clusters200ResponseAllOfDataInnerSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *GetKsqldbcmV2Cluster200Response) SetSpec(v ListKsqldbcmV2Clusters200ResponseAllOfDataInnerSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *GetKsqldbcmV2Cluster200Response) GetStatus() KsqldbcmV2ClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetKsqldbcmV2Cluster200Response) GetStatusOk() (*KsqldbcmV2ClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetKsqldbcmV2Cluster200Response) SetStatus(v KsqldbcmV2ClusterStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


