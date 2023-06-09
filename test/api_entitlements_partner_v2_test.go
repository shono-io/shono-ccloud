/*
Confluent Cloud APIs

Testing EntitlementsPartnerV2ApiService

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

func Test_openapi_EntitlementsPartnerV2ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EntitlementsPartnerV2ApiService CreatePartnerV2Entitlement", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EntitlementsPartnerV2Api.CreatePartnerV2Entitlement(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsPartnerV2ApiService GetPartnerV2Entitlement", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.EntitlementsPartnerV2Api.GetPartnerV2Entitlement(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsPartnerV2ApiService ListPartnerV2Entitlements", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EntitlementsPartnerV2Api.ListPartnerV2Entitlements(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
