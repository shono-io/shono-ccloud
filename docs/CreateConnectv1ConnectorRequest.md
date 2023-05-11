# CreateConnectv1ConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the connector to create. | [optional] 
**Config** | Pointer to [**CreateConnectv1ConnectorRequestConfig**](CreateConnectv1ConnectorRequestConfig.md) |  | [optional] 

## Methods

### NewCreateConnectv1ConnectorRequest

`func NewCreateConnectv1ConnectorRequest() *CreateConnectv1ConnectorRequest`

NewCreateConnectv1ConnectorRequest instantiates a new CreateConnectv1ConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConnectv1ConnectorRequestWithDefaults

`func NewCreateConnectv1ConnectorRequestWithDefaults() *CreateConnectv1ConnectorRequest`

NewCreateConnectv1ConnectorRequestWithDefaults instantiates a new CreateConnectv1ConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateConnectv1ConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateConnectv1ConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateConnectv1ConnectorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateConnectv1ConnectorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *CreateConnectv1ConnectorRequest) GetConfig() CreateConnectv1ConnectorRequestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateConnectv1ConnectorRequest) GetConfigOk() (*CreateConnectv1ConnectorRequestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateConnectv1ConnectorRequest) SetConfig(v CreateConnectv1ConnectorRequestConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateConnectv1ConnectorRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


