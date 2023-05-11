# NetworkingV1PeeringSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the peering | [optional] 
**Cloud** | Pointer to [**NetworkingV1PeeringSpecCloud**](NetworkingV1PeeringSpecCloud.md) |  | [optional] 
**Environment** | Pointer to [**NetworkingV1NetworkSpecEnvironment**](NetworkingV1NetworkSpecEnvironment.md) |  | [optional] 
**Network** | Pointer to [**NetworkingV1PeeringSpecNetwork**](NetworkingV1PeeringSpecNetwork.md) |  | [optional] 

## Methods

### NewNetworkingV1PeeringSpec

`func NewNetworkingV1PeeringSpec() *NetworkingV1PeeringSpec`

NewNetworkingV1PeeringSpec instantiates a new NetworkingV1PeeringSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PeeringSpecWithDefaults

`func NewNetworkingV1PeeringSpecWithDefaults() *NetworkingV1PeeringSpec`

NewNetworkingV1PeeringSpecWithDefaults instantiates a new NetworkingV1PeeringSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1PeeringSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1PeeringSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1PeeringSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1PeeringSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingV1PeeringSpec) GetCloud() NetworkingV1PeeringSpecCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingV1PeeringSpec) GetCloudOk() (*NetworkingV1PeeringSpecCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingV1PeeringSpec) SetCloud(v NetworkingV1PeeringSpecCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingV1PeeringSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1PeeringSpec) GetEnvironment() NetworkingV1NetworkSpecEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1PeeringSpec) GetEnvironmentOk() (*NetworkingV1NetworkSpecEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1PeeringSpec) SetEnvironment(v NetworkingV1NetworkSpecEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1PeeringSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetNetwork

`func (o *NetworkingV1PeeringSpec) GetNetwork() NetworkingV1PeeringSpecNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkingV1PeeringSpec) GetNetworkOk() (*NetworkingV1PeeringSpecNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkingV1PeeringSpec) SetNetwork(v NetworkingV1PeeringSpecNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkingV1PeeringSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


