# CmkV2ClusterSpecByok

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the referred resource | 
**Related** | **string** | API URL for accessing or modifying the referred object | [readonly] 
**ResourceName** | **string** | CRN reference to the referred resource | [readonly] 

## Methods

### NewCmkV2ClusterSpecByok

`func NewCmkV2ClusterSpecByok(id string, related string, resourceName string, ) *CmkV2ClusterSpecByok`

NewCmkV2ClusterSpecByok instantiates a new CmkV2ClusterSpecByok object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmkV2ClusterSpecByokWithDefaults

`func NewCmkV2ClusterSpecByokWithDefaults() *CmkV2ClusterSpecByok`

NewCmkV2ClusterSpecByokWithDefaults instantiates a new CmkV2ClusterSpecByok object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CmkV2ClusterSpecByok) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CmkV2ClusterSpecByok) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CmkV2ClusterSpecByok) SetId(v string)`

SetId sets Id field to given value.


### GetRelated

`func (o *CmkV2ClusterSpecByok) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *CmkV2ClusterSpecByok) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *CmkV2ClusterSpecByok) SetRelated(v string)`

SetRelated sets Related field to given value.


### GetResourceName

`func (o *CmkV2ClusterSpecByok) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *CmkV2ClusterSpecByok) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *CmkV2ClusterSpecByok) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


