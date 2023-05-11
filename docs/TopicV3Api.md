# \TopicV3Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKafkaTopic**](TopicV3Api.md#CreateKafkaTopic) | **Post** /kafka/v3/clusters/{cluster_id}/topics | Create Topic
[**DeleteKafkaTopic**](TopicV3Api.md#DeleteKafkaTopic) | **Delete** /kafka/v3/clusters/{cluster_id}/topics/{topic_name} | Delete Topic
[**GetKafkaTopic**](TopicV3Api.md#GetKafkaTopic) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name} | Get Topic
[**ListKafkaTopics**](TopicV3Api.md#ListKafkaTopics) | **Get** /kafka/v3/clusters/{cluster_id}/topics | List Topics
[**UpdatePartitionCountKafkaTopic**](TopicV3Api.md#UpdatePartitionCountKafkaTopic) | **Patch** /kafka/v3/clusters/{cluster_id}/topics/{topic_name} | Update partition count



## CreateKafkaTopic

> TopicData CreateKafkaTopic(ctx, clusterId).CreateTopicRequestData(createTopicRequestData).Execute()

Create Topic



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    createTopicRequestData := *openapiclient.NewCreateTopicRequestData("TopicName_example") // CreateTopicRequestData | The topic creation request. Note that Confluent Cloud allows only specific replication factor values. Because of that the replication factor field should either be omitted or it should use one of the allowed values (see https://docs.confluent.io/cloud/current/client-apps/optimizing/durability.html). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicV3Api.CreateKafkaTopic(context.Background(), clusterId).CreateTopicRequestData(createTopicRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicV3Api.CreateKafkaTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKafkaTopic`: TopicData
    fmt.Fprintf(os.Stdout, "Response from `TopicV3Api.CreateKafkaTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKafkaTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTopicRequestData** | [**CreateTopicRequestData**](CreateTopicRequestData.md) | The topic creation request. Note that Confluent Cloud allows only specific replication factor values. Because of that the replication factor field should either be omitted or it should use one of the allowed values (see https://docs.confluent.io/cloud/current/client-apps/optimizing/durability.html). | 

### Return type

[**TopicData**](TopicData.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaTopic

> DeleteKafkaTopic(ctx, clusterId, topicName).IncludeAuthorizedOperations(includeAuthorizedOperations).Execute()

Delete Topic



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    topicName := "topic-1" // string | The topic name.
    includeAuthorizedOperations := true // bool | Specify if authorized operations should be included in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TopicV3Api.DeleteKafkaTopic(context.Background(), clusterId, topicName).IncludeAuthorizedOperations(includeAuthorizedOperations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicV3Api.DeleteKafkaTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**topicName** | **string** | The topic name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKafkaTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeAuthorizedOperations** | **bool** | Specify if authorized operations should be included in the response. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaTopic

> TopicData GetKafkaTopic(ctx, clusterId, topicName).IncludeAuthorizedOperations(includeAuthorizedOperations).Execute()

Get Topic



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    topicName := "topic-1" // string | The topic name.
    includeAuthorizedOperations := true // bool | Specify if authorized operations should be included in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicV3Api.GetKafkaTopic(context.Background(), clusterId, topicName).IncludeAuthorizedOperations(includeAuthorizedOperations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicV3Api.GetKafkaTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaTopic`: TopicData
    fmt.Fprintf(os.Stdout, "Response from `TopicV3Api.GetKafkaTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**topicName** | **string** | The topic name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeAuthorizedOperations** | **bool** | Specify if authorized operations should be included in the response. | 

### Return type

[**TopicData**](TopicData.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaTopics

> TopicDataList ListKafkaTopics(ctx, clusterId).Execute()

List Topics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicV3Api.ListKafkaTopics(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicV3Api.ListKafkaTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaTopics`: TopicDataList
    fmt.Fprintf(os.Stdout, "Response from `TopicV3Api.ListKafkaTopics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TopicDataList**](TopicDataList.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartitionCountKafkaTopic

> TopicData UpdatePartitionCountKafkaTopic(ctx, clusterId, topicName).IncludeAuthorizedOperations(includeAuthorizedOperations).UpdatePartitionCountRequestData(updatePartitionCountRequestData).Execute()

Update partition count



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    topicName := "topic-1" // string | The topic name.
    includeAuthorizedOperations := true // bool | Specify if authorized operations should be included in the response. (optional)
    updatePartitionCountRequestData := *openapiclient.NewUpdatePartitionCountRequestData(int32(123)) // UpdatePartitionCountRequestData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicV3Api.UpdatePartitionCountKafkaTopic(context.Background(), clusterId, topicName).IncludeAuthorizedOperations(includeAuthorizedOperations).UpdatePartitionCountRequestData(updatePartitionCountRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicV3Api.UpdatePartitionCountKafkaTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePartitionCountKafkaTopic`: TopicData
    fmt.Fprintf(os.Stdout, "Response from `TopicV3Api.UpdatePartitionCountKafkaTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**topicName** | **string** | The topic name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePartitionCountKafkaTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeAuthorizedOperations** | **bool** | Specify if authorized operations should be included in the response. | 
 **updatePartitionCountRequestData** | [**UpdatePartitionCountRequestData**](UpdatePartitionCountRequestData.md) |  | 

### Return type

[**TopicData**](TopicData.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

