# UpdateKafkaQuotasV1ClientQuotaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**KafkaQuotasV1ClientQuotaMetadata**](KafkaQuotasV1ClientQuotaMetadata.md) |  | [optional] 
**Spec** | [**UpdateCmkV2ClusterRequestAllOfSpec**](UpdateCmkV2ClusterRequestAllOfSpec.md) |  | 

## Methods

### NewUpdateKafkaQuotasV1ClientQuotaRequest

`func NewUpdateKafkaQuotasV1ClientQuotaRequest(spec UpdateCmkV2ClusterRequestAllOfSpec, ) *UpdateKafkaQuotasV1ClientQuotaRequest`

NewUpdateKafkaQuotasV1ClientQuotaRequest instantiates a new UpdateKafkaQuotasV1ClientQuotaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateKafkaQuotasV1ClientQuotaRequestWithDefaults

`func NewUpdateKafkaQuotasV1ClientQuotaRequestWithDefaults() *UpdateKafkaQuotasV1ClientQuotaRequest`

NewUpdateKafkaQuotasV1ClientQuotaRequestWithDefaults instantiates a new UpdateKafkaQuotasV1ClientQuotaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) GetMetadata() KafkaQuotasV1ClientQuotaMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) GetMetadataOk() (*KafkaQuotasV1ClientQuotaMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) SetMetadata(v KafkaQuotasV1ClientQuotaMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) GetSpec() UpdateCmkV2ClusterRequestAllOfSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) GetSpecOk() (*UpdateCmkV2ClusterRequestAllOfSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UpdateKafkaQuotasV1ClientQuotaRequest) SetSpec(v UpdateCmkV2ClusterRequestAllOfSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


