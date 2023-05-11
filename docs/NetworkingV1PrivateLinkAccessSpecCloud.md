# NetworkingV1PrivateLinkAccessSpecCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLink kind type. | 
**Account** | **string** | The AWS account ID | 
**Subscription** | **string** | The Azure subscription ID for the account containing the VNets you want to connect from using Azure Private Link. You can find your Azure subscription ID in the subscription section of your [Microsoft Azure Portal](https://portal.azure.com/#blade/Microsoft_Azure_Billing/SubscriptionsBlade). Must be a valid **32 character UUID string**.  | 
**Project** | **string** | The GCP project ID for the account containing the VPCs that you want to connect from using Private Service Connect. You can find your Google Cloud Project ID under **Project ID** section of your [Google Cloud Console dashboard](https://console.cloud.google.com/home/dashboard).  | 

## Methods

### NewNetworkingV1PrivateLinkAccessSpecCloud

`func NewNetworkingV1PrivateLinkAccessSpecCloud(kind string, account string, subscription string, project string, ) *NetworkingV1PrivateLinkAccessSpecCloud`

NewNetworkingV1PrivateLinkAccessSpecCloud instantiates a new NetworkingV1PrivateLinkAccessSpecCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PrivateLinkAccessSpecCloudWithDefaults

`func NewNetworkingV1PrivateLinkAccessSpecCloudWithDefaults() *NetworkingV1PrivateLinkAccessSpecCloud`

NewNetworkingV1PrivateLinkAccessSpecCloudWithDefaults instantiates a new NetworkingV1PrivateLinkAccessSpecCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1PrivateLinkAccessSpecCloud) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1PrivateLinkAccessSpecCloud) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1PrivateLinkAccessSpecCloud) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAccount

`func (o *NetworkingV1PrivateLinkAccessSpecCloud) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkingV1PrivateLinkAccessSpecCloud) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkingV1PrivateLinkAccessSpecCloud) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetSubscription

`func (o *NetworkingV1PrivateLinkAccessSpecCloud) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *NetworkingV1PrivateLinkAccessSpecCloud) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *NetworkingV1PrivateLinkAccessSpecCloud) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.


### GetProject

`func (o *NetworkingV1PrivateLinkAccessSpecCloud) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *NetworkingV1PrivateLinkAccessSpecCloud) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *NetworkingV1PrivateLinkAccessSpecCloud) SetProject(v string)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


