# ReadConnectv1ConnectorStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the connector. | 
**Type** | **string** | Type of connector, sink or source. | 
**Connector** | [**ReadConnectv1ConnectorStatus200ResponseConnector**](ReadConnectv1ConnectorStatus200ResponseConnector.md) |  | 
**Tasks** | Pointer to [**[]ReadConnectv1ConnectorStatus200ResponseTasksInner**](ReadConnectv1ConnectorStatus200ResponseTasksInner.md) | The map containing the task status. | [optional] 

## Methods

### NewReadConnectv1ConnectorStatus200Response

`func NewReadConnectv1ConnectorStatus200Response(name string, type_ string, connector ReadConnectv1ConnectorStatus200ResponseConnector, ) *ReadConnectv1ConnectorStatus200Response`

NewReadConnectv1ConnectorStatus200Response instantiates a new ReadConnectv1ConnectorStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadConnectv1ConnectorStatus200ResponseWithDefaults

`func NewReadConnectv1ConnectorStatus200ResponseWithDefaults() *ReadConnectv1ConnectorStatus200Response`

NewReadConnectv1ConnectorStatus200ResponseWithDefaults instantiates a new ReadConnectv1ConnectorStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReadConnectv1ConnectorStatus200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReadConnectv1ConnectorStatus200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReadConnectv1ConnectorStatus200Response) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ReadConnectv1ConnectorStatus200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReadConnectv1ConnectorStatus200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReadConnectv1ConnectorStatus200Response) SetType(v string)`

SetType sets Type field to given value.


### GetConnector

`func (o *ReadConnectv1ConnectorStatus200Response) GetConnector() ReadConnectv1ConnectorStatus200ResponseConnector`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ReadConnectv1ConnectorStatus200Response) GetConnectorOk() (*ReadConnectv1ConnectorStatus200ResponseConnector, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ReadConnectv1ConnectorStatus200Response) SetConnector(v ReadConnectv1ConnectorStatus200ResponseConnector)`

SetConnector sets Connector field to given value.


### GetTasks

`func (o *ReadConnectv1ConnectorStatus200Response) GetTasks() []ReadConnectv1ConnectorStatus200ResponseTasksInner`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ReadConnectv1ConnectorStatus200Response) GetTasksOk() (*[]ReadConnectv1ConnectorStatus200ResponseTasksInner, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ReadConnectv1ConnectorStatus200Response) SetTasks(v []ReadConnectv1ConnectorStatus200ResponseTasksInner)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ReadConnectv1ConnectorStatus200Response) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


