# ConnectV1ConnectorTasksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | **string** | The name of the connector the task belongs to | 
**Task** | **int32** | Task ID within the connector | 

## Methods

### NewConnectV1ConnectorTasksInner

`func NewConnectV1ConnectorTasksInner(connector string, task int32, ) *ConnectV1ConnectorTasksInner`

NewConnectV1ConnectorTasksInner instantiates a new ConnectV1ConnectorTasksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1ConnectorTasksInnerWithDefaults

`func NewConnectV1ConnectorTasksInnerWithDefaults() *ConnectV1ConnectorTasksInner`

NewConnectV1ConnectorTasksInnerWithDefaults instantiates a new ConnectV1ConnectorTasksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *ConnectV1ConnectorTasksInner) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ConnectV1ConnectorTasksInner) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ConnectV1ConnectorTasksInner) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetTask

`func (o *ConnectV1ConnectorTasksInner) GetTask() int32`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ConnectV1ConnectorTasksInner) GetTaskOk() (*int32, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ConnectV1ConnectorTasksInner) SetTask(v int32)`

SetTask sets Task field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


