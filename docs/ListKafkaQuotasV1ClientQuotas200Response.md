# ListKafkaQuotasV1ClientQuotas200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**KafkaQuotasV1ClientQuotaListMetadata**](KafkaQuotasV1ClientQuotaListMetadata.md) |  | 
**Data** | [**[]ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInner**](ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInner.md) |  | 

## Methods

### NewListKafkaQuotasV1ClientQuotas200Response

`func NewListKafkaQuotasV1ClientQuotas200Response(apiVersion string, kind string, metadata KafkaQuotasV1ClientQuotaListMetadata, data []ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInner, ) *ListKafkaQuotasV1ClientQuotas200Response`

NewListKafkaQuotasV1ClientQuotas200Response instantiates a new ListKafkaQuotasV1ClientQuotas200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKafkaQuotasV1ClientQuotas200ResponseWithDefaults

`func NewListKafkaQuotasV1ClientQuotas200ResponseWithDefaults() *ListKafkaQuotasV1ClientQuotas200Response`

NewListKafkaQuotasV1ClientQuotas200ResponseWithDefaults instantiates a new ListKafkaQuotasV1ClientQuotas200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListKafkaQuotasV1ClientQuotas200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListKafkaQuotasV1ClientQuotas200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListKafkaQuotasV1ClientQuotas200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ListKafkaQuotasV1ClientQuotas200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListKafkaQuotasV1ClientQuotas200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListKafkaQuotasV1ClientQuotas200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListKafkaQuotasV1ClientQuotas200Response) GetMetadata() KafkaQuotasV1ClientQuotaListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListKafkaQuotasV1ClientQuotas200Response) GetMetadataOk() (*KafkaQuotasV1ClientQuotaListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListKafkaQuotasV1ClientQuotas200Response) SetMetadata(v KafkaQuotasV1ClientQuotaListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListKafkaQuotasV1ClientQuotas200Response) GetData() []ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListKafkaQuotasV1ClientQuotas200Response) GetDataOk() (*[]ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListKafkaQuotasV1ClientQuotas200Response) SetData(v []ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


