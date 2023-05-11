# \IdentityPoolsIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIamV2IdentityPool**](IdentityPoolsIamV2Api.md#CreateIamV2IdentityPool) | **Post** /iam/v2/identity-providers/{provider_id}/identity-pools | Create an Identity Pool
[**DeleteIamV2IdentityPool**](IdentityPoolsIamV2Api.md#DeleteIamV2IdentityPool) | **Delete** /iam/v2/identity-providers/{provider_id}/identity-pools/{id} | Delete an Identity Pool
[**GetIamV2IdentityPool**](IdentityPoolsIamV2Api.md#GetIamV2IdentityPool) | **Get** /iam/v2/identity-providers/{provider_id}/identity-pools/{id} | Read an Identity Pool
[**ListIamV2IdentityPools**](IdentityPoolsIamV2Api.md#ListIamV2IdentityPools) | **Get** /iam/v2/identity-providers/{provider_id}/identity-pools | List of Identity Pools
[**UpdateIamV2IdentityPool**](IdentityPoolsIamV2Api.md#UpdateIamV2IdentityPool) | **Patch** /iam/v2/identity-providers/{provider_id}/identity-pools/{id} | Update an Identity Pool



## CreateIamV2IdentityPool

> CreateIamV2IdentityPoolRequest CreateIamV2IdentityPool(ctx, providerId).CreateIamV2IdentityPoolRequest(createIamV2IdentityPoolRequest).Execute()

Create an Identity Pool



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
    providerId := "providerId_example" // string | The Provider
    createIamV2IdentityPoolRequest := *openapiclient.NewCreateIamV2IdentityPoolRequest("My Identity Pool", "Prod Access to Kafka clusters to Release Engineering", "claims.sub", "claims.aud=="confluent" && claims.group!="invalid_group"") // CreateIamV2IdentityPoolRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPoolsIamV2Api.CreateIamV2IdentityPool(context.Background(), providerId).CreateIamV2IdentityPoolRequest(createIamV2IdentityPoolRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPoolsIamV2Api.CreateIamV2IdentityPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamV2IdentityPool`: CreateIamV2IdentityPoolRequest
    fmt.Fprintf(os.Stdout, "Response from `IdentityPoolsIamV2Api.CreateIamV2IdentityPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** | The Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamV2IdentityPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIamV2IdentityPoolRequest** | [**CreateIamV2IdentityPoolRequest**](CreateIamV2IdentityPoolRequest.md) |  | 

### Return type

[**CreateIamV2IdentityPoolRequest**](CreateIamV2IdentityPoolRequest.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamV2IdentityPool

> DeleteIamV2IdentityPool(ctx, providerId, id).Execute()

Delete an Identity Pool



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
    providerId := "providerId_example" // string | The Provider
    id := "id_example" // string | The unique identifier for the identity pool.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdentityPoolsIamV2Api.DeleteIamV2IdentityPool(context.Background(), providerId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPoolsIamV2Api.DeleteIamV2IdentityPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** | The Provider | 
**id** | **string** | The unique identifier for the identity pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamV2IdentityPoolRequest struct via the builder pattern


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


## GetIamV2IdentityPool

> GetIamV2IdentityPool200Response GetIamV2IdentityPool(ctx, providerId, id).Execute()

Read an Identity Pool



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
    providerId := "providerId_example" // string | The Provider
    id := "id_example" // string | The unique identifier for the identity pool.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPoolsIamV2Api.GetIamV2IdentityPool(context.Background(), providerId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPoolsIamV2Api.GetIamV2IdentityPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2IdentityPool`: GetIamV2IdentityPool200Response
    fmt.Fprintf(os.Stdout, "Response from `IdentityPoolsIamV2Api.GetIamV2IdentityPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** | The Provider | 
**id** | **string** | The unique identifier for the identity pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2IdentityPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetIamV2IdentityPool200Response**](GetIamV2IdentityPool200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIamV2IdentityPools

> ListIamV2IdentityPools200Response ListIamV2IdentityPools(ctx, providerId).PageSize(pageSize).PageToken(pageToken).Execute()

List of Identity Pools



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
    providerId := "providerId_example" // string | The Provider
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPoolsIamV2Api.ListIamV2IdentityPools(context.Background(), providerId).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPoolsIamV2Api.ListIamV2IdentityPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIamV2IdentityPools`: ListIamV2IdentityPools200Response
    fmt.Fprintf(os.Stdout, "Response from `IdentityPoolsIamV2Api.ListIamV2IdentityPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** | The Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIamV2IdentityPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ListIamV2IdentityPools200Response**](ListIamV2IdentityPools200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIamV2IdentityPool

> GetIamV2IdentityPool200Response UpdateIamV2IdentityPool(ctx, providerId, id).IamV2IdentityPool(iamV2IdentityPool).Execute()

Update an Identity Pool



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
    providerId := "providerId_example" // string | The Provider
    id := "id_example" // string | The unique identifier for the identity pool.
    iamV2IdentityPool := *openapiclient.NewIamV2IdentityPool() // IamV2IdentityPool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPoolsIamV2Api.UpdateIamV2IdentityPool(context.Background(), providerId, id).IamV2IdentityPool(iamV2IdentityPool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPoolsIamV2Api.UpdateIamV2IdentityPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIamV2IdentityPool`: GetIamV2IdentityPool200Response
    fmt.Fprintf(os.Stdout, "Response from `IdentityPoolsIamV2Api.UpdateIamV2IdentityPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** | The Provider | 
**id** | **string** | The unique identifier for the identity pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIamV2IdentityPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamV2IdentityPool** | [**IamV2IdentityPool**](IamV2IdentityPool.md) |  | 

### Return type

[**GetIamV2IdentityPool200Response**](GetIamV2IdentityPool200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

