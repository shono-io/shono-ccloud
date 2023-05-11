# GetServiceQuotaV1AppliedQuota200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**ServiceQuotaV1AppliedQuotaMetadata**](ServiceQuotaV1AppliedQuotaMetadata.md) |  | [optional] 
**Scope** | **string** | The applied scope that this quota belongs to. | 
**DisplayName** | **string** | A human-readable name for the quota type name. | 
**DefaultLimit** | **int32** | The default service quota value.  | 
**AppliedLimit** | **int32** | The latest applied service quota value, taking into account any limit adjustments.  | 
**Usage** | Pointer to **int32** | Show the quota usage value if the quota usage is available for this quota.  | [optional] 
**User** | Pointer to **interface{}** |  | [optional] 
**Organization** | Pointer to **interface{}** |  | [optional] 
**Environment** | Pointer to **interface{}** |  | [optional] 
**Network** | Pointer to **interface{}** |  | [optional] 
**KafkaCluster** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewGetServiceQuotaV1AppliedQuota200Response

`func NewGetServiceQuotaV1AppliedQuota200Response(apiVersion string, kind string, id string, scope string, displayName string, defaultLimit int32, appliedLimit int32, ) *GetServiceQuotaV1AppliedQuota200Response`

NewGetServiceQuotaV1AppliedQuota200Response instantiates a new GetServiceQuotaV1AppliedQuota200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceQuotaV1AppliedQuota200ResponseWithDefaults

`func NewGetServiceQuotaV1AppliedQuota200ResponseWithDefaults() *GetServiceQuotaV1AppliedQuota200Response`

NewGetServiceQuotaV1AppliedQuota200ResponseWithDefaults instantiates a new GetServiceQuotaV1AppliedQuota200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetMetadata() ServiceQuotaV1AppliedQuotaMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetMetadataOk() (*ServiceQuotaV1AppliedQuotaMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetMetadata(v ServiceQuotaV1AppliedQuotaMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetServiceQuotaV1AppliedQuota200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetScope

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetScope(v string)`

SetScope sets Scope field to given value.


### GetDisplayName

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDefaultLimit

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetDefaultLimit() int32`

GetDefaultLimit returns the DefaultLimit field if non-nil, zero value otherwise.

### GetDefaultLimitOk

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetDefaultLimitOk() (*int32, bool)`

GetDefaultLimitOk returns a tuple with the DefaultLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLimit

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetDefaultLimit(v int32)`

SetDefaultLimit sets DefaultLimit field to given value.


### GetAppliedLimit

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetAppliedLimit() int32`

GetAppliedLimit returns the AppliedLimit field if non-nil, zero value otherwise.

### GetAppliedLimitOk

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetAppliedLimitOk() (*int32, bool)`

GetAppliedLimitOk returns a tuple with the AppliedLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedLimit

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetAppliedLimit(v int32)`

SetAppliedLimit sets AppliedLimit field to given value.


### GetUsage

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *GetServiceQuotaV1AppliedQuota200Response) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUser

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetUser() interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetUserOk() (*interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetUser(v interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *GetServiceQuotaV1AppliedQuota200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GetServiceQuotaV1AppliedQuota200Response) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetOrganization

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetOrganization() interface{}`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetOrganizationOk() (*interface{}, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetOrganization(v interface{})`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *GetServiceQuotaV1AppliedQuota200Response) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *GetServiceQuotaV1AppliedQuota200Response) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetEnvironment

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetEnvironment() interface{}`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetEnvironmentOk() (*interface{}, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetEnvironment(v interface{})`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *GetServiceQuotaV1AppliedQuota200Response) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *GetServiceQuotaV1AppliedQuota200Response) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetNetwork

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetNetwork() interface{}`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetNetworkOk() (*interface{}, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetNetwork(v interface{})`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetServiceQuotaV1AppliedQuota200Response) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *GetServiceQuotaV1AppliedQuota200Response) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetKafkaCluster

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetKafkaCluster() interface{}`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *GetServiceQuotaV1AppliedQuota200Response) GetKafkaClusterOk() (*interface{}, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetKafkaCluster(v interface{})`

SetKafkaCluster sets KafkaCluster field to given value.

### HasKafkaCluster

`func (o *GetServiceQuotaV1AppliedQuota200Response) HasKafkaCluster() bool`

HasKafkaCluster returns a boolean if a field has been set.

### SetKafkaClusterNil

`func (o *GetServiceQuotaV1AppliedQuota200Response) SetKafkaClusterNil(b bool)`

 SetKafkaClusterNil sets the value for KafkaCluster to be an explicit nil

### UnsetKafkaCluster
`func (o *GetServiceQuotaV1AppliedQuota200Response) UnsetKafkaCluster()`

UnsetKafkaCluster ensures that no value is present for KafkaCluster, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


