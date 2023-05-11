# CreateIamV2ApiKey202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**IamV2ApiKeyMetadata**](IamV2ApiKeyMetadata.md) |  | [optional] 
**Spec** | [**ListIamV2ApiKeys200ResponseAllOfDataInnerSpec**](ListIamV2ApiKeys200ResponseAllOfDataInnerSpec.md) |  | 

## Methods

### NewCreateIamV2ApiKey202Response

`func NewCreateIamV2ApiKey202Response(spec ListIamV2ApiKeys200ResponseAllOfDataInnerSpec, ) *CreateIamV2ApiKey202Response`

NewCreateIamV2ApiKey202Response instantiates a new CreateIamV2ApiKey202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIamV2ApiKey202ResponseWithDefaults

`func NewCreateIamV2ApiKey202ResponseWithDefaults() *CreateIamV2ApiKey202Response`

NewCreateIamV2ApiKey202ResponseWithDefaults instantiates a new CreateIamV2ApiKey202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateIamV2ApiKey202Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateIamV2ApiKey202Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateIamV2ApiKey202Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateIamV2ApiKey202Response) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CreateIamV2ApiKey202Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateIamV2ApiKey202Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateIamV2ApiKey202Response) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateIamV2ApiKey202Response) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CreateIamV2ApiKey202Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateIamV2ApiKey202Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateIamV2ApiKey202Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateIamV2ApiKey202Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateIamV2ApiKey202Response) GetMetadata() IamV2ApiKeyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateIamV2ApiKey202Response) GetMetadataOk() (*IamV2ApiKeyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateIamV2ApiKey202Response) SetMetadata(v IamV2ApiKeyMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateIamV2ApiKey202Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *CreateIamV2ApiKey202Response) GetSpec() ListIamV2ApiKeys200ResponseAllOfDataInnerSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreateIamV2ApiKey202Response) GetSpecOk() (*ListIamV2ApiKeys200ResponseAllOfDataInnerSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreateIamV2ApiKey202Response) SetSpec(v ListIamV2ApiKeys200ResponseAllOfDataInnerSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

