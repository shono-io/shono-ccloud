# ListNetworkingV1NetworkLinkServices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**NetworkingV1NetworkLinkServiceListMetadata**](NetworkingV1NetworkLinkServiceListMetadata.md) |  | 
**Data** | [**[]ListNetworkingV1Peerings200ResponseAllOfDataInner**](ListNetworkingV1Peerings200ResponseAllOfDataInner.md) |  | 

## Methods

### NewListNetworkingV1NetworkLinkServices200Response

`func NewListNetworkingV1NetworkLinkServices200Response(apiVersion string, kind string, metadata NetworkingV1NetworkLinkServiceListMetadata, data []ListNetworkingV1Peerings200ResponseAllOfDataInner, ) *ListNetworkingV1NetworkLinkServices200Response`

NewListNetworkingV1NetworkLinkServices200Response instantiates a new ListNetworkingV1NetworkLinkServices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkingV1NetworkLinkServices200ResponseWithDefaults

`func NewListNetworkingV1NetworkLinkServices200ResponseWithDefaults() *ListNetworkingV1NetworkLinkServices200Response`

NewListNetworkingV1NetworkLinkServices200ResponseWithDefaults instantiates a new ListNetworkingV1NetworkLinkServices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListNetworkingV1NetworkLinkServices200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListNetworkingV1NetworkLinkServices200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListNetworkingV1NetworkLinkServices200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ListNetworkingV1NetworkLinkServices200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListNetworkingV1NetworkLinkServices200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListNetworkingV1NetworkLinkServices200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListNetworkingV1NetworkLinkServices200Response) GetMetadata() NetworkingV1NetworkLinkServiceListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListNetworkingV1NetworkLinkServices200Response) GetMetadataOk() (*NetworkingV1NetworkLinkServiceListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListNetworkingV1NetworkLinkServices200Response) SetMetadata(v NetworkingV1NetworkLinkServiceListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListNetworkingV1NetworkLinkServices200Response) GetData() []ListNetworkingV1Peerings200ResponseAllOfDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListNetworkingV1NetworkLinkServices200Response) GetDataOk() (*[]ListNetworkingV1Peerings200ResponseAllOfDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListNetworkingV1NetworkLinkServices200Response) SetData(v []ListNetworkingV1Peerings200ResponseAllOfDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


