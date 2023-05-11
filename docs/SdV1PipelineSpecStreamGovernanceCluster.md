# SdV1PipelineSpecStreamGovernanceCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the referred resource | 
**Environment** | Pointer to **string** | Environment of the referred resource, if env-scoped | [optional] 
**Related** | **string** | API URL for accessing or modifying the referred object | [readonly] 
**ResourceName** | **string** | CRN reference to the referred resource | [readonly] 
**ApiVersion** | Pointer to **string** | API group and version of the referred resource | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the referred resource | [optional] [readonly] 

## Methods

### NewSdV1PipelineSpecStreamGovernanceCluster

`func NewSdV1PipelineSpecStreamGovernanceCluster(id string, related string, resourceName string, ) *SdV1PipelineSpecStreamGovernanceCluster`

NewSdV1PipelineSpecStreamGovernanceCluster instantiates a new SdV1PipelineSpecStreamGovernanceCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdV1PipelineSpecStreamGovernanceClusterWithDefaults

`func NewSdV1PipelineSpecStreamGovernanceClusterWithDefaults() *SdV1PipelineSpecStreamGovernanceCluster`

NewSdV1PipelineSpecStreamGovernanceClusterWithDefaults instantiates a new SdV1PipelineSpecStreamGovernanceCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SdV1PipelineSpecStreamGovernanceCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SdV1PipelineSpecStreamGovernanceCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SdV1PipelineSpecStreamGovernanceCluster) SetId(v string)`

SetId sets Id field to given value.


### GetEnvironment

`func (o *SdV1PipelineSpecStreamGovernanceCluster) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SdV1PipelineSpecStreamGovernanceCluster) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SdV1PipelineSpecStreamGovernanceCluster) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SdV1PipelineSpecStreamGovernanceCluster) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRelated

`func (o *SdV1PipelineSpecStreamGovernanceCluster) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *SdV1PipelineSpecStreamGovernanceCluster) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *SdV1PipelineSpecStreamGovernanceCluster) SetRelated(v string)`

SetRelated sets Related field to given value.


### GetResourceName

`func (o *SdV1PipelineSpecStreamGovernanceCluster) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *SdV1PipelineSpecStreamGovernanceCluster) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *SdV1PipelineSpecStreamGovernanceCluster) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetApiVersion

`func (o *SdV1PipelineSpecStreamGovernanceCluster) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SdV1PipelineSpecStreamGovernanceCluster) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SdV1PipelineSpecStreamGovernanceCluster) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SdV1PipelineSpecStreamGovernanceCluster) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SdV1PipelineSpecStreamGovernanceCluster) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SdV1PipelineSpecStreamGovernanceCluster) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SdV1PipelineSpecStreamGovernanceCluster) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SdV1PipelineSpecStreamGovernanceCluster) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


