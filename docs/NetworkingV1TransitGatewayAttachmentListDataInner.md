# NetworkingV1TransitGatewayAttachmentListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | [**NetworkingV1TransitGatewayAttachmentMetadata**](NetworkingV1TransitGatewayAttachmentMetadata.md) |  | 
**Spec** | **map[string]interface{}** |  | 
**Status** | [**NetworkingV1TransitGatewayAttachmentStatus**](NetworkingV1TransitGatewayAttachmentStatus.md) |  | 

## Methods

### NewNetworkingV1TransitGatewayAttachmentListDataInner

`func NewNetworkingV1TransitGatewayAttachmentListDataInner(id string, metadata NetworkingV1TransitGatewayAttachmentMetadata, spec map[string]interface{}, status NetworkingV1TransitGatewayAttachmentStatus, ) *NetworkingV1TransitGatewayAttachmentListDataInner`

NewNetworkingV1TransitGatewayAttachmentListDataInner instantiates a new NetworkingV1TransitGatewayAttachmentListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1TransitGatewayAttachmentListDataInnerWithDefaults

`func NewNetworkingV1TransitGatewayAttachmentListDataInnerWithDefaults() *NetworkingV1TransitGatewayAttachmentListDataInner`

NewNetworkingV1TransitGatewayAttachmentListDataInnerWithDefaults instantiates a new NetworkingV1TransitGatewayAttachmentListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) GetMetadata() NetworkingV1TransitGatewayAttachmentMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) GetMetadataOk() (*NetworkingV1TransitGatewayAttachmentMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) SetMetadata(v NetworkingV1TransitGatewayAttachmentMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) GetStatus() NetworkingV1TransitGatewayAttachmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) GetStatusOk() (*NetworkingV1TransitGatewayAttachmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkingV1TransitGatewayAttachmentListDataInner) SetStatus(v NetworkingV1TransitGatewayAttachmentStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


