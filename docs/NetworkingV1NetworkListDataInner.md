# NetworkingV1NetworkListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | [**NetworkingV1NetworkMetadata**](NetworkingV1NetworkMetadata.md) |  | 
**Spec** | **map[string]interface{}** |  | 
**Status** | [**NetworkingV1NetworkStatus**](NetworkingV1NetworkStatus.md) |  | 

## Methods

### NewNetworkingV1NetworkListDataInner

`func NewNetworkingV1NetworkListDataInner(id string, metadata NetworkingV1NetworkMetadata, spec map[string]interface{}, status NetworkingV1NetworkStatus, ) *NetworkingV1NetworkListDataInner`

NewNetworkingV1NetworkListDataInner instantiates a new NetworkingV1NetworkListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkListDataInnerWithDefaults

`func NewNetworkingV1NetworkListDataInnerWithDefaults() *NetworkingV1NetworkListDataInner`

NewNetworkingV1NetworkListDataInnerWithDefaults instantiates a new NetworkingV1NetworkListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingV1NetworkListDataInner) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingV1NetworkListDataInner) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingV1NetworkListDataInner) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkingV1NetworkListDataInner) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkingV1NetworkListDataInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1NetworkListDataInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1NetworkListDataInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkingV1NetworkListDataInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *NetworkingV1NetworkListDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkingV1NetworkListDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkingV1NetworkListDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *NetworkingV1NetworkListDataInner) GetMetadata() NetworkingV1NetworkMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingV1NetworkListDataInner) GetMetadataOk() (*NetworkingV1NetworkMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingV1NetworkListDataInner) SetMetadata(v NetworkingV1NetworkMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *NetworkingV1NetworkListDataInner) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkingV1NetworkListDataInner) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkingV1NetworkListDataInner) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *NetworkingV1NetworkListDataInner) GetStatus() NetworkingV1NetworkStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkingV1NetworkListDataInner) GetStatusOk() (*NetworkingV1NetworkStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkingV1NetworkListDataInner) SetStatus(v NetworkingV1NetworkStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


