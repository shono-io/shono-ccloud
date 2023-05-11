# ListByokV1Keys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ByokV1KeyListMetadata**](ByokV1KeyListMetadata.md) |  | 
**Data** | [**[]ByokV1KeyListDataInner**](ByokV1KeyListDataInner.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewListByokV1Keys200Response

`func NewListByokV1Keys200Response(apiVersion string, kind string, metadata ByokV1KeyListMetadata, data []ByokV1KeyListDataInner, ) *ListByokV1Keys200Response`

NewListByokV1Keys200Response instantiates a new ListByokV1Keys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListByokV1Keys200ResponseWithDefaults

`func NewListByokV1Keys200ResponseWithDefaults() *ListByokV1Keys200Response`

NewListByokV1Keys200ResponseWithDefaults instantiates a new ListByokV1Keys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListByokV1Keys200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListByokV1Keys200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListByokV1Keys200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ListByokV1Keys200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListByokV1Keys200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListByokV1Keys200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListByokV1Keys200Response) GetMetadata() ByokV1KeyListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListByokV1Keys200Response) GetMetadataOk() (*ByokV1KeyListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListByokV1Keys200Response) SetMetadata(v ByokV1KeyListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListByokV1Keys200Response) GetData() []ByokV1KeyListDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListByokV1Keys200Response) GetDataOk() (*[]ByokV1KeyListDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListByokV1Keys200Response) SetData(v []ByokV1KeyListDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


