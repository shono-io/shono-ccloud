# CreateConnectv1ConnectorRequestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorClass** | **string** | The connector class name. E.g. BigQuerySink, GcsSink, etc. | 
**Name** | **string** | Name or alias of the class (plugin) for this connector. | 
**KafkaApiKey** | **string** | The kafka cluster api key. | 
**KafkaApiSecret** | **string** | The kafka cluster api secret key. | 

## Methods

### NewCreateConnectv1ConnectorRequestConfig

`func NewCreateConnectv1ConnectorRequestConfig(connectorClass string, name string, kafkaApiKey string, kafkaApiSecret string, ) *CreateConnectv1ConnectorRequestConfig`

NewCreateConnectv1ConnectorRequestConfig instantiates a new CreateConnectv1ConnectorRequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConnectv1ConnectorRequestConfigWithDefaults

`func NewCreateConnectv1ConnectorRequestConfigWithDefaults() *CreateConnectv1ConnectorRequestConfig`

NewCreateConnectv1ConnectorRequestConfigWithDefaults instantiates a new CreateConnectv1ConnectorRequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorClass

`func (o *CreateConnectv1ConnectorRequestConfig) GetConnectorClass() string`

GetConnectorClass returns the ConnectorClass field if non-nil, zero value otherwise.

### GetConnectorClassOk

`func (o *CreateConnectv1ConnectorRequestConfig) GetConnectorClassOk() (*string, bool)`

GetConnectorClassOk returns a tuple with the ConnectorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorClass

`func (o *CreateConnectv1ConnectorRequestConfig) SetConnectorClass(v string)`

SetConnectorClass sets ConnectorClass field to given value.


### GetName

`func (o *CreateConnectv1ConnectorRequestConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateConnectv1ConnectorRequestConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateConnectv1ConnectorRequestConfig) SetName(v string)`

SetName sets Name field to given value.


### GetKafkaApiKey

`func (o *CreateConnectv1ConnectorRequestConfig) GetKafkaApiKey() string`

GetKafkaApiKey returns the KafkaApiKey field if non-nil, zero value otherwise.

### GetKafkaApiKeyOk

`func (o *CreateConnectv1ConnectorRequestConfig) GetKafkaApiKeyOk() (*string, bool)`

GetKafkaApiKeyOk returns a tuple with the KafkaApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaApiKey

`func (o *CreateConnectv1ConnectorRequestConfig) SetKafkaApiKey(v string)`

SetKafkaApiKey sets KafkaApiKey field to given value.


### GetKafkaApiSecret

`func (o *CreateConnectv1ConnectorRequestConfig) GetKafkaApiSecret() string`

GetKafkaApiSecret returns the KafkaApiSecret field if non-nil, zero value otherwise.

### GetKafkaApiSecretOk

`func (o *CreateConnectv1ConnectorRequestConfig) GetKafkaApiSecretOk() (*string, bool)`

GetKafkaApiSecretOk returns a tuple with the KafkaApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaApiSecret

`func (o *CreateConnectv1ConnectorRequestConfig) SetKafkaApiSecret(v string)`

SetKafkaApiSecret sets KafkaApiSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


