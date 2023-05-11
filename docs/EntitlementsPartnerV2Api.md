# \EntitlementsPartnerV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePartnerV2Entitlement**](EntitlementsPartnerV2Api.md#CreatePartnerV2Entitlement) | **Post** /partner/v2/entitlements | Create an Entitlement
[**GetPartnerV2Entitlement**](EntitlementsPartnerV2Api.md#GetPartnerV2Entitlement) | **Get** /partner/v2/entitlements/{id} | Read an Entitlement
[**ListPartnerV2Entitlements**](EntitlementsPartnerV2Api.md#ListPartnerV2Entitlements) | **Get** /partner/v2/entitlements | List of Entitlements



## CreatePartnerV2Entitlement

> CreatePartnerV2EntitlementRequest CreatePartnerV2Entitlement(ctx).CreatePartnerV2EntitlementRequest(createPartnerV2EntitlementRequest).Execute()

Create an Entitlement



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
    createPartnerV2EntitlementRequest := *openapiclient.NewCreatePartnerV2EntitlementRequest("1111-2222-3333-4444", "Acme Prod Entitlement", "confluent-cloud-payg-prod", "confluent-cloud-kafka-service-azure") // CreatePartnerV2EntitlementRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsPartnerV2Api.CreatePartnerV2Entitlement(context.Background()).CreatePartnerV2EntitlementRequest(createPartnerV2EntitlementRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsPartnerV2Api.CreatePartnerV2Entitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePartnerV2Entitlement`: CreatePartnerV2EntitlementRequest
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsPartnerV2Api.CreatePartnerV2Entitlement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePartnerV2EntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPartnerV2EntitlementRequest** | [**CreatePartnerV2EntitlementRequest**](CreatePartnerV2EntitlementRequest.md) |  | 

### Return type

[**CreatePartnerV2EntitlementRequest**](CreatePartnerV2EntitlementRequest.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartnerV2Entitlement

> CreatePartnerV2EntitlementRequest GetPartnerV2Entitlement(ctx, id).OrganizationId(organizationId).Execute()

Read an Entitlement



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
    id := "id_example" // string | The unique identifier for the entitlement.
    organizationId := "b3a17773-05cc-4431-9560-433fb4613da8" // string | Scope the operation to the given organization.id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsPartnerV2Api.GetPartnerV2Entitlement(context.Background(), id).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsPartnerV2Api.GetPartnerV2Entitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartnerV2Entitlement`: CreatePartnerV2EntitlementRequest
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsPartnerV2Api.GetPartnerV2Entitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the entitlement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnerV2EntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationId** | **string** | Scope the operation to the given organization.id. | 

### Return type

[**CreatePartnerV2EntitlementRequest**](CreatePartnerV2EntitlementRequest.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPartnerV2Entitlements

> PartnerV2EntitlementList ListPartnerV2Entitlements(ctx).OrganizationId(organizationId).PageSize(pageSize).PageToken(pageToken).Execute()

List of Entitlements



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
    organizationId := "b3a17773-05cc-4431-9560-433fb4613da8" // string | Filter the results by exact match for organization.id. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsPartnerV2Api.ListPartnerV2Entitlements(context.Background()).OrganizationId(organizationId).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsPartnerV2Api.ListPartnerV2Entitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPartnerV2Entitlements`: PartnerV2EntitlementList
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsPartnerV2Api.ListPartnerV2Entitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPartnerV2EntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **string** | Filter the results by exact match for organization.id. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**PartnerV2EntitlementList**](PartnerV2EntitlementList.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

