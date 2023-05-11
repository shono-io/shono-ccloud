# NetworkingV1NetworkStatusCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Network kind type. | 
**Vpc** | **string** | The Confluent Cloud VPC ID. | [readonly] 
**Account** | **string** | The AWS account ID associated with the Confluent Cloud VPC. | [readonly] 
**PrivateLinkEndpointService** | Pointer to **string** | The endpoint service of the Confluent Cloud VPC. (used for PrivateLink) if available. | [optional] [readonly] 
**Project** | **string** | The GCP Project ID associated with the Confluent Cloud VPC. | [readonly] 
**VpcNetwork** | **string** | The network name of the Confluent Cloud VPC. | [readonly] 
**PrivateServiceConnectServiceAttachments** | Pointer to **map[string]string** | The mapping of zones to Private Service Connect Service Attachments if available. Keys are zones and values are [GCP Private Service Connect Service Attachment](https://cloud.google.com/vpc/docs/configure-private-service-connect-producer#api_7)  | [optional] [readonly] 
**Vnet** | **string** | The resource ID of the Confluent Cloud VNet. | [readonly] 
**Subscription** | **string** | The Azure Subscription ID associated with the Confluent Cloud VPC. | [readonly] 
**PrivateLinkServiceAliases** | Pointer to **map[string]string** | The mapping of zones to Private Link Service Aliases if available. Keys are zones and values are [Azure Private Link Service Aliases](https://docs.microsoft.com/en-us/azure/private-link/private-link-service-overview#share-your-service).  | [optional] [readonly] 
**PrivateLinkServiceResourceIds** | Pointer to **map[string]string** | The mapping of zones to Private Link Service Resource IDs if available. Keys are zones and values are [Azure Private Link Service Resource IDs](https://docs.microsoft.com/en-us/azure/private-link/private-link-service-overview#share-your-service).  | [optional] [readonly] 

## Methods

### NewNetworkingV1NetworkStatusCloud

`func NewNetworkingV1NetworkStatusCloud(kind string, vpc string, account string, project string, vpcNetwork string, vnet string, subscription string, ) *NetworkingV1NetworkStatusCloud`

NewNetworkingV1NetworkStatusCloud instantiates a new NetworkingV1NetworkStatusCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkStatusCloudWithDefaults

`func NewNetworkingV1NetworkStatusCloudWithDefaults() *NetworkingV1NetworkStatusCloud`

NewNetworkingV1NetworkStatusCloudWithDefaults instantiates a new NetworkingV1NetworkStatusCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1NetworkStatusCloud) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1NetworkStatusCloud) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1NetworkStatusCloud) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVpc

`func (o *NetworkingV1NetworkStatusCloud) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *NetworkingV1NetworkStatusCloud) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *NetworkingV1NetworkStatusCloud) SetVpc(v string)`

SetVpc sets Vpc field to given value.


### GetAccount

`func (o *NetworkingV1NetworkStatusCloud) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkingV1NetworkStatusCloud) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkingV1NetworkStatusCloud) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetPrivateLinkEndpointService

`func (o *NetworkingV1NetworkStatusCloud) GetPrivateLinkEndpointService() string`

GetPrivateLinkEndpointService returns the PrivateLinkEndpointService field if non-nil, zero value otherwise.

### GetPrivateLinkEndpointServiceOk

`func (o *NetworkingV1NetworkStatusCloud) GetPrivateLinkEndpointServiceOk() (*string, bool)`

GetPrivateLinkEndpointServiceOk returns a tuple with the PrivateLinkEndpointService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkEndpointService

`func (o *NetworkingV1NetworkStatusCloud) SetPrivateLinkEndpointService(v string)`

SetPrivateLinkEndpointService sets PrivateLinkEndpointService field to given value.

### HasPrivateLinkEndpointService

`func (o *NetworkingV1NetworkStatusCloud) HasPrivateLinkEndpointService() bool`

HasPrivateLinkEndpointService returns a boolean if a field has been set.

### GetProject

`func (o *NetworkingV1NetworkStatusCloud) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *NetworkingV1NetworkStatusCloud) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *NetworkingV1NetworkStatusCloud) SetProject(v string)`

SetProject sets Project field to given value.


### GetVpcNetwork

`func (o *NetworkingV1NetworkStatusCloud) GetVpcNetwork() string`

GetVpcNetwork returns the VpcNetwork field if non-nil, zero value otherwise.

### GetVpcNetworkOk

`func (o *NetworkingV1NetworkStatusCloud) GetVpcNetworkOk() (*string, bool)`

GetVpcNetworkOk returns a tuple with the VpcNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcNetwork

`func (o *NetworkingV1NetworkStatusCloud) SetVpcNetwork(v string)`

SetVpcNetwork sets VpcNetwork field to given value.


### GetPrivateServiceConnectServiceAttachments

`func (o *NetworkingV1NetworkStatusCloud) GetPrivateServiceConnectServiceAttachments() map[string]string`

GetPrivateServiceConnectServiceAttachments returns the PrivateServiceConnectServiceAttachments field if non-nil, zero value otherwise.

### GetPrivateServiceConnectServiceAttachmentsOk

`func (o *NetworkingV1NetworkStatusCloud) GetPrivateServiceConnectServiceAttachmentsOk() (*map[string]string, bool)`

GetPrivateServiceConnectServiceAttachmentsOk returns a tuple with the PrivateServiceConnectServiceAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServiceConnectServiceAttachments

`func (o *NetworkingV1NetworkStatusCloud) SetPrivateServiceConnectServiceAttachments(v map[string]string)`

SetPrivateServiceConnectServiceAttachments sets PrivateServiceConnectServiceAttachments field to given value.

### HasPrivateServiceConnectServiceAttachments

`func (o *NetworkingV1NetworkStatusCloud) HasPrivateServiceConnectServiceAttachments() bool`

HasPrivateServiceConnectServiceAttachments returns a boolean if a field has been set.

### GetVnet

`func (o *NetworkingV1NetworkStatusCloud) GetVnet() string`

GetVnet returns the Vnet field if non-nil, zero value otherwise.

### GetVnetOk

`func (o *NetworkingV1NetworkStatusCloud) GetVnetOk() (*string, bool)`

GetVnetOk returns a tuple with the Vnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnet

`func (o *NetworkingV1NetworkStatusCloud) SetVnet(v string)`

SetVnet sets Vnet field to given value.


### GetSubscription

`func (o *NetworkingV1NetworkStatusCloud) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *NetworkingV1NetworkStatusCloud) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *NetworkingV1NetworkStatusCloud) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.


### GetPrivateLinkServiceAliases

`func (o *NetworkingV1NetworkStatusCloud) GetPrivateLinkServiceAliases() map[string]string`

GetPrivateLinkServiceAliases returns the PrivateLinkServiceAliases field if non-nil, zero value otherwise.

### GetPrivateLinkServiceAliasesOk

`func (o *NetworkingV1NetworkStatusCloud) GetPrivateLinkServiceAliasesOk() (*map[string]string, bool)`

GetPrivateLinkServiceAliasesOk returns a tuple with the PrivateLinkServiceAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceAliases

`func (o *NetworkingV1NetworkStatusCloud) SetPrivateLinkServiceAliases(v map[string]string)`

SetPrivateLinkServiceAliases sets PrivateLinkServiceAliases field to given value.

### HasPrivateLinkServiceAliases

`func (o *NetworkingV1NetworkStatusCloud) HasPrivateLinkServiceAliases() bool`

HasPrivateLinkServiceAliases returns a boolean if a field has been set.

### GetPrivateLinkServiceResourceIds

`func (o *NetworkingV1NetworkStatusCloud) GetPrivateLinkServiceResourceIds() map[string]string`

GetPrivateLinkServiceResourceIds returns the PrivateLinkServiceResourceIds field if non-nil, zero value otherwise.

### GetPrivateLinkServiceResourceIdsOk

`func (o *NetworkingV1NetworkStatusCloud) GetPrivateLinkServiceResourceIdsOk() (*map[string]string, bool)`

GetPrivateLinkServiceResourceIdsOk returns a tuple with the PrivateLinkServiceResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceResourceIds

`func (o *NetworkingV1NetworkStatusCloud) SetPrivateLinkServiceResourceIds(v map[string]string)`

SetPrivateLinkServiceResourceIds sets PrivateLinkServiceResourceIds field to given value.

### HasPrivateLinkServiceResourceIds

`func (o *NetworkingV1NetworkStatusCloud) HasPrivateLinkServiceResourceIds() bool`

HasPrivateLinkServiceResourceIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


