# CreateNetworkingV1Network202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**NetworkingV1NetworkMetadata**](NetworkingV1NetworkMetadata.md) |  | [optional] 
**Spec** | [**ListNetworkingV1Networks200ResponseAllOfDataInnerSpec**](ListNetworkingV1Networks200ResponseAllOfDataInnerSpec.md) |  | 
**Status** | [**NetworkingV1NetworkStatus**](NetworkingV1NetworkStatus.md) |  | 

## Methods

### NewCreateNetworkingV1Network202Response

`func NewCreateNetworkingV1Network202Response(spec ListNetworkingV1Networks200ResponseAllOfDataInnerSpec, status NetworkingV1NetworkStatus, ) *CreateNetworkingV1Network202Response`

NewCreateNetworkingV1Network202Response instantiates a new CreateNetworkingV1Network202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkingV1Network202ResponseWithDefaults

`func NewCreateNetworkingV1Network202ResponseWithDefaults() *CreateNetworkingV1Network202Response`

NewCreateNetworkingV1Network202ResponseWithDefaults instantiates a new CreateNetworkingV1Network202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateNetworkingV1Network202Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateNetworkingV1Network202Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateNetworkingV1Network202Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateNetworkingV1Network202Response) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CreateNetworkingV1Network202Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateNetworkingV1Network202Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateNetworkingV1Network202Response) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateNetworkingV1Network202Response) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CreateNetworkingV1Network202Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNetworkingV1Network202Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNetworkingV1Network202Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNetworkingV1Network202Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateNetworkingV1Network202Response) GetMetadata() NetworkingV1NetworkMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateNetworkingV1Network202Response) GetMetadataOk() (*NetworkingV1NetworkMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateNetworkingV1Network202Response) SetMetadata(v NetworkingV1NetworkMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateNetworkingV1Network202Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *CreateNetworkingV1Network202Response) GetSpec() ListNetworkingV1Networks200ResponseAllOfDataInnerSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreateNetworkingV1Network202Response) GetSpecOk() (*ListNetworkingV1Networks200ResponseAllOfDataInnerSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreateNetworkingV1Network202Response) SetSpec(v ListNetworkingV1Networks200ResponseAllOfDataInnerSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *CreateNetworkingV1Network202Response) GetStatus() NetworkingV1NetworkStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateNetworkingV1Network202Response) GetStatusOk() (*NetworkingV1NetworkStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateNetworkingV1Network202Response) SetStatus(v NetworkingV1NetworkStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


