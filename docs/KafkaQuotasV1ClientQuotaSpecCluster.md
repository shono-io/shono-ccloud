# KafkaQuotasV1ClientQuotaSpecCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the referred resource | 
**Environment** | Pointer to **string** | Environment of the referred resource, if env-scoped | [optional] 
**Related** | **string** | API URL for accessing or modifying the referred object | [readonly] 
**ResourceName** | **string** | CRN reference to the referred resource | [readonly] 

## Methods

### NewKafkaQuotasV1ClientQuotaSpecCluster

`func NewKafkaQuotasV1ClientQuotaSpecCluster(id string, related string, resourceName string, ) *KafkaQuotasV1ClientQuotaSpecCluster`

NewKafkaQuotasV1ClientQuotaSpecCluster instantiates a new KafkaQuotasV1ClientQuotaSpecCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaQuotasV1ClientQuotaSpecClusterWithDefaults

`func NewKafkaQuotasV1ClientQuotaSpecClusterWithDefaults() *KafkaQuotasV1ClientQuotaSpecCluster`

NewKafkaQuotasV1ClientQuotaSpecClusterWithDefaults instantiates a new KafkaQuotasV1ClientQuotaSpecCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KafkaQuotasV1ClientQuotaSpecCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KafkaQuotasV1ClientQuotaSpecCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KafkaQuotasV1ClientQuotaSpecCluster) SetId(v string)`

SetId sets Id field to given value.


### GetEnvironment

`func (o *KafkaQuotasV1ClientQuotaSpecCluster) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KafkaQuotasV1ClientQuotaSpecCluster) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KafkaQuotasV1ClientQuotaSpecCluster) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *KafkaQuotasV1ClientQuotaSpecCluster) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRelated

`func (o *KafkaQuotasV1ClientQuotaSpecCluster) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *KafkaQuotasV1ClientQuotaSpecCluster) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *KafkaQuotasV1ClientQuotaSpecCluster) SetRelated(v string)`

SetRelated sets Related field to given value.


### GetResourceName

`func (o *KafkaQuotasV1ClientQuotaSpecCluster) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *KafkaQuotasV1ClientQuotaSpecCluster) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *KafkaQuotasV1ClientQuotaSpecCluster) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


