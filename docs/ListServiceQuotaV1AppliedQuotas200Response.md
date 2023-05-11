# ListServiceQuotaV1AppliedQuotas200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ServiceQuotaV1AppliedQuotaListMetadata**](ServiceQuotaV1AppliedQuotaListMetadata.md) |  | 
**Data** | [**[]ListServiceQuotaV1AppliedQuotas200ResponseAllOfDataInner**](ListServiceQuotaV1AppliedQuotas200ResponseAllOfDataInner.md) |  | 

## Methods

### NewListServiceQuotaV1AppliedQuotas200Response

`func NewListServiceQuotaV1AppliedQuotas200Response(apiVersion string, kind string, metadata ServiceQuotaV1AppliedQuotaListMetadata, data []ListServiceQuotaV1AppliedQuotas200ResponseAllOfDataInner, ) *ListServiceQuotaV1AppliedQuotas200Response`

NewListServiceQuotaV1AppliedQuotas200Response instantiates a new ListServiceQuotaV1AppliedQuotas200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceQuotaV1AppliedQuotas200ResponseWithDefaults

`func NewListServiceQuotaV1AppliedQuotas200ResponseWithDefaults() *ListServiceQuotaV1AppliedQuotas200Response`

NewListServiceQuotaV1AppliedQuotas200ResponseWithDefaults instantiates a new ListServiceQuotaV1AppliedQuotas200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListServiceQuotaV1AppliedQuotas200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListServiceQuotaV1AppliedQuotas200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListServiceQuotaV1AppliedQuotas200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ListServiceQuotaV1AppliedQuotas200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListServiceQuotaV1AppliedQuotas200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListServiceQuotaV1AppliedQuotas200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListServiceQuotaV1AppliedQuotas200Response) GetMetadata() ServiceQuotaV1AppliedQuotaListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListServiceQuotaV1AppliedQuotas200Response) GetMetadataOk() (*ServiceQuotaV1AppliedQuotaListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListServiceQuotaV1AppliedQuotas200Response) SetMetadata(v ServiceQuotaV1AppliedQuotaListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListServiceQuotaV1AppliedQuotas200Response) GetData() []ListServiceQuotaV1AppliedQuotas200ResponseAllOfDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListServiceQuotaV1AppliedQuotas200Response) GetDataOk() (*[]ListServiceQuotaV1AppliedQuotas200ResponseAllOfDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListServiceQuotaV1AppliedQuotas200Response) SetData(v []ListServiceQuotaV1AppliedQuotas200ResponseAllOfDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


