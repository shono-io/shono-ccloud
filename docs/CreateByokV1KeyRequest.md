# CreateByokV1KeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ByokV1KeyMetadata**](ByokV1KeyMetadata.md) |  | [optional] 
**Key** | [**ByokV1KeyKey**](ByokV1KeyKey.md) |  | 
**Provider** | Pointer to **string** | The cloud provider of the Key. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the key:   AVAILABLE: key can be used for a Kafka cluster provisioning   IN_USE: key is already in use by a Kafka cluster provisioning  | [optional] [readonly] 

## Methods

### NewCreateByokV1KeyRequest

`func NewCreateByokV1KeyRequest(key ByokV1KeyKey, ) *CreateByokV1KeyRequest`

NewCreateByokV1KeyRequest instantiates a new CreateByokV1KeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateByokV1KeyRequestWithDefaults

`func NewCreateByokV1KeyRequestWithDefaults() *CreateByokV1KeyRequest`

NewCreateByokV1KeyRequestWithDefaults instantiates a new CreateByokV1KeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateByokV1KeyRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateByokV1KeyRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateByokV1KeyRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateByokV1KeyRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CreateByokV1KeyRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateByokV1KeyRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateByokV1KeyRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateByokV1KeyRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CreateByokV1KeyRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateByokV1KeyRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateByokV1KeyRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateByokV1KeyRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateByokV1KeyRequest) GetMetadata() ByokV1KeyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateByokV1KeyRequest) GetMetadataOk() (*ByokV1KeyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateByokV1KeyRequest) SetMetadata(v ByokV1KeyMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateByokV1KeyRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetKey

`func (o *CreateByokV1KeyRequest) GetKey() ByokV1KeyKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateByokV1KeyRequest) GetKeyOk() (*ByokV1KeyKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateByokV1KeyRequest) SetKey(v ByokV1KeyKey)`

SetKey sets Key field to given value.


### GetProvider

`func (o *CreateByokV1KeyRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateByokV1KeyRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateByokV1KeyRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CreateByokV1KeyRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetState

`func (o *CreateByokV1KeyRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateByokV1KeyRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateByokV1KeyRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateByokV1KeyRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


