# GetNetworkingV1PrivateLinkAccess200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**NetworkingV1PrivateLinkAccessMetadata**](NetworkingV1PrivateLinkAccessMetadata.md) |  | [optional] 
**Spec** | [**ListNetworkingV1Peerings200ResponseAllOfDataInnerSpec**](ListNetworkingV1Peerings200ResponseAllOfDataInnerSpec.md) |  | 
**Status** | [**NetworkingV1PrivateLinkAccessStatus**](NetworkingV1PrivateLinkAccessStatus.md) |  | 

## Methods

### NewGetNetworkingV1PrivateLinkAccess200Response

`func NewGetNetworkingV1PrivateLinkAccess200Response(apiVersion string, kind string, id string, spec ListNetworkingV1Peerings200ResponseAllOfDataInnerSpec, status NetworkingV1PrivateLinkAccessStatus, ) *GetNetworkingV1PrivateLinkAccess200Response`

NewGetNetworkingV1PrivateLinkAccess200Response instantiates a new GetNetworkingV1PrivateLinkAccess200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkingV1PrivateLinkAccess200ResponseWithDefaults

`func NewGetNetworkingV1PrivateLinkAccess200ResponseWithDefaults() *GetNetworkingV1PrivateLinkAccess200Response`

NewGetNetworkingV1PrivateLinkAccess200ResponseWithDefaults instantiates a new GetNetworkingV1PrivateLinkAccess200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetNetworkingV1PrivateLinkAccess200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetNetworkingV1PrivateLinkAccess200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetNetworkingV1PrivateLinkAccess200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetNetworkingV1PrivateLinkAccess200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetNetworkingV1PrivateLinkAccess200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetNetworkingV1PrivateLinkAccess200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetNetworkingV1PrivateLinkAccess200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkingV1PrivateLinkAccess200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkingV1PrivateLinkAccess200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetNetworkingV1PrivateLinkAccess200Response) GetMetadata() NetworkingV1PrivateLinkAccessMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetNetworkingV1PrivateLinkAccess200Response) GetMetadataOk() (*NetworkingV1PrivateLinkAccessMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetNetworkingV1PrivateLinkAccess200Response) SetMetadata(v NetworkingV1PrivateLinkAccessMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetNetworkingV1PrivateLinkAccess200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *GetNetworkingV1PrivateLinkAccess200Response) GetSpec() ListNetworkingV1Peerings200ResponseAllOfDataInnerSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *GetNetworkingV1PrivateLinkAccess200Response) GetSpecOk() (*ListNetworkingV1Peerings200ResponseAllOfDataInnerSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *GetNetworkingV1PrivateLinkAccess200Response) SetSpec(v ListNetworkingV1Peerings200ResponseAllOfDataInnerSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *GetNetworkingV1PrivateLinkAccess200Response) GetStatus() NetworkingV1PrivateLinkAccessStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNetworkingV1PrivateLinkAccess200Response) GetStatusOk() (*NetworkingV1PrivateLinkAccessStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNetworkingV1PrivateLinkAccess200Response) SetStatus(v NetworkingV1PrivateLinkAccessStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


