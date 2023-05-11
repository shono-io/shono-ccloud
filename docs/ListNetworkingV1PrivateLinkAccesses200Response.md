# ListNetworkingV1PrivateLinkAccesses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**NetworkingV1PrivateLinkAccessListMetadata**](NetworkingV1PrivateLinkAccessListMetadata.md) |  | 
**Data** | [**[]ListNetworkingV1Peerings200ResponseAllOfDataInner**](ListNetworkingV1Peerings200ResponseAllOfDataInner.md) |  | 

## Methods

### NewListNetworkingV1PrivateLinkAccesses200Response

`func NewListNetworkingV1PrivateLinkAccesses200Response(apiVersion string, kind string, metadata NetworkingV1PrivateLinkAccessListMetadata, data []ListNetworkingV1Peerings200ResponseAllOfDataInner, ) *ListNetworkingV1PrivateLinkAccesses200Response`

NewListNetworkingV1PrivateLinkAccesses200Response instantiates a new ListNetworkingV1PrivateLinkAccesses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkingV1PrivateLinkAccesses200ResponseWithDefaults

`func NewListNetworkingV1PrivateLinkAccesses200ResponseWithDefaults() *ListNetworkingV1PrivateLinkAccesses200Response`

NewListNetworkingV1PrivateLinkAccesses200ResponseWithDefaults instantiates a new ListNetworkingV1PrivateLinkAccesses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListNetworkingV1PrivateLinkAccesses200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListNetworkingV1PrivateLinkAccesses200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListNetworkingV1PrivateLinkAccesses200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ListNetworkingV1PrivateLinkAccesses200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListNetworkingV1PrivateLinkAccesses200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListNetworkingV1PrivateLinkAccesses200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListNetworkingV1PrivateLinkAccesses200Response) GetMetadata() NetworkingV1PrivateLinkAccessListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListNetworkingV1PrivateLinkAccesses200Response) GetMetadataOk() (*NetworkingV1PrivateLinkAccessListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListNetworkingV1PrivateLinkAccesses200Response) SetMetadata(v NetworkingV1PrivateLinkAccessListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListNetworkingV1PrivateLinkAccesses200Response) GetData() []ListNetworkingV1Peerings200ResponseAllOfDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListNetworkingV1PrivateLinkAccesses200Response) GetDataOk() (*[]ListNetworkingV1Peerings200ResponseAllOfDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListNetworkingV1PrivateLinkAccesses200Response) SetData(v []ListNetworkingV1Peerings200ResponseAllOfDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


