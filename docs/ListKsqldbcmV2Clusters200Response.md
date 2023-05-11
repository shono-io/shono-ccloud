# ListKsqldbcmV2Clusters200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**KsqldbcmV2ClusterListMetadata**](KsqldbcmV2ClusterListMetadata.md) |  | 
**Data** | [**[]ListKsqldbcmV2Clusters200ResponseAllOfDataInner**](ListKsqldbcmV2Clusters200ResponseAllOfDataInner.md) |  | 

## Methods

### NewListKsqldbcmV2Clusters200Response

`func NewListKsqldbcmV2Clusters200Response(apiVersion string, kind string, metadata KsqldbcmV2ClusterListMetadata, data []ListKsqldbcmV2Clusters200ResponseAllOfDataInner, ) *ListKsqldbcmV2Clusters200Response`

NewListKsqldbcmV2Clusters200Response instantiates a new ListKsqldbcmV2Clusters200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKsqldbcmV2Clusters200ResponseWithDefaults

`func NewListKsqldbcmV2Clusters200ResponseWithDefaults() *ListKsqldbcmV2Clusters200Response`

NewListKsqldbcmV2Clusters200ResponseWithDefaults instantiates a new ListKsqldbcmV2Clusters200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListKsqldbcmV2Clusters200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListKsqldbcmV2Clusters200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListKsqldbcmV2Clusters200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ListKsqldbcmV2Clusters200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListKsqldbcmV2Clusters200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListKsqldbcmV2Clusters200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListKsqldbcmV2Clusters200Response) GetMetadata() KsqldbcmV2ClusterListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListKsqldbcmV2Clusters200Response) GetMetadataOk() (*KsqldbcmV2ClusterListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListKsqldbcmV2Clusters200Response) SetMetadata(v KsqldbcmV2ClusterListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListKsqldbcmV2Clusters200Response) GetData() []ListKsqldbcmV2Clusters200ResponseAllOfDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListKsqldbcmV2Clusters200Response) GetDataOk() (*[]ListKsqldbcmV2Clusters200ResponseAllOfDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListKsqldbcmV2Clusters200Response) SetData(v []ListKsqldbcmV2Clusters200ResponseAllOfDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


