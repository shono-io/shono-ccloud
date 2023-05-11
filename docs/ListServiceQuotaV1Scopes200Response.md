# ListServiceQuotaV1Scopes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ServiceQuotaV1ScopeListMetadata**](ServiceQuotaV1ScopeListMetadata.md) |  | 
**Data** | [**[]ServiceQuotaV1ScopeListDataInner**](ServiceQuotaV1ScopeListDataInner.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewListServiceQuotaV1Scopes200Response

`func NewListServiceQuotaV1Scopes200Response(apiVersion string, kind string, metadata ServiceQuotaV1ScopeListMetadata, data []ServiceQuotaV1ScopeListDataInner, ) *ListServiceQuotaV1Scopes200Response`

NewListServiceQuotaV1Scopes200Response instantiates a new ListServiceQuotaV1Scopes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceQuotaV1Scopes200ResponseWithDefaults

`func NewListServiceQuotaV1Scopes200ResponseWithDefaults() *ListServiceQuotaV1Scopes200Response`

NewListServiceQuotaV1Scopes200ResponseWithDefaults instantiates a new ListServiceQuotaV1Scopes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListServiceQuotaV1Scopes200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListServiceQuotaV1Scopes200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListServiceQuotaV1Scopes200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ListServiceQuotaV1Scopes200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListServiceQuotaV1Scopes200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListServiceQuotaV1Scopes200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListServiceQuotaV1Scopes200Response) GetMetadata() ServiceQuotaV1ScopeListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListServiceQuotaV1Scopes200Response) GetMetadataOk() (*ServiceQuotaV1ScopeListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListServiceQuotaV1Scopes200Response) SetMetadata(v ServiceQuotaV1ScopeListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListServiceQuotaV1Scopes200Response) GetData() []ServiceQuotaV1ScopeListDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListServiceQuotaV1Scopes200Response) GetDataOk() (*[]ServiceQuotaV1ScopeListDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListServiceQuotaV1Scopes200Response) SetData(v []ServiceQuotaV1ScopeListDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


