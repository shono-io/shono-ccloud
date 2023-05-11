# \SignupPartnerV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSignup**](SignupPartnerV2Api.md#ActivateSignup) | **Post** /partner/v2/signup/activate | Activate an Incomplete Signup
[**Signup**](SignupPartnerV2Api.md#Signup) | **Post** /partner/v2/signup | Signup an Organization on behalf of a Customer
[**SignupPartnerV2Link**](SignupPartnerV2Api.md#SignupPartnerV2Link) | **Post** /partner/v2/signup/link | Signup a Customer by Linking to an Existing Organization



## ActivateSignup

> PartnerSignupResponse ActivateSignup(ctx).ActivatePartnerSignupRequest(activatePartnerSignupRequest).Execute()

Activate an Incomplete Signup



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
    activatePartnerSignupRequest := *openapiclient.NewActivatePartnerSignupRequest(interface{}(123), "b3a17773-05cc-4431-9560-433fb4613da8") // ActivatePartnerSignupRequest | A JSON object containing signup information (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignupPartnerV2Api.ActivateSignup(context.Background()).ActivatePartnerSignupRequest(activatePartnerSignupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignupPartnerV2Api.ActivateSignup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateSignup`: PartnerSignupResponse
    fmt.Fprintf(os.Stdout, "Response from `SignupPartnerV2Api.ActivateSignup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivateSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activatePartnerSignupRequest** | [**ActivatePartnerSignupRequest**](ActivatePartnerSignupRequest.md) | A JSON object containing signup information | 

### Return type

[**PartnerSignupResponse**](PartnerSignupResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Signup

> PartnerSignupResponse Signup(ctx).DryRun(dryRun).PartnerSignupRequest(partnerSignupRequest).Execute()

Signup an Organization on behalf of a Customer



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
    dryRun := true // bool | If true, only perform validation of signup (optional)
    partnerSignupRequest := *openapiclient.NewPartnerSignupRequest(*openapiclient.NewPartnerSignupRequestOrganization("Acme Organization", openapiclient.partner_v2_Organization_sso_config{AzureSSOConfig: openapiclient.NewAzureSSOConfig("AzureSSOConfig", "b3a17773-05cc-4431-9560-433fb4613da8")}), openapiclient.PartnerSignupRequest_entitlement{PartnerSignupRequestEntitlementOneOf: openapiclient.NewPartnerSignupRequestEntitlementOneOf("1111-2222-3333-4444", "Acme Prod Entitlement", "confluent-cloud-payg-prod", "confluent-cloud-kafka-service-azure")}) // PartnerSignupRequest | A JSON object containing signup information (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignupPartnerV2Api.Signup(context.Background()).DryRun(dryRun).PartnerSignupRequest(partnerSignupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignupPartnerV2Api.Signup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Signup`: PartnerSignupResponse
    fmt.Fprintf(os.Stdout, "Response from `SignupPartnerV2Api.Signup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dryRun** | **bool** | If true, only perform validation of signup | 
 **partnerSignupRequest** | [**PartnerSignupRequest**](PartnerSignupRequest.md) | A JSON object containing signup information | 

### Return type

[**PartnerSignupResponse**](PartnerSignupResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignupPartnerV2Link

> PartnerSignupResponse SignupPartnerV2Link(ctx).DryRun(dryRun).PartnerLinkRequest(partnerLinkRequest).Execute()

Signup a Customer by Linking to an Existing Organization



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
    dryRun := true // bool | If true, only perform validation of signup (optional)
    partnerLinkRequest := *openapiclient.NewPartnerLinkRequest("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", *openapiclient.NewPartnerLinkRequestOrganization(openapiclient.partner_v2_Organization_sso_config{AzureSSOConfig: openapiclient.NewAzureSSOConfig("AzureSSOConfig", "b3a17773-05cc-4431-9560-433fb4613da8")}), openapiclient.PartnerSignupRequest_entitlement{PartnerSignupRequestEntitlementOneOf: openapiclient.NewPartnerSignupRequestEntitlementOneOf("1111-2222-3333-4444", "Acme Prod Entitlement", "confluent-cloud-payg-prod", "confluent-cloud-kafka-service-azure")}) // PartnerLinkRequest | A JSON object containing signup information (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignupPartnerV2Api.SignupPartnerV2Link(context.Background()).DryRun(dryRun).PartnerLinkRequest(partnerLinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignupPartnerV2Api.SignupPartnerV2Link``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignupPartnerV2Link`: PartnerSignupResponse
    fmt.Fprintf(os.Stdout, "Response from `SignupPartnerV2Api.SignupPartnerV2Link`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupPartnerV2LinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dryRun** | **bool** | If true, only perform validation of signup | 
 **partnerLinkRequest** | [**PartnerLinkRequest**](PartnerLinkRequest.md) | A JSON object containing signup information | 

### Return type

[**PartnerSignupResponse**](PartnerSignupResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

