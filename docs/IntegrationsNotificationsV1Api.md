# \IntegrationsNotificationsV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationsV1Integration**](IntegrationsNotificationsV1Api.md#CreateNotificationsV1Integration) | **Post** /notifications/v1/integrations | Create an Integration
[**DeleteNotificationsV1Integration**](IntegrationsNotificationsV1Api.md#DeleteNotificationsV1Integration) | **Delete** /notifications/v1/integrations/{id} | Delete an Integration
[**GetNotificationsV1Integration**](IntegrationsNotificationsV1Api.md#GetNotificationsV1Integration) | **Get** /notifications/v1/integrations/{id} | Read an Integration
[**ListNotificationsV1Integrations**](IntegrationsNotificationsV1Api.md#ListNotificationsV1Integrations) | **Get** /notifications/v1/integrations | List of Integrations
[**TestNotificationsV1Integration**](IntegrationsNotificationsV1Api.md#TestNotificationsV1Integration) | **Post** /notifications/v1/integrations:test | Test a Webhook, Slack or Microsoft Teams integration
[**UpdateNotificationsV1Integration**](IntegrationsNotificationsV1Api.md#UpdateNotificationsV1Integration) | **Patch** /notifications/v1/integrations/{id} | Update an Integration



## CreateNotificationsV1Integration

> CreateNotificationsV1IntegrationRequest CreateNotificationsV1Integration(ctx).CreateNotificationsV1IntegrationRequest(createNotificationsV1IntegrationRequest).Execute()

Create an Integration



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
    createNotificationsV1IntegrationRequest := *openapiclient.NewCreateNotificationsV1IntegrationRequest("Slack integration", *openapiclient.NewNotificationsV1IntegrationTarget("MsTeams", "https://admin.webhook.office.com/webhookb2/{id}/IncomingWebhook/{id}", "OrganizationAdmin", "u-temp1", "https://my.webhook.url/{id}")) // CreateNotificationsV1IntegrationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsNotificationsV1Api.CreateNotificationsV1Integration(context.Background()).CreateNotificationsV1IntegrationRequest(createNotificationsV1IntegrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsNotificationsV1Api.CreateNotificationsV1Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsV1Integration`: CreateNotificationsV1IntegrationRequest
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsNotificationsV1Api.CreateNotificationsV1Integration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsV1IntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNotificationsV1IntegrationRequest** | [**CreateNotificationsV1IntegrationRequest**](CreateNotificationsV1IntegrationRequest.md) |  | 

### Return type

[**CreateNotificationsV1IntegrationRequest**](CreateNotificationsV1IntegrationRequest.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationsV1Integration

> DeleteNotificationsV1Integration(ctx, id).Execute()

Delete an Integration



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
    id := "id_example" // string | The unique identifier for the integration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationsNotificationsV1Api.DeleteNotificationsV1Integration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsNotificationsV1Api.DeleteNotificationsV1Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationsV1IntegrationRequest struct via the builder pattern


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


## GetNotificationsV1Integration

> GetNotificationsV1Integration200Response GetNotificationsV1Integration(ctx, id).Execute()

Read an Integration



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
    id := "id_example" // string | The unique identifier for the integration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsNotificationsV1Api.GetNotificationsV1Integration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsNotificationsV1Api.GetNotificationsV1Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsV1Integration`: GetNotificationsV1Integration200Response
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsNotificationsV1Api.GetNotificationsV1Integration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsV1IntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNotificationsV1Integration200Response**](GetNotificationsV1Integration200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationsV1Integrations

> ListNotificationsV1Integrations200Response ListNotificationsV1Integrations(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Integrations



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
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 100)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsNotificationsV1Api.ListNotificationsV1Integrations(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsNotificationsV1Api.ListNotificationsV1Integrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotificationsV1Integrations`: ListNotificationsV1Integrations200Response
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsNotificationsV1Api.ListNotificationsV1Integrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationsV1IntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 100]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ListNotificationsV1Integrations200Response**](ListNotificationsV1Integrations200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestNotificationsV1Integration

> TestNotificationsV1Integration(ctx).CreateNotificationsV1IntegrationRequest(createNotificationsV1IntegrationRequest).Execute()

Test a Webhook, Slack or Microsoft Teams integration



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
    createNotificationsV1IntegrationRequest := *openapiclient.NewCreateNotificationsV1IntegrationRequest("Slack integration", *openapiclient.NewNotificationsV1IntegrationTarget("MsTeams", "https://admin.webhook.office.com/webhookb2/{id}/IncomingWebhook/{id}", "OrganizationAdmin", "u-temp1", "https://my.webhook.url/{id}")) // CreateNotificationsV1IntegrationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationsNotificationsV1Api.TestNotificationsV1Integration(context.Background()).CreateNotificationsV1IntegrationRequest(createNotificationsV1IntegrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsNotificationsV1Api.TestNotificationsV1Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestNotificationsV1IntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNotificationsV1IntegrationRequest** | [**CreateNotificationsV1IntegrationRequest**](CreateNotificationsV1IntegrationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationsV1Integration

> GetNotificationsV1Integration200Response UpdateNotificationsV1Integration(ctx, id).NotificationsV1Integration(notificationsV1Integration).Execute()

Update an Integration



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
    id := "id_example" // string | The unique identifier for the integration.
    notificationsV1Integration := *openapiclient.NewNotificationsV1Integration() // NotificationsV1Integration |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsNotificationsV1Api.UpdateNotificationsV1Integration(context.Background(), id).NotificationsV1Integration(notificationsV1Integration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsNotificationsV1Api.UpdateNotificationsV1Integration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationsV1Integration`: GetNotificationsV1Integration200Response
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsNotificationsV1Api.UpdateNotificationsV1Integration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationsV1IntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationsV1Integration** | [**NotificationsV1Integration**](NotificationsV1Integration.md) |  | 

### Return type

[**GetNotificationsV1Integration200Response**](GetNotificationsV1Integration200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

