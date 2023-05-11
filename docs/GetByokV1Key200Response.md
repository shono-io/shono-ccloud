# GetByokV1Key200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**ByokV1KeyMetadata**](ByokV1KeyMetadata.md) |  | [optional] 
**Key** | [**ByokV1KeyKey**](ByokV1KeyKey.md) |  | 
**Provider** | **string** | The cloud provider of the Key. | [readonly] 
**State** | **string** | The state of the key:   AVAILABLE: key can be used for a Kafka cluster provisioning   IN_USE: key is already in use by a Kafka cluster provisioning  | [readonly] 

## Methods

### NewGetByokV1Key200Response

`func NewGetByokV1Key200Response(apiVersion string, kind string, id string, key ByokV1KeyKey, provider string, state string, ) *GetByokV1Key200Response`

NewGetByokV1Key200Response instantiates a new GetByokV1Key200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetByokV1Key200ResponseWithDefaults

`func NewGetByokV1Key200ResponseWithDefaults() *GetByokV1Key200Response`

NewGetByokV1Key200ResponseWithDefaults instantiates a new GetByokV1Key200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetByokV1Key200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetByokV1Key200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetByokV1Key200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetByokV1Key200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetByokV1Key200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetByokV1Key200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetByokV1Key200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetByokV1Key200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetByokV1Key200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetByokV1Key200Response) GetMetadata() ByokV1KeyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetByokV1Key200Response) GetMetadataOk() (*ByokV1KeyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetByokV1Key200Response) SetMetadata(v ByokV1KeyMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetByokV1Key200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetKey

`func (o *GetByokV1Key200Response) GetKey() ByokV1KeyKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetByokV1Key200Response) GetKeyOk() (*ByokV1KeyKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetByokV1Key200Response) SetKey(v ByokV1KeyKey)`

SetKey sets Key field to given value.


### GetProvider

`func (o *GetByokV1Key200Response) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GetByokV1Key200Response) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GetByokV1Key200Response) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetState

`func (o *GetByokV1Key200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetByokV1Key200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetByokV1Key200Response) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


