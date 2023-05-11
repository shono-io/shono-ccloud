# \ScopesServiceQuotaV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceQuotaV1Scope**](ScopesServiceQuotaV1Api.md#GetServiceQuotaV1Scope) | **Get** /service-quota/v1/scopes/{id} | Read a Scope
[**ListServiceQuotaV1Scopes**](ScopesServiceQuotaV1Api.md#ListServiceQuotaV1Scopes) | **Get** /service-quota/v1/scopes | List of Scopes



## GetServiceQuotaV1Scope

> GetServiceQuotaV1Scope200Response GetServiceQuotaV1Scope(ctx, id).Execute()

Read a Scope



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
    id := "id_example" // string | The unique identifier for the scope.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScopesServiceQuotaV1Api.GetServiceQuotaV1Scope(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopesServiceQuotaV1Api.GetServiceQuotaV1Scope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceQuotaV1Scope`: GetServiceQuotaV1Scope200Response
    fmt.Fprintf(os.Stdout, "Response from `ScopesServiceQuotaV1Api.GetServiceQuotaV1Scope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceQuotaV1ScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServiceQuotaV1Scope200Response**](GetServiceQuotaV1Scope200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceQuotaV1Scopes

> ListServiceQuotaV1Scopes200Response ListServiceQuotaV1Scopes(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Scopes



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
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScopesServiceQuotaV1Api.ListServiceQuotaV1Scopes(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopesServiceQuotaV1Api.ListServiceQuotaV1Scopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceQuotaV1Scopes`: ListServiceQuotaV1Scopes200Response
    fmt.Fprintf(os.Stdout, "Response from `ScopesServiceQuotaV1Api.ListServiceQuotaV1Scopes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServiceQuotaV1ScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ListServiceQuotaV1Scopes200Response**](ListServiceQuotaV1Scopes200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

