# CreateNetworkingV1NetworkLinkEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**NetworkingV1NetworkLinkEndpointMetadata**](NetworkingV1NetworkLinkEndpointMetadata.md) |  | [optional] 
**Spec** | [**CreateNetworkingV1NetworkLinkEndpointRequestAllOf1Spec**](CreateNetworkingV1NetworkLinkEndpointRequestAllOf1Spec.md) |  | 
**Status** | Pointer to [**NetworkingV1NetworkLinkEndpointStatus**](NetworkingV1NetworkLinkEndpointStatus.md) |  | [optional] 

## Methods

### NewCreateNetworkingV1NetworkLinkEndpointRequest

`func NewCreateNetworkingV1NetworkLinkEndpointRequest(spec CreateNetworkingV1NetworkLinkEndpointRequestAllOf1Spec, ) *CreateNetworkingV1NetworkLinkEndpointRequest`

NewCreateNetworkingV1NetworkLinkEndpointRequest instantiates a new CreateNetworkingV1NetworkLinkEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkingV1NetworkLinkEndpointRequestWithDefaults

`func NewCreateNetworkingV1NetworkLinkEndpointRequestWithDefaults() *CreateNetworkingV1NetworkLinkEndpointRequest`

NewCreateNetworkingV1NetworkLinkEndpointRequestWithDefaults instantiates a new CreateNetworkingV1NetworkLinkEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) GetMetadata() NetworkingV1NetworkLinkEndpointMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) GetMetadataOk() (*NetworkingV1NetworkLinkEndpointMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) SetMetadata(v NetworkingV1NetworkLinkEndpointMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) GetSpec() CreateNetworkingV1NetworkLinkEndpointRequestAllOf1Spec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) GetSpecOk() (*CreateNetworkingV1NetworkLinkEndpointRequestAllOf1Spec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) SetSpec(v CreateNetworkingV1NetworkLinkEndpointRequestAllOf1Spec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) GetStatus() NetworkingV1NetworkLinkEndpointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) GetStatusOk() (*NetworkingV1NetworkLinkEndpointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) SetStatus(v NetworkingV1NetworkLinkEndpointStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateNetworkingV1NetworkLinkEndpointRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


