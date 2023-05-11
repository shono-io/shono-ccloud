# CmkV2ClusterSpecConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Dedicated cluster type.  | 
**Cku** | **int32** |  | 
**EncryptionKey** | Pointer to **string** | The id of the encryption key that is used to encrypt the data in the Kafka cluster. (e.g. for Amazon Web Services, the Amazon Resource Name of the key).  | [optional] 
**Zones** | Pointer to **[]string** | The list of zones the cluster is in.  On AWS, zones are AWS [AZ IDs](https://docs.aws.amazon.com/ram/latest/userguide/working-with-az-ids.html)  (e.g. use1-az3)  On GCP, zones are GCP [zones](https://cloud.google.com/compute/docs/regions-zones)  (e.g. us-central1-c).  | [optional] [readonly] 

## Methods

### NewCmkV2ClusterSpecConfig

`func NewCmkV2ClusterSpecConfig(kind string, cku int32, ) *CmkV2ClusterSpecConfig`

NewCmkV2ClusterSpecConfig instantiates a new CmkV2ClusterSpecConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmkV2ClusterSpecConfigWithDefaults

`func NewCmkV2ClusterSpecConfigWithDefaults() *CmkV2ClusterSpecConfig`

NewCmkV2ClusterSpecConfigWithDefaults instantiates a new CmkV2ClusterSpecConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CmkV2ClusterSpecConfig) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CmkV2ClusterSpecConfig) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CmkV2ClusterSpecConfig) SetKind(v string)`

SetKind sets Kind field to given value.


### GetCku

`func (o *CmkV2ClusterSpecConfig) GetCku() int32`

GetCku returns the Cku field if non-nil, zero value otherwise.

### GetCkuOk

`func (o *CmkV2ClusterSpecConfig) GetCkuOk() (*int32, bool)`

GetCkuOk returns a tuple with the Cku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCku

`func (o *CmkV2ClusterSpecConfig) SetCku(v int32)`

SetCku sets Cku field to given value.


### GetEncryptionKey

`func (o *CmkV2ClusterSpecConfig) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *CmkV2ClusterSpecConfig) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *CmkV2ClusterSpecConfig) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *CmkV2ClusterSpecConfig) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetZones

`func (o *CmkV2ClusterSpecConfig) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *CmkV2ClusterSpecConfig) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *CmkV2ClusterSpecConfig) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *CmkV2ClusterSpecConfig) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


