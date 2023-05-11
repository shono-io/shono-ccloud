# GetIamV2ApiKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**IamV2ApiKeyMetadata**](IamV2ApiKeyMetadata.md) |  | [optional] 
**Spec** | [**ListIamV2ApiKeys200ResponseAllOfDataInnerSpec**](ListIamV2ApiKeys200ResponseAllOfDataInnerSpec.md) |  | 

## Methods

### NewGetIamV2ApiKey200Response

`func NewGetIamV2ApiKey200Response(apiVersion string, kind string, id string, spec ListIamV2ApiKeys200ResponseAllOfDataInnerSpec, ) *GetIamV2ApiKey200Response`

NewGetIamV2ApiKey200Response instantiates a new GetIamV2ApiKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIamV2ApiKey200ResponseWithDefaults

`func NewGetIamV2ApiKey200ResponseWithDefaults() *GetIamV2ApiKey200Response`

NewGetIamV2ApiKey200ResponseWithDefaults instantiates a new GetIamV2ApiKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetIamV2ApiKey200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetIamV2ApiKey200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetIamV2ApiKey200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetIamV2ApiKey200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetIamV2ApiKey200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetIamV2ApiKey200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetIamV2ApiKey200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIamV2ApiKey200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIamV2ApiKey200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetIamV2ApiKey200Response) GetMetadata() IamV2ApiKeyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetIamV2ApiKey200Response) GetMetadataOk() (*IamV2ApiKeyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetIamV2ApiKey200Response) SetMetadata(v IamV2ApiKeyMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetIamV2ApiKey200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *GetIamV2ApiKey200Response) GetSpec() ListIamV2ApiKeys200ResponseAllOfDataInnerSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *GetIamV2ApiKey200Response) GetSpecOk() (*ListIamV2ApiKeys200ResponseAllOfDataInnerSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *GetIamV2ApiKey200Response) SetSpec(v ListIamV2ApiKeys200ResponseAllOfDataInnerSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


