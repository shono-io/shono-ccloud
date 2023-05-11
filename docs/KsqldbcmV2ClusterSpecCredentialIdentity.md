# KsqldbcmV2ClusterSpecCredentialIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the referred resource | 
**Related** | **string** | API URL for accessing or modifying the referred object | [readonly] 
**ResourceName** | **string** | CRN reference to the referred resource | [readonly] 
**ApiVersion** | Pointer to **string** | API group and version of the referred resource | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the referred resource | [optional] [readonly] 

## Methods

### NewKsqldbcmV2ClusterSpecCredentialIdentity

`func NewKsqldbcmV2ClusterSpecCredentialIdentity(id string, related string, resourceName string, ) *KsqldbcmV2ClusterSpecCredentialIdentity`

NewKsqldbcmV2ClusterSpecCredentialIdentity instantiates a new KsqldbcmV2ClusterSpecCredentialIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKsqldbcmV2ClusterSpecCredentialIdentityWithDefaults

`func NewKsqldbcmV2ClusterSpecCredentialIdentityWithDefaults() *KsqldbcmV2ClusterSpecCredentialIdentity`

NewKsqldbcmV2ClusterSpecCredentialIdentityWithDefaults instantiates a new KsqldbcmV2ClusterSpecCredentialIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) SetId(v string)`

SetId sets Id field to given value.


### GetRelated

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) SetRelated(v string)`

SetRelated sets Related field to given value.


### GetResourceName

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetApiVersion

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KsqldbcmV2ClusterSpecCredentialIdentity) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


