# CreateNetworkingV1NetworkLinkEndpoint202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**NetworkingV1NetworkLinkEndpointMetadata**](NetworkingV1NetworkLinkEndpointMetadata.md) |  | [optional] 
**Spec** | [**ListNetworkingV1NetworkLinkEndpoints200ResponseAllOfDataInnerSpec**](ListNetworkingV1NetworkLinkEndpoints200ResponseAllOfDataInnerSpec.md) |  | 
**Status** | [**NetworkingV1NetworkLinkEndpointStatus**](NetworkingV1NetworkLinkEndpointStatus.md) |  | 

## Methods

### NewCreateNetworkingV1NetworkLinkEndpoint202Response

`func NewCreateNetworkingV1NetworkLinkEndpoint202Response(spec ListNetworkingV1NetworkLinkEndpoints200ResponseAllOfDataInnerSpec, status NetworkingV1NetworkLinkEndpointStatus, ) *CreateNetworkingV1NetworkLinkEndpoint202Response`

NewCreateNetworkingV1NetworkLinkEndpoint202Response instantiates a new CreateNetworkingV1NetworkLinkEndpoint202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkingV1NetworkLinkEndpoint202ResponseWithDefaults

`func NewCreateNetworkingV1NetworkLinkEndpoint202ResponseWithDefaults() *CreateNetworkingV1NetworkLinkEndpoint202Response`

NewCreateNetworkingV1NetworkLinkEndpoint202ResponseWithDefaults instantiates a new CreateNetworkingV1NetworkLinkEndpoint202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) GetMetadata() NetworkingV1NetworkLinkEndpointMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) GetMetadataOk() (*NetworkingV1NetworkLinkEndpointMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) SetMetadata(v NetworkingV1NetworkLinkEndpointMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) GetSpec() ListNetworkingV1NetworkLinkEndpoints200ResponseAllOfDataInnerSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) GetSpecOk() (*ListNetworkingV1NetworkLinkEndpoints200ResponseAllOfDataInnerSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) SetSpec(v ListNetworkingV1NetworkLinkEndpoints200ResponseAllOfDataInnerSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) GetStatus() NetworkingV1NetworkLinkEndpointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) GetStatusOk() (*NetworkingV1NetworkLinkEndpointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateNetworkingV1NetworkLinkEndpoint202Response) SetStatus(v NetworkingV1NetworkLinkEndpointStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


