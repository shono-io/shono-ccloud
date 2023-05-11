# CreateOrUpdateConnectv1ConnectorConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorClass** | **string** | The connector class name. E.g. BigQuerySink, GcsSink, etc. | 
**Name** | **string** | Name or alias of the class (plugin) for this connector. | 
**KafkaApiKey** | **string** | The kafka cluster api key. | 
**KafkaApiSecret** | **string** | The kafka cluster api secret key. | 

## Methods

### NewCreateOrUpdateConnectv1ConnectorConfigRequest

`func NewCreateOrUpdateConnectv1ConnectorConfigRequest(connectorClass string, name string, kafkaApiKey string, kafkaApiSecret string, ) *CreateOrUpdateConnectv1ConnectorConfigRequest`

NewCreateOrUpdateConnectv1ConnectorConfigRequest instantiates a new CreateOrUpdateConnectv1ConnectorConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateConnectv1ConnectorConfigRequestWithDefaults

`func NewCreateOrUpdateConnectv1ConnectorConfigRequestWithDefaults() *CreateOrUpdateConnectv1ConnectorConfigRequest`

NewCreateOrUpdateConnectv1ConnectorConfigRequestWithDefaults instantiates a new CreateOrUpdateConnectv1ConnectorConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorClass

`func (o *CreateOrUpdateConnectv1ConnectorConfigRequest) GetConnectorClass() string`

GetConnectorClass returns the ConnectorClass field if non-nil, zero value otherwise.

### GetConnectorClassOk

`func (o *CreateOrUpdateConnectv1ConnectorConfigRequest) GetConnectorClassOk() (*string, bool)`

GetConnectorClassOk returns a tuple with the ConnectorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorClass

`func (o *CreateOrUpdateConnectv1ConnectorConfigRequest) SetConnectorClass(v string)`

SetConnectorClass sets ConnectorClass field to given value.


### GetName

`func (o *CreateOrUpdateConnectv1ConnectorConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateConnectv1ConnectorConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateConnectv1ConnectorConfigRequest) SetName(v string)`

SetName sets Name field to given value.


### GetKafkaApiKey

`func (o *CreateOrUpdateConnectv1ConnectorConfigRequest) GetKafkaApiKey() string`

GetKafkaApiKey returns the KafkaApiKey field if non-nil, zero value otherwise.

### GetKafkaApiKeyOk

`func (o *CreateOrUpdateConnectv1ConnectorConfigRequest) GetKafkaApiKeyOk() (*string, bool)`

GetKafkaApiKeyOk returns a tuple with the KafkaApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaApiKey

`func (o *CreateOrUpdateConnectv1ConnectorConfigRequest) SetKafkaApiKey(v string)`

SetKafkaApiKey sets KafkaApiKey field to given value.


### GetKafkaApiSecret

`func (o *CreateOrUpdateConnectv1ConnectorConfigRequest) GetKafkaApiSecret() string`

GetKafkaApiSecret returns the KafkaApiSecret field if non-nil, zero value otherwise.

### GetKafkaApiSecretOk

`func (o *CreateOrUpdateConnectv1ConnectorConfigRequest) GetKafkaApiSecretOk() (*string, bool)`

GetKafkaApiSecretOk returns a tuple with the KafkaApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaApiSecret

`func (o *CreateOrUpdateConnectv1ConnectorConfigRequest) SetKafkaApiSecret(v string)`

SetKafkaApiSecret sets KafkaApiSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


