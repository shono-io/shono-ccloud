# ServiceQuotaV1AppliedQuotaKafkaCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the referred resource | 
**Environment** | Pointer to **string** | Environment of the referred resource, if env-scoped | [optional] 
**Related** | **string** | API URL for accessing or modifying the referred object | [readonly] 
**ResourceName** | **string** | CRN reference to the referred resource | [readonly] 

## Methods

### NewServiceQuotaV1AppliedQuotaKafkaCluster

`func NewServiceQuotaV1AppliedQuotaKafkaCluster(id string, related string, resourceName string, ) *ServiceQuotaV1AppliedQuotaKafkaCluster`

NewServiceQuotaV1AppliedQuotaKafkaCluster instantiates a new ServiceQuotaV1AppliedQuotaKafkaCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceQuotaV1AppliedQuotaKafkaClusterWithDefaults

`func NewServiceQuotaV1AppliedQuotaKafkaClusterWithDefaults() *ServiceQuotaV1AppliedQuotaKafkaCluster`

NewServiceQuotaV1AppliedQuotaKafkaClusterWithDefaults instantiates a new ServiceQuotaV1AppliedQuotaKafkaCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceQuotaV1AppliedQuotaKafkaCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceQuotaV1AppliedQuotaKafkaCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceQuotaV1AppliedQuotaKafkaCluster) SetId(v string)`

SetId sets Id field to given value.


### GetEnvironment

`func (o *ServiceQuotaV1AppliedQuotaKafkaCluster) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ServiceQuotaV1AppliedQuotaKafkaCluster) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ServiceQuotaV1AppliedQuotaKafkaCluster) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ServiceQuotaV1AppliedQuotaKafkaCluster) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRelated

`func (o *ServiceQuotaV1AppliedQuotaKafkaCluster) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *ServiceQuotaV1AppliedQuotaKafkaCluster) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *ServiceQuotaV1AppliedQuotaKafkaCluster) SetRelated(v string)`

SetRelated sets Related field to given value.


### GetResourceName

`func (o *ServiceQuotaV1AppliedQuotaKafkaCluster) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ServiceQuotaV1AppliedQuotaKafkaCluster) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ServiceQuotaV1AppliedQuotaKafkaCluster) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


