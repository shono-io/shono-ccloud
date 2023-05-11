# \ClientQuotasKafkaQuotasV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKafkaQuotasV1ClientQuota**](ClientQuotasKafkaQuotasV1Api.md#CreateKafkaQuotasV1ClientQuota) | **Post** /kafka-quotas/v1/client-quotas | Create a Client Quota
[**DeleteKafkaQuotasV1ClientQuota**](ClientQuotasKafkaQuotasV1Api.md#DeleteKafkaQuotasV1ClientQuota) | **Delete** /kafka-quotas/v1/client-quotas/{id} | Delete a Client Quota
[**GetKafkaQuotasV1ClientQuota**](ClientQuotasKafkaQuotasV1Api.md#GetKafkaQuotasV1ClientQuota) | **Get** /kafka-quotas/v1/client-quotas/{id} | Read a Client Quota
[**ListKafkaQuotasV1ClientQuotas**](ClientQuotasKafkaQuotasV1Api.md#ListKafkaQuotasV1ClientQuotas) | **Get** /kafka-quotas/v1/client-quotas | List of Client Quotas
[**UpdateKafkaQuotasV1ClientQuota**](ClientQuotasKafkaQuotasV1Api.md#UpdateKafkaQuotasV1ClientQuota) | **Patch** /kafka-quotas/v1/client-quotas/{id} | Update a Client Quota



## CreateKafkaQuotasV1ClientQuota

> CreateKafkaQuotasV1ClientQuota202Response CreateKafkaQuotasV1ClientQuota(ctx).CreateKafkaQuotasV1ClientQuotaRequest(createKafkaQuotasV1ClientQuotaRequest).Execute()

Create a Client Quota



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
    createKafkaQuotasV1ClientQuotaRequest := *openapiclient.NewCreateKafkaQuotasV1ClientQuotaRequest(*openapiclient.NewCreateKafkaQuotasV1ClientQuotaRequestAllOf1Spec()) // CreateKafkaQuotasV1ClientQuotaRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientQuotasKafkaQuotasV1Api.CreateKafkaQuotasV1ClientQuota(context.Background()).CreateKafkaQuotasV1ClientQuotaRequest(createKafkaQuotasV1ClientQuotaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientQuotasKafkaQuotasV1Api.CreateKafkaQuotasV1ClientQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKafkaQuotasV1ClientQuota`: CreateKafkaQuotasV1ClientQuota202Response
    fmt.Fprintf(os.Stdout, "Response from `ClientQuotasKafkaQuotasV1Api.CreateKafkaQuotasV1ClientQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKafkaQuotasV1ClientQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKafkaQuotasV1ClientQuotaRequest** | [**CreateKafkaQuotasV1ClientQuotaRequest**](CreateKafkaQuotasV1ClientQuotaRequest.md) |  | 

### Return type

[**CreateKafkaQuotasV1ClientQuota202Response**](CreateKafkaQuotasV1ClientQuota202Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaQuotasV1ClientQuota

> DeleteKafkaQuotasV1ClientQuota(ctx, id).Execute()

Delete a Client Quota



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
    id := "id_example" // string | The unique identifier for the client quota.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ClientQuotasKafkaQuotasV1Api.DeleteKafkaQuotasV1ClientQuota(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientQuotasKafkaQuotasV1Api.DeleteKafkaQuotasV1ClientQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the client quota. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKafkaQuotasV1ClientQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaQuotasV1ClientQuota

> GetKafkaQuotasV1ClientQuota200Response GetKafkaQuotasV1ClientQuota(ctx, id).Execute()

Read a Client Quota



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
    id := "id_example" // string | The unique identifier for the client quota.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientQuotasKafkaQuotasV1Api.GetKafkaQuotasV1ClientQuota(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientQuotasKafkaQuotasV1Api.GetKafkaQuotasV1ClientQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaQuotasV1ClientQuota`: GetKafkaQuotasV1ClientQuota200Response
    fmt.Fprintf(os.Stdout, "Response from `ClientQuotasKafkaQuotasV1Api.GetKafkaQuotasV1ClientQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the client quota. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaQuotasV1ClientQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetKafkaQuotasV1ClientQuota200Response**](GetKafkaQuotasV1ClientQuota200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaQuotasV1ClientQuotas

> ListKafkaQuotasV1ClientQuotas200Response ListKafkaQuotasV1ClientQuotas(ctx).SpecCluster(specCluster).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

List of Client Quotas



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
    specCluster := "lkc-xxxxx" // string | Filter the results by exact match for spec.cluster.
    environment := "env-xxxxx" // string | Filter the results by exact match for environment.
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientQuotasKafkaQuotasV1Api.ListKafkaQuotasV1ClientQuotas(context.Background()).SpecCluster(specCluster).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientQuotasKafkaQuotasV1Api.ListKafkaQuotasV1ClientQuotas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaQuotasV1ClientQuotas`: ListKafkaQuotasV1ClientQuotas200Response
    fmt.Fprintf(os.Stdout, "Response from `ClientQuotasKafkaQuotasV1Api.ListKafkaQuotasV1ClientQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaQuotasV1ClientQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **specCluster** | **string** | Filter the results by exact match for spec.cluster. | 
 **environment** | **string** | Filter the results by exact match for environment. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ListKafkaQuotasV1ClientQuotas200Response**](ListKafkaQuotasV1ClientQuotas200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaQuotasV1ClientQuota

> GetKafkaQuotasV1ClientQuota200Response UpdateKafkaQuotasV1ClientQuota(ctx, id).UpdateKafkaQuotasV1ClientQuotaRequest(updateKafkaQuotasV1ClientQuotaRequest).Execute()

Update a Client Quota



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
    id := "id_example" // string | The unique identifier for the client quota.
    updateKafkaQuotasV1ClientQuotaRequest := *openapiclient.NewUpdateKafkaQuotasV1ClientQuotaRequest(*openapiclient.NewUpdateCmkV2ClusterRequestAllOfSpec(interface{}({"id":"env-00000"}))) // UpdateKafkaQuotasV1ClientQuotaRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientQuotasKafkaQuotasV1Api.UpdateKafkaQuotasV1ClientQuota(context.Background(), id).UpdateKafkaQuotasV1ClientQuotaRequest(updateKafkaQuotasV1ClientQuotaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientQuotasKafkaQuotasV1Api.UpdateKafkaQuotasV1ClientQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKafkaQuotasV1ClientQuota`: GetKafkaQuotasV1ClientQuota200Response
    fmt.Fprintf(os.Stdout, "Response from `ClientQuotasKafkaQuotasV1Api.UpdateKafkaQuotasV1ClientQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the client quota. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaQuotasV1ClientQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateKafkaQuotasV1ClientQuotaRequest** | [**UpdateKafkaQuotasV1ClientQuotaRequest**](UpdateKafkaQuotasV1ClientQuotaRequest.md) |  | 

### Return type

[**GetKafkaQuotasV1ClientQuota200Response**](GetKafkaQuotasV1ClientQuota200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

