# \JwksIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RefreshIamV2JsonWebKeySet**](JwksIamV2Api.md#RefreshIamV2JsonWebKeySet) | **Patch** /iam/v2/identity-providers/{provider_id}/jwks | Refresh a provider&#39;s JWKS



## RefreshIamV2JsonWebKeySet

> RefreshIamV2JsonWebKeySet200Response RefreshIamV2JsonWebKeySet(ctx, providerId).IamV2Jwks(iamV2Jwks).Execute()

Refresh a provider's JWKS



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
    iamV2Jwks := *openapiclient.NewIamV2Jwks() // IamV2Jwks |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JwksIamV2Api.RefreshIamV2JsonWebKeySet(context.Background(), providerId).IamV2Jwks(iamV2Jwks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JwksIamV2Api.RefreshIamV2JsonWebKeySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshIamV2JsonWebKeySet`: RefreshIamV2JsonWebKeySet200Response
    fmt.Fprintf(os.Stdout, "Response from `JwksIamV2Api.RefreshIamV2JsonWebKeySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** | The Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshIamV2JsonWebKeySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamV2Jwks** | [**IamV2Jwks**](IamV2Jwks.md) |  | 

### Return type

[**RefreshIamV2JsonWebKeySet200Response**](RefreshIamV2JsonWebKeySet200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

