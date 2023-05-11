# \OrganizationsPartnerV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPartnerV2Organization**](OrganizationsPartnerV2Api.md#GetPartnerV2Organization) | **Get** /partner/v2/organizations/{id} | Read an Organization
[**ListPartnerV2Organizations**](OrganizationsPartnerV2Api.md#ListPartnerV2Organizations) | **Get** /partner/v2/organizations | List of Organizations



## GetPartnerV2Organization

> GetPartnerV2Organization200Response GetPartnerV2Organization(ctx, id).Execute()

Read an Organization



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
    id := "id_example" // string | The unique identifier for the organization.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsPartnerV2Api.GetPartnerV2Organization(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsPartnerV2Api.GetPartnerV2Organization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartnerV2Organization`: GetPartnerV2Organization200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsPartnerV2Api.GetPartnerV2Organization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnerV2OrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPartnerV2Organization200Response**](GetPartnerV2Organization200Response.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPartnerV2Organizations

> PartnerV2OrganizationList ListPartnerV2Organizations(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Organizations



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
    resp, r, err := apiClient.OrganizationsPartnerV2Api.ListPartnerV2Organizations(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsPartnerV2Api.ListPartnerV2Organizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPartnerV2Organizations`: PartnerV2OrganizationList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsPartnerV2Api.ListPartnerV2Organizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPartnerV2OrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**PartnerV2OrganizationList**](PartnerV2OrganizationList.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

