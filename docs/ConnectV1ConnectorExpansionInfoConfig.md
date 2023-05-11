# ConnectV1ConnectorExpansionInfoConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudEnvironment** | **string** | The cloud environment type. | 
**CloudProvider** | **string** | The cloud service provider, e.g. aws, azure, etc. | 
**ConnectorClass** | **string** | The connector class name. E.g. BigQuerySink, GcsSink, etc. | 
**Name** | **string** | Name or alias of the class (plugin) for this connector. | 
**KafkaEndpoint** | **string** | The kafka cluster endpoint. | 
**KafkaRegion** | **string** | The kafka cluster region. | 
**KafkaApiKey** | **string** | The kafka cluster api key. | 
**KafkaApiSecret** | **string** | The kafka cluster api secret key. | 

## Methods

### NewConnectV1ConnectorExpansionInfoConfig

`func NewConnectV1ConnectorExpansionInfoConfig(cloudEnvironment string, cloudProvider string, connectorClass string, name string, kafkaEndpoint string, kafkaRegion string, kafkaApiKey string, kafkaApiSecret string, ) *ConnectV1ConnectorExpansionInfoConfig`

NewConnectV1ConnectorExpansionInfoConfig instantiates a new ConnectV1ConnectorExpansionInfoConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1ConnectorExpansionInfoConfigWithDefaults

`func NewConnectV1ConnectorExpansionInfoConfigWithDefaults() *ConnectV1ConnectorExpansionInfoConfig`

NewConnectV1ConnectorExpansionInfoConfigWithDefaults instantiates a new ConnectV1ConnectorExpansionInfoConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudEnvironment

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetCloudEnvironment() string`

GetCloudEnvironment returns the CloudEnvironment field if non-nil, zero value otherwise.

### GetCloudEnvironmentOk

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetCloudEnvironmentOk() (*string, bool)`

GetCloudEnvironmentOk returns a tuple with the CloudEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudEnvironment

`func (o *ConnectV1ConnectorExpansionInfoConfig) SetCloudEnvironment(v string)`

SetCloudEnvironment sets CloudEnvironment field to given value.


### GetCloudProvider

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ConnectV1ConnectorExpansionInfoConfig) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetConnectorClass

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetConnectorClass() string`

GetConnectorClass returns the ConnectorClass field if non-nil, zero value otherwise.

### GetConnectorClassOk

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetConnectorClassOk() (*string, bool)`

GetConnectorClassOk returns a tuple with the ConnectorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorClass

`func (o *ConnectV1ConnectorExpansionInfoConfig) SetConnectorClass(v string)`

SetConnectorClass sets ConnectorClass field to given value.


### GetName

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectV1ConnectorExpansionInfoConfig) SetName(v string)`

SetName sets Name field to given value.


### GetKafkaEndpoint

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetKafkaEndpoint() string`

GetKafkaEndpoint returns the KafkaEndpoint field if non-nil, zero value otherwise.

### GetKafkaEndpointOk

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetKafkaEndpointOk() (*string, bool)`

GetKafkaEndpointOk returns a tuple with the KafkaEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaEndpoint

`func (o *ConnectV1ConnectorExpansionInfoConfig) SetKafkaEndpoint(v string)`

SetKafkaEndpoint sets KafkaEndpoint field to given value.


### GetKafkaRegion

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetKafkaRegion() string`

GetKafkaRegion returns the KafkaRegion field if non-nil, zero value otherwise.

### GetKafkaRegionOk

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetKafkaRegionOk() (*string, bool)`

GetKafkaRegionOk returns a tuple with the KafkaRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaRegion

`func (o *ConnectV1ConnectorExpansionInfoConfig) SetKafkaRegion(v string)`

SetKafkaRegion sets KafkaRegion field to given value.


### GetKafkaApiKey

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetKafkaApiKey() string`

GetKafkaApiKey returns the KafkaApiKey field if non-nil, zero value otherwise.

### GetKafkaApiKeyOk

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetKafkaApiKeyOk() (*string, bool)`

GetKafkaApiKeyOk returns a tuple with the KafkaApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaApiKey

`func (o *ConnectV1ConnectorExpansionInfoConfig) SetKafkaApiKey(v string)`

SetKafkaApiKey sets KafkaApiKey field to given value.


### GetKafkaApiSecret

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetKafkaApiSecret() string`

GetKafkaApiSecret returns the KafkaApiSecret field if non-nil, zero value otherwise.

### GetKafkaApiSecretOk

`func (o *ConnectV1ConnectorExpansionInfoConfig) GetKafkaApiSecretOk() (*string, bool)`

GetKafkaApiSecretOk returns a tuple with the KafkaApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaApiSecret

`func (o *ConnectV1ConnectorExpansionInfoConfig) SetKafkaApiSecret(v string)`

SetKafkaApiSecret sets KafkaApiSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


