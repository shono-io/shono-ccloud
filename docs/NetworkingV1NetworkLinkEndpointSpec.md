# NetworkingV1NetworkLinkEndpointSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the network link endpoint | [optional] 
**Description** | Pointer to **string** | The description of the network link endpoint | [optional] 
**Environment** | Pointer to [**SrcmV2ClusterSpecEnvironment**](SrcmV2ClusterSpecEnvironment.md) |  | [optional] 
**Network** | Pointer to [**NetworkingV1NetworkLinkServiceSpecNetwork**](NetworkingV1NetworkLinkServiceSpecNetwork.md) |  | [optional] 
**NetworkLinkService** | Pointer to [**NetworkingV1NetworkLinkEndpointSpecNetworkLinkService**](NetworkingV1NetworkLinkEndpointSpecNetworkLinkService.md) |  | [optional] 

## Methods

### NewNetworkingV1NetworkLinkEndpointSpec

`func NewNetworkingV1NetworkLinkEndpointSpec() *NetworkingV1NetworkLinkEndpointSpec`

NewNetworkingV1NetworkLinkEndpointSpec instantiates a new NetworkingV1NetworkLinkEndpointSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkLinkEndpointSpecWithDefaults

`func NewNetworkingV1NetworkLinkEndpointSpecWithDefaults() *NetworkingV1NetworkLinkEndpointSpec`

NewNetworkingV1NetworkLinkEndpointSpecWithDefaults instantiates a new NetworkingV1NetworkLinkEndpointSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1NetworkLinkEndpointSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1NetworkLinkEndpointSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1NetworkLinkEndpointSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1NetworkLinkEndpointSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkingV1NetworkLinkEndpointSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkingV1NetworkLinkEndpointSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkingV1NetworkLinkEndpointSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkingV1NetworkLinkEndpointSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1NetworkLinkEndpointSpec) GetEnvironment() SrcmV2ClusterSpecEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1NetworkLinkEndpointSpec) GetEnvironmentOk() (*SrcmV2ClusterSpecEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1NetworkLinkEndpointSpec) SetEnvironment(v SrcmV2ClusterSpecEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1NetworkLinkEndpointSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetNetwork

`func (o *NetworkingV1NetworkLinkEndpointSpec) GetNetwork() NetworkingV1NetworkLinkServiceSpecNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkingV1NetworkLinkEndpointSpec) GetNetworkOk() (*NetworkingV1NetworkLinkServiceSpecNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkingV1NetworkLinkEndpointSpec) SetNetwork(v NetworkingV1NetworkLinkServiceSpecNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkingV1NetworkLinkEndpointSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkLinkService

`func (o *NetworkingV1NetworkLinkEndpointSpec) GetNetworkLinkService() NetworkingV1NetworkLinkEndpointSpecNetworkLinkService`

GetNetworkLinkService returns the NetworkLinkService field if non-nil, zero value otherwise.

### GetNetworkLinkServiceOk

`func (o *NetworkingV1NetworkLinkEndpointSpec) GetNetworkLinkServiceOk() (*NetworkingV1NetworkLinkEndpointSpecNetworkLinkService, bool)`

GetNetworkLinkServiceOk returns a tuple with the NetworkLinkService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLinkService

`func (o *NetworkingV1NetworkLinkEndpointSpec) SetNetworkLinkService(v NetworkingV1NetworkLinkEndpointSpecNetworkLinkService)`

SetNetworkLinkService sets NetworkLinkService field to given value.

### HasNetworkLinkService

`func (o *NetworkingV1NetworkLinkEndpointSpec) HasNetworkLinkService() bool`

HasNetworkLinkService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


