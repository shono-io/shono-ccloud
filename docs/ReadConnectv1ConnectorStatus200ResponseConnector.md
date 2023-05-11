# ReadConnectv1ConnectorStatus200ResponseConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The state of the connector. | 
**WorkerId** | **string** | The worker ID of the connector. | 
**Trace** | Pointer to **string** | The exception name in case of error. | [optional] 

## Methods

### NewReadConnectv1ConnectorStatus200ResponseConnector

`func NewReadConnectv1ConnectorStatus200ResponseConnector(state string, workerId string, ) *ReadConnectv1ConnectorStatus200ResponseConnector`

NewReadConnectv1ConnectorStatus200ResponseConnector instantiates a new ReadConnectv1ConnectorStatus200ResponseConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadConnectv1ConnectorStatus200ResponseConnectorWithDefaults

`func NewReadConnectv1ConnectorStatus200ResponseConnectorWithDefaults() *ReadConnectv1ConnectorStatus200ResponseConnector`

NewReadConnectv1ConnectorStatus200ResponseConnectorWithDefaults instantiates a new ReadConnectv1ConnectorStatus200ResponseConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ReadConnectv1ConnectorStatus200ResponseConnector) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReadConnectv1ConnectorStatus200ResponseConnector) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReadConnectv1ConnectorStatus200ResponseConnector) SetState(v string)`

SetState sets State field to given value.


### GetWorkerId

`func (o *ReadConnectv1ConnectorStatus200ResponseConnector) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *ReadConnectv1ConnectorStatus200ResponseConnector) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *ReadConnectv1ConnectorStatus200ResponseConnector) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.


### GetTrace

`func (o *ReadConnectv1ConnectorStatus200ResponseConnector) GetTrace() string`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *ReadConnectv1ConnectorStatus200ResponseConnector) GetTraceOk() (*string, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *ReadConnectv1ConnectorStatus200ResponseConnector) SetTrace(v string)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *ReadConnectv1ConnectorStatus200ResponseConnector) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


