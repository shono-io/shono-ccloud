# ReadConnectv1ConnectorStatus200ResponseTasksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of task. | 
**State** | **string** | The state of the task. | 
**WorkerId** | **string** | The worker ID of the task. | 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewReadConnectv1ConnectorStatus200ResponseTasksInner

`func NewReadConnectv1ConnectorStatus200ResponseTasksInner(id int32, state string, workerId string, ) *ReadConnectv1ConnectorStatus200ResponseTasksInner`

NewReadConnectv1ConnectorStatus200ResponseTasksInner instantiates a new ReadConnectv1ConnectorStatus200ResponseTasksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadConnectv1ConnectorStatus200ResponseTasksInnerWithDefaults

`func NewReadConnectv1ConnectorStatus200ResponseTasksInnerWithDefaults() *ReadConnectv1ConnectorStatus200ResponseTasksInner`

NewReadConnectv1ConnectorStatus200ResponseTasksInnerWithDefaults instantiates a new ReadConnectv1ConnectorStatus200ResponseTasksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReadConnectv1ConnectorStatus200ResponseTasksInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReadConnectv1ConnectorStatus200ResponseTasksInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReadConnectv1ConnectorStatus200ResponseTasksInner) SetId(v int32)`

SetId sets Id field to given value.


### GetState

`func (o *ReadConnectv1ConnectorStatus200ResponseTasksInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReadConnectv1ConnectorStatus200ResponseTasksInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReadConnectv1ConnectorStatus200ResponseTasksInner) SetState(v string)`

SetState sets State field to given value.


### GetWorkerId

`func (o *ReadConnectv1ConnectorStatus200ResponseTasksInner) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *ReadConnectv1ConnectorStatus200ResponseTasksInner) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *ReadConnectv1ConnectorStatus200ResponseTasksInner) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.


### GetMsg

`func (o *ReadConnectv1ConnectorStatus200ResponseTasksInner) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ReadConnectv1ConnectorStatus200ResponseTasksInner) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ReadConnectv1ConnectorStatus200ResponseTasksInner) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ReadConnectv1ConnectorStatus200ResponseTasksInner) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


