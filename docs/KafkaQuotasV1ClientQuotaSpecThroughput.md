# KafkaQuotasV1ClientQuotaSpecThroughput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngressByteRate** | **string** | Ingress throughput limit for principals specified in bytes per second. | 
**EgressByteRate** | **string** | Egress throughput limit for principals specified in bytes per second. | 

## Methods

### NewKafkaQuotasV1ClientQuotaSpecThroughput

`func NewKafkaQuotasV1ClientQuotaSpecThroughput(ingressByteRate string, egressByteRate string, ) *KafkaQuotasV1ClientQuotaSpecThroughput`

NewKafkaQuotasV1ClientQuotaSpecThroughput instantiates a new KafkaQuotasV1ClientQuotaSpecThroughput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaQuotasV1ClientQuotaSpecThroughputWithDefaults

`func NewKafkaQuotasV1ClientQuotaSpecThroughputWithDefaults() *KafkaQuotasV1ClientQuotaSpecThroughput`

NewKafkaQuotasV1ClientQuotaSpecThroughputWithDefaults instantiates a new KafkaQuotasV1ClientQuotaSpecThroughput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngressByteRate

`func (o *KafkaQuotasV1ClientQuotaSpecThroughput) GetIngressByteRate() string`

GetIngressByteRate returns the IngressByteRate field if non-nil, zero value otherwise.

### GetIngressByteRateOk

`func (o *KafkaQuotasV1ClientQuotaSpecThroughput) GetIngressByteRateOk() (*string, bool)`

GetIngressByteRateOk returns a tuple with the IngressByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressByteRate

`func (o *KafkaQuotasV1ClientQuotaSpecThroughput) SetIngressByteRate(v string)`

SetIngressByteRate sets IngressByteRate field to given value.


### GetEgressByteRate

`func (o *KafkaQuotasV1ClientQuotaSpecThroughput) GetEgressByteRate() string`

GetEgressByteRate returns the EgressByteRate field if non-nil, zero value otherwise.

### GetEgressByteRateOk

`func (o *KafkaQuotasV1ClientQuotaSpecThroughput) GetEgressByteRateOk() (*string, bool)`

GetEgressByteRateOk returns a tuple with the EgressByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressByteRate

`func (o *KafkaQuotasV1ClientQuotaSpecThroughput) SetEgressByteRate(v string)`

SetEgressByteRate sets EgressByteRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


