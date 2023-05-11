# NetworkingV1TransitGatewayAttachmentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the TGW attachment | [optional] 
**Cloud** | Pointer to [**NetworkingV1TransitGatewayAttachmentSpecCloud**](NetworkingV1TransitGatewayAttachmentSpecCloud.md) |  | [optional] 
**Environment** | Pointer to [**NetworkingV1NetworkSpecEnvironment**](NetworkingV1NetworkSpecEnvironment.md) |  | [optional] 
**Network** | Pointer to [**NetworkingV1PeeringSpecNetwork**](NetworkingV1PeeringSpecNetwork.md) |  | [optional] 

## Methods

### NewNetworkingV1TransitGatewayAttachmentSpec

`func NewNetworkingV1TransitGatewayAttachmentSpec() *NetworkingV1TransitGatewayAttachmentSpec`

NewNetworkingV1TransitGatewayAttachmentSpec instantiates a new NetworkingV1TransitGatewayAttachmentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1TransitGatewayAttachmentSpecWithDefaults

`func NewNetworkingV1TransitGatewayAttachmentSpecWithDefaults() *NetworkingV1TransitGatewayAttachmentSpec`

NewNetworkingV1TransitGatewayAttachmentSpecWithDefaults instantiates a new NetworkingV1TransitGatewayAttachmentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1TransitGatewayAttachmentSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1TransitGatewayAttachmentSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetCloud() NetworkingV1TransitGatewayAttachmentSpecCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetCloudOk() (*NetworkingV1TransitGatewayAttachmentSpecCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingV1TransitGatewayAttachmentSpec) SetCloud(v NetworkingV1TransitGatewayAttachmentSpecCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingV1TransitGatewayAttachmentSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetEnvironment() NetworkingV1NetworkSpecEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetEnvironmentOk() (*NetworkingV1NetworkSpecEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1TransitGatewayAttachmentSpec) SetEnvironment(v NetworkingV1NetworkSpecEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1TransitGatewayAttachmentSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetNetwork

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetNetwork() NetworkingV1PeeringSpecNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetNetworkOk() (*NetworkingV1PeeringSpecNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkingV1TransitGatewayAttachmentSpec) SetNetwork(v NetworkingV1PeeringSpecNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkingV1TransitGatewayAttachmentSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


