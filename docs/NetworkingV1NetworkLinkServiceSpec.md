# NetworkingV1NetworkLinkServiceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the network link service | [optional] 
**Description** | Pointer to **string** | The description of the network link service | [optional] 
**Accept** | Pointer to [**NetworkingV1NetworkLinkServiceSpecAccept**](NetworkingV1NetworkLinkServiceSpecAccept.md) |  | [optional] 
**Environment** | Pointer to [**SrcmV2ClusterSpecEnvironment**](SrcmV2ClusterSpecEnvironment.md) |  | [optional] 
**Network** | Pointer to [**NetworkingV1NetworkLinkServiceSpecNetwork**](NetworkingV1NetworkLinkServiceSpecNetwork.md) |  | [optional] 

## Methods

### NewNetworkingV1NetworkLinkServiceSpec

`func NewNetworkingV1NetworkLinkServiceSpec() *NetworkingV1NetworkLinkServiceSpec`

NewNetworkingV1NetworkLinkServiceSpec instantiates a new NetworkingV1NetworkLinkServiceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkLinkServiceSpecWithDefaults

`func NewNetworkingV1NetworkLinkServiceSpecWithDefaults() *NetworkingV1NetworkLinkServiceSpec`

NewNetworkingV1NetworkLinkServiceSpecWithDefaults instantiates a new NetworkingV1NetworkLinkServiceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1NetworkLinkServiceSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1NetworkLinkServiceSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1NetworkLinkServiceSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1NetworkLinkServiceSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkingV1NetworkLinkServiceSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkingV1NetworkLinkServiceSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkingV1NetworkLinkServiceSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkingV1NetworkLinkServiceSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccept

`func (o *NetworkingV1NetworkLinkServiceSpec) GetAccept() NetworkingV1NetworkLinkServiceSpecAccept`

GetAccept returns the Accept field if non-nil, zero value otherwise.

### GetAcceptOk

`func (o *NetworkingV1NetworkLinkServiceSpec) GetAcceptOk() (*NetworkingV1NetworkLinkServiceSpecAccept, bool)`

GetAcceptOk returns a tuple with the Accept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccept

`func (o *NetworkingV1NetworkLinkServiceSpec) SetAccept(v NetworkingV1NetworkLinkServiceSpecAccept)`

SetAccept sets Accept field to given value.

### HasAccept

`func (o *NetworkingV1NetworkLinkServiceSpec) HasAccept() bool`

HasAccept returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1NetworkLinkServiceSpec) GetEnvironment() SrcmV2ClusterSpecEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1NetworkLinkServiceSpec) GetEnvironmentOk() (*SrcmV2ClusterSpecEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1NetworkLinkServiceSpec) SetEnvironment(v SrcmV2ClusterSpecEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1NetworkLinkServiceSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetNetwork

`func (o *NetworkingV1NetworkLinkServiceSpec) GetNetwork() NetworkingV1NetworkLinkServiceSpecNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkingV1NetworkLinkServiceSpec) GetNetworkOk() (*NetworkingV1NetworkLinkServiceSpecNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkingV1NetworkLinkServiceSpec) SetNetwork(v NetworkingV1NetworkLinkServiceSpecNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkingV1NetworkLinkServiceSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


