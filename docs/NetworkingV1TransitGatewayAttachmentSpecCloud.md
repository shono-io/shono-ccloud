# NetworkingV1TransitGatewayAttachmentSpecCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AWS Transit Gateway Attachment kind type. | 
**RamShareArn** | **string** | The full AWS Resource Name (ARN) for the AWS Resource Access Manager (RAM) Share of the Transit Gateways that you want Confluent Cloud to be attached to. | 
**TransitGatewayId** | **string** | The ID of the AWS Transit Gateway that you want Confluent CLoud to be attached to. | 
**Routes** | **[]string** | List of destination routes. | 

## Methods

### NewNetworkingV1TransitGatewayAttachmentSpecCloud

`func NewNetworkingV1TransitGatewayAttachmentSpecCloud(kind string, ramShareArn string, transitGatewayId string, routes []string, ) *NetworkingV1TransitGatewayAttachmentSpecCloud`

NewNetworkingV1TransitGatewayAttachmentSpecCloud instantiates a new NetworkingV1TransitGatewayAttachmentSpecCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1TransitGatewayAttachmentSpecCloudWithDefaults

`func NewNetworkingV1TransitGatewayAttachmentSpecCloudWithDefaults() *NetworkingV1TransitGatewayAttachmentSpecCloud`

NewNetworkingV1TransitGatewayAttachmentSpecCloudWithDefaults instantiates a new NetworkingV1TransitGatewayAttachmentSpecCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1TransitGatewayAttachmentSpecCloud) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1TransitGatewayAttachmentSpecCloud) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1TransitGatewayAttachmentSpecCloud) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRamShareArn

`func (o *NetworkingV1TransitGatewayAttachmentSpecCloud) GetRamShareArn() string`

GetRamShareArn returns the RamShareArn field if non-nil, zero value otherwise.

### GetRamShareArnOk

`func (o *NetworkingV1TransitGatewayAttachmentSpecCloud) GetRamShareArnOk() (*string, bool)`

GetRamShareArnOk returns a tuple with the RamShareArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamShareArn

`func (o *NetworkingV1TransitGatewayAttachmentSpecCloud) SetRamShareArn(v string)`

SetRamShareArn sets RamShareArn field to given value.


### GetTransitGatewayId

`func (o *NetworkingV1TransitGatewayAttachmentSpecCloud) GetTransitGatewayId() string`

GetTransitGatewayId returns the TransitGatewayId field if non-nil, zero value otherwise.

### GetTransitGatewayIdOk

`func (o *NetworkingV1TransitGatewayAttachmentSpecCloud) GetTransitGatewayIdOk() (*string, bool)`

GetTransitGatewayIdOk returns a tuple with the TransitGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitGatewayId

`func (o *NetworkingV1TransitGatewayAttachmentSpecCloud) SetTransitGatewayId(v string)`

SetTransitGatewayId sets TransitGatewayId field to given value.


### GetRoutes

`func (o *NetworkingV1TransitGatewayAttachmentSpecCloud) GetRoutes() []string`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *NetworkingV1TransitGatewayAttachmentSpecCloud) GetRoutesOk() (*[]string, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *NetworkingV1TransitGatewayAttachmentSpecCloud) SetRoutes(v []string)`

SetRoutes sets Routes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


