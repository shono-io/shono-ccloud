# NetworkingV1NetworkLinkServiceSpecNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the referred resource | 
**Environment** | Pointer to **string** | Environment of the referred resource, if env-scoped | [optional] 
**Related** | **string** | API URL for accessing or modifying the referred object | [readonly] 
**ResourceName** | **string** | CRN reference to the referred resource | [readonly] 

## Methods

### NewNetworkingV1NetworkLinkServiceSpecNetwork

`func NewNetworkingV1NetworkLinkServiceSpecNetwork(id string, related string, resourceName string, ) *NetworkingV1NetworkLinkServiceSpecNetwork`

NewNetworkingV1NetworkLinkServiceSpecNetwork instantiates a new NetworkingV1NetworkLinkServiceSpecNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkLinkServiceSpecNetworkWithDefaults

`func NewNetworkingV1NetworkLinkServiceSpecNetworkWithDefaults() *NetworkingV1NetworkLinkServiceSpecNetwork`

NewNetworkingV1NetworkLinkServiceSpecNetworkWithDefaults instantiates a new NetworkingV1NetworkLinkServiceSpecNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkingV1NetworkLinkServiceSpecNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkingV1NetworkLinkServiceSpecNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkingV1NetworkLinkServiceSpecNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetEnvironment

`func (o *NetworkingV1NetworkLinkServiceSpecNetwork) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1NetworkLinkServiceSpecNetwork) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1NetworkLinkServiceSpecNetwork) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1NetworkLinkServiceSpecNetwork) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRelated

`func (o *NetworkingV1NetworkLinkServiceSpecNetwork) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *NetworkingV1NetworkLinkServiceSpecNetwork) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *NetworkingV1NetworkLinkServiceSpecNetwork) SetRelated(v string)`

SetRelated sets Related field to given value.


### GetResourceName

`func (o *NetworkingV1NetworkLinkServiceSpecNetwork) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *NetworkingV1NetworkLinkServiceSpecNetwork) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *NetworkingV1NetworkLinkServiceSpecNetwork) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


