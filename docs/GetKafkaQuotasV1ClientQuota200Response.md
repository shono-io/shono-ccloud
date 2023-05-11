# GetKafkaQuotasV1ClientQuota200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**KafkaQuotasV1ClientQuotaMetadata**](KafkaQuotasV1ClientQuotaMetadata.md) |  | [optional] 
**Spec** | [**ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInnerSpec**](ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInnerSpec.md) |  | 

## Methods

### NewGetKafkaQuotasV1ClientQuota200Response

`func NewGetKafkaQuotasV1ClientQuota200Response(apiVersion string, kind string, id string, spec ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInnerSpec, ) *GetKafkaQuotasV1ClientQuota200Response`

NewGetKafkaQuotasV1ClientQuota200Response instantiates a new GetKafkaQuotasV1ClientQuota200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKafkaQuotasV1ClientQuota200ResponseWithDefaults

`func NewGetKafkaQuotasV1ClientQuota200ResponseWithDefaults() *GetKafkaQuotasV1ClientQuota200Response`

NewGetKafkaQuotasV1ClientQuota200ResponseWithDefaults instantiates a new GetKafkaQuotasV1ClientQuota200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetKafkaQuotasV1ClientQuota200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetKafkaQuotasV1ClientQuota200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetKafkaQuotasV1ClientQuota200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetKafkaQuotasV1ClientQuota200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetKafkaQuotasV1ClientQuota200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetKafkaQuotasV1ClientQuota200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetKafkaQuotasV1ClientQuota200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetKafkaQuotasV1ClientQuota200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetKafkaQuotasV1ClientQuota200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetKafkaQuotasV1ClientQuota200Response) GetMetadata() KafkaQuotasV1ClientQuotaMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetKafkaQuotasV1ClientQuota200Response) GetMetadataOk() (*KafkaQuotasV1ClientQuotaMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetKafkaQuotasV1ClientQuota200Response) SetMetadata(v KafkaQuotasV1ClientQuotaMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetKafkaQuotasV1ClientQuota200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *GetKafkaQuotasV1ClientQuota200Response) GetSpec() ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInnerSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *GetKafkaQuotasV1ClientQuota200Response) GetSpecOk() (*ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInnerSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *GetKafkaQuotasV1ClientQuota200Response) SetSpec(v ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInnerSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


