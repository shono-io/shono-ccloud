# ServiceQuotaV1AppliedQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ServiceQuotaV1AppliedQuotaMetadata**](ServiceQuotaV1AppliedQuotaMetadata.md) |  | [optional] 
**Scope** | Pointer to **string** | The applied scope that this quota belongs to. | [optional] 
**DisplayName** | Pointer to **string** | A human-readable name for the quota type name. | [optional] 
**DefaultLimit** | Pointer to **int32** | The default service quota value.  | [optional] 
**AppliedLimit** | Pointer to **int32** | The latest applied service quota value, taking into account any limit adjustments.  | [optional] 
**Usage** | Pointer to **int32** | Show the quota usage value if the quota usage is available for this quota.  | [optional] 
**User** | Pointer to [**ServiceQuotaV1AppliedQuotaUser**](ServiceQuotaV1AppliedQuotaUser.md) |  | [optional] 
**Organization** | Pointer to [**NullableServiceQuotaV1AppliedQuotaOrganization**](ServiceQuotaV1AppliedQuotaOrganization.md) |  | [optional] 
**Environment** | Pointer to [**NullableServiceQuotaV1AppliedQuotaEnvironment**](ServiceQuotaV1AppliedQuotaEnvironment.md) |  | [optional] 
**Network** | Pointer to [**NullableServiceQuotaV1AppliedQuotaNetwork**](ServiceQuotaV1AppliedQuotaNetwork.md) |  | [optional] 
**KafkaCluster** | Pointer to [**NullableServiceQuotaV1AppliedQuotaKafkaCluster**](ServiceQuotaV1AppliedQuotaKafkaCluster.md) |  | [optional] 

## Methods

### NewServiceQuotaV1AppliedQuota

`func NewServiceQuotaV1AppliedQuota() *ServiceQuotaV1AppliedQuota`

NewServiceQuotaV1AppliedQuota instantiates a new ServiceQuotaV1AppliedQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceQuotaV1AppliedQuotaWithDefaults

`func NewServiceQuotaV1AppliedQuotaWithDefaults() *ServiceQuotaV1AppliedQuota`

NewServiceQuotaV1AppliedQuotaWithDefaults instantiates a new ServiceQuotaV1AppliedQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ServiceQuotaV1AppliedQuota) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ServiceQuotaV1AppliedQuota) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ServiceQuotaV1AppliedQuota) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ServiceQuotaV1AppliedQuota) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ServiceQuotaV1AppliedQuota) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ServiceQuotaV1AppliedQuota) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ServiceQuotaV1AppliedQuota) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ServiceQuotaV1AppliedQuota) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ServiceQuotaV1AppliedQuota) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceQuotaV1AppliedQuota) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceQuotaV1AppliedQuota) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceQuotaV1AppliedQuota) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceQuotaV1AppliedQuota) GetMetadata() ServiceQuotaV1AppliedQuotaMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceQuotaV1AppliedQuota) GetMetadataOk() (*ServiceQuotaV1AppliedQuotaMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceQuotaV1AppliedQuota) SetMetadata(v ServiceQuotaV1AppliedQuotaMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceQuotaV1AppliedQuota) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetScope

`func (o *ServiceQuotaV1AppliedQuota) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ServiceQuotaV1AppliedQuota) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ServiceQuotaV1AppliedQuota) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ServiceQuotaV1AppliedQuota) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDisplayName

`func (o *ServiceQuotaV1AppliedQuota) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ServiceQuotaV1AppliedQuota) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ServiceQuotaV1AppliedQuota) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ServiceQuotaV1AppliedQuota) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDefaultLimit

`func (o *ServiceQuotaV1AppliedQuota) GetDefaultLimit() int32`

GetDefaultLimit returns the DefaultLimit field if non-nil, zero value otherwise.

### GetDefaultLimitOk

`func (o *ServiceQuotaV1AppliedQuota) GetDefaultLimitOk() (*int32, bool)`

GetDefaultLimitOk returns a tuple with the DefaultLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLimit

`func (o *ServiceQuotaV1AppliedQuota) SetDefaultLimit(v int32)`

SetDefaultLimit sets DefaultLimit field to given value.

### HasDefaultLimit

`func (o *ServiceQuotaV1AppliedQuota) HasDefaultLimit() bool`

HasDefaultLimit returns a boolean if a field has been set.

### GetAppliedLimit

`func (o *ServiceQuotaV1AppliedQuota) GetAppliedLimit() int32`

GetAppliedLimit returns the AppliedLimit field if non-nil, zero value otherwise.

### GetAppliedLimitOk

`func (o *ServiceQuotaV1AppliedQuota) GetAppliedLimitOk() (*int32, bool)`

GetAppliedLimitOk returns a tuple with the AppliedLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedLimit

`func (o *ServiceQuotaV1AppliedQuota) SetAppliedLimit(v int32)`

SetAppliedLimit sets AppliedLimit field to given value.

### HasAppliedLimit

`func (o *ServiceQuotaV1AppliedQuota) HasAppliedLimit() bool`

HasAppliedLimit returns a boolean if a field has been set.

### GetUsage

`func (o *ServiceQuotaV1AppliedQuota) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ServiceQuotaV1AppliedQuota) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ServiceQuotaV1AppliedQuota) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ServiceQuotaV1AppliedQuota) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUser

`func (o *ServiceQuotaV1AppliedQuota) GetUser() ServiceQuotaV1AppliedQuotaUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ServiceQuotaV1AppliedQuota) GetUserOk() (*ServiceQuotaV1AppliedQuotaUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ServiceQuotaV1AppliedQuota) SetUser(v ServiceQuotaV1AppliedQuotaUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ServiceQuotaV1AppliedQuota) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetOrganization

`func (o *ServiceQuotaV1AppliedQuota) GetOrganization() ServiceQuotaV1AppliedQuotaOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ServiceQuotaV1AppliedQuota) GetOrganizationOk() (*ServiceQuotaV1AppliedQuotaOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ServiceQuotaV1AppliedQuota) SetOrganization(v ServiceQuotaV1AppliedQuotaOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ServiceQuotaV1AppliedQuota) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ServiceQuotaV1AppliedQuota) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ServiceQuotaV1AppliedQuota) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetEnvironment

`func (o *ServiceQuotaV1AppliedQuota) GetEnvironment() ServiceQuotaV1AppliedQuotaEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ServiceQuotaV1AppliedQuota) GetEnvironmentOk() (*ServiceQuotaV1AppliedQuotaEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ServiceQuotaV1AppliedQuota) SetEnvironment(v ServiceQuotaV1AppliedQuotaEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ServiceQuotaV1AppliedQuota) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ServiceQuotaV1AppliedQuota) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ServiceQuotaV1AppliedQuota) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetNetwork

`func (o *ServiceQuotaV1AppliedQuota) GetNetwork() ServiceQuotaV1AppliedQuotaNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ServiceQuotaV1AppliedQuota) GetNetworkOk() (*ServiceQuotaV1AppliedQuotaNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ServiceQuotaV1AppliedQuota) SetNetwork(v ServiceQuotaV1AppliedQuotaNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ServiceQuotaV1AppliedQuota) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *ServiceQuotaV1AppliedQuota) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *ServiceQuotaV1AppliedQuota) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetKafkaCluster

`func (o *ServiceQuotaV1AppliedQuota) GetKafkaCluster() ServiceQuotaV1AppliedQuotaKafkaCluster`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *ServiceQuotaV1AppliedQuota) GetKafkaClusterOk() (*ServiceQuotaV1AppliedQuotaKafkaCluster, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *ServiceQuotaV1AppliedQuota) SetKafkaCluster(v ServiceQuotaV1AppliedQuotaKafkaCluster)`

SetKafkaCluster sets KafkaCluster field to given value.

### HasKafkaCluster

`func (o *ServiceQuotaV1AppliedQuota) HasKafkaCluster() bool`

HasKafkaCluster returns a boolean if a field has been set.

### SetKafkaClusterNil

`func (o *ServiceQuotaV1AppliedQuota) SetKafkaClusterNil(b bool)`

 SetKafkaClusterNil sets the value for KafkaCluster to be an explicit nil

### UnsetKafkaCluster
`func (o *ServiceQuotaV1AppliedQuota) UnsetKafkaCluster()`

UnsetKafkaCluster ensures that no value is present for KafkaCluster, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


