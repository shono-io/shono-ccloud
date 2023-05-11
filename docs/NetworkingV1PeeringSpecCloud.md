# NetworkingV1PeeringSpecCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Peering kind type. | 
**Account** | **string** | The AWS account ID | 
**Vpc** | **string** | The VPC ID you are peering with Confluent Cloud network. | 
**Routes** | **[]string** | The [CIDR blocks](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) of the VPC you are peering with Confluent Cloud network. This is used by Confluent Cloud network to route traffic back to your network. The CIDR block must be a private range and cannot overlap with the Confluent Cloud CIDR block.  | 
**CustomerRegion** | **string** | The region of the VNet you are peering with Confluent Cloud network. | 
**Project** | **string** | The Google Cloud project ID associated with the VPC that you are peering with Confluent Cloud network.  | 
**VpcNetwork** | **string** | The name of the VPC that you are peering with Confluent Cloud network. | 
**ImportCustomRoutes** | Pointer to **bool** | Enable customer route import. For more information, see [Importing custom routes](https://cloud.google.com/vpc/docs/vpc-peering#importing-exporting-routes).  | [optional] [default to false]
**Tenant** | **string** | The Azure Tenant ID in which your Azure Subscription exists. Represents an organization in Azure Active Directory. You can find your Azure Tenant ID in the Azure Portal under [Azure Active Directory](https://portal.azure.com/#blade/Microsoft_AAD_IAM/ActiveDirectoryMenuBlade/Overview). Must be a valid **32 character UUID string**.  | 
**Vnet** | **string** | The resource ID of the VNet that you are peering with Confluent Cloud. You can find the name of your Azure VNet in the [Azure Portal on the Overview tab of your Azure Virtual Network](https://portal.azure.com/#blade/HubsExtension/BrowseResource/resourceType/Microsoft.Network%2FvirtualNetworks). | 

## Methods

### NewNetworkingV1PeeringSpecCloud

`func NewNetworkingV1PeeringSpecCloud(kind string, account string, vpc string, routes []string, customerRegion string, project string, vpcNetwork string, tenant string, vnet string, ) *NetworkingV1PeeringSpecCloud`

NewNetworkingV1PeeringSpecCloud instantiates a new NetworkingV1PeeringSpecCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PeeringSpecCloudWithDefaults

`func NewNetworkingV1PeeringSpecCloudWithDefaults() *NetworkingV1PeeringSpecCloud`

NewNetworkingV1PeeringSpecCloudWithDefaults instantiates a new NetworkingV1PeeringSpecCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1PeeringSpecCloud) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1PeeringSpecCloud) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1PeeringSpecCloud) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAccount

`func (o *NetworkingV1PeeringSpecCloud) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkingV1PeeringSpecCloud) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkingV1PeeringSpecCloud) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetVpc

`func (o *NetworkingV1PeeringSpecCloud) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *NetworkingV1PeeringSpecCloud) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *NetworkingV1PeeringSpecCloud) SetVpc(v string)`

SetVpc sets Vpc field to given value.


### GetRoutes

`func (o *NetworkingV1PeeringSpecCloud) GetRoutes() []string`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *NetworkingV1PeeringSpecCloud) GetRoutesOk() (*[]string, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *NetworkingV1PeeringSpecCloud) SetRoutes(v []string)`

SetRoutes sets Routes field to given value.


### GetCustomerRegion

`func (o *NetworkingV1PeeringSpecCloud) GetCustomerRegion() string`

GetCustomerRegion returns the CustomerRegion field if non-nil, zero value otherwise.

### GetCustomerRegionOk

`func (o *NetworkingV1PeeringSpecCloud) GetCustomerRegionOk() (*string, bool)`

GetCustomerRegionOk returns a tuple with the CustomerRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRegion

`func (o *NetworkingV1PeeringSpecCloud) SetCustomerRegion(v string)`

SetCustomerRegion sets CustomerRegion field to given value.


### GetProject

`func (o *NetworkingV1PeeringSpecCloud) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *NetworkingV1PeeringSpecCloud) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *NetworkingV1PeeringSpecCloud) SetProject(v string)`

SetProject sets Project field to given value.


### GetVpcNetwork

`func (o *NetworkingV1PeeringSpecCloud) GetVpcNetwork() string`

GetVpcNetwork returns the VpcNetwork field if non-nil, zero value otherwise.

### GetVpcNetworkOk

`func (o *NetworkingV1PeeringSpecCloud) GetVpcNetworkOk() (*string, bool)`

GetVpcNetworkOk returns a tuple with the VpcNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcNetwork

`func (o *NetworkingV1PeeringSpecCloud) SetVpcNetwork(v string)`

SetVpcNetwork sets VpcNetwork field to given value.


### GetImportCustomRoutes

`func (o *NetworkingV1PeeringSpecCloud) GetImportCustomRoutes() bool`

GetImportCustomRoutes returns the ImportCustomRoutes field if non-nil, zero value otherwise.

### GetImportCustomRoutesOk

`func (o *NetworkingV1PeeringSpecCloud) GetImportCustomRoutesOk() (*bool, bool)`

GetImportCustomRoutesOk returns a tuple with the ImportCustomRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportCustomRoutes

`func (o *NetworkingV1PeeringSpecCloud) SetImportCustomRoutes(v bool)`

SetImportCustomRoutes sets ImportCustomRoutes field to given value.

### HasImportCustomRoutes

`func (o *NetworkingV1PeeringSpecCloud) HasImportCustomRoutes() bool`

HasImportCustomRoutes returns a boolean if a field has been set.

### GetTenant

`func (o *NetworkingV1PeeringSpecCloud) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *NetworkingV1PeeringSpecCloud) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *NetworkingV1PeeringSpecCloud) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetVnet

`func (o *NetworkingV1PeeringSpecCloud) GetVnet() string`

GetVnet returns the Vnet field if non-nil, zero value otherwise.

### GetVnetOk

`func (o *NetworkingV1PeeringSpecCloud) GetVnetOk() (*string, bool)`

GetVnetOk returns a tuple with the Vnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnet

`func (o *NetworkingV1PeeringSpecCloud) SetVnet(v string)`

SetVnet sets Vnet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


