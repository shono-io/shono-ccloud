# GetConnectv1ConnectorConfig200Response

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

### NewGetConnectv1ConnectorConfig200Response

`func NewGetConnectv1ConnectorConfig200Response(cloudEnvironment string, cloudProvider string, connectorClass string, name string, kafkaEndpoint string, kafkaRegion string, kafkaApiKey string, kafkaApiSecret string, ) *GetConnectv1ConnectorConfig200Response`

NewGetConnectv1ConnectorConfig200Response instantiates a new GetConnectv1ConnectorConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectv1ConnectorConfig200ResponseWithDefaults

`func NewGetConnectv1ConnectorConfig200ResponseWithDefaults() *GetConnectv1ConnectorConfig200Response`

NewGetConnectv1ConnectorConfig200ResponseWithDefaults instantiates a new GetConnectv1ConnectorConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudEnvironment

`func (o *GetConnectv1ConnectorConfig200Response) GetCloudEnvironment() string`

GetCloudEnvironment returns the CloudEnvironment field if non-nil, zero value otherwise.

### GetCloudEnvironmentOk

`func (o *GetConnectv1ConnectorConfig200Response) GetCloudEnvironmentOk() (*string, bool)`

GetCloudEnvironmentOk returns a tuple with the CloudEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudEnvironment

`func (o *GetConnectv1ConnectorConfig200Response) SetCloudEnvironment(v string)`

SetCloudEnvironment sets CloudEnvironment field to given value.


### GetCloudProvider

`func (o *GetConnectv1ConnectorConfig200Response) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *GetConnectv1ConnectorConfig200Response) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *GetConnectv1ConnectorConfig200Response) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetConnectorClass

`func (o *GetConnectv1ConnectorConfig200Response) GetConnectorClass() string`

GetConnectorClass returns the ConnectorClass field if non-nil, zero value otherwise.

### GetConnectorClassOk

`func (o *GetConnectv1ConnectorConfig200Response) GetConnectorClassOk() (*string, bool)`

GetConnectorClassOk returns a tuple with the ConnectorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorClass

`func (o *GetConnectv1ConnectorConfig200Response) SetConnectorClass(v string)`

SetConnectorClass sets ConnectorClass field to given value.


### GetName

`func (o *GetConnectv1ConnectorConfig200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetConnectv1ConnectorConfig200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetConnectv1ConnectorConfig200Response) SetName(v string)`

SetName sets Name field to given value.


### GetKafkaEndpoint

`func (o *GetConnectv1ConnectorConfig200Response) GetKafkaEndpoint() string`

GetKafkaEndpoint returns the KafkaEndpoint field if non-nil, zero value otherwise.

### GetKafkaEndpointOk

`func (o *GetConnectv1ConnectorConfig200Response) GetKafkaEndpointOk() (*string, bool)`

GetKafkaEndpointOk returns a tuple with the KafkaEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaEndpoint

`func (o *GetConnectv1ConnectorConfig200Response) SetKafkaEndpoint(v string)`

SetKafkaEndpoint sets KafkaEndpoint field to given value.


### GetKafkaRegion

`func (o *GetConnectv1ConnectorConfig200Response) GetKafkaRegion() string`

GetKafkaRegion returns the KafkaRegion field if non-nil, zero value otherwise.

### GetKafkaRegionOk

`func (o *GetConnectv1ConnectorConfig200Response) GetKafkaRegionOk() (*string, bool)`

GetKafkaRegionOk returns a tuple with the KafkaRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaRegion

`func (o *GetConnectv1ConnectorConfig200Response) SetKafkaRegion(v string)`

SetKafkaRegion sets KafkaRegion field to given value.


### GetKafkaApiKey

`func (o *GetConnectv1ConnectorConfig200Response) GetKafkaApiKey() string`

GetKafkaApiKey returns the KafkaApiKey field if non-nil, zero value otherwise.

### GetKafkaApiKeyOk

`func (o *GetConnectv1ConnectorConfig200Response) GetKafkaApiKeyOk() (*string, bool)`

GetKafkaApiKeyOk returns a tuple with the KafkaApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaApiKey

`func (o *GetConnectv1ConnectorConfig200Response) SetKafkaApiKey(v string)`

SetKafkaApiKey sets KafkaApiKey field to given value.


### GetKafkaApiSecret

`func (o *GetConnectv1ConnectorConfig200Response) GetKafkaApiSecret() string`

GetKafkaApiSecret returns the KafkaApiSecret field if non-nil, zero value otherwise.

### GetKafkaApiSecretOk

`func (o *GetConnectv1ConnectorConfig200Response) GetKafkaApiSecretOk() (*string, bool)`

GetKafkaApiSecretOk returns a tuple with the KafkaApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaApiSecret

`func (o *GetConnectv1ConnectorConfig200Response) SetKafkaApiSecret(v string)`

SetKafkaApiSecret sets KafkaApiSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


