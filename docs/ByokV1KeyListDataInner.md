# ByokV1KeyListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | [**ByokV1KeyMetadata**](ByokV1KeyMetadata.md) |  | 
**Key** | [**ByokV1KeyKey**](ByokV1KeyKey.md) |  | 
**Provider** | **string** | The cloud provider of the Key. | [readonly] 
**State** | **string** | The state of the key:   AVAILABLE: key can be used for a Kafka cluster provisioning   IN_USE: key is already in use by a Kafka cluster provisioning  | [readonly] 

## Methods

### NewByokV1KeyListDataInner

`func NewByokV1KeyListDataInner(id string, metadata ByokV1KeyMetadata, key ByokV1KeyKey, provider string, state string, ) *ByokV1KeyListDataInner`

NewByokV1KeyListDataInner instantiates a new ByokV1KeyListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByokV1KeyListDataInnerWithDefaults

`func NewByokV1KeyListDataInnerWithDefaults() *ByokV1KeyListDataInner`

NewByokV1KeyListDataInnerWithDefaults instantiates a new ByokV1KeyListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ByokV1KeyListDataInner) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ByokV1KeyListDataInner) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ByokV1KeyListDataInner) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ByokV1KeyListDataInner) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ByokV1KeyListDataInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ByokV1KeyListDataInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ByokV1KeyListDataInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ByokV1KeyListDataInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ByokV1KeyListDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ByokV1KeyListDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ByokV1KeyListDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *ByokV1KeyListDataInner) GetMetadata() ByokV1KeyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ByokV1KeyListDataInner) GetMetadataOk() (*ByokV1KeyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ByokV1KeyListDataInner) SetMetadata(v ByokV1KeyMetadata)`

SetMetadata sets Metadata field to given value.


### GetKey

`func (o *ByokV1KeyListDataInner) GetKey() ByokV1KeyKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ByokV1KeyListDataInner) GetKeyOk() (*ByokV1KeyKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ByokV1KeyListDataInner) SetKey(v ByokV1KeyKey)`

SetKey sets Key field to given value.


### GetProvider

`func (o *ByokV1KeyListDataInner) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ByokV1KeyListDataInner) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ByokV1KeyListDataInner) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetState

`func (o *ByokV1KeyListDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ByokV1KeyListDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ByokV1KeyListDataInner) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


