# NetworkingV1NetworkSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the network | [optional] 
**ZonesInfo** | Pointer to [**[]NetworkingV1ZoneInfo**](NetworkingV1ZoneInfo.md) | Cloud provider zones metadata. | [optional] 
**Environment** | Pointer to [**NetworkingV1NetworkSpecEnvironment**](NetworkingV1NetworkSpecEnvironment.md) |  | [optional] 

## Methods

### NewNetworkingV1NetworkSpecUpdate

`func NewNetworkingV1NetworkSpecUpdate() *NetworkingV1NetworkSpecUpdate`

NewNetworkingV1NetworkSpecUpdate instantiates a new NetworkingV1NetworkSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkSpecUpdateWithDefaults

`func NewNetworkingV1NetworkSpecUpdateWithDefaults() *NetworkingV1NetworkSpecUpdate`

NewNetworkingV1NetworkSpecUpdateWithDefaults instantiates a new NetworkingV1NetworkSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1NetworkSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1NetworkSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1NetworkSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1NetworkSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetZonesInfo

`func (o *NetworkingV1NetworkSpecUpdate) GetZonesInfo() []NetworkingV1ZoneInfo`

GetZonesInfo returns the ZonesInfo field if non-nil, zero value otherwise.

### GetZonesInfoOk

`func (o *NetworkingV1NetworkSpecUpdate) GetZonesInfoOk() (*[]NetworkingV1ZoneInfo, bool)`

GetZonesInfoOk returns a tuple with the ZonesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonesInfo

`func (o *NetworkingV1NetworkSpecUpdate) SetZonesInfo(v []NetworkingV1ZoneInfo)`

SetZonesInfo sets ZonesInfo field to given value.

### HasZonesInfo

`func (o *NetworkingV1NetworkSpecUpdate) HasZonesInfo() bool`

HasZonesInfo returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1NetworkSpecUpdate) GetEnvironment() NetworkingV1NetworkSpecEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1NetworkSpecUpdate) GetEnvironmentOk() (*NetworkingV1NetworkSpecEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1NetworkSpecUpdate) SetEnvironment(v NetworkingV1NetworkSpecEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1NetworkSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


