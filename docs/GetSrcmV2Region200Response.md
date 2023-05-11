# GetSrcmV2Region200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**SrcmV2RegionMetadata**](SrcmV2RegionMetadata.md) |  | [optional] 
**Spec** | **map[string]interface{}** |  | 

## Methods

### NewGetSrcmV2Region200Response

`func NewGetSrcmV2Region200Response(apiVersion string, kind string, id string, spec map[string]interface{}, ) *GetSrcmV2Region200Response`

NewGetSrcmV2Region200Response instantiates a new GetSrcmV2Region200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSrcmV2Region200ResponseWithDefaults

`func NewGetSrcmV2Region200ResponseWithDefaults() *GetSrcmV2Region200Response`

NewGetSrcmV2Region200ResponseWithDefaults instantiates a new GetSrcmV2Region200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetSrcmV2Region200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetSrcmV2Region200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetSrcmV2Region200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetSrcmV2Region200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetSrcmV2Region200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetSrcmV2Region200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetSrcmV2Region200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSrcmV2Region200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSrcmV2Region200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetSrcmV2Region200Response) GetMetadata() SrcmV2RegionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetSrcmV2Region200Response) GetMetadataOk() (*SrcmV2RegionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetSrcmV2Region200Response) SetMetadata(v SrcmV2RegionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetSrcmV2Region200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *GetSrcmV2Region200Response) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *GetSrcmV2Region200Response) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *GetSrcmV2Region200Response) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


