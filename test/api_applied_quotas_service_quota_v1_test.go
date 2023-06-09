/*
Confluent Cloud APIs

Testing AppliedQuotasServiceQuotaV1ApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/shono-io/shono-ccloud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_AppliedQuotasServiceQuotaV1ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppliedQuotasServiceQuotaV1ApiService GetServiceQuotaV1AppliedQuota", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppliedQuotasServiceQuotaV1Api.GetServiceQuotaV1AppliedQuota(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppliedQuotasServiceQuotaV1ApiService ListServiceQuotaV1AppliedQuotas", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AppliedQuotasServiceQuotaV1Api.ListServiceQuotaV1AppliedQuotas(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
